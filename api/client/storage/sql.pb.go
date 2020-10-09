// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: storage/sql.proto

package storage

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type SQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Default database connection pool.
	Default *ConnectionPool `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	// Migration database connection pool.
	Migration *ConnectionPool `protobuf:"bytes,3,opt,name=migration,proto3" json:"migration,omitempty"`
}

func (x *SQL) Reset() {
	*x = SQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_sql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQL) ProtoMessage() {}

func (x *SQL) ProtoReflect() protoreflect.Message {
	mi := &file_storage_sql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQL.ProtoReflect.Descriptor instead.
func (*SQL) Descriptor() ([]byte, []int) {
	return file_storage_sql_proto_rawDescGZIP(), []int{0}
}

func (x *SQL) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *SQL) GetDefault() *ConnectionPool {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *SQL) GetMigration() *ConnectionPool {
	if x != nil {
		return x.Migration
	}
	return nil
}

// Default connection pool to Keel's datastore
type ConnectionPools struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default *ConnectionPool `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *ConnectionPools) Reset() {
	*x = ConnectionPools{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_sql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionPools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionPools) ProtoMessage() {}

func (x *ConnectionPools) ProtoReflect() protoreflect.Message {
	mi := &file_storage_sql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionPools.ProtoReflect.Descriptor instead.
func (*ConnectionPools) Descriptor() ([]byte, []int) {
	return file_storage_sql_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionPools) GetDefault() *ConnectionPool {
	if x != nil {
		return x.Default
	}
	return nil
}

// ConnectionPool confifugration for the SQL server
type ConnectionPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Database username
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Database password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Database connection string. This needs to include server port
	// and database name as well
	JdbcUrl string `protobuf:"bytes,3,opt,name=jdbcUrl,proto3" json:"jdbcUrl,omitempty"`
	// Database connection timeout in milliseconds
	ConnectionTimeout int32 `protobuf:"varint,4,opt,name=connectionTimeout,proto3" json:"connectionTimeout,omitempty"`
	// maxLifetime controls the maximum lifetime of a connection in
	// the pool in milliseconds.
	MaxLifetime int32 `protobuf:"varint,5,opt,name=maxLifetime,proto3" json:"maxLifetime,omitempty"`
	// Maximum number of connections stored in the connection pool
	MaxPoolSize int32 `protobuf:"varint,6,opt,name=maxPoolSize,proto3" json:"maxPoolSize,omitempty"`
}

func (x *ConnectionPool) Reset() {
	*x = ConnectionPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_sql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionPool) ProtoMessage() {}

func (x *ConnectionPool) ProtoReflect() protoreflect.Message {
	mi := &file_storage_sql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionPool.ProtoReflect.Descriptor instead.
func (*ConnectionPool) Descriptor() ([]byte, []int) {
	return file_storage_sql_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionPool) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ConnectionPool) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConnectionPool) GetJdbcUrl() string {
	if x != nil {
		return x.JdbcUrl
	}
	return ""
}

func (x *ConnectionPool) GetConnectionTimeout() int32 {
	if x != nil {
		return x.ConnectionTimeout
	}
	return 0
}

func (x *ConnectionPool) GetMaxLifetime() int32 {
	if x != nil {
		return x.MaxLifetime
	}
	return 0
}

func (x *ConnectionPool) GetMaxPoolSize() int32 {
	if x != nil {
		return x.MaxPoolSize
	}
	return 0
}

var File_storage_sql_proto protoreflect.FileDescriptor

var file_storage_sql_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x03, 0x53, 0x51, 0x4c, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x37, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x09, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x64, 0x62, 0x63, 0x55, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x64, 0x62, 0x63, 0x55, 0x72, 0x6c, 0x12,
	0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_sql_proto_rawDescOnce sync.Once
	file_storage_sql_proto_rawDescData = file_storage_sql_proto_rawDesc
)

func file_storage_sql_proto_rawDescGZIP() []byte {
	file_storage_sql_proto_rawDescOnce.Do(func() {
		file_storage_sql_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_sql_proto_rawDescData)
	})
	return file_storage_sql_proto_rawDescData
}

var file_storage_sql_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storage_sql_proto_goTypes = []interface{}{
	(*SQL)(nil),                // 0: proto.storage.SQL
	(*ConnectionPools)(nil),    // 1: proto.storage.ConnectionPools
	(*ConnectionPool)(nil),     // 2: proto.storage.ConnectionPool
	(*wrappers.BoolValue)(nil), // 3: google.protobuf.BoolValue
}
var file_storage_sql_proto_depIdxs = []int32{
	3, // 0: proto.storage.SQL.enabled:type_name -> google.protobuf.BoolValue
	2, // 1: proto.storage.SQL.default:type_name -> proto.storage.ConnectionPool
	2, // 2: proto.storage.SQL.migration:type_name -> proto.storage.ConnectionPool
	2, // 3: proto.storage.ConnectionPools.default:type_name -> proto.storage.ConnectionPool
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_sql_proto_init() }
func file_storage_sql_proto_init() {
	if File_storage_sql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_sql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQL); i {
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
		file_storage_sql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionPools); i {
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
		file_storage_sql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionPool); i {
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
			RawDescriptor: file_storage_sql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_sql_proto_goTypes,
		DependencyIndexes: file_storage_sql_proto_depIdxs,
		MessageInfos:      file_storage_sql_proto_msgTypes,
	}.Build()
	File_storage_sql_proto = out.File
	file_storage_sql_proto_rawDesc = nil
	file_storage_sql_proto_goTypes = nil
	file_storage_sql_proto_depIdxs = nil
}
