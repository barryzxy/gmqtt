// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.13.0
// source: publish.proto

package admin

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	Payload   string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Qos       uint32 `protobuf:"varint,3,opt,name=qos,proto3" json:"qos,omitempty"`
	Retained  bool   `protobuf:"varint,4,opt,name=retained,proto3" json:"retained,omitempty"`
	// the following fields are using in v5 client.
	ContentType     string            `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	CorrelationData string            `protobuf:"bytes,6,opt,name=correlation_data,json=correlationData,proto3" json:"correlation_data,omitempty"`
	MessageExpiry   uint32            `protobuf:"varint,7,opt,name=message_expiry,json=messageExpiry,proto3" json:"message_expiry,omitempty"`
	PayloadFormat   uint32            `protobuf:"varint,8,opt,name=payload_format,json=payloadFormat,proto3" json:"payload_format,omitempty"`
	ResponseTopic   string            `protobuf:"bytes,9,opt,name=response_topic,json=responseTopic,proto3" json:"response_topic,omitempty"`
	UserProperties  []*UserProperties `protobuf:"bytes,10,rep,name=user_properties,json=userProperties,proto3" json:"user_properties,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *PublishRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *PublishRequest) GetQos() uint32 {
	if x != nil {
		return x.Qos
	}
	return 0
}

func (x *PublishRequest) GetRetained() bool {
	if x != nil {
		return x.Retained
	}
	return false
}

func (x *PublishRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *PublishRequest) GetCorrelationData() string {
	if x != nil {
		return x.CorrelationData
	}
	return ""
}

func (x *PublishRequest) GetMessageExpiry() uint32 {
	if x != nil {
		return x.MessageExpiry
	}
	return 0
}

func (x *PublishRequest) GetPayloadFormat() uint32 {
	if x != nil {
		return x.PayloadFormat
	}
	return 0
}

func (x *PublishRequest) GetResponseTopic() string {
	if x != nil {
		return x.ResponseTopic
	}
	return ""
}

func (x *PublishRequest) GetUserProperties() []*UserProperties {
	if x != nil {
		return x.UserProperties
	}
	return nil
}

type UserProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K []byte `protobuf:"bytes,1,opt,name=K,proto3" json:"K,omitempty"`
	V []byte `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *UserProperties) Reset() {
	*x = UserProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProperties) ProtoMessage() {}

func (x *UserProperties) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProperties.ProtoReflect.Descriptor instead.
func (*UserProperties) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{1}
}

func (x *UserProperties) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *UserProperties) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

var File_publish_proto protoreflect.FileDescriptor

var file_publish_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x0e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x48, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x2c, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x4b, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x56,
	0x32, 0x6c, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1f, 0x2e,
	0x67, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_publish_proto_rawDescOnce sync.Once
	file_publish_proto_rawDescData = file_publish_proto_rawDesc
)

func file_publish_proto_rawDescGZIP() []byte {
	file_publish_proto_rawDescOnce.Do(func() {
		file_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_publish_proto_rawDescData)
	})
	return file_publish_proto_rawDescData
}

var file_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_publish_proto_goTypes = []interface{}{
	(*PublishRequest)(nil), // 0: gmqtt.admin.api.PublishRequest
	(*UserProperties)(nil), // 1: gmqtt.admin.api.UserProperties
	(*empty.Empty)(nil),    // 2: google.protobuf.Empty
}
var file_publish_proto_depIdxs = []int32{
	1, // 0: gmqtt.admin.api.PublishRequest.user_properties:type_name -> gmqtt.admin.api.UserProperties
	0, // 1: gmqtt.admin.api.PublishService.Publish:input_type -> gmqtt.admin.api.PublishRequest
	2, // 2: gmqtt.admin.api.PublishService.Publish:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_publish_proto_init() }
func file_publish_proto_init() {
	if File_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProperties); i {
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
			RawDescriptor: file_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publish_proto_goTypes,
		DependencyIndexes: file_publish_proto_depIdxs,
		MessageInfos:      file_publish_proto_msgTypes,
	}.Build()
	File_publish_proto = out.File
	file_publish_proto_rawDesc = nil
	file_publish_proto_goTypes = nil
	file_publish_proto_depIdxs = nil
}
