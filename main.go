package main

import (
	"usegolang/views"
	"net/http"
	//"html/template"
	"github.com/gorilla/mux"
)

var homeTemplate *views.View 
var contactTemplate *views.View


func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTemplate.Template.ExecuteTemplate(w, homeTemplate.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate.Template.ExecuteTemplate(w, contactTemplate.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {

	homeTemplate = views.NewView("bootstrap", "views/home.gohtml")
	contactTemplate = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/contact", contactFunc)
	http.ListenAndServe(":8000", r)
}