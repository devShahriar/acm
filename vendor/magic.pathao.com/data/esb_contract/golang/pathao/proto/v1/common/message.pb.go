// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pathao/proto/v1/common/message.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	EventType string                 `protobuf:"bytes,3,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Event     []byte                 `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pathao_proto_v1_common_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pathao_proto_v1_common_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pathao_proto_v1_common_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Message) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Message) GetEvent() []byte {
	if x != nil {
		return x.Event
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude        float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude       float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Timestamp       int64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	StartNode       int64   `protobuf:"varint,4,opt,name=start_node,json=startNode,proto3" json:"start_node,omitempty"`
	EndNode         int64   `protobuf:"varint,5,opt,name=end_node,json=endNode,proto3" json:"end_node,omitempty"`
	WayId           int64   `protobuf:"varint,6,opt,name=way_id,json=wayId,proto3" json:"way_id,omitempty"`
	DeviceTimestamp int64   `protobuf:"varint,7,opt,name=device_timestamp,json=deviceTimestamp,proto3" json:"device_timestamp,omitempty"`
	GpsTimestamp    int64   `protobuf:"varint,8,opt,name=gps_timestamp,json=gpsTimestamp,proto3" json:"gps_timestamp,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pathao_proto_v1_common_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_pathao_proto_v1_common_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_pathao_proto_v1_common_message_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Location) GetStartNode() int64 {
	if x != nil {
		return x.StartNode
	}
	return 0
}

func (x *Location) GetEndNode() int64 {
	if x != nil {
		return x.EndNode
	}
	return 0
}

func (x *Location) GetWayId() int64 {
	if x != nil {
		return x.WayId
	}
	return 0
}

func (x *Location) GetDeviceTimestamp() int64 {
	if x != nil {
		return x.DeviceTimestamp
	}
	return 0
}

func (x *Location) GetGpsTimestamp() int64 {
	if x != nil {
		return x.GpsTimestamp
	}
	return 0
}

type H3Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level_7 string `protobuf:"bytes,1,opt,name=level_7,json=level7,proto3" json:"level_7,omitempty"`
	Level_8 string `protobuf:"bytes,2,opt,name=level_8,json=level8,proto3" json:"level_8,omitempty"`
	Level_9 string `protobuf:"bytes,3,opt,name=level_9,json=level9,proto3" json:"level_9,omitempty"`
}

func (x *H3Index) Reset() {
	*x = H3Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pathao_proto_v1_common_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H3Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H3Index) ProtoMessage() {}

func (x *H3Index) ProtoReflect() protoreflect.Message {
	mi := &file_pathao_proto_v1_common_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H3Index.ProtoReflect.Descriptor instead.
func (*H3Index) Descriptor() ([]byte, []int) {
	return file_pathao_proto_v1_common_message_proto_rawDescGZIP(), []int{2}
}

func (x *H3Index) GetLevel_7() string {
	if x != nil {
		return x.Level_7
	}
	return ""
}

func (x *H3Index) GetLevel_8() string {
	if x != nil {
		return x.Level_8
	}
	return ""
}

func (x *H3Index) GetLevel_9() string {
	if x != nil {
		return x.Level_9
	}
	return ""
}

type GeoPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *GeoPoint) Reset() {
	*x = GeoPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pathao_proto_v1_common_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoPoint) ProtoMessage() {}

func (x *GeoPoint) ProtoReflect() protoreflect.Message {
	mi := &file_pathao_proto_v1_common_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoPoint.ProtoReflect.Descriptor instead.
func (*GeoPoint) Descriptor() ([]byte, []int) {
	return file_pathao_proto_v1_common_message_proto_rawDescGZIP(), []int{3}
}

func (x *GeoPoint) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GeoPoint) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

var File_pathao_proto_v1_common_message_proto protoreflect.FileDescriptor

var file_pathao_proto_v1_common_message_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x61, 0x74, 0x68, 0x61, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x61, 0x74, 0x68, 0x61, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x98, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x83, 0x02, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x61, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x67,
	0x70, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x67, 0x70, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x54, 0x0a, 0x07, 0x48, 0x33, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x5f, 0x37, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x37, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x38, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x38, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x39, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x39, 0x22, 0x2e, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x61, 0x74, 0x68, 0x61, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pathao_proto_v1_common_message_proto_rawDescOnce sync.Once
	file_pathao_proto_v1_common_message_proto_rawDescData = file_pathao_proto_v1_common_message_proto_rawDesc
)

func file_pathao_proto_v1_common_message_proto_rawDescGZIP() []byte {
	file_pathao_proto_v1_common_message_proto_rawDescOnce.Do(func() {
		file_pathao_proto_v1_common_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pathao_proto_v1_common_message_proto_rawDescData)
	})
	return file_pathao_proto_v1_common_message_proto_rawDescData
}

var file_pathao_proto_v1_common_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pathao_proto_v1_common_message_proto_goTypes = []interface{}{
	(*Message)(nil),               // 0: pathao.proto.v1.common.Message
	(*Location)(nil),              // 1: pathao.proto.v1.common.Location
	(*H3Index)(nil),               // 2: pathao.proto.v1.common.H3Index
	(*GeoPoint)(nil),              // 3: pathao.proto.v1.common.GeoPoint
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_pathao_proto_v1_common_message_proto_depIdxs = []int32{
	4, // 0: pathao.proto.v1.common.Message.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pathao_proto_v1_common_message_proto_init() }
func file_pathao_proto_v1_common_message_proto_init() {
	if File_pathao_proto_v1_common_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pathao_proto_v1_common_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pathao_proto_v1_common_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_pathao_proto_v1_common_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H3Index); i {
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
		file_pathao_proto_v1_common_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoPoint); i {
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
			RawDescriptor: file_pathao_proto_v1_common_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pathao_proto_v1_common_message_proto_goTypes,
		DependencyIndexes: file_pathao_proto_v1_common_message_proto_depIdxs,
		MessageInfos:      file_pathao_proto_v1_common_message_proto_msgTypes,
	}.Build()
	File_pathao_proto_v1_common_message_proto = out.File
	file_pathao_proto_v1_common_message_proto_rawDesc = nil
	file_pathao_proto_v1_common_message_proto_goTypes = nil
	file_pathao_proto_v1_common_message_proto_depIdxs = nil
}
