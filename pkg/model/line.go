package model

import (
	"log"
)

type Line struct {
	Channels map[int]*Channel
}

func (this *Line) OpenChannel() {
	this.Channels[len(this.Channels)] = &Channel{}
	log.Println("Opened a channel")
}

func (this *Line) CloseChannel(id int) {
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
	line = &Line{Channels: make(map[int]*Channel)}
	return
}
