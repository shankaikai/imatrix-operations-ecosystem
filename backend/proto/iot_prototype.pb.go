// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/iot_prototype.proto

package operations_ecosys

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

type GateState_GatePosition int32

const (
	GateState_CLOSED  GateState_GatePosition = 0
	GateState_OPEN    GateState_GatePosition = 1
	GateState_ERROR   GateState_GatePosition = 2
	GateState_INITIAL GateState_GatePosition = 3
)

// Enum value maps for GateState_GatePosition.
var (
	GateState_GatePosition_name = map[int32]string{
		0: "CLOSED",
		1: "OPEN",
		2: "ERROR",
		3: "INITIAL",
	}
	GateState_GatePosition_value = map[string]int32{
		"CLOSED":  0,
		"OPEN":    1,
		"ERROR":   2,
		"INITIAL": 3,
	}
)

func (x GateState_GatePosition) Enum() *GateState_GatePosition {
	p := new(GateState_GatePosition)
	*p = x
	return p
}

func (x GateState_GatePosition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GateState_GatePosition) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_iot_prototype_proto_enumTypes[0].Descriptor()
}

func (GateState_GatePosition) Type() protoreflect.EnumType {
	return &file_proto_iot_prototype_proto_enumTypes[0]
}

func (x GateState_GatePosition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GateState_GatePosition.Descriptor instead.
func (GateState_GatePosition) EnumDescriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{1, 0}
}

type FireAlarmState_AlarmState int32

const (
	FireAlarmState_OFF   FireAlarmState_AlarmState = 0
	FireAlarmState_ON    FireAlarmState_AlarmState = 1
	FireAlarmState_ERROR FireAlarmState_AlarmState = 2
)

// Enum value maps for FireAlarmState_AlarmState.
var (
	FireAlarmState_AlarmState_name = map[int32]string{
		0: "OFF",
		1: "ON",
		2: "ERROR",
	}
	FireAlarmState_AlarmState_value = map[string]int32{
		"OFF":   0,
		"ON":    1,
		"ERROR": 2,
	}
)

func (x FireAlarmState_AlarmState) Enum() *FireAlarmState_AlarmState {
	p := new(FireAlarmState_AlarmState)
	*p = x
	return p
}

func (x FireAlarmState_AlarmState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FireAlarmState_AlarmState) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_iot_prototype_proto_enumTypes[1].Descriptor()
}

func (FireAlarmState_AlarmState) Type() protoreflect.EnumType {
	return &file_proto_iot_prototype_proto_enumTypes[1]
}

func (x FireAlarmState_AlarmState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FireAlarmState_AlarmState.Descriptor instead.
func (FireAlarmState_AlarmState) EnumDescriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{3, 0}
}

// The request message containing the stepper's name.
type Gate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Gate) Reset() {
	*x = Gate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_iot_prototype_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gate) ProtoMessage() {}

func (x *Gate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_iot_prototype_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gate.ProtoReflect.Descriptor instead.
func (*Gate) Descriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{0}
}

func (x *Gate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The response message containing the state of the stepper
type GateState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	State GateState_GatePosition `protobuf:"varint,2,opt,name=state,proto3,enum=gate_prototype.GateState_GatePosition" json:"state,omitempty"`
}

func (x *GateState) Reset() {
	*x = GateState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_iot_prototype_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateState) ProtoMessage() {}

func (x *GateState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_iot_prototype_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateState.ProtoReflect.Descriptor instead.
func (*GateState) Descriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{1}
}

func (x *GateState) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GateState) GetState() GateState_GatePosition {
	if x != nil {
		return x.State
	}
	return GateState_CLOSED
}

// The request message containing the fire alarm's id.
type FireAlarm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FireAlarm) Reset() {
	*x = FireAlarm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_iot_prototype_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FireAlarm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FireAlarm) ProtoMessage() {}

func (x *FireAlarm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_iot_prototype_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FireAlarm.ProtoReflect.Descriptor instead.
func (*FireAlarm) Descriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{2}
}

func (x *FireAlarm) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The response message containing the state of the stepper
type FireAlarmState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	State FireAlarmState_AlarmState `protobuf:"varint,2,opt,name=state,proto3,enum=gate_prototype.FireAlarmState_AlarmState" json:"state,omitempty"`
}

func (x *FireAlarmState) Reset() {
	*x = FireAlarmState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_iot_prototype_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FireAlarmState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FireAlarmState) ProtoMessage() {}

func (x *FireAlarmState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_iot_prototype_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FireAlarmState.ProtoReflect.Descriptor instead.
func (*FireAlarmState) Descriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{3}
}

func (x *FireAlarmState) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FireAlarmState) GetState() FireAlarmState_AlarmState {
	if x != nil {
		return x.State
	}
	return FireAlarmState_OFF
}

// The request message containing the fire alarm's id.
type CpuTemp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CpuTemp) Reset() {
	*x = CpuTemp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_iot_prototype_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuTemp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuTemp) ProtoMessage() {}

func (x *CpuTemp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_iot_prototype_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuTemp.ProtoReflect.Descriptor instead.
func (*CpuTemp) Descriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{4}
}

func (x *CpuTemp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The response message containing the state of the stepper
type CpuTempState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Temp float64 `protobuf:"fixed64,2,opt,name=temp,proto3" json:"temp,omitempty"`
}

func (x *CpuTempState) Reset() {
	*x = CpuTempState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_iot_prototype_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuTempState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuTempState) ProtoMessage() {}

func (x *CpuTempState) ProtoReflect() protoreflect.Message {
	mi := &file_proto_iot_prototype_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuTempState.ProtoReflect.Descriptor instead.
func (*CpuTempState) Descriptor() ([]byte, []int) {
	return file_proto_iot_prototype_proto_rawDescGZIP(), []int{5}
}

func (x *CpuTempState) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CpuTempState) GetTemp() float64 {
	if x != nil {
		return x.Temp
	}
	return 0
}

var File_proto_iot_prototype_proto protoreflect.FileDescriptor

var file_proto_iot_prototype_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6f, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x47,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x3c, 0x0a, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x22, 0x1b, 0x0a,
	0x09, 0x46, 0x69, 0x72, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x46,
	0x69, 0x72, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69,
	0x72, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x28,
	0x0a, 0x0a, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x19, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x54,
	0x65, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0c, 0x43, 0x70, 0x75, 0x54, 0x65, 0x6d, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x32, 0xca, 0x02, 0x0a, 0x1a, 0x49, 0x6f, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x1a, 0x19, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0c, 0x53,
	0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x72, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69,
	0x72, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x1a, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x43, 0x70, 0x75, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x70, 0x75, 0x54, 0x65,
	0x6d, 0x70, 0x1a, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x70, 0x75, 0x54, 0x65, 0x6d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_iot_prototype_proto_rawDescOnce sync.Once
	file_proto_iot_prototype_proto_rawDescData = file_proto_iot_prototype_proto_rawDesc
)

func file_proto_iot_prototype_proto_rawDescGZIP() []byte {
	file_proto_iot_prototype_proto_rawDescOnce.Do(func() {
		file_proto_iot_prototype_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_iot_prototype_proto_rawDescData)
	})
	return file_proto_iot_prototype_proto_rawDescData
}

var file_proto_iot_prototype_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_iot_prototype_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_iot_prototype_proto_goTypes = []interface{}{
	(GateState_GatePosition)(0),    // 0: gate_prototype.GateState.GatePosition
	(FireAlarmState_AlarmState)(0), // 1: gate_prototype.FireAlarmState.AlarmState
	(*Gate)(nil),                   // 2: gate_prototype.Gate
	(*GateState)(nil),              // 3: gate_prototype.GateState
	(*FireAlarm)(nil),              // 4: gate_prototype.FireAlarm
	(*FireAlarmState)(nil),         // 5: gate_prototype.FireAlarmState
	(*CpuTemp)(nil),                // 6: gate_prototype.CpuTemp
	(*CpuTempState)(nil),           // 7: gate_prototype.CpuTempState
}
var file_proto_iot_prototype_proto_depIdxs = []int32{
	0, // 0: gate_prototype.GateState.state:type_name -> gate_prototype.GateState.GatePosition
	1, // 1: gate_prototype.FireAlarmState.state:type_name -> gate_prototype.FireAlarmState.AlarmState
	2, // 2: gate_prototype.IotControlPrototypeService.GetGateState:input_type -> gate_prototype.Gate
	3, // 3: gate_prototype.IotControlPrototypeService.SetGateState:input_type -> gate_prototype.GateState
	4, // 4: gate_prototype.IotControlPrototypeService.GetFireAlarmState:input_type -> gate_prototype.FireAlarm
	6, // 5: gate_prototype.IotControlPrototypeService.GetCpuTemp:input_type -> gate_prototype.CpuTemp
	3, // 6: gate_prototype.IotControlPrototypeService.GetGateState:output_type -> gate_prototype.GateState
	3, // 7: gate_prototype.IotControlPrototypeService.SetGateState:output_type -> gate_prototype.GateState
	5, // 8: gate_prototype.IotControlPrototypeService.GetFireAlarmState:output_type -> gate_prototype.FireAlarmState
	7, // 9: gate_prototype.IotControlPrototypeService.GetCpuTemp:output_type -> gate_prototype.CpuTempState
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_iot_prototype_proto_init() }
func file_proto_iot_prototype_proto_init() {
	if File_proto_iot_prototype_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_iot_prototype_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gate); i {
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
		file_proto_iot_prototype_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateState); i {
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
		file_proto_iot_prototype_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FireAlarm); i {
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
		file_proto_iot_prototype_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FireAlarmState); i {
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
		file_proto_iot_prototype_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuTemp); i {
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
		file_proto_iot_prototype_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuTempState); i {
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
			RawDescriptor: file_proto_iot_prototype_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_iot_prototype_proto_goTypes,
		DependencyIndexes: file_proto_iot_prototype_proto_depIdxs,
		EnumInfos:         file_proto_iot_prototype_proto_enumTypes,
		MessageInfos:      file_proto_iot_prototype_proto_msgTypes,
	}.Build()
	File_proto_iot_prototype_proto = out.File
	file_proto_iot_prototype_proto_rawDesc = nil
	file_proto_iot_prototype_proto_goTypes = nil
	file_proto_iot_prototype_proto_depIdxs = nil
}
