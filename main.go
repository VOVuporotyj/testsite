package main

import (
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func indexPage(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "indexPage", nil)
}

func loginPage(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "loginPage", nil)

}

func main() {

	m := martini.Classic()
	m.Get("/", indexPage)
	m.Get("/login", loginPage)
	m.Run()
}
