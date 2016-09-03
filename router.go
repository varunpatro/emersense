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
	asrc := http.StripPrefix("/src/admin-panel/", http.FileServer(http.Dir("./admin-panel-app/build/bundled/src/admin-panel/")))
	usrc := http.StripPrefix("/src/user-view/", http.FileServer(http.Dir("./user-view-app/build/bundled/src/user-view/")))
	router.PathPrefix("/bower_components/").Handler(bc)
	router.PathPrefix("/src/admin-panel").Handler(asrc)
	router.PathPrefix("/src/user-view").Handler(usrc)

	return router
}
