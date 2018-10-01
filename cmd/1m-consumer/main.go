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
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/api/option"
)

const (
	demoPubsubTopic  = "echo"
	envGcpProject    = "GCP_PROJECT"
	envGcpAPIKeyFile = "GCP_KEYJSON"
)

func envParse() error {
	if os.Getenv(envGcpProject) == "" {
		return fmt.Errorf("Environment variable %s is required", envGcpProject)
	}

	if os.Getenv(envGcpAPIKeyFile) == "" {
		return fmt.Errorf("Environment variable %s is required", envGcpAPIKeyFile)
	}

	return nil
}

func getSubscription(ctx context.Context, client *pubsub.Client) (*pubsub.Subscription, error) {
	id, err := os.Hostname()

	if err != nil {
		return nil, err
	}

	subscription := client.Subscription(id)
	exist, err := subscription.Exists(ctx)

	if err != nil {
		return nil, err
	}

	if exist {
		return subscription, nil
	}

	topic := client.Topic(demoPubsubTopic)
	option := pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 60 * time.Second,
	}

	return client.CreateSubscription(ctx, id, option)
}

// Broker ...
type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

// ServeHTTP ...
func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		// Make sure that the writer supports flushing.
		//
		flusher, ok := rw.(http.Flusher)

		if !ok {
			http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "text/event-stream")
		rw.Header().Set("Cache-Control", "no-cache")
		rw.Header().Set("Connection", "keep-alive")
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		// Each connection registers its own message channel with the Broker's connections registry
		messageChan := make(chan []byte)

		// Signal the broker that we have a new connection
		broker.newClients <- messageChan

		// Remove this client from the map of connected clients
		// when this handler exits.
		defer func() {
			broker.closingClients <- messageChan
		}()

		// Listen to connection close and un-register messageChan
		notify := rw.(http.CloseNotifier).CloseNotify()

		go func() {
			<-notify
			broker.closingClients <- messageChan
		}()

		for {

			// Write to the ResponseWriter
			// Server Sent Events compatible
			fmt.Fprintf(rw, "%s\n\n", <-messageChan)

			// Flush the data immediatly instead of buffering it for later.
			flusher.Flush()
		}
	} else {
		http.Error(rw, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("BROKER | added a client . %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("BROKER | removed a client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan := range broker.clients {
				clientMessageChan <- event
			}
		}
	}

}

func (broker *Broker) prosume() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv(envGcpProject), option.WithCredentialsFile(os.Getenv(envGcpAPIKeyFile)))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	sub, err := getSubscription(ctx, client)
	if err != nil {
		log.Fatalln(err)
	}

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		log.Printf("consumed a message from Pub/Sub topic | ID %s", msg.ID)

		eventString := fmt.Sprintf("ID %s | Data %s", msg.ID, msg.Data)
		broker.Notifier <- []byte(eventString)
	})
	if err != nil {
		log.Fatalln(err)
	}

}

func sseHandler() http.Handler {
	broker := &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	// Keep prosuming - consuming messages from a Cloud Pub/Sub topic then producing them
	go broker.prosume()

	return broker
}

func main() {
	if err := envParse(); err != nil {
		log.Fatalln(err)
	}

	log.Println("HTTP SERVER | sse endpoint :8080/sse")
	http.Handle("/sse", sseHandler())

	log.Println("HTTP SERVER | prometheus metrics endpoint :8080/metrics")
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
