/*
Implementations of http file uploader client and server
See https://gist.github.com/ebraminio/576fdfdff425bf3335b51a191a65dbdb
*/
package route

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

var users [9]user

func RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// HTTP GET to visit welcome page
			// The "/" pattern matches everything, so we need to check
			// that we're at the root here.

			if r.URL.Path != "/" {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Page Not Found")
			} else {
				// e.g.,
				// curl -v -X GET -L http://localhost:8080
				// curl -v -X GET -L http://localhost:8080/

				fmt.Fprintf(w, "Welcome")
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	})

	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/files/" {
			if r.Method == http.MethodPost {
				// e.g.,
				// curl -v -X POST -L http://localhost:8080/files/ -H 'Content-Type: application/octet-stream' --data-binary '@README.md'

				if file, err := os.Create("/tmp/result"); err != nil { // TODO: replace with a random filename
					log.Println(err)
				} else {
					if n, err := io.Copy(file, r.Body); err != nil {
						log.Println(err)
					} else {
						log.Printf("%d bytes are recieved.\n", n)
						fmt.Fprintf(w, "%d bytes are recieved.\n", n)
					}
				}
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Page Not Found")
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// The "/users/" pattern matches everything prefixed, so we need to check
		// that we're at the ? here.

		if r.URL.Path == "/users/" {
			// /users will be redirected to /users/

			if r.Method == http.MethodGet {
				// e.g.,
				// curl -v -X GET -L http://localhost:8080/users
				// curl -v -X GET -L http://localhost:8080/users/

				for _, u := range users {
					fmt.Fprintln(w, u)
				}
			} else if r.Method == http.MethodPost {
				// e.g.,
				// curl -v -X POST -L http://localhost:8080/users/ -F 'id=1' -F 'name=cclin'

				id, _ := strconv.Atoi(r.FormValue("id"))
				idx := id - 1
				users[idx].id = id
				users[idx].name = r.FormValue("name")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if r.URL.Path == "/users/new" {
			if r.Method == http.MethodGet {
				// e.g.,
				// curl -v -X GET -L http://localhost:8080/users/new

				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, "<form>name: <input /><button>Create</button></form>")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodGet {
				// e.g.,
				// curl -v -X GET -L http://localhost:8080/users/1

				idx := _idx(r.URL.Path)
				fmt.Fprintf(w, "%v", users[idx])
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodGet {
				// e.g.,
				// curl -v -X GET -L http://localhost:8080/users/1/edit

				idx := _idx(r.URL.Path)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, "<form><div>id: %v</div>name: <input value='%v'/><button>Update</button></form>", users[idx].id, users[idx].name)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodPut {
				// e.g.,
				// curl -v -X PUT -L http://localhost:8080/users/1/update -F 'name=cc lin'

				idx := _idx(r.URL.Path)
				users[idx].name = r.FormValue("name")
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodDelete {
				// e.g.,
				// curl -v -X DELETE -L http://localhost:8080/users/1/delete

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
}