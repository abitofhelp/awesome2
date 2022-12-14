// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: awesome2/v1/awesome2.proto

package awesome2v1

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

// Awesome2ServiceClient is the client API for Awesome2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Awesome2ServiceClient interface {
	FindReportByOwnerName(ctx context.Context, in *FindReportByOwnerNameRequest, opts ...grpc.CallOption) (*FindReportByOwnerNameResponse, error)
}

type awesome2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAwesome2ServiceClient(cc grpc.ClientConnInterface) Awesome2ServiceClient {
	return &awesome2ServiceClient{cc}
}

func (c *awesome2ServiceClient) FindReportByOwnerName(ctx context.Context, in *FindReportByOwnerNameRequest, opts ...grpc.CallOption) (*FindReportByOwnerNameResponse, error) {
	out := new(FindReportByOwnerNameResponse)
	err := c.cc.Invoke(ctx, "/awesome2.v1.Awesome2Service/FindReportByOwnerName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Awesome2ServiceServer is the server API for Awesome2Service service.
// All implementations should embed UnimplementedAwesome2ServiceServer
// for forward compatibility
type Awesome2ServiceServer interface {
	FindReportByOwnerName(context.Context, *FindReportByOwnerNameRequest) (*FindReportByOwnerNameResponse, error)
}

// UnimplementedAwesome2ServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAwesome2ServiceServer struct {
}

func (UnimplementedAwesome2ServiceServer) FindReportByOwnerName(context.Context, *FindReportByOwnerNameRequest) (*FindReportByOwnerNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindReportByOwnerName not implemented")
}

// UnsafeAwesome2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Awesome2ServiceServer will
// result in compilation errors.
type UnsafeAwesome2ServiceServer interface {
	mustEmbedUnimplementedAwesome2ServiceServer()
}

func RegisterAwesome2ServiceServer(s grpc.ServiceRegistrar, srv Awesome2ServiceServer) {
	s.RegisterService(&Awesome2Service_ServiceDesc, srv)
}

func _Awesome2Service_FindReportByOwnerName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindReportByOwnerNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Awesome2ServiceServer).FindReportByOwnerName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awesome2.v1.Awesome2Service/FindReportByOwnerName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Awesome2ServiceServer).FindReportByOwnerName(ctx, req.(*FindReportByOwnerNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Awesome2Service_ServiceDesc is the grpc.ServiceDesc for Awesome2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Awesome2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "awesome2.v1.Awesome2Service",
	HandlerType: (*Awesome2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindReportByOwnerName",
			Handler:    _Awesome2Service_FindReportByOwnerName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awesome2/v1/awesome2.proto",
}
