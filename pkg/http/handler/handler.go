package handler

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Dummy ...
func Dummy(w http.ResponseWriter, r *http.Request) {
}

// FilesAPIHandler ...
func FilesAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/files/" {
		if r.Method == http.MethodPost {
			// e.g.,
			// curl -v -X POST -L http://localhost:8080/files/ -H 'Content-Type: application/octet-stream' --data-binary '@README.md'

			_, cancelCtx := context.WithCancel(context.Background())
			defer cancelCtx()

			if file, err := ioutil.TempFile("/tmp", "upload-"); err != nil {
				log.Println(err)
			} else {
				if n, err := io.Copy(file, r.Body); err != nil {
					log.Println(err)
				} else {
					log.Printf("%d bytes are recieved. Saved as %s\n", n, file.Name())
					fmt.Fprintf(w, "%d bytes are recieved.\n", n)
				}
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method Not Allowed")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page Not Found")
	}
}
