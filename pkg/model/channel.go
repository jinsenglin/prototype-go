package model

import (
	"context"
	"log"
)

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

func (this *Channel) Listen() {
	for {
		select {
		case <-this.Context.Done():
			for s := range this.ClosingClients {
				delete(this.Clients, s)
				log.Printf("Removed client. %d registered clients", len(this.Clients))
			}
			log.Printf("Stopped a channel")
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
			log.Printf("Broadcasted %s", event)
		}
	}
}

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
