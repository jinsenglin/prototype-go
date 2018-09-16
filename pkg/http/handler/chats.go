package handler

import (
	"fmt"
	"net/http"
)

func ChatsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chats/" {
		if r.Method == http.MethodPost {
			if channel := linectl.GetChannel(0); channel == nil { // TODO: fix channel id
				http.Error(w, "Channel 0 is not opened.", http.StatusInternalServerError)
			} else {
				channel.Notifier <- []byte(r.FormValue("chat"))
				http.Redirect(w, r, "/chats/new", 301)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/chats/new" {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<form action='/chats/' method='POST'>chat: <input name='chat'/><button>Submit</button></form>")
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
