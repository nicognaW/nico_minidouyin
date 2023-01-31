// idl/api.proto; 注解拓展

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_api_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50101,
		Name:          "api.raw_body",
		Tag:           "bytes,50101,opt,name=raw_body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50102,
		Name:          "api.query",
		Tag:           "bytes,50102,opt,name=query",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50103,
		Name:          "api.header",
		Tag:           "bytes,50103,opt,name=header",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50104,
		Name:          "api.cookie",
		Tag:           "bytes,50104,opt,name=cookie",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50105,
		Name:          "api.body",
		Tag:           "bytes,50105,opt,name=body",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50106,
		Name:          "api.path",
		Tag:           "bytes,50106,opt,name=path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50107,
		Name:          "api.vd",
		Tag:           "bytes,50107,opt,name=vd",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50108,
		Name:          "api.form",
		Tag:           "bytes,50108,opt,name=form",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51001,
		Name:          "api.go_tag",
		Tag:           "bytes,51001,opt,name=go_tag",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50109,
		Name:          "api.js_conv",
		Tag:           "bytes,50109,opt,name=js_conv",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50201,
		Name:          "api.get",
		Tag:           "bytes,50201,opt,name=get",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50202,
		Name:          "api.post",
		Tag:           "bytes,50202,opt,name=post",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50203,
		Name:          "api.put",
		Tag:           "bytes,50203,opt,name=put",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50204,
		Name:          "api.delete",
		Tag:           "bytes,50204,opt,name=delete",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50205,
		Name:          "api.patch",
		Tag:           "bytes,50205,opt,name=patch",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50206,
		Name:          "api.options",
		Tag:           "bytes,50206,opt,name=options",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50207,
		Name:          "api.head",
		Tag:           "bytes,50207,opt,name=head",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50208,
		Name:          "api.any",
		Tag:           "bytes,50208,opt,name=any",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50301,
		Name:          "api.gen_path",
		Tag:           "bytes,50301,opt,name=gen_path",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50302,
		Name:          "api.api_version",
		Tag:           "bytes,50302,opt,name=api_version",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50303,
		Name:          "api.tag",
		Tag:           "bytes,50303,opt,name=tag",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50304,
		Name:          "api.name",
		Tag:           "bytes,50304,opt,name=name",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50305,
		Name:          "api.api_level",
		Tag:           "bytes,50305,opt,name=api_level",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50306,
		Name:          "api.serializer",
		Tag:           "bytes,50306,opt,name=serializer",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50307,
		Name:          "api.param",
		Tag:           "bytes,50307,opt,name=param",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50308,
		Name:          "api.baseurl",
		Tag:           "bytes,50308,opt,name=baseurl",
		Filename:      "api.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50401,
		Name:          "api.http_code",
		Tag:           "varint,50401,opt,name=http_code",
		Filename:      "api.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string raw_body = 50101;
	E_RawBody = &file_api_proto_extTypes[0]
	// optional string query = 50102;
	E_Query = &file_api_proto_extTypes[1]
	// optional string header = 50103;
	E_Header = &file_api_proto_extTypes[2]
	// optional string cookie = 50104;
	E_Cookie = &file_api_proto_extTypes[3]
	// optional string body = 50105;
	E_Body = &file_api_proto_extTypes[4]
	// optional string path = 50106;
	E_Path = &file_api_proto_extTypes[5]
	// optional string vd = 50107;
	E_Vd = &file_api_proto_extTypes[6]
	// optional string form = 50108;
	E_Form = &file_api_proto_extTypes[7]
	// optional string go_tag = 51001;
	E_GoTag = &file_api_proto_extTypes[8]
	// optional string js_conv = 50109;
	E_JsConv = &file_api_proto_extTypes[9]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string get = 50201;
	E_Get = &file_api_proto_extTypes[10]
	// optional string post = 50202;
	E_Post = &file_api_proto_extTypes[11]
	// optional string put = 50203;
	E_Put = &file_api_proto_extTypes[12]
	// optional string delete = 50204;
	E_Delete = &file_api_proto_extTypes[13]
	// optional string patch = 50205;
	E_Patch = &file_api_proto_extTypes[14]
	// optional string options = 50206;
	E_Options = &file_api_proto_extTypes[15]
	// optional string head = 50207;
	E_Head = &file_api_proto_extTypes[16]
	// optional string any = 50208;
	E_Any = &file_api_proto_extTypes[17]
	// optional string gen_path = 50301;
	E_GenPath = &file_api_proto_extTypes[18]
	// optional string api_version = 50302;
	E_ApiVersion = &file_api_proto_extTypes[19]
	// optional string tag = 50303;
	E_Tag = &file_api_proto_extTypes[20]
	// optional string name = 50304;
	E_Name = &file_api_proto_extTypes[21]
	// optional string api_level = 50305;
	E_ApiLevel = &file_api_proto_extTypes[22]
	// optional string serializer = 50306;
	E_Serializer = &file_api_proto_extTypes[23]
	// optional string param = 50307;
	E_Param = &file_api_proto_extTypes[24]
	// optional string baseurl = 50308;
	E_Baseurl = &file_api_proto_extTypes[25]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional int32 http_code = 50401;
	E_HttpCode = &file_api_proto_extTypes[26]
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x3a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x87,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x42, 0x6f, 0x64, 0x79, 0x3a, 0x35,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb6, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x3a, 0x37,
	0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x3a, 0x33, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9,
	0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x33, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xba, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x3a, 0x2f, 0x0a, 0x02, 0x76, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbb, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x76, 0x64, 0x3a, 0x33, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbc, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x36, 0x0a, 0x06, 0x67, 0x6f, 0x5f, 0x74, 0x61,
	0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb9, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6f, 0x54, 0x61, 0x67, 0x3a,
	0x38, 0x0a, 0x07, 0x6a, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0x87, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6a, 0x73, 0x43, 0x6f, 0x6e, 0x76, 0x3a, 0x32, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x99, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x74, 0x3a, 0x34, 0x0a,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9a, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x3a, 0x32, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9b, 0x88, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x75, 0x74, 0x3a, 0x38, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x9c, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x3a, 0x36, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9d, 0x88, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x3a, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9e, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x34, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9f, 0x88,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x3a, 0x32, 0x0a, 0x03, 0x61,
	0x6e, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa0, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x3a,
	0x3b, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0x88, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x3a, 0x41, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0x88, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a,
	0x32, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff, 0x88, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x3a, 0x34, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x80, 0x89, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x3d, 0x0a, 0x09, 0x61, 0x70, 0x69,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x81, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x70, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x40, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x82, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x3a, 0x36, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x83, 0x89, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x3a, 0x3a, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0x89,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x75, 0x72, 0x6c, 0x3a, 0x40,
	0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe1,
	0x89, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x42, 0x25, 0x5a, 0x23, 0x6e, 0x69, 0x63, 0x6f, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x77, 0x65, 0x62, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
}

var file_api_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil),     // 0: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),    // 1: google.protobuf.MethodOptions
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.raw_body:extendee -> google.protobuf.FieldOptions
	0,  // 1: api.query:extendee -> google.protobuf.FieldOptions
	0,  // 2: api.header:extendee -> google.protobuf.FieldOptions
	0,  // 3: api.cookie:extendee -> google.protobuf.FieldOptions
	0,  // 4: api.body:extendee -> google.protobuf.FieldOptions
	0,  // 5: api.path:extendee -> google.protobuf.FieldOptions
	0,  // 6: api.vd:extendee -> google.protobuf.FieldOptions
	0,  // 7: api.form:extendee -> google.protobuf.FieldOptions
	0,  // 8: api.go_tag:extendee -> google.protobuf.FieldOptions
	0,  // 9: api.js_conv:extendee -> google.protobuf.FieldOptions
	1,  // 10: api.get:extendee -> google.protobuf.MethodOptions
	1,  // 11: api.post:extendee -> google.protobuf.MethodOptions
	1,  // 12: api.put:extendee -> google.protobuf.MethodOptions
	1,  // 13: api.delete:extendee -> google.protobuf.MethodOptions
	1,  // 14: api.patch:extendee -> google.protobuf.MethodOptions
	1,  // 15: api.options:extendee -> google.protobuf.MethodOptions
	1,  // 16: api.head:extendee -> google.protobuf.MethodOptions
	1,  // 17: api.any:extendee -> google.protobuf.MethodOptions
	1,  // 18: api.gen_path:extendee -> google.protobuf.MethodOptions
	1,  // 19: api.api_version:extendee -> google.protobuf.MethodOptions
	1,  // 20: api.tag:extendee -> google.protobuf.MethodOptions
	1,  // 21: api.name:extendee -> google.protobuf.MethodOptions
	1,  // 22: api.api_level:extendee -> google.protobuf.MethodOptions
	1,  // 23: api.serializer:extendee -> google.protobuf.MethodOptions
	1,  // 24: api.param:extendee -> google.protobuf.MethodOptions
	1,  // 25: api.baseurl:extendee -> google.protobuf.MethodOptions
	2,  // 26: api.http_code:extendee -> google.protobuf.EnumValueOptions
	27, // [27:27] is the sub-list for method output_type
	27, // [27:27] is the sub-list for method input_type
	27, // [27:27] is the sub-list for extension type_name
	0,  // [0:27] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 27,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		ExtensionInfos:    file_api_proto_extTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
