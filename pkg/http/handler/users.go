//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jinsenglin/prototype-go/pkg/http/context"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func user_idx(path string) (index int) {
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
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/users/new" {
		if r.Method == http.MethodGet {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/users/new

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form>name: <input /><button>Create</button></form>")
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/users/1

			idx := user_idx(r.URL.Path)
			reqcontext.SetUserIndex(r.Context(), idx)
			fmt.Fprintf(w, "%v", data.Get(idx))
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/edit$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/users/1/edit

			idx := user_idx(r.URL.Path)
			reqcontext.SetUserIndex(r.Context(), idx)
			u := data.Get(idx)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form><div>id: %v</div>name: <input value='%v'/><button>Update</button></form>", u.Id, u.Name)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/update$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodPut {
			// e.g.,
			// curl -v -X PUT -L http://localhost:8080/users/1/update -F 'name=cc lin'

			idx := user_idx(r.URL.Path)
			reqcontext.SetUserIndex(r.Context(), idx)
			data.Update(idx, r.FormValue("name"))
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/users/[1-9]/delete$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodDelete {
			// e.g.,
			// curl -v -X DELETE -L http://localhost:8080/users/1/delete

			idx := user_idx(r.URL.Path)
			reqcontext.SetUserIndex(r.Context(), idx)
			data.Delete(idx)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
