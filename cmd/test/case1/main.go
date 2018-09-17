package main

import (
	"fmt"
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
	log.Println("Consumer a is going to connect to the channel 0.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer a is forced to be disconnected from channel 0.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Consumer a is going to disconnect from channel 0.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer a stops receiving message due to channel 0 closed.")
				return
			case s := <-messageChan:
				log.Printf("Consumer a receives a message %s", s)
			}
		}
	}(ch0)

	// T3 :: client a sends a message to channel 0
	log.Println("Producer a is going to send 5 messages to the channel 0.")
	go func(ch *model.Channel) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r) // ch.Notifier is closed.
			}
		}()

		for i := 1; i < 5; i++ {
			select {
			case <-ch.Context.Done():
				log.Println("Producer a stops sending message due to channel 0 closed.")
				return
			case ch.Notifier <- []byte("a"):
				log.Println("Producer a sends a message a.")
				time.Sleep(1e9)
			}
		}
	}(ch0)

	// T4 :: client b connects to channel 0
	log.Println("Consumer b is going to connect to the channel 0.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer b is forced to be disconnected from channel 0.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Consumer b is going to disconnect from channel 0.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer b stops receiving message due to channel 0 closed.")
				return
			case s := <-messageChan:
				log.Printf("Consumer b receives a message %s", s)
			}
		}
	}(ch0)

	// T5 :: client b sends a message to channel 0
	log.Println("Producer b is going to keep sending messages to the channel 0.")
	go func(ch *model.Channel) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r) // ch.Notifier is closed.
			}
		}()

		for {
			select {
			case <-ch.Context.Done():
				log.Println("Producer b stops sending message due to channel 0 closed.")
				return
			case ch.Notifier <- []byte("b"):
				log.Println("Producer b sends a message b.")
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
	log.Println("Consumer c is going to connect to the channel 1.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer c is forced to be disconnected from channel 1.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Consumer c is going to disconnect from channel 1.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer c stops receiving message due to channel 1 closed.")
				return
			case s := <-messageChan:
				log.Printf("Consumer c receives a message %s", s)
			}
		}
	}(ch1)

	// T9 :: client c sends a message to channel 1
	log.Println("Producer c is going to send 5 messages to the channel 1.")
	go func(ch *model.Channel) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r) // ch.Notifier is closed.
			}
		}()

		for i := 1; i < 5; i++ {
			select {
			case <-ch.Context.Done():
				log.Println("Producer c stops sending message due to channel 1 closed.")
				return
			case ch.Notifier <- []byte("c"):
				log.Println("Producer c sends a message c.")
				time.Sleep(1e9)
			}
		}
	}(ch1)

	// T10 :: client d connects to channel 1
	log.Println("Consumer d is going to connect to the channel 1.")
	go func(ch *model.Channel) {
		messageChan := make(chan []byte)
		ch.NewClients <- messageChan
		defer func() {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer d is forced to be disconnected from channel 1.")
			case ch.ClosingClients <- messageChan:
				log.Printf("Consumer d is going to disconnect from channel 1.")
			}
		}()
		for {
			select {
			case <-ch.Context.Done():
				log.Printf("Consumer d stops receiving message due to channel 1 closed.")
				return
			case s := <-messageChan:
				log.Printf("Consumer d receives a message %s", s)
			}
		}
	}(ch1)

	// T11 :: client d sends a message to channel 1
	log.Println("Producer d is going to keep sending messages to the channel 1.")
	go func(ch *model.Channel) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r) // ch.Notifier is closed.
			}
		}()

		for {
			select {
			case <-ch.Context.Done():
				log.Println("Producer d stops sending message due to channel 1 closed.")
				return
			case ch.Notifier <- []byte("d"):
				log.Println("Producer d sends a message d.")
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
