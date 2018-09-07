/*
Additional Resources
- https://golang.org/pkg/net/http/
- https://golang.org/doc/articles/wiki/
- https://gowebexamples.com/http-server/
- http://legendtkl.com/2016/08/21/go-web-server/
- https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type user struct {
	id   int
	name string
}

func _idx(path string) int {
	re, _ := regexp.Compile("[1-9]")
	id, _ := strconv.Atoi(re.FindString(path))
	idx := id - 1
	return idx
}

func main() {
	var users [9]user

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// HTTP GET to visit welcome page
			// The "/" pattern matches everything, so we need to check
			// that we're at the root here.
			if r.URL.Path != "/" {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Page Not Found")
			} else {
				fmt.Fprintf(w, "Welcome")
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// The "/users/" pattern matches everything prefixed, so we need to check
		// that we're at the ? here.

		if r.URL.Path == "/users/" {
			// /users will be redirected to /users/

			if r.Method == http.MethodGet {
				for _, u := range users {
					fmt.Fprintln(w, u)
				}
			} else if r.Method == http.MethodPost {
				fmt.Fprintf(w, "TODO: create a user")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if r.URL.Path == "/users/new" {
			if r.Method == http.MethodGet {
				fmt.Fprintf(w, "TODO: return a form for user creation")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodGet {
				idx := _idx(r.URL.Path)
				fmt.Fprintf(w, "%v", users[idx])
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodGet {
				fmt.Fprintf(w, "TODO: return a form for user modification")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodPut {
				fmt.Fprintf(w, "TODO: update a user")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodDelete {
				idx := _idx(r.URL.Path)
				users[idx].id = 0
				users[idx].name = ""
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Page Not Found")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
