package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
	"github.com/go-martini/martini"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func indexPage(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "indexPage", nil)
}

//RegisPage открывает шаблон regis.
func RegisPage(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "regisPage", nil)
}

//PostRegis записывает в БД лог и пас.
func PostRegis(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println(username)
	fmt.Println(password)

	db, err := bolt.Open("DB/test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(db.Path())

	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Login"))
		if err != nil {
			return err
		}
		if err := b.Put([]byte(username), []byte(password)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	templates.ExecuteTemplate(w, "loginPage", nil)
}

//LoginPage открывает шаблон Login.
func LoginPage(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "loginPage", nil)
}

//Postlogin сравнивает логин и пароль с базой.
func Postlogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println(username)
	fmt.Println(password)
	db, err := bolt.Open("DB/test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(db.Path())

	if err := db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("login")).Get([]byte("username"))
		if string(v) == username {
			templates.ExecuteTemplate(w, "testPage", nil)
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	templates.ExecuteTemplate(w, "testPage", nil)
}

func main() {

	m := martini.Classic()
	m.Get("/", indexPage)
	m.Get("/regis", RegisPage)
	m.Post("/testPage", PostRegis)
	m.Get("/login", LoginPage)
	m.Post("/testPage", Postlogin)
	m.Run()
}
