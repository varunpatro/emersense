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

func GetMockUsers() []models.User {
	return mockUsers
}

func GetUsers() []models.User {
	data := getData()
	users := make([]models.User, 0)
	for _, u := range data {
		users = append(users, u)
	}

	return users
}

