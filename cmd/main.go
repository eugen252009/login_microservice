package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hallo")
	router := makeRouter()
	panic(http.ListenAndServe(":8080", router))
}

func makeRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/register", register)
	router.HandleFunc("/update", update)

	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hallo Index!"))
}
func login(w http.ResponseWriter, r *http.Request) {
	access, err := r.Cookie("access")
	reroll, err := r.Cookie("reroll")
	user, err := r.Cookie("user")
	session, err := r.Cookie("session")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf("access:%s reroll:%s user:%s session:%s", access.Name, reroll.Name, user.Name, session.Name)))
}

// ACCESSTOKEN
// REROLL TOKEN
// USER
// Session

func register(w http.ResponseWriter, r *http.Request) {
	a := uuid.NewString()
	w.Write([]byte(fmt.Sprintf("%s", a)))
}
func update(w http.ResponseWriter, r *http.Request) {
	reroll, err := r.Cookie("reroll")
	user, err := r.Cookie("user")
	session, err := r.Cookie("session")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	cookie := &http.Cookie{
		Name:     "auth",
		Value:    uuid.NewString(),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 15,
		Expires:  time.Now().Add(time.Minute * 15),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(fmt.Sprintf("Hallo Update! %s %s %s", reroll, user, session)))
}
