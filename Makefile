run:
	go run main.go

run-cmd-ls:
	go build -o out/ls cmd/ls/*.go
	./out/ls / /non-exist

run-cmd-cp:
	go build -o out/cp cmd/cp/*.go
	./out/cp README.md /tmp/README.md

run-cmd-httpserv:
	go run cmd/httpserv/main.go

run-cmd-httpsserv:
	go run cmd/httpsserv/main.go

run-cmd-httpsserv2:
	go run cmd/httpsserv2/main.go

run-cmd-proxyserv:
	go run cmd/proxyserv/main.go

run-cmd-proxyserv2:
	go run cmd/proxyserv2/main.go

run-cmd-proxyserv3:
	go run cmd/proxyserv3/main.go

run-cmd-proxyserv4:
	go run cmd/proxyserv4/main.go