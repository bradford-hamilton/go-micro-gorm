FROM alpine
ADD go-micro-gorm /go-micro-gorm
ENTRYPOINT [ "/go-micro-gorm" ]
