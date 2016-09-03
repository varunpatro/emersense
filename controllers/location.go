package controllers

import (
	"../models"
	"encoding/json"
	"github.com/kellydunn/golang-geo"
	"net/http"
	"strconv"
	"time"
)

type status struct {
	status bool `json:"status"`
}

type DangerPoint struct {
	Lat    float64 `json:"latitude"`
	Long   float64 `json:"longitude"`
	Radius float64 `json:"radius"`
}

var uuidToLocation = make(map[string]models.UserLocation)
var DangerZones = getDangerZones()
var thresholdDist = 0.050

func LocationUpdate(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	long, _ := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
	timeVal, _ := strconv.ParseInt(r.URL.Query().Get("timestamp"), 10, 64)
	time := time.Unix(timeVal, 0)

	ul := models.UserLocation{
		Point: geo.NewPoint(lat, long),
		Time:  time,
	}
	uuidToLocation[uuidStr] = ul

	s := status{status: false}
	for _, dz := range DangerZones {
		if ul.Point.GreatCircleDistance(geo.NewPoint(dz.Lat, dz.Long)) > thresholdDist {
			s.status = true
			break
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

func LocationDangerZones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(getDangerZones()); err != nil {
		panic(err)
	}
}

func getDangerZones() []DangerPoint {
	points := make([]DangerPoint, 0)
	points = append(points, DangerPoint{Lat: 1.29565, Long: 103.856611, Radius: 0.05})  // south beach tower
	points = append(points, DangerPoint{Lat: 1.297588, Long: 103.854308, Radius: 0.05}) // national library
	return points
}
