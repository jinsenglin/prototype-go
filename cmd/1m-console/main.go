package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	n int
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] URL\n\n", os.Args[0]) // TODO: remove
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "ENVIRONMENT:")
	fmt.Fprintln(os.Stderr, " HTTP_PROXY proxy for HTTP requests; complete URL or HOST[:PORT]") // TODO: remove
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "SUBCOMMAND:")
	fmt.Fprintln(os.Stderr, " launch ...") // TODO: remove
}

func init() {
	flag.IntVar(&n, "num", 1, "number of consumer workers")
	flag.NewFlagSet("launch", flag.ExitOnError)
	flag.Usage = usage
}

func main() {
	flag.Parse()

	log.Println("ok")
}
