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
	"os"
	"time"

	"cloud.google.com/go/pubsub"
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

func main() {
	if err := envParse(); err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv(envGcpProject), option.WithCredentialsFile(os.Getenv(envGcpAPIKeyFile)))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// TODO: keep producing
	t := client.Topic(demoPubsubTopic)
	for {
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("msg"),
		})

		// Block until the result is returned and a server-generated
		// ID is returned for the published message.
		id, err := result.Get(ctx)

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("produced a message; msg ID: %v\n", id)
		time.Sleep(3 * time.Second)
	}
}
