// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v2/packet.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// PacketStatus specifies the status of a RecvPacketResult.
type PacketStatus int32

const (
	// PACKET_STATUS_UNSPECIFIED indicates an unknown packet status.
	PacketStatus_NONE PacketStatus = 0
	// PACKET_STATUS_SUCCESS indicates a successful packet receipt.
	PacketStatus_Success PacketStatus = 1
	// PACKET_STATUS_FAILURE indicates a failed packet receipt.
	PacketStatus_Failure PacketStatus = 2
	// PACKET_STATUS_ASYNC indicates an async packet receipt.
	PacketStatus_Async PacketStatus = 3
)

var PacketStatus_name = map[int32]string{
	0: "PACKET_STATUS_UNSPECIFIED",
	1: "PACKET_STATUS_SUCCESS",
	2: "PACKET_STATUS_FAILURE",
	3: "PACKET_STATUS_ASYNC",
}

var PacketStatus_value = map[string]int32{
	"PACKET_STATUS_UNSPECIFIED": 0,
	"PACKET_STATUS_SUCCESS":     1,
	"PACKET_STATUS_FAILURE":     2,
	"PACKET_STATUS_ASYNC":       3,
}

func (x PacketStatus) String() string {
	return proto.EnumName(PacketStatus_name, int32(x))
}

func (PacketStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{0}
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	// number corresponds to the order of sends and receives, where a Packet
	// with an earlier sequence number must be sent and received before a Packet
	// with a later sequence number.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// identifies the sending client on the sending chain.
	SourceClient string `protobuf:"bytes,2,opt,name=source_client,json=sourceClient,proto3" json:"source_client,omitempty"`
	// identifies the receiving client on the receiving chain.
	DestinationClient string `protobuf:"bytes,3,opt,name=destination_client,json=destinationClient,proto3" json:"destination_client,omitempty"`
	// timeout timestamp in seconds after which the packet times out.
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty"`
	// a list of payloads, each one for a specific application.
	Payloads []Payload `protobuf:"bytes,5,rep,name=payloads,proto3" json:"payloads"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Packet) GetSourceClient() string {
	if m != nil {
		return m.SourceClient
	}
	return ""
}

func (m *Packet) GetDestinationClient() string {
	if m != nil {
		return m.DestinationClient
	}
	return ""
}

func (m *Packet) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *Packet) GetPayloads() []Payload {
	if m != nil {
		return m.Payloads
	}
	return nil
}

// Payload contains the source and destination ports and payload for the application (version, encoding, raw bytes)
type Payload struct {
	// specifies the source port of the packet.
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// specifies the destination port of the packet.
	DestinationPort string `protobuf:"bytes,2,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// version of the specified application.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// the encoding used for the provided value.
	Encoding string `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding,omitempty"`
	// the raw bytes for the payload.
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{1}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return m.Size()
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *Payload) GetDestinationPort() string {
	if m != nil {
		return m.DestinationPort
	}
	return ""
}

func (m *Payload) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Payload) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *Payload) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Acknowledgement contains a list of all ack results associated with a single packet.
// In the case of a successful receive, the acknowledgement will contain an app acknowledgement
// for each application that received a payload in the same order that the payloads were sent
// in the packet.
// If the receive is not successful, the acknowledgement will contain a single app acknowledgment
// which will be a constant error acknowledgment as defined by the IBC v2 protocol.
type Acknowledgement struct {
	AppAcknowledgements [][]byte `protobuf:"bytes,1,rep,name=app_acknowledgements,json=appAcknowledgements,proto3" json:"app_acknowledgements,omitempty"`
}

func (m *Acknowledgement) Reset()         { *m = Acknowledgement{} }
func (m *Acknowledgement) String() string { return proto.CompactTextString(m) }
func (*Acknowledgement) ProtoMessage()    {}
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{2}
}
func (m *Acknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Acknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Acknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Acknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acknowledgement.Merge(m, src)
}
func (m *Acknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *Acknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_Acknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_Acknowledgement proto.InternalMessageInfo

func (m *Acknowledgement) GetAppAcknowledgements() [][]byte {
	if m != nil {
		return m.AppAcknowledgements
	}
	return nil
}

// RecvPacketResult speecifies the status of a packet as well as the acknowledgement bytes.
type RecvPacketResult struct {
	// status of the packet
	Status PacketStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ibc.core.channel.v2.PacketStatus" json:"status,omitempty"`
	// acknowledgement of the packet
	Acknowledgement []byte `protobuf:"bytes,2,opt,name=acknowledgement,proto3" json:"acknowledgement,omitempty"`
}

func (m *RecvPacketResult) Reset()         { *m = RecvPacketResult{} }
func (m *RecvPacketResult) String() string { return proto.CompactTextString(m) }
func (*RecvPacketResult) ProtoMessage()    {}
func (*RecvPacketResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{3}
}
func (m *RecvPacketResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecvPacketResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecvPacketResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecvPacketResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvPacketResult.Merge(m, src)
}
func (m *RecvPacketResult) XXX_Size() int {
	return m.Size()
}
func (m *RecvPacketResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvPacketResult.DiscardUnknown(m)
}

var xxx_messageInfo_RecvPacketResult proto.InternalMessageInfo

func (m *RecvPacketResult) GetStatus() PacketStatus {
	if m != nil {
		return m.Status
	}
	return PacketStatus_NONE
}

func (m *RecvPacketResult) GetAcknowledgement() []byte {
	if m != nil {
		return m.Acknowledgement
	}
	return nil
}

func init() {
	proto.RegisterEnum("ibc.core.channel.v2.PacketStatus", PacketStatus_name, PacketStatus_value)
	proto.RegisterType((*Packet)(nil), "ibc.core.channel.v2.Packet")
	proto.RegisterType((*Payload)(nil), "ibc.core.channel.v2.Payload")
	proto.RegisterType((*Acknowledgement)(nil), "ibc.core.channel.v2.Acknowledgement")
	proto.RegisterType((*RecvPacketResult)(nil), "ibc.core.channel.v2.RecvPacketResult")
}

func init() { proto.RegisterFile("ibc/core/channel/v2/packet.proto", fileDescriptor_2f814aba9ca97169) }

var fileDescriptor_2f814aba9ca97169 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x4d, 0xd2, 0x36, 0xdb, 0x40, 0xdd, 0x6d, 0x91, 0x4c, 0x84, 0x52, 0x13, 0x24,
	0x08, 0xa0, 0xda, 0x10, 0xb8, 0x54, 0x42, 0x48, 0xa9, 0xeb, 0x4a, 0x15, 0x28, 0x44, 0xeb, 0x04,
	0x09, 0x2e, 0xd1, 0x66, 0xb3, 0x72, 0xad, 0xda, 0x5e, 0xe3, 0x5d, 0xbb, 0xea, 0x2b, 0xf4, 0xc4,
	0x0b, 0xf4, 0xc0, 0x99, 0x17, 0xe9, 0xb1, 0x47, 0x4e, 0x08, 0xb5, 0xe2, 0x3d, 0x90, 0xd7, 0x6e,
	0x95, 0x16, 0x38, 0xd9, 0xf3, 0xcf, 0xf7, 0x7b, 0xf4, 0x8f, 0x35, 0xd0, 0xf0, 0xa7, 0xd4, 0xa2,
	0x3c, 0x61, 0x16, 0x3d, 0x20, 0x51, 0xc4, 0x02, 0x2b, 0xeb, 0x59, 0x31, 0xa1, 0x87, 0x4c, 0x9a,
	0x71, 0xc2, 0x25, 0x47, 0xeb, 0xfe, 0x94, 0x9a, 0x39, 0x61, 0x96, 0x84, 0x99, 0xf5, 0x5a, 0x1b,
	0x1e, 0xf7, 0xb8, 0xea, 0x5b, 0xf9, 0x5b, 0x81, 0x76, 0x7e, 0x03, 0xb8, 0x38, 0x54, 0x5e, 0xd4,
	0x82, 0xcb, 0x82, 0x7d, 0x49, 0x59, 0x44, 0x99, 0x0e, 0x0c, 0xd0, 0xad, 0xe1, 0xeb, 0x1a, 0x3d,
	0x82, 0x77, 0x04, 0x4f, 0x13, 0xca, 0x26, 0x34, 0xf0, 0x59, 0x24, 0xf5, 0x05, 0x03, 0x74, 0x1b,
	0xb8, 0x59, 0x88, 0xb6, 0xd2, 0xd0, 0x16, 0x44, 0x33, 0x26, 0xa4, 0x1f, 0x11, 0xe9, 0xf3, 0xe8,
	0x8a, 0xac, 0x2a, 0x72, 0x6d, 0xae, 0x53, 0xe2, 0xcf, 0xe1, 0x9a, 0xf4, 0x43, 0xc6, 0x53, 0x39,
	0xc9, 0x9f, 0x42, 0x92, 0x30, 0xd6, 0x6b, 0x6a, 0xb0, 0x56, 0x36, 0x46, 0x57, 0x3a, 0x7a, 0x0b,
	0x97, 0x63, 0x72, 0x1c, 0x70, 0x32, 0x13, 0x7a, 0xdd, 0xa8, 0x76, 0x57, 0x7a, 0x0f, 0xcc, 0x7f,
	0xa4, 0x34, 0x87, 0x05, 0xb4, 0x53, 0x3b, 0xfb, 0xb9, 0x59, 0xc1, 0xd7, 0x9e, 0xce, 0x37, 0x00,
	0x97, 0xca, 0x1e, 0xda, 0x84, 0x2b, 0x65, 0x98, 0x98, 0x27, 0x52, 0x65, 0x6d, 0x60, 0x58, 0x48,
	0x43, 0x9e, 0x48, 0xf4, 0x14, 0x6a, 0xf3, 0x41, 0x14, 0x55, 0x04, 0x5e, 0x9d, 0xd3, 0x15, 0xaa,
	0xc3, 0xa5, 0x8c, 0x25, 0xc2, 0xe7, 0x51, 0x19, 0xf4, 0xaa, 0xcc, 0xd7, 0xc9, 0x22, 0xca, 0x67,
	0x7e, 0xe4, 0xa9, 0x54, 0x0d, 0x7c, 0x5d, 0xa3, 0x0d, 0x58, 0xcf, 0x48, 0x90, 0x32, 0xbd, 0x6e,
	0x80, 0x6e, 0x13, 0x17, 0x45, 0x67, 0x17, 0xae, 0xf6, 0xe9, 0x61, 0xc4, 0x8f, 0x02, 0x36, 0xf3,
	0x58, 0x98, 0xef, 0xe8, 0x25, 0xdc, 0x20, 0x71, 0x3c, 0x21, 0x37, 0x65, 0xa1, 0x03, 0xa3, 0xda,
	0x6d, 0xe2, 0x75, 0x12, 0xc7, 0xb7, 0x1c, 0xa2, 0x73, 0x04, 0x35, 0xcc, 0x68, 0x56, 0xfc, 0x54,
	0xcc, 0x44, 0x1a, 0x48, 0xb4, 0x0d, 0x17, 0x85, 0x24, 0x32, 0x15, 0x2a, 0xec, 0xdd, 0xde, 0xc3,
	0xff, 0xec, 0x2e, 0xb7, 0xb8, 0x0a, 0xc4, 0xa5, 0x01, 0x75, 0xe1, 0xea, 0xad, 0xe9, 0x6a, 0x15,
	0x4d, 0x7c, 0x5b, 0x7e, 0xf6, 0x1d, 0xc0, 0xe6, 0xfc, 0x27, 0xd0, 0x13, 0x78, 0x7f, 0xd8, 0xb7,
	0xdf, 0x39, 0xa3, 0x89, 0x3b, 0xea, 0x8f, 0xc6, 0xee, 0x64, 0x3c, 0x70, 0x87, 0x8e, 0xbd, 0xbf,
	0xb7, 0xef, 0xec, 0x6a, 0x95, 0xd6, 0xf2, 0xc9, 0xa9, 0x51, 0x1b, 0x7c, 0x18, 0x38, 0xe8, 0x31,
	0xbc, 0x77, 0x13, 0x74, 0xc7, 0xb6, 0xed, 0xb8, 0xae, 0x06, 0x5a, 0x2b, 0x27, 0xa7, 0xc6, 0x92,
	0x9b, 0x52, 0xca, 0x84, 0xf8, 0x9b, 0xdb, 0xeb, 0xef, 0xbf, 0x1f, 0x63, 0x47, 0x5b, 0x28, 0xb8,
	0x3d, 0xe2, 0x07, 0x69, 0xc2, 0x50, 0x07, 0xae, 0xdf, 0xe4, 0xfa, 0xee, 0xa7, 0x81, 0xad, 0x55,
	0x5b, 0x8d, 0x93, 0x53, 0xa3, 0xde, 0x17, 0xc7, 0x11, 0xdd, 0xf9, 0x78, 0x76, 0xd1, 0x06, 0xe7,
	0x17, 0x6d, 0xf0, 0xeb, 0xa2, 0x0d, 0xbe, 0x5e, 0xb6, 0x2b, 0xe7, 0x97, 0xed, 0xca, 0x8f, 0xcb,
	0x76, 0xe5, 0xf3, 0x1b, 0xcf, 0x97, 0x07, 0xe9, 0xd4, 0xa4, 0x3c, 0xb4, 0x28, 0x17, 0x21, 0x17,
	0x96, 0x3f, 0xa5, 0x5b, 0x1e, 0xb7, 0xb2, 0x6d, 0x2b, 0xe4, 0xb3, 0x34, 0x60, 0xa2, 0xb8, 0xbf,
	0x17, 0xaf, 0xb7, 0xe6, 0x4e, 0x50, 0x1e, 0xc7, 0x4c, 0x4c, 0x17, 0xd5, 0x5d, 0xbd, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0x60, 0x44, 0x2c, 0xa6, 0x03, 0x00, 0x00,
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payloads) > 0 {
		for iNdEx := len(m.Payloads) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Payloads[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.DestinationClient) > 0 {
		i -= len(m.DestinationClient)
		copy(dAtA[i:], m.DestinationClient)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DestinationClient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceClient) > 0 {
		i -= len(m.SourceClient)
		copy(dAtA[i:], m.SourceClient)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourceClient)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Encoding) > 0 {
		i -= len(m.Encoding)
		copy(dAtA[i:], m.Encoding)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Encoding)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestinationPort) > 0 {
		i -= len(m.DestinationPort)
		copy(dAtA[i:], m.DestinationPort)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DestinationPort)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Acknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Acknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Acknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppAcknowledgements) > 0 {
		for iNdEx := len(m.AppAcknowledgements) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AppAcknowledgements[iNdEx])
			copy(dAtA[i:], m.AppAcknowledgements[iNdEx])
			i = encodeVarintPacket(dAtA, i, uint64(len(m.AppAcknowledgements[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RecvPacketResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecvPacketResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecvPacketResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Acknowledgement) > 0 {
		i -= len(m.Acknowledgement)
		copy(dAtA[i:], m.Acknowledgement)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Acknowledgement)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovPacket(uint64(m.Sequence))
	}
	l = len(m.SourceClient)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DestinationClient)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovPacket(uint64(m.TimeoutTimestamp))
	}
	if len(m.Payloads) > 0 {
		for _, e := range m.Payloads {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func (m *Payload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DestinationPort)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Encoding)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *Acknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AppAcknowledgements) > 0 {
		for _, b := range m.AppAcknowledgements {
			l = len(b)
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func (m *RecvPacketResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovPacket(uint64(m.Status))
	}
	l = len(m.Acknowledgement)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceClient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceClient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationClient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationClient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payloads", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payloads = append(m.Payloads, Payload{})
			if err := m.Payloads[len(m.Payloads)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Encoding = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Acknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Acknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Acknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAcknowledgements", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAcknowledgements = append(m.AppAcknowledgements, make([]byte, postIndex-iNdEx))
			copy(m.AppAcknowledgements[len(m.AppAcknowledgements)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RecvPacketResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RecvPacketResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecvPacketResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= PacketStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Acknowledgement", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Acknowledgement = append(m.Acknowledgement[:0], dAtA[iNdEx:postIndex]...)
			if m.Acknowledgement == nil {
				m.Acknowledgement = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
