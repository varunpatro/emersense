package data

import "../models"

var mockUsers = []models.User{{
	Name: "varun",
	//Phone: "+6593716368",
}, {
	Name: "weihan",
}, {
	Name:  "ayush",
	Phone: "+6598970982",
}}

func GetUsers() []models.User {
	return mockUsers
}
