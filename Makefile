GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/messaging/messaging.proto

.PHONY: build-mac
build-mac: proto
	go build -o go-micro-gorm *.go

.PHONY: build-linux
build-linux: proto
	CGO_ENABLED=0 GOOS=linux sudo go build -a -installsuffix cgo -ldflags '-w' -i -o go-micro-gorm ./main.go ./plugins.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t bradfordhamilton/go-micro-gorm:latest
	docker tag bradfordhamilton/go-micro-gorm:latest bradfordhamilton/go-micro-gorm:$(shell git rev-parse HEAD)
	docker push bradfordhamilton/go-micro-gorm:$(shell git rev-parse HEAD)
	docker rmi bradfordhamilton/go-micro-gorm:$(shell git rev-parse HEAD)
	echo 'sha for deployment.yaml: $(shell git rev-parse HEAD)'
