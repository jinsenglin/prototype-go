//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package handler

import (
	"fmt"
	"net/http"
	"time"
)

func ChunkedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// e.g.,
		// curl http://localhost:8080/chunked-response

		if flusher, ok := w.(http.Flusher); ok {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			for i := 1; i <= 10; i++ {
				fmt.Fprintf(w, "Chunk #%d\n", i)
				flusher.Flush() // Trigger "chunked" encoding and send a chunk...
				time.Sleep(500 * time.Millisecond)
			}
		} else {
			http.Error(w, "Chunked response feature is not supported.", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
