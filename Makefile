run:
	go run -race main.go

run-cmd-1m-console:
	go build -race -o out/1m-console cmd/1m-console/*.go
	./out/1m-console -h

run-cmd-ls:
	go build -race -o out/ls cmd/ls/*.go
	./out/ls / /non-exist

run-cmd-cp:
	go build -race -o out/cp cmd/cp/*.go
	./out/cp README.md /tmp/README.md

build-cmd-linebot:
	go build -race -o out/linebot cmd/linebot/*.go

run-test-case1:
	go build -race -o out/test-case1 cmd/test/case1/*.go
	./out/test-case1

build-cmd-httpserv-debug:
	go build -race -gcflags "-N -l" -o out/httpserv-debug cmd/httpserv/main.go

run-cmd-httpserv:
	go run -race cmd/httpserv/main.go

run-cmd-httpsserv:
	go run -race cmd/httpsserv/main.go

run-cmd-httpsserv2:
	go run -race cmd/httpsserv2/main.go

run-cmd-proxyserv:
	go run -race cmd/proxyserv/main.go

run-cmd-proxyserv2:
	go run -race cmd/proxyserv2/main.go

run-cmd-proxyserv3:
	go run -race cmd/proxyserv3/main.go

run-cmd-proxyserv4:
	go run -race cmd/proxyserv4/main.go

run-cmd-proxyserv5:
	go run -race cmd/proxyserv5/main.go

run-cmd-proxyserv6:
	go run -race cmd/proxyserv6/main.go

run-exercise-loop:
	go run -race cmd/go-tour/exercise/loop-and-functions/main.go

test-exercise-loop:
	go test -race github.com/jinsenglin/prototype-go/cmd/go-tour/exercise/loop-and-functions

run-exercise-slices:
	go run -race cmd/go-tour/exercise/slices/main.go

test-exercise-slices:
	go test -race github.com/jinsenglin/prototype-go/cmd/go-tour/exercise/slices

run-exercise-maps:
	go run -race cmd/go-tour/exercise/maps/main.go

run-exercise-fibonacci-closure:
	go run -race cmd/go-tour/exercise/fibonacci-closure/main.go

run-exercise-stringers:
	go run -race cmd/go-tour/exercise/stringers/main.go

run-exercise-errors:
	go run -race cmd/go-tour/exercise/errors/main.go

run-exercise-readers:
	go run -race cmd/go-tour/exercise/readers/main.go

run-exercise-iot13reader:
	go run -race cmd/go-tour/exercise/iot13reader/main.go

run-exercise-images:
	go run -race cmd/go-tour/exercise/images/main.go

run-exercise-equivalent-binary-tree:
	go run -race cmd/go-tour/exercise/equivalent-binary-tree/main.go

run-exercise-web-crawler:
	go run -race cmd/go-tour/exercise/web-crawler/main.go

run-godoc:
	echo "open http://localhost:6060/pkg/github.com/jinsenglin/prototype-go/"
	godoc -http=:6060