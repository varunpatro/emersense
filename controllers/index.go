package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	b, e := ioutil.ReadFile("admin-panel-app/build/bundled/index.html")
	w.Write(b)
	//fmt.Fprint(w, "Welcome!\n")
	fmt.Println(e)
}
