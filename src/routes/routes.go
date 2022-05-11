package routes

import (
	"html/template"
	"log"
	"net/http"
)
func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)

	tmpl, _ := template.ParseFiles("views/index.html")
	tmpl.Execute(w, nil)
}

func Repetition(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)

	tmpl, _ := template.ParseFiles("views/repetition.html")
	tmpl.Execute(w, nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
}