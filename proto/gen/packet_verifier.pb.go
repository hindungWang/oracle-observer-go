// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: packet_verifier.proto

package gen

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

type InvalidPacketReason int32

const (
	InvalidPacketReason_invalid_packet_reason_insufficient_balance InvalidPacketReason = 0
)

// Enum value maps for InvalidPacketReason.
var (
	InvalidPacketReason_name = map[int32]string{
		0: "invalid_packet_reason_insufficient_balance",
	}
	InvalidPacketReason_value = map[string]int32{
		"invalid_packet_reason_insufficient_balance": 0,
	}
)

func (x InvalidPacketReason) Enum() *InvalidPacketReason {
	p := new(InvalidPacketReason)
	*p = x
	return p
}

func (x InvalidPacketReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvalidPacketReason) Descriptor() protoreflect.EnumDescriptor {
	return file_packet_verifier_proto_enumTypes[0].Descriptor()
}

func (InvalidPacketReason) Type() protoreflect.EnumType {
	return &file_packet_verifier_proto_enumTypes[0]
}

func (x InvalidPacketReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvalidPacketReason.Descriptor instead.
func (InvalidPacketReason) EnumDescriptor() ([]byte, []int) {
	return file_packet_verifier_proto_rawDescGZIP(), []int{0}
}

type ValidPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadSize     uint32 `protobuf:"varint,1,opt,name=payload_size,json=payloadSize,proto3" json:"payload_size,omitempty"`
	Gateway         []byte `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	PayloadHash     []byte `protobuf:"bytes,3,opt,name=payload_hash,json=payloadHash,proto3" json:"payload_hash,omitempty"`
	NumDcs          uint32 `protobuf:"varint,4,opt,name=num_dcs,json=numDcs,proto3" json:"num_dcs,omitempty"`
	PacketTimestamp uint64 `protobuf:"varint,5,opt,name=packet_timestamp,json=packetTimestamp,proto3" json:"packet_timestamp,omitempty"`
}

func (x *ValidPacket) Reset() {
	*x = ValidPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_verifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidPacket) ProtoMessage() {}

func (x *ValidPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packet_verifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidPacket.ProtoReflect.Descriptor instead.
func (*ValidPacket) Descriptor() ([]byte, []int) {
	return file_packet_verifier_proto_rawDescGZIP(), []int{0}
}

func (x *ValidPacket) GetPayloadSize() uint32 {
	if x != nil {
		return x.PayloadSize
	}
	return 0
}

func (x *ValidPacket) GetGateway() []byte {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *ValidPacket) GetPayloadHash() []byte {
	if x != nil {
		return x.PayloadHash
	}
	return nil
}

func (x *ValidPacket) GetNumDcs() uint32 {
	if x != nil {
		return x.NumDcs
	}
	return 0
}

func (x *ValidPacket) GetPacketTimestamp() uint64 {
	if x != nil {
		return x.PacketTimestamp
	}
	return 0
}

type InvalidPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadSize uint32              `protobuf:"varint,1,opt,name=payload_size,json=payloadSize,proto3" json:"payload_size,omitempty"`
	Gateway     []byte              `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	PayloadHash []byte              `protobuf:"bytes,3,opt,name=payload_hash,json=payloadHash,proto3" json:"payload_hash,omitempty"`
	Reason      InvalidPacketReason `protobuf:"varint,4,opt,name=reason,proto3,enum=helium.packet_verifier.InvalidPacketReason" json:"reason,omitempty"`
}

func (x *InvalidPacket) Reset() {
	*x = InvalidPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_verifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidPacket) ProtoMessage() {}

func (x *InvalidPacket) ProtoReflect() protoreflect.Message {
	mi := &file_packet_verifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidPacket.ProtoReflect.Descriptor instead.
func (*InvalidPacket) Descriptor() ([]byte, []int) {
	return file_packet_verifier_proto_rawDescGZIP(), []int{1}
}

func (x *InvalidPacket) GetPayloadSize() uint32 {
	if x != nil {
		return x.PayloadSize
	}
	return 0
}

func (x *InvalidPacket) GetGateway() []byte {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *InvalidPacket) GetPayloadHash() []byte {
	if x != nil {
		return x.PayloadHash
	}
	return nil
}

func (x *InvalidPacket) GetReason() InvalidPacketReason {
	if x != nil {
		return x.Reason
	}
	return InvalidPacketReason_invalid_packet_reason_insufficient_balance
}

type ValidDataTransferSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey        []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	UploadBytes   uint64 `protobuf:"varint,2,opt,name=upload_bytes,json=uploadBytes,proto3" json:"upload_bytes,omitempty"`
	DownloadBytes uint64 `protobuf:"varint,3,opt,name=download_bytes,json=downloadBytes,proto3" json:"download_bytes,omitempty"`
	NumDcs        uint64 `protobuf:"varint,4,opt,name=num_dcs,json=numDcs,proto3" json:"num_dcs,omitempty"`
	Payer         []byte `protobuf:"bytes,5,opt,name=payer,proto3" json:"payer,omitempty"`
	// Timestamp in millis of the first ingest file we found a data transfer
	// session in
	FirstTimestamp uint64 `protobuf:"varint,6,opt,name=first_timestamp,json=firstTimestamp,proto3" json:"first_timestamp,omitempty"`
	// Timestamp in millis of the last ingest file we found a data transfer
	// session in
	LastTimestamp uint64 `protobuf:"varint,7,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
}

func (x *ValidDataTransferSession) Reset() {
	*x = ValidDataTransferSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_verifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidDataTransferSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidDataTransferSession) ProtoMessage() {}

func (x *ValidDataTransferSession) ProtoReflect() protoreflect.Message {
	mi := &file_packet_verifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidDataTransferSession.ProtoReflect.Descriptor instead.
func (*ValidDataTransferSession) Descriptor() ([]byte, []int) {
	return file_packet_verifier_proto_rawDescGZIP(), []int{2}
}

func (x *ValidDataTransferSession) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *ValidDataTransferSession) GetUploadBytes() uint64 {
	if x != nil {
		return x.UploadBytes
	}
	return 0
}

func (x *ValidDataTransferSession) GetDownloadBytes() uint64 {
	if x != nil {
		return x.DownloadBytes
	}
	return 0
}

func (x *ValidDataTransferSession) GetNumDcs() uint64 {
	if x != nil {
		return x.NumDcs
	}
	return 0
}

func (x *ValidDataTransferSession) GetPayer() []byte {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *ValidDataTransferSession) GetFirstTimestamp() uint64 {
	if x != nil {
		return x.FirstTimestamp
	}
	return 0
}

func (x *ValidDataTransferSession) GetLastTimestamp() uint64 {
	if x != nil {
		return x.LastTimestamp
	}
	return 0
}

var File_packet_verifier_proto protoreflect.FileDescriptor

var file_packet_verifier_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22,
	0xb2, 0x01, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x44, 0x63, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xff,
	0x01, 0x0a, 0x1b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x44, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2a, 0x47, 0x0a, 0x15, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x2a, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6e, 0x64, 0x75, 0x6e, 0x67, 0x57,
	0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packet_verifier_proto_rawDescOnce sync.Once
	file_packet_verifier_proto_rawDescData = file_packet_verifier_proto_rawDesc
)

func file_packet_verifier_proto_rawDescGZIP() []byte {
	file_packet_verifier_proto_rawDescOnce.Do(func() {
		file_packet_verifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_verifier_proto_rawDescData)
	})
	return file_packet_verifier_proto_rawDescData
}

var file_packet_verifier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_packet_verifier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packet_verifier_proto_goTypes = []interface{}{
	(InvalidPacketReason)(0),         // 0: helium.packet_verifier.invalid_packet_reason
	(*ValidPacket)(nil),              // 1: helium.packet_verifier.valid_packet
	(*InvalidPacket)(nil),            // 2: helium.packet_verifier.invalid_packet
	(*ValidDataTransferSession)(nil), // 3: helium.packet_verifier.valid_data_transfer_session
}
var file_packet_verifier_proto_depIdxs = []int32{
	0, // 0: helium.packet_verifier.invalid_packet.reason:type_name -> helium.packet_verifier.invalid_packet_reason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_packet_verifier_proto_init() }
func file_packet_verifier_proto_init() {
	if File_packet_verifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packet_verifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidPacket); i {
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
		file_packet_verifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidPacket); i {
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
		file_packet_verifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidDataTransferSession); i {
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
			RawDescriptor: file_packet_verifier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_verifier_proto_goTypes,
		DependencyIndexes: file_packet_verifier_proto_depIdxs,
		EnumInfos:         file_packet_verifier_proto_enumTypes,
		MessageInfos:      file_packet_verifier_proto_msgTypes,
	}.Build()
	File_packet_verifier_proto = out.File
	file_packet_verifier_proto_rawDesc = nil
	file_packet_verifier_proto_goTypes = nil
	file_packet_verifier_proto_depIdxs = nil
}