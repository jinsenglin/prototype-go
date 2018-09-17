/*
Additional Resources
- https://golang.org/pkg/net/http/
- https://golang.org/pkg/html/template/
- https://golang.org/doc/articles/wiki/
- https://gowebexamples.com/http-server/
- http://legendtkl.com/2016/08/21/go-web-server/
- http://fuxiaohei.me/2016/9/24/go-and-fasthttp-server.html
*/

package main

import (
	"log"
	"net/http"

	"github.com/jinsenglin/prototype-go/pkg/http/route"
)

func main() {
	log.Println("Try `curl http://localhost:8080`")
	route.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil)) // TODO: refactor with https. See httpsserv
}
