// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: tasks/tasks.proto

package tasksv1

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
	Tasks_GetAllTasks_FullMethodName        = "/tasks.Tasks/GetAllTasks"
	Tasks_IsTaskCompleted_FullMethodName    = "/tasks.Tasks/IsTaskCompleted"
	Tasks_GetTasksDaysStreak_FullMethodName = "/tasks.Tasks/GetTasksDaysStreak"
)

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for managing tasks
type TasksClient interface {
	// Get all tasks with pagination
	GetAllTasks(ctx context.Context, in *GetAllTasksRequest, opts ...grpc.CallOption) (*GetAllTasksResponse, error)
	// Check if a task is completed
	IsTaskCompleted(ctx context.Context, in *IsTaskCompletedRequest, opts ...grpc.CallOption) (*IsTaskCompletedResponse, error)
	// Get the number of consecutive days a user completed tasks
	GetTasksDaysStreak(ctx context.Context, in *GetTasksDaysStreakRequest, opts ...grpc.CallOption) (*GetTasksDaysStreakResponse, error)
}

type tasksClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksClient(cc grpc.ClientConnInterface) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) GetAllTasks(ctx context.Context, in *GetAllTasksRequest, opts ...grpc.CallOption) (*GetAllTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTasksResponse)
	err := c.cc.Invoke(ctx, Tasks_GetAllTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) IsTaskCompleted(ctx context.Context, in *IsTaskCompletedRequest, opts ...grpc.CallOption) (*IsTaskCompletedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsTaskCompletedResponse)
	err := c.cc.Invoke(ctx, Tasks_IsTaskCompleted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) GetTasksDaysStreak(ctx context.Context, in *GetTasksDaysStreakRequest, opts ...grpc.CallOption) (*GetTasksDaysStreakResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTasksDaysStreakResponse)
	err := c.cc.Invoke(ctx, Tasks_GetTasksDaysStreak_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
// All implementations must embed UnimplementedTasksServer
// for forward compatibility.
//
// Service for managing tasks
type TasksServer interface {
	// Get all tasks with pagination
	GetAllTasks(context.Context, *GetAllTasksRequest) (*GetAllTasksResponse, error)
	// Check if a task is completed
	IsTaskCompleted(context.Context, *IsTaskCompletedRequest) (*IsTaskCompletedResponse, error)
	// Get the number of consecutive days a user completed tasks
	GetTasksDaysStreak(context.Context, *GetTasksDaysStreakRequest) (*GetTasksDaysStreakResponse, error)
	mustEmbedUnimplementedTasksServer()
}

// UnimplementedTasksServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTasksServer struct{}

func (UnimplementedTasksServer) GetAllTasks(context.Context, *GetAllTasksRequest) (*GetAllTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTasks not implemented")
}
func (UnimplementedTasksServer) IsTaskCompleted(context.Context, *IsTaskCompletedRequest) (*IsTaskCompletedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsTaskCompleted not implemented")
}
func (UnimplementedTasksServer) GetTasksDaysStreak(context.Context, *GetTasksDaysStreakRequest) (*GetTasksDaysStreakResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasksDaysStreak not implemented")
}
func (UnimplementedTasksServer) mustEmbedUnimplementedTasksServer() {}
func (UnimplementedTasksServer) testEmbeddedByValue()               {}

// UnsafeTasksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServer will
// result in compilation errors.
type UnsafeTasksServer interface {
	mustEmbedUnimplementedTasksServer()
}

func RegisterTasksServer(s grpc.ServiceRegistrar, srv TasksServer) {
	// If the following call pancis, it indicates UnimplementedTasksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Tasks_ServiceDesc, srv)
}

func _Tasks_GetAllTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetAllTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tasks_GetAllTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetAllTasks(ctx, req.(*GetAllTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_IsTaskCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsTaskCompletedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).IsTaskCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tasks_IsTaskCompleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).IsTaskCompleted(ctx, req.(*IsTaskCompletedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_GetTasksDaysStreak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksDaysStreakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetTasksDaysStreak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tasks_GetTasksDaysStreak_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetTasksDaysStreak(ctx, req.(*GetTasksDaysStreakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tasks_ServiceDesc is the grpc.ServiceDesc for Tasks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tasks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tasks.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTasks",
			Handler:    _Tasks_GetAllTasks_Handler,
		},
		{
			MethodName: "IsTaskCompleted",
			Handler:    _Tasks_IsTaskCompleted_Handler,
		},
		{
			MethodName: "GetTasksDaysStreak",
			Handler:    _Tasks_GetTasksDaysStreak_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks/tasks.proto",
}
