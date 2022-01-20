package greeter

import (
	"context"
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
)

type server struct {
	UnimplementedGreeterServer
}

func NewGreeterServer() GreeterServer {
	return &server{}
}

func (s *server) Greet(ctx context.Context, req *GreetRequest) (*flatbuffers.Builder, error) {
	fmt.Println(string(req.Greeting()))

	builder := flatbuffers.NewBuilder(16)
	ack := builder.CreateString("ack")
	GreetResponseStart(builder)
	GreetResponseAddAck(builder, ack)
	end := GreetResponseEnd(builder)
	builder.Finish(end)

	return builder, nil
}
