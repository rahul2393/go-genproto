// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/account_budget_proposal_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The possible statuses of an AccountBudgetProposal.
type AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus int32

const (
	// Not specified.
	AccountBudgetProposalStatusEnum_UNSPECIFIED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AccountBudgetProposalStatusEnum_UNKNOWN AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 1
	// The proposal is pending approval.
	AccountBudgetProposalStatusEnum_PENDING AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 2
	// The proposal has been approved but the corresponding billing setup
	// has not.  This can occur for proposals that set up the first budget
	// when signing up for billing or when performing a change of bill-to
	// operation.
	AccountBudgetProposalStatusEnum_APPROVED_HELD AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 3
	// The proposal has been approved.
	AccountBudgetProposalStatusEnum_APPROVED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 4
	// The proposal has been cancelled by the user.
	AccountBudgetProposalStatusEnum_CANCELLED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 5
	// The proposal has been rejected by the user, e.g. by rejecting an
	// acceptance email.
	AccountBudgetProposalStatusEnum_REJECTED AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus = 6
)

var AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED_HELD",
	4: "APPROVED",
	5: "CANCELLED",
	6: "REJECTED",
}

var AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"PENDING":       2,
	"APPROVED_HELD": 3,
	"APPROVED":      4,
	"CANCELLED":     5,
	"REJECTED":      6,
}

func (x AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus) String() string {
	return proto.EnumName(AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_name, int32(x))
}

func (AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0efe433c1cc6a5f9, []int{0, 0}
}

// Message describing AccountBudgetProposal statuses.
type AccountBudgetProposalStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetProposalStatusEnum) Reset()         { *m = AccountBudgetProposalStatusEnum{} }
func (m *AccountBudgetProposalStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposalStatusEnum) ProtoMessage()    {}
func (*AccountBudgetProposalStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0efe433c1cc6a5f9, []int{0}
}

func (m *AccountBudgetProposalStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposalStatusEnum.Unmarshal(m, b)
}
func (m *AccountBudgetProposalStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposalStatusEnum.Marshal(b, m, deterministic)
}
func (m *AccountBudgetProposalStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposalStatusEnum.Merge(m, src)
}
func (m *AccountBudgetProposalStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposalStatusEnum.Size(m)
}
func (m *AccountBudgetProposalStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposalStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposalStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus", AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_name, AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus_value)
	proto.RegisterType((*AccountBudgetProposalStatusEnum)(nil), "google.ads.googleads.v3.enums.AccountBudgetProposalStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/account_budget_proposal_status.proto", fileDescriptor_0efe433c1cc6a5f9)
}

var fileDescriptor_0efe433c1cc6a5f9 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0xdd, 0x84, 0x5d, 0x76, 0xd7, 0x14, 0x35, 0xcd, 0xb1, 0x2d, 0x6a, 0xe1, 0x03, 0x9c, 0x43,
	0x6e, 0xee, 0xc9, 0x49, 0x5c, 0x4a, 0x8b, 0x42, 0x04, 0x25, 0x95, 0xaa, 0x48, 0x91, 0x21, 0x91,
	0x85, 0x04, 0x76, 0x84, 0x1d, 0x3e, 0xa1, 0x1f, 0xd2, 0x23, 0x9f, 0xd2, 0x4f, 0xe9, 0xbd, 0xf7,
	0x2a, 0x0e, 0xe1, 0x56, 0x2e, 0xd6, 0xf3, 0xbc, 0x99, 0x79, 0x33, 0x6f, 0x80, 0xc7, 0x84, 0x60,
	0xeb, 0xdc, 0xa1, 0x99, 0x74, 0x6a, 0x58, 0xa1, 0x9d, 0xeb, 0xe4, 0xbc, 0xdc, 0x48, 0x87, 0x2e,
	0x97, 0xa2, 0xe4, 0x2a, 0x5d, 0x94, 0x19, 0xcb, 0x55, 0x5a, 0x6c, 0x45, 0x21, 0x24, 0x5d, 0xa7,
	0x52, 0x51, 0x55, 0x4a, 0x58, 0x6c, 0x85, 0x12, 0x76, 0xaf, 0x2e, 0x84, 0x34, 0x93, 0xf0, 0xd8,
	0x03, 0xee, 0x5c, 0xa8, 0x7b, 0x5c, 0x5e, 0x37, 0x12, 0xc5, 0xca, 0xa1, 0x9c, 0x0b, 0x45, 0xd5,
	0x4a, 0xf0, 0x43, 0xf1, 0x60, 0x6f, 0x80, 0x1b, 0x5c, 0xab, 0x78, 0x5a, 0x24, 0x3a, 0x68, 0xcc,
	0xb4, 0x04, 0xe1, 0xe5, 0x66, 0xf0, 0x66, 0x80, 0xab, 0x13, 0x39, 0xf6, 0x39, 0xe8, 0xcc, 0xc3,
	0x59, 0x44, 0xfc, 0xd1, 0xfd, 0x88, 0x04, 0xd6, 0x2f, 0xbb, 0x03, 0xfe, 0xce, 0xc3, 0xa7, 0x70,
	0xf2, 0x12, 0x5a, 0x46, 0xf5, 0x89, 0x48, 0x18, 0x8c, 0xc2, 0xa1, 0x65, 0xda, 0x17, 0xa0, 0x8b,
	0xa3, 0x68, 0x3a, 0x89, 0x49, 0x90, 0x3e, 0x90, 0x71, 0x60, 0xb5, 0xec, 0x33, 0xf0, 0xaf, 0x09,
	0x59, 0xbf, 0xed, 0x2e, 0xf8, 0xef, 0xe3, 0xd0, 0x27, 0xe3, 0x31, 0x09, 0xac, 0x3f, 0x15, 0x39,
	0x25, 0x8f, 0xc4, 0x7f, 0x26, 0x81, 0xd5, 0xf6, 0xbe, 0x0c, 0xd0, 0x5f, 0x8a, 0x0d, 0x3c, 0xb9,
	0xb0, 0x77, 0x7b, 0x62, 0xd6, 0xa8, 0x5a, 0x3a, 0x32, 0x5e, 0x0f, 0xbe, 0x43, 0x26, 0xd6, 0x94,
	0x33, 0x28, 0xb6, 0xcc, 0x61, 0x39, 0xd7, 0x96, 0x34, 0x77, 0x28, 0x56, 0xf2, 0x87, 0xb3, 0xdc,
	0xe9, 0xf7, 0xdd, 0x6c, 0x0d, 0x31, 0xde, 0x9b, 0xbd, 0x61, 0xdd, 0x0a, 0x67, 0x12, 0xd6, 0xb0,
	0x42, 0xb1, 0x0b, 0x2b, 0xef, 0xe4, 0x47, 0xc3, 0x27, 0x38, 0x93, 0xc9, 0x91, 0x4f, 0x62, 0x37,
	0xd1, 0xfc, 0xa7, 0xd9, 0xaf, 0x83, 0x08, 0xe1, 0x4c, 0x22, 0x74, 0xcc, 0x40, 0x28, 0x76, 0x11,
	0xd2, 0x39, 0x8b, 0xb6, 0x1e, 0xcc, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x9c, 0x74, 0x91,
	0x2e, 0x02, 0x00, 0x00,
}