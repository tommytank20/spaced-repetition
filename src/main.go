package main

import (
	"net/http"
	"spaced-repetition/src/routes"
)
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", routes.Redirect)
	http.HandleFunc("/index", routes.Index)
	http.HandleFunc("/repetition", routes.Repetition)
	http.ListenAndServe(":80", nil)
}