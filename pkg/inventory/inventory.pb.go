// Copyright (c) 2022 Vitality South, LLC <chris@vitalitysouth.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: inventory.proto

package inventory

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

type Inventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamo:"FabricProductCode"
	ProductCode string `protobuf:"bytes,1,opt,name=ProductCode,proto3" json:"ProductCode,omitempty" dynamo:"FabricProductCode"`
	// @gotags: dynamo:"InventoryRolls"
	Rolls []*Inventory_Roll `protobuf:"bytes,2,rep,name=Rolls,proto3" json:"Rolls,omitempty" dynamo:"InventoryRolls"`
	// @gotags: dynamo:"InventoryTotal"
	Total string `protobuf:"bytes,3,opt,name=Total,proto3" json:"Total,omitempty" dynamo:"InventoryTotal"`
	// @gotags: dynamo:"InventoryPO"
	PurchaseOrders []*Inventory_PO `protobuf:"bytes,4,rep,name=PurchaseOrders,proto3" json:"PurchaseOrders,omitempty" dynamo:"InventoryPO"`
	// @gotags: dynamo:"InventoryTotalOnPO"
	TotalOnPO string `protobuf:"bytes,5,opt,name=TotalOnPO,proto3" json:"TotalOnPO,omitempty" dynamo:"InventoryTotalOnPO"`
	// @gotags: dynamo:"InventoryOnBackorder"
	OnBackorder string `protobuf:"bytes,6,opt,name=OnBackorder,proto3" json:"OnBackorder,omitempty" dynamo:"InventoryOnBackorder"`
}

func (x *Inventory) Reset() {
	*x = Inventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory) ProtoMessage() {}

func (x *Inventory) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory.ProtoReflect.Descriptor instead.
func (*Inventory) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Inventory) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *Inventory) GetRolls() []*Inventory_Roll {
	if x != nil {
		return x.Rolls
	}
	return nil
}

func (x *Inventory) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *Inventory) GetPurchaseOrders() []*Inventory_PO {
	if x != nil {
		return x.PurchaseOrders
	}
	return nil
}

func (x *Inventory) GetTotalOnPO() string {
	if x != nil {
		return x.TotalOnPO
	}
	return ""
}

func (x *Inventory) GetOnBackorder() string {
	if x != nil {
		return x.OnBackorder
	}
	return ""
}

type Inventory_Roll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamo:"InventoryRollID"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty" dynamo:"InventoryRollID"`
	// @gotags: dynamo:"InventoryRollQuantity"
	Quantity string `protobuf:"bytes,2,opt,name=Quantity,proto3" json:"Quantity,omitempty" dynamo:"InventoryRollQuantity"`
}

func (x *Inventory_Roll) Reset() {
	*x = Inventory_Roll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory_Roll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory_Roll) ProtoMessage() {}

func (x *Inventory_Roll) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory_Roll.ProtoReflect.Descriptor instead.
func (*Inventory_Roll) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Inventory_Roll) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Inventory_Roll) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type Inventory_PO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamo:"InventoryPOETA"
	ETA string `protobuf:"bytes,1,opt,name=ETA,proto3" json:"ETA,omitempty" dynamo:"InventoryPOETA"`
	// @gotags: dynamo:"InventoryPOAmount"
	Amount string `protobuf:"bytes,2,opt,name=Amount,proto3" json:"Amount,omitempty" dynamo:"InventoryPOAmount"`
}

func (x *Inventory_PO) Reset() {
	*x = Inventory_PO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory_PO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory_PO) ProtoMessage() {}

func (x *Inventory_PO) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory_PO.ProtoReflect.Descriptor instead.
func (*Inventory_PO) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Inventory_PO) GetETA() string {
	if x != nil {
		return x.ETA
	}
	return ""
}

func (x *Inventory_PO) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x6d, 0x61, 0x67, 0x6e, 0x6f, 0x6c, 0x69, 0x61, 0x66, 0x61, 0x62, 0x72, 0x69,
	0x63, 0x73, 0x22, 0xe5, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x61, 0x67, 0x6e, 0x6f, 0x6c, 0x69, 0x61, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x6f,
	0x6c, 0x6c, 0x52, 0x05, 0x52, 0x6f, 0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x45, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6e, 0x6f, 0x6c,
	0x69, 0x61, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x50, 0x4f, 0x52, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4f,
	0x6e, 0x50, 0x4f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x4f, 0x6e, 0x50, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x6e, 0x42, 0x61, 0x63,
	0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x32, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x2e, 0x0a, 0x02, 0x50, 0x4f,
	0x12, 0x10, 0x0a, 0x03, 0x45, 0x54, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45,
	0x54, 0x41, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x2d, 0x53, 0x6f, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x67, 0x6e, 0x6f, 0x6c, 0x69, 0x61,
	0x2d, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_inventory_proto_goTypes = []interface{}{
	(*Inventory)(nil),      // 0: magnoliafabrics.Inventory
	(*Inventory_Roll)(nil), // 1: magnoliafabrics.Inventory.Roll
	(*Inventory_PO)(nil),   // 2: magnoliafabrics.Inventory.PO
}
var file_inventory_proto_depIdxs = []int32{
	1, // 0: magnoliafabrics.Inventory.Rolls:type_name -> magnoliafabrics.Inventory.Roll
	2, // 1: magnoliafabrics.Inventory.PurchaseOrders:type_name -> magnoliafabrics.Inventory.PO
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory_Roll); i {
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
		file_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory_PO); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}
