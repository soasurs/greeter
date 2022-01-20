package main

import (
	"greeter/greeter"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterCodec(flatbuffers.FlatbuffersCodec{})
}

func main() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, greeter.NewGreeterServer())
	s.Serve(listen)
}
