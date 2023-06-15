// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: api/proto/seleniumserver.proto

package selserv

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

// SeleniumServerClient is the client API for SeleniumServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeleniumServerClient interface {
	GetHtml(ctx context.Context, in *HtmlArgs, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type seleniumServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSeleniumServerClient(cc grpc.ClientConnInterface) SeleniumServerClient {
	return &seleniumServerClient{cc}
}

func (c *seleniumServerClient) GetHtml(ctx context.Context, in *HtmlArgs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/selserv.SeleniumServer/GetHtml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeleniumServerServer is the server API for SeleniumServer service.
// All implementations must embed UnimplementedSeleniumServerServer
// for forward compatibility
type SeleniumServerServer interface {
	GetHtml(context.Context, *HtmlArgs) (*emptypb.Empty, error)
	mustEmbedUnimplementedSeleniumServerServer()
}

// UnimplementedSeleniumServerServer must be embedded to have forward compatible implementations.
type UnimplementedSeleniumServerServer struct {
}

func (UnimplementedSeleniumServerServer) GetHtml(context.Context, *HtmlArgs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHtml not implemented")
}
func (UnimplementedSeleniumServerServer) mustEmbedUnimplementedSeleniumServerServer() {}

// UnsafeSeleniumServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeleniumServerServer will
// result in compilation errors.
type UnsafeSeleniumServerServer interface {
	mustEmbedUnimplementedSeleniumServerServer()
}

func RegisterSeleniumServerServer(s grpc.ServiceRegistrar, srv SeleniumServerServer) {
	s.RegisterService(&SeleniumServer_ServiceDesc, srv)
}

func _SeleniumServer_GetHtml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HtmlArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeleniumServerServer).GetHtml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selserv.SeleniumServer/GetHtml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeleniumServerServer).GetHtml(ctx, req.(*HtmlArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// SeleniumServer_ServiceDesc is the grpc.ServiceDesc for SeleniumServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeleniumServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "selserv.SeleniumServer",
	HandlerType: (*SeleniumServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHtml",
			Handler:    _SeleniumServer_GetHtml_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/seleniumserver.proto",
}