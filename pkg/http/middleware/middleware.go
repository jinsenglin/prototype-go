package middleware

import (
	"log"
	"net/http"
	"time"
)

// Dummy ...
func Dummy(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

// Timed ...
func Timed(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		end := time.Now()
		log.Println("The request took", end.Sub(start))
	}
}
