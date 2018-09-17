/*
Additional Resources
- https://golang.org/pkg/crypto/tls/
- https://github.com/denji/golang-tls
- http://www.hydrogen18.com/blog/your-own-pki-tls-golang.html
- http://www.bite-code.com/2015/06/25/tls-mutual-auth-in-golang/
- https://github.com/cclin81922/tls
- https://github.com/cclin81922/pki
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinsenglin/prototype-go/pkg/http/route"
)

func main() {
	log.Println("Try `curl -L --cert pki/client.cert.pem --key pki/client.key.pem --cacert pki/ca.cert.pem https://localhost.localdomain:8443`")
	route.RegisterRoutes()

	if caCert, err := ioutil.ReadFile("pki/ca.cert.pem"); err != nil {
		log.Fatal(err)
	} else {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		tlsConfig.BuildNameToCertificate()

		server := &http.Server{
			Addr:      ":8443",
			TLSConfig: tlsConfig,
		}

		log.Fatal(server.ListenAndServeTLS("pki/server.cert.pem", "pki/server.key.pem"))
	}
}
