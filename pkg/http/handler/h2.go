package handler

import (
	"fmt"
	"log"
	"net/http"
)

const indexHTML = `<html>
<head>
	<title>Demo HTTP/2 Server Push</title>
	<script src="/static/app.js"></script>
	<link rel="stylesheet" href="/static/style.css"">
</head>
<body>
Demo HTTP/2 Server Push
</body>
</html>
`

func H2Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/h2-server-push/" {
		if r.Method == http.MethodGet {
			// NOTE: MUST USE HTTPS e.g.,
			// curl --http2 -v -X GET -L -k https://localhost:8443/h2-server-push/

			if pusher, ok := w.(http.Pusher); ok {
				if err := pusher.Push("/static/app.js", nil); err != nil {
					log.Printf("Failed to push: %v", err)

				}
				if err := pusher.Push("/static/style.css", nil); err != nil {
					log.Printf("Failed to push: %v", err)
				}

				// WARNING: Failed to push when using curl

				fmt.Fprintf(w, indexHTML)
			} else {
				http.Error(w, "HTTP/2 server push feature is not supported.", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.NotFound(w, r)
	}
}
