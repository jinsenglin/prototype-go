package main

import (
	"log"
	"time"

	"github.com/jinsenglin/prototype-go/pkg/controller/line"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func main() {
	line := line.Run()

	// T0 :: system opens channel 0
	ch0 := model.NewChannel(0)
	line.OpenChannel <- ch0

	// T1 :: client a connects to channel 0
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client a received %s", <-messageChan)
		}
	}(ch0)

	// T2 :: client a sends a message to channel 0
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("a")
			log.Println("Sent a")
			time.Sleep(1e9)
		}
	}(ch0)

	// T3 :: client b connects to channel 0
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client b received %s", <-messageChan)
		}
	}(ch0)

	// T4 :: client b sends a message to channel 0
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("b")
			log.Println("Sent b")
			time.Sleep(1e9)
		}
	}(ch0)

	// T5 :: system is idle for a while
	time.Sleep(10e9)

	// T6 :: system opens channel 1
	ch1 := model.NewChannel(1)
	line.OpenChannel <- ch1

	// T7 :: client c connects to channel 1
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client c received %s", <-messageChan)
		}
	}(ch1)

	// T8 :: client c sends a message to channel 1
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("c")
			log.Println("Sent c")
			time.Sleep(1e9)
		}
	}(ch1)

	// T9 :: client d connects to channel 1
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			ch.ClosingClients <- messageChan
		}()
		for {
			log.Printf("Client d received %s", <-messageChan)
		}
	}(ch1)

	// T10 :: client d sends a message to channel 1
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("d")
			log.Println("Sent d")
			time.Sleep(1e9)
		}
	}(ch1)

	// T11 :: system is idle for a while
	time.Sleep(10e9)

	// T12 :: system closes channel 0
	line.CloseChannel <- ch0

	// T13 :: system is idle for a while
	time.Sleep(10e9)

	// T14 :: system closes channel 1
	line.CloseChannel <- ch1

	// T15 :: system is idle for a while
	time.Sleep(10e9)

	log.Println("ok")
}
