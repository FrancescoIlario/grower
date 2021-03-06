// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: scheduler.proto

package schedulerpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListSchedulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSchedulesRequest) Reset() {
	*x = ListSchedulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchedulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchedulesRequest) ProtoMessage() {}

func (x *ListSchedulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchedulesRequest.ProtoReflect.Descriptor instead.
func (*ListSchedulesRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{0}
}

type ListSchedulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedules []*Schedule `protobuf:"bytes,1,rep,name=Schedules,proto3" json:"Schedules,omitempty"`
}

func (x *ListSchedulesResponse) Reset() {
	*x = ListSchedulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchedulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchedulesResponse) ProtoMessage() {}

func (x *ListSchedulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchedulesResponse.ProtoReflect.Descriptor instead.
func (*ListSchedulesResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *ListSchedulesResponse) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type GetScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource id  of the Schedule to be returned.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetScheduleRequest) Reset() {
	*x = GetScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScheduleRequest) ProtoMessage() {}

func (x *GetScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScheduleRequest.ProtoReflect.Descriptor instead.
func (*GetScheduleRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *GetScheduleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenTime  *TimePoint `protobuf:"bytes,1,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
	CloseTime *TimePoint `protobuf:"bytes,2,opt,name=CloseTime,proto3" json:"CloseTime,omitempty"`
}

func (x *CreateScheduleRequest) Reset() {
	*x = CreateScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleRequest) ProtoMessage() {}

func (x *CreateScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleRequest.ProtoReflect.Descriptor instead.
func (*CreateScheduleRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{3}
}

func (x *CreateScheduleRequest) GetOpenTime() *TimePoint {
	if x != nil {
		return x.OpenTime
	}
	return nil
}

func (x *CreateScheduleRequest) GetCloseTime() *TimePoint {
	if x != nil {
		return x.CloseTime
	}
	return nil
}

type UpdateScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Schedule resource which replaces the resource on the server.
	Schedule *Schedule `protobuf:"bytes,1,opt,name=Schedule,proto3" json:"Schedule,omitempty"`
}

func (x *UpdateScheduleRequest) Reset() {
	*x = UpdateScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScheduleRequest) ProtoMessage() {}

func (x *UpdateScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScheduleRequest.ProtoReflect.Descriptor instead.
func (*UpdateScheduleRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateScheduleRequest) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type DeleteScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource id  of the Schedule to be deleted.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteScheduleRequest) Reset() {
	*x = DeleteScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScheduleRequest) ProtoMessage() {}

func (x *DeleteScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScheduleRequest.ProtoReflect.Descriptor instead.
func (*DeleteScheduleRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteScheduleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreationTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	OpenTime     *TimePoint           `protobuf:"bytes,3,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
	CloseTime    *TimePoint           `protobuf:"bytes,4,opt,name=CloseTime,proto3" json:"CloseTime,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{6}
}

func (x *Schedule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Schedule) GetCreationTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *Schedule) GetOpenTime() *TimePoint {
	if x != nil {
		return x.OpenTime
	}
	return nil
}

func (x *Schedule) GetCloseTime() *TimePoint {
	if x != nil {
		return x.CloseTime
	}
	return nil
}

type TimePoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hours   int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
}

func (x *TimePoint) Reset() {
	*x = TimePoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimePoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimePoint) ProtoMessage() {}

func (x *TimePoint) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimePoint.ProtoReflect.Descriptor instead.
func (*TimePoint) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{7}
}

func (x *TimePoint) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *TimePoint) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

var File_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x81, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x4f, 0x70,
	0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc4, 0x01, 0x0a, 0x08, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x65, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a,
	0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x13, 0xe2, 0xdf,
	0x1f, 0x0f, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x18, 0x18, 0x20,
	0x01, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x10,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x18, 0x3c, 0x20, 0x01, 0x52, 0x07,
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x32, 0xd3, 0x02, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x35, 0x48,
	0x03, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x63, 0x6f, 0x49, 0x6c, 0x61, 0x72, 0x69, 0x6f, 0x2f, 0x67, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_proto_rawDescData = file_scheduler_proto_rawDesc
)

func file_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_proto_rawDescData)
	})
	return file_scheduler_proto_rawDescData
}

var file_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_scheduler_proto_goTypes = []interface{}{
	(*ListSchedulesRequest)(nil),  // 0: schedulerpb.ListSchedulesRequest
	(*ListSchedulesResponse)(nil), // 1: schedulerpb.ListSchedulesResponse
	(*GetScheduleRequest)(nil),    // 2: schedulerpb.GetScheduleRequest
	(*CreateScheduleRequest)(nil), // 3: schedulerpb.CreateScheduleRequest
	(*UpdateScheduleRequest)(nil), // 4: schedulerpb.UpdateScheduleRequest
	(*DeleteScheduleRequest)(nil), // 5: schedulerpb.DeleteScheduleRequest
	(*Schedule)(nil),              // 6: schedulerpb.Schedule
	(*TimePoint)(nil),             // 7: schedulerpb.TimePoint
	(*timestamp.Timestamp)(nil),   // 8: google.protobuf.Timestamp
	(*empty.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_scheduler_proto_depIdxs = []int32{
	6,  // 0: schedulerpb.ListSchedulesResponse.Schedules:type_name -> schedulerpb.Schedule
	7,  // 1: schedulerpb.CreateScheduleRequest.OpenTime:type_name -> schedulerpb.TimePoint
	7,  // 2: schedulerpb.CreateScheduleRequest.CloseTime:type_name -> schedulerpb.TimePoint
	6,  // 3: schedulerpb.UpdateScheduleRequest.Schedule:type_name -> schedulerpb.Schedule
	8,  // 4: schedulerpb.Schedule.CreationTime:type_name -> google.protobuf.Timestamp
	7,  // 5: schedulerpb.Schedule.OpenTime:type_name -> schedulerpb.TimePoint
	7,  // 6: schedulerpb.Schedule.CloseTime:type_name -> schedulerpb.TimePoint
	0,  // 7: schedulerpb.ScheduleService.ListSchedules:input_type -> schedulerpb.ListSchedulesRequest
	2,  // 8: schedulerpb.ScheduleService.GetSchedule:input_type -> schedulerpb.GetScheduleRequest
	3,  // 9: schedulerpb.ScheduleService.CreateSchedule:input_type -> schedulerpb.CreateScheduleRequest
	5,  // 10: schedulerpb.ScheduleService.DeleteSchedule:input_type -> schedulerpb.DeleteScheduleRequest
	1,  // 11: schedulerpb.ScheduleService.ListSchedules:output_type -> schedulerpb.ListSchedulesResponse
	6,  // 12: schedulerpb.ScheduleService.GetSchedule:output_type -> schedulerpb.Schedule
	6,  // 13: schedulerpb.ScheduleService.CreateSchedule:output_type -> schedulerpb.Schedule
	9,  // 14: schedulerpb.ScheduleService.DeleteSchedule:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_scheduler_proto_init() }
func file_scheduler_proto_init() {
	if File_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchedulesRequest); i {
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
		file_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchedulesResponse); i {
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
		file_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScheduleRequest); i {
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
		file_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScheduleRequest); i {
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
		file_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScheduleRequest); i {
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
		file_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScheduleRequest); i {
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
		file_scheduler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_scheduler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimePoint); i {
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
			RawDescriptor: file_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_proto = out.File
	file_scheduler_proto_rawDesc = nil
	file_scheduler_proto_goTypes = nil
	file_scheduler_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	ListSchedules(ctx context.Context, in *ListSchedulesRequest, opts ...grpc.CallOption) (*ListSchedulesResponse, error)
	GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*Schedule, error)
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*Schedule, error)
	DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) ListSchedules(ctx context.Context, in *ListSchedulesRequest, opts ...grpc.CallOption) (*ListSchedulesResponse, error) {
	out := new(ListSchedulesResponse)
	err := c.cc.Invoke(ctx, "/schedulerpb.ScheduleService/ListSchedules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*Schedule, error) {
	out := new(Schedule)
	err := c.cc.Invoke(ctx, "/schedulerpb.ScheduleService/GetSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*Schedule, error) {
	out := new(Schedule)
	err := c.cc.Invoke(ctx, "/schedulerpb.ScheduleService/CreateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/schedulerpb.ScheduleService/DeleteSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
type ScheduleServiceServer interface {
	ListSchedules(context.Context, *ListSchedulesRequest) (*ListSchedulesResponse, error)
	GetSchedule(context.Context, *GetScheduleRequest) (*Schedule, error)
	CreateSchedule(context.Context, *CreateScheduleRequest) (*Schedule, error)
	DeleteSchedule(context.Context, *DeleteScheduleRequest) (*empty.Empty, error)
}

// UnimplementedScheduleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScheduleServiceServer struct {
}

func (*UnimplementedScheduleServiceServer) ListSchedules(context.Context, *ListSchedulesRequest) (*ListSchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchedules not implemented")
}
func (*UnimplementedScheduleServiceServer) GetSchedule(context.Context, *GetScheduleRequest) (*Schedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (*UnimplementedScheduleServiceServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*Schedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (*UnimplementedScheduleServiceServer) DeleteSchedule(context.Context, *DeleteScheduleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchedule not implemented")
}

func RegisterScheduleServiceServer(s *grpc.Server, srv ScheduleServiceServer) {
	s.RegisterService(&_ScheduleService_serviceDesc, srv)
}

func _ScheduleService_ListSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).ListSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulerpb.ScheduleService/ListSchedules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).ListSchedules(ctx, req.(*ListSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulerpb.ScheduleService/GetSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetSchedule(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulerpb.ScheduleService/CreateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_DeleteSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).DeleteSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulerpb.ScheduleService/DeleteSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).DeleteSchedule(ctx, req.(*DeleteScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScheduleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schedulerpb.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSchedules",
			Handler:    _ScheduleService_ListSchedules_Handler,
		},
		{
			MethodName: "GetSchedule",
			Handler:    _ScheduleService_GetSchedule_Handler,
		},
		{
			MethodName: "CreateSchedule",
			Handler:    _ScheduleService_CreateSchedule_Handler,
		},
		{
			MethodName: "DeleteSchedule",
			Handler:    _ScheduleService_DeleteSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}
