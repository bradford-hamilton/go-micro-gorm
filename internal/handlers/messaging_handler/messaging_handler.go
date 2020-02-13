package messaginghandler

import (
	"context"
	"fmt"
	"strings"

	"github.com/bradford-hamilton/go-micro-gorm/internal/db"
	"github.com/bradford-hamilton/go-micro-gorm/pkg/json"
	proto_messaging "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

// Messaging defines the struct to which we attach different handlers to
type Messaging struct {
	Db *gorm.DB
}

// Call is a single request handler called via client.Call or the generated client code
func (m *Messaging) Call(
	ctx context.Context,
	req *proto_messaging.MessagingRequest,
	res *proto_messaging.MessagingResponse,
) error {
	log.Log("Received Messaging.Call request")
	res.Msg = fmt.Sprintf("Hey there %s", req.Name)
	return nil
}

// List all messages - currently takes a "message_type" and filters by that enum
func (m *Messaging) List(
	ctx context.Context,
	req *proto_messaging.MessagingListRequest,
	res *proto_messaging.MessagingListResponse,
) error {
	log.Log("Received Messaging.List request")

	// Create a slice of db.Message for gorm to use
	Messages := []db.Message{}

	// Query our db for any messages where MsgType == the request's asked MsgType
	if errs := m.
		Db.
		Where(&db.Message{MsgType: req.GetMessageType().String()}).
		Find(&Messages).
		GetErrors(); len(errs) > 0 {
		res.Errors = errSliceToString(errs)
		return nil
	}

	// Attmpt to marshall the Messages into bytes
	msgBytes, err := json.API.Marshal(&Messages)
	if err != nil {
		log.Errorf("Error marshalling json in List endpoint: %v", err)
		res.Errors = err.Error()
		return nil
	}

	// Set the found messages to our response
	res.Messages = string(msgBytes)

	return nil
}

// DestroyByID is the handler for destroying a message by its ID
func (m *Messaging) DestroyByID(
	ctx context.Context,
	req *proto_messaging.MessagingRequest,
	res *proto_messaging.MessagingResponse,
) error {
	log.Log("Received Messaging.DestroyByID request")
	ID := req.GetID()

	// Destroy the record, catch errors and returning them early if there are any
	if errs := m.Db.Where("ID = ?", ID).Delete(&db.Message{}).GetErrors(); len(errs) > 0 {
		res.Errors = errSliceToString(errs)
		return nil
	}

	// Respond with successful message deletion
	res.Msg = fmt.Sprintf("Successfully destroyed message %d", ID)

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (m *Messaging) Stream(
	ctx context.Context,
	req *proto_messaging.StreamingRequest,
	stream proto_messaging.MessagingService_StreamStream,
) error {
	log.Logf("Received Messaging.Stream request with count: %d", req.Count)

	// Loop over the request's Count and send that `n` number of messages back through the stream with a Count of i
	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&proto_messaging.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (m *Messaging) PingPong(ctx context.Context, stream proto_messaging.MessagingService_PingPongStream) error {
	// Continuously send & receive from the stream
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&proto_messaging.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

// errSliceToString is a small helper function for now to join a slice of errors, which
// is the type that is returned from a gorm db call with GetErrors() chained.
func errSliceToString(errs []error) string {
	var errStrings []string
	for _, e := range errs {
		errStrings = append(errStrings, e.Error())
	}
	return strings.Join(errStrings, ", ")
}
