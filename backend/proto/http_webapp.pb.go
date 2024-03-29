// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/http_webapp.proto

package operations_ecosys

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HTTPMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//status==0 denotes no issues. status > 0 denotes some issue.
	Status   int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Value    string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValueArr []string `protobuf:"bytes,3,rep,name=valueArr,proto3" json:"valueArr,omitempty"`
}

func (x *HTTPMessage) Reset() {
	*x = HTTPMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_http_webapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPMessage) ProtoMessage() {}

func (x *HTTPMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_http_webapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPMessage.ProtoReflect.Descriptor instead.
func (*HTTPMessage) Descriptor() ([]byte, []int) {
	return file_proto_http_webapp_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPMessage) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HTTPMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *HTTPMessage) GetValueArr() []string {
	if x != nil {
		return x.ValueArr
	}
	return nil
}

type HTTPAssignmentsGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Twan       string `protobuf:"bytes,1,opt,name=twan,proto3" json:"twan,omitempty"`
	TeleUserId int64  `protobuf:"varint,2,opt,name=tele_user_id,json=teleUserId,proto3" json:"tele_user_id,omitempty"`
	StartDate  string `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate    string `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *HTTPAssignmentsGetRequest) Reset() {
	*x = HTTPAssignmentsGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_http_webapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPAssignmentsGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPAssignmentsGetRequest) ProtoMessage() {}

func (x *HTTPAssignmentsGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_http_webapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPAssignmentsGetRequest.ProtoReflect.Descriptor instead.
func (*HTTPAssignmentsGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_http_webapp_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPAssignmentsGetRequest) GetTwan() string {
	if x != nil {
		return x.Twan
	}
	return ""
}

func (x *HTTPAssignmentsGetRequest) GetTeleUserId() int64 {
	if x != nil {
		return x.TeleUserId
	}
	return 0
}

func (x *HTTPAssignmentsGetRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *HTTPAssignmentsGetRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type HTTPAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *HTTPMessage          `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Rosters  []*HTTPRosterResponse `protobuf:"bytes,2,rep,name=rosters,proto3" json:"rosters,omitempty"`
}

func (x *HTTPAssignmentResponse) Reset() {
	*x = HTTPAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_http_webapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPAssignmentResponse) ProtoMessage() {}

func (x *HTTPAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_http_webapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPAssignmentResponse.ProtoReflect.Descriptor instead.
func (*HTTPAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_proto_http_webapp_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPAssignmentResponse) GetResponse() *HTTPMessage {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *HTTPAssignmentResponse) GetRosters() []*HTTPRosterResponse {
	if x != nil {
		return x.Rosters
	}
	return nil
}

type HTTPRosterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AifsId        int64    `protobuf:"varint,2,opt,name=aifs_id,json=aifsId,proto3" json:"aifs_id,omitempty"`
	StartDatetime string   `protobuf:"bytes,3,opt,name=startDatetime,proto3" json:"startDatetime,omitempty"`
	EndDatetime   string   `protobuf:"bytes,4,opt,name=endDatetime,proto3" json:"endDatetime,omitempty"`
	Addresses     []string `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *HTTPRosterResponse) Reset() {
	*x = HTTPRosterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_http_webapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRosterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRosterResponse) ProtoMessage() {}

func (x *HTTPRosterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_http_webapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRosterResponse.ProtoReflect.Descriptor instead.
func (*HTTPRosterResponse) Descriptor() ([]byte, []int) {
	return file_proto_http_webapp_proto_rawDescGZIP(), []int{3}
}

func (x *HTTPRosterResponse) GetAifsId() int64 {
	if x != nil {
		return x.AifsId
	}
	return 0
}

func (x *HTTPRosterResponse) GetStartDatetime() string {
	if x != nil {
		return x.StartDatetime
	}
	return ""
}

func (x *HTTPRosterResponse) GetEndDatetime() string {
	if x != nil {
		return x.EndDatetime
	}
	return ""
}

func (x *HTTPRosterResponse) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type HTTPReportPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Twan                  string `protobuf:"bytes,1,opt,name=twan,proto3" json:"twan,omitempty"`
	TeleUserId            int64  `protobuf:"varint,2,opt,name=tele_user_id,json=teleUserId,proto3" json:"tele_user_id,omitempty"`
	Title                 string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	ReportType            string `protobuf:"bytes,4,opt,name=report_type,json=reportType,proto3" json:"report_type,omitempty"`
	Address               string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Date                  string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Time                  string `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Details               string `protobuf:"bytes,8,opt,name=details,proto3" json:"details,omitempty"`
	IsPeopleInjured       bool   `protobuf:"varint,9,opt,name=isPeopleInjured,proto3" json:"isPeopleInjured,omitempty"`
	InjuryDetails         string `protobuf:"bytes,10,opt,name=injuryDetails,proto3" json:"injuryDetails,omitempty"`
	IsPoliceNotified      bool   `protobuf:"varint,11,opt,name=isPoliceNotified,proto3" json:"isPoliceNotified,omitempty"`
	IsPropertyStolen      bool   `protobuf:"varint,12,opt,name=isPropertyStolen,proto3" json:"isPropertyStolen,omitempty"`
	PropertyStolenDetails string `protobuf:"bytes,13,opt,name=propertyStolenDetails,proto3" json:"propertyStolenDetails,omitempty"`
	IsActionTaken         bool   `protobuf:"varint,14,opt,name=isActionTaken,proto3" json:"isActionTaken,omitempty"`
	ActionDetails         string `protobuf:"bytes,15,opt,name=actionDetails,proto3" json:"actionDetails,omitempty"`
}

func (x *HTTPReportPostRequest) Reset() {
	*x = HTTPReportPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_http_webapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPReportPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPReportPostRequest) ProtoMessage() {}

func (x *HTTPReportPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_http_webapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPReportPostRequest.ProtoReflect.Descriptor instead.
func (*HTTPReportPostRequest) Descriptor() ([]byte, []int) {
	return file_proto_http_webapp_proto_rawDescGZIP(), []int{4}
}

func (x *HTTPReportPostRequest) GetTwan() string {
	if x != nil {
		return x.Twan
	}
	return ""
}

func (x *HTTPReportPostRequest) GetTeleUserId() int64 {
	if x != nil {
		return x.TeleUserId
	}
	return 0
}

func (x *HTTPReportPostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *HTTPReportPostRequest) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

func (x *HTTPReportPostRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HTTPReportPostRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *HTTPReportPostRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *HTTPReportPostRequest) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *HTTPReportPostRequest) GetIsPeopleInjured() bool {
	if x != nil {
		return x.IsPeopleInjured
	}
	return false
}

func (x *HTTPReportPostRequest) GetInjuryDetails() string {
	if x != nil {
		return x.InjuryDetails
	}
	return ""
}

func (x *HTTPReportPostRequest) GetIsPoliceNotified() bool {
	if x != nil {
		return x.IsPoliceNotified
	}
	return false
}

func (x *HTTPReportPostRequest) GetIsPropertyStolen() bool {
	if x != nil {
		return x.IsPropertyStolen
	}
	return false
}

func (x *HTTPReportPostRequest) GetPropertyStolenDetails() string {
	if x != nil {
		return x.PropertyStolenDetails
	}
	return ""
}

func (x *HTTPReportPostRequest) GetIsActionTaken() bool {
	if x != nil {
		return x.IsActionTaken
	}
	return false
}

func (x *HTTPReportPostRequest) GetActionDetails() string {
	if x != nil {
		return x.ActionDetails
	}
	return ""
}

type HTTPRegistrationFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	TeleUserId       int64  `protobuf:"varint,2,opt,name=tele_user_id,json=teleUserId,proto3" json:"tele_user_id,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email            string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber      string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	LoginString      string `protobuf:"bytes,6,opt,name=loginString,proto3" json:"loginString,omitempty"`
	HasedLoginString string `protobuf:"bytes,7,opt,name=hasedLoginString,proto3" json:"hasedLoginString,omitempty"`
	IsPartTime       bool   `protobuf:"varint,8,opt,name=isPartTime,proto3" json:"isPartTime,omitempty"`
	TeleHandle       string `protobuf:"bytes,9,opt,name=tele_handle,json=teleHandle,proto3" json:"tele_handle,omitempty"`
}

func (x *HTTPRegistrationFormRequest) Reset() {
	*x = HTTPRegistrationFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_http_webapp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRegistrationFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRegistrationFormRequest) ProtoMessage() {}

func (x *HTTPRegistrationFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_http_webapp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRegistrationFormRequest.ProtoReflect.Descriptor instead.
func (*HTTPRegistrationFormRequest) Descriptor() ([]byte, []int) {
	return file_proto_http_webapp_proto_rawDescGZIP(), []int{5}
}

func (x *HTTPRegistrationFormRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *HTTPRegistrationFormRequest) GetTeleUserId() int64 {
	if x != nil {
		return x.TeleUserId
	}
	return 0
}

func (x *HTTPRegistrationFormRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HTTPRegistrationFormRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *HTTPRegistrationFormRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *HTTPRegistrationFormRequest) GetLoginString() string {
	if x != nil {
		return x.LoginString
	}
	return ""
}

func (x *HTTPRegistrationFormRequest) GetHasedLoginString() string {
	if x != nil {
		return x.HasedLoginString
	}
	return ""
}

func (x *HTTPRegistrationFormRequest) GetIsPartTime() bool {
	if x != nil {
		return x.IsPartTime
	}
	return false
}

func (x *HTTPRegistrationFormRequest) GetTeleHandle() string {
	if x != nil {
		return x.TeleHandle
	}
	return ""
}

var File_proto_http_webapp_proto protoreflect.FileDescriptor

var file_proto_http_webapp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65, 0x62,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x72, 0x72, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x72, 0x72, 0x22, 0x89, 0x01,
	0x0a, 0x19, 0x48, 0x54, 0x54, 0x50, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x77, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x77, 0x61, 0x6e, 0x12,
	0x20, 0x0a, 0x0c, 0x74, 0x65, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x16, 0x48, 0x54,
	0x54, 0x50, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65,
	0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x48, 0x54, 0x54, 0x50, 0x52, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x69, 0x66, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x69, 0x66, 0x73, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x8a, 0x04, 0x0a, 0x15,
	0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x77, 0x61, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x77, 0x61, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x65, 0x6c,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x65, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x69, 0x73, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x6a, 0x75, 0x72, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x49, 0x6e, 0x6a, 0x75, 0x72, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x6a, 0x75,
	0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x6a, 0x75, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x73,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53, 0x74, 0x6f, 0x6c, 0x65, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x53, 0x74, 0x6f, 0x6c, 0x65, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x53, 0x74, 0x6f, 0x6c, 0x65, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53,
	0x74, 0x6f, 0x6c, 0x65, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x6b,
	0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xae, 0x02, 0x0a, 0x1b, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0c,
	0x74, 0x65, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x10,
	0x68, 0x61, 0x73, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x61, 0x73, 0x65, 0x64, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x50, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x50, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6c, 0x65,
	0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x65, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x32, 0x9a, 0x03, 0x0a, 0x0e, 0x57, 0x65,
	0x62, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x87, 0x01, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x12, 0x26,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65,
	0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x73, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x57, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x12,
	0x22, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x70,
	0x70, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x88, 0x01, 0x0a, 0x1e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x12, 0x28,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x14, 0x5a, 0x12, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_http_webapp_proto_rawDescOnce sync.Once
	file_proto_http_webapp_proto_rawDescData = file_proto_http_webapp_proto_rawDesc
)

func file_proto_http_webapp_proto_rawDescGZIP() []byte {
	file_proto_http_webapp_proto_rawDescOnce.Do(func() {
		file_proto_http_webapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_http_webapp_proto_rawDescData)
	})
	return file_proto_http_webapp_proto_rawDescData
}

var file_proto_http_webapp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_http_webapp_proto_goTypes = []interface{}{
	(*HTTPMessage)(nil),                 // 0: http_webapp.HTTPMessage
	(*HTTPAssignmentsGetRequest)(nil),   // 1: http_webapp.HTTPAssignmentsGetRequest
	(*HTTPAssignmentResponse)(nil),      // 2: http_webapp.HTTPAssignmentResponse
	(*HTTPRosterResponse)(nil),          // 3: http_webapp.HTTPRosterResponse
	(*HTTPReportPostRequest)(nil),       // 4: http_webapp.HTTPReportPostRequest
	(*HTTPRegistrationFormRequest)(nil), // 5: http_webapp.HTTPRegistrationFormRequest
}
var file_proto_http_webapp_proto_depIdxs = []int32{
	0, // 0: http_webapp.HTTPAssignmentResponse.response:type_name -> http_webapp.HTTPMessage
	3, // 1: http_webapp.HTTPAssignmentResponse.rosters:type_name -> http_webapp.HTTPRosterResponse
	1, // 2: http_webapp.WebAppServices.GetRosterAssignmentsForWebApp:input_type -> http_webapp.HTTPAssignmentsGetRequest
	4, // 3: http_webapp.WebAppServices.PostWReportFromWebApp:input_type -> http_webapp.HTTPReportPostRequest
	5, // 4: http_webapp.WebAppServices.PostRegistrationFormFromWebApp:input_type -> http_webapp.HTTPRegistrationFormRequest
	2, // 5: http_webapp.WebAppServices.GetRosterAssignmentsForWebApp:output_type -> http_webapp.HTTPAssignmentResponse
	0, // 6: http_webapp.WebAppServices.PostWReportFromWebApp:output_type -> http_webapp.HTTPMessage
	0, // 7: http_webapp.WebAppServices.PostRegistrationFormFromWebApp:output_type -> http_webapp.HTTPMessage
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_http_webapp_proto_init() }
func file_proto_http_webapp_proto_init() {
	if File_proto_http_webapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_http_webapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPMessage); i {
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
		file_proto_http_webapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPAssignmentsGetRequest); i {
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
		file_proto_http_webapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPAssignmentResponse); i {
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
		file_proto_http_webapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRosterResponse); i {
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
		file_proto_http_webapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPReportPostRequest); i {
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
		file_proto_http_webapp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRegistrationFormRequest); i {
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
			RawDescriptor: file_proto_http_webapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_http_webapp_proto_goTypes,
		DependencyIndexes: file_proto_http_webapp_proto_depIdxs,
		MessageInfos:      file_proto_http_webapp_proto_msgTypes,
	}.Build()
	File_proto_http_webapp_proto = out.File
	file_proto_http_webapp_proto_rawDesc = nil
	file_proto_http_webapp_proto_goTypes = nil
	file_proto_http_webapp_proto_depIdxs = nil
}
