package route

import (
	"net/http"

	"github.com/jinsenglin/prototype-go/pkg/http/handler"
	"github.com/jinsenglin/prototype-go/pkg/http/middleware"
)

// RegisterRoutes ...
func RegisterRoutes() {

	http.HandleFunc("/", handler.RootHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/users/", handler.UsersAPIHandler)

	http.HandleFunc("/files/", middleware.Timed(handler.FilesAPIHandler))

	http.HandleFunc("/channels/", handler.ChannelsAPIHandler)

	http.HandleFunc("/chats/", handler.ChatsAPIHandler)

	http.HandleFunc("/chunked-response/", handler.ChunkedHandler)

	http.HandleFunc("/sse/", handler.SSEHandler)

	http.HandleFunc("/h2-server-push/", handler.H2Handler)

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
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

	http.Handle("/dummy/", middleware.WithSession(http.HandlerFunc(handler.DummyHandler)))

}
