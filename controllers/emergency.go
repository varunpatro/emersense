package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/sfreiberg/gotwilio"

	"../data"
	"../models"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"strings"
)

var twilio *gotwilio.Twilio

func init() {
	accountSid := "AC30009862e9b061e21fcde75bc13cb49b"
	authToken := "95b081654eb589a8397ae216067502c8"
	twilio = gotwilio.NewTwilioClient(accountSid, authToken)
}

var emergencies = make(map[int]models.Emergency)

var uuidToEmergency = make(map[string]*models.Emergency)
var uuidToUser = make(map[string]*models.User)

func EmergencyAll(w http.ResponseWriter, r *http.Request) {
	emap := make(map[int]string)
	keys := make([]int, 0, len(emap))
	for k := range emap {
		keys = append(keys, k)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(keys); err != nil {
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

	e := emergencies[id]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
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
	newEmergency := models.Emergency{
		Id:          emergencyCount,
		CreatedAt:   time.Now(),
		PendingList: pendingUserList,
	}
	emergencies[emergencyCount] = newEmergency

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newEmergency); err != nil {
		panic(err)
	}

	sendPhoneNotifs(&newEmergency)
}

func EmergencyRespond(w http.ResponseWriter, r *http.Request) {
	uuidStr := mux.Vars(r)["uuid"]
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

	for i, val := range emergency.PendingList {
		if val.User.Id == user.Id {
			emergency.PendingList[i] = emergency.PendingList[len(emergency.PendingList)-1]
			//emergency.PendingList[len(emergency.PendingList)-1] = nil
			emergency.PendingList = emergency.PendingList[:len(emergency.PendingList)-1]
			break
		}
	}

	emergency.SafeList = append(emergency.SafeList, models.UserStatus{
		User:      *user,
		UpdatedAt: time.Now(),
	})

	w.Write([]byte("safely responded"))
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
	message := "Hi " + user.Name + ", click here to alert your safety: \"http://localhost:8080/emergency/respond/" + uuidStr + "\""
	fmt.Println(twilio.SendSMS(from, to, message, "", ""))
}
