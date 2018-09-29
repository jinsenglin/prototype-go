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
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	demoPubsubTopic      = "xxxx-5678"
	demoContainerCluster = "xxxx-5678"
	envGcpProject        = "GCP_PROJECT" // TODO: remove
	subcommandUp         = "up"
	subcommandMore       = "more"
	subcommandDown       = "down"
)

var (
	n int
)

const environmentUsage = "\n" +
	`ENVIRONMENT:` + "\n" +
	`  %-12s GCP project name to use for this demo. Required` + "\n"

const subcommandUsage = "\n" +
	`SUBCOMMAND:` + "\n" +
	`  %-12s launch a GKE cluster` + "\n" +
	`  %-12s resize GKE cluster by adding one more node` + "\n" +
	`  %-12s shutdown GKE cluster` + "\n"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [SUBCOMMAND] [OPTIONS]\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, environmentUsage, envGcpProject)
	fmt.Fprintf(os.Stderr, subcommandUsage, subcommandUp, subcommandMore, subcommandDown)
}

func init() {
	flag.IntVar(&n, "num", 1, "number of consumer workers") // TODO: remove
	flag.Usage = usage
}

func envParse() error {
	if os.Getenv(envGcpProject) == "" {
		return fmt.Errorf("Environment variable %s is required", envGcpProject)
	}

	return nil
}

func execCommandStart(stdin []byte, name string, arg ...string) error {
	log.Println(name, arg)

	cmd := exec.Command(name, arg...)
	cmdStdin, _ := cmd.StdinPipe()
	cmdStdout, _ := cmd.StdoutPipe()
	cmdStderr, _ := cmd.StderrPipe()
	cmd.Start()
	cmdStdin.Write(stdin)
	cmdStdoutBytes, _ := ioutil.ReadAll(cmdStdout)
	cmdStderrBytes, _ := ioutil.ReadAll(cmdStderr)
	log.Printf("STDOUT\n%s", cmdStdoutBytes)
	log.Printf("STDERR\n%s", cmdStderrBytes)

	err := cmd.Wait()
	return err
}

func gcloudPubsubTopicsDelete() error {
	return execCommandStart([]byte{}, "gcloud", "pubsub", "topics", "delete", demoPubsubTopic)
}

func gcloudPubsubTopicsCreate() error {
	return execCommandStart([]byte{}, "gcloud", "pubsub", "topics", "create", demoPubsubTopic)
}

func gcloudContainerClustersDelete() error {
	return execCommandStart([]byte("Y\n"), "gcloud", "container", "clusters", "delete", demoContainerCluster)
}

func gcloudContainerClustersCreate() error {
	cmdTemplate := `container clusters create %s --num-nodes 1`
	cmdString := fmt.Sprintf(cmdTemplate, demoContainerCluster)
	cmdArg := strings.Split(cmdString, " ")
	return execCommandStart([]byte{}, "gcloud", cmdArg...)

	// NOTE
	// spend 2 mins 20 secs for 1 node
	// auto-config kubeconfig after cluster creation
}

func kubectlApply() error {
	// TODO
	return nil
}

func up() {
	// Create a Cloud Pub/Sub topic
	// No sdk to do this. Use gcloud command-line tool instead.
	if err := gcloudPubsubTopicsCreate(); err != nil {
		log.Fatalln(err)
	}

	// Launch a GKE cluster
	// No sdk to do this. Use gcloud command-line tool instead.
	if err := gcloudContainerClustersCreate(); err != nil {
		log.Fatalln(err)
	}

	// Deploy workload to GKE cluster
	// Use kubectl command-line tool, which is an easier way compared to client-go library.
	if err := kubectlApply(); err != nil {
		log.Fatalln(err)
	}
}

func more() {
	// TODO: resize GKE cluster by adding one more node
	// TODO: reshape workload by adding more consumer pods
}

func down() {
	// Delete Cloud Pub/Sub topic
	// No sdk to do this. Use gcloud command-line tool instead.
	if err := gcloudPubsubTopicsDelete(); err != nil {
		log.Fatalln(err)
	}

	// Shutdown GKE cluster
	// No sdk to do this. Use gcloud command-line tool instead.
	if err := gcloudContainerClustersDelete(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	flag.Parse()
	if err := envParse(); err != nil {
		log.Fatalln(err)
	}

	switch len(os.Args) {
	case 1:
		usage()
	default:
		switch os.Args[1] {
		case subcommandUp:
			up()
		case subcommandMore:
			more()
		case subcommandDown:
			down()
		default:
			usage()
		}
	}
}
