package main

import (
	"log"
	"time"

	"github.com/jinsenglin/prototype-go/pkg/controller/line"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func main() {
	line := line.Run()

	// DEMO CODE

	// T0
	ch0 := &model.Channel{
		Id:             0,
		Notifier:       make(chan []byte, 100),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool)}
	line.OpenChannel <- ch0
	//line.Dump()

	go func(ch *model.Channel) {
		// Simulate client of a
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client a received %s", <-messageChan)
		}
	}(ch0)
	go func(ch *model.Channel) {
		// Simulate post a message to channel 0 from a
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("a")
			log.Println("Sent a")
			time.Sleep(1e9)
		}
	}(ch0)
	go func(ch *model.Channel) {
		// Simulate client of b
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client b received %s", <-messageChan)
		}
	}(ch0)
	go func(ch *model.Channel) {
		// Simulate post a message to channel 0 from b
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("b")
			log.Println("Sent b")
			time.Sleep(1e9)
		}
	}(ch0)

	time.Sleep(10e9)

	// T1
	ch1 := &model.Channel{
		Id:             1,
		Notifier:       make(chan []byte, 100),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool)}
	line.OpenChannel <- ch1
	//line.Dump()

	go func(ch *model.Channel) {
		// Simulate client of c
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client c received %s", <-messageChan)
		}
	}(ch1)
	go func(ch *model.Channel) {
		// Simulate post a message to channel 1 from c
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("c")
			log.Println("Sent c")
			time.Sleep(1e9)
		}
	}(ch1)
	go func(ch *model.Channel) {
		// Simulate client of d
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client d received %s", <-messageChan)
		}
	}(ch1)
	go func(ch *model.Channel) {
		// Simulate post a message to channel 1 from d
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("d")
			log.Println("Sent d")
			time.Sleep(1e9)
		}
	}(ch1)

	time.Sleep(10e9)

	// T2
	line.CloseChannel <- ch0
	//line.Dump()
	time.Sleep(10e9)

	// T3
	line.CloseChannel <- ch1
	//line.Dump()
	time.Sleep(10e9)

	// END DEMO

	log.Println("ok")
}
