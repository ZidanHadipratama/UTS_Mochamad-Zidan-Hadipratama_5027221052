// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pcmgmt/grpc.proto

package UTS_Mochamad_Zidan_Hadipratama_5027221052

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PRCServicesClient is the client API for PRCServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PRCServicesClient interface {
	CreatePRC(ctx context.Context, in *PRC, opts ...grpc.CallOption) (*PRCResponse, error)
	ReadPRC(ctx context.Context, in *PRCRequest, opts ...grpc.CallOption) (*PRC, error)
	UpdatePRC(ctx context.Context, in *PRC, opts ...grpc.CallOption) (*PRCResponse, error)
	DeletePRC(ctx context.Context, in *PRCRequest, opts ...grpc.CallOption) (*PRCResponse, error)
	ListPRCs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPRCsResponse, error)
}

type pRCServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewPRCServicesClient(cc grpc.ClientConnInterface) PRCServicesClient {
	return &pRCServicesClient{cc}
}

func (c *pRCServicesClient) CreatePRC(ctx context.Context, in *PRC, opts ...grpc.CallOption) (*PRCResponse, error) {
	out := new(PRCResponse)
	err := c.cc.Invoke(ctx, "/pcmanagement.PRCServices/CreatePRC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRCServicesClient) ReadPRC(ctx context.Context, in *PRCRequest, opts ...grpc.CallOption) (*PRC, error) {
	out := new(PRC)
	err := c.cc.Invoke(ctx, "/pcmanagement.PRCServices/ReadPRC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRCServicesClient) UpdatePRC(ctx context.Context, in *PRC, opts ...grpc.CallOption) (*PRCResponse, error) {
	out := new(PRCResponse)
	err := c.cc.Invoke(ctx, "/pcmanagement.PRCServices/UpdatePRC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRCServicesClient) DeletePRC(ctx context.Context, in *PRCRequest, opts ...grpc.CallOption) (*PRCResponse, error) {
	out := new(PRCResponse)
	err := c.cc.Invoke(ctx, "/pcmanagement.PRCServices/DeletePRC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pRCServicesClient) ListPRCs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPRCsResponse, error) {
	out := new(ListPRCsResponse)
	err := c.cc.Invoke(ctx, "/pcmanagement.PRCServices/ListPRCs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PRCServicesServer is the server API for PRCServices service.
// All implementations must embed UnimplementedPRCServicesServer
// for forward compatibility
type PRCServicesServer interface {
	CreatePRC(context.Context, *PRC) (*PRCResponse, error)
	ReadPRC(context.Context, *PRCRequest) (*PRC, error)
	UpdatePRC(context.Context, *PRC) (*PRCResponse, error)
	DeletePRC(context.Context, *PRCRequest) (*PRCResponse, error)
	ListPRCs(context.Context, *emptypb.Empty) (*ListPRCsResponse, error)
	mustEmbedUnimplementedPRCServicesServer()
}

// UnimplementedPRCServicesServer must be embedded to have forward compatible implementations.
type UnimplementedPRCServicesServer struct {
}

func (UnimplementedPRCServicesServer) CreatePRC(context.Context, *PRC) (*PRCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePRC not implemented")
}
func (UnimplementedPRCServicesServer) ReadPRC(context.Context, *PRCRequest) (*PRC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPRC not implemented")
}
func (UnimplementedPRCServicesServer) UpdatePRC(context.Context, *PRC) (*PRCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePRC not implemented")
}
func (UnimplementedPRCServicesServer) DeletePRC(context.Context, *PRCRequest) (*PRCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePRC not implemented")
}
func (UnimplementedPRCServicesServer) ListPRCs(context.Context, *emptypb.Empty) (*ListPRCsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPRCs not implemented")
}
func (UnimplementedPRCServicesServer) mustEmbedUnimplementedPRCServicesServer() {}

// UnsafePRCServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PRCServicesServer will
// result in compilation errors.
type UnsafePRCServicesServer interface {
	mustEmbedUnimplementedPRCServicesServer()
}

func RegisterPRCServicesServer(s grpc.ServiceRegistrar, srv PRCServicesServer) {
	s.RegisterService(&PRCServices_ServiceDesc, srv)
}

func _PRCServices_CreatePRC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PRC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRCServicesServer).CreatePRC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcmanagement.PRCServices/CreatePRC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRCServicesServer).CreatePRC(ctx, req.(*PRC))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRCServices_ReadPRC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PRCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRCServicesServer).ReadPRC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcmanagement.PRCServices/ReadPRC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRCServicesServer).ReadPRC(ctx, req.(*PRCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRCServices_UpdatePRC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PRC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRCServicesServer).UpdatePRC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcmanagement.PRCServices/UpdatePRC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRCServicesServer).UpdatePRC(ctx, req.(*PRC))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRCServices_DeletePRC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PRCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRCServicesServer).DeletePRC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcmanagement.PRCServices/DeletePRC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRCServicesServer).DeletePRC(ctx, req.(*PRCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PRCServices_ListPRCs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PRCServicesServer).ListPRCs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pcmanagement.PRCServices/ListPRCs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PRCServicesServer).ListPRCs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PRCServices_ServiceDesc is the grpc.ServiceDesc for PRCServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PRCServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pcmanagement.PRCServices",
	HandlerType: (*PRCServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePRC",
			Handler:    _PRCServices_CreatePRC_Handler,
		},
		{
			MethodName: "ReadPRC",
			Handler:    _PRCServices_ReadPRC_Handler,
		},
		{
			MethodName: "UpdatePRC",
			Handler:    _PRCServices_UpdatePRC_Handler,
		},
		{
			MethodName: "DeletePRC",
			Handler:    _PRCServices_DeletePRC_Handler,
		},
		{
			MethodName: "ListPRCs",
			Handler:    _PRCServices_ListPRCs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pcmgmt/grpc.proto",
}
