package handler

import (
	"fmt"
	"log"
	"net/http"
)

func ChatsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chats/" {
		if r.Method == http.MethodPost {
			log.Printf("TODO: add a chat '%s' to a channel", r.FormValue("chat")) // TODO
			http.Redirect(w, r, "/chats/new", 301)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if r.URL.Path == "/chats/new" {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form action='/chats/' method='POST'>chat: <input name='chat'/><button>Submit</button></form>") // TODO: fix form action and method.
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found")
	}
}
