// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/userpb/v3/group.proto

package userv3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v3 "github.com/paralus/paralus/proto/types/commonpb/v3"
	_ "github.com/paralus/paralus/proto/types/rolepb/v3"
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

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string       `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string       `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec       *GroupSpec   `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Status     *v3.Status   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Group) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Group) GetMetadata() *v3.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Group) GetSpec() *GroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Group) GetStatus() *v3.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type ProjectNamespaceRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   *string `protobuf:"bytes,1,opt,name=project,proto3,oneof" json:"project,omitempty"`
	Namespace *string `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Role      string  `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Group     *string `protobuf:"bytes,4,opt,name=group,proto3,oneof" json:"group,omitempty"`
}

func (x *ProjectNamespaceRole) Reset() {
	*x = ProjectNamespaceRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectNamespaceRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectNamespaceRole) ProtoMessage() {}

func (x *ProjectNamespaceRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectNamespaceRole.ProtoReflect.Descriptor instead.
func (*ProjectNamespaceRole) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_group_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectNamespaceRole) GetProject() string {
	if x != nil && x.Project != nil {
		return *x.Project
	}
	return ""
}

func (x *ProjectNamespaceRole) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *ProjectNamespaceRole) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ProjectNamespaceRole) GetGroup() string {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return ""
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project     *string  `protobuf:"bytes,1,opt,name=project,proto3,oneof" json:"project,omitempty"`
	Namespace   *string  `protobuf:"bytes,2,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Role        string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Permissions []string `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_group_proto_rawDescGZIP(), []int{2}
}

func (x *Permission) GetProject() string {
	if x != nil && x.Project != nil {
		return *x.Project
	}
	return ""
}

func (x *Permission) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *Permission) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Permission) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type GroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectNamespaceRoles []*ProjectNamespaceRole `protobuf:"bytes,1,rep,name=projectNamespaceRoles,proto3" json:"projectNamespaceRoles,omitempty"`
	Users                 []string                `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Type                  string                  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GroupSpec) Reset() {
	*x = GroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSpec) ProtoMessage() {}

func (x *GroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSpec.ProtoReflect.Descriptor instead.
func (*GroupSpec) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_group_proto_rawDescGZIP(), []int{3}
}

func (x *GroupSpec) GetProjectNamespaceRoles() []*ProjectNamespaceRole {
	if x != nil {
		return x.ProjectNamespaceRoles
	}
	return nil
}

func (x *GroupSpec) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GroupSpec) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GroupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string           `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string           `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *v3.ListMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Items      []*Group         `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GroupList) Reset() {
	*x = GroupList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_userpb_v3_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupList) ProtoMessage() {}

func (x *GroupList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_userpb_v3_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupList.ProtoReflect.Descriptor instead.
func (*GroupList) Descriptor() ([]byte, []int) {
	return file_proto_types_userpb_v3_group_proto_rawDescGZIP(), []int{4}
}

func (x *GroupList) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *GroupList) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *GroupList) GetMetadata() *v3.ListMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GroupList) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_types_userpb_v3_group_proto protoreflect.FileDescriptor

var file_proto_types_userpb_v3_group_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x24, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x04, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x69,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x49, 0x92, 0x41, 0x46, 0x2a, 0x0b, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x32, 0x21, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x6b, 0x38, 0x73, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x33, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0x92, 0x41, 0x29, 0x2a, 0x04, 0x4b, 0x69,
	0x6e, 0x64, 0x32, 0x1a, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x05,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x6e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x2d, 0x92, 0x41, 0x2a, 0x2a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x1e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x65, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x42, 0x2d, 0x92,
	0x41, 0x2a, 0x2a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x1e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x60, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x25, 0x92, 0x41, 0x22, 0x2a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0x16, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x3a, 0x39, 0x92, 0x41, 0x36, 0x0a, 0x34, 0x2a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x32, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xd2, 0x01, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xd2, 0x01, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0xd2, 0x01, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xd2, 0x01, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22,
	0xd2, 0x02, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12, 0x2a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0x92, 0x41, 0x16, 0x2a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x32, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x01, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x2a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x32, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x2c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0x92, 0x41, 0x0e, 0x2a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x48, 0x02, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x88, 0x01, 0x01, 0x3a,
	0x4f, 0x92, 0x41, 0x4c, 0x0a, 0x4a, 0x2a, 0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x32, 0x32, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2c, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x20, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0xb2, 0x02, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12, 0x2a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x32, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0x92, 0x41,
	0x16, 0x2a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x2a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x32, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x2e, 0x92, 0x41, 0x2b, 0x2a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x1c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x72, 0x6f, 0x6c,
	0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x1f,
	0x92, 0x41, 0x1c, 0x0a, 0x1a, 0x2a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xd8, 0x02, 0x0a, 0x09, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0xb0, 0x01, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x4b, 0x92, 0x41, 0x48, 0x2a, 0x15, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x32, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2c, 0x20, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2c, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x61, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x23, 0x92, 0x41, 0x20, 0x2a, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x32, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0x92, 0x41, 0x15, 0x2a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x32, 0x0d,
	0x54, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x3a, 0x2f, 0x92, 0x41, 0x2c, 0x0a, 0x2a, 0x2a, 0x13, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x13, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa5, 0x03, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x5a, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x37, 0x2a, 0x0b, 0x41, 0x50, 0x49,
	0x20, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x26, 0x41, 0x50, 0x49, 0x20, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x40, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0x92, 0x41,
	0x29, 0x2a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x32, 0x1f, 0x4b, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x79, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x34, 0x92, 0x41, 0x31,
	0x2a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x23, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40,
	0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5f, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x29, 0x92, 0x41, 0x26, 0x2a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x40, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x1e, 0x92, 0x41,
	0x1b, 0x0a, 0x19, 0x2a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x0a,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x40, 0x01, 0x42, 0xe9, 0x01, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x4c, 0x61, 0x62, 0x73,
	0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f,
	0x76, 0x33, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x33, 0xa2, 0x02, 0x04, 0x52, 0x44, 0x54, 0x55,
	0xaa, 0x02, 0x17, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x17, 0x52, 0x61, 0x66,
	0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x23, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76,
	0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x33, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x52, 0x61, 0x66,
	0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a,
	0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_userpb_v3_group_proto_rawDescOnce sync.Once
	file_proto_types_userpb_v3_group_proto_rawDescData = file_proto_types_userpb_v3_group_proto_rawDesc
)

func file_proto_types_userpb_v3_group_proto_rawDescGZIP() []byte {
	file_proto_types_userpb_v3_group_proto_rawDescOnce.Do(func() {
		file_proto_types_userpb_v3_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_userpb_v3_group_proto_rawDescData)
	})
	return file_proto_types_userpb_v3_group_proto_rawDescData
}

var file_proto_types_userpb_v3_group_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_types_userpb_v3_group_proto_goTypes = []interface{}{
	(*Group)(nil),                // 0: paralus.dev.types.user.v3.Group
	(*ProjectNamespaceRole)(nil), // 1: paralus.dev.types.user.v3.ProjectNamespaceRole
	(*Permission)(nil),           // 2: paralus.dev.types.user.v3.Permission
	(*GroupSpec)(nil),            // 3: paralus.dev.types.user.v3.GroupSpec
	(*GroupList)(nil),            // 4: paralus.dev.types.user.v3.GroupList
	(*v3.Metadata)(nil),          // 5: paralus.dev.types.common.v3.Metadata
	(*v3.Status)(nil),            // 6: paralus.dev.types.common.v3.Status
	(*v3.ListMetadata)(nil),      // 7: paralus.dev.types.common.v3.ListMetadata
}
var file_proto_types_userpb_v3_group_proto_depIdxs = []int32{
	5, // 0: paralus.dev.types.user.v3.Group.metadata:type_name -> paralus.dev.types.common.v3.Metadata
	3, // 1: paralus.dev.types.user.v3.Group.spec:type_name -> paralus.dev.types.user.v3.GroupSpec
	6, // 2: paralus.dev.types.user.v3.Group.status:type_name -> paralus.dev.types.common.v3.Status
	1, // 3: paralus.dev.types.user.v3.GroupSpec.projectNamespaceRoles:type_name -> paralus.dev.types.user.v3.ProjectNamespaceRole
	7, // 4: paralus.dev.types.user.v3.GroupList.metadata:type_name -> paralus.dev.types.common.v3.ListMetadata
	0, // 5: paralus.dev.types.user.v3.GroupList.items:type_name -> paralus.dev.types.user.v3.Group
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_types_userpb_v3_group_proto_init() }
func file_proto_types_userpb_v3_group_proto_init() {
	if File_proto_types_userpb_v3_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_userpb_v3_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_proto_types_userpb_v3_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectNamespaceRole); i {
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
		file_proto_types_userpb_v3_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_proto_types_userpb_v3_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupSpec); i {
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
		file_proto_types_userpb_v3_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupList); i {
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
	file_proto_types_userpb_v3_group_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_proto_types_userpb_v3_group_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_types_userpb_v3_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_userpb_v3_group_proto_goTypes,
		DependencyIndexes: file_proto_types_userpb_v3_group_proto_depIdxs,
		MessageInfos:      file_proto_types_userpb_v3_group_proto_msgTypes,
	}.Build()
	File_proto_types_userpb_v3_group_proto = out.File
	file_proto_types_userpb_v3_group_proto_rawDesc = nil
	file_proto_types_userpb_v3_group_proto_goTypes = nil
	file_proto_types_userpb_v3_group_proto_depIdxs = nil
}
