package main

import (
	"log"
	"time"

	"github.com/jinsenglin/prototype-go/pkg/controller/line"
)

func main() {
	line := line.Run()

	// DEMO CODE

	// T0
	line.OpenChannel <- 0
	line.Dump()
	time.Sleep(10e9)

	// T1
	line.OpenChannel <- 1
	line.Dump()
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
