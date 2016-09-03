package main

import (
	//"./data"
	"log"
	"net/http"
)

func main() {
	//data.GetData()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
