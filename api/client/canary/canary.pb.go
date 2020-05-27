// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: canary/canary.proto

package canary

import (
	proto "github.com/golang/protobuf/proto"
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

// Configuration for Spinnaker's canary service.
type Canary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the canary service is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Canary service integrations. You must configure at least one account of
	// each canary.SupportedType (METRICS_STORE, CONFIGURATION_STORE,
	// OBJECT_STORE) in order to use Spinnaker's canary service.
	ServiceIntegrations *Canary_ServiceIntegrations `protobuf:"bytes,2,opt,name=serviceIntegrations,proto3" json:"serviceIntegrations,omitempty"`
}

func (x *Canary) Reset() {
	*x = Canary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_canary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Canary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Canary) ProtoMessage() {}

func (x *Canary) ProtoReflect() protoreflect.Message {
	mi := &file_canary_canary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Canary.ProtoReflect.Descriptor instead.
func (*Canary) Descriptor() ([]byte, []int) {
	return file_canary_canary_proto_rawDescGZIP(), []int{0}
}

func (x *Canary) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Canary) GetServiceIntegrations() *Canary_ServiceIntegrations {
	if x != nil {
		return x.ServiceIntegrations
	}
	return nil
}

type Canary_ServiceIntegrations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aws        *Aws        `protobuf:"bytes,1,opt,name=aws,proto3" json:"aws,omitempty"`
	Datadog    *Datadog    `protobuf:"bytes,2,opt,name=datadog,proto3" json:"datadog,omitempty"`
	Google     *Google     `protobuf:"bytes,3,opt,name=google,proto3" json:"google,omitempty"`
	Newrelic   *NewRelic   `protobuf:"bytes,4,opt,name=newrelic,proto3" json:"newrelic,omitempty"`
	Prometheus *Prometheus `protobuf:"bytes,5,opt,name=prometheus,proto3" json:"prometheus,omitempty"`
	Signalfx   *SignalFx   `protobuf:"bytes,6,opt,name=signalfx,proto3" json:"signalfx,omitempty"`
}

func (x *Canary_ServiceIntegrations) Reset() {
	*x = Canary_ServiceIntegrations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_canary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Canary_ServiceIntegrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Canary_ServiceIntegrations) ProtoMessage() {}

func (x *Canary_ServiceIntegrations) ProtoReflect() protoreflect.Message {
	mi := &file_canary_canary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Canary_ServiceIntegrations.ProtoReflect.Descriptor instead.
func (*Canary_ServiceIntegrations) Descriptor() ([]byte, []int) {
	return file_canary_canary_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Canary_ServiceIntegrations) GetAws() *Aws {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *Canary_ServiceIntegrations) GetDatadog() *Datadog {
	if x != nil {
		return x.Datadog
	}
	return nil
}

func (x *Canary_ServiceIntegrations) GetGoogle() *Google {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *Canary_ServiceIntegrations) GetNewrelic() *NewRelic {
	if x != nil {
		return x.Newrelic
	}
	return nil
}

func (x *Canary_ServiceIntegrations) GetPrometheus() *Prometheus {
	if x != nil {
		return x.Prometheus
	}
	return nil
}

func (x *Canary_ServiceIntegrations) GetSignalfx() *SignalFx {
	if x != nil {
		return x.Signalfx
	}
	return nil
}

var File_canary_canary_proto protoreflect.FileDescriptor

var file_canary_canary_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x1a, 0x10, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x61, 0x77, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x66,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x03, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x5a, 0x0a, 0x13,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xbb, 0x02, 0x0a, 0x13, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x23, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x77, 0x73,
	0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x52, 0x07, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x12, 0x2c, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x06, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x08,
	0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65,
	0x75, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x66, 0x78, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x46, 0x78, 0x52, 0x08, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x66, 0x78, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b,
	0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_canary_canary_proto_rawDescOnce sync.Once
	file_canary_canary_proto_rawDescData = file_canary_canary_proto_rawDesc
)

func file_canary_canary_proto_rawDescGZIP() []byte {
	file_canary_canary_proto_rawDescOnce.Do(func() {
		file_canary_canary_proto_rawDescData = protoimpl.X.CompressGZIP(file_canary_canary_proto_rawDescData)
	})
	return file_canary_canary_proto_rawDescData
}

var file_canary_canary_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_canary_canary_proto_goTypes = []interface{}{
	(*Canary)(nil),                     // 0: proto.canary.Canary
	(*Canary_ServiceIntegrations)(nil), // 1: proto.canary.Canary.ServiceIntegrations
	(*Aws)(nil),                        // 2: proto.canary.Aws
	(*Datadog)(nil),                    // 3: proto.canary.Datadog
	(*Google)(nil),                     // 4: proto.canary.Google
	(*NewRelic)(nil),                   // 5: proto.canary.NewRelic
	(*Prometheus)(nil),                 // 6: proto.canary.Prometheus
	(*SignalFx)(nil),                   // 7: proto.canary.SignalFx
}
var file_canary_canary_proto_depIdxs = []int32{
	1, // 0: proto.canary.Canary.serviceIntegrations:type_name -> proto.canary.Canary.ServiceIntegrations
	2, // 1: proto.canary.Canary.ServiceIntegrations.aws:type_name -> proto.canary.Aws
	3, // 2: proto.canary.Canary.ServiceIntegrations.datadog:type_name -> proto.canary.Datadog
	4, // 3: proto.canary.Canary.ServiceIntegrations.google:type_name -> proto.canary.Google
	5, // 4: proto.canary.Canary.ServiceIntegrations.newrelic:type_name -> proto.canary.NewRelic
	6, // 5: proto.canary.Canary.ServiceIntegrations.prometheus:type_name -> proto.canary.Prometheus
	7, // 6: proto.canary.Canary.ServiceIntegrations.signalfx:type_name -> proto.canary.SignalFx
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_canary_canary_proto_init() }
func file_canary_canary_proto_init() {
	if File_canary_canary_proto != nil {
		return
	}
	file_canary_aws_proto_init()
	file_canary_datadog_proto_init()
	file_canary_google_proto_init()
	file_canary_newrelic_proto_init()
	file_canary_prometheus_proto_init()
	file_canary_signalfx_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_canary_canary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Canary); i {
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
		file_canary_canary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Canary_ServiceIntegrations); i {
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
			RawDescriptor: file_canary_canary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_canary_canary_proto_goTypes,
		DependencyIndexes: file_canary_canary_proto_depIdxs,
		MessageInfos:      file_canary_canary_proto_msgTypes,
	}.Build()
	File_canary_canary_proto = out.File
	file_canary_canary_proto_rawDesc = nil
	file_canary_canary_proto_goTypes = nil
	file_canary_canary_proto_depIdxs = nil
}
