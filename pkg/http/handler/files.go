package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// FilesAPIHandler ...
func FilesAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// e.g.,
		// curl -X POST http://localhost:8080/files -H 'Content-Type: application/octet-stream' --data-binary '@README.md'

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
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
