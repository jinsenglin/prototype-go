package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func space_id(path string) (id int) {
	re, _ := regexp.Compile("[1-9]")
	id, _ = strconv.Atoi(re.FindString(path))
	return
}

func SpacesAPIHandler(w http.ResponseWriter, r *http.Request) {
	if re, _ := regexp.Compile("^/spaces/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			id := space_id(r.URL.Path)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html><frameset rows='*,50px'><frame src='/channels/%d/chats'><frame src='/chats/new?ch_id=%d'></frameset></html>", id, id)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		log.Println(r.URL.Path)
		http.NotFound(w, r)
	}
}
