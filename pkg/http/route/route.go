/*
Package route ...

Implementations of http file uploader client and server
See https://gist.github.com/ebraminio/576fdfdff425bf3335b51a191a65dbdb

Copy a struct instance
See https://flaviocopes.com/go-copying-structs/

Convert an array to a slice
See https://stackoverflow.com/questions/28886616/convert-array-to-slice-in-go

Deep copy a struct instance
- https://gobyexample.com/json
- https://blog.golang.org/json-and-go
- https://www.jianshu.com/p/f1cdb1bc1b74

With context
See https://blog.golang.org/context

HTTP/2 Server Push
See https://blog.golang.org/h2push
*/
package route

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/jinsenglin/prototype-go/pkg/http/handler"
	"github.com/jinsenglin/prototype-go/pkg/http/middleware"
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func _idx(path string) int {
	re, _ := regexp.Compile("[1-9]")
	id, _ := strconv.Atoi(re.FindString(path))
	idx := id - 1
	return idx
}

var data = model.Users{}

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

	http.HandleFunc("/h2/", handler.H2Handler)

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
