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

// Package route ...
package route

import (
	"net/http"

	"github.com/jinsenglin/prototype-go/pkg/http/handler"
	"github.com/jinsenglin/prototype-go/pkg/http/middleware"
)

// RegisterRoutes ...
func RegisterRoutes() {

	// equals to '^'
	http.HandleFunc("/", handler.RootHandler)

	// equals to '^/static'
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// equals to '^/users'
	http.HandleFunc("/users/", handler.UsersAPIHandler)

	// exactly '/files'
	http.HandleFunc("/files", middleware.Timed(handler.FilesAPIHandler))

	// equals to '^/spaces'
	http.HandleFunc("/spaces/", handler.SpacesAPIHandler)

	// equals to '^/channels'
	http.HandleFunc("/channels/", handler.ChannelsAPIHandler)

	// equals to '^/chats'
	http.HandleFunc("/chats/", handler.ChatsAPIHandler)

	// exactly '/chunked-response'
	http.HandleFunc("/chunked-response", handler.ChunkedHandler)

	// exactly '/sse'
	http.HandleFunc("/sse", handler.SSEServer.ServeHTTP)

	// exactly '/h2-server-push'
	http.HandleFunc("/h2-server-push", handler.H2Handler)

	// exactly '/ws'
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// TODO websocket

		// e.g.,
		/* curl --header "Connection: Upgrade" \
		--header "Upgrade: websocket" \
		--header "Host: localhost:8080" \
		--header "Origin: http://localhost:80" \
		--header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
		--header "Sec-WebSocket-Version: 13" \
		-v -X GET -L http://localhost:8080/ws/
		*/
	})

	// exactly 'login'
	http.HandleFunc("/login", handler.LoginHandler)

	// exactly 'logout'
	http.HandleFunc("/logout", handler.LogoutHandler)

	// exactly '/dummy'
	http.Handle("/dummy", middleware.WithSession(http.HandlerFunc(handler.DummyHandler)))

	// exactly '/dummy2'
	http.Handle("/dummy2", middleware.TLSAuthLogged(http.HandlerFunc(handler.DummyHandler)))

	// exactly '/dummy3'
	http.Handle("/dummy3", middleware.BasicAuthLogged(http.HandlerFunc(handler.DummyHandler)))

	// exactly '/dummy4'
	http.Handle("/dummy5", middleware.FormAuthLogged(http.HandlerFunc(handler.DummyHandler)))

	// exactly '/dummy5'
	http.Handle("/dummy4", middleware.Authed(http.HandlerFunc(handler.DummyHandler)))

}
