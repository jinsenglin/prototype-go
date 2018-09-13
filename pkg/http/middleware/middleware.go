package middleware

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jinsenglin/prototype-go/pkg/http/context"
	"github.com/jinsenglin/prototype-go/pkg/http/session"
)

// Dummy ...
func Dummy(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

// Timed ...
func Timed(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		end := time.Now()
		log.Println("The request took", end.Sub(start))
	}
}

// WithSession ...
func WithSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		// e.g.,
		// curl --cookie "sid=1"

		if cookie, err := r.Cookie("sid"); err != nil {
			sid := strconv.Itoa(rand.Int())
			expiration := time.Now().Add(365 * 24 * time.Hour)
			c := http.Cookie{Name: "sid", Value: sid, Expires: expiration, Path: "/"}

			http.SetCookie(w, &c)

			ctx = reqcontext.SetSession(r.Context(), session.Session{ID: sid})
		} else {
			ctx = reqcontext.SetSession(r.Context(), session.Session{ID: cookie.Value})
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
