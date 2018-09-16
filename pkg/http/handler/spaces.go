package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func SpacesAPIHandler(w http.ResponseWriter, r *http.Request) {
	if re, _ := regexp.Compile("^/spaces/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html><frameset rows='*,50px'><frame src='/channels/1/chats'><frame src='/chats/new'></frameset></html>") // TODO: fix channel id
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		log.Println(r.URL.Path)
		http.NotFound(w, r)
	}
}
