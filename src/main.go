package main

import (
	"net/http"
	"spaced-repetition/src/routes"
)
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", routes.Logging(routes.Redirect))
	http.Handle("/index", routes.Logging(routes.Index))
	http.Handle("/repetition", routes.Logging(routes.Repetition))
	http.ListenAndServe(":80", nil)
}