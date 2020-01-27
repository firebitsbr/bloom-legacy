// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/billing/service.proto

package billing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Plan struct {
	Id                       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Storage                  uint64   `protobuf:"varint,3,opt,name=storage,proto3" json:"storage,omitempty"`
	BitflowParallelDownloads uint64   `protobuf:"varint,4,opt,name=bitflow_parallel_downloads,json=bitflowParallelDownloads,proto3" json:"bitflow_parallel_downloads,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{1}
}

func (m *Plan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plan.Unmarshal(m, b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return xxx_messageInfo_Plan.Size(m)
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Plan) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Plan) GetBitflowParallelDownloads() uint64 {
	if m != nil {
		return m.BitflowParallelDownloads
	}
	return 0
}

type PlanList struct {
	Plans                []*Plan  `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanList) Reset()         { *m = PlanList{} }
func (m *PlanList) String() string { return proto.CompactTextString(m) }
func (*PlanList) ProtoMessage()    {}
func (*PlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{2}
}

func (m *PlanList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanList.Unmarshal(m, b)
}
func (m *PlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanList.Marshal(b, m, deterministic)
}
func (m *PlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanList.Merge(m, src)
}
func (m *PlanList) XXX_Size() int {
	return xxx_messageInfo_PlanList.Size(m)
}
func (m *PlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanList.DiscardUnknown(m)
}

var xxx_messageInfo_PlanList proto.InternalMessageInfo

func (m *PlanList) GetPlans() []*Plan {
	if m != nil {
		return m.Plans
	}
	return nil
}

type DetailedPlan struct {
	Id                       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StripeId                 string   `protobuf:"bytes,3,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	Storage                  uint64   `protobuf:"varint,4,opt,name=storage,proto3" json:"storage,omitempty"`
	BitflowParallelDownloads uint64   `protobuf:"varint,5,opt,name=bitflow_parallel_downloads,json=bitflowParallelDownloads,proto3" json:"bitflow_parallel_downloads,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *DetailedPlan) Reset()         { *m = DetailedPlan{} }
func (m *DetailedPlan) String() string { return proto.CompactTextString(m) }
func (*DetailedPlan) ProtoMessage()    {}
func (*DetailedPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{3}
}

func (m *DetailedPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailedPlan.Unmarshal(m, b)
}
func (m *DetailedPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailedPlan.Marshal(b, m, deterministic)
}
func (m *DetailedPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailedPlan.Merge(m, src)
}
func (m *DetailedPlan) XXX_Size() int {
	return xxx_messageInfo_DetailedPlan.Size(m)
}
func (m *DetailedPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailedPlan.DiscardUnknown(m)
}

var xxx_messageInfo_DetailedPlan proto.InternalMessageInfo

func (m *DetailedPlan) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DetailedPlan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DetailedPlan) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *DetailedPlan) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *DetailedPlan) GetBitflowParallelDownloads() uint64 {
	if m != nil {
		return m.BitflowParallelDownloads
	}
	return 0
}

type DetailedPlanList struct {
	Plans                []*DetailedPlan `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DetailedPlanList) Reset()         { *m = DetailedPlanList{} }
func (m *DetailedPlanList) String() string { return proto.CompactTextString(m) }
func (*DetailedPlanList) ProtoMessage()    {}
func (*DetailedPlanList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{4}
}

func (m *DetailedPlanList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailedPlanList.Unmarshal(m, b)
}
func (m *DetailedPlanList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailedPlanList.Marshal(b, m, deterministic)
}
func (m *DetailedPlanList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailedPlanList.Merge(m, src)
}
func (m *DetailedPlanList) XXX_Size() int {
	return xxx_messageInfo_DetailedPlanList.Size(m)
}
func (m *DetailedPlanList) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailedPlanList.DiscardUnknown(m)
}

var xxx_messageInfo_DetailedPlanList proto.InternalMessageInfo

func (m *DetailedPlanList) GetPlans() []*DetailedPlan {
	if m != nil {
		return m.Plans
	}
	return nil
}

type CreatePlanParams struct {
	Name                     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StripeId                 string   `protobuf:"bytes,2,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	Storage                  uint64   `protobuf:"varint,3,opt,name=storage,proto3" json:"storage,omitempty"`
	BitflowParallelDownloads uint64   `protobuf:"varint,4,opt,name=bitflow_parallel_downloads,json=bitflowParallelDownloads,proto3" json:"bitflow_parallel_downloads,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *CreatePlanParams) Reset()         { *m = CreatePlanParams{} }
func (m *CreatePlanParams) String() string { return proto.CompactTextString(m) }
func (*CreatePlanParams) ProtoMessage()    {}
func (*CreatePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{5}
}

func (m *CreatePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlanParams.Unmarshal(m, b)
}
func (m *CreatePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlanParams.Marshal(b, m, deterministic)
}
func (m *CreatePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlanParams.Merge(m, src)
}
func (m *CreatePlanParams) XXX_Size() int {
	return xxx_messageInfo_CreatePlanParams.Size(m)
}
func (m *CreatePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlanParams proto.InternalMessageInfo

func (m *CreatePlanParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePlanParams) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *CreatePlanParams) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *CreatePlanParams) GetBitflowParallelDownloads() uint64 {
	if m != nil {
		return m.BitflowParallelDownloads
	}
	return 0
}

type UpdatePlanParams struct {
	PlanId                   string   `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Name                     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StripeId                 string   `protobuf:"bytes,3,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	Storage                  uint64   `protobuf:"varint,4,opt,name=storage,proto3" json:"storage,omitempty"`
	BitflowParallelDownloads uint64   `protobuf:"varint,5,opt,name=bitflow_parallel_downloads,json=bitflowParallelDownloads,proto3" json:"bitflow_parallel_downloads,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *UpdatePlanParams) Reset()         { *m = UpdatePlanParams{} }
func (m *UpdatePlanParams) String() string { return proto.CompactTextString(m) }
func (*UpdatePlanParams) ProtoMessage()    {}
func (*UpdatePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{6}
}

func (m *UpdatePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlanParams.Unmarshal(m, b)
}
func (m *UpdatePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlanParams.Marshal(b, m, deterministic)
}
func (m *UpdatePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlanParams.Merge(m, src)
}
func (m *UpdatePlanParams) XXX_Size() int {
	return xxx_messageInfo_UpdatePlanParams.Size(m)
}
func (m *UpdatePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlanParams proto.InternalMessageInfo

func (m *UpdatePlanParams) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *UpdatePlanParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePlanParams) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *UpdatePlanParams) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *UpdatePlanParams) GetBitflowParallelDownloads() uint64 {
	if m != nil {
		return m.BitflowParallelDownloads
	}
	return 0
}

type DeletePlanParams struct {
	PlanId               string   `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePlanParams) Reset()         { *m = DeletePlanParams{} }
func (m *DeletePlanParams) String() string { return proto.CompactTextString(m) }
func (*DeletePlanParams) ProtoMessage()    {}
func (*DeletePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{7}
}

func (m *DeletePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePlanParams.Unmarshal(m, b)
}
func (m *DeletePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePlanParams.Marshal(b, m, deterministic)
}
func (m *DeletePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlanParams.Merge(m, src)
}
func (m *DeletePlanParams) XXX_Size() int {
	return xxx_messageInfo_DeletePlanParams.Size(m)
}
func (m *DeletePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlanParams proto.InternalMessageInfo

func (m *DeletePlanParams) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

type ChangePlanParams struct {
	PlanId               string   `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	StripePaymentMethod  string   `protobuf:"bytes,2,opt,name=stripe_payment_method,json=stripePaymentMethod,proto3" json:"stripe_payment_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePlanParams) Reset()         { *m = ChangePlanParams{} }
func (m *ChangePlanParams) String() string { return proto.CompactTextString(m) }
func (*ChangePlanParams) ProtoMessage()    {}
func (*ChangePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{8}
}

func (m *ChangePlanParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePlanParams.Unmarshal(m, b)
}
func (m *ChangePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePlanParams.Marshal(b, m, deterministic)
}
func (m *ChangePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePlanParams.Merge(m, src)
}
func (m *ChangePlanParams) XXX_Size() int {
	return xxx_messageInfo_ChangePlanParams.Size(m)
}
func (m *ChangePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePlanParams proto.InternalMessageInfo

func (m *ChangePlanParams) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *ChangePlanParams) GetStripePaymentMethod() string {
	if m != nil {
		return m.StripePaymentMethod
	}
	return ""
}

type AddPaymentMethodParams struct {
	StripePaymentMethod  string   `protobuf:"bytes,1,opt,name=stripe_payment_method,json=stripePaymentMethod,proto3" json:"stripe_payment_method,omitempty"`
	SetAsDefault         bool     `protobuf:"varint,2,opt,name=set_as_default,json=setAsDefault,proto3" json:"set_as_default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPaymentMethodParams) Reset()         { *m = AddPaymentMethodParams{} }
func (m *AddPaymentMethodParams) String() string { return proto.CompactTextString(m) }
func (*AddPaymentMethodParams) ProtoMessage()    {}
func (*AddPaymentMethodParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{9}
}

func (m *AddPaymentMethodParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPaymentMethodParams.Unmarshal(m, b)
}
func (m *AddPaymentMethodParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPaymentMethodParams.Marshal(b, m, deterministic)
}
func (m *AddPaymentMethodParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPaymentMethodParams.Merge(m, src)
}
func (m *AddPaymentMethodParams) XXX_Size() int {
	return xxx_messageInfo_AddPaymentMethodParams.Size(m)
}
func (m *AddPaymentMethodParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPaymentMethodParams.DiscardUnknown(m)
}

var xxx_messageInfo_AddPaymentMethodParams proto.InternalMessageInfo

func (m *AddPaymentMethodParams) GetStripePaymentMethod() string {
	if m != nil {
		return m.StripePaymentMethod
	}
	return ""
}

func (m *AddPaymentMethodParams) GetSetAsDefault() bool {
	if m != nil {
		return m.SetAsDefault
	}
	return false
}

// TODO
type PaymentMethod struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsDefault            bool     `protobuf:"varint,2,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	StripeId             string   `protobuf:"bytes,3,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardLast_4           string   `protobuf:"bytes,4,opt,name=card_last_4,json=cardLast4,proto3" json:"card_last_4,omitempty"`
	ExpirationMonth      int64    `protobuf:"varint,5,opt,name=expiration_month,json=expirationMonth,proto3" json:"expiration_month,omitempty"`
	ExpirationYear       int64    `protobuf:"varint,6,opt,name=expiration_year,json=expirationYear,proto3" json:"expiration_year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentMethod) Reset()         { *m = PaymentMethod{} }
func (m *PaymentMethod) String() string { return proto.CompactTextString(m) }
func (*PaymentMethod) ProtoMessage()    {}
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{10}
}

func (m *PaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethod.Unmarshal(m, b)
}
func (m *PaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethod.Marshal(b, m, deterministic)
}
func (m *PaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethod.Merge(m, src)
}
func (m *PaymentMethod) XXX_Size() int {
	return xxx_messageInfo_PaymentMethod.Size(m)
}
func (m *PaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethod proto.InternalMessageInfo

func (m *PaymentMethod) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentMethod) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *PaymentMethod) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *PaymentMethod) GetCardLast_4() string {
	if m != nil {
		return m.CardLast_4
	}
	return ""
}

func (m *PaymentMethod) GetExpirationMonth() int64 {
	if m != nil {
		return m.ExpirationMonth
	}
	return 0
}

func (m *PaymentMethod) GetExpirationYear() int64 {
	if m != nil {
		return m.ExpirationYear
	}
	return 0
}

type PaymentMethodList struct {
	PaymentMethods       []*PaymentMethod `protobuf:"bytes,1,rep,name=payment_methods,json=paymentMethods,proto3" json:"payment_methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PaymentMethodList) Reset()         { *m = PaymentMethodList{} }
func (m *PaymentMethodList) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodList) ProtoMessage()    {}
func (*PaymentMethodList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{11}
}

func (m *PaymentMethodList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodList.Unmarshal(m, b)
}
func (m *PaymentMethodList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodList.Marshal(b, m, deterministic)
}
func (m *PaymentMethodList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodList.Merge(m, src)
}
func (m *PaymentMethodList) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodList.Size(m)
}
func (m *PaymentMethodList) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodList.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodList proto.InternalMessageInfo

func (m *PaymentMethodList) GetPaymentMethods() []*PaymentMethod {
	if m != nil {
		return m.PaymentMethods
	}
	return nil
}

type RemovePaymentMethodParams struct {
	PaymentMethodId      string   `protobuf:"bytes,1,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePaymentMethodParams) Reset()         { *m = RemovePaymentMethodParams{} }
func (m *RemovePaymentMethodParams) String() string { return proto.CompactTextString(m) }
func (*RemovePaymentMethodParams) ProtoMessage()    {}
func (*RemovePaymentMethodParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{12}
}

func (m *RemovePaymentMethodParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePaymentMethodParams.Unmarshal(m, b)
}
func (m *RemovePaymentMethodParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePaymentMethodParams.Marshal(b, m, deterministic)
}
func (m *RemovePaymentMethodParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePaymentMethodParams.Merge(m, src)
}
func (m *RemovePaymentMethodParams) XXX_Size() int {
	return xxx_messageInfo_RemovePaymentMethodParams.Size(m)
}
func (m *RemovePaymentMethodParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePaymentMethodParams.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePaymentMethodParams proto.InternalMessageInfo

func (m *RemovePaymentMethodParams) GetPaymentMethodId() string {
	if m != nil {
		return m.PaymentMethodId
	}
	return ""
}

type UpdateDefaultPaymentMethodParams struct {
	PaymentMethodId      string   `protobuf:"bytes,1,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDefaultPaymentMethodParams) Reset()         { *m = UpdateDefaultPaymentMethodParams{} }
func (m *UpdateDefaultPaymentMethodParams) String() string { return proto.CompactTextString(m) }
func (*UpdateDefaultPaymentMethodParams) ProtoMessage()    {}
func (*UpdateDefaultPaymentMethodParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{13}
}

func (m *UpdateDefaultPaymentMethodParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDefaultPaymentMethodParams.Unmarshal(m, b)
}
func (m *UpdateDefaultPaymentMethodParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDefaultPaymentMethodParams.Marshal(b, m, deterministic)
}
func (m *UpdateDefaultPaymentMethodParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDefaultPaymentMethodParams.Merge(m, src)
}
func (m *UpdateDefaultPaymentMethodParams) XXX_Size() int {
	return xxx_messageInfo_UpdateDefaultPaymentMethodParams.Size(m)
}
func (m *UpdateDefaultPaymentMethodParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDefaultPaymentMethodParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDefaultPaymentMethodParams proto.InternalMessageInfo

func (m *UpdateDefaultPaymentMethodParams) GetPaymentMethodId() string {
	if m != nil {
		return m.PaymentMethodId
	}
	return ""
}

// TODO
type Invoice struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{14}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type InvoiceList struct {
	Invoices             []*Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *InvoiceList) Reset()         { *m = InvoiceList{} }
func (m *InvoiceList) String() string { return proto.CompactTextString(m) }
func (*InvoiceList) ProtoMessage()    {}
func (*InvoiceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c569122374058fad, []int{15}
}

func (m *InvoiceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceList.Unmarshal(m, b)
}
func (m *InvoiceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceList.Marshal(b, m, deterministic)
}
func (m *InvoiceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceList.Merge(m, src)
}
func (m *InvoiceList) XXX_Size() int {
	return xxx_messageInfo_InvoiceList.Size(m)
}
func (m *InvoiceList) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceList.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceList proto.InternalMessageInfo

func (m *InvoiceList) GetInvoices() []*Invoice {
	if m != nil {
		return m.Invoices
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "com.bloom42.billing.Empty")
	proto.RegisterType((*Plan)(nil), "com.bloom42.billing.Plan")
	proto.RegisterType((*PlanList)(nil), "com.bloom42.billing.PlanList")
	proto.RegisterType((*DetailedPlan)(nil), "com.bloom42.billing.DetailedPlan")
	proto.RegisterType((*DetailedPlanList)(nil), "com.bloom42.billing.DetailedPlanList")
	proto.RegisterType((*CreatePlanParams)(nil), "com.bloom42.billing.CreatePlanParams")
	proto.RegisterType((*UpdatePlanParams)(nil), "com.bloom42.billing.UpdatePlanParams")
	proto.RegisterType((*DeletePlanParams)(nil), "com.bloom42.billing.DeletePlanParams")
	proto.RegisterType((*ChangePlanParams)(nil), "com.bloom42.billing.ChangePlanParams")
	proto.RegisterType((*AddPaymentMethodParams)(nil), "com.bloom42.billing.AddPaymentMethodParams")
	proto.RegisterType((*PaymentMethod)(nil), "com.bloom42.billing.PaymentMethod")
	proto.RegisterType((*PaymentMethodList)(nil), "com.bloom42.billing.PaymentMethodList")
	proto.RegisterType((*RemovePaymentMethodParams)(nil), "com.bloom42.billing.RemovePaymentMethodParams")
	proto.RegisterType((*UpdateDefaultPaymentMethodParams)(nil), "com.bloom42.billing.UpdateDefaultPaymentMethodParams")
	proto.RegisterType((*Invoice)(nil), "com.bloom42.billing.Invoice")
	proto.RegisterType((*InvoiceList)(nil), "com.bloom42.billing.InvoiceList")
}

func init() { proto.RegisterFile("rpc/billing/service.proto", fileDescriptor_c569122374058fad) }

var fileDescriptor_c569122374058fad = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xba, 0x6e, 0x6d, 0xce, 0x46, 0xd7, 0x79, 0x02, 0xd2, 0xc0, 0x50, 0x89, 0x18, 0x0c,
	0x26, 0x75, 0x52, 0x19, 0x02, 0x09, 0x6e, 0xf6, 0x03, 0x53, 0xb5, 0x1f, 0x46, 0x24, 0x10, 0x70,
	0x13, 0xdc, 0xc6, 0xdb, 0x2c, 0x39, 0x3f, 0x8a, 0xcd, 0xc6, 0x6e, 0x79, 0x0d, 0x1e, 0x80, 0x17,
	0xe0, 0x3d, 0x78, 0x03, 0x9e, 0x05, 0xd9, 0x49, 0xd7, 0x26, 0x4a, 0xd2, 0x22, 0xb8, 0xe0, 0x6a,
	0xf3, 0xf1, 0xf1, 0xe7, 0xf3, 0x9d, 0x2f, 0xe7, 0x73, 0xa1, 0x15, 0x85, 0x83, 0x8d, 0x3e, 0x65,
	0x8c, 0xfa, 0xa7, 0x1b, 0x9c, 0x44, 0xe7, 0x74, 0x40, 0x3a, 0x61, 0x14, 0x88, 0x00, 0x2d, 0x0f,
	0x02, 0xaf, 0xd3, 0x67, 0x41, 0xe0, 0x6d, 0x76, 0x3b, 0x49, 0x8a, 0x55, 0x83, 0xd9, 0x97, 0x5e,
	0x28, 0x2e, 0xad, 0xaf, 0x1a, 0x54, 0x8f, 0x19, 0xf6, 0x51, 0x03, 0x2a, 0xd4, 0x35, 0xb4, 0xb6,
	0xb6, 0xa6, 0xdb, 0x15, 0xea, 0x22, 0x04, 0x55, 0x1f, 0x7b, 0xc4, 0xa8, 0xa8, 0x88, 0xfa, 0x1f,
	0x19, 0x50, 0xe3, 0x22, 0x88, 0xf0, 0x29, 0x31, 0x66, 0xda, 0xda, 0x5a, 0xd5, 0x1e, 0x2e, 0xd1,
	0x0b, 0x30, 0xfb, 0x54, 0x9c, 0xb0, 0xe0, 0xc2, 0x09, 0x71, 0x84, 0x19, 0x23, 0xcc, 0x71, 0x83,
	0x0b, 0x9f, 0x05, 0xd8, 0xe5, 0x46, 0x55, 0x25, 0x1b, 0x49, 0xc6, 0x71, 0x92, 0xb0, 0x3b, 0xdc,
	0xb7, 0x9e, 0x43, 0x5d, 0xd6, 0x70, 0x40, 0xb9, 0x40, 0x1b, 0x30, 0x1b, 0x32, 0xec, 0x73, 0x43,
	0x6b, 0xcf, 0xac, 0xcd, 0x77, 0x5b, 0x9d, 0x9c, 0xf2, 0x3b, 0x32, 0xdb, 0x8e, 0xf3, 0xac, 0xef,
	0x1a, 0x2c, 0xec, 0x12, 0x81, 0x29, 0x23, 0xee, 0xd4, 0x4c, 0x6e, 0x81, 0xce, 0x45, 0x44, 0x43,
	0xe2, 0x50, 0x57, 0x71, 0xd1, 0xed, 0x7a, 0x1c, 0xe8, 0xb9, 0xe3, 0x34, 0xab, 0x7f, 0x42, 0x73,
	0x76, 0x02, 0xcd, 0x7d, 0x68, 0x8e, 0x17, 0xaa, 0xe8, 0x3e, 0x4d, 0xd3, 0xbd, 0x9b, 0x4b, 0x77,
	0xfc, 0xd4, 0x90, 0xf6, 0x37, 0x0d, 0x9a, 0x3b, 0x11, 0xc1, 0x82, 0xc8, 0xa8, 0xbc, 0xcc, 0xe3,
	0x57, 0x54, 0xb5, 0x22, 0xaa, 0x95, 0x62, 0xaa, 0xff, 0x54, 0xd1, 0x1f, 0x1a, 0x34, 0xdf, 0x86,
	0x6e, 0xba, 0xba, 0x9b, 0x50, 0x93, 0xb5, 0x3b, 0x57, 0xea, 0xcc, 0xc9, 0x65, 0xef, 0x3f, 0x52,
	0x68, 0x5d, 0x2a, 0xc4, 0xc8, 0x54, 0x55, 0x5b, 0x0e, 0x34, 0x77, 0xce, 0xb0, 0x7f, 0x3a, 0x15,
	0xc5, 0x2e, 0x5c, 0x4f, 0xe8, 0x84, 0xf8, 0xd2, 0x23, 0xbe, 0x70, 0x3c, 0x22, 0xce, 0x82, 0xa1,
	0x22, 0xcb, 0xf1, 0xe6, 0x71, 0xbc, 0x77, 0xa8, 0xb6, 0xac, 0x08, 0x6e, 0x6c, 0xb9, 0x6e, 0x2a,
	0x96, 0x5c, 0x53, 0x88, 0xa6, 0x15, 0xa2, 0xa1, 0x7b, 0xd0, 0xe0, 0x44, 0x38, 0x98, 0x3b, 0x2e,
	0x39, 0xc1, 0x9f, 0x99, 0x50, 0x57, 0xd7, 0xed, 0x05, 0x4e, 0xc4, 0x16, 0xdf, 0x8d, 0x63, 0xd6,
	0x4f, 0x0d, 0xae, 0xa5, 0xcf, 0x65, 0xc7, 0x69, 0x05, 0x80, 0x66, 0x31, 0x74, 0x3a, 0x04, 0x28,
	0xd7, 0xed, 0x0e, 0xcc, 0x0f, 0x70, 0xe4, 0x3a, 0x0c, 0x73, 0xe1, 0x6c, 0x2a, 0xed, 0x74, 0x5b,
	0x97, 0xa1, 0x03, 0xcc, 0xc5, 0x26, 0x7a, 0x08, 0x4d, 0xf2, 0x25, 0xa4, 0x11, 0x16, 0x34, 0xf0,
	0x1d, 0x2f, 0xf0, 0xc5, 0x99, 0xd2, 0x6c, 0xc6, 0x5e, 0x1c, 0xc5, 0x0f, 0x65, 0x18, 0x3d, 0x80,
	0xb1, 0x90, 0x73, 0x49, 0x70, 0x64, 0xcc, 0xa9, 0xcc, 0xc6, 0x28, 0xfc, 0x81, 0xe0, 0xc8, 0xfa,
	0x04, 0x4b, 0x29, 0x42, 0x6a, 0xec, 0xf6, 0x61, 0x31, 0xdd, 0xb9, 0xe1, 0x00, 0x5a, 0xf9, 0x7e,
	0x33, 0x0e, 0x60, 0x37, 0xc2, 0xf1, 0x25, 0xb7, 0xf6, 0xa0, 0x65, 0x13, 0x2f, 0x38, 0x27, 0x79,
	0x52, 0x3d, 0x82, 0xa5, 0xf4, 0x4d, 0xa3, 0x6f, 0x63, 0x31, 0x85, 0xd3, 0x73, 0xad, 0x23, 0x68,
	0xc7, 0x43, 0x93, 0x34, 0xf3, 0x6f, 0xf1, 0x5a, 0x50, 0xeb, 0xf9, 0xe7, 0x01, 0x1d, 0x90, 0xac,
	0x8a, 0xd6, 0x1e, 0xcc, 0x27, 0x5b, 0xaa, 0x1f, 0xcf, 0xa0, 0x4e, 0xe3, 0xe5, 0xb0, 0x11, 0xb7,
	0x73, 0x1b, 0x91, 0x9c, 0xb1, 0xaf, 0xb2, 0xbb, 0xbf, 0x6a, 0x50, 0xdb, 0x8e, 0x77, 0xd1, 0x3b,
	0x80, 0x91, 0x25, 0xa1, 0xd5, 0x5c, 0x84, 0xac, 0x67, 0x99, 0x93, 0x2d, 0x4f, 0xe2, 0x8e, 0xcc,
	0xa4, 0x00, 0x37, 0xeb, 0x36, 0xd3, 0xe0, 0xbe, 0x06, 0x18, 0x8d, 0x7b, 0x01, 0x6e, 0xd6, 0x0f,
	0x4c, 0x33, 0x37, 0x4d, 0xbd, 0xa6, 0xe8, 0x0d, 0x2c, 0xc8, 0x76, 0x6e, 0x31, 0x26, 0x0f, 0x70,
	0x54, 0x92, 0x6b, 0xae, 0x4e, 0xac, 0x4f, 0x29, 0xf3, 0x0a, 0x74, 0xf9, 0x77, 0x32, 0xde, 0x4a,
	0xe1, 0x4b, 0xa9, 0x70, 0x8e, 0x00, 0x46, 0x6e, 0x55, 0xa4, 0x4d, 0xc6, 0xce, 0xcc, 0xe2, 0xd7,
	0x17, 0x0d, 0xa0, 0x99, 0x35, 0x27, 0xb4, 0x9e, 0x9b, 0x9e, 0xef, 0x61, 0xe6, 0x14, 0x93, 0x86,
	0x30, 0x2c, 0xe7, 0x4c, 0x16, 0xea, 0xe4, 0x1e, 0x2d, 0x9c, 0xc1, 0x52, 0xc9, 0x2e, 0xc0, 0x2c,
	0x9e, 0x39, 0xf4, 0xa4, 0xe4, 0x5b, 0x2b, 0x1e, 0xd2, 0xa9, 0xb8, 0xbd, 0x07, 0xa4, 0x84, 0x4d,
	0x79, 0x49, 0xa9, 0xc2, 0xf7, 0x27, 0xa3, 0x2a, 0xa9, 0x0f, 0xe2, 0xaf, 0x30, 0x99, 0xd5, 0x72,
	0xcc, 0x76, 0xd9, 0x98, 0x4b, 0x94, 0x6d, 0xfd, 0x63, 0x2d, 0x09, 0xf7, 0xe7, 0xd4, 0x2f, 0xca,
	0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0x8f, 0x67, 0x4c, 0x6e, 0x0a, 0x00, 0x00,
}
