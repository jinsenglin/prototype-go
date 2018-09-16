package handler

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// HTTP GET to visit welcome page
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.

		if r.URL.Path != "/" {
			http.NotFound(w, r)
		} else {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080
			// curl -v -X GET -L http://localhost:8080/

			fmt.Fprintf(w, "Welcome")
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
