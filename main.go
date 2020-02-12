package main

import (
	"github.com/bradford-hamilton/go-micro-gorm/handler"
	proto_messaging "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	"github.com/bradford-hamilton/go-micro-gorm/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// Create new micro service
	service := micro.NewService(
		micro.Name("go.micro.srv.messaging"),
		micro.Version("latest"),
	)

	// Initialize the service
	service.Init()

	// Register Handler
	proto_messaging.RegisterMessagingHandler(service.Server(), new(handler.Messaging))

	// Register Messaging struct as subscriber
	micro.RegisterSubscriber("go.micro.srv.messaging", service.Server(), new(subscriber.Messaging))

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
