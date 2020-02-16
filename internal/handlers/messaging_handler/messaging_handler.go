package messaginghandler

import (
	"context"
	"fmt"

	"github.com/bradford-hamilton/go-micro-gorm/internal/db"
	"github.com/bradford-hamilton/go-micro-gorm/pkg/json"
	pb "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

// Messaging defines the struct to which we attach different handlers to
type Messaging struct {
	DB *gorm.DB
}

// Call is a single request handler called via client.Call or the generated client code
func (m *Messaging) Call(ctx context.Context, req *pb.Request, res *pb.Response) error {
	log.Log("Received Messaging.Call request")
	res.Msg = fmt.Sprintf("Hey there %s", req.Name)
	return nil
}

// List all messages - currently takes a "message_type" and filters by that enum
func (m *Messaging) List(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	log.Log("Received Messaging.List request")

	// Create a slice of db.Message for gorm to use
	Messages := []db.Message{}

	// Query our db for any messages where MsgType == the request's asked MsgType
	if errs := m.
		DB.
		Where(&db.Message{MsgType: req.GetMessageType().String()}).
		Find(&Messages).
		GetErrors(); len(errs) > 0 {
		res.Errors = db.ErrSliceToString(errs)
		return nil
	}

	// Attmpt to marshall the Messages into bytes
	msgBytes, err := json.API.Marshal(&Messages)
	if err != nil {
		log.Errorf("Error marshalling json in List endpoint: %v", err)
		res.Errors = []string{err.Error()}
		return nil
	}

	// Set the found messages to our response
	res.Messages = string(msgBytes)

	return nil
}

// DestroyByID is the handler for destroying a message by its ID
func (m *Messaging) DestroyByID(ctx context.Context, req *pb.Request, res *pb.Response) error {
	log.Log("Received Messaging.DestroyByID request")
	ID := req.GetID()

	// Destroy the record, catch errors and return them early if there are any
	if errs := m.DB.Where("ID = ?", ID).Delete(&db.Message{}).GetErrors(); len(errs) > 0 {
		res.Errors = db.ErrSliceToString(errs)
		return nil
	}

	// Respond with successful message deletion
	res.Msg = fmt.Sprintf("Successfully destroyed message %d", ID)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (m *Messaging) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.MessagingService_StreamStream) error {
	log.Logf("Received Messaging.Stream request with count: %d", req.Count)

	// Loop over the request's Count and send that `n` number of messages back through the stream with a Count of i
	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{Count: int64(i)}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (m *Messaging) PingPong(ctx context.Context, stream pb.MessagingService_PingPongStream) error {
	// Continuously listen and respond to anything incoming
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
