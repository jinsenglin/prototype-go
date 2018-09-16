package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func ChatsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chats/" {
		if r.Method == http.MethodPost {
			ch_id, _ := strconv.Atoi(r.FormValue("ch_id"))
			if channel := linectl.GetChannel(ch_id); channel == nil {
				message := fmt.Sprintf("Channel %d is not opened.", ch_id)
				http.Error(w, message, http.StatusInternalServerError)
			} else {
				channel.Notifier <- []byte(r.FormValue("chat"))
				path := fmt.Sprintf("/chats/new?ch_id=%d", ch_id)
				http.Redirect(w, r, path, 301)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/chats/new" {
		if r.Method == http.MethodGet {
			ch_id, _ := strconv.Atoi(r.URL.Query().Get("ch_id"))
			if channel := linectl.GetChannel(ch_id); channel == nil {
				message := fmt.Sprintf("Channel %d is not opened.", ch_id)
				http.Error(w, message, http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				fmt.Fprintf(w, "<form action='/chats/' method='POST'>chat: <input name='chat'/><input name='ch_id' value='%d' type='hidden'/><button>Submit</button></form>", ch_id)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
