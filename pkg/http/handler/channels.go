package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/jinsenglin/prototype-go/pkg/controller/line"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

var linectl = line.Run()

func ChannelsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/channels/" {
		if r.Method == http.MethodPost {
			channel := model.NewChannel(0) // TODO: fix channel id
			linectl.OpenChannel <- channel
			fmt.Fprintf(w, "Channel 0 is created.")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]/chats$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			if channel := linectl.GetChannel(0); channel == nil { // TODO: fix channel id
				http.Error(w, "Channel 0 is not opened.", http.StatusInternalServerError)
			} else {
				flusher, ok := w.(http.Flusher)

				if !ok {
					http.Error(w, "Streaming is unsupported!", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")
				w.Header().Set("Access-Control-Allow-Origin", "*")

				messageChan := make(chan []byte)
				channel.NewClients <- messageChan
				defer func() {
					channel.ClosingClients <- messageChan
				}()
				for {
					fmt.Fprintf(w, "%s\n", <-messageChan)
					flusher.Flush()
				}
			}
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
