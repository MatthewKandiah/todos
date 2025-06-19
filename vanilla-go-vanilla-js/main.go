package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// let's just pretend this is a database!
	db := make(map[string][]Todo)
	todo := Todo{description: "Buy milk"}
	db["bob"] = append(db["bob"], todo)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("loginTemplate").Parse(loginTemplate))
		t.Execute(w, nil)
	})

	r.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		todos := db[name]
		data := UserTodosData{
			Name:  name,
			Todos: todos,
		}
		t := template.Must(template.New("userTodosTemplate").Parse(userTodosTemplate))
		t.Execute(w, data)
	}).Methods("GET").Schemes("http")

	r.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		redirectUrl := fmt.Sprintf("/user/%s", r.Form["loginName"][0])
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
	}).Methods("POST").Schemes("http")

	fmt.Println("Starting server")
	http.ListenAndServe(":8913", r)
}

type Todo struct {
	description string
}

type UserTodosData struct {
	Name  string
	Todos []Todo
}

const loginTemplate = `
<h1>Login</h1>
<form action="/user/login" method="post">
  <label for="loginName"></label>
  <input type="text" id="loginName" name="loginName"></input>
  <br/>
  <input type="submit" value="Submit"/>
</form>
	`

const userTodosTemplate = `
<h1>{{.Name}}</h1>
<ul>
  <li>{{.Todos}}</li>
</ul>
	`
