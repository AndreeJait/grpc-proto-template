// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: todo.proto

package todo

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	GetAllToDo(ctx context.Context, in *GetAllToDoRequest, opts ...grpc.CallOption) (*ToDoResponse, error)
	DoAddTodo(ctx context.Context, in *ToDoRequest, opts ...grpc.CallOption) (*Empty, error)
	GetToDo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*ToDo, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) GetAllToDo(ctx context.Context, in *GetAllToDoRequest, opts ...grpc.CallOption) (*ToDoResponse, error) {
	out := new(ToDoResponse)
	err := c.cc.Invoke(ctx, "/todo.TodoService/GetAllToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DoAddTodo(ctx context.Context, in *ToDoRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todo.TodoService/DoAddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetToDo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*ToDo, error) {
	out := new(ToDo)
	err := c.cc.Invoke(ctx, "/todo.TodoService/GetToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations should embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	GetAllToDo(context.Context, *GetAllToDoRequest) (*ToDoResponse, error)
	DoAddTodo(context.Context, *ToDoRequest) (*Empty, error)
	GetToDo(context.Context, *GetTodoRequest) (*ToDo, error)
}

// UnimplementedTodoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) GetAllToDo(context.Context, *GetAllToDoRequest) (*ToDoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllToDo not implemented")
}
func (UnimplementedTodoServiceServer) DoAddTodo(context.Context, *ToDoRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoAddTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetToDo(context.Context, *GetTodoRequest) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToDo not implemented")
}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_GetAllToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetAllToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/GetAllToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetAllToDo(ctx, req.(*GetAllToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DoAddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DoAddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/DoAddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DoAddTodo(ctx, req.(*ToDoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/GetToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetToDo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllToDo",
			Handler:    _TodoService_GetAllToDo_Handler,
		},
		{
			MethodName: "DoAddTodo",
			Handler:    _TodoService_DoAddTodo_Handler,
		},
		{
			MethodName: "GetToDo",
			Handler:    _TodoService_GetToDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
