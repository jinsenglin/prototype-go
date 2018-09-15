package model

import (
	"log"
	"sync"
)

type Channel struct {
	Id       int
	Chats    []*Chat
	Pipeline chan *Chat
}

func (this *Channel) Listen() {
	for {
		select {}
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
