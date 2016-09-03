package models

import "time"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserStatus struct {
	User      User      `json:"user"`
	Status    bool      `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
}
