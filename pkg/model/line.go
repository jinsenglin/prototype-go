package model

import (
	"context"
	"log"
)

type Line struct {
	Channels     map[int]*Channel
	OpenChannel  chan *Channel
	CloseChannel chan *Channel
}

func (this *Line) Listen() {
	for {
		select {
		case ch := <-this.OpenChannel:
			this.openChannel(ch)
		case ch := <-this.CloseChannel:
			this.closeChannel(ch)
		}
	}
}

func (this *Line) openChannel(ch *Channel) {
	ch.Context, ch.Cancel = context.WithCancel(context.Background())

	go ch.Listen()

	this.Channels[ch.Id] = ch
	log.Println("Opened a channel")
}

func (this *Line) closeChannel(ch *Channel) {
	ch.Cancel()

	delete(this.Channels, ch.Id)
	log.Println("Closed a channel")
}

func (this *Line) GetChannel(id int) (channel *Channel) {
	channel = this.Channels[id]
	return
}

func (this *Line) Dump() {
	log.Println("Channels:")
	for k, v := range this.Channels {
		log.Printf("id: %v ch: %v", k, v)
	}
}

func NewLine() (line *Line) {
	line = &Line{
		Channels:     make(map[int]*Channel),
		OpenChannel:  make(chan *Channel),
		CloseChannel: make(chan *Channel)}
	return
}
