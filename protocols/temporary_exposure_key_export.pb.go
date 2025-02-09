//https://static.googleusercontent.com/media/www.google.com/en//covid19/exposurenotifications/pdfs/Exposure-Key-File-Format-and-Verification.pdf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: temporary_exposure_key_export.proto

package protocols

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

type TemporaryExposureKeyExport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time window of keys in the file, based on arrival
	// at the server, in UTC seconds.
	StartTimestamp *uint64 `protobuf:"fixed64,1,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	EndTimestamp   *uint64 `protobuf:"fixed64,2,opt,name=end_timestamp,json=endTimestamp" json:"end_timestamp,omitempty"`
	// Region from which these keys came (for example, MCC).
	Region *string `protobuf:"bytes,3,opt,name=region" json:"region,omitempty"`
	// Reserved for future use. Both batch_num and batch_size
	// must be set to a value of 1.
	BatchNum  *int32 `protobuf:"varint,4,opt,name=batch_num,json=batchNum" json:"batch_num,omitempty"`
	BatchSize *int32 `protobuf:"varint,5,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
	// Information about associated signatures.
	SignatureInfos []*SignatureInfo `protobuf:"bytes,6,rep,name=signature_infos,json=signatureInfos" json:"signature_infos,omitempty"`
	// The temporary exposure keys themselves.
	Keys []*TemporaryExposureKey `protobuf:"bytes,7,rep,name=keys" json:"keys,omitempty"`
}

func (x *TemporaryExposureKeyExport) Reset() {
	*x = TemporaryExposureKeyExport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporary_exposure_key_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporaryExposureKeyExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporaryExposureKeyExport) ProtoMessage() {}

func (x *TemporaryExposureKeyExport) ProtoReflect() protoreflect.Message {
	mi := &file_temporary_exposure_key_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporaryExposureKeyExport.ProtoReflect.Descriptor instead.
func (*TemporaryExposureKeyExport) Descriptor() ([]byte, []int) {
	return file_temporary_exposure_key_export_proto_rawDescGZIP(), []int{0}
}

func (x *TemporaryExposureKeyExport) GetStartTimestamp() uint64 {
	if x != nil && x.StartTimestamp != nil {
		return *x.StartTimestamp
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetEndTimestamp() uint64 {
	if x != nil && x.EndTimestamp != nil {
		return *x.EndTimestamp
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetRegion() string {
	if x != nil && x.Region != nil {
		return *x.Region
	}
	return ""
}

func (x *TemporaryExposureKeyExport) GetBatchNum() int32 {
	if x != nil && x.BatchNum != nil {
		return *x.BatchNum
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetBatchSize() int32 {
	if x != nil && x.BatchSize != nil {
		return *x.BatchSize
	}
	return 0
}

func (x *TemporaryExposureKeyExport) GetSignatureInfos() []*SignatureInfo {
	if x != nil {
		return x.SignatureInfos
	}
	return nil
}

func (x *TemporaryExposureKeyExport) GetKeys() []*TemporaryExposureKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type SignatureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// App Store app bundle ID.
	AppBundleId *string `protobuf:"bytes,1,opt,name=app_bundle_id,json=appBundleId" json:"app_bundle_id,omitempty"`
	// Android app package name.
	AndroidPackage *string `protobuf:"bytes,2,opt,name=android_package,json=androidPackage" json:"android_package,omitempty"`
	// Key version in case the EN server signing key is rotated.
	VerificationKeyVersion *string `protobuf:"bytes,3,opt,name=verification_key_version,json=verificationKeyVersion" json:"verification_key_version,omitempty"`
	// Three-digit mobile country code (MCC) for validating the key file.
	// If a region has more than one MCC, the server can choose
	// which MCC to use. This value does not have to match the client's MCC,
	// but must correspond to one of the supported MCCs for its region.
	VerificationKeyId *string `protobuf:"bytes,4,opt,name=verification_key_id,json=verificationKeyId" json:"verification_key_id,omitempty"`
	// All keys must be signed using the SHA-256 with ECDSA algorithm.
	// This field must contain the string "1.2.840.10045.4.3.2".
	SignatureAlgorithm *string `protobuf:"bytes,5,opt,name=signature_algorithm,json=signatureAlgorithm" json:"signature_algorithm,omitempty"`
}

func (x *SignatureInfo) Reset() {
	*x = SignatureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporary_exposure_key_export_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureInfo) ProtoMessage() {}

func (x *SignatureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporary_exposure_key_export_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureInfo.ProtoReflect.Descriptor instead.
func (*SignatureInfo) Descriptor() ([]byte, []int) {
	return file_temporary_exposure_key_export_proto_rawDescGZIP(), []int{1}
}

func (x *SignatureInfo) GetAppBundleId() string {
	if x != nil && x.AppBundleId != nil {
		return *x.AppBundleId
	}
	return ""
}

func (x *SignatureInfo) GetAndroidPackage() string {
	if x != nil && x.AndroidPackage != nil {
		return *x.AndroidPackage
	}
	return ""
}

func (x *SignatureInfo) GetVerificationKeyVersion() string {
	if x != nil && x.VerificationKeyVersion != nil {
		return *x.VerificationKeyVersion
	}
	return ""
}

func (x *SignatureInfo) GetVerificationKeyId() string {
	if x != nil && x.VerificationKeyId != nil {
		return *x.VerificationKeyId
	}
	return ""
}

func (x *SignatureInfo) GetSignatureAlgorithm() string {
	if x != nil && x.SignatureAlgorithm != nil {
		return *x.SignatureAlgorithm
	}
	return ""
}

type TemporaryExposureKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of infected user
	KeyData []byte `protobuf:"bytes,1,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	// Varying risk associated with a key depending on diagnosis method
	TransmissionRiskLevel *int32 `protobuf:"varint,2,opt,name=transmission_risk_level,json=transmissionRiskLevel" json:"transmission_risk_level,omitempty"`
	// The interval number since epoch for which a key starts
	RollingStartIntervalNumber *int32 `protobuf:"varint,3,opt,name=rolling_start_interval_number,json=rollingStartIntervalNumber" json:"rolling_start_interval_number,omitempty"`
	// Increments of 10 minutes describing how long a key is valid
	RollingPeriod *int32 `protobuf:"varint,4,opt,name=rolling_period,json=rollingPeriod,def=144" json:"rolling_period,omitempty"` // defaults to 24 hours
	// Type of diagnosis associated with a key.
	ReportType *ReportType `protobuf:"varint,5,opt,name=report_type,json=reportType,enum=protocols.ReportType,def=1" json:"report_type,omitempty"`
	// Number of days elapsed between symptom onset and the TEK being used.
	// E.g. 2 means TEK is 2 days after onset of symptoms.
	DaysSinceOnsetOfSymptoms *int32 `protobuf:"zigzag32,6,opt,name=days_since_onset_of_symptoms,json=daysSinceOnsetOfSymptoms" json:"days_since_onset_of_symptoms,omitempty"`
}

// Default values for TemporaryExposureKey fields.
const (
	Default_TemporaryExposureKey_RollingPeriod = int32(144)
	Default_TemporaryExposureKey_ReportType    = ReportType_CONFIRMED_TEST
)

func (x *TemporaryExposureKey) Reset() {
	*x = TemporaryExposureKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporary_exposure_key_export_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporaryExposureKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporaryExposureKey) ProtoMessage() {}

func (x *TemporaryExposureKey) ProtoReflect() protoreflect.Message {
	mi := &file_temporary_exposure_key_export_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporaryExposureKey.ProtoReflect.Descriptor instead.
func (*TemporaryExposureKey) Descriptor() ([]byte, []int) {
	return file_temporary_exposure_key_export_proto_rawDescGZIP(), []int{2}
}

func (x *TemporaryExposureKey) GetKeyData() []byte {
	if x != nil {
		return x.KeyData
	}
	return nil
}

func (x *TemporaryExposureKey) GetTransmissionRiskLevel() int32 {
	if x != nil && x.TransmissionRiskLevel != nil {
		return *x.TransmissionRiskLevel
	}
	return 0
}

func (x *TemporaryExposureKey) GetRollingStartIntervalNumber() int32 {
	if x != nil && x.RollingStartIntervalNumber != nil {
		return *x.RollingStartIntervalNumber
	}
	return 0
}

func (x *TemporaryExposureKey) GetRollingPeriod() int32 {
	if x != nil && x.RollingPeriod != nil {
		return *x.RollingPeriod
	}
	return Default_TemporaryExposureKey_RollingPeriod
}

func (x *TemporaryExposureKey) GetReportType() ReportType {
	if x != nil && x.ReportType != nil {
		return *x.ReportType
	}
	return Default_TemporaryExposureKey_ReportType
}

func (x *TemporaryExposureKey) GetDaysSinceOnsetOfSymptoms() int32 {
	if x != nil && x.DaysSinceOnsetOfSymptoms != nil {
		return *x.DaysSinceOnsetOfSymptoms
	}
	return 0
}

var File_temporary_exposure_key_export_proto protoreflect.FileDescriptor

var file_temporary_exposure_key_export_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x70, 0x6f,
	0x73, 0x75, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x1a, 0x19, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x69, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x1a,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x06, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x41, 0x0a, 0x0f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x33, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x22, 0xf7, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x5f, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x70, 0x70, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x13, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x22, 0xe0,
	0x02, 0x0a, 0x14, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6f,
	0x73, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x36, 0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x69, 0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x69, 0x73, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x6f,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x1a, 0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x0e, 0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x31, 0x34, 0x34, 0x52, 0x0d, 0x72, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x46, 0x0a, 0x0b, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x0e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f,
	0x6f, 0x6e, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52, 0x18, 0x64, 0x61, 0x79, 0x73, 0x53, 0x69, 0x6e,
	0x63, 0x65, 0x4f, 0x6e, 0x73, 0x65, 0x74, 0x4f, 0x66, 0x53, 0x79, 0x6d, 0x70, 0x74, 0x6f, 0x6d,
	0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
}

var (
	file_temporary_exposure_key_export_proto_rawDescOnce sync.Once
	file_temporary_exposure_key_export_proto_rawDescData = file_temporary_exposure_key_export_proto_rawDesc
)

func file_temporary_exposure_key_export_proto_rawDescGZIP() []byte {
	file_temporary_exposure_key_export_proto_rawDescOnce.Do(func() {
		file_temporary_exposure_key_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporary_exposure_key_export_proto_rawDescData)
	})
	return file_temporary_exposure_key_export_proto_rawDescData
}

var file_temporary_exposure_key_export_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporary_exposure_key_export_proto_goTypes = []interface{}{
	(*TemporaryExposureKeyExport)(nil), // 0: protocols.TemporaryExposureKeyExport
	(*SignatureInfo)(nil),              // 1: protocols.SignatureInfo
	(*TemporaryExposureKey)(nil),       // 2: protocols.TemporaryExposureKey
	(ReportType)(0),                    // 3: protocols.ReportType
}
var file_temporary_exposure_key_export_proto_depIdxs = []int32{
	1, // 0: protocols.TemporaryExposureKeyExport.signature_infos:type_name -> protocols.SignatureInfo
	2, // 1: protocols.TemporaryExposureKeyExport.keys:type_name -> protocols.TemporaryExposureKey
	3, // 2: protocols.TemporaryExposureKey.report_type:type_name -> protocols.ReportType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_temporary_exposure_key_export_proto_init() }
func file_temporary_exposure_key_export_proto_init() {
	if File_temporary_exposure_key_export_proto != nil {
		return
	}
	file_diagnosis_key_batch_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_temporary_exposure_key_export_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporaryExposureKeyExport); i {
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
		file_temporary_exposure_key_export_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureInfo); i {
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
		file_temporary_exposure_key_export_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemporaryExposureKey); i {
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
			RawDescriptor: file_temporary_exposure_key_export_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporary_exposure_key_export_proto_goTypes,
		DependencyIndexes: file_temporary_exposure_key_export_proto_depIdxs,
		MessageInfos:      file_temporary_exposure_key_export_proto_msgTypes,
	}.Build()
	File_temporary_exposure_key_export_proto = out.File
	file_temporary_exposure_key_export_proto_rawDesc = nil
	file_temporary_exposure_key_export_proto_goTypes = nil
	file_temporary_exposure_key_export_proto_depIdxs = nil
}
