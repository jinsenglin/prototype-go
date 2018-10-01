package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Event ...
type Event struct {
	Name string
	ID   string
	Data []byte
}

func watch(endpoint string) (events chan Event, err error) {

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("sse http status code %d", resp.StatusCode)
	}

	events = make(chan Event)

	go func() {
		ev := Event{}

		reader := bufio.NewReader(resp.Body)

		for {
			line, err := reader.ReadBytes('\n')

			if err != nil {
				log.Fatalln(err)
			}

			switch {
			case bytes.HasPrefix(line, []byte("data:")):
				ev.Data = line
			case bytes.Equal(line, []byte("\n")):
				events <- ev
				ev = Event{}
			default:
				log.Fatalln("GG")
			}
		}
	}()

	return events, nil
}

func virtualClient(endpoint string) {
	events, err := watch(endpoint)
	if err != nil {
		log.Fatalln(err)

	}

	for event := range events {
		log.Printf("CLIENT | received a message | NAME %s | DATA %s", event.Name, event.Data)
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Usage: 1m-client <sse server endpoint>")
	}

	// TODO: simulate 10 virtual clients
	sseServerEndpoint := os.Args[1]
	for i := 0; i < 1; i++ {
		go virtualClient(sseServerEndpoint)
	}

	log.Println("HTTP SERVER | prometheus metrics endpoint :8080/metrics")
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
