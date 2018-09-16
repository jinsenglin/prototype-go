package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jinsenglin/prototype-go/pkg/controller/line"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

var linectl = line.Run()

func channel_id(path string) (id int) {
	re, _ := regexp.Compile("[1-9]")
	id, _ = strconv.Atoi(re.FindString(path))
	return
}

func ChannelsAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/channels/" {
		if r.Method == http.MethodPost {
			id, _ := strconv.Atoi(r.FormValue("id"))
			channel := model.NewChannel(id)
			linectl.OpenChannel <- channel
			message := fmt.Sprintf("Channel %d is opened.", id)
			fmt.Fprintf(w, message)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else if re, _ := regexp.Compile("^/channels/[1-9]/chats$"); re.MatchString(r.URL.Path) {
		if r.Method == http.MethodGet {
			id := channel_id(r.URL.Path)
			if channel := linectl.GetChannel(id); channel == nil {
				message := fmt.Sprintf("Channel %d is not opened.", id)
				http.Error(w, message, http.StatusInternalServerError)
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
							// TODO bug fix
							message := fmt.Sprintf("Channel %d is closed.", id)
							fmt.Fprintf(w, message)
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
			id := channel_id(r.URL.Path)
			if channel := linectl.GetChannel(id); channel == nil {
				message := fmt.Sprintf("Channel %d is not opened.", id)
				http.Error(w, message, http.StatusInternalServerError)
			} else {
				message := fmt.Sprintf("Channel %d is closed.", id)
				linectl.CloseChannel <- channel
				fmt.Fprintf(w, message)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
