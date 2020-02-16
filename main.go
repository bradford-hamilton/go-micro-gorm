package main

import (
	"github.com/bradford-hamilton/go-micro-gorm/internal/db"
	msginghandler "github.com/bradford-hamilton/go-micro-gorm/internal/handlers/messaging_handler"
	userhandler "github.com/bradford-hamilton/go-micro-gorm/internal/handlers/user_handler"
	msgingsubscriber "github.com/bradford-hamilton/go-micro-gorm/internal/subscribers/messaging_subscriber"
	msgpb "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	userpb "github.com/bradford-hamilton/go-micro-gorm/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

const serviceName = "go.micro.srv.webapp"

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create new micro service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
	)

	// Initialize the service
	service.Init()

	// Register and attach the db to the handler for messaging service
	msgHandler := msginghandler.Messaging{DB: db}
	err = msgpb.RegisterMessagingServiceHandler(service.Server(), &msgHandler)
	if err != nil {
		log.Fatal(err)
	}

	// Reigster and attach the db to the handler for user service
	userHandler := userhandler.User{DB: db}
	err = userpb.RegisterUserServiceHandler(service.Server(), &userHandler)
	if err != nil {
		log.Fatal(err)
	}

	// Register and attach the db to to the subscriber for msg subscription service
	msgSubscriber := msgingsubscriber.Messaging{DB: db}
	err = micro.RegisterSubscriber(serviceName, service.Server(), &msgSubscriber)
	if err != nil {
		log.Fatal(err)
	}

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
