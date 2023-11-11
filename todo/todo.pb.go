// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: todo.proto

package todo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	DateAt  string `protobuf:"bytes,4,opt,name=date_at,json=dateAt,proto3" json:"date_at,omitempty"`
	TimeAt  string `protobuf:"bytes,5,opt,name=time_at,json=timeAt,proto3" json:"time_at,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *ToDo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ToDo) GetDateAt() string {
	if x != nil {
		return x.DateAt
	}
	return ""
}

func (x *ToDo) GetTimeAt() string {
	if x != nil {
		return x.TimeAt
	}
	return ""
}

type ToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	DateAt  string `protobuf:"bytes,3,opt,name=date_at,json=dateAt,proto3" json:"date_at,omitempty"`
	TimeAt  string `protobuf:"bytes,4,opt,name=time_at,json=timeAt,proto3" json:"time_at,omitempty"`
}

func (x *ToDoRequest) Reset() {
	*x = ToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoRequest) ProtoMessage() {}

func (x *ToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoRequest.ProtoReflect.Descriptor instead.
func (*ToDoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *ToDoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ToDoRequest) GetDateAt() string {
	if x != nil {
		return x.DateAt
	}
	return ""
}

func (x *ToDoRequest) GetTimeAt() string {
	if x != nil {
		return x.TimeAt
	}
	return ""
}

type ToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*ToDo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *ToDoResponse) Reset() {
	*x = ToDoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDoResponse) ProtoMessage() {}

func (x *ToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDoResponse.ProtoReflect.Descriptor instead.
func (*ToDoResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *ToDoResponse) GetTodos() []*ToDo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type GetAllToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *GetAllToDoRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAllToDoRequest) Reset() {
	*x = GetAllToDoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllToDoRequest) ProtoMessage() {}

func (x *GetAllToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllToDoRequest.ProtoReflect.Descriptor instead.
func (*GetAllToDoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllToDoRequest) GetFilter() *GetAllToDoRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId int64 `protobuf:"varint,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *GetTodoRequest) GetTodoId() int64 {
	if x != nil {
		return x.TodoId
	}
	return 0
}

type GetAllToDoRequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	LastId int64  `protobuf:"varint,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetAllToDoRequest_Filter) Reset() {
	*x = GetAllToDoRequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllToDoRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllToDoRequest_Filter) ProtoMessage() {}

func (x *GetAllToDoRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllToDoRequest_Filter.ProtoReflect.Descriptor instead.
func (*GetAllToDoRequest_Filter) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetAllToDoRequest_Filter) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllToDoRequest_Filter) GetLastId() int64 {
	if x != nil {
		return x.LastId
	}
	return 0
}

func (x *GetAllToDoRequest_Filter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x78, 0x0a, 0x04, 0x54,
	0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x69, 0x6d, 0x65, 0x41, 0x74, 0x22, 0x6f, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x6d, 0x65, 0x41, 0x74, 0x22, 0x30, 0x0a, 0x0c, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44,
	0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x4b, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x32, 0xa7,
	0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x17, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x09, 0x44, 0x6f,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54,
	0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x44, 0x6f, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_todo_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: todo.Empty
	(*ToDo)(nil),                     // 1: todo.ToDo
	(*ToDoRequest)(nil),              // 2: todo.ToDoRequest
	(*ToDoResponse)(nil),             // 3: todo.ToDoResponse
	(*GetAllToDoRequest)(nil),        // 4: todo.GetAllToDoRequest
	(*GetTodoRequest)(nil),           // 5: todo.GetTodoRequest
	(*GetAllToDoRequest_Filter)(nil), // 6: todo.GetAllToDoRequest.Filter
}
var file_todo_proto_depIdxs = []int32{
	1, // 0: todo.ToDoResponse.todos:type_name -> todo.ToDo
	6, // 1: todo.GetAllToDoRequest.filter:type_name -> todo.GetAllToDoRequest.Filter
	4, // 2: todo.TodoService.GetAllToDo:input_type -> todo.GetAllToDoRequest
	2, // 3: todo.TodoService.DoAddTodo:input_type -> todo.ToDoRequest
	5, // 4: todo.TodoService.GetToDo:input_type -> todo.GetTodoRequest
	2, // 5: todo.TodoService.GetAllToDo:output_type -> todo.ToDoRequest
	0, // 6: todo.TodoService.DoAddTodo:output_type -> todo.Empty
	1, // 7: todo.TodoService.GetToDo:output_type -> todo.ToDo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllToDoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllToDoRequest_Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
