package messagingsubscriber

import (
	"context"

	pb "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

// Messaging defines the struct to which we attach different handlers to
type Messaging struct {
	DB *gorm.DB
}

// Handle message publish
func (m *Messaging) Handle(ctx context.Context, msg *pb.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

// Handler that just prints our message
func Handler(ctx context.Context, msg *pb.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
