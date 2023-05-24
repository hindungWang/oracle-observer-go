// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: blockchain_txn_vars_v1.proto

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

type BlockchainVarV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BlockchainVarV1) Reset() {
	*x = BlockchainVarV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_vars_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainVarV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainVarV1) ProtoMessage() {}

func (x *BlockchainVarV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_vars_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainVarV1.ProtoReflect.Descriptor instead.
func (*BlockchainVarV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_vars_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainVarV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlockchainVarV1) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BlockchainVarV1) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type BlockchainTxnVarsV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vars             []*BlockchainVarV1 `protobuf:"bytes,1,rep,name=vars,proto3" json:"vars,omitempty"`
	VersionPredicate uint32             `protobuf:"varint,2,opt,name=version_predicate,json=versionPredicate,proto3" json:"version_predicate,omitempty"`
	Proof            []byte             `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	MasterKey        []byte             `protobuf:"bytes,4,opt,name=master_key,json=masterKey,proto3" json:"master_key,omitempty"`
	KeyProof         []byte             `protobuf:"bytes,5,opt,name=key_proof,json=keyProof,proto3" json:"key_proof,omitempty"`
	Cancels          [][]byte           `protobuf:"bytes,6,rep,name=cancels,proto3" json:"cancels,omitempty"`
	Unsets           [][]byte           `protobuf:"bytes,7,rep,name=unsets,proto3" json:"unsets,omitempty"`
	Nonce            uint32             `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	MultiKeys        [][]byte           `protobuf:"bytes,9,rep,name=multi_keys,json=multiKeys,proto3" json:"multi_keys,omitempty"`
	MultiProofs      [][]byte           `protobuf:"bytes,10,rep,name=multi_proofs,json=multiProofs,proto3" json:"multi_proofs,omitempty"`
	MultiKeyProofs   [][]byte           `protobuf:"bytes,11,rep,name=multi_key_proofs,json=multiKeyProofs,proto3" json:"multi_key_proofs,omitempty"`
}

func (x *BlockchainTxnVarsV1) Reset() {
	*x = BlockchainTxnVarsV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_vars_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnVarsV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnVarsV1) ProtoMessage() {}

func (x *BlockchainTxnVarsV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_vars_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnVarsV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnVarsV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_vars_v1_proto_rawDescGZIP(), []int{1}
}

func (x *BlockchainTxnVarsV1) GetVars() []*BlockchainVarV1 {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetVersionPredicate() uint32 {
	if x != nil {
		return x.VersionPredicate
	}
	return 0
}

func (x *BlockchainTxnVarsV1) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetMasterKey() []byte {
	if x != nil {
		return x.MasterKey
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetKeyProof() []byte {
	if x != nil {
		return x.KeyProof
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetCancels() [][]byte {
	if x != nil {
		return x.Cancels
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetUnsets() [][]byte {
	if x != nil {
		return x.Unsets
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockchainTxnVarsV1) GetMultiKeys() [][]byte {
	if x != nil {
		return x.MultiKeys
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetMultiProofs() [][]byte {
	if x != nil {
		return x.MultiProofs
	}
	return nil
}

func (x *BlockchainTxnVarsV1) GetMultiKeyProofs() [][]byte {
	if x != nil {
		return x.MultiKeyProofs
	}
	return nil
}

var File_blockchain_txn_vars_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_vars_v1_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x76, 0x61, 0x72, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0x51, 0x0a, 0x11, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x76, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xfa, 0x02, 0x0a, 0x16, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x76, 0x61, 0x72,
	0x73, 0x5f, 0x76, 0x31, 0x12, 0x2d, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x76, 0x31, 0x52, 0x04, 0x76,
	0x61, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x6e, 0x73, 0x65, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x6e,
	0x73, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x4b, 0x65, 0x79,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6e, 0x64, 0x75, 0x6e, 0x67, 0x57, 0x61, 0x6e, 0x67,
	0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_vars_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_vars_v1_proto_rawDescData = file_blockchain_txn_vars_v1_proto_rawDesc
)

func file_blockchain_txn_vars_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_vars_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_vars_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_vars_v1_proto_rawDescData)
	})
	return file_blockchain_txn_vars_v1_proto_rawDescData
}

var file_blockchain_txn_vars_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blockchain_txn_vars_v1_proto_goTypes = []interface{}{
	(*BlockchainVarV1)(nil),     // 0: helium.blockchain_var_v1
	(*BlockchainTxnVarsV1)(nil), // 1: helium.blockchain_txn_vars_v1
}
var file_blockchain_txn_vars_v1_proto_depIdxs = []int32{
	0, // 0: helium.blockchain_txn_vars_v1.vars:type_name -> helium.blockchain_var_v1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blockchain_txn_vars_v1_proto_init() }
func file_blockchain_txn_vars_v1_proto_init() {
	if File_blockchain_txn_vars_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_vars_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainVarV1); i {
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
		file_blockchain_txn_vars_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnVarsV1); i {
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
			RawDescriptor: file_blockchain_txn_vars_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_vars_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_vars_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_vars_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_vars_v1_proto = out.File
	file_blockchain_txn_vars_v1_proto_rawDesc = nil
	file_blockchain_txn_vars_v1_proto_goTypes = nil
	file_blockchain_txn_vars_v1_proto_depIdxs = nil
}
