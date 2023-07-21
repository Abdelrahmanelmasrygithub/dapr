//
//Copyright 2022 The Dapr Authors
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: dapr/proto/components/v1/pubsub.proto

package components

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

const (
	PubSub_Init_FullMethodName         = "/dapr.proto.components.v1.PubSub/Init"
	PubSub_Features_FullMethodName     = "/dapr.proto.components.v1.PubSub/Features"
	PubSub_Publish_FullMethodName      = "/dapr.proto.components.v1.PubSub/Publish"
	PubSub_BulkPublish_FullMethodName  = "/dapr.proto.components.v1.PubSub/BulkPublish"
	PubSub_PullMessages_FullMethodName = "/dapr.proto.components.v1.PubSub/PullMessages"
	PubSub_Ping_FullMethodName         = "/dapr.proto.components.v1.PubSub/Ping"
)

// PubSubClient is the client API for PubSub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubClient interface {
	// Initializes the pubsub component with the given metadata.
	Init(ctx context.Context, in *PubSubInitRequest, opts ...grpc.CallOption) (*PubSubInitResponse, error)
	// Returns a list of implemented pubsub features.
	Features(ctx context.Context, in *FeaturesRequest, opts ...grpc.CallOption) (*FeaturesResponse, error)
	// Publish publishes a new message for the given topic.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	BulkPublish(ctx context.Context, in *BulkPublishRequest, opts ...grpc.CallOption) (*BulkPublishResponse, error)
	// Establishes a stream with the server (PubSub component), which sends
	// messages down to the client (daprd). The client streams acknowledgements
	// back to the server. The server will close the stream and return the status
	// on any error. In case of closed connection, the client should re-establish
	// the stream. The first message MUST contain a `topic` attribute on it that
	// should be used for the entire streaming pull.
	PullMessages(ctx context.Context, opts ...grpc.CallOption) (PubSub_PullMessagesClient, error)
	// Ping the pubsub. Used for liveness porpuses.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type pubSubClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubClient(cc grpc.ClientConnInterface) PubSubClient {
	return &pubSubClient{cc}
}

func (c *pubSubClient) Init(ctx context.Context, in *PubSubInitRequest, opts ...grpc.CallOption) (*PubSubInitResponse, error) {
	out := new(PubSubInitResponse)
	err := c.cc.Invoke(ctx, PubSub_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) Features(ctx context.Context, in *FeaturesRequest, opts ...grpc.CallOption) (*FeaturesResponse, error) {
	out := new(FeaturesResponse)
	err := c.cc.Invoke(ctx, PubSub_Features_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, PubSub_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) BulkPublish(ctx context.Context, in *BulkPublishRequest, opts ...grpc.CallOption) (*BulkPublishResponse, error) {
	out := new(BulkPublishResponse)
	err := c.cc.Invoke(ctx, PubSub_BulkPublish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) PullMessages(ctx context.Context, opts ...grpc.CallOption) (PubSub_PullMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubSub_ServiceDesc.Streams[0], PubSub_PullMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubPullMessagesClient{stream}
	return x, nil
}

type PubSub_PullMessagesClient interface {
	Send(*PullMessagesRequest) error
	Recv() (*PullMessagesResponse, error)
	grpc.ClientStream
}

type pubSubPullMessagesClient struct {
	grpc.ClientStream
}

func (x *pubSubPullMessagesClient) Send(m *PullMessagesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubPullMessagesClient) Recv() (*PullMessagesResponse, error) {
	m := new(PullMessagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, PubSub_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServer is the server API for PubSub service.
// All implementations must embed UnimplementedPubSubServer
// for forward compatibility
type PubSubServer interface {
	// Initializes the pubsub component with the given metadata.
	Init(context.Context, *PubSubInitRequest) (*PubSubInitResponse, error)
	// Returns a list of implemented pubsub features.
	Features(context.Context, *FeaturesRequest) (*FeaturesResponse, error)
	// Publish publishes a new message for the given topic.
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	BulkPublish(context.Context, *BulkPublishRequest) (*BulkPublishResponse, error)
	// Establishes a stream with the server (PubSub component), which sends
	// messages down to the client (daprd). The client streams acknowledgements
	// back to the server. The server will close the stream and return the status
	// on any error. In case of closed connection, the client should re-establish
	// the stream. The first message MUST contain a `topic` attribute on it that
	// should be used for the entire streaming pull.
	PullMessages(PubSub_PullMessagesServer) error
	// Ping the pubsub. Used for liveness porpuses.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedPubSubServer()
}

// UnimplementedPubSubServer must be embedded to have forward compatible implementations.
type UnimplementedPubSubServer struct {
}

func (UnimplementedPubSubServer) Init(context.Context, *PubSubInitRequest) (*PubSubInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedPubSubServer) Features(context.Context, *FeaturesRequest) (*FeaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Features not implemented")
}
func (UnimplementedPubSubServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPubSubServer) BulkPublish(context.Context, *BulkPublishRequest) (*BulkPublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkPublish not implemented")
}
func (UnimplementedPubSubServer) PullMessages(PubSub_PullMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullMessages not implemented")
}
func (UnimplementedPubSubServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPubSubServer) mustEmbedUnimplementedPubSubServer() {}

// UnsafePubSubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServer will
// result in compilation errors.
type UnsafePubSubServer interface {
	mustEmbedUnimplementedPubSubServer()
}

func RegisterPubSubServer(s grpc.ServiceRegistrar, srv PubSubServer) {
	s.RegisterService(&PubSub_ServiceDesc, srv)
}

func _PubSub_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubSubInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSub_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Init(ctx, req.(*PubSubInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_Features_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Features(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSub_Features_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Features(ctx, req.(*FeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSub_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_BulkPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkPublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).BulkPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSub_BulkPublish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).BulkPublish(ctx, req.(*BulkPublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_PullMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServer).PullMessages(&pubSubPullMessagesServer{stream})
}

type PubSub_PullMessagesServer interface {
	Send(*PullMessagesResponse) error
	Recv() (*PullMessagesRequest, error)
	grpc.ServerStream
}

type pubSubPullMessagesServer struct {
	grpc.ServerStream
}

func (x *pubSubPullMessagesServer) Send(m *PullMessagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubPullMessagesServer) Recv() (*PullMessagesRequest, error) {
	m := new(PullMessagesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PubSub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubSub_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PubSub_ServiceDesc is the grpc.ServiceDesc for PubSub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.components.v1.PubSub",
	HandlerType: (*PubSubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _PubSub_Init_Handler,
		},
		{
			MethodName: "Features",
			Handler:    _PubSub_Features_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _PubSub_Publish_Handler,
		},
		{
			MethodName: "BulkPublish",
			Handler:    _PubSub_BulkPublish_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _PubSub_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullMessages",
			Handler:       _PubSub_PullMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dapr/proto/components/v1/pubsub.proto",
}