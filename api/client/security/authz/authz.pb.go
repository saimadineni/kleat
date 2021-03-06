// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: security/authz/authz.proto

package authz

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

// Configuration for which role provider to use for authorization decisions.
type GroupMembership_RoleProviderType int32

const (
	// Unspecified. Do not directly use, instead omit the field.
	GroupMembership_UNSPECIFIED GroupMembership_RoleProviderType = 0
	// File-based role provider.
	GroupMembership_FILE GroupMembership_RoleProviderType = 1
	// Google role provider.
	GroupMembership_GOOGLE GroupMembership_RoleProviderType = 2
	// GitHub role provider.
	GroupMembership_GITHUB GroupMembership_RoleProviderType = 3
	// LDAP role provider.
	GroupMembership_LDAP GroupMembership_RoleProviderType = 4
)

// Enum value maps for GroupMembership_RoleProviderType.
var (
	GroupMembership_RoleProviderType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "FILE",
		2: "GOOGLE",
		3: "GITHUB",
		4: "LDAP",
	}
	GroupMembership_RoleProviderType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"FILE":        1,
		"GOOGLE":      2,
		"GITHUB":      3,
		"LDAP":        4,
	}
)

func (x GroupMembership_RoleProviderType) Enum() *GroupMembership_RoleProviderType {
	p := new(GroupMembership_RoleProviderType)
	*p = x
	return p
}

func (x GroupMembership_RoleProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupMembership_RoleProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_security_authz_authz_proto_enumTypes[0].Descriptor()
}

func (GroupMembership_RoleProviderType) Type() protoreflect.EnumType {
	return &file_security_authz_authz_proto_enumTypes[0]
}

func (x GroupMembership_RoleProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupMembership_RoleProviderType.Descriptor instead.
func (GroupMembership_RoleProviderType) EnumDescriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{1, 0}
}

// Configuration for what resources users of Spinnaker can read and modify.
type Authorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Spinnaker's role-based authorization is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Configuration role providers that map users to groups.
	GroupMembership *GroupMembership `protobuf:"bytes,2,opt,name=groupMembership,proto3" json:"groupMembership,omitempty"`
}

func (x *Authorization) Reset() {
	*x = Authorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_authz_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization) ProtoMessage() {}

func (x *Authorization) ProtoReflect() protoreflect.Message {
	mi := &file_security_authz_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization.ProtoReflect.Descriptor instead.
func (*Authorization) Descriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{0}
}

func (x *Authorization) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Authorization) GetGroupMembership() *GroupMembership {
	if x != nil {
		return x.GroupMembership
	}
	return nil
}

// Configuration role providers that map users to groups.
type GroupMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for which role provider to use for authorization decisions.
	// Each role provider has a corresponding field; configuration specific to
	// the role provider you are using should be added to the appropriate field.
	Service GroupMembership_RoleProviderType `protobuf:"varint,1,opt,name=service,proto3,enum=proto.security.authz.GroupMembership_RoleProviderType" json:"service,omitempty"`
	// Configuration for the Google role provider.
	Google *GoogleRoleProvider `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	// Configuration for the GitHub role provider.
	Github *GithubRoleProvider `protobuf:"bytes,3,opt,name=github,proto3" json:"github,omitempty"`
	// Configuration for the file-based role provider.
	File *FileRoleProvider `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	// Configuration for the LDAP role provider.
	Ldap *LdapRoleProvider `protobuf:"bytes,5,opt,name=ldap,proto3" json:"ldap,omitempty"`
}

func (x *GroupMembership) Reset() {
	*x = GroupMembership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_authz_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembership) ProtoMessage() {}

func (x *GroupMembership) ProtoReflect() protoreflect.Message {
	mi := &file_security_authz_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembership.ProtoReflect.Descriptor instead.
func (*GroupMembership) Descriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{1}
}

func (x *GroupMembership) GetService() GroupMembership_RoleProviderType {
	if x != nil {
		return x.Service
	}
	return GroupMembership_UNSPECIFIED
}

func (x *GroupMembership) GetGoogle() *GoogleRoleProvider {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *GroupMembership) GetGithub() *GithubRoleProvider {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *GroupMembership) GetFile() *FileRoleProvider {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *GroupMembership) GetLdap() *LdapRoleProvider {
	if x != nil {
		return x.Ldap
	}
	return nil
}

// Configuration for the Google role provider.
type GoogleRoleProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A path to a valid json service account that can authenticate against the
	// Google role provider.
	CredentialPath string `protobuf:"bytes,1,opt,name=credentialPath,proto3" json:"credentialPath,omitempty"`
	// Your role provider's admin username e.g. admin@myorg.net.
	AdminUsername string `protobuf:"bytes,2,opt,name=adminUsername,proto3" json:"adminUsername,omitempty"`
	// The domain your role provider is configured for e.g. myorg.net.
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *GoogleRoleProvider) Reset() {
	*x = GoogleRoleProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_authz_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleRoleProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleRoleProvider) ProtoMessage() {}

func (x *GoogleRoleProvider) ProtoReflect() protoreflect.Message {
	mi := &file_security_authz_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleRoleProvider.ProtoReflect.Descriptor instead.
func (*GoogleRoleProvider) Descriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{2}
}

func (x *GoogleRoleProvider) GetCredentialPath() string {
	if x != nil {
		return x.CredentialPath
	}
	return ""
}

func (x *GoogleRoleProvider) GetAdminUsername() string {
	if x != nil {
		return x.AdminUsername
	}
	return ""
}

func (x *GoogleRoleProvider) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// Configuration for the GitHub role provider.
type GithubRoleProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used if using GitHub enterprise some other non github.com GitHub installation.
	BaseUrl string `protobuf:"bytes,1,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	// A personal access token of an account with access to your organization's
	// GitHub Teams structure.
	AccessToken string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	// The GitHub organization under which to query for GitHub Teams.
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *GithubRoleProvider) Reset() {
	*x = GithubRoleProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_authz_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubRoleProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubRoleProvider) ProtoMessage() {}

func (x *GithubRoleProvider) ProtoReflect() protoreflect.Message {
	mi := &file_security_authz_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubRoleProvider.ProtoReflect.Descriptor instead.
func (*GithubRoleProvider) Descriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{3}
}

func (x *GithubRoleProvider) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *GithubRoleProvider) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GithubRoleProvider) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

// Configuration for the file-based role provider.
type FileRoleProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A path to a file describing the roles of each user.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileRoleProvider) Reset() {
	*x = FileRoleProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_authz_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRoleProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRoleProvider) ProtoMessage() {}

func (x *FileRoleProvider) ProtoReflect() protoreflect.Message {
	mi := &file_security_authz_authz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRoleProvider.ProtoReflect.Descriptor instead.
func (*FileRoleProvider) Descriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{4}
}

func (x *FileRoleProvider) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Configuration for the LDAP role provider.
type LdapRoleProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ldap:// or ldaps:// url of the LDAP server.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The manager user's distinguished name (principal) to use for querying LDAP groups.
	ManagerDn string `protobuf:"bytes,2,opt,name=managerDn,proto3" json:"managerDn,omitempty"`
	// The manager user's password to use for querying LDAP groups.
	ManagerPassword string `protobuf:"bytes,3,opt,name=managerPassword,proto3" json:"managerPassword,omitempty"`
	// The pattern for finding a user's DN using simple pattern matching. For example,
	// if your LDAP server has the URL ldap://mysite.com/dc=spinnaker,dc=org, and
	// you have the pattern 'uid={0},ou=members', 'me' will map to a DN
	// uid=me,ou=members,dc=spinnaker,dc=org. If no match is found, will try to
	// find the user using -user-search-filter, if set.
	UserDnPattern string `protobuf:"bytes,4,opt,name=userDnPattern,proto3" json:"userDnPattern,omitempty"`
	// The part of the directory tree under which user searches should be performed.
	// If -user-search-base isn't supplied, the search will be performed from the root.
	UserSearchBase string `protobuf:"bytes,5,opt,name=userSearchBase,proto3" json:"userSearchBase,omitempty"`
	// The part of the directory tree under which group searches should be performed.
	GroupSearchBase string `protobuf:"bytes,6,opt,name=groupSearchBase,proto3" json:"groupSearchBase,omitempty"`
	// The filter to use when searching for a user's DN. Will search either from
	// -user-search-base (if specified) or root for entries matching the filter.
	UserSearchFilter string `protobuf:"bytes,7,opt,name=userSearchFilter,proto3" json:"userSearchFilter,omitempty"`
	// The filter which is used to search for group membership. The default is
	// 'uniqueMember={0}', corresponding to the groupOfUniqueMembers LDAP class. In
	// this case, the substituted parameter is the full distinguished name of the
	// user. The parameter '{1}' can be used if you want to filter on the login name.
	GroupSearchFilter string `protobuf:"bytes,8,opt,name=groupSearchFilter,proto3" json:"groupSearchFilter,omitempty"`
	// The attribute which contains the name of the authority defined by the group
	// entry. Defaults to 'cn'.
	GroupRoleAttributes string `protobuf:"bytes,9,opt,name=groupRoleAttributes,proto3" json:"groupRoleAttributes,omitempty"`
}

func (x *LdapRoleProvider) Reset() {
	*x = LdapRoleProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_authz_authz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LdapRoleProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LdapRoleProvider) ProtoMessage() {}

func (x *LdapRoleProvider) ProtoReflect() protoreflect.Message {
	mi := &file_security_authz_authz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LdapRoleProvider.ProtoReflect.Descriptor instead.
func (*LdapRoleProvider) Descriptor() ([]byte, []int) {
	return file_security_authz_authz_proto_rawDescGZIP(), []int{5}
}

func (x *LdapRoleProvider) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LdapRoleProvider) GetManagerDn() string {
	if x != nil {
		return x.ManagerDn
	}
	return ""
}

func (x *LdapRoleProvider) GetManagerPassword() string {
	if x != nil {
		return x.ManagerPassword
	}
	return ""
}

func (x *LdapRoleProvider) GetUserDnPattern() string {
	if x != nil {
		return x.UserDnPattern
	}
	return ""
}

func (x *LdapRoleProvider) GetUserSearchBase() string {
	if x != nil {
		return x.UserSearchBase
	}
	return ""
}

func (x *LdapRoleProvider) GetGroupSearchBase() string {
	if x != nil {
		return x.GroupSearchBase
	}
	return ""
}

func (x *LdapRoleProvider) GetUserSearchFilter() string {
	if x != nil {
		return x.UserSearchFilter
	}
	return ""
}

func (x *LdapRoleProvider) GetGroupSearchFilter() string {
	if x != nil {
		return x.GroupSearchFilter
	}
	return ""
}

func (x *LdapRoleProvider) GetGroupRoleAttributes() string {
	if x != nil {
		return x.GroupRoleAttributes
	}
	return ""
}

var File_security_authz_authz_proto protoreflect.FileDescriptor

var file_security_authz_authz_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x4f, 0x0a, 0x0f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x22, 0xb0, 0x03, 0x0a, 0x0f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x50, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x3a, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x4c, 0x64, 0x61, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x22, 0x4f, 0x0a,
	0x10, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48,
	0x55, 0x42, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x44, 0x41, 0x50, 0x10, 0x04, 0x22, 0x7a,
	0x0a, 0x12, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x74, 0x0a, 0x12, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x26, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xf0, 0x02, 0x0a, 0x10, 0x4c, 0x64, 0x61,
	0x70, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x6e, 0x12, 0x28, 0x0a,
	0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x44,
	0x6e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x6e, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x26, 0x0a,
	0x0e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x73, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x42, 0x61, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x6f, 0x6c,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_authz_authz_proto_rawDescOnce sync.Once
	file_security_authz_authz_proto_rawDescData = file_security_authz_authz_proto_rawDesc
)

func file_security_authz_authz_proto_rawDescGZIP() []byte {
	file_security_authz_authz_proto_rawDescOnce.Do(func() {
		file_security_authz_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_authz_authz_proto_rawDescData)
	})
	return file_security_authz_authz_proto_rawDescData
}

var file_security_authz_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_security_authz_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_security_authz_authz_proto_goTypes = []interface{}{
	(GroupMembership_RoleProviderType)(0), // 0: proto.security.authz.GroupMembership.RoleProviderType
	(*Authorization)(nil),                 // 1: proto.security.authz.Authorization
	(*GroupMembership)(nil),               // 2: proto.security.authz.GroupMembership
	(*GoogleRoleProvider)(nil),            // 3: proto.security.authz.GoogleRoleProvider
	(*GithubRoleProvider)(nil),            // 4: proto.security.authz.GithubRoleProvider
	(*FileRoleProvider)(nil),              // 5: proto.security.authz.FileRoleProvider
	(*LdapRoleProvider)(nil),              // 6: proto.security.authz.LdapRoleProvider
	(*wrappers.BoolValue)(nil),            // 7: google.protobuf.BoolValue
}
var file_security_authz_authz_proto_depIdxs = []int32{
	7, // 0: proto.security.authz.Authorization.enabled:type_name -> google.protobuf.BoolValue
	2, // 1: proto.security.authz.Authorization.groupMembership:type_name -> proto.security.authz.GroupMembership
	0, // 2: proto.security.authz.GroupMembership.service:type_name -> proto.security.authz.GroupMembership.RoleProviderType
	3, // 3: proto.security.authz.GroupMembership.google:type_name -> proto.security.authz.GoogleRoleProvider
	4, // 4: proto.security.authz.GroupMembership.github:type_name -> proto.security.authz.GithubRoleProvider
	5, // 5: proto.security.authz.GroupMembership.file:type_name -> proto.security.authz.FileRoleProvider
	6, // 6: proto.security.authz.GroupMembership.ldap:type_name -> proto.security.authz.LdapRoleProvider
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_security_authz_authz_proto_init() }
func file_security_authz_authz_proto_init() {
	if File_security_authz_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_authz_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authorization); i {
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
		file_security_authz_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMembership); i {
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
		file_security_authz_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleRoleProvider); i {
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
		file_security_authz_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubRoleProvider); i {
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
		file_security_authz_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRoleProvider); i {
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
		file_security_authz_authz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LdapRoleProvider); i {
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
			RawDescriptor: file_security_authz_authz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_authz_authz_proto_goTypes,
		DependencyIndexes: file_security_authz_authz_proto_depIdxs,
		EnumInfos:         file_security_authz_authz_proto_enumTypes,
		MessageInfos:      file_security_authz_authz_proto_msgTypes,
	}.Build()
	File_security_authz_authz_proto = out.File
	file_security_authz_authz_proto_rawDesc = nil
	file_security_authz_authz_proto_goTypes = nil
	file_security_authz_authz_proto_depIdxs = nil
}
