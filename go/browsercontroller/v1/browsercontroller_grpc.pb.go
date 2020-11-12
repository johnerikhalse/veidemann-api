// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package browsercontroller

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BrowserControllerClient is the client API for BrowserController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrowserControllerClient interface {
	Do(ctx context.Context, opts ...grpc.CallOption) (BrowserController_DoClient, error)
}

type browserControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrowserControllerClient(cc grpc.ClientConnInterface) BrowserControllerClient {
	return &browserControllerClient{cc}
}

func (c *browserControllerClient) Do(ctx context.Context, opts ...grpc.CallOption) (BrowserController_DoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BrowserController_serviceDesc.Streams[0], "/veidemann.api.browsercontroller.v1.BrowserController/do", opts...)
	if err != nil {
		return nil, err
	}
	x := &browserControllerDoClient{stream}
	return x, nil
}

type BrowserController_DoClient interface {
	Send(*DoRequest) error
	Recv() (*DoReply, error)
	grpc.ClientStream
}

type browserControllerDoClient struct {
	grpc.ClientStream
}

func (x *browserControllerDoClient) Send(m *DoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *browserControllerDoClient) Recv() (*DoReply, error) {
	m := new(DoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrowserControllerServer is the server API for BrowserController service.
// All implementations must embed UnimplementedBrowserControllerServer
// for forward compatibility
type BrowserControllerServer interface {
	Do(BrowserController_DoServer) error
	mustEmbedUnimplementedBrowserControllerServer()
}

// UnimplementedBrowserControllerServer must be embedded to have forward compatible implementations.
type UnimplementedBrowserControllerServer struct {
}

func (UnimplementedBrowserControllerServer) Do(BrowserController_DoServer) error {
	return status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedBrowserControllerServer) mustEmbedUnimplementedBrowserControllerServer() {}

// UnsafeBrowserControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrowserControllerServer will
// result in compilation errors.
type UnsafeBrowserControllerServer interface {
	mustEmbedUnimplementedBrowserControllerServer()
}

func RegisterBrowserControllerServer(s grpc.ServiceRegistrar, srv BrowserControllerServer) {
	s.RegisterService(&_BrowserController_serviceDesc, srv)
}

func _BrowserController_Do_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrowserControllerServer).Do(&browserControllerDoServer{stream})
}

type BrowserController_DoServer interface {
	Send(*DoReply) error
	Recv() (*DoRequest, error)
	grpc.ServerStream
}

type browserControllerDoServer struct {
	grpc.ServerStream
}

func (x *browserControllerDoServer) Send(m *DoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *browserControllerDoServer) Recv() (*DoRequest, error) {
	m := new(DoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BrowserController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.browsercontroller.v1.BrowserController",
	HandlerType: (*BrowserControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "do",
			Handler:       _BrowserController_Do_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "browsercontroller/v1/browsercontroller.proto",
}
