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
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func subscribeIfNotExits(client *pubsub.Client) error {
	// TODO
	return nil
}

func main() {
	// TODO: get project name from environment variable
	// TODO: get credential file path from environment variable
	pubsubClient, err := pubsub.NewClient(context.Background(), "k8s-project-199813", option.WithCredentialsFile("/Users/cclin/key.json"))
	if err != nil {
		log.Fatalln(err)
	}

	// Create a subscription if not exits
	if err := subscribeIfNotExits(pubsubClient); err != nil {
		log.Fatalln(err)
	}

	// TODO: keep consuming
	for {
		time.Sleep(1 * time.Second)
		log.Println("consumed a message")
	}
}
