// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rollup.proto

package rollup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	structures "github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	valprotocol "github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeBuf struct {
	PrevHash             *common.HashBuf                   `protobuf:"bytes,1,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Deadline             *common.TimeTicksBuf              `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	DisputableNode       *valprotocol.DisputableNodeBuf    `protobuf:"bytes,3,opt,name=disputableNode,proto3" json:"disputableNode,omitempty"`
	LinkType             uint32                            `protobuf:"varint,4,opt,name=linkType,proto3" json:"linkType,omitempty"`
	VmProtoData          *valprotocol.VMProtoDataBuf       `protobuf:"bytes,5,opt,name=vmProtoData,proto3" json:"vmProtoData,omitempty"`
	MachineHash          *common.HashBuf                   `protobuf:"bytes,6,opt,name=machineHash,proto3" json:"machineHash,omitempty"`
	Assertion            *structures.ExecutionAssertionBuf `protobuf:"bytes,7,opt,name=assertion,proto3" json:"assertion,omitempty"`
	Depth                uint64                            `protobuf:"varint,8,opt,name=depth,proto3" json:"depth,omitempty"`
	NodeDataHash         *common.HashBuf                   `protobuf:"bytes,9,opt,name=nodeDataHash,proto3" json:"nodeDataHash,omitempty"`
	InnerHash            *common.HashBuf                   `protobuf:"bytes,10,opt,name=innerHash,proto3" json:"innerHash,omitempty"`
	Hash                 *common.HashBuf                   `protobuf:"bytes,11,opt,name=hash,proto3" json:"hash,omitempty"`
	AssertionTxHash      *common.HashBuf                   `protobuf:"bytes,12,opt,name=assertionTxHash,proto3" json:"assertionTxHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *NodeBuf) Reset()         { *m = NodeBuf{} }
func (m *NodeBuf) String() string { return proto.CompactTextString(m) }
func (*NodeBuf) ProtoMessage()    {}
func (*NodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{0}
}

func (m *NodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeBuf.Unmarshal(m, b)
}
func (m *NodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeBuf.Marshal(b, m, deterministic)
}
func (m *NodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeBuf.Merge(m, src)
}
func (m *NodeBuf) XXX_Size() int {
	return xxx_messageInfo_NodeBuf.Size(m)
}
func (m *NodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_NodeBuf proto.InternalMessageInfo

func (m *NodeBuf) GetPrevHash() *common.HashBuf {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *NodeBuf) GetDeadline() *common.TimeTicksBuf {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *NodeBuf) GetDisputableNode() *valprotocol.DisputableNodeBuf {
	if m != nil {
		return m.DisputableNode
	}
	return nil
}

func (m *NodeBuf) GetLinkType() uint32 {
	if m != nil {
		return m.LinkType
	}
	return 0
}

func (m *NodeBuf) GetVmProtoData() *valprotocol.VMProtoDataBuf {
	if m != nil {
		return m.VmProtoData
	}
	return nil
}

func (m *NodeBuf) GetMachineHash() *common.HashBuf {
	if m != nil {
		return m.MachineHash
	}
	return nil
}

func (m *NodeBuf) GetAssertion() *structures.ExecutionAssertionBuf {
	if m != nil {
		return m.Assertion
	}
	return nil
}

func (m *NodeBuf) GetDepth() uint64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *NodeBuf) GetNodeDataHash() *common.HashBuf {
	if m != nil {
		return m.NodeDataHash
	}
	return nil
}

func (m *NodeBuf) GetInnerHash() *common.HashBuf {
	if m != nil {
		return m.InnerHash
	}
	return nil
}

func (m *NodeBuf) GetHash() *common.HashBuf {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *NodeBuf) GetAssertionTxHash() *common.HashBuf {
	if m != nil {
		return m.AssertionTxHash
	}
	return nil
}

type NodeGraphBuf struct {
	Nodes                []*NodeBuf                  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	OldestNodeHash       *common.HashBuf             `protobuf:"bytes,2,opt,name=oldestNodeHash,proto3" json:"oldestNodeHash,omitempty"`
	LatestConfirmedHash  *common.HashBuf             `protobuf:"bytes,3,opt,name=latestConfirmedHash,proto3" json:"latestConfirmedHash,omitempty"`
	LeafHashes           []*common.HashBuf           `protobuf:"bytes,4,rep,name=leafHashes,proto3" json:"leafHashes,omitempty"`
	Params               *valprotocol.ChainParamsBuf `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NodeGraphBuf) Reset()         { *m = NodeGraphBuf{} }
func (m *NodeGraphBuf) String() string { return proto.CompactTextString(m) }
func (*NodeGraphBuf) ProtoMessage()    {}
func (*NodeGraphBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{1}
}

func (m *NodeGraphBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGraphBuf.Unmarshal(m, b)
}
func (m *NodeGraphBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGraphBuf.Marshal(b, m, deterministic)
}
func (m *NodeGraphBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGraphBuf.Merge(m, src)
}
func (m *NodeGraphBuf) XXX_Size() int {
	return xxx_messageInfo_NodeGraphBuf.Size(m)
}
func (m *NodeGraphBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGraphBuf.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGraphBuf proto.InternalMessageInfo

func (m *NodeGraphBuf) GetNodes() []*NodeBuf {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *NodeGraphBuf) GetOldestNodeHash() *common.HashBuf {
	if m != nil {
		return m.OldestNodeHash
	}
	return nil
}

func (m *NodeGraphBuf) GetLatestConfirmedHash() *common.HashBuf {
	if m != nil {
		return m.LatestConfirmedHash
	}
	return nil
}

func (m *NodeGraphBuf) GetLeafHashes() []*common.HashBuf {
	if m != nil {
		return m.LeafHashes
	}
	return nil
}

func (m *NodeGraphBuf) GetParams() *valprotocol.ChainParamsBuf {
	if m != nil {
		return m.Params
	}
	return nil
}

type StakedNodeGraphBuf struct {
	NodeGraph            *NodeGraphBuf   `protobuf:"bytes,1,opt,name=nodeGraph,proto3" json:"nodeGraph,omitempty"`
	Stakers              []*StakerBuf    `protobuf:"bytes,2,rep,name=stakers,proto3" json:"stakers,omitempty"`
	Challenges           []*ChallengeBuf `protobuf:"bytes,3,rep,name=challenges,proto3" json:"challenges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StakedNodeGraphBuf) Reset()         { *m = StakedNodeGraphBuf{} }
func (m *StakedNodeGraphBuf) String() string { return proto.CompactTextString(m) }
func (*StakedNodeGraphBuf) ProtoMessage()    {}
func (*StakedNodeGraphBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{2}
}

func (m *StakedNodeGraphBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakedNodeGraphBuf.Unmarshal(m, b)
}
func (m *StakedNodeGraphBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakedNodeGraphBuf.Marshal(b, m, deterministic)
}
func (m *StakedNodeGraphBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakedNodeGraphBuf.Merge(m, src)
}
func (m *StakedNodeGraphBuf) XXX_Size() int {
	return xxx_messageInfo_StakedNodeGraphBuf.Size(m)
}
func (m *StakedNodeGraphBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_StakedNodeGraphBuf.DiscardUnknown(m)
}

var xxx_messageInfo_StakedNodeGraphBuf proto.InternalMessageInfo

func (m *StakedNodeGraphBuf) GetNodeGraph() *NodeGraphBuf {
	if m != nil {
		return m.NodeGraph
	}
	return nil
}

func (m *StakedNodeGraphBuf) GetStakers() []*StakerBuf {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func (m *StakedNodeGraphBuf) GetChallenges() []*ChallengeBuf {
	if m != nil {
		return m.Challenges
	}
	return nil
}

type ChainObserverBuf struct {
	StakedNodeGraph      *StakedNodeGraphBuf  `protobuf:"bytes,1,opt,name=stakedNodeGraph,proto3" json:"stakedNodeGraph,omitempty"`
	ContractAddress      *common.AddressBuf   `protobuf:"bytes,2,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Inbox                *structures.InboxBuf `protobuf:"bytes,3,opt,name=inbox,proto3" json:"inbox,omitempty"`
	KnownValidNode       *common.HashBuf      `protobuf:"bytes,4,opt,name=knownValidNode,proto3" json:"knownValidNode,omitempty"`
	CalculatedValidNode  *common.HashBuf      `protobuf:"bytes,5,opt,name=calculatedValidNode,proto3" json:"calculatedValidNode,omitempty"`
	LatestBlockId        *common.BlockIdBuf   `protobuf:"bytes,6,opt,name=latestBlockId,proto3" json:"latestBlockId,omitempty"`
	IsOpinionated        bool                 `protobuf:"varint,7,opt,name=isOpinionated,proto3" json:"isOpinionated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChainObserverBuf) Reset()         { *m = ChainObserverBuf{} }
func (m *ChainObserverBuf) String() string { return proto.CompactTextString(m) }
func (*ChainObserverBuf) ProtoMessage()    {}
func (*ChainObserverBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{3}
}

func (m *ChainObserverBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainObserverBuf.Unmarshal(m, b)
}
func (m *ChainObserverBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainObserverBuf.Marshal(b, m, deterministic)
}
func (m *ChainObserverBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainObserverBuf.Merge(m, src)
}
func (m *ChainObserverBuf) XXX_Size() int {
	return xxx_messageInfo_ChainObserverBuf.Size(m)
}
func (m *ChainObserverBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainObserverBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainObserverBuf proto.InternalMessageInfo

func (m *ChainObserverBuf) GetStakedNodeGraph() *StakedNodeGraphBuf {
	if m != nil {
		return m.StakedNodeGraph
	}
	return nil
}

func (m *ChainObserverBuf) GetContractAddress() *common.AddressBuf {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

func (m *ChainObserverBuf) GetInbox() *structures.InboxBuf {
	if m != nil {
		return m.Inbox
	}
	return nil
}

func (m *ChainObserverBuf) GetKnownValidNode() *common.HashBuf {
	if m != nil {
		return m.KnownValidNode
	}
	return nil
}

func (m *ChainObserverBuf) GetCalculatedValidNode() *common.HashBuf {
	if m != nil {
		return m.CalculatedValidNode
	}
	return nil
}

func (m *ChainObserverBuf) GetLatestBlockId() *common.BlockIdBuf {
	if m != nil {
		return m.LatestBlockId
	}
	return nil
}

func (m *ChainObserverBuf) GetIsOpinionated() bool {
	if m != nil {
		return m.IsOpinionated
	}
	return false
}

type StakerBuf struct {
	Address              *common.AddressBuf   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location             *common.HashBuf      `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CreationTime         *common.TimeTicksBuf `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	ChallengeAddr        *common.AddressBuf   `protobuf:"bytes,4,opt,name=challengeAddr,proto3" json:"challengeAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StakerBuf) Reset()         { *m = StakerBuf{} }
func (m *StakerBuf) String() string { return proto.CompactTextString(m) }
func (*StakerBuf) ProtoMessage()    {}
func (*StakerBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{4}
}

func (m *StakerBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakerBuf.Unmarshal(m, b)
}
func (m *StakerBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakerBuf.Marshal(b, m, deterministic)
}
func (m *StakerBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerBuf.Merge(m, src)
}
func (m *StakerBuf) XXX_Size() int {
	return xxx_messageInfo_StakerBuf.Size(m)
}
func (m *StakerBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerBuf.DiscardUnknown(m)
}

var xxx_messageInfo_StakerBuf proto.InternalMessageInfo

func (m *StakerBuf) GetAddress() *common.AddressBuf {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *StakerBuf) GetLocation() *common.HashBuf {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *StakerBuf) GetCreationTime() *common.TimeTicksBuf {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *StakerBuf) GetChallengeAddr() *common.AddressBuf {
	if m != nil {
		return m.ChallengeAddr
	}
	return nil
}

type ChallengeBuf struct {
	BlockId              *common.BlockIdBuf `protobuf:"bytes,1,opt,name=blockId,proto3" json:"blockId,omitempty"`
	LogIndex             uint64             `protobuf:"varint,2,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Asserter             *common.AddressBuf `protobuf:"bytes,3,opt,name=asserter,proto3" json:"asserter,omitempty"`
	Challenger           *common.AddressBuf `protobuf:"bytes,4,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Contract             *common.AddressBuf `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
	ConflictNodeHash     *common.HashBuf    `protobuf:"bytes,6,opt,name=conflictNodeHash,proto3" json:"conflictNodeHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ChallengeBuf) Reset()         { *m = ChallengeBuf{} }
func (m *ChallengeBuf) String() string { return proto.CompactTextString(m) }
func (*ChallengeBuf) ProtoMessage()    {}
func (*ChallengeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{5}
}

func (m *ChallengeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeBuf.Unmarshal(m, b)
}
func (m *ChallengeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeBuf.Marshal(b, m, deterministic)
}
func (m *ChallengeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeBuf.Merge(m, src)
}
func (m *ChallengeBuf) XXX_Size() int {
	return xxx_messageInfo_ChallengeBuf.Size(m)
}
func (m *ChallengeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeBuf proto.InternalMessageInfo

func (m *ChallengeBuf) GetBlockId() *common.BlockIdBuf {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *ChallengeBuf) GetLogIndex() uint64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *ChallengeBuf) GetAsserter() *common.AddressBuf {
	if m != nil {
		return m.Asserter
	}
	return nil
}

func (m *ChallengeBuf) GetChallenger() *common.AddressBuf {
	if m != nil {
		return m.Challenger
	}
	return nil
}

func (m *ChallengeBuf) GetContract() *common.AddressBuf {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ChallengeBuf) GetConflictNodeHash() *common.HashBuf {
	if m != nil {
		return m.ConflictNodeHash
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeBuf)(nil), "rollup.NodeBuf")
	proto.RegisterType((*NodeGraphBuf)(nil), "rollup.NodeGraphBuf")
	proto.RegisterType((*StakedNodeGraphBuf)(nil), "rollup.StakedNodeGraphBuf")
	proto.RegisterType((*ChainObserverBuf)(nil), "rollup.ChainObserverBuf")
	proto.RegisterType((*StakerBuf)(nil), "rollup.StakerBuf")
	proto.RegisterType((*ChallengeBuf)(nil), "rollup.ChallengeBuf")
}

func init() { proto.RegisterFile("rollup.proto", fileDescriptor_037f188b50610c79) }

var fileDescriptor_037f188b50610c79 = []byte{
	// 858 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x5b, 0x6f, 0x23, 0x35,
	0x14, 0x80, 0x35, 0xb9, 0xb4, 0xc9, 0x69, 0xba, 0x5d, 0xbc, 0x7d, 0x18, 0x75, 0x05, 0x2a, 0x01,
	0xa4, 0x8a, 0x65, 0x13, 0x68, 0x91, 0x58, 0xc4, 0xb5, 0x17, 0x2e, 0x7d, 0x80, 0x5d, 0x85, 0x6a,
	0x1f, 0x78, 0x73, 0x3c, 0x4e, 0x63, 0xc5, 0x63, 0x8f, 0x6c, 0x4f, 0x28, 0xbf, 0x84, 0x37, 0xde,
	0x11, 0xbf, 0x08, 0x7e, 0x0d, 0xf2, 0x19, 0xcf, 0x64, 0x26, 0x4c, 0xf6, 0x69, 0x6c, 0x9f, 0xef,
	0x5c, 0x7d, 0x7c, 0x06, 0x46, 0x46, 0x4b, 0x99, 0x67, 0x93, 0xcc, 0x68, 0xa7, 0xc9, 0x5e, 0xb1,
	0x3b, 0x79, 0xc2, 0x74, 0x9a, 0x6a, 0x35, 0x2d, 0x3e, 0x85, 0xf0, 0xe4, 0xa9, 0x75, 0x26, 0x67,
	0x2e, 0x37, 0xdc, 0x4e, 0x37, 0xcb, 0x20, 0x7c, 0x7b, 0x4d, 0x25, 0xae, 0x98, 0x96, 0xd3, 0xda,
	0xba, 0x10, 0x8f, 0xff, 0xe9, 0xc1, 0xfe, 0xcf, 0x3a, 0xe1, 0x57, 0xf9, 0x82, 0x3c, 0x83, 0x41,
	0x66, 0xf8, 0xfa, 0x47, 0x6a, 0x97, 0x71, 0x74, 0x1a, 0x9d, 0x1d, 0x9c, 0x1f, 0x4d, 0x82, 0x23,
	0x7f, 0x76, 0x95, 0x2f, 0x66, 0x15, 0x40, 0x3e, 0x86, 0x41, 0xc2, 0x69, 0x22, 0x85, 0xe2, 0x71,
	0x07, 0xe1, 0xe3, 0x12, 0xbe, 0x13, 0x29, 0xbf, 0x13, 0x6c, 0x65, 0x51, 0xa3, 0xa4, 0xc8, 0xf7,
	0xf0, 0x28, 0x11, 0x36, 0xcb, 0x1d, 0x9d, 0x4b, 0xee, 0x7d, 0xc6, 0x5d, 0xd4, 0x7b, 0x67, 0x52,
	0x0f, 0xeb, 0xa6, 0x81, 0x78, 0x0b, 0x5b, 0x5a, 0xe4, 0x04, 0x06, 0x52, 0xa8, 0xd5, 0xdd, 0xef,
	0x19, 0x8f, 0x7b, 0xa7, 0xd1, 0xd9, 0xe1, 0xac, 0xda, 0x93, 0xaf, 0xe0, 0x60, 0x9d, 0xbe, 0xf2,
	0xb6, 0x6e, 0xa8, 0xa3, 0x71, 0x1f, 0x1d, 0x3c, 0x6d, 0x38, 0x78, 0xfd, 0x53, 0x25, 0xf7, 0xd6,
	0xeb, 0x3c, 0xf9, 0x04, 0x0e, 0x52, 0xca, 0x96, 0x42, 0x71, 0x2c, 0xc2, 0x5e, 0x7b, 0x11, 0xea,
	0x0c, 0xf9, 0x06, 0x86, 0xd4, 0x5a, 0x6e, 0x9c, 0xd0, 0x2a, 0xde, 0x47, 0x85, 0x77, 0x27, 0xb5,
	0x5b, 0xf8, 0xee, 0x81, 0xb3, 0xdc, 0x0b, 0x2f, 0x4b, 0xca, 0x9b, 0xd8, 0xe8, 0x90, 0x63, 0xe8,
	0x27, 0x3c, 0x73, 0xcb, 0x78, 0x70, 0x1a, 0x9d, 0xf5, 0x66, 0xc5, 0x86, 0x5c, 0xc0, 0x48, 0xe9,
	0x84, 0xfb, 0xa8, 0x30, 0x94, 0x61, 0x7b, 0x28, 0x0d, 0x88, 0x3c, 0x87, 0xa1, 0x50, 0x8a, 0x1b,
	0xd4, 0x80, 0x76, 0x8d, 0x0d, 0x41, 0xde, 0x83, 0xde, 0xd2, 0x93, 0x07, 0xed, 0x24, 0x0a, 0xc9,
	0xe7, 0x70, 0x54, 0xc5, 0x7a, 0xf7, 0x80, 0x96, 0x47, 0xed, 0xfc, 0x36, 0x37, 0xfe, 0xa3, 0x03,
	0x23, 0x7f, 0x63, 0x3f, 0x18, 0x9a, 0x79, 0x82, 0x7c, 0x00, 0x7d, 0x1f, 0xaf, 0x8d, 0xa3, 0xd3,
	0x2e, 0x5a, 0x08, 0x3d, 0x5e, 0xde, 0x74, 0x21, 0x25, 0x9f, 0xc1, 0x23, 0x2d, 0x13, 0x6e, 0x9d,
	0x3f, 0x47, 0x8f, 0x9d, 0x76, 0x8f, 0x5b, 0x18, 0xb9, 0x84, 0x27, 0x92, 0x3a, 0x6e, 0xdd, 0xb5,
	0x56, 0x0b, 0x61, 0x52, 0x9e, 0xa0, 0x76, 0xb7, 0x5d, 0xbb, 0x8d, 0x25, 0x53, 0x00, 0xc9, 0xe9,
	0xc2, 0xaf, 0xb9, 0x8d, 0x7b, 0x21, 0xce, 0x2d, 0xcd, 0x1a, 0x42, 0x2e, 0x60, 0x2f, 0xa3, 0x86,
	0xa6, 0xb6, 0xb5, 0xd9, 0xae, 0x97, 0x54, 0xa8, 0x57, 0x28, 0xf7, 0x8a, 0x01, 0x1d, 0xff, 0x1d,
	0x01, 0xf9, 0xc5, 0xd1, 0x15, 0x4f, 0x1a, 0xf5, 0x39, 0x87, 0xa1, 0x2a, 0xf7, 0xe1, 0x05, 0x1e,
	0xd7, 0x6b, 0x54, 0x82, 0xb3, 0x0d, 0x46, 0x9e, 0xc1, 0xbe, 0xf5, 0x96, 0x8c, 0x8d, 0x3b, 0x18,
	0xed, 0x5b, 0xa5, 0x06, 0x3a, 0x30, 0x1e, 0x2f, 0x09, 0xf2, 0x29, 0x00, 0x5b, 0x52, 0x29, 0xb9,
	0xba, 0xe7, 0x36, 0xee, 0x22, 0x5f, 0x79, 0xb8, 0x2e, 0x25, 0x98, 0xe2, 0x86, 0x1b, 0xff, 0xd9,
	0x85, 0xc7, 0x98, 0xc8, 0xcb, 0xb9, 0xe5, 0x66, 0x8d, 0x36, 0xc9, 0x0d, 0x1c, 0xd9, 0x66, 0x06,
	0x21, 0xe2, 0x93, 0x86, 0xff, 0x46, 0x82, 0xb3, 0x6d, 0x15, 0xf2, 0x25, 0x1c, 0x31, 0xad, 0x9c,
	0xa1, 0xcc, 0x5d, 0x26, 0x89, 0xe1, 0xd6, 0x86, 0xbb, 0x26, 0x65, 0xcd, 0xc3, 0x31, 0x6a, 0x6f,
	0xa1, 0xe4, 0x43, 0xe8, 0x0b, 0x35, 0xd7, 0x0f, 0xe1, 0x86, 0x8f, 0xeb, 0xef, 0xee, 0xd6, 0x0b,
	0xb0, 0xa9, 0x10, 0xf1, 0x4d, 0xb5, 0x52, 0xfa, 0x37, 0xf5, 0x9a, 0x4a, 0x81, 0x01, 0xe0, 0xec,
	0x68, 0x6b, 0xaa, 0x26, 0xe6, 0x9b, 0x8a, 0x51, 0xc9, 0x72, 0xdf, 0x2d, 0xc9, 0x46, 0xbb, 0xbf,
	0xa3, 0xa9, 0x5a, 0x58, 0xf2, 0x02, 0x0e, 0x8b, 0x5e, 0xbb, 0x92, 0x9a, 0xad, 0x6e, 0x93, 0x30,
	0x58, 0xaa, 0x1c, 0xc3, 0xb1, 0xd7, 0x6f, 0x82, 0xe4, 0x7d, 0x38, 0x14, 0xf6, 0x65, 0x26, 0x94,
	0xd0, 0xca, 0xdb, 0xc4, 0x09, 0x33, 0x98, 0x35, 0x0f, 0xc7, 0xff, 0x46, 0x30, 0xac, 0x6e, 0x9b,
	0x7c, 0x04, 0xfb, 0x34, 0xd4, 0x32, 0xda, 0x59, 0xcb, 0x12, 0xf1, 0x43, 0x5f, 0x6a, 0x46, 0x71,
	0x7c, 0xed, 0x78, 0x66, 0x15, 0x40, 0x5e, 0xc0, 0x88, 0x19, 0x8e, 0x6b, 0x3f, 0xe4, 0xab, 0xba,
	0xb7, 0x0d, 0xfe, 0x06, 0xe9, 0x4b, 0x50, 0x75, 0x94, 0x0f, 0x23, 0x54, 0xbf, 0x2d, 0xb4, 0x26,
	0x38, 0xfe, 0xab, 0x03, 0xa3, 0x7a, 0x6b, 0xfa, 0xfc, 0xe6, 0xa1, 0x8e, 0xd1, 0xce, 0x3a, 0x96,
	0x08, 0xfe, 0x2d, 0xf4, 0xfd, 0xad, 0x4a, 0xf8, 0x03, 0xe6, 0xd7, 0x9b, 0x55, 0x7b, 0x32, 0x81,
	0x41, 0x31, 0xb3, 0xb8, 0x09, 0xa9, 0xb4, 0xc5, 0x53, 0x31, 0xe4, 0xbc, 0xf6, 0x7c, 0xde, 0x94,
	0x41, 0x8d, 0xf2, 0x3e, 0xca, 0xb6, 0x0d, 0x3d, 0xd3, 0xea, 0xa3, 0x64, 0xc8, 0x17, 0xf0, 0x98,
	0x69, 0xb5, 0x90, 0x82, 0x6d, 0xc6, 0xdf, 0x8e, 0xff, 0xd0, 0xff, 0xc0, 0xab, 0x6f, 0x7f, 0xfd,
	0xfa, 0x5e, 0xb8, 0x65, 0x3e, 0xf7, 0xe8, 0x54, 0x2f, 0x16, 0xcc, 0x3f, 0x5b, 0x49, 0xe7, 0x76,
	0x4a, 0xcd, 0x5c, 0x38, 0x93, 0xa7, 0xd3, 0x8c, 0xb2, 0x15, 0xbd, 0xe7, 0x78, 0xf2, 0x7c, 0xed,
	0x7b, 0x94, 0x3a, 0x6d, 0xa6, 0xc5, 0xa3, 0x9d, 0xef, 0xe1, 0xe8, 0xba, 0xf8, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x22, 0x13, 0xbe, 0xd1, 0x7f, 0x08, 0x00, 0x00,
}