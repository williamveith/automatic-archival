// protoc --go_out=. protobuf/schema.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.28.2
// source: protobuf/schema.proto

package database

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

type LineItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Time          string                 `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Prof          string                 `protobuf:"bytes,3,opt,name=prof,proto3" json:"prof,omitempty"`
	User          string                 `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Tool          string                 `protobuf:"bytes,5,opt,name=tool,proto3" json:"tool,omitempty"`
	Usage         float64                `protobuf:"fixed64,6,opt,name=usage,proto3" json:"usage,omitempty"`
	Rate          float64                `protobuf:"fixed64,7,opt,name=rate,proto3" json:"rate,omitempty"`
	Cost          float64                `protobuf:"fixed64,8,opt,name=cost,proto3" json:"cost,omitempty"`
	AppliedCost   float64                `protobuf:"fixed64,9,opt,name=applied_cost,json=appliedCost,proto3" json:"applied_cost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	mi := &file_protobuf_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_schema_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_protobuf_schema_proto_rawDescGZIP(), []int{0}
}

func (x *LineItem) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *LineItem) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *LineItem) GetProf() string {
	if x != nil {
		return x.Prof
	}
	return ""
}

func (x *LineItem) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *LineItem) GetTool() string {
	if x != nil {
		return x.Tool
	}
	return ""
}

func (x *LineItem) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *LineItem) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *LineItem) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *LineItem) GetAppliedCost() float64 {
	if x != nil {
		return x.AppliedCost
	}
	return 0
}

type Invoice struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Period        string                 `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	Group         string                 `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Usage         float64                `protobuf:"fixed64,3,opt,name=usage,proto3" json:"usage,omitempty"`
	Cost          float64                `protobuf:"fixed64,4,opt,name=cost,proto3" json:"cost,omitempty"`
	AppliedCost   float64                `protobuf:"fixed64,5,opt,name=applied_cost,json=appliedCost,proto3" json:"applied_cost,omitempty"`
	Surcharge     float64                `protobuf:"fixed64,6,opt,name=surcharge,proto3" json:"surcharge,omitempty"`
	Tax           float64                `protobuf:"fixed64,7,opt,name=tax,proto3" json:"tax,omitempty"`
	Lineitems     []*LineItem            `protobuf:"bytes,8,rep,name=lineitems,proto3" json:"lineitems,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	mi := &file_protobuf_schema_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_schema_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_protobuf_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Invoice) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Invoice) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Invoice) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *Invoice) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Invoice) GetAppliedCost() float64 {
	if x != nil {
		return x.AppliedCost
	}
	return 0
}

func (x *Invoice) GetSurcharge() float64 {
	if x != nil {
		return x.Surcharge
	}
	return 0
}

func (x *Invoice) GetTax() float64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Invoice) GetLineitems() []*LineItem {
	if x != nil {
		return x.Lineitems
	}
	return nil
}

type Year struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Period        string                 `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	Usage         float64                `protobuf:"fixed64,2,opt,name=usage,proto3" json:"usage,omitempty"`
	Cost          float64                `protobuf:"fixed64,3,opt,name=cost,proto3" json:"cost,omitempty"`
	AppliedCost   float64                `protobuf:"fixed64,4,opt,name=applied_cost,json=appliedCost,proto3" json:"applied_cost,omitempty"`
	Surcharge     float64                `protobuf:"fixed64,5,opt,name=surcharge,proto3" json:"surcharge,omitempty"`
	Tax           float64                `protobuf:"fixed64,6,opt,name=tax,proto3" json:"tax,omitempty"`
	Invoices      []*Invoice             `protobuf:"bytes,7,rep,name=invoices,proto3" json:"invoices,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Year) Reset() {
	*x = Year{}
	mi := &file_protobuf_schema_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Year) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Year) ProtoMessage() {}

func (x *Year) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_schema_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Year.ProtoReflect.Descriptor instead.
func (*Year) Descriptor() ([]byte, []int) {
	return file_protobuf_schema_proto_rawDescGZIP(), []int{2}
}

func (x *Year) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Year) GetUsage() float64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *Year) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Year) GetAppliedCost() float64 {
	if x != nil {
		return x.AppliedCost
	}
	return 0
}

func (x *Year) GetSurcharge() float64 {
	if x != nil {
		return x.Surcharge
	}
	return 0
}

func (x *Year) GetTax() float64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Year) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

var File_protobuf_schema_proto protoreflect.FileDescriptor

var file_protobuf_schema_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x6f, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x72, 0x6f, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f,
	0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x43,
	0x6f, 0x73, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x30, 0x0a, 0x09, 0x6c, 0x69,
	0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xca, 0x01, 0x0a,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x2d, 0x0a, 0x08, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_schema_proto_rawDescOnce sync.Once
	file_protobuf_schema_proto_rawDescData = file_protobuf_schema_proto_rawDesc
)

func file_protobuf_schema_proto_rawDescGZIP() []byte {
	file_protobuf_schema_proto_rawDescOnce.Do(func() {
		file_protobuf_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_schema_proto_rawDescData)
	})
	return file_protobuf_schema_proto_rawDescData
}

var file_protobuf_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_schema_proto_goTypes = []any{
	(*LineItem)(nil), // 0: database.LineItem
	(*Invoice)(nil),  // 1: database.Invoice
	(*Year)(nil),     // 2: database.Year
}
var file_protobuf_schema_proto_depIdxs = []int32{
	0, // 0: database.Invoice.lineitems:type_name -> database.LineItem
	1, // 1: database.Year.invoices:type_name -> database.Invoice
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_schema_proto_init() }
func file_protobuf_schema_proto_init() {
	if File_protobuf_schema_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_schema_proto_goTypes,
		DependencyIndexes: file_protobuf_schema_proto_depIdxs,
		MessageInfos:      file_protobuf_schema_proto_msgTypes,
	}.Build()
	File_protobuf_schema_proto = out.File
	file_protobuf_schema_proto_rawDesc = nil
	file_protobuf_schema_proto_goTypes = nil
	file_protobuf_schema_proto_depIdxs = nil
}
