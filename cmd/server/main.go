package main

import (
	"hello/greeter"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	greeter.RegisterGreeterServer(s, greeter.NewGreeterServer())
	s.Serve(listen)
}
