package routes

import (
	"html/template"
	"log"
	"net/http"
)

func Logging(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})

}



func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/index.html")
	tmpl.Execute(w, nil)
}

func Repetition(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("views/repetition.html")
	tmpl.Execute(w, nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
}