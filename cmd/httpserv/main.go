package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to list users
		fmt.Fprintf(w, "%+v", r.Method)

		// HTTP POST to create a user

		// HTTP PUT to update a user

		// HTTP DELETE to delete a user
	})

	http.HandleFunc("/user/id/1", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a user
		if r.Method == "GET" {
			fmt.Fprintf(w, "%+v", r.Method)
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a form to create a user
		if r.Method == "GET" {
			fmt.Fprintf(w, "Form to create a user")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/user/edit/1", func(w http.ResponseWriter, r *http.Request) {
		// HTTP GET to get a form to update a user
		if r.Method == "GET" {
			fmt.Fprintf(w, "Form to update a user")
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
