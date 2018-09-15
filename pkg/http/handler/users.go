package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jinsenglin/prototype-go/pkg/http/context"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func _idx(path string) (index int) {
	re, _ := regexp.Compile("[1-9]")
	id, _ := strconv.Atoi(re.FindString(path))
	index = id - 1

	return
}

var data = model.Users{}

func UsersAPIHandler(w http.ResponseWriter, r *http.Request) {
	// The "/users/" pattern matches everything prefixed, so we need to check
	// that we're at the ? here.

	if r.URL.Path == "/users/" {
		// /users will be redirected to /users/

		if r.Method == http.MethodGet {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/users
			// curl -v -X GET -L http://localhost:8080/users/

			for _, u := range data.List() {
				fmt.Fprintln(w, u)
			}
		} else if r.Method == http.MethodPost {
			// e.g.,
			// curl -v -X POST -L http://localhost:8080/users/ -F 'id=1' -F 'name=cclin'

			id, _ := strconv.Atoi(r.FormValue("id"))
			idx := id - 1
			reqcontext.SetUserIndex(r.Context(), idx)
			data.Create(idx, id, r.FormValue("name"))
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
			reqcontext.SetUserIndex(r.Context(), idx)
			fmt.Fprintf(w, "%v", data.Get(idx))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/users/1/edit

			idx := _idx(r.URL.Path)
			reqcontext.SetUserIndex(r.Context(), idx)
			u := data.Get(idx)
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
			reqcontext.SetUserIndex(r.Context(), idx)
			data.Update(idx, r.FormValue("name"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodDelete {
			// e.g.,
			// curl -v -X DELETE -L http://localhost:8080/users/1/delete

			idx := _idx(r.URL.Path)
			reqcontext.SetUserIndex(r.Context(), idx)
			data.Delete(idx)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found")
	}
}
