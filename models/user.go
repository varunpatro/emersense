package models

import (
	"github.com/kellydunn/golang-geo"
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
	Point *geo.Point
	Time  time.Time
}
