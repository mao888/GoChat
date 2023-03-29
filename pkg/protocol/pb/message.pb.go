// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: message.proto

package pb

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

// 上行消息类型
type CmdType int32

const (
	CmdType_UnKnow    CmdType = 0 // 未知
	CmdType_Login     CmdType = 1 // 连接注册，客户端向服务端发送，建立连接
	CmdType_Heartbeat CmdType = 2 // 心跳，客户端向服务端发送，连接保活
	CmdType_SYNC      CmdType = 3 // 离线消息同步，客户端向服务端发送，拉取离线消息
	CmdType_ACK       CmdType = 4 // ACK 消息，回复
	CmdType_Up        CmdType = 5 // 消息上行投递，指从客户端向服务端发送的消息
	CmdType_Push      CmdType = 6 // 消息下行投递，指从服务端向客户端发送的消息
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "UnKnow",
		1: "Login",
		2: "Heartbeat",
		3: "SYNC",
		4: "ACK",
		5: "Up",
		6: "Push",
	}
	CmdType_value = map[string]int32{
		"UnKnow":    0,
		"Login":     1,
		"Heartbeat": 2,
		"SYNC":      3,
		"ACK":       4,
		"Up":        5,
		"Push":      6,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

// 顶层消息
type CmdMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type CmdType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.CmdType" json:"type,omitempty"` // 消息类型，根据不同消息类型，解析成下面不同的类型
	Data []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                  // 数据
}

func (x *CmdMsg) Reset() {
	*x = CmdMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdMsg) ProtoMessage() {}

func (x *CmdMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdMsg.ProtoReflect.Descriptor instead.
func (*CmdMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *CmdMsg) GetType() CmdType {
	if x != nil {
		return x.Type
	}
	return CmdType_UnKnow
}

func (x *CmdMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 登录
type LoginMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // token
}

func (x *LoginMsg) Reset() {
	*x = LoginMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsg) ProtoMessage() {}

func (x *LoginMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsg.ProtoReflect.Descriptor instead.
func (*LoginMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *LoginMsg) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LoginMsg) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

// 心跳
type HeartBeatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartBeatMsg) Reset() {
	*x = HeartBeatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatMsg) ProtoMessage() {}

func (x *HeartBeatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatMsg.ProtoReflect.Descriptor instead.
func (*HeartBeatMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

// 离线消息同步
type SyncMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncMsg) Reset() {
	*x = SyncMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMsg) ProtoMessage() {}

func (x *SyncMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMsg.ProtoReflect.Descriptor instead.
func (*SyncMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

// ACK
type ACKMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ACKMsg) Reset() {
	*x = ACKMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACKMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACKMsg) ProtoMessage() {}

func (x *ACKMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACKMsg.ProtoReflect.Descriptor instead.
func (*ACKMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

// 上行消息
type UpMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId uint64 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 实际数据
}

func (x *UpMsg) Reset() {
	*x = UpMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpMsg) ProtoMessage() {}

func (x *UpMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpMsg.ProtoReflect.Descriptor instead.
func (*UpMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *UpMsg) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *UpMsg) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// 下行消息
type PushMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId   uint64 `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"` // 消息id
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`           // 数据
}

func (x *PushMsg) Reset() {
	*x = PushMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsg) ProtoMessage() {}

func (x *PushMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsg.ProtoReflect.Descriptor instead.
func (*PushMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *PushMsg) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *PushMsg) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x3d, 0x0a, 0x06, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x1f, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x39, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0e, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x09, 0x0a,
	0x07, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x73, 0x67, 0x22, 0x08, 0x0a, 0x06, 0x41, 0x43, 0x4b, 0x4d,
	0x73, 0x67, 0x22, 0x3e, 0x0a, 0x05, 0x55, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x3a, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x15, 0x0a,
	0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x54,
	0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x4b,
	0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b,
	0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x70, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x75,
	0x73, 0x68, 0x10, 0x06, 0x42, 0x18, 0x5a, 0x16, 0x47, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_message_proto_goTypes = []interface{}{
	(CmdType)(0),         // 0: pb.CmdType
	(*CmdMsg)(nil),       // 1: pb.CmdMsg
	(*LoginMsg)(nil),     // 2: pb.LoginMsg
	(*HeartBeatMsg)(nil), // 3: pb.HeartBeatMsg
	(*SyncMsg)(nil),      // 4: pb.SyncMsg
	(*ACKMsg)(nil),       // 5: pb.ACKMsg
	(*UpMsg)(nil),        // 6: pb.UpMsg
	(*PushMsg)(nil),      // 7: pb.PushMsg
}
var file_message_proto_depIdxs = []int32{
	0, // 0: pb.CmdMsg.type:type_name -> pb.CmdType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdMsg); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMsg); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatMsg); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMsg); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACKMsg); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpMsg); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsg); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
