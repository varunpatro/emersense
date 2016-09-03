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
		"UserAll",
		"GET",
		"/user/all",
		controllers.UserAll,
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
		"EmergencyRespondSafe",
		"GET",
		"/emergency/respond/safe",
		controllers.EmergencyRespondSafe,
	},
	Route{
		"EmergencyRespondUnsafe",
		"GET",
		"/emergency/respond/unsafe",
		controllers.EmergencyRespondUnsafe,
	},
}
