// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: validator/v1/validator.proto

package v1

import (
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

type AverageOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId         string  `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ContractThreshold float64 `protobuf:"fixed64,2,opt,name=contract_threshold,json=contractThreshold,proto3" json:"contract_threshold,omitempty"`
	Baseline          float64 `protobuf:"fixed64,3,opt,name=baseline,proto3" json:"baseline,omitempty"`
	AverageOutput     float64 `protobuf:"fixed64,4,opt,name=average_output,json=averageOutput,proto3" json:"average_output,omitempty"`
	StartTime         string  `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime           string  `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *AverageOutput) Reset() {
	*x = AverageOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_v1_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageOutput) ProtoMessage() {}

func (x *AverageOutput) ProtoReflect() protoreflect.Message {
	mi := &file_validator_v1_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageOutput.ProtoReflect.Descriptor instead.
func (*AverageOutput) Descriptor() ([]byte, []int) {
	return file_validator_v1_validator_proto_rawDescGZIP(), []int{0}
}

func (x *AverageOutput) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *AverageOutput) GetContractThreshold() float64 {
	if x != nil {
		return x.ContractThreshold
	}
	return 0
}

func (x *AverageOutput) GetBaseline() float64 {
	if x != nil {
		return x.Baseline
	}
	return 0
}

func (x *AverageOutput) GetAverageOutput() float64 {
	if x != nil {
		return x.AverageOutput
	}
	return 0
}

func (x *AverageOutput) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AverageOutput) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ValidateAverageOutputsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AverageOutputs []*AverageOutput `protobuf:"bytes,1,rep,name=average_outputs,json=averageOutputs,proto3" json:"average_outputs,omitempty"`
}

func (x *ValidateAverageOutputsRequest) Reset() {
	*x = ValidateAverageOutputsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_v1_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAverageOutputsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAverageOutputsRequest) ProtoMessage() {}

func (x *ValidateAverageOutputsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_validator_v1_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAverageOutputsRequest.ProtoReflect.Descriptor instead.
func (*ValidateAverageOutputsRequest) Descriptor() ([]byte, []int) {
	return file_validator_v1_validator_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateAverageOutputsRequest) GetAverageOutputs() []*AverageOutput {
	if x != nil {
		return x.AverageOutputs
	}
	return nil
}

type ValidateAverageOutputsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	NotValid bool               `protobuf:"varint,2,opt,name=not_valid,json=notValid,proto3" json:"not_valid,omitempty"`
	Errors   []*ValidationError `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ValidateAverageOutputsResponse) Reset() {
	*x = ValidateAverageOutputsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_v1_validator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAverageOutputsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAverageOutputsResponse) ProtoMessage() {}

func (x *ValidateAverageOutputsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validator_v1_validator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAverageOutputsResponse.ProtoReflect.Descriptor instead.
func (*ValidateAverageOutputsResponse) Descriptor() ([]byte, []int) {
	return file_validator_v1_validator_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateAverageOutputsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ValidateAverageOutputsResponse) GetNotValid() bool {
	if x != nil {
		return x.NotValid
	}
	return false
}

func (x *ValidateAverageOutputsResponse) GetErrors() []*ValidationError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type ValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ValidationError) Reset() {
	*x = ValidationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_v1_validator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationError) ProtoMessage() {}

func (x *ValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_validator_v1_validator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationError.ProtoReflect.Descriptor instead.
func (*ValidationError) Descriptor() ([]byte, []int) {
	return file_validator_v1_validator_proto_rawDescGZIP(), []int{3}
}

func (x *ValidationError) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ValidationError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_validator_v1_validator_proto protoreflect.FileDescriptor

var file_validator_v1_validator_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xda, 0x01, 0x0a,
	0x0d, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x1d, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0f, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x0e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x22, 0x8e, 0x01, 0x0a, 0x1e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x6e, 0x6f, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x22, 0x4a, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x89, 0x01,
	0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x75, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x69, 0x64, 0x2d, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x69, 0x64, 0x2d, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_validator_v1_validator_proto_rawDescOnce sync.Once
	file_validator_v1_validator_proto_rawDescData = file_validator_v1_validator_proto_rawDesc
)

func file_validator_v1_validator_proto_rawDescGZIP() []byte {
	file_validator_v1_validator_proto_rawDescOnce.Do(func() {
		file_validator_v1_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_validator_v1_validator_proto_rawDescData)
	})
	return file_validator_v1_validator_proto_rawDescData
}

var file_validator_v1_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_validator_v1_validator_proto_goTypes = []interface{}{
	(*AverageOutput)(nil),                  // 0: validator.v1.AverageOutput
	(*ValidateAverageOutputsRequest)(nil),  // 1: validator.v1.ValidateAverageOutputsRequest
	(*ValidateAverageOutputsResponse)(nil), // 2: validator.v1.ValidateAverageOutputsResponse
	(*ValidationError)(nil),                // 3: validator.v1.ValidationError
}
var file_validator_v1_validator_proto_depIdxs = []int32{
	0, // 0: validator.v1.ValidateAverageOutputsRequest.average_outputs:type_name -> validator.v1.AverageOutput
	3, // 1: validator.v1.ValidateAverageOutputsResponse.errors:type_name -> validator.v1.ValidationError
	1, // 2: validator.v1.ValidatorService.ValidateAverageOutputs:input_type -> validator.v1.ValidateAverageOutputsRequest
	2, // 3: validator.v1.ValidatorService.ValidateAverageOutputs:output_type -> validator.v1.ValidateAverageOutputsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_validator_v1_validator_proto_init() }
func file_validator_v1_validator_proto_init() {
	if File_validator_v1_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validator_v1_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageOutput); i {
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
		file_validator_v1_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAverageOutputsRequest); i {
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
		file_validator_v1_validator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAverageOutputsResponse); i {
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
		file_validator_v1_validator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationError); i {
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
			RawDescriptor: file_validator_v1_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_validator_v1_validator_proto_goTypes,
		DependencyIndexes: file_validator_v1_validator_proto_depIdxs,
		MessageInfos:      file_validator_v1_validator_proto_msgTypes,
	}.Build()
	File_validator_v1_validator_proto = out.File
	file_validator_v1_validator_proto_rawDesc = nil
	file_validator_v1_validator_proto_goTypes = nil
	file_validator_v1_validator_proto_depIdxs = nil
}
