package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// HTTP GET to list users
			fmt.Fprintf(w, "Got users")
		} else if r.Method == "POST" {
			// HTTP POST to create a user
			fmt.Fprintf(w, "Created a user")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a user
		if r.Method == "GET" {
			fmt.Fprintf(w, "Got user 1")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/new", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a form to create a user
		if r.Method == "GET" {
			fmt.Fprintf(w, "Form to create a user")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/1/edit", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a form to update a user
		if r.Method == "GET" {
			fmt.Fprintf(w, "Form to update user 1")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/1/update", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a form to update a user
		if r.Method == "PUT" {
			fmt.Fprintf(w, "Updated user 1")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/1/delete", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a form to update a user
		if r.Method == "PUT" {
			fmt.Fprintf(w, "Deleted user 1")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
