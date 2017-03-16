package main

import (
	"html/template"
	"net/http"

	"fmt"

	"github.com/go-martini/martini"
)

func indexPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "indexPage", nil)

}

func loginPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "loginPage", nil)

}

func main() {

	m := martini.Classic()
	m.Get("/", indexPage)
	m.Get("/login", loginPage)
	m.Run()
}
