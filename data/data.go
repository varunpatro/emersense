package data

import "../models"

var mockUsers = []models.User{{
	Name: "varun",
}, {
	Name: "weihan",
}}

func GetUsers() []models.User {
	return mockUsers
}
