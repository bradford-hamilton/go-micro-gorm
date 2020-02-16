FROM ubuntu
COPY cmd/go-micro-gorm/go-micro-gorm /go-micro-gorm
ENTRYPOINT [ "/go-micro-gorm", "--registry=etcd" ]
