// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: storage/gcs.proto

package storage

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

// Configuration for a Google Cloud Storage persistent store
type Gcs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// A path to a JSON service account with permission to read and write to the bucket to be used as a backing store.
	JsonPath string `protobuf:"bytes,2,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// The Google Cloud Platform project you are using to host the GCS bucket as a backing store.
	Project string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	// The name of a storage bucket that your specified account has access to.
	Bucket string `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// The root folder in the chosen bucket to place all of Spinnaker's persistent data in.
	RootFolder string `protobuf:"bytes,5,opt,name=rootFolder,proto3" json:"rootFolder,omitempty"`
	// This is only required if the bucket you specify does not exist yet.
	BucketLocation string `protobuf:"bytes,6,opt,name=bucketLocation,proto3" json:"bucketLocation,omitempty"`
}

func (x *Gcs) Reset() {
	*x = Gcs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_gcs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gcs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gcs) ProtoMessage() {}

func (x *Gcs) ProtoReflect() protoreflect.Message {
	mi := &file_storage_gcs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gcs.ProtoReflect.Descriptor instead.
func (*Gcs) Descriptor() ([]byte, []int) {
	return file_storage_gcs_proto_rawDescGZIP(), []int{0}
}

func (x *Gcs) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Gcs) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *Gcs) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Gcs) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Gcs) GetRootFolder() string {
	if x != nil {
		return x.RootFolder
	}
	return ""
}

func (x *Gcs) GetBucketLocation() string {
	if x != nil {
		return x.BucketLocation
	}
	return ""
}

var File_storage_gcs_proto protoreflect.FileDescriptor

var file_storage_gcs_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x03, 0x47, 0x63, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b,
	0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_storage_gcs_proto_rawDescOnce sync.Once
	file_storage_gcs_proto_rawDescData = file_storage_gcs_proto_rawDesc
)

func file_storage_gcs_proto_rawDescGZIP() []byte {
	file_storage_gcs_proto_rawDescOnce.Do(func() {
		file_storage_gcs_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_gcs_proto_rawDescData)
	})
	return file_storage_gcs_proto_rawDescData
}

var file_storage_gcs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_gcs_proto_goTypes = []interface{}{
	(*Gcs)(nil), // 0: proto.storage.Gcs
}
var file_storage_gcs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_gcs_proto_init() }
func file_storage_gcs_proto_init() {
	if File_storage_gcs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_gcs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gcs); i {
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
			RawDescriptor: file_storage_gcs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_gcs_proto_goTypes,
		DependencyIndexes: file_storage_gcs_proto_depIdxs,
		MessageInfos:      file_storage_gcs_proto_msgTypes,
	}.Build()
	File_storage_gcs_proto = out.File
	file_storage_gcs_proto_rawDesc = nil
	file_storage_gcs_proto_goTypes = nil
	file_storage_gcs_proto_depIdxs = nil
}
