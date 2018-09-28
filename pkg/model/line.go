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
	"log"
)

// Line ...
type Line struct {
	Channels     map[int]*Channel
	OpenChannel  chan *Channel
	CloseChannel chan *Channel
}

// Listen ...
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
	go ch.Listen()
	this.Channels[ch.Id] = ch
	log.Println("Opened a channel")
}

func (this *Line) closeChannel(ch *Channel) {
	ch.Cancel()
	delete(this.Channels, ch.Id)
	log.Println("Closed a channel")
}

// GetChannel ...
func (this *Line) GetChannel(id int) (channel *Channel) {
	channel = this.Channels[id]
	return
}

// Dump ...
func (this *Line) Dump() {
	log.Println("Channels:")
	for k, v := range this.Channels {
		log.Printf("id: %v ch: %v", k, v)
	}
}

// NewLine ...
func NewLine() (line *Line) {
	line = &Line{
		Channels:     make(map[int]*Channel),
		OpenChannel:  make(chan *Channel),
		CloseChannel: make(chan *Channel)}
	return
}
