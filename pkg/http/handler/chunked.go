package handler

import (
	"fmt"
	"net/http"
	"time"
)

func ChunkedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/chunked-response/" {
			http.NotFound(w, r)
		} else {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/chunked-response/

			if flusher, ok := w.(http.Flusher); ok {
				w.Header().Set("X-Content-Type-Options", "nosniff")
				for i := 1; i <= 10; i++ {
					fmt.Fprintf(w, "Chunk #%d\n", i)
					flusher.Flush() // Trigger "chunked" encoding and send a chunk...
					time.Sleep(500 * time.Millisecond)
				}
			} else {
				fmt.Fprintf(w, "Chunked response feature is not supported.")
			}
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
