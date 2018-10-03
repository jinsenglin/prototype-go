package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	flagSSEURL        string
	flagVirtualClient int
)

func init() {
	flag.StringVar(&flagSSEURL, "url", "http://localhost:8082/sse", "SSE server url")
	flag.IntVar(&flagVirtualClient, "client", 1, "number of virtual clients")
}

// Event ...
type Event struct {
	Name string
	ID   string
	Data []byte
}

func watch() (events chan Event, err error) {

	resp, err := http.Get(flagSSEURL)

	// Block until a first message responsed

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

func virtualClient(clientID string) {
	log.Printf("CLIENT %s | receiving ...", clientID)

	events, err := watch()
	if err != nil {
		log.Fatalln(err)
	}

	for event := range events {
		t := time.Now()
		log.Printf("CLIENT %s | %s | received a message | DATA %s", clientID, t.Format(time.RFC3339), event.Data)
	}
}

func main() {
	flag.Parse()

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < flagVirtualClient; i++ {
		clientID := fmt.Sprintf("%s-%d", hostname, i)
		go virtualClient(clientID)
	}

	log.Println("HTTP SERVER | prometheus metrics endpoint :8083/metrics")
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8083", nil))
}
