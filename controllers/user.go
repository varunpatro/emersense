package controllers

import (
	"encoding/json"
	"net/http"
	"../data"
)

func UserAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	allResidents := data.GetUsers()
	if err := json.NewEncoder(w).Encode(allResidents); err != nil {
		panic(err)
	}
}
