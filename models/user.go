package models

import "time"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type UserStatus struct {
	User      User            `json:"user"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
