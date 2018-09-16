package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

func SpacesAPIHandler(w http.ResponseWriter, r *http.Request) {
	if re, _ := regexp.Compile("^/spaces/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html></html>") // TODO
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
