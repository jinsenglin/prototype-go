package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/jinsenglin/prototype-go/pkg/model"
)

var channels = model.Channels{Items: make([]*model.Channel, 1)} // TODO

func ChannelsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/channels/" {
		if r.Method == http.MethodPost {
			// TODO: create a channel

			_channel := &model.Channel{Pipeline: make(chan model.Chat)}
			channels.Items[0] = _channel // TODO: use the real channel.
			go _channel.Consume()        // TODO: exit goroutine when channel is deleted.
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
