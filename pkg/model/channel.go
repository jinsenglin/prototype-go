package model

import (
	"log"
	"sync"
)

type Channel struct {
	Id       int
	Chats    []*Chat
	Pipeline chan *Chat

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	NewClients chan chan []byte

	// Closed client connections
	ClosingClients chan chan []byte

	// Client connections registry
	Clients map[chan []byte]bool
}

func (this *Channel) Listen() {
	for {
		select {
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

func (this *Channel) Produce(chat *Chat) {
	log.Printf("Producing %v", chat)
	this.Pipeline <- chat
	log.Printf("Produced %v", chat)
}

func (this *Channel) Consume() (chat *Chat) {
	log.Printf("Consuming")
	chat = <-this.Pipeline
	log.Printf("Consumed %v", chat)

	this.Chats = append(this.Chats, chat)
	return
}

type Channels struct {
	Items []*Channel
	mux   sync.Mutex
}
