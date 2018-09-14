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
}
