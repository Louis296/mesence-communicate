// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
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

type RequestStatus int32

const (
	RequestStatus_Accepted RequestStatus = 0
	RequestStatus_Refused  RequestStatus = 1
	RequestStatus_Waiting  RequestStatus = 2
)

// Enum value maps for RequestStatus.
var (
	RequestStatus_name = map[int32]string{
		0: "Accepted",
		1: "Refused",
		2: "Waiting",
	}
	RequestStatus_value = map[string]int32{
		"Accepted": 0,
		"Refused":  1,
		"Waiting":  2,
	}
)

func (x RequestStatus) Enum() *RequestStatus {
	p := new(RequestStatus)
	*p = x
	return p
}

func (x RequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (RequestStatus) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x RequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestStatus.Descriptor instead.
func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_Word          Type = 0
	Type_Online        Type = 1
	Type_Offline       Type = 2
	Type_FriendRequest Type = 3
	Type_HeartPackage  Type = 4
	Type_Offer         Type = 5
	Type_Answer        Type = 6
	Type_Candidate     Type = 7
	Type_GetMaxSeq     Type = 8
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Word",
		1: "Online",
		2: "Offline",
		3: "FriendRequest",
		4: "HeartPackage",
		5: "Offer",
		6: "Answer",
		7: "Candidate",
		8: "GetMaxSeq",
	}
	Type_value = map[string]int32{
		"Word":          0,
		"Online":        1,
		"Offline":       2,
		"FriendRequest": 3,
		"HeartPackage":  4,
		"Offer":         5,
		"Answer":        6,
		"Candidate":     7,
		"GetMaxSeq":     8,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type  `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Seq  int64 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Word
}

func (x *Msg) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Msg) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To            string        `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	From          string        `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Content       string        `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	SendTime      string        `protobuf:"bytes,4,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
	OnlineUsers   []string      `protobuf:"bytes,5,rep,name=onlineUsers,proto3" json:"onlineUsers,omitempty"`
	OfflineUsers  []string      `protobuf:"bytes,6,rep,name=offlineUsers,proto3" json:"offlineUsers,omitempty"`
	Candidate     string        `protobuf:"bytes,7,opt,name=candidate,proto3" json:"candidate,omitempty"`
	RequestStatus RequestStatus `protobuf:"varint,8,opt,name=requestStatus,proto3,enum=RequestStatus" json:"requestStatus,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Data) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Data) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Data) GetSendTime() string {
	if x != nil {
		return x.SendTime
	}
	return ""
}

func (x *Data) GetOnlineUsers() []string {
	if x != nil {
		return x.OnlineUsers
	}
	return nil
}

func (x *Data) GetOfflineUsers() []string {
	if x != nil {
		return x.OfflineUsers
	}
	return nil
}

func (x *Data) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

func (x *Data) GetRequestStatus() RequestStatus {
	if x != nil {
		return x.RequestStatus
	}
	return RequestStatus_Accepted
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrStr string `protobuf:"bytes,1,opt,name=errStr,proto3" json:"errStr,omitempty"`
	Msg    *Msg   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Resp) GetErrStr() string {
	if x != nil {
		return x.ErrStr
	}
	return ""
}

func (x *Resp) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4d, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x22, 0xfa,
	0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x04, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x53, 0x74, 0x72, 0x12, 0x16, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x2a, 0x37, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x75, 0x73, 0x65, 0x64, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x2a, 0x83, 0x01, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x71,
	0x10, 0x08, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x6f, 0x75, 0x69, 0x73, 0x32, 0x39, 0x36, 0x2f, 0x6d, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_proto_goTypes = []interface{}{
	(RequestStatus)(0), // 0: RequestStatus
	(Type)(0),          // 1: Type
	(*Msg)(nil),        // 2: Msg
	(*Data)(nil),       // 3: Data
	(*Resp)(nil),       // 4: Resp
}
var file_message_proto_depIdxs = []int32{
	1, // 0: Msg.type:type_name -> Type
	3, // 1: Msg.data:type_name -> Data
	0, // 2: Data.requestStatus:type_name -> RequestStatus
	2, // 3: Resp.msg:type_name -> Msg
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			switch v := v.(*Resp); i {
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
			NumEnums:      2,
			NumMessages:   3,
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
