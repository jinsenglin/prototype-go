/*
Addiontional Resources
- https://gist.github.com/spikebike/2232102
*/

package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
)

func handleConn(from net.Conn) {
	if cert, err := tls.LoadX509KeyPair("pki/client.cert.pem", "pki/client.key.pem"); err != nil {
		log.Fatalf("%v", err)
	} else {
		config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true} // TODO: refactor with a CA file.
		if to, err := tls.Dial("tcp", ":8443", &config); err != nil {
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
}

func worker(id int, conns <-chan net.Conn) {
	log.Printf("worker %v started", id)
	for conn := range conns {
		log.Printf("worker %v got a job", id)
		handleConn(conn)
		log.Printf("worker %v done a job", id)
	}
}

func main() {
	conns := make(chan net.Conn, 10)
	for w := 1; w < 4; w++ {
		go worker(w, conns)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%v", err)
		} else {
			conns <- conn
		}
	}
}
