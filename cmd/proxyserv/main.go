package main

import (
	"io"
	"log"
	"net"
)

func handleConn(from net.Conn) {
	to, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Printf("%v", err)
	} else {
		done := make(chan struct{})
		go func() {
			defer from.Close()
			defer to.Close()
			io.Copy(from, to)
			done <- struct{}{}
		}()

		go func() {
			defer from.Close()
			defer to.Close()
			io.Copy(to, from)
			done <- struct{}{}
		}()

		<-done
		<-done
	}
}

func main() {
	log.Println("Try `curl -v -X GET -L http://localhost:8081`")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%v", err)
		} else {
			go handleConn(conn) // TODO: refactor with a Goroutine Pool. See proxyserv2
		}
	}
}
