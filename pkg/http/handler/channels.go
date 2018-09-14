package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

func ChannelsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/channels/" {
		if r.Method == http.MethodPost {
			// TODO: create a channel
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			// TODO: check channel existence
			// TODO: return a list of chats
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]/update$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodPut {
			// TODO: check channel existence
			// TODO: update a channel by adding a chat
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]/delete$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodDelete {
			// TODO: check channel existence
			// TODO: delete a channel
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found")
	}
}
