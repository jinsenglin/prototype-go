/*
Addiontional Resources
- https://gist.github.com/spikebike/2232102
- https://gobyexample.com/rate-limiting
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func handleConn(from net.Conn) {
	if cert, err := tls.LoadX509KeyPair("pki/client.cert.pem", "pki/client.key.pem"); err != nil {
		log.Fatalf("%v", err)
	} else {
		if ca, err := ioutil.ReadFile("pki/ca.cert.pem"); err != nil {
			log.Fatalf("%v", err)
		} else {
			roots := x509.NewCertPool()
			roots.AppendCertsFromPEM(ca)

			config := tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs:      roots}

			if to, err := tls.Dial("tcp", "localhost.localdomain:8443", &config); err != nil {
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
}

func worker(id int, conns <-chan net.Conn) {
	log.Printf("worker %v started", id)
	limiter := time.Tick(200 * time.Millisecond)
	for conn := range conns {
		<-limiter // TODO: refactor with a burst.
		log.Printf("worker %v got a job", id)
		handleConn(conn)
		log.Printf("worker %v done a job", id)
	}
}

func main() {
	log.Println("Try `curl -v -X GET -L http://localhost:8081`")
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
