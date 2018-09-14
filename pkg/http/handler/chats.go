package handler

import (
	"fmt"
	"net/http"
)

func ChatsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chats/new" {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form>chat: <input /><button>Submit</button></form>") // TODO: fix form action and method.
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found")
	}
}
