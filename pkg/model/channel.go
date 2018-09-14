package model

import (
	"log"
	"sync"
)

type Channel struct {
	Chats    []*Chat
	Pipeline chan Chat
}

func (this *Channel) Produce(chat Chat) {
	log.Printf("Producing %v", chat)
	this.Pipeline <- chat
	log.Printf("Produced %v", chat)
}

func (this *Channel) Consume() Chat {
	log.Printf("Consuming")
	chat := <-this.Pipeline
	log.Printf("Consumed %v", chat)
	return chat
}

type Channels struct {
	Items []*Channel
	mux   sync.Mutex
}
