package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nu7hatch/gouuid"

	"../data"
	"../models"
)

var emergencies = make(map[int]models.Emergency)


func EmergencyCreate(w http.ResponseWriter, r *http.Request) {
	emergencyCount := len(emergencies)
	users := data.GetUsers()
	pendingUserList := make([]models.UserStatus, 0)
	for _, user := range users {
		append(pendingUserList, models.UserStatus{
			User:      user,
			UpdatedAt: time.Now(),
		})
	}
	newEmergency := models.Emergency{
		CreatedAt:   time.Now(),
		PendingList: pendingUserList,
	}
	emergencies[emergencyCount] = newEmergency

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newEmergency); err != nil {
		panic(err)
	}

	sendPhoneNotifs(newEmergency.PendingList)
}

func sendPhoneNotifs(userList []models.UserStatus) {
	for _, userStatus := range userList {
		user := userStatus.User

	}
}
