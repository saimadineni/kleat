// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: pubsub/google.proto

package pubsub

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

// Represents the format of an incoming pub/sub message.
type MessageFormat int32

const (
	MessageFormat_CUSTOM MessageFormat = 0
	MessageFormat_GCB    MessageFormat = 1
	MessageFormat_GCS    MessageFormat = 2
	MessageFormat_GCR    MessageFormat = 3
)

// Enum value maps for MessageFormat.
var (
	MessageFormat_name = map[int32]string{
		0: "CUSTOM",
		1: "GCB",
		2: "GCS",
		3: "GCR",
	}
	MessageFormat_value = map[string]int32{
		"CUSTOM": 0,
		"GCB":    1,
		"GCS":    2,
		"GCR":    3,
	}
)

func (x MessageFormat) Enum() *MessageFormat {
	p := new(MessageFormat)
	*p = x
	return p
}

func (x MessageFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_pubsub_google_proto_enumTypes[0].Descriptor()
}

func (MessageFormat) Type() protoreflect.EnumType {
	return &file_pubsub_google_proto_enumTypes[0]
}

func (x MessageFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageFormat.Descriptor instead.
func (MessageFormat) EnumDescriptor() ([]byte, []int) {
	return file_pubsub_google_proto_rawDescGZIP(), []int{0}
}

// Configuration for Google Cloud Pub/Sub integration.
type Google struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Google Cloud Pub/Sub is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured subscriptions.
	Subscriptions []*GoogleSubscriber `protobuf:"bytes,2,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	// The list of configured publishers.
	Publishers []*GooglePublisher `protobuf:"bytes,3,rep,name=publishers,proto3" json:"publishers,omitempty"`
}

func (x *Google) Reset() {
	*x = Google{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_google_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Google) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Google) ProtoMessage() {}

func (x *Google) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_google_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Google.ProtoReflect.Descriptor instead.
func (*Google) Descriptor() ([]byte, []int) {
	return file_pubsub_google_proto_rawDescGZIP(), []int{0}
}

func (x *Google) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Google) GetSubscriptions() []*GoogleSubscriber {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *Google) GetPublishers() []*GooglePublisher {
	if x != nil {
		return x.Publishers
	}
	return nil
}

// Configuration for a Google Cloud Pub/Sub subscriber.
type GoogleSubscriber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the subscriber account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the GCP project your subscription lives in.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the subscription to listen to. This identifier does not include
	// the name of the project, and must already be configured.
	SubscriptionName string `protobuf:"bytes,3,opt,name=subscriptionName,proto3" json:"subscriptionName,omitempty"`
	// The path to a JSON service account that Spinnaker will use as credentials.
	// This is only needed if Spinnaker is not deployed on a Google Compute Engine
	// VM, or needs permissions not afforded to the VM it is running on.
	// See https://cloud.google.com/compute/docs/access/service-accounts for more information.
	JsonPath string `protobuf:"bytes,4,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// The acknowledgement deadline as configured on the Pub/Sub subscription.
	AckDeadlineSeconds int32 `protobuf:"varint,5,opt,name=ackDeadlineSeconds,proto3" json:"ackDeadlineSeconds,omitempty"`
	// The format of the incoming message. Used to translate the incoming message
	// into Spinnaker artifacts.
	MessageFormat MessageFormat `protobuf:"varint,6,opt,name=messageFormat,proto3,enum=proto.pubsub.MessageFormat" json:"messageFormat,omitempty"`
	// A path to a jinja template that specifies how artifacts from this pubsub system
	// are interpreted and transformed into Spinnaker artifacts. Only used if
	// messageFormat is set to CUSTOM.
	TemplatePath string `protobuf:"bytes,7,opt,name=templatePath,proto3" json:"templatePath,omitempty"`
}

func (x *GoogleSubscriber) Reset() {
	*x = GoogleSubscriber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_google_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleSubscriber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleSubscriber) ProtoMessage() {}

func (x *GoogleSubscriber) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_google_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleSubscriber.ProtoReflect.Descriptor instead.
func (*GoogleSubscriber) Descriptor() ([]byte, []int) {
	return file_pubsub_google_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleSubscriber) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleSubscriber) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GoogleSubscriber) GetSubscriptionName() string {
	if x != nil {
		return x.SubscriptionName
	}
	return ""
}

func (x *GoogleSubscriber) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *GoogleSubscriber) GetAckDeadlineSeconds() int32 {
	if x != nil {
		return x.AckDeadlineSeconds
	}
	return 0
}

func (x *GoogleSubscriber) GetMessageFormat() MessageFormat {
	if x != nil {
		return x.MessageFormat
	}
	return MessageFormat_CUSTOM
}

func (x *GoogleSubscriber) GetTemplatePath() string {
	if x != nil {
		return x.TemplatePath
	}
	return ""
}

// Configuration for a Google Cloud Pub/Sub publisher.
type GooglePublisher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the publisher account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the GCP project your topic lives in.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the topic to publish to. This identifier does not include the
	// name of the project, and must already be configured.
	TopicName string `protobuf:"bytes,3,opt,name=topicName,proto3" json:"topicName,omitempty"`
	// The path to a JSON service account that Spinnaker will use as credentials.
	// This is only needed if Spinnaker is not deployed on a Google Compute Engine
	// VM, or needs permissions not afforded to the VM it is running on.
	// See https://cloud.google.com/compute/docs/access/service-accounts for more information.
	JsonPath string `protobuf:"bytes,4,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// The content to publish to the topic. Must be one of ALL or NOTIFICATIONS.
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GooglePublisher) Reset() {
	*x = GooglePublisher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_google_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GooglePublisher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GooglePublisher) ProtoMessage() {}

func (x *GooglePublisher) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_google_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GooglePublisher.ProtoReflect.Descriptor instead.
func (*GooglePublisher) Descriptor() ([]byte, []int) {
	return file_pubsub_google_proto_rawDescGZIP(), []int{2}
}

func (x *GooglePublisher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GooglePublisher) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GooglePublisher) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *GooglePublisher) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *GooglePublisher) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_pubsub_google_proto protoreflect.FileDescriptor

var file_pubsub_google_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x22, 0xa7, 0x01, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x52, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x22, 0x9f, 0x02,
	0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x63, 0x6b, 0x44,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x93, 0x01, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x36, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x43, 0x42, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x43, 0x53, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x43, 0x52, 0x10, 0x03, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e,
	0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pubsub_google_proto_rawDescOnce sync.Once
	file_pubsub_google_proto_rawDescData = file_pubsub_google_proto_rawDesc
)

func file_pubsub_google_proto_rawDescGZIP() []byte {
	file_pubsub_google_proto_rawDescOnce.Do(func() {
		file_pubsub_google_proto_rawDescData = protoimpl.X.CompressGZIP(file_pubsub_google_proto_rawDescData)
	})
	return file_pubsub_google_proto_rawDescData
}

var file_pubsub_google_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pubsub_google_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pubsub_google_proto_goTypes = []interface{}{
	(MessageFormat)(0),       // 0: proto.pubsub.MessageFormat
	(*Google)(nil),           // 1: proto.pubsub.Google
	(*GoogleSubscriber)(nil), // 2: proto.pubsub.GoogleSubscriber
	(*GooglePublisher)(nil),  // 3: proto.pubsub.GooglePublisher
}
var file_pubsub_google_proto_depIdxs = []int32{
	2, // 0: proto.pubsub.Google.subscriptions:type_name -> proto.pubsub.GoogleSubscriber
	3, // 1: proto.pubsub.Google.publishers:type_name -> proto.pubsub.GooglePublisher
	0, // 2: proto.pubsub.GoogleSubscriber.messageFormat:type_name -> proto.pubsub.MessageFormat
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pubsub_google_proto_init() }
func file_pubsub_google_proto_init() {
	if File_pubsub_google_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pubsub_google_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Google); i {
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
		file_pubsub_google_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleSubscriber); i {
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
		file_pubsub_google_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GooglePublisher); i {
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
			RawDescriptor: file_pubsub_google_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pubsub_google_proto_goTypes,
		DependencyIndexes: file_pubsub_google_proto_depIdxs,
		EnumInfos:         file_pubsub_google_proto_enumTypes,
		MessageInfos:      file_pubsub_google_proto_msgTypes,
	}.Build()
	File_pubsub_google_proto = out.File
	file_pubsub_google_proto_rawDesc = nil
	file_pubsub_google_proto_goTypes = nil
	file_pubsub_google_proto_depIdxs = nil
}
