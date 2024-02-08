// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.2
// source: proto/server.proto

package proto

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

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os              string   `protobuf:"bytes,13,opt,name=os,proto3" json:"os,omitempty"`
	Platform        string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	PlatformVersion string   `protobuf:"bytes,2,opt,name=platform_version,json=platformVersion,proto3" json:"platform_version,omitempty"`
	Cpu             []string `protobuf:"bytes,3,rep,name=cpu,proto3" json:"cpu,omitempty"`
	MemTotal        uint64   `protobuf:"varint,4,opt,name=mem_total,json=memTotal,proto3" json:"mem_total,omitempty"`
	DiskTotal       uint64   `protobuf:"varint,5,opt,name=disk_total,json=diskTotal,proto3" json:"disk_total,omitempty"`
	SwapTotal       uint64   `protobuf:"varint,6,opt,name=swap_total,json=swapTotal,proto3" json:"swap_total,omitempty"`
	Arch            string   `protobuf:"bytes,7,opt,name=arch,proto3" json:"arch,omitempty"`
	Virtualization  string   `protobuf:"bytes,8,opt,name=virtualization,proto3" json:"virtualization,omitempty"`
	BootTime        uint64   `protobuf:"varint,9,opt,name=boot_time,json=bootTime,proto3" json:"boot_time,omitempty"`
	Ip              string   `protobuf:"bytes,10,opt,name=ip,proto3" json:"ip,omitempty"`
	CountryCode     string   `protobuf:"bytes,11,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Version         string   `protobuf:"bytes,12,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Host) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Host) GetPlatformVersion() string {
	if x != nil {
		return x.PlatformVersion
	}
	return ""
}

func (x *Host) GetCpu() []string {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Host) GetMemTotal() uint64 {
	if x != nil {
		return x.MemTotal
	}
	return 0
}

func (x *Host) GetDiskTotal() uint64 {
	if x != nil {
		return x.DiskTotal
	}
	return 0
}

func (x *Host) GetSwapTotal() uint64 {
	if x != nil {
		return x.SwapTotal
	}
	return 0
}

func (x *Host) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Host) GetVirtualization() string {
	if x != nil {
		return x.Virtualization
	}
	return ""
}

func (x *Host) GetBootTime() uint64 {
	if x != nil {
		return x.BootTime
	}
	return 0
}

func (x *Host) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Host) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Host) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu            float64 `protobuf:"fixed64,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	MemUsed        uint64  `protobuf:"varint,3,opt,name=mem_used,json=memUsed,proto3" json:"mem_used,omitempty"`
	SwapUsed       uint64  `protobuf:"varint,4,opt,name=swap_used,json=swapUsed,proto3" json:"swap_used,omitempty"`
	DiskUsed       uint64  `protobuf:"varint,5,opt,name=disk_used,json=diskUsed,proto3" json:"disk_used,omitempty"`
	NetInTransfer  uint64  `protobuf:"varint,6,opt,name=net_in_transfer,json=netInTransfer,proto3" json:"net_in_transfer,omitempty"`
	NetOutTransfer uint64  `protobuf:"varint,7,opt,name=net_out_transfer,json=netOutTransfer,proto3" json:"net_out_transfer,omitempty"`
	NetInSpeed     uint64  `protobuf:"varint,8,opt,name=net_in_speed,json=netInSpeed,proto3" json:"net_in_speed,omitempty"`
	NetOutSpeed    uint64  `protobuf:"varint,9,opt,name=net_out_speed,json=netOutSpeed,proto3" json:"net_out_speed,omitempty"`
	Uptime         uint64  `protobuf:"varint,10,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Load1          float64 `protobuf:"fixed64,11,opt,name=load1,proto3" json:"load1,omitempty"`
	Load5          float64 `protobuf:"fixed64,12,opt,name=load5,proto3" json:"load5,omitempty"`
	Load15         float64 `protobuf:"fixed64,13,opt,name=load15,proto3" json:"load15,omitempty"`
	TcpConnCount   uint64  `protobuf:"varint,14,opt,name=tcp_conn_count,json=tcpConnCount,proto3" json:"tcp_conn_count,omitempty"`
	UdpConnCount   uint64  `protobuf:"varint,15,opt,name=udp_conn_count,json=udpConnCount,proto3" json:"udp_conn_count,omitempty"`
	ProcessCount   uint64  `protobuf:"varint,16,opt,name=process_count,json=processCount,proto3" json:"process_count,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *State) GetCpu() float64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *State) GetMemUsed() uint64 {
	if x != nil {
		return x.MemUsed
	}
	return 0
}

func (x *State) GetSwapUsed() uint64 {
	if x != nil {
		return x.SwapUsed
	}
	return 0
}

func (x *State) GetDiskUsed() uint64 {
	if x != nil {
		return x.DiskUsed
	}
	return 0
}

func (x *State) GetNetInTransfer() uint64 {
	if x != nil {
		return x.NetInTransfer
	}
	return 0
}

func (x *State) GetNetOutTransfer() uint64 {
	if x != nil {
		return x.NetOutTransfer
	}
	return 0
}

func (x *State) GetNetInSpeed() uint64 {
	if x != nil {
		return x.NetInSpeed
	}
	return 0
}

func (x *State) GetNetOutSpeed() uint64 {
	if x != nil {
		return x.NetOutSpeed
	}
	return 0
}

func (x *State) GetUptime() uint64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *State) GetLoad1() float64 {
	if x != nil {
		return x.Load1
	}
	return 0
}

func (x *State) GetLoad5() float64 {
	if x != nil {
		return x.Load5
	}
	return 0
}

func (x *State) GetLoad15() float64 {
	if x != nil {
		return x.Load15
	}
	return 0
}

func (x *State) GetTcpConnCount() uint64 {
	if x != nil {
		return x.TcpConnCount
	}
	return 0
}

func (x *State) GetUdpConnCount() uint64 {
	if x != nil {
		return x.UdpConnCount
	}
	return 0
}

func (x *State) GetProcessCount() uint64 {
	if x != nil {
		return x.ProcessCount
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Data     string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	ServerId uint64 `protobuf:"varint,4,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Task) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Task) GetServerId() uint64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type       uint64  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Delay      float32 `protobuf:"fixed32,3,opt,name=delay,proto3" json:"delay,omitempty"`
	Data       string  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Successful bool    `protobuf:"varint,5,opt,name=successful,proto3" json:"successful,omitempty"`
	ServerId   uint64  `protobuf:"varint,6,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *TaskResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskResult) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TaskResult) GetDelay() float32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *TaskResult) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TaskResult) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

func (x *TaskResult) GetServerId() uint64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proced bool `protobuf:"varint,1,opt,name=proced,proto3" json:"proced,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{4}
}

func (x *Receipt) GetProced() bool {
	if x != nil {
		return x.Proced
	}
	return false
}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x02, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x70, 0x75, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x64, 0x69, 0x73, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x77, 0x61,
	0x70, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x77, 0x61, 0x70, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x0e,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd3,
	0x03, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x75, 0x73,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x77, 0x61, 0x70, 0x55, 0x73,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x5f, 0x6f,
	0x75, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x49, 0x6e, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x4f,
	0x75, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x6c, 0x6f, 0x61, 0x64, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x35, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x61, 0x64, 0x31, 0x35, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x61,
	0x64, 0x31, 0x35, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x63, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x64, 0x70,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x75, 0x64, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x07, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x32, 0xd7,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x33, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_server_proto_goTypes = []interface{}{
	(*Host)(nil),       // 0: proto.Host
	(*State)(nil),      // 1: proto.State
	(*Task)(nil),       // 2: proto.Task
	(*TaskResult)(nil), // 3: proto.TaskResult
	(*Receipt)(nil),    // 4: proto.Receipt
}
var file_proto_server_proto_depIdxs = []int32{
	1, // 0: proto.ServerService.ReportSystemState:input_type -> proto.State
	0, // 1: proto.ServerService.ReportSystemInfo:input_type -> proto.Host
	3, // 2: proto.ServerService.ReportTask:input_type -> proto.TaskResult
	0, // 3: proto.ServerService.RequestTask:input_type -> proto.Host
	4, // 4: proto.ServerService.ReportSystemState:output_type -> proto.Receipt
	4, // 5: proto.ServerService.ReportSystemInfo:output_type -> proto.Receipt
	4, // 6: proto.ServerService.ReportTask:output_type -> proto.Receipt
	2, // 7: proto.ServerService.RequestTask:output_type -> proto.Task
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResult); i {
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
		file_proto_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerServiceClient is the client API for ServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerServiceClient interface {
	ReportSystemState(ctx context.Context, in *State, opts ...grpc.CallOption) (*Receipt, error)
	ReportSystemInfo(ctx context.Context, in *Host, opts ...grpc.CallOption) (*Receipt, error)
	ReportTask(ctx context.Context, in *TaskResult, opts ...grpc.CallOption) (*Receipt, error)
	RequestTask(ctx context.Context, in *Host, opts ...grpc.CallOption) (ServerService_RequestTaskClient, error)
}

type serverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerServiceClient(cc grpc.ClientConnInterface) ServerServiceClient {
	return &serverServiceClient{cc}
}

func (c *serverServiceClient) ReportSystemState(ctx context.Context, in *State, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/proto.ServerService/ReportSystemState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ReportSystemInfo(ctx context.Context, in *Host, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/proto.ServerService/ReportSystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ReportTask(ctx context.Context, in *TaskResult, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, "/proto.ServerService/ReportTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) RequestTask(ctx context.Context, in *Host, opts ...grpc.CallOption) (ServerService_RequestTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerService_serviceDesc.Streams[0], "/proto.ServerService/RequestTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverServiceRequestTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerService_RequestTaskClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type serverServiceRequestTaskClient struct {
	grpc.ClientStream
}

func (x *serverServiceRequestTaskClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerServiceServer is the server API for ServerService service.
type ServerServiceServer interface {
	ReportSystemState(context.Context, *State) (*Receipt, error)
	ReportSystemInfo(context.Context, *Host) (*Receipt, error)
	ReportTask(context.Context, *TaskResult) (*Receipt, error)
	RequestTask(*Host, ServerService_RequestTaskServer) error
}

// UnimplementedServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerServiceServer struct {
}

func (*UnimplementedServerServiceServer) ReportSystemState(context.Context, *State) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportSystemState not implemented")
}
func (*UnimplementedServerServiceServer) ReportSystemInfo(context.Context, *Host) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportSystemInfo not implemented")
}
func (*UnimplementedServerServiceServer) ReportTask(context.Context, *TaskResult) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTask not implemented")
}
func (*UnimplementedServerServiceServer) RequestTask(*Host, ServerService_RequestTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestTask not implemented")
}

func RegisterServerServiceServer(s *grpc.Server, srv ServerServiceServer) {
	s.RegisterService(&_ServerService_serviceDesc, srv)
}

func _ServerService_ReportSystemState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ReportSystemState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServerService/ReportSystemState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ReportSystemState(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ReportSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Host)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ReportSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServerService/ReportSystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ReportSystemInfo(ctx, req.(*Host))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ReportTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ReportTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServerService/ReportTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ReportTask(ctx, req.(*TaskResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_RequestTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Host)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerServiceServer).RequestTask(m, &serverServiceRequestTaskServer{stream})
}

type ServerService_RequestTaskServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type serverServiceRequestTaskServer struct {
	grpc.ServerStream
}

func (x *serverServiceRequestTaskServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

var _ServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServerService",
	HandlerType: (*ServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportSystemState",
			Handler:    _ServerService_ReportSystemState_Handler,
		},
		{
			MethodName: "ReportSystemInfo",
			Handler:    _ServerService_ReportSystemInfo_Handler,
		},
		{
			MethodName: "ReportTask",
			Handler:    _ServerService_ReportTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestTask",
			Handler:       _ServerService_RequestTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/server.proto",
}
