package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
)

func indexPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(t.ExecuteTemplate(w, "indexPage", nil))

}

func main() {

	m := martini.Classic()
	m.Get("/", indexPage)
	m.Get("/regis", reg.regisPage)
	m.Post("/testPage", reg.postRegis)
	m.Get("/login", reg.loginPage)
	m.Post("/testPage", reg.postlogin)
	m.Run()
}
