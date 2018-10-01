package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	flagSSEURL string
)

func init() {
	flag.StringVar(&flagSSEURL, "url", "http://localhost:8082/sse", "SSE server url")
}

// Event ...
type Event struct {
	Name string
	ID   string
	Data []byte
}

func watch() (events chan Event, err error) {

	resp, err := http.Get(flagSSEURL)

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
				log.Fatalf("FATAL | LINE %s", line)
			}
		}
	}()

	return events, nil
}

func virtualClient() {
	events, err := watch()
	if err != nil {
		log.Fatalln(err)

	}

	for event := range events {
		log.Printf("CLIENT | received a message | DATA %s", event.Data)
	}
}

func main() {
	// TODO: simulate 10 virtual clients
	for i := 0; i < 1; i++ {
		go virtualClient()
	}

	log.Println("HTTP SERVER | prometheus metrics endpoint :8083/metrics")
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8083", nil))
}
