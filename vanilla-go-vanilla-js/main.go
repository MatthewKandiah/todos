package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<a href=\"https:localhost:8913/books/moby-dick/page/456\">Wanna hear about Moby Dick?</a>")
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		data := TemplateData{
			BookName: title,
			Page:     page,
		}
		t := template.Must(template.ParseFiles("template.html"))
		fmt.Fprintf(w, "<!DOCTYPE html><html lang=\"en\"><head></head><body>")
		fmt.Fprintf(w, "<h4>Template contents start...</h4>")
		t.Execute(w, data)
		fmt.Fprintf(w, "<h4>Template contents finished!</h4>")
		fmt.Fprintf(w, "</body></html>")
	})

	fmt.Println("Starting server")
	http.ListenAndServe(":8913", r)
}

type TemplateData struct {
	BookName string
	Page     string
}
