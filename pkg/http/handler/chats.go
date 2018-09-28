//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ChatsAPIHandler ...
func ChatsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chats/" {
		if r.Method == http.MethodPost {
			ch_id, _ := strconv.Atoi(r.FormValue("ch_id"))
			if channel := linectl.GetChannel(ch_id); channel == nil {
				message := fmt.Sprintf("Channel %d is not opened.", ch_id)
				http.Error(w, message, http.StatusInternalServerError)
			} else {
				select {
				case <-channel.Context.Done():
					log.Println("Producer stops sending message due to channel closed.")
					message := fmt.Sprintf("Channel %d is closed.", ch_id)
					http.Error(w, message, http.StatusInternalServerError)
				case channel.Notifier <- []byte(r.FormValue("chat")):
					log.Println("Producer sends a message.")
					path := fmt.Sprintf("/chats/new?ch_id=%d", ch_id)
					http.Redirect(w, r, path, 301)
				}
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
