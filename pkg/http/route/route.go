/*
Implementations of http file uploader client and server
See https://gist.github.com/ebraminio/576fdfdff425bf3335b51a191a65dbdb

Copy a struct instance
See https://flaviocopes.com/go-copying-structs/

Convert an array to a slice
See https://stackoverflow.com/questions/28886616/convert-array-to-slice-in-go
*/
package route

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

type user struct {
	Id   int
	Name string
	Lang []string
}

func (u1 *user) copy() (u2 user) {
	// deep copy
	b, _ := json.Marshal(u1)
	json.Unmarshal(b, &u2)
	return u2
}

type users struct {
	items [9]user
	mux   sync.Mutex
}

func (data *users) list() []user {
	data.mux.Lock()
	us := data.items[:]
	data.mux.Unlock()

	return us
}

func (data *users) get(idx int) user {
	data.mux.Lock()
	u := data.items[idx]
	data.mux.Unlock()

	return u
}

func (data *users) create(idx int, id int, name string) {
	data.mux.Lock()
	data.items[idx].Id = id
	data.items[idx].Name = name
	data.mux.Unlock()
}

func (data *users) update(idx int, name string) {
	data.mux.Lock()
	data.items[idx].Name = name
	data.mux.Unlock()
}

func (data *users) delete(idx int) {
	data.mux.Lock()
	data.items[idx].Id = 0
	data.items[idx].Name = ""
	data.mux.Unlock()
}

func _idx(path string) int {
	re, _ := regexp.Compile("[1-9]")
	id, _ := strconv.Atoi(re.FindString(path))
	idx := id - 1
	return idx
}

var data = users{}

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

				if file, err := ioutil.TempFile("/tmp", "upload-"); err != nil {
					log.Println(err)
				} else {
					if n, err := io.Copy(file, r.Body); err != nil {
						log.Println(err)
					} else {
						log.Printf("%d bytes are recieved. Saved as %s\n", n, file.Name())
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

				for _, u := range data.list() {
					fmt.Fprintln(w, u)
				}
			} else if r.Method == http.MethodPost {
				// e.g.,
				// curl -v -X POST -L http://localhost:8080/users/ -F 'id=1' -F 'name=cclin'

				id, _ := strconv.Atoi(r.FormValue("id"))
				idx := id - 1
				data.create(idx, id, r.FormValue("name"))
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
				fmt.Fprintf(w, "%v", data.get(idx))
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodGet {
				// e.g.,
				// curl -v -X GET -L http://localhost:8080/users/1/edit

				idx := _idx(r.URL.Path)
				u := data.get(idx)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, "<form><div>id: %v</div>name: <input value='%v'/><button>Update</button></form>", u.Id, u.Name)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodPut {
				// e.g.,
				// curl -v -X PUT -L http://localhost:8080/users/1/update -F 'name=cc lin'

				idx := _idx(r.URL.Path)
				data.update(idx, r.FormValue("name"))
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method Not Allowed")
			}
		} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
			if r.Method == http.MethodDelete {
				// e.g.,
				// curl -v -X DELETE -L http://localhost:8080/users/1/delete

				idx := _idx(r.URL.Path)
				data.delete(idx)
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
