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
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]/chats$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			if channel := linectl.GetChannel(0); channel == nil { // TODO: fix channel id
				http.Error(w, "Channel 0 is not opened.", http.StatusInternalServerError)
			} else {
				if flusher, ok := w.(http.Flusher); ok {
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
						select {
						case message := <-messageChan:
							fmt.Fprintf(w, "%s\n", message)
							flusher.Flush()
						case <-channel.Context.Done():
							fmt.Fprintf(w, "Channel 0 is closed.")
							return
						}
					}
				} else {
					http.Error(w, "Streaming is unsupported.", http.StatusBadRequest)
				}
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]/delete$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodDelete {
			if channel := linectl.GetChannel(0); channel == nil { // TODO: fix channel id
				http.Error(w, "Channel 0 is not opened.", http.StatusInternalServerError)
			} else {
				linectl.CloseChannel <- channel
				fmt.Fprintf(w, "Channel 0 is closed.")
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
