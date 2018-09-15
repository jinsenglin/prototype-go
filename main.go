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
	line.OpenChannel <- 0
	line.Dump()

	ch0 := line.GetChannel(0)
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("a")
			log.Println("Sent a")
			time.Sleep(1e9)
		}
	}(ch0)
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("b")
			log.Println("Sent b")
			time.Sleep(1e9)
		}
	}(ch0)

	time.Sleep(10e9)

	// T1
	line.OpenChannel <- 1
	line.Dump()

	ch1 := line.GetChannel(1)
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("c")
			log.Println("Sent c")
			time.Sleep(1e9)
		}
	}(ch1)
	go func(ch *model.Channel) {
		for i := 1; i < 5; i++ {
			ch.Notifier <- []byte("d")
			log.Println("Sent d")
			time.Sleep(1e9)
		}
	}(ch1)

	time.Sleep(10e9)

	// T2
	line.CloseChannel <- 1
	line.Dump()
	time.Sleep(10e9)

	// T3
	line.CloseChannel <- 0
	line.Dump()
	time.Sleep(10e9)

	// END DEMO

	log.Println("ok")
}
