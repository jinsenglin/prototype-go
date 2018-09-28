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
	for conn := range conns { // TODO: refactor with rate limiting. See proxyserv5
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
