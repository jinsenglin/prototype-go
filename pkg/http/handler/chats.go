package handler

import (
	"fmt"
	"net/http"

	"github.com/jinsenglin/prototype-go/pkg/model"
)

func ChatsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chats/" {
		if r.Method == http.MethodPost {

			_channel := channels.Items[0] // TODO: use the real channel.
			_chats := _channel.Chats
			_chats = append(_chats, model.Chat{Message: r.FormValue("chat")}) // TOOD: refactor with a chan for goroutine safe.

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
