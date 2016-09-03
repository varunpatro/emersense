package data

import "../models"

var mockUsers = []models.User{{
	Id:    1,
	Name:  "Ayush",
	Phone: "+6598970982",
}, {
	Id:    2,
	Name:  "Weihan",
	Phone: "+90011282",
},
}

func GetMockUsers() []models.User {
	return mockUsers
}

func GetUsers() []models.User {
	return mockUsers
	data := GetData()
	users := make([]models.User, 0)
	for _, u := range data {
		users = append(users, u)
	}

	return users
}
