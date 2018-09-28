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

// Package middleware ...
package middleware

import (
	"context"
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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

// Authed ...
func Authed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cookie-based auth :: Check authed Cookie
		if cookie, err := r.Cookie("authed"); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			// TODO: more check against fake cookie.
			next.ServeHTTP(w, r)
			log.Printf("cookie.Value = %v", cookie.Value)
		}

		// TODO: token-based auth :: Check authed Header
	})
}

// BasicAuthLogged ...
func BasicAuthLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// e.g.,
		// curl -u user:pass ...

		auth := r.Header.Get("Authorization")
		basic_auth := strings.SplitN(auth, " ", 2)
		payload, _ := base64.StdEncoding.DecodeString(basic_auth[1])
		user_pass := strings.SplitN(string(payload), ":", 2)
		user := user_pass[0]
		pass := user_pass[1]
		log.Printf("Basic Auth user: %s | pass: %s", user, pass)

		next.ServeHTTP(w, r)
	})
}

// FormAuthLogged ...
func FormAuthLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// e.g.,
		// curl -F 'user=user' -F 'pass=pass' ...

		user := r.FormValue("user")
		pass := r.FormValue("pass")
		log.Printf("Form Auth user: %s | pass: %s", user, pass)

		next.ServeHTTP(w, r)
	})
}

// TLSAuthLogged ...
func TLSAuthLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		certs := r.TLS.PeerCertificates
		for _, cert := range certs {
			log.Printf("TLS Auth common name: %s", cert.Subject.CommonName)
		}

		next.ServeHTTP(w, r)
	})
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
