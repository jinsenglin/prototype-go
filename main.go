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

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func main() {
	/*
		gcloud iam service-accounts create xxxx-5678 --display-name xxxx-5678
		gcloud iam service-accounts keys create key.json --iam-account=xxxx-5678@k8s-project-199813.iam.gserviceaccount.com
		gcloud projects add-iam-policy-binding k8s-project-199813 --member=serviceAccount:xxxx-5678@k8s-project-199813.iam.gserviceaccount.com --role=roles/pubsub.admin
	*/

	pubsubClient, err := pubsub.NewClient(context.Background(), "k8s-project-199813", option.WithCredentialsFile("/Users/cclin/key.json"))
	if err != nil {
		log.Fatalln(err)
	}

	if topic, err := pubsubClient.CreateTopic(context.Background(), "topic-name"); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(topic)
	}
}
