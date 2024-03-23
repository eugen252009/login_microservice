package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var LoginQuery = "select id,name from nutzer where id=$1"

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error Occured while parsing form\n %s", err)
	}
	log.Printf(r.FormValue("user"))
	log.Printf(r.FormValue("password"))
	uuidstring := uuid.NewString()
	w.Header().Add("Authorization", uuidstring)
	w.WriteHeader(200)
	w.Write([]byte(uuidstring))
	// http.SetCookie(w, &http.Cookie{Name: "session", Value: uuidstring})
}

func register(w http.ResponseWriter, r *http.Request) {
	username, passowrd, ok := r.BasicAuth()
	if ok {
		fmt.Println(username, passowrd)
	}
	a := uuid.NewString()
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%s", a)))
}

// func update(w http.ResponseWriter, r *http.Request) {
// 	reroll, err := r.Cookie("reroll")
// 	user, err := r.Cookie("user")
// 	session, err := r.Cookie("session")
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	cookie := &http.Cookie{
// 		Name:     "auth",
// 		Value:    uuid.NewString(),
// 		Path:     "/",
// 		HttpOnly: true,
// 		MaxAge:   60 * 15,
// 		Expires:  time.Now().Add(time.Minute * 15),
// 	}
// 	http.SetCookie(w, cookie)
// 	w.Write([]byte(fmt.Sprintf("Hallo Update! %s %s %s", reroll, user, session)))
// }
