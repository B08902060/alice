// Copyright © 2021 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/getamis/alice/crypto/bip32/child/message.proto

package child

import (
	circuit "github.com/getamis/alice/crypto/circuit"
	ot "github.com/getamis/alice/crypto/ot"
	zkproof "github.com/getamis/alice/crypto/zkproof"
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

type Type int32

const (
	Type_Initial        Type = 0
	Type_OtReceiver     Type = 1
	Type_OtSendResponse Type = 2
	Type_EncH           Type = 3
	Type_Sh2Hash        Type = 4
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Initial",
		1: "OtReceiver",
		2: "OtSendResponse",
		3: "EncH",
		4: "Sh2Hash",
	}
	Type_value = map[string]int32{
		"Initial":        0,
		"OtReceiver":     1,
		"OtSendResponse": 2,
		"EncH":           3,
		"Sh2Hash":        4,
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
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_github_com_getamis_alice_crypto_bip32_child_message_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=getamis.alice.crypto.bip32.child.Type" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Body:
	//	*Message_Initial
	//	*Message_OtReceiver
	//	*Message_OtSendResponse
	//	*Message_EncH
	//	*Message_Sh2Hash
	Body isMessage_Body `protobuf_oneof:"body"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Initial
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Message) GetInitial() *BodyInitial {
	if x, ok := x.GetBody().(*Message_Initial); ok {
		return x.Initial
	}
	return nil
}

func (x *Message) GetOtReceiver() *BodyOtReceiver {
	if x, ok := x.GetBody().(*Message_OtReceiver); ok {
		return x.OtReceiver
	}
	return nil
}

func (x *Message) GetOtSendResponse() *BodyOtSendResponse {
	if x, ok := x.GetBody().(*Message_OtSendResponse); ok {
		return x.OtSendResponse
	}
	return nil
}

func (x *Message) GetEncH() *BodyEncH {
	if x, ok := x.GetBody().(*Message_EncH); ok {
		return x.EncH
	}
	return nil
}

func (x *Message) GetSh2Hash() *BodySh2Hash {
	if x, ok := x.GetBody().(*Message_Sh2Hash); ok {
		return x.Sh2Hash
	}
	return nil
}

type isMessage_Body interface {
	isMessage_Body()
}

type Message_Initial struct {
	Initial *BodyInitial `protobuf:"bytes,3,opt,name=initial,proto3,oneof"`
}

type Message_OtReceiver struct {
	OtReceiver *BodyOtReceiver `protobuf:"bytes,4,opt,name=otReceiver,proto3,oneof"`
}

type Message_OtSendResponse struct {
	OtSendResponse *BodyOtSendResponse `protobuf:"bytes,5,opt,name=otSendResponse,proto3,oneof"`
}

type Message_EncH struct {
	EncH *BodyEncH `protobuf:"bytes,6,opt,name=encH,proto3,oneof"`
}

type Message_Sh2Hash struct {
	Sh2Hash *BodySh2Hash `protobuf:"bytes,7,opt,name=sh2Hash,proto3,oneof"`
}

func (*Message_Initial) isMessage_Body() {}

func (*Message_OtReceiver) isMessage_Body() {}

func (*Message_OtSendResponse) isMessage_Body() {}

func (*Message_EncH) isMessage_Body() {}

func (*Message_Sh2Hash) isMessage_Body() {}

type BodyInitial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtRecMsg       *ot.OtReceiverMessage         `protobuf:"bytes,1,opt,name=otRecMsg,proto3" json:"otRecMsg,omitempty"`
	GarcirMsg      *circuit.GarbleCircuitMessage `protobuf:"bytes,2,opt,name=garcirMsg,proto3" json:"garcirMsg,omitempty"`
	OtherInfoWire  [][]byte                      `protobuf:"bytes,3,rep,name=otherInfoWire,proto3" json:"otherInfoWire,omitempty"`
	PubKey         []byte                        `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	PubKeyN        []byte                        `protobuf:"bytes,5,opt,name=pubKeyN,proto3" json:"pubKeyN,omitempty"`
	ShareGProofMsg *zkproof.SchnorrProofMessage  `protobuf:"bytes,6,opt,name=shareGProofMsg,proto3" json:"shareGProofMsg,omitempty"`
}

func (x *BodyInitial) Reset() {
	*x = BodyInitial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInitial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInitial) ProtoMessage() {}

func (x *BodyInitial) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInitial.ProtoReflect.Descriptor instead.
func (*BodyInitial) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{1}
}

func (x *BodyInitial) GetOtRecMsg() *ot.OtReceiverMessage {
	if x != nil {
		return x.OtRecMsg
	}
	return nil
}

func (x *BodyInitial) GetGarcirMsg() *circuit.GarbleCircuitMessage {
	if x != nil {
		return x.GarcirMsg
	}
	return nil
}

func (x *BodyInitial) GetOtherInfoWire() [][]byte {
	if x != nil {
		return x.OtherInfoWire
	}
	return nil
}

func (x *BodyInitial) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *BodyInitial) GetPubKeyN() []byte {
	if x != nil {
		return x.PubKeyN
	}
	return nil
}

func (x *BodyInitial) GetShareGProofMsg() *zkproof.SchnorrProofMessage {
	if x != nil {
		return x.ShareGProofMsg
	}
	return nil
}

type BodyOtReceiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtExtReceiveMsg *ot.OtExtReceiveMessage `protobuf:"bytes,1,opt,name=otExtReceiveMsg,proto3" json:"otExtReceiveMsg,omitempty"`
}

func (x *BodyOtReceiver) Reset() {
	*x = BodyOtReceiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyOtReceiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyOtReceiver) ProtoMessage() {}

func (x *BodyOtReceiver) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyOtReceiver.ProtoReflect.Descriptor instead.
func (*BodyOtReceiver) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{2}
}

func (x *BodyOtReceiver) GetOtExtReceiveMsg() *ot.OtExtReceiveMessage {
	if x != nil {
		return x.OtExtReceiveMsg
	}
	return nil
}

type BodyOtSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtExtSendResponseMsg *ot.OtExtSendResponseMessage `protobuf:"bytes,1,opt,name=otExtSendResponseMsg,proto3" json:"otExtSendResponseMsg,omitempty"`
}

func (x *BodyOtSendResponse) Reset() {
	*x = BodyOtSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyOtSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyOtSendResponse) ProtoMessage() {}

func (x *BodyOtSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyOtSendResponse.ProtoReflect.Descriptor instead.
func (*BodyOtSendResponse) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{3}
}

func (x *BodyOtSendResponse) GetOtExtSendResponseMsg() *ot.OtExtSendResponseMessage {
	if x != nil {
		return x.OtExtSendResponseMsg
	}
	return nil
}

type BodyEncH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncH []byte `protobuf:"bytes,1,opt,name=encH,proto3" json:"encH,omitempty"`
}

func (x *BodyEncH) Reset() {
	*x = BodyEncH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyEncH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyEncH) ProtoMessage() {}

func (x *BodyEncH) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyEncH.ProtoReflect.Descriptor instead.
func (*BodyEncH) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{4}
}

func (x *BodyEncH) GetEncH() []byte {
	if x != nil {
		return x.EncH
	}
	return nil
}

type BodySh2Hash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Sh2Hash []byte `protobuf:"bytes,2,opt,name=sh2Hash,proto3" json:"sh2Hash,omitempty"`
}

func (x *BodySh2Hash) Reset() {
	*x = BodySh2Hash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodySh2Hash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodySh2Hash) ProtoMessage() {}

func (x *BodySh2Hash) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodySh2Hash.ProtoReflect.Descriptor instead.
func (*BodySh2Hash) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP(), []int{5}
}

func (x *BodySh2Hash) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *BodySh2Hash) GetSh2Hash() []byte {
	if x != nil {
		return x.Sh2Hash
	}
	return nil
}

var File_github_com_getamis_alice_crypto_bip32_child_message_proto protoreflect.FileDescriptor

var file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x65, 0x74,
	0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x1a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69,
	0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6f,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61,
	0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x7a, 0x6b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x03,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62,
	0x69, 0x70, 0x33, 0x32, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73,
	0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69,
	0x70, 0x33, 0x32, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x12, 0x52, 0x0a, 0x0a, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69, 0x70, 0x33,
	0x32, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x0e, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67,
	0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2e,
	0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x65, 0x6e, 0x63, 0x48, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2e,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x6e, 0x63, 0x48, 0x48, 0x00,
	0x52, 0x04, 0x65, 0x6e, 0x63, 0x48, 0x12, 0x49, 0x0a, 0x07, 0x73, 0x68, 0x32, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62,
	0x69, 0x70, 0x33, 0x32, 0x2e, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x53,
	0x68, 0x32, 0x48, 0x61, 0x73, 0x68, 0x48, 0x00, 0x52, 0x07, 0x73, 0x68, 0x32, 0x48, 0x61, 0x73,
	0x68, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xda, 0x02, 0x0a, 0x0b, 0x42, 0x6f,
	0x64, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x08, 0x6f, 0x74, 0x52,
	0x65, 0x63, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x65,
	0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x6f, 0x74, 0x2e, 0x4f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x4d, 0x73,
	0x67, 0x12, 0x50, 0x0a, 0x09, 0x67, 0x61, 0x72, 0x63, 0x69, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x72, 0x63,
	0x75, 0x69, 0x74, 0x2e, 0x67, 0x61, 0x72, 0x62, 0x6c, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x67, 0x61, 0x72, 0x63, 0x69, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x57, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x69, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x4e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x4e, 0x12, 0x59, 0x0a, 0x0e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x47, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4d, 0x73, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x7a, 0x6b, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x2e, 0x53, 0x63, 0x68, 0x6e, 0x6f, 0x72, 0x72, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x4d, 0x73, 0x67, 0x22, 0x68, 0x0a, 0x0e, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0f, 0x6f, 0x74, 0x45, 0x78,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6f, 0x74, 0x2e, 0x4f, 0x74, 0x45, 0x78,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0f, 0x6f, 0x74, 0x45, 0x78, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x73, 0x67,
	0x22, 0x7b, 0x0a, 0x12, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x14, 0x6f, 0x74, 0x45, 0x78, 0x74, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6f, 0x74, 0x2e, 0x4f,
	0x74, 0x45, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x14, 0x6f, 0x74, 0x45, 0x78, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x1e, 0x0a,
	0x08, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x6e, 0x63, 0x48, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x63,
	0x48, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x65, 0x6e, 0x63, 0x48, 0x22, 0x3f, 0x0a,
	0x0b, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x68, 0x32, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x32, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x68, 0x32, 0x48, 0x61, 0x73, 0x68, 0x2a, 0x4e,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x6e, 0x63, 0x48, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x68, 0x32, 0x48, 0x61, 0x73, 0x68, 0x10, 0x04, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x62, 0x69, 0x70, 0x33, 0x32, 0x2f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescOnce sync.Once
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescData = file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDesc
)

func file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescGZIP() []byte {
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescOnce.Do(func() {
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescData)
	})
	return file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDescData
}

var file_github_com_getamis_alice_crypto_bip32_child_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_getamis_alice_crypto_bip32_child_message_proto_goTypes = []interface{}{
	(Type)(0),                            // 0: getamis.alice.crypto.bip32.child.Type
	(*Message)(nil),                      // 1: getamis.alice.crypto.bip32.child.Message
	(*BodyInitial)(nil),                  // 2: getamis.alice.crypto.bip32.child.BodyInitial
	(*BodyOtReceiver)(nil),               // 3: getamis.alice.crypto.bip32.child.BodyOtReceiver
	(*BodyOtSendResponse)(nil),           // 4: getamis.alice.crypto.bip32.child.BodyOtSendResponse
	(*BodyEncH)(nil),                     // 5: getamis.alice.crypto.bip32.child.BodyEncH
	(*BodySh2Hash)(nil),                  // 6: getamis.alice.crypto.bip32.child.BodySh2Hash
	(*ot.OtReceiverMessage)(nil),         // 7: getamis.alice.crypto.ot.OtReceiverMessage
	(*circuit.GarbleCircuitMessage)(nil), // 8: getamis.alice.crypto.circuit.garbleCircuitMessage
	(*zkproof.SchnorrProofMessage)(nil),  // 9: getamis.alice.crypto.zkproof.SchnorrProofMessage
	(*ot.OtExtReceiveMessage)(nil),       // 10: getamis.alice.crypto.ot.OtExtReceiveMessage
	(*ot.OtExtSendResponseMessage)(nil),  // 11: getamis.alice.crypto.ot.OtExtSendResponseMessage
}
var file_github_com_getamis_alice_crypto_bip32_child_message_proto_depIdxs = []int32{
	0,  // 0: getamis.alice.crypto.bip32.child.Message.type:type_name -> getamis.alice.crypto.bip32.child.Type
	2,  // 1: getamis.alice.crypto.bip32.child.Message.initial:type_name -> getamis.alice.crypto.bip32.child.BodyInitial
	3,  // 2: getamis.alice.crypto.bip32.child.Message.otReceiver:type_name -> getamis.alice.crypto.bip32.child.BodyOtReceiver
	4,  // 3: getamis.alice.crypto.bip32.child.Message.otSendResponse:type_name -> getamis.alice.crypto.bip32.child.BodyOtSendResponse
	5,  // 4: getamis.alice.crypto.bip32.child.Message.encH:type_name -> getamis.alice.crypto.bip32.child.BodyEncH
	6,  // 5: getamis.alice.crypto.bip32.child.Message.sh2Hash:type_name -> getamis.alice.crypto.bip32.child.BodySh2Hash
	7,  // 6: getamis.alice.crypto.bip32.child.BodyInitial.otRecMsg:type_name -> getamis.alice.crypto.ot.OtReceiverMessage
	8,  // 7: getamis.alice.crypto.bip32.child.BodyInitial.garcirMsg:type_name -> getamis.alice.crypto.circuit.garbleCircuitMessage
	9,  // 8: getamis.alice.crypto.bip32.child.BodyInitial.shareGProofMsg:type_name -> getamis.alice.crypto.zkproof.SchnorrProofMessage
	10, // 9: getamis.alice.crypto.bip32.child.BodyOtReceiver.otExtReceiveMsg:type_name -> getamis.alice.crypto.ot.OtExtReceiveMessage
	11, // 10: getamis.alice.crypto.bip32.child.BodyOtSendResponse.otExtSendResponseMsg:type_name -> getamis.alice.crypto.ot.OtExtSendResponseMessage
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_github_com_getamis_alice_crypto_bip32_child_message_proto_init() }
func file_github_com_getamis_alice_crypto_bip32_child_message_proto_init() {
	if File_github_com_getamis_alice_crypto_bip32_child_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInitial); i {
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
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyOtReceiver); i {
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
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyOtSendResponse); i {
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
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyEncH); i {
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
		file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodySh2Hash); i {
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
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Initial)(nil),
		(*Message_OtReceiver)(nil),
		(*Message_OtSendResponse)(nil),
		(*Message_EncH)(nil),
		(*Message_Sh2Hash)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_getamis_alice_crypto_bip32_child_message_proto_goTypes,
		DependencyIndexes: file_github_com_getamis_alice_crypto_bip32_child_message_proto_depIdxs,
		EnumInfos:         file_github_com_getamis_alice_crypto_bip32_child_message_proto_enumTypes,
		MessageInfos:      file_github_com_getamis_alice_crypto_bip32_child_message_proto_msgTypes,
	}.Build()
	File_github_com_getamis_alice_crypto_bip32_child_message_proto = out.File
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_rawDesc = nil
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_goTypes = nil
	file_github_com_getamis_alice_crypto_bip32_child_message_proto_depIdxs = nil
}
