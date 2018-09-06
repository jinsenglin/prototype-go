package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// HTTP GET to visit welcome page
			// The "/" pattern matches everything, so we need to check
			// that we're at the root here.
			if r.URL.Path != "/" {
				w.WriteHeader(404)
				fmt.Fprintf(w, "Page Not Found")
			} else {
				fmt.Fprintf(w, "Welcome")
			}
		} else {
			w.WriteHeader(405)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// The "/users/" pattern matches everything prefixed, so we need to check
		// that we're at the ? here.

		if r.URL.Path == "/users/" {
			// /users will be redirected to /users/
			fmt.Fprintf(w, "/users/")
		} else if r.URL.Path == "/users/new" {
			fmt.Fprintf(w, "/users/new")
		} else if re, _ := regexp.Compile("^/users/[1-9]$"); re.MatchString(r.URL.Path) {
			fmt.Fprintf(w, "get a user")
		} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
			fmt.Fprintf(w, "edit a user")
		} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
			fmt.Fprintf(w, "update a user")
		} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
			fmt.Fprintf(w, "delete a user")
		} else {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Page Not Found")
		}
	})
	/*
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
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}
