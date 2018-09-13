package middleware

import (
	"context"
	"log"
	"net/http"
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
			expiration := time.Now().Add(365 * 24 * time.Hour)
			sid := http.Cookie{Name: "sid", Value: "1", Expires: expiration} // TODO: refactor with a random session ID.
			http.SetCookie(w, &sid)
			ctx = reqcontext.SetSession(r.Context(), session.Session{ID: "1"})
		} else {
			ctx = reqcontext.SetSession(r.Context(), session.Session{ID: cookie.Value})
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
