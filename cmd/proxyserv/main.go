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
