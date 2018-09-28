package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	n int
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [SUBCOMMAND] [OPTIONS]\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "ENVIRONMENT:")
	fmt.Fprintln(os.Stderr, " HTTP_PROXY proxy for HTTP requests; complete URL or HOST[:PORT]") // TODO: remove
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "SUBCOMMAND:")
	fmt.Fprintln(os.Stderr, " up   launch a GKE cluster")
	fmt.Fprintln(os.Stderr, " more resize GKE cluster by adding one more node")
	fmt.Fprintln(os.Stderr, " down shutdown GKE cluster")
}

func init() {
	flag.IntVar(&n, "num", 1, "number of consumer workers") // TODO: remove
	flag.Usage = usage
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

	switch os.Args[1] {
	case "up":
		up()
	case "more":
		more()
	case "down":
		down()
	default:
		os.Exit(1)
	}
}
