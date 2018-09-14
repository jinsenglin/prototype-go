package handler

import (
	"fmt"
	"net/http"
	"time"
)

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/sse/" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Page Not Found")
		} else {
			// e.g.,
			// curl -v -X GET -L http://localhost:8080/sse/

			if flusher, ok := w.(http.Flusher); ok {
				w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")
				w.Header().Set("Access-Control-Allow-Origin", "*")

				for i := 1; i <= 10; i++ {
					fmt.Fprintf(w, "Event #%d\n", i)
					flusher.Flush() // Flush the data immediatly instead of buffering it for later.
					time.Sleep(500 * time.Millisecond)
				}
			} else {
				fmt.Fprintf(w, "Server sent event feature is not supported.")
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
	}
}
