run:
	go run -race main.go

run-cmd-1m-console:
	go build -race -o out/1m-console cmd/1m-console/*.go
	./out/1m-console

run-cmd-1m-client:
	go build -race -o out/1m-client cmd/1m-client/*.go
	./out/1m-client

run-cmd-1m-consumer:
	go build -race -o out/1m-consumer cmd/1m-consumer/*.go
	./out/1m-consumer

run-cmd-1m-producer:
	go build -race -o out/1m-producer cmd/1m-producer/*.go
	./out/1m-producer

build-image-1m-client:
	GOOS=linux GOARCH=amd64 go build -o out/1m-client-linux-amd64 cmd/1m-client/*.go
	docker build -f dockerfile/1m-client/Dockerfile -t jinsenglin/1m-client:latest .

build-image-1m-consumer:
	GOOS=linux GOARCH=amd64 go build -o out/1m-consumer-linux-amd64 cmd/1m-consumer/*.go
	docker build -f dockerfile/1m-consumer/Dockerfile -t jinsenglin/1m-consumer:latest .

build-image-1m-producer:
	GOOS=linux GOARCH=amd64 go build -o out/1m-producer-linux-amd64 cmd/1m-producer/*.go
	docker build -f dockerfile/1m-producer/Dockerfile -t jinsenglin/1m-producer:latest .

run-image-1m-client:
	docker run --rm --name 1m-client jinsenglin/1m-client:latest /app -url http://172.17.0.2:8082/sse

run-image-1m-consumer:
	docker run --rm -P -v ${GCP_KEYJSON}:${GCP_KEYJSON} -e GCP_PROJECT=${GCP_PROJECT} -e GCP_KEYJSON=${GCP_KEYJSON} -h onem-consumer --name 1m-consumer jinsenglin/1m-consumer:latest

run-image-1m-producer:
	docker run --rm -P -v ${GCP_KEYJSON}:${GCP_KEYJSON} -e GCP_PROJECT=${GCP_PROJECT} -e GCP_KEYJSON=${GCP_KEYJSON} -h onem-producer --name 1m-producer jinsenglin/1m-producer:latest

run-image-1m-client-in-k8s-by-docker-for-mac:
	helm install --name onem-client k8s/DockerForMac/onem-client

	# CLEANUP
	# helm delete --purge onem-client

run-image-1m-consumer-in-k8s-by-docker-for-mac:
	kubectl create secret generic key-json --from-file=${GCP_KEYJSON}
	kubectl create configmap gcp-project --from-literal=gcp-project-id=${GCP_PROJECT}
	helm install --name onem-consumer k8s/DockerForMac/onem-consumer
	
	# CLEANUP
	# helm delete --purge onem-consumer
	# kubectl delete configmap gcp-project
	# kubectl delete secret key-json

run-image-1m-producer-in-k8s-by-docker-for-mac:
	kubectl create secret generic key-json --from-file=${GCP_KEYJSON}
	kubectl create configmap gcp-project --from-literal=gcp-project-id=${GCP_PROJECT}
	helm install --name onem-producer k8s/DockerForMac/onem-producer

	# CLEANUP
	# helm delete --purge onem-producer
	# kubectl delete configmap gcp-project
	# kubectl delete secret key-json

push-image-1m-client-to-gcp:
	docker tag jinsenglin/1m-client:latest asia.gcr.io/${GCP_PROJECT}/1m-client:latest
	gcloud auth configure-docker
	docker push asia.gcr.io/${GCP_PROJECT}/1m-client:latest

push-image-1m-consumer-to-gcp:
	docker tag jinsenglin/1m-consumer:latest asia.gcr.io/${GCP_PROJECT}/1m-consumer:latest
	gcloud auth configure-docker
	docker push asia.gcr.io/${GCP_PROJECT}/1m-consumer:latest

push-image-1m-producer-to-gcp:
	docker tag jinsenglin/1m-producer:latest asia.gcr.io/${GCP_PROJECT}/1m-producer:latest
	gcloud auth configure-docker
	docker push asia.gcr.io/${GCP_PROJECT}/1m-producer:latest

up-gke-stage:
	# SPEC
	# allocatable:
    #   cpu: 940m
    #   memory: 2708916Ki
    #   pods: "110"
    # capacity:
    #   cpu: "1"
    #   memory: 3794356Ki
    #   pods: "110"
	
	gcloud container clusters create k8s-1m --num-nodes 1
	kubectl apply -f k8s/GKE/service-account-helm.yaml
	
	# NOTE: keep watching until all system pods are running
	# watch -n 5 'kubectl get po -n kube-system'

	helm init --service-account helm

	# NOTE: keep watching until the system tiller deploy is available
	# watch -n 5 'kubectl get deploy tiller-deploy -n kube-system'

	# CLEANUP
	# gcloud container clusters delete k8s-1m

up-gke-prod:
	echo "TODO"

run-image-1m-client-in-k8s-by-gke:
	helm install --name onem-client k8s/GKE/onem-client --set image.repository=asia.gcr.io/${GCP_PROJECT}/1m-client
	
	# CLEANUP STEPS
	# helm delete --purge onem-client

run-image-1m-consumer-in-k8s-by-gke:
	kubectl create secret generic key-json --from-file=${GCP_KEYJSON}
	kubectl create configmap gcp-project --from-literal=gcp-project-id=${GCP_PROJECT}
	helm install --name onem-consumer k8s/GKE/onem-consumer --set image.repository=asia.gcr.io/${GCP_PROJECT}/1m-consumer
	
	# CLEANUP
	# helm delete --purge onem-consumer
	# kubectl delete configmap gcp-project
	# kubectl delete secret key-json

run-image-1m-producer-in-k8s-by-gke:
	kubectl create secret generic key-json --from-file=${GCP_KEYJSON}
	kubectl create configmap gcp-project --from-literal=gcp-project-id=${GCP_PROJECT}
	helm install --name onem-producer k8s/GKE/onem-producer --set image.repository=asia.gcr.io/${GCP_PROJECT}/1m-producer
	
	# CLEANUP STEPS
	# helm delete --purge onem-producer
	# kubectl delete configmap gcp-project
	# kubectl delete secret key-json

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