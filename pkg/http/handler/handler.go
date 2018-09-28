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
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// DummyHandler ...
func DummyHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/dummy

	fmt.Fprintf(w, "Dummy")
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/login

	// Set authed Cookie
	sid := strconv.Itoa(rand.Int())
	expiration := time.Now().Add(365 * 24 * time.Hour)
	c := http.Cookie{Name: "authed", Value: sid, Expires: expiration, Path: "/"}
	http.SetCookie(w, &c)

	fmt.Fprintf(w, "Logined")
}

// LogoutHandler ...
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// e.g.,
	// curl http://localhost:8080/logout

	// Remove authed Cookie
	sid := strconv.Itoa(rand.Int())
	expiration := time.Now().Add(-365 * 24 * time.Hour)
	c := http.Cookie{Name: "authed", Value: sid, Expires: expiration, Path: "/"}
	http.SetCookie(w, &c)

	fmt.Fprintf(w, "Logouted")
}
