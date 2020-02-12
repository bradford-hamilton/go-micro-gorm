package handler

import (
	"context"

	proto_messaging "github.com/bradford-hamilton/go-micro-gorm/proto/messaging"
	"github.com/micro/go-micro/v2/util/log"
)

// Messaging defines the struct to which we attach different handlers to
type Messaging struct{}

// Call is a single request handler called via client.Call or the generated client code
func (m *Messaging) Call(ctx context.Context, req *proto_messaging.MessagingRequest, rsp *proto_messaging.MessagingResponse) error {
	log.Log("Received Messaging.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (m *Messaging) Stream(ctx context.Context, req *proto_messaging.StreamingRequest, stream proto_messaging.Messaging_StreamStream) error {
	log.Logf("Received Messaging.Stream request with count: %d", req.Count)

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
func (m *Messaging) PingPong(ctx context.Context, stream proto_messaging.Messaging_PingPongStream) error {
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
