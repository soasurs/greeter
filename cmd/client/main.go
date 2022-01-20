package main

import (
	"context"
	"fmt"
	"greeter/greeter"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{}))}
	)

	conn, err := grpc.Dial(":3000", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := greeter.NewGreeterClient(conn)

	builder := flatbuffers.NewBuilder(16)
	greeting := builder.CreateByteString([]byte("Hello World"))
	greeter.GreetRequestStart(builder)
	greeter.GreetRequestAddGreeting(builder, greeting)
	end := greeter.GreetRequestEnd(builder)
	builder.Finish(end)

	resp, err := client.Greet(context.Background(), builder)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Ack())
}
