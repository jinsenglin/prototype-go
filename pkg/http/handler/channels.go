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
						select {
						case <-channel.Context.Done():
							log.Printf("Consumer is forced to be disconnected from channel.")
						case channel.ClosingClients <- messageChan:
							log.Printf("Consumer says goodbay for disconnection from channel.")
						}
						log.Printf("Consumer has cleanup.")
					}()

					notify := w.(http.CloseNotifier).CloseNotify()
					go func() {
						<-notify
						select {
						case <-channel.Context.Done():
							log.Printf("Consumer is forced to be disconnected from channel.")
						case channel.ClosingClients <- messageChan:
							log.Printf("Consumer says goodbay for disconnection from channel.")
						}
						log.Printf("Consumer has cleanup.")
					}()

					for message := range messageChan {
						fmt.Fprintf(w, "%s\n", message)
						flusher.Flush()
					}
					message := fmt.Sprintf("Channel %d is closed.", id)
					fmt.Fprintf(w, message)
					flusher.Flush()
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
