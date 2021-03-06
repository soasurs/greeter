//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: greeter

package greeter

import (
	context "context"

	flatbuffers "github.com/google/flatbuffers/go"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client API for Greeter service
type GreeterClient interface {
	Greet(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*GreetResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Greet(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greeter.Greeter/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service
type GreeterServer interface {
	Greet(context.Context, *GreetRequest) (*flatbuffers.Builder, error)
	mustEmbedUnimplementedGreeterServer()
}

type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Greet(context.Context, *GreetRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Greet_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.Greeter/Greet",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Greeter_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
