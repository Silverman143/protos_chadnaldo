// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: flappySiuGame/games.proto

package flappySiuGamev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FlappySiuGame_CreateSession_FullMethodName         = "/flappySiuGame.FlappySiuGame/CreateSession"
	FlappySiuGame_SubmitScore_FullMethodName           = "/flappySiuGame.FlappySiuGame/SubmitScore"
	FlappySiuGame_GetBombOrBonusStatus_FullMethodName  = "/flappySiuGame.FlappySiuGame/GetBombOrBonusStatus"
	FlappySiuGame_BombOrBonusRevealCard_FullMethodName = "/flappySiuGame.FlappySiuGame/BombOrBonusRevealCard"
)

// FlappySiuGameClient is the client API for FlappySiuGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for managing the mining process
type FlappySiuGameClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	SubmitScore(ctx context.Context, in *SubmitScoreRequest, opts ...grpc.CallOption) (*SubmitScoreResponse, error)
	GetBombOrBonusStatus(ctx context.Context, in *GetBombOrBonusStatusRequest, opts ...grpc.CallOption) (*GetBombOrBonusStatusResponse, error)
	BombOrBonusRevealCard(ctx context.Context, in *BombOrBonusRevealCardRequest, opts ...grpc.CallOption) (*BombOrBonusRevealCardResponse, error)
}

type flappySiuGameClient struct {
	cc grpc.ClientConnInterface
}

func NewFlappySiuGameClient(cc grpc.ClientConnInterface) FlappySiuGameClient {
	return &flappySiuGameClient{cc}
}

func (c *flappySiuGameClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, FlappySiuGame_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flappySiuGameClient) SubmitScore(ctx context.Context, in *SubmitScoreRequest, opts ...grpc.CallOption) (*SubmitScoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitScoreResponse)
	err := c.cc.Invoke(ctx, FlappySiuGame_SubmitScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flappySiuGameClient) GetBombOrBonusStatus(ctx context.Context, in *GetBombOrBonusStatusRequest, opts ...grpc.CallOption) (*GetBombOrBonusStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBombOrBonusStatusResponse)
	err := c.cc.Invoke(ctx, FlappySiuGame_GetBombOrBonusStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flappySiuGameClient) BombOrBonusRevealCard(ctx context.Context, in *BombOrBonusRevealCardRequest, opts ...grpc.CallOption) (*BombOrBonusRevealCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BombOrBonusRevealCardResponse)
	err := c.cc.Invoke(ctx, FlappySiuGame_BombOrBonusRevealCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlappySiuGameServer is the server API for FlappySiuGame service.
// All implementations must embed UnimplementedFlappySiuGameServer
// for forward compatibility.
//
// Service for managing the mining process
type FlappySiuGameServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	SubmitScore(context.Context, *SubmitScoreRequest) (*SubmitScoreResponse, error)
	GetBombOrBonusStatus(context.Context, *GetBombOrBonusStatusRequest) (*GetBombOrBonusStatusResponse, error)
	BombOrBonusRevealCard(context.Context, *BombOrBonusRevealCardRequest) (*BombOrBonusRevealCardResponse, error)
	mustEmbedUnimplementedFlappySiuGameServer()
}

// UnimplementedFlappySiuGameServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFlappySiuGameServer struct{}

func (UnimplementedFlappySiuGameServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedFlappySiuGameServer) SubmitScore(context.Context, *SubmitScoreRequest) (*SubmitScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScore not implemented")
}
func (UnimplementedFlappySiuGameServer) GetBombOrBonusStatus(context.Context, *GetBombOrBonusStatusRequest) (*GetBombOrBonusStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBombOrBonusStatus not implemented")
}
func (UnimplementedFlappySiuGameServer) BombOrBonusRevealCard(context.Context, *BombOrBonusRevealCardRequest) (*BombOrBonusRevealCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BombOrBonusRevealCard not implemented")
}
func (UnimplementedFlappySiuGameServer) mustEmbedUnimplementedFlappySiuGameServer() {}
func (UnimplementedFlappySiuGameServer) testEmbeddedByValue()                       {}

// UnsafeFlappySiuGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlappySiuGameServer will
// result in compilation errors.
type UnsafeFlappySiuGameServer interface {
	mustEmbedUnimplementedFlappySiuGameServer()
}

func RegisterFlappySiuGameServer(s grpc.ServiceRegistrar, srv FlappySiuGameServer) {
	// If the following call pancis, it indicates UnimplementedFlappySiuGameServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FlappySiuGame_ServiceDesc, srv)
}

func _FlappySiuGame_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlappySiuGameServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlappySiuGame_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlappySiuGameServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlappySiuGame_SubmitScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlappySiuGameServer).SubmitScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlappySiuGame_SubmitScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlappySiuGameServer).SubmitScore(ctx, req.(*SubmitScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlappySiuGame_GetBombOrBonusStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBombOrBonusStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlappySiuGameServer).GetBombOrBonusStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlappySiuGame_GetBombOrBonusStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlappySiuGameServer).GetBombOrBonusStatus(ctx, req.(*GetBombOrBonusStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlappySiuGame_BombOrBonusRevealCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BombOrBonusRevealCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlappySiuGameServer).BombOrBonusRevealCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlappySiuGame_BombOrBonusRevealCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlappySiuGameServer).BombOrBonusRevealCard(ctx, req.(*BombOrBonusRevealCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlappySiuGame_ServiceDesc is the grpc.ServiceDesc for FlappySiuGame service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlappySiuGame_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flappySiuGame.FlappySiuGame",
	HandlerType: (*FlappySiuGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _FlappySiuGame_CreateSession_Handler,
		},
		{
			MethodName: "SubmitScore",
			Handler:    _FlappySiuGame_SubmitScore_Handler,
		},
		{
			MethodName: "GetBombOrBonusStatus",
			Handler:    _FlappySiuGame_GetBombOrBonusStatus_Handler,
		},
		{
			MethodName: "BombOrBonusRevealCard",
			Handler:    _FlappySiuGame_BombOrBonusRevealCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flappySiuGame/games.proto",
}
