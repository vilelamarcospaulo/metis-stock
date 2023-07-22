// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: stock_picker.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStockFluctuationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ticker string                 `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Begin  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=begin,proto3" json:"begin,omitempty"`
	End    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetStockFluctuationRequest) Reset() {
	*x = GetStockFluctuationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_picker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockFluctuationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockFluctuationRequest) ProtoMessage() {}

func (x *GetStockFluctuationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_picker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockFluctuationRequest.ProtoReflect.Descriptor instead.
func (*GetStockFluctuationRequest) Descriptor() ([]byte, []int) {
	return file_stock_picker_proto_rawDescGZIP(), []int{0}
}

func (x *GetStockFluctuationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStockFluctuationRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *GetStockFluctuationRequest) GetBegin() *timestamppb.Timestamp {
	if x != nil {
		return x.Begin
	}
	return nil
}

func (x *GetStockFluctuationRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type GetStockFluctuationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before float64 `protobuf:"fixed64,1,opt,name=before,proto3" json:"before,omitempty"`
	After  float64 `protobuf:"fixed64,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *GetStockFluctuationResponse) Reset() {
	*x = GetStockFluctuationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_picker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStockFluctuationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStockFluctuationResponse) ProtoMessage() {}

func (x *GetStockFluctuationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_picker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStockFluctuationResponse.ProtoReflect.Descriptor instead.
func (*GetStockFluctuationResponse) Descriptor() ([]byte, []int) {
	return file_stock_picker_proto_rawDescGZIP(), []int{1}
}

func (x *GetStockFluctuationResponse) GetBefore() float64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *GetStockFluctuationResponse) GetAfter() float64 {
	if x != nil {
		return x.After
	}
	return 0
}

var File_stock_picker_proto protoreflect.FileDescriptor

var file_stock_picker_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x46, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x4b,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x46, 0x6c, 0x75, 0x63, 0x74, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x32, 0x87, 0x01, 0x0a, 0x12,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x71, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x46, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x46, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x5f, 0x70, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x46, 0x6c, 0x75, 0x63, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x70,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_picker_proto_rawDescOnce sync.Once
	file_stock_picker_proto_rawDescData = file_stock_picker_proto_rawDesc
)

func file_stock_picker_proto_rawDescGZIP() []byte {
	file_stock_picker_proto_rawDescOnce.Do(func() {
		file_stock_picker_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_picker_proto_rawDescData)
	})
	return file_stock_picker_proto_rawDescData
}

var file_stock_picker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stock_picker_proto_goTypes = []interface{}{
	(*GetStockFluctuationRequest)(nil),  // 0: stock_picker.GetStockFluctuationRequest
	(*GetStockFluctuationResponse)(nil), // 1: stock_picker.GetStockFluctuationResponse
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
}
var file_stock_picker_proto_depIdxs = []int32{
	2, // 0: stock_picker.GetStockFluctuationRequest.begin:type_name -> google.protobuf.Timestamp
	2, // 1: stock_picker.GetStockFluctuationRequest.end:type_name -> google.protobuf.Timestamp
	0, // 2: stock_picker.StockPickerService.GetStockPriceFluctuation:input_type -> stock_picker.GetStockFluctuationRequest
	1, // 3: stock_picker.StockPickerService.GetStockPriceFluctuation:output_type -> stock_picker.GetStockFluctuationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stock_picker_proto_init() }
func file_stock_picker_proto_init() {
	if File_stock_picker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_picker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockFluctuationRequest); i {
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
		file_stock_picker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStockFluctuationResponse); i {
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
			RawDescriptor: file_stock_picker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_picker_proto_goTypes,
		DependencyIndexes: file_stock_picker_proto_depIdxs,
		MessageInfos:      file_stock_picker_proto_msgTypes,
	}.Build()
	File_stock_picker_proto = out.File
	file_stock_picker_proto_rawDesc = nil
	file_stock_picker_proto_goTypes = nil
	file_stock_picker_proto_depIdxs = nil
}