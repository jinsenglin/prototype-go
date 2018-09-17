package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// DummyHandler ...
func DummyHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/dummy

	fmt.Fprintf(w, "Dummy")
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/login

	sid := strconv.Itoa(rand.Int())
	expiration := time.Now().Add(365 * 24 * time.Hour)
	c := http.Cookie{Name: "authed", Value: sid, Expires: expiration, Path: "/"}
	http.SetCookie(w, &c)

	fmt.Fprintf(w, "Logined")
}

// LogoutHandler ...
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/logout

	sid := strconv.Itoa(rand.Int())
	expiration := time.Now().Add(-365 * 24 * time.Hour)
	c := http.Cookie{Name: "authed", Value: sid, Expires: expiration, Path: "/"}
	http.SetCookie(w, &c)

	fmt.Fprintf(w, "Logouted")
}
