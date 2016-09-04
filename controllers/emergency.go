package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kr/pretty"
	"github.com/nu7hatch/gouuid"
	"github.com/sfreiberg/gotwilio"

	"../data"
	"../models"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
)

var twilio *gotwilio.Twilio

func init() {
	accountSid := "AC30009862e9b061e21fcde75bc13cb49b"
	authToken := "95b081654eb589a8397ae216067502c8"
	twilio = gotwilio.NewTwilioClient(accountSid, authToken)
}

var emergencies = make(map[int]*models.Emergency)

var uuidToEmergency = make(map[string]*models.Emergency)
var uuidToUser = make(map[string]*models.User)

func EmergencyAll(w http.ResponseWriter, r *http.Request) {
	keys := make([]int, 0)
	for k := range emergencies {
		keys = append(keys, k)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println(keys)
	if err := json.NewEncoder(w).Encode(pretty.Sprint(keys)); err != nil {
		panic(err)
	}
}

func EmergencyGet(w http.ResponseWriter, r *http.Request) {
	eId := mux.Vars(r)["eid"]
	id, err := strconv.Atoi(eId)

	if err != nil {
		w.Write([]byte("error parsing eid"))
		return
	}

	e, ok := emergencies[id]
	if !ok {
		w.Write([]byte("emergency not present"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(*e); err != nil {
		panic(err)
	}
}

func EmergencyCreate(w http.ResponseWriter, r *http.Request) {
	emergencyCount := len(emergencies)
	users := data.GetUsers()
	pendingUserList := make([]models.UserStatus, 0)
	for _, user := range users {
		pendingUserList = append(pendingUserList, models.UserStatus{
			User:      user,
			UpdatedAt: time.Now(),
		})
	}
	st, su := data.NewSheet()
	newEmergency := models.Emergency{
		Id:          emergencyCount,
		CreatedAt:   time.Now(),
		PendingList: pendingUserList,
		SheetTitle:  st,
		SheetURL:    su,
	}
	emergencies[emergencyCount] = &newEmergency

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newEmergency); err != nil {
		panic(err)
	}

	sendPhoneNotifs(&newEmergency)
}

func EmergencyRespondSafe(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")
	if uuidStr == "" {
		w.Write([]byte("error in getting uuid"))
		return
	}

	emergency, ok := uuidToEmergency[uuidStr]

	if !ok {
		w.Write([]byte("invalid uuid"))
		return
	}

	user := uuidToUser[uuidStr]

	emergency.SafeList = append(emergency.SafeList, models.UserStatus{
		User:      *user,
		UpdatedAt: time.Now(),
	})

	for i, val := range emergency.PendingList {
		if val.User.Id == user.Id {
			emergency.PendingList[i] = emergency.PendingList[len(emergency.PendingList)-1]
			//emergency.PendingList[len(emergency.PendingList)-1] = models.UserStatus{}
			emergency.PendingList = emergency.PendingList[:len(emergency.PendingList)-1]
			break
		}
	}

	data.UpdateStatus(user.Id, true, emergency.SheetTitle)

	b, _ := ioutil.ReadFile("user-view-app/build/unbundled/index.html")
	w.Write(b)
	//w.Write([]byte("safely responded"))
	//delete(uuidToEmergency, uuidStr)
	//delete(uuidToEmergency, uuidStr)
	fmt.Println(*emergency)
}

func EmergencyRespondUnsafe(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")
	if uuidStr == "" {
		w.Write([]byte("error in getting uuid"))
		return
	}

	emergency, ok := uuidToEmergency[uuidStr]

	if !ok {
		w.Write([]byte("invalid uuid"))
		return
	}

	user := uuidToUser[uuidStr]

	emergency.UnsafeList = append(emergency.UnsafeList, models.UserStatus{
		User:      *user,
		UpdatedAt: time.Now(),
	})

	for i, val := range emergency.PendingList {
		if val.User.Id == user.Id {
			emergency.PendingList[i] = emergency.PendingList[len(emergency.PendingList)-1]
			//emergency.PendingList[len(emergency.PendingList)-1] = models.UserStatus{}
			emergency.PendingList = emergency.PendingList[:len(emergency.PendingList)-1]
			break
		}
	}

	data.UpdateStatus(user.Id, false, emergency.SheetTitle)

	b, _ := ioutil.ReadFile("user-view-app/build/unbundled/index.html")
	w.Write(b)
	//w.Write([]byte("oh no! help coming to you soon!"))
	//delete(uuidToEmergency, uuidStr)
	//delete(uuidToEmergency, uuidStr)
	fmt.Println(*emergency)
}

func sendPhoneNotifs(emergency *models.Emergency) {
	for _, userStatus := range emergency.PendingList {
		user := userStatus.User
		u, _ := uuid.NewV4()
		uuidStr := strings.Replace(u.String(), "-", "", -1)
		uuidToEmergency[uuidStr] = emergency
		uuidToUser[uuidStr] = &user
		sendClickLink(user, uuidStr)
	}
}

func sendClickLink(user models.User, uuidStr string) {
	from := "+12016056631"
	to := user.Phone
	message := "Hi " + user.Name + ", click here to alert your safety: http://" + os.Getenv("EXTERNAL_IP") + ":8080/emergency/respond/safe?uuid=" + uuidStr
	fmt.Println(twilio.SendSMS(from, to, message, "", ""))
}
