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

	// exactly '/dummy'
	http.Handle("/dummy", middleware.WithSession(http.HandlerFunc(handler.DummyHandler)))

}
