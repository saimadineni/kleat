// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: config/halconfig.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	client "github.com/spinnaker/kleat/api/client"
	artifact "github.com/spinnaker/kleat/api/client/artifact"
	canary "github.com/spinnaker/kleat/api/client/canary"
	ci "github.com/spinnaker/kleat/api/client/ci"
	cloudprovider "github.com/spinnaker/kleat/api/client/cloudprovider"
	notification "github.com/spinnaker/kleat/api/client/notification"
	pubsub "github.com/spinnaker/kleat/api/client/pubsub"
	security "github.com/spinnaker/kleat/api/client/security"
	storage "github.com/spinnaker/kleat/api/client/storage"
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

// Configuration for a Spinnaker installation.
type Hal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersistentStorage *storage.PersistentStorage  `protobuf:"bytes,1,opt,name=persistentStorage,proto3" json:"persistentStorage,omitempty"`
	Providers         *cloudprovider.Providers    `protobuf:"bytes,2,opt,name=providers,proto3" json:"providers,omitempty"`
	Artifacts         *artifact.Artifacts         `protobuf:"bytes,3,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Notifications     *notification.Notifications `protobuf:"bytes,4,opt,name=notifications,proto3" json:"notifications,omitempty"`
	Pubsub            *pubsub.Pubsub              `protobuf:"bytes,5,opt,name=pubsub,proto3" json:"pubsub,omitempty"`
	Ci                *ci.Ci                      `protobuf:"bytes,6,opt,name=ci,proto3" json:"ci,omitempty"`
	Stats             *client.Stats               `protobuf:"bytes,7,opt,name=stats,proto3" json:"stats,omitempty"`
	Features          *client.Features            `protobuf:"bytes,8,opt,name=features,proto3" json:"features,omitempty"`
	Webhook           *security.WebhookConfig     `protobuf:"bytes,9,opt,name=webhook,proto3" json:"webhook,omitempty"`
	Security          *security.Security          `protobuf:"bytes,10,opt,name=security,proto3" json:"security,omitempty"`
	Canary            *canary.Canary              `protobuf:"bytes,11,opt,name=canary,proto3" json:"canary,omitempty"`
	// The timezone in which your Spinnaker instance runs. This affects what the
	// UI will display as well as how CRON triggers are run.
	Timezone string `protobuf:"bytes,12,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Top-level Spinnaker version.
	Version string `protobuf:"bytes,13,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Hal) Reset() {
	*x = Hal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_halconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hal) ProtoMessage() {}

func (x *Hal) ProtoReflect() protoreflect.Message {
	mi := &file_config_halconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hal.ProtoReflect.Descriptor instead.
func (*Hal) Descriptor() ([]byte, []int) {
	return file_config_halconfig_proto_rawDescGZIP(), []int{0}
}

func (x *Hal) GetPersistentStorage() *storage.PersistentStorage {
	if x != nil {
		return x.PersistentStorage
	}
	return nil
}

func (x *Hal) GetProviders() *cloudprovider.Providers {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *Hal) GetArtifacts() *artifact.Artifacts {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *Hal) GetNotifications() *notification.Notifications {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *Hal) GetPubsub() *pubsub.Pubsub {
	if x != nil {
		return x.Pubsub
	}
	return nil
}

func (x *Hal) GetCi() *ci.Ci {
	if x != nil {
		return x.Ci
	}
	return nil
}

func (x *Hal) GetStats() *client.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Hal) GetFeatures() *client.Features {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Hal) GetWebhook() *security.WebhookConfig {
	if x != nil {
		return x.Webhook
	}
	return nil
}

func (x *Hal) GetSecurity() *security.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Hal) GetCanary() *canary.Canary {
	if x != nil {
		return x.Canary
	}
	return nil
}

func (x *Hal) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Hal) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_config_halconfig_proto protoreflect.FileDescriptor

var file_config_halconfig_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x68, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f,
	0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x05, 0x0a,
	0x03, 0x48, 0x61, 0x6c, 0x12, 0x4e, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x11, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x52, 0x06, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x12, 0x1c, 0x0a, 0x02, 0x63, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x43, 0x69, 0x52, 0x02, 0x63, 0x69,
	0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x37, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x2c, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_halconfig_proto_rawDescOnce sync.Once
	file_config_halconfig_proto_rawDescData = file_config_halconfig_proto_rawDesc
)

func file_config_halconfig_proto_rawDescGZIP() []byte {
	file_config_halconfig_proto_rawDescOnce.Do(func() {
		file_config_halconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_halconfig_proto_rawDescData)
	})
	return file_config_halconfig_proto_rawDescData
}

var file_config_halconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_halconfig_proto_goTypes = []interface{}{
	(*Hal)(nil),                        // 0: proto.config.Hal
	(*storage.PersistentStorage)(nil),  // 1: proto.storage.PersistentStorage
	(*cloudprovider.Providers)(nil),    // 2: proto.cloudprovider.Providers
	(*artifact.Artifacts)(nil),         // 3: proto.artifact.Artifacts
	(*notification.Notifications)(nil), // 4: proto.notification.Notifications
	(*pubsub.Pubsub)(nil),              // 5: proto.pubsub.Pubsub
	(*ci.Ci)(nil),                      // 6: proto.ci.Ci
	(*client.Stats)(nil),               // 7: proto.Stats
	(*client.Features)(nil),            // 8: proto.Features
	(*security.WebhookConfig)(nil),     // 9: proto.security.WebhookConfig
	(*security.Security)(nil),          // 10: proto.security.Security
	(*canary.Canary)(nil),              // 11: proto.canary.Canary
}
var file_config_halconfig_proto_depIdxs = []int32{
	1,  // 0: proto.config.Hal.persistentStorage:type_name -> proto.storage.PersistentStorage
	2,  // 1: proto.config.Hal.providers:type_name -> proto.cloudprovider.Providers
	3,  // 2: proto.config.Hal.artifacts:type_name -> proto.artifact.Artifacts
	4,  // 3: proto.config.Hal.notifications:type_name -> proto.notification.Notifications
	5,  // 4: proto.config.Hal.pubsub:type_name -> proto.pubsub.Pubsub
	6,  // 5: proto.config.Hal.ci:type_name -> proto.ci.Ci
	7,  // 6: proto.config.Hal.stats:type_name -> proto.Stats
	8,  // 7: proto.config.Hal.features:type_name -> proto.Features
	9,  // 8: proto.config.Hal.webhook:type_name -> proto.security.WebhookConfig
	10, // 9: proto.config.Hal.security:type_name -> proto.security.Security
	11, // 10: proto.config.Hal.canary:type_name -> proto.canary.Canary
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_config_halconfig_proto_init() }
func file_config_halconfig_proto_init() {
	if File_config_halconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_halconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hal); i {
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
			RawDescriptor: file_config_halconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_halconfig_proto_goTypes,
		DependencyIndexes: file_config_halconfig_proto_depIdxs,
		MessageInfos:      file_config_halconfig_proto_msgTypes,
	}.Build()
	File_config_halconfig_proto = out.File
	file_config_halconfig_proto_rawDesc = nil
	file_config_halconfig_proto_goTypes = nil
	file_config_halconfig_proto_depIdxs = nil
}
