// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabric.proto

package fabric

import (
	fmt "fmt"
	inventory "github.com/Vitality-South/magnolia-fabrics-api/pkg/inventory"
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

type Fabric struct {
	// @gotags: dynamo:"FabricSKU"
	Sku string `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty" dynamo:"FabricSKU"`
	// @gotags: dynamo:"FabricProductCode"
	ProductCode string `protobuf:"bytes,2,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty" dynamo:"FabricProductCode"`
	// @gotags: dynamo:"FabricPattern"
	Pattern string `protobuf:"bytes,3,opt,name=pattern,proto3" json:"pattern,omitempty" dynamo:"FabricPattern"`
	// @gotags: dynamo:"FabricPrimaryColor"
	Color string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty" dynamo:"FabricPrimaryColor"`
	// @gotags: dynamo:"FabricPatternColorCombo"
	PatternColorCombo string `protobuf:"bytes,5,opt,name=pattern_color_combo,json=patternColorCombo,proto3" json:"pattern_color_combo,omitempty" dynamo:"FabricPatternColorCombo"`
	// @gotags: dynamo:"FabricStatus"
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty" dynamo:"FabricStatus"`
	// @gotags: dynamo:"FabricIntroDate"
	IntroDate string `protobuf:"bytes,7,opt,name=intro_date,json=introDate,proto3" json:"intro_date,omitempty" dynamo:"FabricIntroDate"`
	// @gotags: dynamo:"FabricContents"
	Contents string `protobuf:"bytes,8,opt,name=contents,proto3" json:"contents,omitempty" dynamo:"FabricContents"`
	// @gotags: dynamo:"FabricWidth"
	Width string `protobuf:"bytes,9,opt,name=width,proto3" json:"width,omitempty" dynamo:"FabricWidth"`
	// @gotags: dynamo:"FabricDirection"
	Direction string `protobuf:"bytes,10,opt,name=direction,proto3" json:"direction,omitempty" dynamo:"FabricDirection"`
	// @gotags: dynamo:"FabricHRepeat"
	HRepeat string `protobuf:"bytes,11,opt,name=h_repeat,json=hRepeat,proto3" json:"h_repeat,omitempty" dynamo:"FabricHRepeat"`
	// @gotags: dynamo:"FabricVRepeat"
	VRepeat string `protobuf:"bytes,12,opt,name=v_repeat,json=vRepeat,proto3" json:"v_repeat,omitempty" dynamo:"FabricVRepeat"`
	// @gotags: dynamo:"FabricCleaningCode"
	CleaningCode string `protobuf:"bytes,13,opt,name=cleaning_code,json=cleaningCode,proto3" json:"cleaning_code,omitempty" dynamo:"FabricCleaningCode"`
	// @gotags: dynamo:"FabricDoubleRubs"
	DoubleRubs string `protobuf:"bytes,14,opt,name=double_rubs,json=doubleRubs,proto3" json:"double_rubs,omitempty" dynamo:"FabricDoubleRubs"`
	// @gotags: dynamo:"FabricMisc"
	Misc string `protobuf:"bytes,15,opt,name=misc,proto3" json:"misc,omitempty" dynamo:"FabricMisc"`
	// @gotags: dynamo:"FabricDisclaimer"
	Disclaimer string `protobuf:"bytes,16,opt,name=disclaimer,proto3" json:"disclaimer,omitempty" dynamo:"FabricDisclaimer"`
	// @gotags: dynamo:"FabricUses"
	Uses []string `protobuf:"bytes,17,rep,name=uses,proto3" json:"uses,omitempty" dynamo:"FabricUses"`
	// @gotags: dynamo:"FabricDesigns"
	Designs []string `protobuf:"bytes,18,rep,name=designs,proto3" json:"designs,omitempty" dynamo:"FabricDesigns"`
	// @gotags: dynamo:"FabricColors"
	Colors []string `protobuf:"bytes,19,rep,name=colors,proto3" json:"colors,omitempty" dynamo:"FabricColors"`
	// @gotags: dynamo:"FabricOrigin"
	Origin string `protobuf:"bytes,20,opt,name=origin,proto3" json:"origin,omitempty" dynamo:"FabricOrigin"`
	// @gotags: dynamo:"FabricPillingGrade"
	PillingGrade string `protobuf:"bytes,21,opt,name=pilling_grade,json=pillingGrade,proto3" json:"pilling_grade,omitempty" dynamo:"FabricPillingGrade"`
	// @gotags: dynamo:"FabricFireCode"
	FireCode string `protobuf:"bytes,22,opt,name=fire_code,json=fireCode,proto3" json:"fire_code,omitempty" dynamo:"FabricFireCode"`
	// @gotags: dynamo:"FabricCategories"
	Categories []string `protobuf:"bytes,23,rep,name=categories,proto3" json:"categories,omitempty" dynamo:"FabricCategories"`
	// @gotags: dynamo:"FabricBrand"
	Brand string `protobuf:"bytes,24,opt,name=brand,proto3" json:"brand,omitempty" dynamo:"FabricBrand"`
	// @gotags: dynamo:"FabricIsDropShipped"
	IsDropShipped bool `protobuf:"varint,25,opt,name=is_drop_shipped,json=isDropShipped,proto3" json:"is_drop_shipped,omitempty" dynamo:"FabricIsDropShipped"`
	// @gotags: dynamo:"FabricMainImage"
	Image string `protobuf:"bytes,26,opt,name=image,proto3" json:"image,omitempty" dynamo:"FabricMainImage"`
	// @gotags: dynamo:"FabricAPIDisplayPrice"
	DisplayPrice string `protobuf:"bytes,27,opt,name=display_price,json=displayPrice,proto3" json:"display_price,omitempty" dynamo:"FabricAPIDisplayPrice"`
	// @gotags: dynamo:"FabricAPIPrice"
	Price int32 `protobuf:"varint,28,opt,name=price,proto3" json:"price,omitempty" dynamo:"FabricAPIPrice"`
	// @gotags: dynamo:"FabricOtherImages"
	SupplementalImages []string `protobuf:"bytes,29,rep,name=supplemental_images,json=supplementalImages,proto3" json:"supplemental_images,omitempty" dynamo:"FabricOtherImages"`
	// @gotags: dynamo:"FabricFullWidthImages"
	FullWidthImages []string `protobuf:"bytes,30,rep,name=full_width_images,json=fullWidthImages,proto3" json:"full_width_images,omitempty" dynamo:"FabricFullWidthImages"`
	// @gotags: dynamo:"-"
	Inventory            *inventory.Inventory `protobuf:"bytes,31,opt,name=inventory,proto3" json:"inventory,omitempty" dynamo:"-"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Fabric) Reset()         { *m = Fabric{} }
func (m *Fabric) String() string { return proto.CompactTextString(m) }
func (*Fabric) ProtoMessage()    {}
func (*Fabric) Descriptor() ([]byte, []int) {
	return fileDescriptor_284efff686d8e9bf, []int{0}
}

func (m *Fabric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fabric.Unmarshal(m, b)
}
func (m *Fabric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fabric.Marshal(b, m, deterministic)
}
func (m *Fabric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fabric.Merge(m, src)
}
func (m *Fabric) XXX_Size() int {
	return xxx_messageInfo_Fabric.Size(m)
}
func (m *Fabric) XXX_DiscardUnknown() {
	xxx_messageInfo_Fabric.DiscardUnknown(m)
}

var xxx_messageInfo_Fabric proto.InternalMessageInfo

func (m *Fabric) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *Fabric) GetProductCode() string {
	if m != nil {
		return m.ProductCode
	}
	return ""
}

func (m *Fabric) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *Fabric) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Fabric) GetPatternColorCombo() string {
	if m != nil {
		return m.PatternColorCombo
	}
	return ""
}

func (m *Fabric) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Fabric) GetIntroDate() string {
	if m != nil {
		return m.IntroDate
	}
	return ""
}

func (m *Fabric) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *Fabric) GetWidth() string {
	if m != nil {
		return m.Width
	}
	return ""
}

func (m *Fabric) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *Fabric) GetHRepeat() string {
	if m != nil {
		return m.HRepeat
	}
	return ""
}

func (m *Fabric) GetVRepeat() string {
	if m != nil {
		return m.VRepeat
	}
	return ""
}

func (m *Fabric) GetCleaningCode() string {
	if m != nil {
		return m.CleaningCode
	}
	return ""
}

func (m *Fabric) GetDoubleRubs() string {
	if m != nil {
		return m.DoubleRubs
	}
	return ""
}

func (m *Fabric) GetMisc() string {
	if m != nil {
		return m.Misc
	}
	return ""
}

func (m *Fabric) GetDisclaimer() string {
	if m != nil {
		return m.Disclaimer
	}
	return ""
}

func (m *Fabric) GetUses() []string {
	if m != nil {
		return m.Uses
	}
	return nil
}

func (m *Fabric) GetDesigns() []string {
	if m != nil {
		return m.Designs
	}
	return nil
}

func (m *Fabric) GetColors() []string {
	if m != nil {
		return m.Colors
	}
	return nil
}

func (m *Fabric) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Fabric) GetPillingGrade() string {
	if m != nil {
		return m.PillingGrade
	}
	return ""
}

func (m *Fabric) GetFireCode() string {
	if m != nil {
		return m.FireCode
	}
	return ""
}

func (m *Fabric) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Fabric) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Fabric) GetIsDropShipped() bool {
	if m != nil {
		return m.IsDropShipped
	}
	return false
}

func (m *Fabric) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Fabric) GetDisplayPrice() string {
	if m != nil {
		return m.DisplayPrice
	}
	return ""
}

func (m *Fabric) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Fabric) GetSupplementalImages() []string {
	if m != nil {
		return m.SupplementalImages
	}
	return nil
}

func (m *Fabric) GetFullWidthImages() []string {
	if m != nil {
		return m.FullWidthImages
	}
	return nil
}

func (m *Fabric) GetInventory() *inventory.Inventory {
	if m != nil {
		return m.Inventory
	}
	return nil
}

func init() {
	proto.RegisterType((*Fabric)(nil), "magnoliafabrics.Fabric")
}

func init() {
	proto.RegisterFile("fabric.proto", fileDescriptor_284efff686d8e9bf)
}

var fileDescriptor_284efff686d8e9bf = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x94, 0x5f, 0x4f, 0x1b, 0x3b,
	0x10, 0xc5, 0x95, 0x0b, 0x84, 0xc4, 0x09, 0x37, 0x60, 0xb8, 0xdc, 0x21, 0xfc, 0x4b, 0x5b, 0xa9,
	0x8a, 0x2a, 0x91, 0x48, 0xed, 0x4b, 0xab, 0xbe, 0x15, 0xd4, 0x8a, 0xb7, 0x2a, 0x48, 0xad, 0xd4,
	0x97, 0x95, 0x77, 0x6d, 0x36, 0x23, 0x76, 0x6d, 0xcb, 0xf6, 0x52, 0xf1, 0xd5, 0xfa, 0xe9, 0x2a,
	0x8f, 0x77, 0x53, 0xd4, 0x37, 0x9f, 0xf3, 0x3b, 0x63, 0x66, 0x3c, 0x6c, 0xd8, 0xf8, 0x5e, 0xe4,
	0x0e, 0x8b, 0x85, 0x75, 0x26, 0x18, 0x3e, 0xa9, 0x45, 0xa9, 0x4d, 0x85, 0x22, 0xb9, 0x7e, 0x3a,
	0x41, 0xfd, 0xa8, 0x74, 0x30, 0xee, 0x29, 0x25, 0x5e, 0xfe, 0xda, 0x65, 0xfd, 0xcf, 0x04, 0xf9,
	0x3e, 0xdb, 0xf2, 0x0f, 0x0d, 0xf4, 0x66, 0xbd, 0xf9, 0x70, 0x15, 0x8f, 0xfc, 0x05, 0x1b, 0x5b,
	0x67, 0x64, 0x53, 0x84, 0xac, 0x30, 0x52, 0xc1, 0x3f, 0x84, 0x46, 0xad, 0x77, 0x6d, 0xa4, 0xe2,
	0xc0, 0x76, 0xad, 0x08, 0x41, 0x39, 0x0d, 0x5b, 0x44, 0x3b, 0xc9, 0x8f, 0xd8, 0x4e, 0x61, 0x2a,
	0xe3, 0x60, 0x9b, 0xfc, 0x24, 0xf8, 0x82, 0x1d, 0xb6, 0x81, 0x8c, 0x8c, 0xac, 0x30, 0x75, 0x6e,
	0x60, 0x87, 0x32, 0x07, 0x2d, 0xba, 0x8e, 0xe4, 0x3a, 0x02, 0x7e, 0xcc, 0xfa, 0x3e, 0x88, 0xd0,
	0x78, 0xe8, 0x53, 0xa4, 0x55, 0xfc, 0x9c, 0x31, 0xd4, 0xc1, 0x99, 0x4c, 0x8a, 0xa0, 0x60, 0x97,
	0xd8, 0x90, 0x9c, 0x1b, 0x11, 0x14, 0x9f, 0xb2, 0x41, 0x61, 0x74, 0x50, 0x3a, 0x78, 0x18, 0x10,
	0xdc, 0xe8, 0xd8, 0xd8, 0x4f, 0x94, 0x61, 0x0d, 0xc3, 0xd4, 0x18, 0x09, 0x7e, 0xc6, 0x86, 0x12,
	0x9d, 0x2a, 0x02, 0x1a, 0x0d, 0x2c, 0xdd, 0xb7, 0x31, 0xf8, 0x09, 0x1b, 0xac, 0x33, 0xa7, 0xac,
	0x12, 0x01, 0x46, 0x69, 0xce, 0xf5, 0x8a, 0x64, 0x44, 0x8f, 0x1d, 0x1a, 0x27, 0xf4, 0xd8, 0xa2,
	0x57, 0x6c, 0xaf, 0xa8, 0x94, 0xd0, 0xa8, 0xcb, 0xf4, 0x80, 0x7b, 0xc4, 0xc7, 0x9d, 0x49, 0x2f,
	0x78, 0xc9, 0x46, 0xd2, 0x34, 0x79, 0xa5, 0x32, 0xd7, 0xe4, 0x1e, 0xfe, 0xa5, 0x08, 0x4b, 0xd6,
	0xaa, 0xc9, 0x3d, 0xe7, 0x6c, 0xbb, 0x46, 0x5f, 0xc0, 0x84, 0x08, 0x9d, 0xf9, 0x05, 0x63, 0x12,
	0x7d, 0x51, 0x09, 0xac, 0x95, 0x83, 0xfd, 0xb6, 0x66, 0xe3, 0xc4, 0x9a, 0xc6, 0x2b, 0x0f, 0x07,
	0xb3, 0xad, 0x58, 0x13, 0xcf, 0x71, 0x55, 0x52, 0x79, 0x2c, 0xb5, 0x07, 0x4e, 0x76, 0x27, 0xe3,
	0x23, 0xd3, 0x32, 0x3c, 0x1c, 0x12, 0x68, 0x55, 0xf4, 0x8d, 0xc3, 0x12, 0x35, 0x1c, 0xa5, 0xc7,
	0x4f, 0x2a, 0xce, 0x65, 0xb1, 0xaa, 0xe2, 0x58, 0xa5, 0x13, 0x52, 0xc1, 0x7f, 0x69, 0xae, 0xd6,
	0xfc, 0x12, 0x3d, 0x7e, 0xca, 0x86, 0xf7, 0xe8, 0x54, 0x1a, 0xfc, 0x38, 0xed, 0x20, 0x1a, 0x34,
	0xf4, 0x05, 0x63, 0x85, 0x08, 0xaa, 0x34, 0x0e, 0x95, 0x87, 0xff, 0xe9, 0xaf, 0x3e, 0x73, 0xe2,
	0x8e, 0x72, 0x27, 0xb4, 0x04, 0x48, 0x3b, 0x22, 0xc1, 0x5f, 0xb3, 0x09, 0xfa, 0x4c, 0x3a, 0x63,
	0x33, 0xbf, 0x46, 0x6b, 0x95, 0x84, 0x93, 0x59, 0x6f, 0x3e, 0x58, 0xed, 0xa1, 0xbf, 0x71, 0xc6,
	0xde, 0x25, 0x33, 0x56, 0x63, 0x2d, 0x4a, 0x05, 0xd3, 0x54, 0x4d, 0x22, 0x76, 0x2d, 0xd1, 0xdb,
	0x4a, 0x3c, 0x65, 0xd6, 0x61, 0xa1, 0xe0, 0x34, 0x75, 0xdd, 0x9a, 0x5f, 0xa3, 0x17, 0x4b, 0x13,
	0x3c, 0x9b, 0xf5, 0xe6, 0x3b, 0xab, 0x24, 0xf8, 0x92, 0x1d, 0xfa, 0xc6, 0xda, 0x4a, 0xd5, 0x4a,
	0x07, 0x51, 0x65, 0x74, 0xa1, 0x87, 0x73, 0xea, 0x9b, 0x3f, 0x47, 0xb7, 0x44, 0xf8, 0x1b, 0x76,
	0x70, 0xdf, 0x54, 0x55, 0x46, 0xff, 0x5b, 0x5d, 0xfc, 0x82, 0xe2, 0x93, 0x08, 0xbe, 0x47, 0xbf,
	0xcd, 0xbe, 0x67, 0xc3, 0xcd, 0x57, 0x09, 0x97, 0xb3, 0xde, 0x7c, 0xf4, 0x76, 0xba, 0xf8, 0xeb,
	0xc3, 0x5d, 0xdc, 0x76, 0x89, 0xd5, 0x9f, 0xf0, 0xa7, 0x8f, 0x3f, 0x3e, 0x94, 0x18, 0xd6, 0x4d,
	0xbe, 0x28, 0x4c, 0xbd, 0xfc, 0x86, 0x41, 0x54, 0x18, 0x9e, 0xae, 0xee, 0x4c, 0x13, 0xd6, 0xcb,
	0xee, 0x86, 0xab, 0xf6, 0x8a, 0x2b, 0x61, 0x71, 0x69, 0x1f, 0xca, 0x65, 0xd2, 0x79, 0x9f, 0x7e,
	0x00, 0xde, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x82, 0x9e, 0xf2, 0xc5, 0x32, 0x04, 0x00, 0x00,
}
