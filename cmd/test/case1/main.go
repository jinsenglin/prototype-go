package main

import (
	"log"
	"time"

	"github.com/jinsenglin/prototype-go/pkg/controller/line"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func main() {
	// T0 :: system starts a line.
	log.Println("A line controller is going to run.")
	line := line.Run()

	// T1 :: system opens a channel 0
	log.Println("A channel of id 0 is going to be open.")
	ch0 := model.NewChannel(0)
	line.OpenChannel <- ch0

	// T2 :: client a connects to channel 0
	log.Println("A client of id a is going to connect to the channel 0.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Client a receives a done signal from channel 0.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Client a is leaving.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Client a stops receiving message due to channel 0 closed.")
				return
			case s := <-messageChan:
				log.Printf("Client a receives a message %s", s)
			}
		}
	}(ch0)

	// T3 :: client a sends a message to channel 0
	log.Println("A client of id a is going to send 5 messages to the channel 0.")
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			select {
			case <-ch.Context.Done():
				log.Println("Client a stops sending message due to channel 0 closed.")
				return
			case ch.Notifier <- []byte("a"):
				log.Println("Client a sends a message a.")
				time.Sleep(1e9)
			}
		}
	}(ch0)

	// T4 :: client b connects to channel 0
	log.Println("A client of id b is going to connect to the channel 0.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Client b stops receiving message due to channel 0 closed.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Client b is leaving.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Client b receives a done signal from channel 0.")
				return
			case s := <-messageChan:
				log.Printf("Client b receives a message %s", s)
			}
		}
	}(ch0)

	// T5 :: client b sends a message to channel 0
	log.Println("A client of id b is going to keep sending messages to the channel 0.")
	go func(ch *model.Channel) {
		for {
			select {
			case <-ch.Context.Done():
				log.Println("Client b stops sending message due to channel 0 closed.")
				return
			case ch.Notifier <- []byte("b"):
				log.Println("Client b sends a message b.")
				time.Sleep(1e9)
			}
		}
	}(ch0)

	// T6 :: system is idle for a while
	log.Println("Main func sleep 10 seconds.")
	time.Sleep(10e9)

	// T7 :: system opens channel 1
	log.Println("A channel of id 1 is going to be open.")
	ch1 := model.NewChannel(1)
	line.OpenChannel <- ch1

	// T8 :: client c connects to channel 1
	log.Println("A client of id c is going to connect to the channel 1.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Client c receives a done signal from channel 1.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Client c is leaving.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Client c stops receiving message due to channel 1 closed.")
				return
			case s := <-messageChan:
				log.Printf("Client c receives a message %s", s)
			}
		}
	}(ch1)

	// T9 :: client c sends a message to channel 1
	log.Println("A client of id c is going to send 5 messages to the channel 1.")
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			select {
			case <-ch.Context.Done():
				log.Println("Client c stops sending message due to channel 1 closed.")
				return
			case ch.Notifier <- []byte("c"):
				log.Println("Client c sends a message c.")
				time.Sleep(1e9)
			}
		}
	}(ch1)

	// T10 :: client d connects to channel 1
	log.Println("A client of id d is going to connect to the channel 1.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Client d receives a done signal from channel 1.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Client d is leaving.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Client d stops receiving message due to channel 1 closed.")
				return
			case s := <-messageChan:
				log.Printf("Client d receives a message %s", s)
			}
		}
	}(ch1)

	// T11 :: client d sends a message to channel 1
	log.Println("A client of id d is going to keep sending messages to the channel 1.")
	go func(ch *model.Channel) {
		for {
			select {
			case <-ch.Context.Done():
				log.Println("Client d stops sending message due to channel 1 closed.")
				return
			case ch.Notifier <- []byte("d"):
				log.Println("Client d sends a message d.")
				time.Sleep(1e9)
			}
		}
	}(ch1)

	// T12 :: system is idle for a while
	log.Println("Main func sleep 10 seconds.")
	time.Sleep(10e9)

	// T13 :: system closes channel 0
	log.Println("A channel of id 0 is going to be closed.")
	line.CloseChannel <- ch0

	// T14 :: system is idle for a while
	log.Println("Main func sleep 10 seconds.")
	time.Sleep(10e9)

	// T15 :: system closes channel 1
	log.Println("A channel of id 1 is going to be closed.")
	line.CloseChannel <- ch1

	// T16 :: system is idle for a while
	log.Println("Main func sleep 10 seconds.")
	time.Sleep(10e9)

	log.Println("ok")
}
