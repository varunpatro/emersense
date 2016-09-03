package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	bc := http.StripPrefix("/bower_components/", http.FileServer(http.Dir("./admin-panel-app/build/bundled/bower_components/")))
	src := http.StripPrefix("/src/", http.FileServer(http.Dir("./admin-panel-app/build/bundled/src/")))
	router.PathPrefix("/bower_components/").Handler(bc)
	router.PathPrefix("/src/").Handler(src)

	return router
}
