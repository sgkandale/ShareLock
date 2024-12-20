// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: sharelock.proto

package sharelockPB

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

type Status int32

const (
	Status_Unknown     Status = 0
	Status_Acquired    Status = 1
	Status_NotAcquired Status = 2
	Status_Released    Status = 3
	Status_Timeout     Status = 4
	Status_UnknownLock Status = 5
	Status_InvalidData Status = 6
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Unknown",
		1: "Acquired",
		2: "NotAcquired",
		3: "Released",
		4: "Timeout",
		5: "UnknownLock",
		6: "InvalidData",
	}
	Status_value = map[string]int32{
		"Unknown":     0,
		"Acquired":    1,
		"NotAcquired": 2,
		"Released":    3,
		"Timeout":     4,
		"UnknownLock": 5,
		"InvalidData": 6,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_sharelock_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_sharelock_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{0}
}

type ShareLockPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ShareLockPingRequest) Reset() {
	*x = ShareLockPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharelock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareLockPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareLockPingRequest) ProtoMessage() {}

func (x *ShareLockPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sharelock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareLockPingRequest.ProtoReflect.Descriptor instead.
func (*ShareLockPingRequest) Descriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{0}
}

func (x *ShareLockPingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ShareLockPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ShareLockPingResponse) Reset() {
	*x = ShareLockPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharelock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareLockPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareLockPingResponse) ProtoMessage() {}

func (x *ShareLockPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sharelock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareLockPingResponse.ProtoReflect.Descriptor instead.
func (*ShareLockPingResponse) Descriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{1}
}

func (x *ShareLockPingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	TimeoutMs int32  `protobuf:"varint,2,opt,name=timeoutMs,proto3" json:"timeoutMs,omitempty"`
}

func (x *LockRequest) Reset() {
	*x = LockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharelock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockRequest) ProtoMessage() {}

func (x *LockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sharelock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockRequest.ProtoReflect.Descriptor instead.
func (*LockRequest) Descriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{2}
}

func (x *LockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LockRequest) GetTimeoutMs() int32 {
	if x != nil {
		return x.TimeoutMs
	}
	return 0
}

type LockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=sharelock.Status" json:"status,omitempty"`
}

func (x *LockResponse) Reset() {
	*x = LockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharelock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockResponse) ProtoMessage() {}

func (x *LockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sharelock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockResponse.ProtoReflect.Descriptor instead.
func (*LockResponse) Descriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{3}
}

func (x *LockResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

type UnlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UnlockRequest) Reset() {
	*x = UnlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharelock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockRequest) ProtoMessage() {}

func (x *UnlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sharelock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockRequest.ProtoReflect.Descriptor instead.
func (*UnlockRequest) Descriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{4}
}

func (x *UnlockRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type UnlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=sharelock.Status" json:"status,omitempty"`
}

func (x *UnlockResponse) Reset() {
	*x = UnlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sharelock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockResponse) ProtoMessage() {}

func (x *UnlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sharelock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockResponse.ProtoReflect.Descriptor instead.
func (*UnlockResponse) Descriptor() ([]byte, []int) {
	return file_sharelock_proto_rawDescGZIP(), []int{5}
}

func (x *UnlockResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Unknown
}

var File_sharelock_proto protoreflect.FileDescriptor

var file_sharelock_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x30, 0x0a, 0x14,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31,
	0x0a, 0x15, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73,
	0x22, 0x39, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3b,
	0x0a, 0x0e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x71, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x05, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x10, 0x06, 0x32, 0xdb,
	0x01, 0x0a, 0x10, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x6f, 0x63,
	0x6b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x6f,
	0x63, 0x6b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x42, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sharelock_proto_rawDescOnce sync.Once
	file_sharelock_proto_rawDescData = file_sharelock_proto_rawDesc
)

func file_sharelock_proto_rawDescGZIP() []byte {
	file_sharelock_proto_rawDescOnce.Do(func() {
		file_sharelock_proto_rawDescData = protoimpl.X.CompressGZIP(file_sharelock_proto_rawDescData)
	})
	return file_sharelock_proto_rawDescData
}

var file_sharelock_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sharelock_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sharelock_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: sharelock.Status
	(*ShareLockPingRequest)(nil),  // 1: sharelock.ShareLockPingRequest
	(*ShareLockPingResponse)(nil), // 2: sharelock.ShareLockPingResponse
	(*LockRequest)(nil),           // 3: sharelock.LockRequest
	(*LockResponse)(nil),          // 4: sharelock.LockResponse
	(*UnlockRequest)(nil),         // 5: sharelock.UnlockRequest
	(*UnlockResponse)(nil),        // 6: sharelock.UnlockResponse
}
var file_sharelock_proto_depIdxs = []int32{
	0, // 0: sharelock.LockResponse.status:type_name -> sharelock.Status
	0, // 1: sharelock.UnlockResponse.status:type_name -> sharelock.Status
	1, // 2: sharelock.ShareLockService.Ping:input_type -> sharelock.ShareLockPingRequest
	3, // 3: sharelock.ShareLockService.Lock:input_type -> sharelock.LockRequest
	5, // 4: sharelock.ShareLockService.Unlock:input_type -> sharelock.UnlockRequest
	2, // 5: sharelock.ShareLockService.Ping:output_type -> sharelock.ShareLockPingResponse
	4, // 6: sharelock.ShareLockService.Lock:output_type -> sharelock.LockResponse
	6, // 7: sharelock.ShareLockService.Unlock:output_type -> sharelock.UnlockResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sharelock_proto_init() }
func file_sharelock_proto_init() {
	if File_sharelock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sharelock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareLockPingRequest); i {
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
		file_sharelock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareLockPingResponse); i {
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
		file_sharelock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockRequest); i {
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
		file_sharelock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockResponse); i {
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
		file_sharelock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockRequest); i {
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
		file_sharelock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockResponse); i {
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
			RawDescriptor: file_sharelock_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sharelock_proto_goTypes,
		DependencyIndexes: file_sharelock_proto_depIdxs,
		EnumInfos:         file_sharelock_proto_enumTypes,
		MessageInfos:      file_sharelock_proto_msgTypes,
	}.Build()
	File_sharelock_proto = out.File
	file_sharelock_proto_rawDesc = nil
	file_sharelock_proto_goTypes = nil
	file_sharelock_proto_depIdxs = nil
}
