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

			if r.Method == "GET" {
				fmt.Fprintf(w, "TODO: return a list of users")
			} else if r.Method == "POST" {
				fmt.Fprintf(w, "TODO: create a user")
			} else {
				w.WriteHeader(405)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if r.URL.Path == "/users/new" {
			if r.Method == "GET" {
				fmt.Fprintf(w, "TODO: return a form for user creation")
			} else {
				w.WriteHeader(405)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]$"); re.MatchString(r.URL.Path) {
			if r.Method == "GET" {
				fmt.Fprintf(w, "TODO: return a user")
			} else {
				w.WriteHeader(405)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
			if r.Method == "GET" {
				fmt.Fprintf(w, "TODO: return a form for user modification")
			} else {
				w.WriteHeader(405)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
			if r.Method == "PUT" {
				fmt.Fprintf(w, "TODO: update a user")
			} else {
				w.WriteHeader(405)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
			if r.Method == "DELETE" {
				fmt.Fprintf(w, "TODO: delete a user")
			} else {
				w.WriteHeader(405)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Page Not Found")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
