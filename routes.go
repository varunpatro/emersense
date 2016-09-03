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
	//Route{
	//	"ResidentsReport",
	//	"POST",
	//	"/residents/report",
	//	controllers.ResidentsReport,
	//},
}
