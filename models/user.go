package models

import (
	"time"
	"github.com/kellydunn/golang-geo"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserStatus struct {
	User      User      `json:"user"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLocation struct {
	Point *geo.Point
	Time  time.Time
}
