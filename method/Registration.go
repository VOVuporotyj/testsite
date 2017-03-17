package registration

import (
	"fmt"
	"html/template"
	"net/http"

	"log"
	"os"

	"github.com/boltdb/bolt"
)

func regisPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/regis.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(t.ExecuteTemplate(w, "regisPage", nil))

}

func postRegis(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println(username)
	fmt.Println(password)

	db, err := bolt.Open("DB/test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(db.Path())

	// Execute several commands within a read-write transaction.
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

	t, err := template.ParseFiles("templates/Login.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(t.ExecuteTemplate(w, "loginPage", nil))

}

func loginPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	fmt.Println(t.ExecuteTemplate(w, "loginPage", nil))
}

func postlogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println(username)
	fmt.Println(password)
	t, err := template.ParseFiles("templates/testPage.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(t.ExecuteTemplate(w, "testPage", nil))
}
