package controllers

import (
	"encoding/json"
	"github.com/kellydunn/golang-geo"
	"net/http"
	"strconv"
	"time"
)

type UserLocation struct {
	Point *geo.Point
	Time  time.Time
}

var uuidToLocation = make(map[string]UserLocation)
var DangerZones = getDangerZones()
var thresholdDist = 0.050

func LocationUpdate(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	long, _ := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
	timeVal, _ := strconv.ParseInt(r.URL.Query().Get("timestamp"), 10, 64)
	time := time.Unix(timeVal, 0)

	ul := UserLocation{
		Point: geo.NewPoint(lat, long),
		Time:  time,
	}
	uuidToLocation[uuidStr] = ul

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	type status struct {
		status bool `json:"status"`
	}
	s := status{status: false}
	for _, pt := range DangerZones {
		if ul.Point.GreatCircleDistance(pt) > thresholdDist {
			s.status = true
			break
		}
	}
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

func getDangerZones() []*geo.Point {
	points := make([]*geo.Point, 0)
	points = append(points, geo.NewPoint(1.29565, 103.856611))  // south beach tower
	points = append(points, geo.NewPoint(1.297588, 103.854308)) // national library
	return points
}
