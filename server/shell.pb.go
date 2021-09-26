// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: shell.proto

package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type Std struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *Std) Reset() {
	*x = Std{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Std) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Std) ProtoMessage() {}

func (x *Std) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Std.ProtoReflect.Descriptor instead.
func (*Std) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{1}
}

func (x *Std) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *Std) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

var File_shell_proto protoreflect.FileDescriptor

var file_shell_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6c,
	0x6e, 0x6b, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x68, 0x30, 0x32, 0x31, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x22, 0x17, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22, 0x35, 0x0a, 0x03, 0x53, 0x74, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x32, 0x58, 0x0a,
	0x05, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x4f, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x23, 0x2e,
	0x6c, 0x6e, 0x6b, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6c, 0x68, 0x30, 0x32, 0x31, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x4d,
	0x73, 0x67, 0x1a, 0x23, 0x2e, 0x6c, 0x6e, 0x6b, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x68, 0x30, 0x32, 0x31, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x68,
	0x65, 0x6c, 0x6c, 0x2e, 0x53, 0x74, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shell_proto_rawDescOnce sync.Once
	file_shell_proto_rawDescData = file_shell_proto_rawDesc
)

func file_shell_proto_rawDescGZIP() []byte {
	file_shell_proto_rawDescOnce.Do(func() {
		file_shell_proto_rawDescData = protoimpl.X.CompressGZIP(file_shell_proto_rawDescData)
	})
	return file_shell_proto_rawDescData
}

var file_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shell_proto_goTypes = []interface{}{
	(*Msg)(nil), // 0: lnks_homecloud.clh021.webshell.Msg
	(*Std)(nil), // 1: lnks_homecloud.clh021.webshell.Std
}
var file_shell_proto_depIdxs = []int32{
	0, // 0: lnks_homecloud.clh021.webshell.Shell.Run:input_type -> lnks_homecloud.clh021.webshell.Msg
	1, // 1: lnks_homecloud.clh021.webshell.Shell.Run:output_type -> lnks_homecloud.clh021.webshell.Std
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shell_proto_init() }
func file_shell_proto_init() {
	if File_shell_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_shell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Std); i {
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
			RawDescriptor: file_shell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shell_proto_goTypes,
		DependencyIndexes: file_shell_proto_depIdxs,
		MessageInfos:      file_shell_proto_msgTypes,
	}.Build()
	File_shell_proto = out.File
	file_shell_proto_rawDesc = nil
	file_shell_proto_goTypes = nil
	file_shell_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShellClient is the client API for Shell service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShellClient interface {
	Run(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Std, error)
}

type shellClient struct {
	cc grpc.ClientConnInterface
}

func NewShellClient(cc grpc.ClientConnInterface) ShellClient {
	return &shellClient{cc}
}

func (c *shellClient) Run(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Std, error) {
	out := new(Std)
	err := c.cc.Invoke(ctx, "/lnks_homecloud.clh021.webshell.Shell/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShellServer is the server API for Shell service.
type ShellServer interface {
	Run(context.Context, *Msg) (*Std, error)
}

// UnimplementedShellServer can be embedded to have forward compatible implementations.
type UnimplementedShellServer struct {
}

func (*UnimplementedShellServer) Run(context.Context, *Msg) (*Std, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterShellServer(s *grpc.Server, srv ShellServer) {
	s.RegisterService(&_Shell_serviceDesc, srv)
}

func _Shell_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShellServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lnks_homecloud.clh021.webshell.Shell/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShellServer).Run(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shell_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lnks_homecloud.clh021.webshell.Shell",
	HandlerType: (*ShellServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Shell_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shell.proto",
}