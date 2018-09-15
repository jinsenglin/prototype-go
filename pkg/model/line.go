package model

import (
	"log"
)

type Line struct {
	Channels     map[int]*Channel
	OpenChannel  chan int
	CloseChannel chan int
}

func (this *Line) Listen() {
	for {
		select {
		case id := <-this.OpenChannel:
			this.openChannel(id)
		case id := <-this.CloseChannel:
			this.closeChannel(id)
		}
	}
}

func (this *Line) openChannel(id int) {
	this.Channels[id] = &Channel{}
	log.Println("Opened a channel")
}

func (this *Line) closeChannel(id int) {
	delete(this.Channels, id)
	log.Println("Closed a channel")
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
		OpenChannel:  make(chan int),
		CloseChannel: make(chan int)}
	return
}
