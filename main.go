package main

import (
	"github.com/bradford-hamilton/go-micro-gorm/interal/db"
	messaginghandler "github.com/bradford-hamilton/go-micro-gorm/internal/handlers/messaging_handler"
	messagingsubscriber "github.com/bradford-hamilton/go-micro-gorm/internal/subscribers/messaging_subscriber"
	proto_messaging "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create new micro service
	service := micro.NewService(
		micro.Name("go.micro.srv.messaging"),
		micro.Version("latest"),
	)

	// Initialize the service
	service.Init()

	// Register Handler
	proto_messaging.RegisterMessagingServiceHandler(
		service.Server(),
		&messaginghandler.Messaging{Db: db},
	)

	// Register Messaging struct as subscriber
	micro.RegisterSubscriber(
		"go.micro.srv.messaging",
		service.Server(),
		new(messagingsubscriber.Messaging),
	)

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
