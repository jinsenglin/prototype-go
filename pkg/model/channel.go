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

package model

import (
	"context"
	"log"
)

// Channel ...
type Channel struct {
	Id      int
	Context context.Context
	Cancel  context.CancelFunc

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	NewClients chan chan []byte

	// Closed client connections
	ClosingClients chan chan []byte

	// Client connections registry
	Clients map[chan []byte]bool
}

// Listen ...
func (this *Channel) Listen() {
	log.Printf("Started listening")
	for {
		select {
		case <-this.Context.Done():
			// Notify all connected clients
			for clientMessageChan, _ := range this.Clients {
				close(clientMessageChan)
			}
			log.Printf("Stopped listening")
			return

		case s := <-this.NewClients:
			// A new client has connected.
			// Register their message channel
			this.Clients[s] = true
			log.Printf("Client added. %d registered clients", len(this.Clients))

		case s := <-this.ClosingClients:
			// A client has dettached and we want to
			// stop sending them messages.
			delete(this.Clients, s)
			log.Printf("Removed client. %d registered clients", len(this.Clients))

		case event := <-this.Notifier:
			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan, _ := range this.Clients {
				clientMessageChan <- event
			}
		}
	}
}

// NewChannel ...
func NewChannel(id int) (channel *Channel) {
	ctx, cancel := context.WithCancel(context.Background())
	channel = &Channel{
		Id:             id,
		Context:        ctx,
		Cancel:         cancel,
		Notifier:       make(chan []byte, 2),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool)}
	return
}
