// Code generated by protoc-gen-go. DO NOT EDIT.
// source: magnolia_service.proto

package service

import (
	fmt "fmt"
	cleaningcode "github.com/Vitality-South/magfabrics-api/pkg/cleaningcode"
	fabric "github.com/Vitality-South/magfabrics-api/pkg/fabric"
	inventory "github.com/Vitality-South/magfabrics-api/pkg/inventory"
	taxonomy "github.com/Vitality-South/magfabrics-api/pkg/taxonomy"
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

type GetAllFabricsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllFabricsRequest) Reset()         { *m = GetAllFabricsRequest{} }
func (m *GetAllFabricsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllFabricsRequest) ProtoMessage()    {}
func (*GetAllFabricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{0}
}

func (m *GetAllFabricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFabricsRequest.Unmarshal(m, b)
}
func (m *GetAllFabricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFabricsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllFabricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFabricsRequest.Merge(m, src)
}
func (m *GetAllFabricsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllFabricsRequest.Size(m)
}
func (m *GetAllFabricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFabricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFabricsRequest proto.InternalMessageInfo

type GetAllFabricsResponse struct {
	Fabrics              []*fabric.Fabric `protobuf:"bytes,1,rep,name=fabrics,proto3" json:"fabrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetAllFabricsResponse) Reset()         { *m = GetAllFabricsResponse{} }
func (m *GetAllFabricsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllFabricsResponse) ProtoMessage()    {}
func (*GetAllFabricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{1}
}

func (m *GetAllFabricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFabricsResponse.Unmarshal(m, b)
}
func (m *GetAllFabricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFabricsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllFabricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFabricsResponse.Merge(m, src)
}
func (m *GetAllFabricsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllFabricsResponse.Size(m)
}
func (m *GetAllFabricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFabricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFabricsResponse proto.InternalMessageInfo

func (m *GetAllFabricsResponse) GetFabrics() []*fabric.Fabric {
	if m != nil {
		return m.Fabrics
	}
	return nil
}

type GetAllFabricsWithoutInventoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllFabricsWithoutInventoryRequest) Reset()         { *m = GetAllFabricsWithoutInventoryRequest{} }
func (m *GetAllFabricsWithoutInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllFabricsWithoutInventoryRequest) ProtoMessage()    {}
func (*GetAllFabricsWithoutInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{2}
}

func (m *GetAllFabricsWithoutInventoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFabricsWithoutInventoryRequest.Unmarshal(m, b)
}
func (m *GetAllFabricsWithoutInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFabricsWithoutInventoryRequest.Marshal(b, m, deterministic)
}
func (m *GetAllFabricsWithoutInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFabricsWithoutInventoryRequest.Merge(m, src)
}
func (m *GetAllFabricsWithoutInventoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllFabricsWithoutInventoryRequest.Size(m)
}
func (m *GetAllFabricsWithoutInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFabricsWithoutInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFabricsWithoutInventoryRequest proto.InternalMessageInfo

type GetAllFabricsWithoutInventoryResponse struct {
	Fabrics              []*fabric.Fabric `protobuf:"bytes,1,rep,name=fabrics,proto3" json:"fabrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetAllFabricsWithoutInventoryResponse) Reset()         { *m = GetAllFabricsWithoutInventoryResponse{} }
func (m *GetAllFabricsWithoutInventoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllFabricsWithoutInventoryResponse) ProtoMessage()    {}
func (*GetAllFabricsWithoutInventoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{3}
}

func (m *GetAllFabricsWithoutInventoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFabricsWithoutInventoryResponse.Unmarshal(m, b)
}
func (m *GetAllFabricsWithoutInventoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFabricsWithoutInventoryResponse.Marshal(b, m, deterministic)
}
func (m *GetAllFabricsWithoutInventoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFabricsWithoutInventoryResponse.Merge(m, src)
}
func (m *GetAllFabricsWithoutInventoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllFabricsWithoutInventoryResponse.Size(m)
}
func (m *GetAllFabricsWithoutInventoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFabricsWithoutInventoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFabricsWithoutInventoryResponse proto.InternalMessageInfo

func (m *GetAllFabricsWithoutInventoryResponse) GetFabrics() []*fabric.Fabric {
	if m != nil {
		return m.Fabrics
	}
	return nil
}

type GetAllInventoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllInventoryRequest) Reset()         { *m = GetAllInventoryRequest{} }
func (m *GetAllInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllInventoryRequest) ProtoMessage()    {}
func (*GetAllInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{4}
}

func (m *GetAllInventoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllInventoryRequest.Unmarshal(m, b)
}
func (m *GetAllInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllInventoryRequest.Marshal(b, m, deterministic)
}
func (m *GetAllInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllInventoryRequest.Merge(m, src)
}
func (m *GetAllInventoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllInventoryRequest.Size(m)
}
func (m *GetAllInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllInventoryRequest proto.InternalMessageInfo

type GetAllInventoryResponse struct {
	Inventory            []*inventory.Inventory `protobuf:"bytes,1,rep,name=inventory,proto3" json:"inventory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetAllInventoryResponse) Reset()         { *m = GetAllInventoryResponse{} }
func (m *GetAllInventoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllInventoryResponse) ProtoMessage()    {}
func (*GetAllInventoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{5}
}

func (m *GetAllInventoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllInventoryResponse.Unmarshal(m, b)
}
func (m *GetAllInventoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllInventoryResponse.Marshal(b, m, deterministic)
}
func (m *GetAllInventoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllInventoryResponse.Merge(m, src)
}
func (m *GetAllInventoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllInventoryResponse.Size(m)
}
func (m *GetAllInventoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllInventoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllInventoryResponse proto.InternalMessageInfo

func (m *GetAllInventoryResponse) GetInventory() []*inventory.Inventory {
	if m != nil {
		return m.Inventory
	}
	return nil
}

type GetFabricByIDRequest struct {
	FabricId             string   `protobuf:"bytes,1,opt,name=fabric_id,json=fabricId,proto3" json:"fabric_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFabricByIDRequest) Reset()         { *m = GetFabricByIDRequest{} }
func (m *GetFabricByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetFabricByIDRequest) ProtoMessage()    {}
func (*GetFabricByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{6}
}

func (m *GetFabricByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFabricByIDRequest.Unmarshal(m, b)
}
func (m *GetFabricByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFabricByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetFabricByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFabricByIDRequest.Merge(m, src)
}
func (m *GetFabricByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetFabricByIDRequest.Size(m)
}
func (m *GetFabricByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFabricByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFabricByIDRequest proto.InternalMessageInfo

func (m *GetFabricByIDRequest) GetFabricId() string {
	if m != nil {
		return m.FabricId
	}
	return ""
}

type GetFabricByIDResponse struct {
	Fabric               *fabric.Fabric `protobuf:"bytes,1,opt,name=fabric,proto3" json:"fabric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetFabricByIDResponse) Reset()         { *m = GetFabricByIDResponse{} }
func (m *GetFabricByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetFabricByIDResponse) ProtoMessage()    {}
func (*GetFabricByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{7}
}

func (m *GetFabricByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFabricByIDResponse.Unmarshal(m, b)
}
func (m *GetFabricByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFabricByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetFabricByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFabricByIDResponse.Merge(m, src)
}
func (m *GetFabricByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetFabricByIDResponse.Size(m)
}
func (m *GetFabricByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFabricByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFabricByIDResponse proto.InternalMessageInfo

func (m *GetFabricByIDResponse) GetFabric() *fabric.Fabric {
	if m != nil {
		return m.Fabric
	}
	return nil
}

type GetFabricByNameRequest struct {
	FabricName           string   `protobuf:"bytes,1,opt,name=fabric_name,json=fabricName,proto3" json:"fabric_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFabricByNameRequest) Reset()         { *m = GetFabricByNameRequest{} }
func (m *GetFabricByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetFabricByNameRequest) ProtoMessage()    {}
func (*GetFabricByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{8}
}

func (m *GetFabricByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFabricByNameRequest.Unmarshal(m, b)
}
func (m *GetFabricByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFabricByNameRequest.Marshal(b, m, deterministic)
}
func (m *GetFabricByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFabricByNameRequest.Merge(m, src)
}
func (m *GetFabricByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetFabricByNameRequest.Size(m)
}
func (m *GetFabricByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFabricByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFabricByNameRequest proto.InternalMessageInfo

func (m *GetFabricByNameRequest) GetFabricName() string {
	if m != nil {
		return m.FabricName
	}
	return ""
}

type GetFabricByNameResponse struct {
	Fabric               *fabric.Fabric `protobuf:"bytes,1,opt,name=fabric,proto3" json:"fabric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetFabricByNameResponse) Reset()         { *m = GetFabricByNameResponse{} }
func (m *GetFabricByNameResponse) String() string { return proto.CompactTextString(m) }
func (*GetFabricByNameResponse) ProtoMessage()    {}
func (*GetFabricByNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{9}
}

func (m *GetFabricByNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFabricByNameResponse.Unmarshal(m, b)
}
func (m *GetFabricByNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFabricByNameResponse.Marshal(b, m, deterministic)
}
func (m *GetFabricByNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFabricByNameResponse.Merge(m, src)
}
func (m *GetFabricByNameResponse) XXX_Size() int {
	return xxx_messageInfo_GetFabricByNameResponse.Size(m)
}
func (m *GetFabricByNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFabricByNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFabricByNameResponse proto.InternalMessageInfo

func (m *GetFabricByNameResponse) GetFabric() *fabric.Fabric {
	if m != nil {
		return m.Fabric
	}
	return nil
}

type GetAllFabricTaxonomyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllFabricTaxonomyRequest) Reset()         { *m = GetAllFabricTaxonomyRequest{} }
func (m *GetAllFabricTaxonomyRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllFabricTaxonomyRequest) ProtoMessage()    {}
func (*GetAllFabricTaxonomyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{10}
}

func (m *GetAllFabricTaxonomyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFabricTaxonomyRequest.Unmarshal(m, b)
}
func (m *GetAllFabricTaxonomyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFabricTaxonomyRequest.Marshal(b, m, deterministic)
}
func (m *GetAllFabricTaxonomyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFabricTaxonomyRequest.Merge(m, src)
}
func (m *GetAllFabricTaxonomyRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllFabricTaxonomyRequest.Size(m)
}
func (m *GetAllFabricTaxonomyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFabricTaxonomyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFabricTaxonomyRequest proto.InternalMessageInfo

type GetAllFabricTaxonomyResponse struct {
	Taxonomy             *taxonomy.Taxonomy `protobuf:"bytes,1,opt,name=taxonomy,proto3" json:"taxonomy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAllFabricTaxonomyResponse) Reset()         { *m = GetAllFabricTaxonomyResponse{} }
func (m *GetAllFabricTaxonomyResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllFabricTaxonomyResponse) ProtoMessage()    {}
func (*GetAllFabricTaxonomyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{11}
}

func (m *GetAllFabricTaxonomyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFabricTaxonomyResponse.Unmarshal(m, b)
}
func (m *GetAllFabricTaxonomyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFabricTaxonomyResponse.Marshal(b, m, deterministic)
}
func (m *GetAllFabricTaxonomyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFabricTaxonomyResponse.Merge(m, src)
}
func (m *GetAllFabricTaxonomyResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllFabricTaxonomyResponse.Size(m)
}
func (m *GetAllFabricTaxonomyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFabricTaxonomyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFabricTaxonomyResponse proto.InternalMessageInfo

func (m *GetAllFabricTaxonomyResponse) GetTaxonomy() *taxonomy.Taxonomy {
	if m != nil {
		return m.Taxonomy
	}
	return nil
}

type GetCleaningCodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCleaningCodesRequest) Reset()         { *m = GetCleaningCodesRequest{} }
func (m *GetCleaningCodesRequest) String() string { return proto.CompactTextString(m) }
func (*GetCleaningCodesRequest) ProtoMessage()    {}
func (*GetCleaningCodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{12}
}

func (m *GetCleaningCodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCleaningCodesRequest.Unmarshal(m, b)
}
func (m *GetCleaningCodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCleaningCodesRequest.Marshal(b, m, deterministic)
}
func (m *GetCleaningCodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCleaningCodesRequest.Merge(m, src)
}
func (m *GetCleaningCodesRequest) XXX_Size() int {
	return xxx_messageInfo_GetCleaningCodesRequest.Size(m)
}
func (m *GetCleaningCodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCleaningCodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCleaningCodesRequest proto.InternalMessageInfo

type GetCleaningCodesResponse struct {
	CleaningCodes        map[string]*cleaningcode.CleaningCode `protobuf:"bytes,1,rep,name=cleaning_codes,json=cleaningCodes,proto3" json:"cleaning_codes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *GetCleaningCodesResponse) Reset()         { *m = GetCleaningCodesResponse{} }
func (m *GetCleaningCodesResponse) String() string { return proto.CompactTextString(m) }
func (*GetCleaningCodesResponse) ProtoMessage()    {}
func (*GetCleaningCodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{13}
}

func (m *GetCleaningCodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCleaningCodesResponse.Unmarshal(m, b)
}
func (m *GetCleaningCodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCleaningCodesResponse.Marshal(b, m, deterministic)
}
func (m *GetCleaningCodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCleaningCodesResponse.Merge(m, src)
}
func (m *GetCleaningCodesResponse) XXX_Size() int {
	return xxx_messageInfo_GetCleaningCodesResponse.Size(m)
}
func (m *GetCleaningCodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCleaningCodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCleaningCodesResponse proto.InternalMessageInfo

func (m *GetCleaningCodesResponse) GetCleaningCodes() map[string]*cleaningcode.CleaningCode {
	if m != nil {
		return m.CleaningCodes
	}
	return nil
}

type GetFabricBySKURequest struct {
	FabricSku            string   `protobuf:"bytes,1,opt,name=fabric_sku,json=fabricSku,proto3" json:"fabric_sku,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFabricBySKURequest) Reset()         { *m = GetFabricBySKURequest{} }
func (m *GetFabricBySKURequest) String() string { return proto.CompactTextString(m) }
func (*GetFabricBySKURequest) ProtoMessage()    {}
func (*GetFabricBySKURequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{14}
}

func (m *GetFabricBySKURequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFabricBySKURequest.Unmarshal(m, b)
}
func (m *GetFabricBySKURequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFabricBySKURequest.Marshal(b, m, deterministic)
}
func (m *GetFabricBySKURequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFabricBySKURequest.Merge(m, src)
}
func (m *GetFabricBySKURequest) XXX_Size() int {
	return xxx_messageInfo_GetFabricBySKURequest.Size(m)
}
func (m *GetFabricBySKURequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFabricBySKURequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFabricBySKURequest proto.InternalMessageInfo

func (m *GetFabricBySKURequest) GetFabricSku() string {
	if m != nil {
		return m.FabricSku
	}
	return ""
}

type GetFabricBySKUResponse struct {
	Fabric               *fabric.Fabric `protobuf:"bytes,1,opt,name=fabric,proto3" json:"fabric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetFabricBySKUResponse) Reset()         { *m = GetFabricBySKUResponse{} }
func (m *GetFabricBySKUResponse) String() string { return proto.CompactTextString(m) }
func (*GetFabricBySKUResponse) ProtoMessage()    {}
func (*GetFabricBySKUResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14f18cd129ec495, []int{15}
}

func (m *GetFabricBySKUResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFabricBySKUResponse.Unmarshal(m, b)
}
func (m *GetFabricBySKUResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFabricBySKUResponse.Marshal(b, m, deterministic)
}
func (m *GetFabricBySKUResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFabricBySKUResponse.Merge(m, src)
}
func (m *GetFabricBySKUResponse) XXX_Size() int {
	return xxx_messageInfo_GetFabricBySKUResponse.Size(m)
}
func (m *GetFabricBySKUResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFabricBySKUResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFabricBySKUResponse proto.InternalMessageInfo

func (m *GetFabricBySKUResponse) GetFabric() *fabric.Fabric {
	if m != nil {
		return m.Fabric
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllFabricsRequest)(nil), "magnoliafabrics.GetAllFabricsRequest")
	proto.RegisterType((*GetAllFabricsResponse)(nil), "magnoliafabrics.GetAllFabricsResponse")
	proto.RegisterType((*GetAllFabricsWithoutInventoryRequest)(nil), "magnoliafabrics.GetAllFabricsWithoutInventoryRequest")
	proto.RegisterType((*GetAllFabricsWithoutInventoryResponse)(nil), "magnoliafabrics.GetAllFabricsWithoutInventoryResponse")
	proto.RegisterType((*GetAllInventoryRequest)(nil), "magnoliafabrics.GetAllInventoryRequest")
	proto.RegisterType((*GetAllInventoryResponse)(nil), "magnoliafabrics.GetAllInventoryResponse")
	proto.RegisterType((*GetFabricByIDRequest)(nil), "magnoliafabrics.GetFabricByIDRequest")
	proto.RegisterType((*GetFabricByIDResponse)(nil), "magnoliafabrics.GetFabricByIDResponse")
	proto.RegisterType((*GetFabricByNameRequest)(nil), "magnoliafabrics.GetFabricByNameRequest")
	proto.RegisterType((*GetFabricByNameResponse)(nil), "magnoliafabrics.GetFabricByNameResponse")
	proto.RegisterType((*GetAllFabricTaxonomyRequest)(nil), "magnoliafabrics.GetAllFabricTaxonomyRequest")
	proto.RegisterType((*GetAllFabricTaxonomyResponse)(nil), "magnoliafabrics.GetAllFabricTaxonomyResponse")
	proto.RegisterType((*GetCleaningCodesRequest)(nil), "magnoliafabrics.GetCleaningCodesRequest")
	proto.RegisterType((*GetCleaningCodesResponse)(nil), "magnoliafabrics.GetCleaningCodesResponse")
	proto.RegisterMapType((map[string]*cleaningcode.CleaningCode)(nil), "magnoliafabrics.GetCleaningCodesResponse.CleaningCodesEntry")
	proto.RegisterType((*GetFabricBySKURequest)(nil), "magnoliafabrics.GetFabricBySKURequest")
	proto.RegisterType((*GetFabricBySKUResponse)(nil), "magnoliafabrics.GetFabricBySKUResponse")
}

func init() {
	proto.RegisterFile("magnolia_service.proto", fileDescriptor_d14f18cd129ec495)
}

var fileDescriptor_d14f18cd129ec495 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5d, 0x4f, 0x13, 0x41,
	0x14, 0x65, 0x21, 0xf2, 0x71, 0x09, 0x42, 0x06, 0x2d, 0x65, 0xb1, 0x91, 0x4c, 0x04, 0x6a, 0x22,
	0xad, 0x42, 0x20, 0x68, 0x7c, 0x11, 0xfc, 0x2a, 0x46, 0x1f, 0x5a, 0xd1, 0x84, 0x97, 0x75, 0xd8,
	0x0e, 0x65, 0xd2, 0xdd, 0x9d, 0xda, 0x9d, 0x6d, 0xec, 0x1f, 0xf0, 0xd9, 0x9f, 0xe8, 0x4f, 0x31,
	0xdd, 0x99, 0xd9, 0x76, 0x3b, 0xdb, 0x0f, 0xfb, 0xd6, 0xdc, 0x39, 0xf7, 0x9e, 0x73, 0xef, 0xdc,
	0x39, 0x5d, 0xc8, 0xf9, 0xa4, 0x11, 0x70, 0x8f, 0x11, 0x27, 0xa4, 0xed, 0x0e, 0x73, 0x69, 0xa9,
	0xd5, 0xe6, 0x82, 0xa3, 0x75, 0x1d, 0xbf, 0x25, 0x37, 0x6d, 0xe6, 0x86, 0xf6, 0x26, 0x69, 0x31,
	0x87, 0x05, 0x1d, 0x1a, 0x08, 0xde, 0xee, 0x4a, 0x94, 0xbd, 0xd1, 0x0b, 0x4a, 0x84, 0x8a, 0xa0,
	0x5e, 0x44, 0x90, 0x5f, 0x3c, 0xe0, 0xbe, 0x46, 0x6d, 0xba, 0x1e, 0x25, 0x01, 0x0b, 0x1a, 0x8e,
	0xcb, 0xeb, 0x8a, 0x00, 0xe7, 0xe0, 0xc1, 0x07, 0x2a, 0xde, 0x78, 0xde, 0x7b, 0x49, 0x50, 0xa5,
	0x3f, 0x23, 0x1a, 0x0a, 0x7c, 0x09, 0x0f, 0x87, 0xe2, 0x61, 0x8b, 0x07, 0x21, 0x45, 0x2f, 0x60,
	0x49, 0x69, 0xc9, 0x5b, 0xbb, 0x0b, 0xc5, 0xd5, 0xa3, 0xad, 0xd2, 0x90, 0xc6, 0x92, 0x4c, 0xa9,
	0x6a, 0x1c, 0xde, 0x87, 0x27, 0xa9, 0x5a, 0xdf, 0x99, 0xb8, 0xe3, 0x91, 0xa8, 0xe8, 0x2e, 0x34,
	0xe7, 0x35, 0xec, 0x4d, 0xc0, 0xcd, 0xae, 0x21, 0x0f, 0x39, 0x59, 0xdb, 0x60, 0xad, 0xc1, 0x96,
	0x71, 0xa2, 0x78, 0xce, 0x60, 0x25, 0x19, 0xb5, 0x62, 0xb2, 0x0d, 0xa6, 0x7e, 0x5a, 0x1f, 0x8c,
	0x8f, 0xe3, 0xb1, 0x4a, 0x11, 0xe7, 0xdd, 0xca, 0x5b, 0x45, 0x86, 0x76, 0x60, 0x45, 0xe6, 0x39,
	0xac, 0x9e, 0xb7, 0x76, 0xad, 0xe2, 0x4a, 0x75, 0x59, 0x06, 0x2a, 0x75, 0xfc, 0x31, 0x9e, 0xf9,
	0x60, 0x92, 0xd2, 0x51, 0x86, 0x45, 0x09, 0x8a, 0x53, 0xc6, 0xb4, 0xab, 0x60, 0xf8, 0x65, 0xdc,
	0xad, 0xae, 0xf4, 0x85, 0xf8, 0x54, 0x0b, 0x78, 0x0c, 0xab, 0x4a, 0x40, 0x40, 0x7c, 0xaa, 0x24,
	0x80, 0x0c, 0xf5, 0x70, 0xf8, 0x32, 0x1e, 0x47, 0x3a, 0x75, 0x56, 0x19, 0x05, 0xd8, 0x19, 0xbc,
	0xd0, 0xaf, 0x6a, 0x1f, 0xf5, 0xe4, 0xaf, 0xe0, 0x51, 0xf6, 0xb1, 0xe2, 0x3b, 0x81, 0x65, 0xbd,
	0xc2, 0x8a, 0x71, 0xdb, 0x60, 0x4c, 0x92, 0x12, 0x28, 0xde, 0x8e, 0x3b, 0xb8, 0x50, 0xcb, 0x7e,
	0xc1, 0xeb, 0x34, 0xd9, 0xea, 0xbf, 0x16, 0xe4, 0xcd, 0x33, 0x45, 0xe7, 0xc2, 0xfd, 0xd4, 0x0b,
	0xd1, 0xcb, 0xf5, 0xda, 0x20, 0x1d, 0x55, 0xa2, 0x94, 0x8a, 0xbe, 0x0b, 0x44, 0xbb, 0x5b, 0x5d,
	0x73, 0x07, 0x63, 0xb6, 0x03, 0xc8, 0x04, 0xa1, 0x0d, 0x58, 0x68, 0xd2, 0xae, 0xba, 0x8d, 0xde,
	0x4f, 0x74, 0x0c, 0xf7, 0x3a, 0xc4, 0x8b, 0x68, 0x7e, 0x3e, 0x6e, 0xbc, 0x60, 0x68, 0x18, 0xac,
	0x52, 0x95, 0xd8, 0x57, 0xf3, 0x67, 0x16, 0x3e, 0x4d, 0x2d, 0x51, 0xed, 0xd3, 0x95, 0xbe, 0xf9,
	0x02, 0xa8, 0x6b, 0x76, 0xc2, 0x66, 0xa4, 0xa8, 0xd4, 0x32, 0xd6, 0x9a, 0x11, 0xae, 0xa4, 0x56,
	0x26, 0xce, 0x9b, 0xf1, 0xda, 0x8f, 0x7e, 0x2f, 0x41, 0xee, 0xb3, 0x82, 0xa8, 0xa7, 0x5c, 0x93,
	0xae, 0x86, 0x7e, 0xc0, 0x5a, 0xea, 0x89, 0xa3, 0xbd, 0xac, 0xe1, 0x1a, 0x76, 0x64, 0xef, 0x4f,
	0x82, 0x49, 0xad, 0x78, 0x0e, 0xfd, 0xb1, 0xa0, 0x30, 0xd6, 0x45, 0xd0, 0xc9, 0xf8, 0x5a, 0x23,
	0xdc, 0xc9, 0x3e, 0xfd, 0xdf, 0xb4, 0x44, 0xd2, 0x2d, 0xac, 0x0f, 0x39, 0x0c, 0x3a, 0x18, 0x51,
	0xcc, 0x60, 0x2d, 0x4e, 0x06, 0x26, 0x3c, 0x72, 0xb8, 0x7d, 0xff, 0xc8, 0x1e, 0xae, 0x61, 0x4a,
	0xd9, 0xc3, 0x35, 0x6d, 0x28, 0xe9, 0x64, 0xd0, 0x1c, 0xb2, 0x3b, 0xc9, 0x70, 0x9e, 0xec, 0x4e,
	0xb2, 0x7c, 0x06, 0xcf, 0xf5, 0x9e, 0x62, 0x7a, 0x19, 0xd1, 0x58, 0x8d, 0xfd, 0x2d, 0xb7, 0x0f,
	0x26, 0xe2, 0x12, 0x92, 0x28, 0xfd, 0xd7, 0xa7, 0x9d, 0x04, 0x3d, 0x1b, 0x7b, 0xd1, 0x43, 0x26,
	0x66, 0x1f, 0x4e, 0x89, 0x4e, 0x68, 0x19, 0x6c, 0x0c, 0xfb, 0x07, 0x2a, 0x4e, 0x61, 0x31, 0x92,
	0xee, 0xe9, 0xd4, 0x66, 0x84, 0xe7, 0xce, 0x8f, 0xae, 0x9f, 0x37, 0x98, 0xb8, 0x8b, 0x6e, 0x4a,
	0x2e, 0xf7, 0xcb, 0xdf, 0x98, 0x20, 0x1e, 0x13, 0xdd, 0xc3, 0x1a, 0x8f, 0xc4, 0x5d, 0xd9, 0x27,
	0x0d, 0x55, 0xe2, 0x90, 0xb4, 0x58, 0x59, 0x7d, 0x77, 0xdc, 0x2c, 0xc6, 0xdf, 0x05, 0xc7, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x2e, 0xfd, 0x10, 0x92, 0x08, 0x00, 0x00,
}
