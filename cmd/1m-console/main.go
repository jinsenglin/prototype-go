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
	envGcpProject  = "GCP_PROJECT"
	subcommandUp   = "up"
	subcommandMore = "more"
	subcommandDown = "down"
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

func gcloudContainerClustersCreate() error {
	cmd := exec.Command("gcloud", "container", "clusters", "create") // TODO
	cmdStdout, _ := cmd.StdoutPipe()
	cmdStderr, _ := cmd.StderrPipe()
	cmd.Start()
	cmdStdoutBytes, _ := ioutil.ReadAll(cmdStdout)
	cmdStderrBytes, _ := ioutil.ReadAll(cmdStderr)
	log.Println("gcloud container clusters create")
	log.Printf("STDOUT %s", cmdStdoutBytes)
	log.Printf("STDERR %s", cmdStderrBytes)
	err := cmd.Wait()
	return err
}

func up() {
	// TODO: create a Cloud Pub/Sub topic

	// Launch a GKE cluster
	// No sdk to do this. Use gcloud command-line tool instead.
	if err := gcloudContainerClustersCreate(); err != nil {
		log.Fatalln(err)
	}

	// TODO: wait GKE cluster ready to use

	// TODO: deploy workload to GKE cluster
	// No sdk to do this. Use kubectl command-line tool instead.
}

func more() {
	// TODO: resize GKE cluster by adding one more node
	// TODO: reshape workload by adding more consumer pods
}

func down() {
	// TODO: shutdown GKE cluster
	// TODO: delete Cloud Pub/Sub topic
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
