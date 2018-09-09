/*
Additional Resources
- https://golang.org/pkg/crypto/tls/
- https://github.com/denji/golang-tls
- http://www.hydrogen18.com/blog/your-own-pki-tls-golang.html
- http://www.bite-code.com/2015/06/25/tls-mutual-auth-in-golang/
*/

package main

import (
	"log"
	"net/http"

	"github.com/jinsenglin/prototype-go/pkg/http/route"
)

func main() {
	log.Println("Try `curl -v -X GET -L -k https://localhost:8443`")
	route.RegisterRoutes()
	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)) // TODO: refactor with TLS mutual authN. See httpsserv2
}
