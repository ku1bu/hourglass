// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	Register(ctx context.Context, in *CustomerRegisterRequest, opts ...grpc.CallOption) (Customer_RegisterClient, error)
	Feedback(ctx context.Context, in *FeedbackRequest, opts ...grpc.CallOption) (*CustomerInfo, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Register(ctx context.Context, in *CustomerRegisterRequest, opts ...grpc.CallOption) (Customer_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &Customer_ServiceDesc.Streams[0], "/Customer/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerRegisterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_RegisterClient interface {
	Recv() (*CustomerInfo, error)
	grpc.ClientStream
}

type customerRegisterClient struct {
	grpc.ClientStream
}

func (x *customerRegisterClient) Recv() (*CustomerInfo, error) {
	m := new(CustomerInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) Feedback(ctx context.Context, in *FeedbackRequest, opts ...grpc.CallOption) (*CustomerInfo, error) {
	out := new(CustomerInfo)
	err := c.cc.Invoke(ctx, "/Customer/Feedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility
type CustomerServer interface {
	Register(*CustomerRegisterRequest, Customer_RegisterServer) error
	Feedback(context.Context, *FeedbackRequest) (*CustomerInfo, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (UnimplementedCustomerServer) Register(*CustomerRegisterRequest, Customer_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCustomerServer) Feedback(context.Context, *FeedbackRequest) (*CustomerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feedback not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CustomerRegisterRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).Register(m, &customerRegisterServer{stream})
}

type Customer_RegisterServer interface {
	Send(*CustomerInfo) error
	grpc.ServerStream
}

type customerRegisterServer struct {
	grpc.ServerStream
}

func (x *customerRegisterServer) Send(m *CustomerInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_Feedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Feedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Customer/Feedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Feedback(ctx, req.(*FeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feedback",
			Handler:    _Customer_Feedback_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Customer_Register_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/hourglass/v1/customer.proto",
}
