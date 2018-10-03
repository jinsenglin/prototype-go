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
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/api/option"
)

var (
	flagInterval time.Duration
)

func init() {
	flag.DurationVar(&flagInterval, "interval", 5*time.Second, "interval of publishing a message")
}

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

func publish() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv(envGcpProject), option.WithCredentialsFile(os.Getenv(envGcpAPIKeyFile)))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	t := client.Topic(demoPubsubTopic)
	log.Println("PUBLISHER | publishing ...")
	for {
		time.Sleep(flagInterval)

		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("msg"),
		})

		// Block until the result is returned and a server-generated
		// ID is returned for the published message.
		id, err := result.Get(ctx)

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("PUBLISHER | published a message to Pub/Sub topic | ID %s", id)
	}
}

func main() {
	flag.Parse()

	if err := envParse(); err != nil {
		log.Fatalln(err)
	}

	// Keep publishing
	go publish()

	log.Println("HTTP SERVER | prometheus metrics endpoint :8081/metrics")
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8081", nil))
}
