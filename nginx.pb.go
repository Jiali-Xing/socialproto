// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: nginx.proto

package socialproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_nginx_proto protoreflect.FileDescriptor

var file_nginx_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x02, 0x0a, 0x0c, 0x4e, 0x67, 0x69,
	0x6e, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x1f, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x24, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x48, 0x6f, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x69, 0x61, 0x6c, 0x69, 0x2d, 0x58,
	0x69, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_nginx_proto_goTypes = []interface{}{
	(*ComposePostRequest)(nil),       // 0: socialproto.ComposePostRequest
	(*ReadUserTimelineRequest)(nil),  // 1: socialproto.ReadUserTimelineRequest
	(*ReadHomeTimelineRequest)(nil),  // 2: socialproto.ReadHomeTimelineRequest
	(*ComposePostResponse)(nil),      // 3: socialproto.ComposePostResponse
	(*ReadUserTimelineResponse)(nil), // 4: socialproto.ReadUserTimelineResponse
	(*ReadHomeTimelineResponse)(nil), // 5: socialproto.ReadHomeTimelineResponse
}
var file_nginx_proto_depIdxs = []int32{
	0, // 0: socialproto.NginxService.ForwardComposePost:input_type -> socialproto.ComposePostRequest
	1, // 1: socialproto.NginxService.ForwardUserTimeline:input_type -> socialproto.ReadUserTimelineRequest
	2, // 2: socialproto.NginxService.ForwardHomeTimeline:input_type -> socialproto.ReadHomeTimelineRequest
	3, // 3: socialproto.NginxService.ForwardComposePost:output_type -> socialproto.ComposePostResponse
	4, // 4: socialproto.NginxService.ForwardUserTimeline:output_type -> socialproto.ReadUserTimelineResponse
	5, // 5: socialproto.NginxService.ForwardHomeTimeline:output_type -> socialproto.ReadHomeTimelineResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nginx_proto_init() }
func file_nginx_proto_init() {
	if File_nginx_proto != nil {
		return
	}
	file_compose_post_proto_init()
	file_home_timeline_proto_init()
	file_user_timeline_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nginx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nginx_proto_goTypes,
		DependencyIndexes: file_nginx_proto_depIdxs,
	}.Build()
	File_nginx_proto = out.File
	file_nginx_proto_rawDesc = nil
	file_nginx_proto_goTypes = nil
	file_nginx_proto_depIdxs = nil
}
