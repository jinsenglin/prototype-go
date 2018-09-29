package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	demoPubsubTopic      = "xxxx-5678"
	demoContainerCluster = "xxxx-5678"
	envGcpProject        = "GCP_PROJECT"
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
	cmd := exec.Command(name, arg...)
	cmdStdin, _ := cmd.StdinPipe()
	cmdStdout, _ := cmd.StdoutPipe()
	cmdStderr, _ := cmd.StderrPipe()
	cmd.Start()
	cmdStdin.Write(stdin)
	cmdStdoutBytes, _ := ioutil.ReadAll(cmdStdout)
	cmdStderrBytes, _ := ioutil.ReadAll(cmdStderr)
	log.Println(name, arg)
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
	return execCommandStart([]byte{}, "gcloud", "container", "clusters", "create", demoContainerCluster)
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

	// TODO: deploy workload to GKE cluster
	// No sdk to do this. Use kubectl command-line tool instead.
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
