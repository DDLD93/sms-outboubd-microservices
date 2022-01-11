// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package smspb

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

// SMSServiceApiClient is the client API for SMSServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSServiceApiClient interface {
	Send(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error)
	GetOne(ctx context.Context, in *GetSMSRequest, opts ...grpc.CallOption) (*GetSMSResponse, error)
	GetAll(ctx context.Context, in *GetAllSMSRequest, opts ...grpc.CallOption) (SMSServiceApi_GetAllClient, error)
}

type sMSServiceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSServiceApiClient(cc grpc.ClientConnInterface) SMSServiceApiClient {
	return &sMSServiceApiClient{cc}
}

func (c *sMSServiceApiClient) Send(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error) {
	out := new(SendSMSResponse)
	err := c.cc.Invoke(ctx, "/main.SMSServiceApi/send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSServiceApiClient) GetOne(ctx context.Context, in *GetSMSRequest, opts ...grpc.CallOption) (*GetSMSResponse, error) {
	out := new(GetSMSResponse)
	err := c.cc.Invoke(ctx, "/main.SMSServiceApi/getOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSServiceApiClient) GetAll(ctx context.Context, in *GetAllSMSRequest, opts ...grpc.CallOption) (SMSServiceApi_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &SMSServiceApi_ServiceDesc.Streams[0], "/main.SMSServiceApi/getAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &sMSServiceApiGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SMSServiceApi_GetAllClient interface {
	Recv() (*GetAllSMSResponse, error)
	grpc.ClientStream
}

type sMSServiceApiGetAllClient struct {
	grpc.ClientStream
}

func (x *sMSServiceApiGetAllClient) Recv() (*GetAllSMSResponse, error) {
	m := new(GetAllSMSResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SMSServiceApiServer is the server API for SMSServiceApi service.
// All implementations must embed UnimplementedSMSServiceApiServer
// for forward compatibility
type SMSServiceApiServer interface {
	Send(context.Context, *SendSMSRequest) (*SendSMSResponse, error)
	GetOne(context.Context, *GetSMSRequest) (*GetSMSResponse, error)
	GetAll(*GetAllSMSRequest, SMSServiceApi_GetAllServer) error
	mustEmbedUnimplementedSMSServiceApiServer()
}

// UnimplementedSMSServiceApiServer must be embedded to have forward compatible implementations.
type UnimplementedSMSServiceApiServer struct {
}

func (UnimplementedSMSServiceApiServer) Send(context.Context, *SendSMSRequest) (*SendSMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSMSServiceApiServer) GetOne(context.Context, *GetSMSRequest) (*GetSMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedSMSServiceApiServer) GetAll(*GetAllSMSRequest, SMSServiceApi_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedSMSServiceApiServer) mustEmbedUnimplementedSMSServiceApiServer() {}

// UnsafeSMSServiceApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSServiceApiServer will
// result in compilation errors.
type UnsafeSMSServiceApiServer interface {
	mustEmbedUnimplementedSMSServiceApiServer()
}

func RegisterSMSServiceApiServer(s grpc.ServiceRegistrar, srv SMSServiceApiServer) {
	s.RegisterService(&SMSServiceApi_ServiceDesc, srv)
}

func _SMSServiceApi_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceApiServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.SMSServiceApi/send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceApiServer).Send(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSServiceApi_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceApiServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.SMSServiceApi/getOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceApiServer).GetOne(ctx, req.(*GetSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSServiceApi_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllSMSRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SMSServiceApiServer).GetAll(m, &sMSServiceApiGetAllServer{stream})
}

type SMSServiceApi_GetAllServer interface {
	Send(*GetAllSMSResponse) error
	grpc.ServerStream
}

type sMSServiceApiGetAllServer struct {
	grpc.ServerStream
}

func (x *sMSServiceApiGetAllServer) Send(m *GetAllSMSResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SMSServiceApi_ServiceDesc is the grpc.ServiceDesc for SMSServiceApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMSServiceApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.SMSServiceApi",
	HandlerType: (*SMSServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _SMSServiceApi_Send_Handler,
		},
		{
			MethodName: "getOne",
			Handler:    _SMSServiceApi_GetOne_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getAll",
			Handler:       _SMSServiceApi_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sms.proto",
}