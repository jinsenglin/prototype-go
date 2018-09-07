run:
	go run main.go

run-cmd-ls:
	go build -o out/ls cmd/ls/*.go
	./out/ls / /non-exist

run-cmd-httpserv:
	go run cmd/httpserv/main.go

run-cmd-proxyserv:
	go run cmd/proxyserv/main.go

run-cmd-proxyserv2:
	go run cmd/proxyserv2/main.go

run-cmd-proxyserv3:
	go run cmd/proxyserv3/main.go

run-cmd-sqlclient:
	go run cmd/sqlclient/main.go