package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type TestTemplateData = struct {
	Name string
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/user/{userId}", userHandler)
	router.HandleFunc("/register", registerHandler).Methods("PUT")

	return router
}

/*
TODO:
- show a login form that takes a name and sends you to that user's list on submission
- show a register form that takes a name and creates a user
*/
func homeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/page.html", "template/home.html"))
	t.Execute(w, nil)
}

/*
TODO:
- fetch user's todos and display them, with options to create, update and delete todos
*/
func userHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/page.html", "template/test.html"))
	vars := mux.Vars(r)
	testData := TestTemplateData{
		Name: vars["userId"],
	}
	t.Execute(w, testData)
	fmt.Fprintf(w, "<p>Expecting: Hello %s</p>", testData.Name)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// do some stuff
	fmt.Fprintf(w, "<h2>register response</h2>")
}
