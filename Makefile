GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	GO_MICRO_GORM_ENV=development micro run service

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/messaging/messaging.proto

.PHONY: build-mac
build-mac: proto
	go build -o go-micro-gorm *.go

.PHONY: build-linux
build-linux: proto
	CGO_ENABLED=0 GOOS=linux go build -o go-micro-gorm ./main.go ./plugins.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build-linux
	docker build . -t bradfordhamilton/go-micro-gorm:latest
	docker tag bradfordhamilton/go-micro-gorm:latest bradfordhamilton/go-micro-gorm:$(shell git rev-parse HEAD)
	docker push bradfordhamilton/go-micro-gorm:latest
	docker push bradfordhamilton/go-micro-gorm:$(shell git rev-parse HEAD)
	docker rmi bradfordhamilton/go-micro-gorm:$(shell git rev-parse HEAD)
	echo 'sha for deployment.yaml: $(shell git rev-parse HEAD)'

.PHONY: deploy
deploy:
	GO_MICRO_GORM_SHA=$(shell git rev-parse HEAD) kubectl apply -f k8s/deployment.yaml