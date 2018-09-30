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
		AckDeadline: 10 * time.Second, // TODO
	}

	return client.CreateSubscription(ctx, id, option)
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

	sub, err := getSubscription(ctx, client)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: keep consuming
	if err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("consumed a message %s: %q\n", msg.ID, string(msg.Data))
	}); err != nil {
		log.Fatalln(err)
	}
}
