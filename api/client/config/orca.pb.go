// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: config/orca.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	security "github.com/spinnaker/kleat/api/client/security"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the Orca microservice.
type Orca struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineTemplates *Orca_PipelineTemplates `protobuf:"bytes,1,opt,name=pipelineTemplates,proto3" json:"pipelineTemplates,omitempty"`
	Webhook           *security.WebhookConfig `protobuf:"bytes,2,opt,name=webhook,proto3" json:"webhook,omitempty"`
	Default           *Orca_Defaults          `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	Services          *Orca_Services          `protobuf:"bytes,4,opt,name=services,proto3" json:"services,omitempty"`
}

func (x *Orca) Reset() {
	*x = Orca{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_orca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orca) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orca) ProtoMessage() {}

func (x *Orca) ProtoReflect() protoreflect.Message {
	mi := &file_config_orca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orca.ProtoReflect.Descriptor instead.
func (*Orca) Descriptor() ([]byte, []int) {
	return file_config_orca_proto_rawDescGZIP(), []int{0}
}

func (x *Orca) GetPipelineTemplates() *Orca_PipelineTemplates {
	if x != nil {
		return x.PipelineTemplates
	}
	return nil
}

func (x *Orca) GetWebhook() *security.WebhookConfig {
	if x != nil {
		return x.Webhook
	}
	return nil
}

func (x *Orca) GetDefault() *Orca_Defaults {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *Orca) GetServices() *Orca_Services {
	if x != nil {
		return x.Services
	}
	return nil
}

// Configuration for the status of non-core services.
type Orca_Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kayenta *ServiceSettings `protobuf:"bytes,1,opt,name=kayenta,proto3" json:"kayenta,omitempty"`
}

func (x *Orca_Services) Reset() {
	*x = Orca_Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_orca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orca_Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orca_Services) ProtoMessage() {}

func (x *Orca_Services) ProtoReflect() protoreflect.Message {
	mi := &file_config_orca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orca_Services.ProtoReflect.Descriptor instead.
func (*Orca_Services) Descriptor() ([]byte, []int) {
	return file_config_orca_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Orca_Services) GetKayenta() *ServiceSettings {
	if x != nil {
		return x.Kayenta
	}
	return nil
}

// Configuration for pipeline templates.
type Orca_PipelineTemplates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether pipeline templates are enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Orca_PipelineTemplates) Reset() {
	*x = Orca_PipelineTemplates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_orca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orca_PipelineTemplates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orca_PipelineTemplates) ProtoMessage() {}

func (x *Orca_PipelineTemplates) ProtoReflect() protoreflect.Message {
	mi := &file_config_orca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orca_PipelineTemplates.ProtoReflect.Descriptor instead.
func (*Orca_PipelineTemplates) Descriptor() ([]byte, []int) {
	return file_config_orca_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Orca_PipelineTemplates) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Defaults applicable to the orca microservice.
type Orca_Defaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration of bakery defaults.
	Bake *Orca_Defaults_BakeDefaults `protobuf:"bytes,1,opt,name=bake,proto3" json:"bake,omitempty"`
}

func (x *Orca_Defaults) Reset() {
	*x = Orca_Defaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_orca_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orca_Defaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orca_Defaults) ProtoMessage() {}

func (x *Orca_Defaults) ProtoReflect() protoreflect.Message {
	mi := &file_config_orca_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orca_Defaults.ProtoReflect.Descriptor instead.
func (*Orca_Defaults) Descriptor() ([]byte, []int) {
	return file_config_orca_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Orca_Defaults) GetBake() *Orca_Defaults_BakeDefaults {
	if x != nil {
		return x.Bake
	}
	return nil
}

// Configuration of bakery defaults.
type Orca_Defaults_BakeDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The default account to use for baking.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *Orca_Defaults_BakeDefaults) Reset() {
	*x = Orca_Defaults_BakeDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_orca_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orca_Defaults_BakeDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orca_Defaults_BakeDefaults) ProtoMessage() {}

func (x *Orca_Defaults_BakeDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_config_orca_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orca_Defaults_BakeDefaults.ProtoReflect.Descriptor instead.
func (*Orca_Defaults_BakeDefaults) Descriptor() ([]byte, []int) {
	return file_config_orca_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Orca_Defaults_BakeDefaults) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

var File_config_orca_proto protoreflect.FileDescriptor

var file_config_orca_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x03, 0x0a, 0x04,
	0x4f, 0x72, 0x63, 0x61, 0x12, 0x52, 0x0a, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f,
	0x72, 0x63, 0x61, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x12, 0x35, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f, 0x72, 0x63, 0x61, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x43, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a,
	0x07, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x07, 0x6b,
	0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x1a, 0x2d, 0x0a, 0x11, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0x72, 0x0a, 0x08, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4f,
	0x72, 0x63, 0x61, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x42, 0x61, 0x6b,
	0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x04, 0x62, 0x61, 0x6b, 0x65, 0x1a,
	0x28, 0x0a, 0x0c, 0x42, 0x61, 0x6b, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65,
	0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_orca_proto_rawDescOnce sync.Once
	file_config_orca_proto_rawDescData = file_config_orca_proto_rawDesc
)

func file_config_orca_proto_rawDescGZIP() []byte {
	file_config_orca_proto_rawDescOnce.Do(func() {
		file_config_orca_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_orca_proto_rawDescData)
	})
	return file_config_orca_proto_rawDescData
}

var file_config_orca_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_config_orca_proto_goTypes = []interface{}{
	(*Orca)(nil),                       // 0: proto.config.Orca
	(*Orca_Services)(nil),              // 1: proto.config.Orca.Services
	(*Orca_PipelineTemplates)(nil),     // 2: proto.config.Orca.PipelineTemplates
	(*Orca_Defaults)(nil),              // 3: proto.config.Orca.Defaults
	(*Orca_Defaults_BakeDefaults)(nil), // 4: proto.config.Orca.Defaults.BakeDefaults
	(*security.WebhookConfig)(nil),     // 5: proto.security.WebhookConfig
	(*ServiceSettings)(nil),            // 6: proto.config.ServiceSettings
}
var file_config_orca_proto_depIdxs = []int32{
	2, // 0: proto.config.Orca.pipelineTemplates:type_name -> proto.config.Orca.PipelineTemplates
	5, // 1: proto.config.Orca.webhook:type_name -> proto.security.WebhookConfig
	3, // 2: proto.config.Orca.default:type_name -> proto.config.Orca.Defaults
	1, // 3: proto.config.Orca.services:type_name -> proto.config.Orca.Services
	6, // 4: proto.config.Orca.Services.kayenta:type_name -> proto.config.ServiceSettings
	4, // 5: proto.config.Orca.Defaults.bake:type_name -> proto.config.Orca.Defaults.BakeDefaults
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_config_orca_proto_init() }
func file_config_orca_proto_init() {
	if File_config_orca_proto != nil {
		return
	}
	file_config_service_enabled_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_orca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orca); i {
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
		file_config_orca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orca_Services); i {
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
		file_config_orca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orca_PipelineTemplates); i {
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
		file_config_orca_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orca_Defaults); i {
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
		file_config_orca_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orca_Defaults_BakeDefaults); i {
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
			RawDescriptor: file_config_orca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_orca_proto_goTypes,
		DependencyIndexes: file_config_orca_proto_depIdxs,
		MessageInfos:      file_config_orca_proto_msgTypes,
	}.Build()
	File_config_orca_proto = out.File
	file_config_orca_proto_rawDesc = nil
	file_config_orca_proto_goTypes = nil
	file_config_orca_proto_depIdxs = nil
}
