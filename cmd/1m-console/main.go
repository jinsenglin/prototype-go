package main

import (
	"flag"
	"fmt"
	"os"
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

func up() {
	// TODO: launch a GKE cluster
	// No sdk to do this. Use gcloud command-line tool instead.

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
