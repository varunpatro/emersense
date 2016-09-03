package main

import (
	"net/http"

	"./controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"ResidentsAll",
		"GET",
		"/residents/all",
		controllers.ResidentsAll,
	},
	Route{
		"EmergencyAll",
		"GET",
		"/emergency/all",
		controllers.EmergencyAll,
	},
	Route{
		"EmergencyGet",
		"GET",
		"/emergency/get/{eid}",
		controllers.EmergencyGet,
	},
	Route{
		"EmergencyCreate",
		"GET",
		"/emergency/create",
		controllers.EmergencyCreate,
	},
	Route{
		"EmergencyRespond",
		"GET",
		"/emergency/respond/{uuid}",
		controllers.EmergencyRespond,
	},
	//Route{
	//	"ResidentsReport",
	//	"POST",
	//	"/residents/report",
	//	controllers.ResidentsReport,
	//},
}
