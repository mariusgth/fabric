// Code generated by protoc-gen-go.
// source: orderer/ab.proto
// DO NOT EDIT!

/*
Package orderer is a generated protocol buffer package.

It is generated from these files:
	orderer/ab.proto

It has these top-level messages:
	BroadcastResponse
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
	SeekInfo
	Acknowledgement
	DeliverUpdate
	DeliverResponse
*/
package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/hyperledger/fabric/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN               Status = 0
	Status_SUCCESS               Status = 200
	Status_BAD_REQUEST           Status = 400
	Status_FORBIDDEN             Status = 403
	Status_NOT_FOUND             Status = 404
	Status_INTERNAL_SERVER_ERROR Status = 500
	Status_SERVICE_UNAVAILABLE   Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":               0,
	"SUCCESS":               200,
	"BAD_REQUEST":           400,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"INTERNAL_SERVER_ERROR": 500,
	"SERVICE_UNAVAILABLE":   503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ConfigurationItem_ConfigurationType int32

const (
	ConfigurationItem_Policy  ConfigurationItem_ConfigurationType = 0
	ConfigurationItem_Chain   ConfigurationItem_ConfigurationType = 1
	ConfigurationItem_Orderer ConfigurationItem_ConfigurationType = 2
	ConfigurationItem_Fabric  ConfigurationItem_ConfigurationType = 3
)

var ConfigurationItem_ConfigurationType_name = map[int32]string{
	0: "Policy",
	1: "Chain",
	2: "Orderer",
	3: "Fabric",
}
var ConfigurationItem_ConfigurationType_value = map[string]int32{
	"Policy":  0,
	"Chain":   1,
	"Orderer": 2,
	"Fabric":  3,
}

func (x ConfigurationItem_ConfigurationType) String() string {
	return proto.EnumName(ConfigurationItem_ConfigurationType_name, int32(x))
}
func (ConfigurationItem_ConfigurationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// Start may be specified to a specific block number, or may be request from the newest or oldest available
// The start location is always inclusive, so the first reply from NEWEST will contain the newest block at the time
// of reception, it will must not wait until a new block is created.  Similarly, when SPECIFIED, and SpecifiedNumber = 10
// The first block received must be block 10, not block 11
type SeekInfo_StartType int32

const (
	SeekInfo_NEWEST    SeekInfo_StartType = 0
	SeekInfo_OLDEST    SeekInfo_StartType = 1
	SeekInfo_SPECIFIED SeekInfo_StartType = 2
)

var SeekInfo_StartType_name = map[int32]string{
	0: "NEWEST",
	1: "OLDEST",
	2: "SPECIFIED",
}
var SeekInfo_StartType_value = map[string]int32{
	"NEWEST":    0,
	"OLDEST":    1,
	"SPECIFIED": 2,
}

func (x SeekInfo_StartType) String() string {
	return proto.EnumName(SeekInfo_StartType_name, int32(x))
}
func (SeekInfo_StartType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type BroadcastResponse struct {
	Status Status `protobuf:"varint,1,opt,name=Status,enum=orderer.Status" json:"Status,omitempty"`
}

func (m *BroadcastResponse) Reset()                    { *m = BroadcastResponse{} }
func (m *BroadcastResponse) String() string            { return proto.CompactTextString(m) }
func (*BroadcastResponse) ProtoMessage()               {}
func (*BroadcastResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ConfigurationEnvelope is designed to contain _all_ configuration for a chain with no dependency
// on previous configuration transactions.
//
// It is generated with the following scheme:
//   1. Retrieve the existing configuration
//   2. Note the highest configuration sequence number, store it an increment it by one
//   3. Modify desired ConfigurationItems, setting each LastModified to the stored and incremented sequence number
//   4. Update SignedConfigurationItem with appropriate signatures over the modified ConfigurationItem
//     a) Each signature is of type ConfigurationSignature
//     b) The ConfigurationSignature signature is over the concatenation of signatureHeader and the ConfigurationItem header
//   5. Submit new Configuration for ordering in Envelope signed by submitter
//     a) The common.Envelope common.Payload has data set to the marshaled ConfigurationEnvelope
//     b) The common.Envelope common.Payload has a header of type common.Header.Type.CONFIGURATION_TRANSACTION
//
// The configuration manager will verify:
//   1. All configuration items and the envelope refer to the correct chain
//   2. Some configuration item has been added or modified
//   3. No existing configuration item has been ommitted
//   4. All configuration changes have a LastModification of one more than the last configuration's sequence
//   5. All configuration changes satisfy the corresponding modification policy
type ConfigurationEnvelope struct {
	Items    []*SignedConfigurationItem `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	ChainID  []byte                     `protobuf:"bytes,2,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Sequence uint64                     `protobuf:"varint,3,opt,name=Sequence" json:"Sequence,omitempty"`
}

func (m *ConfigurationEnvelope) Reset()                    { *m = ConfigurationEnvelope{} }
func (m *ConfigurationEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ConfigurationEnvelope) ProtoMessage()               {}
func (*ConfigurationEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConfigurationEnvelope) GetItems() []*SignedConfigurationItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// This message may change slightly depending on the finalization of signature schemes for transactions
type SignedConfigurationItem struct {
	ConfigurationItem []byte                    `protobuf:"bytes,1,opt,name=ConfigurationItem,proto3" json:"ConfigurationItem,omitempty"`
	Signatures        []*ConfigurationSignature `protobuf:"bytes,2,rep,name=Signatures" json:"Signatures,omitempty"`
}

func (m *SignedConfigurationItem) Reset()                    { *m = SignedConfigurationItem{} }
func (m *SignedConfigurationItem) String() string            { return proto.CompactTextString(m) }
func (*SignedConfigurationItem) ProtoMessage()               {}
func (*SignedConfigurationItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SignedConfigurationItem) GetSignatures() []*ConfigurationSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type ConfigurationItem struct {
	Header             *common.ChainHeader                 `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Type               ConfigurationItem_ConfigurationType `protobuf:"varint,2,opt,name=Type,enum=orderer.ConfigurationItem_ConfigurationType" json:"Type,omitempty"`
	LastModified       uint64                              `protobuf:"varint,3,opt,name=LastModified" json:"LastModified,omitempty"`
	ModificationPolicy string                              `protobuf:"bytes,4,opt,name=ModificationPolicy" json:"ModificationPolicy,omitempty"`
	Key                string                              `protobuf:"bytes,5,opt,name=Key" json:"Key,omitempty"`
	Value              []byte                              `protobuf:"bytes,6,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *ConfigurationItem) Reset()                    { *m = ConfigurationItem{} }
func (m *ConfigurationItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigurationItem) ProtoMessage()               {}
func (*ConfigurationItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConfigurationItem) GetHeader() *common.ChainHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ConfigurationSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signatureHeader,proto3" json:"signatureHeader,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfigurationSignature) Reset()                    { *m = ConfigurationSignature{} }
func (m *ConfigurationSignature) String() string            { return proto.CompactTextString(m) }
func (*ConfigurationSignature) ProtoMessage()               {}
func (*ConfigurationSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	// Types that are valid to be assigned to Type:
	//	*Policy_SignaturePolicy
	Type isPolicy_Type `protobuf_oneof:"Type"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isPolicy_Type interface {
	isPolicy_Type()
}

type Policy_SignaturePolicy struct {
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=SignaturePolicy,oneof"`
}

func (*Policy_SignaturePolicy) isPolicy_Type() {}

func (m *Policy) GetType() isPolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Policy) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := m.GetType().(*Policy_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Policy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Policy_OneofMarshaler, _Policy_OneofUnmarshaler, _Policy_OneofSizer, []interface{}{
		(*Policy_SignaturePolicy)(nil),
	}
}

func _Policy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Policy)
	// Type
	switch x := m.Type.(type) {
	case *Policy_SignaturePolicy:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SignaturePolicy); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Policy.Type has unexpected type %T", x)
	}
	return nil
}

func _Policy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Policy)
	switch tag {
	case 1: // Type.SignaturePolicy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicyEnvelope)
		err := b.DecodeMessage(msg)
		m.Type = &Policy_SignaturePolicy{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Policy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Policy)
	// Type
	switch x := m.Type.(type) {
	case *Policy_SignaturePolicy:
		s := proto.Size(x.SignaturePolicy)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
type SignaturePolicyEnvelope struct {
	Version    int32            `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Policy     *SignaturePolicy `protobuf:"bytes,2,opt,name=Policy" json:"Policy,omitempty"`
	Identities [][]byte         `protobuf:"bytes,3,rep,name=Identities,proto3" json:"Identities,omitempty"`
}

func (m *SignaturePolicyEnvelope) Reset()                    { *m = SignaturePolicyEnvelope{} }
func (m *SignaturePolicyEnvelope) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicyEnvelope) ProtoMessage()               {}
func (*SignaturePolicyEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SignaturePolicyEnvelope) GetPolicy() *SignaturePolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
type SignaturePolicy struct {
	// Types that are valid to be assigned to Type:
	//	*SignaturePolicy_SignedBy
	//	*SignaturePolicy_From
	Type isSignaturePolicy_Type `protobuf_oneof:"Type"`
}

func (m *SignaturePolicy) Reset()                    { *m = SignaturePolicy{} }
func (m *SignaturePolicy) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy) ProtoMessage()               {}
func (*SignaturePolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isSignaturePolicy_Type interface {
	isSignaturePolicy_Type()
}

type SignaturePolicy_SignedBy struct {
	SignedBy int32 `protobuf:"varint,1,opt,name=SignedBy,oneof"`
}
type SignaturePolicy_From struct {
	From *SignaturePolicy_NOutOf `protobuf:"bytes,2,opt,name=From,oneof"`
}

func (*SignaturePolicy_SignedBy) isSignaturePolicy_Type() {}
func (*SignaturePolicy_From) isSignaturePolicy_Type()     {}

func (m *SignaturePolicy) GetType() isSignaturePolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SignaturePolicy) GetSignedBy() int32 {
	if x, ok := m.GetType().(*SignaturePolicy_SignedBy); ok {
		return x.SignedBy
	}
	return 0
}

func (m *SignaturePolicy) GetFrom() *SignaturePolicy_NOutOf {
	if x, ok := m.GetType().(*SignaturePolicy_From); ok {
		return x.From
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SignaturePolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SignaturePolicy_OneofMarshaler, _SignaturePolicy_OneofUnmarshaler, _SignaturePolicy_OneofSizer, []interface{}{
		(*SignaturePolicy_SignedBy)(nil),
		(*SignaturePolicy_From)(nil),
	}
}

func _SignaturePolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_From:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.From); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SignaturePolicy.Type has unexpected type %T", x)
	}
	return nil
}

func _SignaturePolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SignaturePolicy)
	switch tag {
	case 1: // Type.SignedBy
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &SignaturePolicy_SignedBy{int32(x)}
		return true, err
	case 2: // Type.From
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicy_NOutOf)
		err := b.DecodeMessage(msg)
		m.Type = &SignaturePolicy_From{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SignaturePolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_From:
		s := proto.Size(x.From)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SignaturePolicy_NOutOf struct {
	N        int32              `protobuf:"varint,1,opt,name=N" json:"N,omitempty"`
	Policies []*SignaturePolicy `protobuf:"bytes,2,rep,name=Policies" json:"Policies,omitempty"`
}

func (m *SignaturePolicy_NOutOf) Reset()                    { *m = SignaturePolicy_NOutOf{} }
func (m *SignaturePolicy_NOutOf) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy_NOutOf) ProtoMessage()               {}
func (*SignaturePolicy_NOutOf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *SignaturePolicy_NOutOf) GetPolicies() []*SignaturePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type SeekInfo struct {
	Start           SeekInfo_StartType `protobuf:"varint,1,opt,name=Start,enum=orderer.SeekInfo_StartType" json:"Start,omitempty"`
	SpecifiedNumber uint64             `protobuf:"varint,2,opt,name=SpecifiedNumber" json:"SpecifiedNumber,omitempty"`
	WindowSize      uint64             `protobuf:"varint,3,opt,name=WindowSize" json:"WindowSize,omitempty"`
	ChainID         []byte             `protobuf:"bytes,4,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
}

func (m *SeekInfo) Reset()                    { *m = SeekInfo{} }
func (m *SeekInfo) String() string            { return proto.CompactTextString(m) }
func (*SeekInfo) ProtoMessage()               {}
func (*SeekInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Acknowledgement struct {
	Number uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
}

func (m *Acknowledgement) Reset()                    { *m = Acknowledgement{} }
func (m *Acknowledgement) String() string            { return proto.CompactTextString(m) }
func (*Acknowledgement) ProtoMessage()               {}
func (*Acknowledgement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// The update message either causes a seek to a new stream start with a new window, or acknowledges a received block and advances the base of the window
type DeliverUpdate struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverUpdate_Acknowledgement
	//	*DeliverUpdate_Seek
	Type isDeliverUpdate_Type `protobuf_oneof:"Type"`
}

func (m *DeliverUpdate) Reset()                    { *m = DeliverUpdate{} }
func (m *DeliverUpdate) String() string            { return proto.CompactTextString(m) }
func (*DeliverUpdate) ProtoMessage()               {}
func (*DeliverUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type isDeliverUpdate_Type interface {
	isDeliverUpdate_Type()
}

type DeliverUpdate_Acknowledgement struct {
	Acknowledgement *Acknowledgement `protobuf:"bytes,1,opt,name=Acknowledgement,oneof"`
}
type DeliverUpdate_Seek struct {
	Seek *SeekInfo `protobuf:"bytes,2,opt,name=Seek,oneof"`
}

func (*DeliverUpdate_Acknowledgement) isDeliverUpdate_Type() {}
func (*DeliverUpdate_Seek) isDeliverUpdate_Type()            {}

func (m *DeliverUpdate) GetType() isDeliverUpdate_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverUpdate) GetAcknowledgement() *Acknowledgement {
	if x, ok := m.GetType().(*DeliverUpdate_Acknowledgement); ok {
		return x.Acknowledgement
	}
	return nil
}

func (m *DeliverUpdate) GetSeek() *SeekInfo {
	if x, ok := m.GetType().(*DeliverUpdate_Seek); ok {
		return x.Seek
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeliverUpdate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeliverUpdate_OneofMarshaler, _DeliverUpdate_OneofUnmarshaler, _DeliverUpdate_OneofSizer, []interface{}{
		(*DeliverUpdate_Acknowledgement)(nil),
		(*DeliverUpdate_Seek)(nil),
	}
}

func _DeliverUpdate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeliverUpdate)
	// Type
	switch x := m.Type.(type) {
	case *DeliverUpdate_Acknowledgement:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Acknowledgement); err != nil {
			return err
		}
	case *DeliverUpdate_Seek:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Seek); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeliverUpdate.Type has unexpected type %T", x)
	}
	return nil
}

func _DeliverUpdate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeliverUpdate)
	switch tag {
	case 1: // Type.Acknowledgement
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Acknowledgement)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverUpdate_Acknowledgement{msg}
		return true, err
	case 2: // Type.Seek
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SeekInfo)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverUpdate_Seek{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeliverUpdate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeliverUpdate)
	// Type
	switch x := m.Type.(type) {
	case *DeliverUpdate_Acknowledgement:
		s := proto.Size(x.Acknowledgement)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DeliverUpdate_Seek:
		s := proto.Size(x.Seek)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DeliverResponse struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverResponse_Error
	//	*DeliverResponse_Block
	Type isDeliverResponse_Type `protobuf_oneof:"Type"`
}

func (m *DeliverResponse) Reset()                    { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string            { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()               {}
func (*DeliverResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Error struct {
	Error Status `protobuf:"varint,1,opt,name=Error,enum=orderer.Status,oneof"`
}
type DeliverResponse_Block struct {
	Block *common.Block `protobuf:"bytes,2,opt,name=Block,oneof"`
}

func (*DeliverResponse_Error) isDeliverResponse_Type() {}
func (*DeliverResponse_Block) isDeliverResponse_Type() {}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverResponse) GetError() Status {
	if x, ok := m.GetType().(*DeliverResponse_Error); ok {
		return x.Error
	}
	return Status_UNKNOWN
}

func (m *DeliverResponse) GetBlock() *common.Block {
	if x, ok := m.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeliverResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeliverResponse_OneofMarshaler, _DeliverResponse_OneofUnmarshaler, _DeliverResponse_OneofSizer, []interface{}{
		(*DeliverResponse_Error)(nil),
		(*DeliverResponse_Block)(nil),
	}
}

func _DeliverResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Error:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Error))
	case *DeliverResponse_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeliverResponse.Type has unexpected type %T", x)
	}
	return nil
}

func _DeliverResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeliverResponse)
	switch tag {
	case 1: // Type.Error
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &DeliverResponse_Error{Status(x)}
		return true, err
	case 2: // Type.Block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.Block)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverResponse_Block{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeliverResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Error:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Error))
	case *DeliverResponse_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BroadcastResponse)(nil), "orderer.BroadcastResponse")
	proto.RegisterType((*ConfigurationEnvelope)(nil), "orderer.ConfigurationEnvelope")
	proto.RegisterType((*SignedConfigurationItem)(nil), "orderer.SignedConfigurationItem")
	proto.RegisterType((*ConfigurationItem)(nil), "orderer.ConfigurationItem")
	proto.RegisterType((*ConfigurationSignature)(nil), "orderer.ConfigurationSignature")
	proto.RegisterType((*Policy)(nil), "orderer.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "orderer.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "orderer.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "orderer.SignaturePolicy.NOutOf")
	proto.RegisterType((*SeekInfo)(nil), "orderer.SeekInfo")
	proto.RegisterType((*Acknowledgement)(nil), "orderer.Acknowledgement")
	proto.RegisterType((*DeliverUpdate)(nil), "orderer.DeliverUpdate")
	proto.RegisterType((*DeliverResponse)(nil), "orderer.DeliverResponse")
	proto.RegisterEnum("orderer.Status", Status_name, Status_value)
	proto.RegisterEnum("orderer.ConfigurationItem_ConfigurationType", ConfigurationItem_ConfigurationType_name, ConfigurationItem_ConfigurationType_value)
	proto.RegisterEnum("orderer.SeekInfo_StartType", SeekInfo_StartType_name, SeekInfo_StartType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AtomicBroadcast service

type AtomicBroadcastClient interface {
	// broadcast receives a reply of Acknowledgement for each common.Envelope in order, indicating success or type of failure
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error)
	// deliver first requires an update containing a seek message, then a stream of block replies is received.
	// The receiver may choose to send an Acknowledgement for any block number it receives, however Acknowledgements must never be more than WindowSize apart
	// To avoid latency, clients will likely acknowledge before the WindowSize has been exhausted, preventing the server from stopping and waiting for an Acknowledgement
	Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error)
}

type atomicBroadcastClient struct {
	cc *grpc.ClientConn
}

func NewAtomicBroadcastClient(cc *grpc.ClientConn) AtomicBroadcastClient {
	return &atomicBroadcastClient{cc}
}

func (c *atomicBroadcastClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[0], c.cc, "/orderer.AtomicBroadcast/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastBroadcastClient{stream}
	return x, nil
}

type AtomicBroadcast_BroadcastClient interface {
	Send(*common.Envelope) error
	Recv() (*BroadcastResponse, error)
	grpc.ClientStream
}

type atomicBroadcastBroadcastClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastBroadcastClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastClient) Recv() (*BroadcastResponse, error) {
	m := new(BroadcastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *atomicBroadcastClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[1], c.cc, "/orderer.AtomicBroadcast/Deliver", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastDeliverClient{stream}
	return x, nil
}

type AtomicBroadcast_DeliverClient interface {
	Send(*DeliverUpdate) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type atomicBroadcastDeliverClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastDeliverClient) Send(m *DeliverUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AtomicBroadcast service

type AtomicBroadcastServer interface {
	// broadcast receives a reply of Acknowledgement for each common.Envelope in order, indicating success or type of failure
	Broadcast(AtomicBroadcast_BroadcastServer) error
	// deliver first requires an update containing a seek message, then a stream of block replies is received.
	// The receiver may choose to send an Acknowledgement for any block number it receives, however Acknowledgements must never be more than WindowSize apart
	// To avoid latency, clients will likely acknowledge before the WindowSize has been exhausted, preventing the server from stopping and waiting for an Acknowledgement
	Deliver(AtomicBroadcast_DeliverServer) error
}

func RegisterAtomicBroadcastServer(s *grpc.Server, srv AtomicBroadcastServer) {
	s.RegisterService(&_AtomicBroadcast_serviceDesc, srv)
}

func _AtomicBroadcast_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Broadcast(&atomicBroadcastBroadcastServer{stream})
}

type AtomicBroadcast_BroadcastServer interface {
	Send(*BroadcastResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type atomicBroadcastBroadcastServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastBroadcastServer) Send(m *BroadcastResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AtomicBroadcast_Deliver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Deliver(&atomicBroadcastDeliverServer{stream})
}

type AtomicBroadcast_DeliverServer interface {
	Send(*DeliverResponse) error
	Recv() (*DeliverUpdate, error)
	grpc.ServerStream
}

type atomicBroadcastDeliverServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastDeliverServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverServer) Recv() (*DeliverUpdate, error) {
	m := new(DeliverUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AtomicBroadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.AtomicBroadcast",
	HandlerType: (*AtomicBroadcastServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _AtomicBroadcast_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Deliver",
			Handler:       _AtomicBroadcast_Deliver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("orderer/ab.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x55, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0xce, 0xd4, 0xb9, 0x34, 0x27, 0x69, 0xe3, 0x4e, 0xd9, 0xae, 0x15, 0x56, 0x6c, 0x64, 0x09,
	0x6d, 0x80, 0x2a, 0x59, 0xc2, 0xe5, 0x09, 0x69, 0xc9, 0xc5, 0x51, 0xac, 0x0d, 0xce, 0x32, 0x6e,
	0x5a, 0x89, 0x07, 0x82, 0x63, 0x4f, 0x5a, 0x6b, 0x13, 0x4f, 0xb0, 0x9d, 0x5d, 0x95, 0x77, 0x10,
	0x12, 0x2f, 0x48, 0xc0, 0xbf, 0xe1, 0x8d, 0x17, 0xfe, 0x00, 0xff, 0x04, 0x89, 0x57, 0x34, 0xe3,
	0x4b, 0x6e, 0xcd, 0x93, 0x7d, 0xbe, 0x73, 0xff, 0xce, 0x99, 0x19, 0x90, 0x99, 0xef, 0x50, 0x9f,
	0xfa, 0x4d, 0x6b, 0xda, 0x58, 0xfa, 0x2c, 0x64, 0xb8, 0x10, 0x23, 0xd5, 0x73, 0x9b, 0x2d, 0x16,
	0xcc, 0x6b, 0x46, 0x9f, 0x48, 0xab, 0x7e, 0x01, 0x67, 0x1d, 0x9f, 0x59, 0x8e, 0x6d, 0x05, 0x21,
	0xa1, 0xc1, 0x92, 0x79, 0x01, 0xc5, 0xcf, 0x20, 0x6f, 0x86, 0x56, 0xb8, 0x0a, 0x14, 0x54, 0x43,
	0xf5, 0xd3, 0x56, 0xa5, 0x11, 0xc7, 0x68, 0x44, 0x30, 0x89, 0xd5, 0xea, 0x8f, 0x08, 0x1e, 0x75,
	0x99, 0x37, 0x73, 0x6f, 0x57, 0xbe, 0x15, 0xba, 0xcc, 0xd3, 0xbc, 0x37, 0x74, 0xce, 0x96, 0x14,
	0x7f, 0x0e, 0x39, 0x3d, 0xa4, 0x0b, 0x1e, 0x41, 0xaa, 0x97, 0x5a, 0xb5, 0x75, 0x04, 0xf7, 0xd6,
	0xa3, 0xce, 0x96, 0x13, 0x37, 0x24, 0x91, 0x39, 0x56, 0xa0, 0xd0, 0xbd, 0xb3, 0x5c, 0x4f, 0xef,
	0x29, 0x47, 0x35, 0x54, 0x2f, 0x93, 0x44, 0xc4, 0x55, 0x38, 0x36, 0xe9, 0xf7, 0x2b, 0xea, 0xd9,
	0x54, 0x91, 0x6a, 0xa8, 0x9e, 0x25, 0xa9, 0xac, 0xfe, 0x8c, 0xe0, 0xf1, 0x81, 0xc0, 0xf8, 0x12,
	0xce, 0xf6, 0x40, 0xd1, 0x57, 0x99, 0xec, 0x2b, 0xf0, 0x0b, 0x00, 0x1e, 0xc8, 0x0a, 0x57, 0x3e,
	0x0d, 0x94, 0x23, 0x51, 0xfc, 0xd3, 0xb4, 0xf8, 0x2d, 0xfb, 0xd4, 0x8e, 0x6c, 0xb8, 0xa8, 0x7f,
	0x1d, 0x3d, 0x90, 0x0f, 0x7f, 0x04, 0xf9, 0x01, 0xb5, 0x1c, 0xea, 0x8b, 0xcc, 0xa5, 0xd6, 0x79,
	0x23, 0x9e, 0x82, 0xe8, 0x2e, 0x52, 0x91, 0xd8, 0x04, 0x7f, 0x09, 0xd9, 0xab, 0xfb, 0x25, 0x15,
	0x04, 0x9c, 0xb6, 0x2e, 0x1f, 0xce, 0xce, 0xc3, 0x6e, 0x23, 0xdc, 0x87, 0x08, 0x4f, 0xac, 0x42,
	0x79, 0x68, 0x05, 0xe1, 0x57, 0xcc, 0x71, 0x67, 0x2e, 0x75, 0x62, 0xbe, 0xb6, 0x30, 0xdc, 0x00,
	0x1c, 0xfd, 0xdb, 0xc2, 0xfb, 0x15, 0x9b, 0xbb, 0xf6, 0xbd, 0x92, 0xad, 0xa1, 0x7a, 0x91, 0x3c,
	0xa0, 0xc1, 0x32, 0x48, 0x2f, 0xe9, 0xbd, 0x92, 0x13, 0x06, 0xfc, 0x17, 0xbf, 0x03, 0xb9, 0x6b,
	0x6b, 0xbe, 0xa2, 0x4a, 0x5e, 0xb0, 0x19, 0x09, 0x6a, 0x77, 0xa7, 0x7f, 0x51, 0x10, 0x40, 0x3e,
	0x0a, 0x23, 0x67, 0x70, 0x11, 0x72, 0xa2, 0x6b, 0x19, 0xe1, 0x12, 0x14, 0x46, 0x51, 0x73, 0xf2,
	0x11, 0xb7, 0xe9, 0x5b, 0x53, 0xdf, 0xb5, 0x65, 0x49, 0xfd, 0x0e, 0x2e, 0x1e, 0xe6, 0x1a, 0xd7,
	0xa1, 0x12, 0x24, 0xc2, 0x06, 0xa5, 0x65, 0xb2, 0x0b, 0xe3, 0x27, 0x50, 0x4c, 0xa1, 0x78, 0x99,
	0xd6, 0x80, 0xfa, 0x6d, 0x52, 0x11, 0x1e, 0x42, 0x25, 0x0d, 0x1f, 0xb3, 0x10, 0x0d, 0x69, 0x7b,
	0x69, 0x37, 0xf4, 0xc9, 0x96, 0x0f, 0x32, 0x64, 0xd7, 0xb5, 0x93, 0x8f, 0x86, 0xc7, 0x8f, 0xc6,
	0xe3, 0x03, 0x6e, 0x7c, 0xc9, 0xaf, 0xa9, 0x1f, 0xb8, 0xcc, 0x13, 0x99, 0x72, 0x24, 0x11, 0xf1,
	0xf3, 0xa4, 0x2a, 0x51, 0x70, 0xa9, 0xa5, 0x1c, 0x2a, 0x81, 0x24, 0xd5, 0xbf, 0x07, 0xa0, 0x3b,
	0xd4, 0x0b, 0xdd, 0xd0, 0xa5, 0x81, 0x22, 0xd5, 0xa4, 0x7a, 0x99, 0x6c, 0x20, 0xea, 0x9f, 0x68,
	0xaf, 0x3d, 0xfc, 0x04, 0x8e, 0xa3, 0xd3, 0xd2, 0x89, 0x5a, 0xcd, 0x0d, 0x32, 0x24, 0x45, 0xf0,
	0x67, 0x90, 0xed, 0xfb, 0x6c, 0x11, 0x57, 0xf0, 0xf4, 0x50, 0x05, 0x0d, 0x63, 0xb4, 0x0a, 0x47,
	0xb3, 0x41, 0x86, 0x08, 0xf3, 0xea, 0x10, 0xf2, 0x11, 0x82, 0xcb, 0x80, 0x8c, 0xb8, 0x31, 0x64,
	0xe0, 0x4f, 0xe1, 0x58, 0x38, 0xb8, 0xe9, 0x79, 0x3a, 0xdc, 0x54, 0x6a, 0x99, 0xd2, 0xf8, 0x0f,
	0xe2, 0xc7, 0x9e, 0xbe, 0xd6, 0xbd, 0x19, 0xc3, 0x1f, 0x43, 0xce, 0x0c, 0x2d, 0x3f, 0x8c, 0xaf,
	0xa5, 0x77, 0xd7, 0x71, 0x62, 0x8b, 0x86, 0x50, 0x8b, 0x83, 0x10, 0x59, 0xf2, 0x75, 0x31, 0x97,
	0xd4, 0x16, 0x2b, 0x6f, 0xac, 0x16, 0x53, 0xea, 0x8b, 0xbe, 0xb2, 0x64, 0x17, 0xe6, 0x44, 0xde,
	0xb8, 0x9e, 0xc3, 0xde, 0x9a, 0xee, 0x0f, 0xc9, 0x0d, 0xb3, 0x81, 0x6c, 0xde, 0x4c, 0xd9, 0xad,
	0x9b, 0x49, 0x6d, 0x41, 0x31, 0xcd, 0xcb, 0xb7, 0xd8, 0xd0, 0x6e, 0x34, 0xf3, 0x4a, 0xce, 0xf0,
	0xff, 0xd1, 0xb0, 0xc7, 0xff, 0x11, 0x3e, 0x81, 0xa2, 0xf9, 0x4a, 0xeb, 0xea, 0x7d, 0x5d, 0xeb,
	0xc9, 0x47, 0xea, 0x07, 0x50, 0x69, 0xdb, 0xaf, 0x3d, 0xf6, 0x76, 0x4e, 0x9d, 0x5b, 0xba, 0xa0,
	0x5e, 0x88, 0x2f, 0x20, 0x1f, 0x57, 0x88, 0x44, 0xf2, 0x58, 0x52, 0x7f, 0x42, 0x70, 0xd2, 0xa3,
	0x73, 0xf7, 0x0d, 0xf5, 0xc7, 0x4b, 0xc7, 0x0a, 0x29, 0xee, 0xed, 0x39, 0xc7, 0x1b, 0xbb, 0x66,
	0x76, 0x47, 0xcf, 0x37, 0x75, 0x37, 0xdf, 0x33, 0xc8, 0x72, 0xde, 0xe2, 0x39, 0x9f, 0xed, 0x91,
	0xc9, 0x27, 0xcb, 0xff, 0xd3, 0x59, 0xb8, 0x50, 0x89, 0xeb, 0xd8, 0x78, 0x29, 0x72, 0x9a, 0xef,
	0x33, 0xff, 0xc0, 0x43, 0x31, 0xc8, 0x90, 0x48, 0x8f, 0xdf, 0x87, 0x5c, 0x67, 0xce, 0xec, 0x24,
	0xdb, 0x49, 0x72, 0xff, 0x09, 0x90, 0x9b, 0x89, 0x9f, 0x24, 0xd5, 0x87, 0xbf, 0xa0, 0xe4, 0x09,
	0xe2, 0x77, 0xc4, 0xd8, 0x78, 0x69, 0x8c, 0x6e, 0x0c, 0x39, 0x83, 0xcb, 0x50, 0x30, 0xc7, 0xdd,
	0xae, 0x66, 0x9a, 0xf2, 0xdf, 0x08, 0xcb, 0x50, 0xea, 0xb4, 0x7b, 0x13, 0xa2, 0x7d, 0x3d, 0xe6,
	0x24, 0xff, 0x2a, 0xe1, 0x53, 0x28, 0xf6, 0x47, 0xa4, 0xa3, 0xf7, 0x7a, 0x9a, 0x21, 0xff, 0x26,
	0x64, 0x63, 0x74, 0x35, 0xe9, 0x8f, 0xc6, 0x46, 0x4f, 0xfe, 0x5d, 0xc2, 0x55, 0x78, 0xa4, 0x1b,
	0x57, 0x1a, 0x31, 0xda, 0xc3, 0x89, 0xa9, 0x91, 0x6b, 0x8d, 0x4c, 0x34, 0x42, 0x46, 0x44, 0xfe,
	0x57, 0xc2, 0x0a, 0x9c, 0x73, 0x48, 0xef, 0x6a, 0x93, 0xb1, 0xd1, 0xbe, 0x6e, 0xeb, 0xc3, 0x76,
	0x67, 0xa8, 0xc9, 0xff, 0x49, 0xad, 0x3f, 0x10, 0x54, 0xda, 0x21, 0x5b, 0xb8, 0x76, 0xfa, 0x56,
	0xe2, 0x17, 0x50, 0x5c, 0x0b, 0x72, 0xd2, 0x4e, 0x72, 0xc4, 0xab, 0xd5, 0x94, 0x89, 0xbd, 0xe7,
	0x55, 0xcd, 0xd4, 0xd1, 0x73, 0x84, 0xdb, 0x50, 0x88, 0xd9, 0xc4, 0x17, 0xa9, 0xf1, 0xd6, 0x9c,
	0xab, 0xca, 0x2e, 0xbe, 0x1d, 0xa2, 0xd3, 0xf8, 0xe6, 0xf2, 0xd6, 0x0d, 0xef, 0x56, 0x53, 0x9e,
	0xbe, 0x79, 0x77, 0xbf, 0xa4, 0xbe, 0x18, 0xaf, 0xdf, 0x9c, 0x89, 0x8b, 0xb4, 0x29, 0xde, 0xf8,
	0xa0, 0x19, 0x47, 0x99, 0xe6, 0x85, 0xfc, 0xc9, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xb3,
	0x73, 0x13, 0x25, 0x08, 0x00, 0x00,
}
