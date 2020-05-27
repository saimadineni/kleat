// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: notification/notifications.proto

package notification

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

// Configuration for notifications.
type Notifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slack        *Slack        `protobuf:"bytes,1,opt,name=slack,proto3" json:"slack,omitempty"`
	Twilio       *Twilio       `protobuf:"bytes,2,opt,name=twilio,proto3" json:"twilio,omitempty"`
	GithubStatus *GithubStatus `protobuf:"bytes,3,opt,name=githubStatus,proto3" json:"githubStatus,omitempty"`
}

func (x *Notifications) Reset() {
	*x = Notifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifications) ProtoMessage() {}

func (x *Notifications) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifications.ProtoReflect.Descriptor instead.
func (*Notifications) Descriptor() ([]byte, []int) {
	return file_notification_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *Notifications) GetSlack() *Slack {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *Notifications) GetTwilio() *Twilio {
	if x != nil {
		return x.Twilio
	}
	return nil
}

func (x *Notifications) GetGithubStatus() *GithubStatus {
	if x != nil {
		return x.GithubStatus
	}
	return nil
}

var File_notification_notifications_proto protoreflect.FileDescriptor

var file_notification_notifications_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01,
	0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2f, 0x0a, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x12, 0x32, 0x0a, 0x06, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x52, 0x06, 0x74, 0x77,
	0x69, 0x6c, 0x69, 0x6f, 0x12, 0x44, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b,
	0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_notifications_proto_rawDescOnce sync.Once
	file_notification_notifications_proto_rawDescData = file_notification_notifications_proto_rawDesc
)

func file_notification_notifications_proto_rawDescGZIP() []byte {
	file_notification_notifications_proto_rawDescOnce.Do(func() {
		file_notification_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_notifications_proto_rawDescData)
	})
	return file_notification_notifications_proto_rawDescData
}

var file_notification_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notification_notifications_proto_goTypes = []interface{}{
	(*Notifications)(nil), // 0: proto.notification.Notifications
	(*Slack)(nil),         // 1: proto.notification.Slack
	(*Twilio)(nil),        // 2: proto.notification.Twilio
	(*GithubStatus)(nil),  // 3: proto.notification.GithubStatus
}
var file_notification_notifications_proto_depIdxs = []int32{
	1, // 0: proto.notification.Notifications.slack:type_name -> proto.notification.Slack
	2, // 1: proto.notification.Notifications.twilio:type_name -> proto.notification.Twilio
	3, // 2: proto.notification.Notifications.githubStatus:type_name -> proto.notification.GithubStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notification_notifications_proto_init() }
func file_notification_notifications_proto_init() {
	if File_notification_notifications_proto != nil {
		return
	}
	file_notification_github_status_proto_init()
	file_notification_slack_proto_init()
	file_notification_twilio_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notification_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifications); i {
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
			RawDescriptor: file_notification_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_notifications_proto_goTypes,
		DependencyIndexes: file_notification_notifications_proto_depIdxs,
		MessageInfos:      file_notification_notifications_proto_msgTypes,
	}.Build()
	File_notification_notifications_proto = out.File
	file_notification_notifications_proto_rawDesc = nil
	file_notification_notifications_proto_goTypes = nil
	file_notification_notifications_proto_depIdxs = nil
}
