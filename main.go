package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/VOVuporotyj/testsite/Reg"
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
	m.Get("/regis", Reg.RegisPage)
	m.Post("/testPage", Reg.PostRegis)
	m.Get("/login", Reg.LoginPage)
	m.Post("/testPage", Reg.Postlogin)
	m.Run()
}
