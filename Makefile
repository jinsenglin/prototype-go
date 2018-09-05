run:
	go run main.go

run-cmd-ls:
	go build -o out/ls cmd/ls/*.go
	./out/ls / /non-exist

run-cmd-httpserv:
	go run cmd/httpserv/main.go

run-cmd-sqlclient:
	go run cmd/sqlclient/main.go