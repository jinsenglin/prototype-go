package handler

import (
	"log"
	"net/http"
)

// DummyHandler ...
func DummyHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/dummy

	log.Println(r.Context())
}
