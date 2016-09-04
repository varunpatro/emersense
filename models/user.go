package models

import (
	"time"
)

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UserStatus struct {
	User      User      `json:"user"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLocation struct {
	Safe bool      `json:"safe"`
	Lat  float64   `json:"latitude"`
	Long float64   `json:"longitude"`
	Time time.Time `json:"time"`
}
