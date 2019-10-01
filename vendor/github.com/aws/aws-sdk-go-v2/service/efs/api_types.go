// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// A description of the file system.
type FileSystemDescription struct {
	_ struct{} `type:"structure"`

	// The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z).
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// The opaque string specified in the request.
	//
	// CreationToken is a required field
	CreationToken *string `min:"1" type:"string" required:"true"`

	// A Boolean value that, if true, indicates that the file system is encrypted.
	Encrypted *bool `type:"boolean"`

	// The ID of the file system, assigned by Amazon EFS.
	//
	// FileSystemId is a required field
	FileSystemId *string `type:"string" required:"true"`

	// The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK)
	// that was used to protect the encrypted file system.
	KmsKeyId *string `min:"1" type:"string"`

	// The lifecycle phase of the file system.
	//
	// LifeCycleState is a required field
	LifeCycleState LifeCycleState `type:"string" required:"true" enum:"true"`

	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem. If the file system has a Name tag, Amazon EFS returns
	// the value in this field.
	Name *string `type:"string"`

	// The current number of mount targets that the file system has. For more information,
	// see CreateMountTarget.
	//
	// NumberOfMountTargets is a required field
	NumberOfMountTargets *int64 `type:"integer" required:"true"`

	// The AWS account that created the file system. If the file system was created
	// by an IAM user, the parent account to which the user belongs is the owner.
	//
	// OwnerId is a required field
	OwnerId *string `type:"string" required:"true"`

	// The performance mode of the file system.
	//
	// PerformanceMode is a required field
	PerformanceMode PerformanceMode `type:"string" required:"true" enum:"true"`

	// The throughput, measured in MiB/s, that you want to provision for a file
	// system. Valid values are 1-1024. Required if ThroughputMode is set to provisioned.
	// The limit on throughput is 1024 MiB/s. You can get these limits increased
	// by contacting AWS Support. For more information, see Amazon EFS Limits That
	// You Can Increase (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits)
	// in the Amazon EFS User Guide.
	ProvisionedThroughputInMibps *float64 `min:"1" type:"double"`

	// The latest known metered size (in bytes) of data stored in the file system,
	// in its Value field, and the time at which that size was determined in its
	// Timestamp field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of
	// a consistent snapshot of the file system, but it is eventually consistent
	// when there are no writes to the file system. That is, SizeInBytes represents
	// actual size only if the file system is not modified for a period longer than
	// a couple of hours. Otherwise, the value is not the exact size that the file
	// system was at any point in time.
	//
	// SizeInBytes is a required field
	SizeInBytes *FileSystemSize `type:"structure" required:"true"`

	// The tags associated with the file system, presented as an array of Tag objects.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`

	// The throughput mode for a file system. There are two throughput modes to
	// choose from for your file system: bursting and provisioned. If you set ThroughputMode
	// to provisioned, you must also set a value for ProvisionedThroughPutInMibps.
	// You can decrease your file system's throughput in Provisioned Throughput
	// mode or change between the throughput modes as long as it’s been more than
	// 24 hours since the last decrease or throughput mode change.
	ThroughputMode ThroughputMode `type:"string" enum:"true"`
}

// String returns the string representation
func (s FileSystemDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s FileSystemDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.CreationToken != nil {
		v := *s.CreationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Encrypted != nil {
		v := *s.Encrypted

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Encrypted", protocol.BoolValue(v), metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KmsKeyId != nil {
		v := *s.KmsKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KmsKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LifeCycleState) > 0 {
		v := s.LifeCycleState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LifeCycleState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NumberOfMountTargets != nil {
		v := *s.NumberOfMountTargets

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NumberOfMountTargets", protocol.Int64Value(v), metadata)
	}
	if s.OwnerId != nil {
		v := *s.OwnerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OwnerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.PerformanceMode) > 0 {
		v := s.PerformanceMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PerformanceMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ProvisionedThroughputInMibps != nil {
		v := *s.ProvisionedThroughputInMibps

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProvisionedThroughputInMibps", protocol.Float64Value(v), metadata)
	}
	if s.SizeInBytes != nil {
		v := s.SizeInBytes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SizeInBytes", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ThroughputMode) > 0 {
		v := s.ThroughputMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThroughputMode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The latest known metered size (in bytes) of data stored in the file system,
// in its Value field, and the time at which that size was determined in its
// Timestamp field. The value doesn't represent the size of a consistent snapshot
// of the file system, but it is eventually consistent when there are no writes
// to the file system. That is, the value represents the actual size only if
// the file system is not modified for a period longer than a couple of hours.
// Otherwise, the value is not necessarily the exact size the file system was
// at any instant in time.
type FileSystemSize struct {
	_ struct{} `type:"structure"`

	// The time at which the size of data, returned in the Value field, was determined.
	// The value is the integer number of seconds since 1970-01-01T00:00:00Z.
	Timestamp *time.Time `type:"timestamp"`

	// The latest known metered size (in bytes) of data stored in the file system.
	//
	// Value is a required field
	Value *int64 `type:"long" required:"true"`

	// The latest known metered size (in bytes) of data stored in the Infrequent
	// Access storage class.
	ValueInIA *int64 `type:"long"`

	// The latest known metered size (in bytes) of data stored in the Standard storage
	// class.
	ValueInStandard *int64 `type:"long"`
}

// String returns the string representation
func (s FileSystemSize) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s FileSystemSize) MarshalFields(e protocol.FieldEncoder) error {
	if s.Timestamp != nil {
		v := *s.Timestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Timestamp",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.Int64Value(v), metadata)
	}
	if s.ValueInIA != nil {
		v := *s.ValueInIA

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ValueInIA", protocol.Int64Value(v), metadata)
	}
	if s.ValueInStandard != nil {
		v := *s.ValueInStandard

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ValueInStandard", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Describes a policy used by EFS lifecycle management to transition files to
// the Infrequent Access (IA) storage class.
type LifecyclePolicy struct {
	_ struct{} `type:"structure"`

	// A value that describes the period of time that a file is not accessed, after
	// which it transitions to the IA storage class. Metadata operations such as
	// listing the contents of a directory don't count as file access events.
	TransitionToIA TransitionToIARules `type:"string" enum:"true"`
}

// String returns the string representation
func (s LifecyclePolicy) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s LifecyclePolicy) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.TransitionToIA) > 0 {
		v := s.TransitionToIA

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TransitionToIA", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Provides a description of a mount target.
type MountTargetDescription struct {
	_ struct{} `type:"structure"`

	// The ID of the file system for which the mount target is intended.
	//
	// FileSystemId is a required field
	FileSystemId *string `type:"string" required:"true"`

	// Address at which the file system can be mounted by using the mount target.
	IpAddress *string `type:"string"`

	// Lifecycle state of the mount target.
	//
	// LifeCycleState is a required field
	LifeCycleState LifeCycleState `type:"string" required:"true" enum:"true"`

	// System-assigned mount target ID.
	//
	// MountTargetId is a required field
	MountTargetId *string `type:"string" required:"true"`

	// The ID of the network interface that Amazon EFS created when it created the
	// mount target.
	NetworkInterfaceId *string `type:"string"`

	// AWS account ID that owns the resource.
	OwnerId *string `type:"string"`

	// The ID of the mount target's subnet.
	//
	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MountTargetDescription) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MountTargetDescription) MarshalFields(e protocol.FieldEncoder) error {
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IpAddress != nil {
		v := *s.IpAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IpAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LifeCycleState) > 0 {
		v := s.LifeCycleState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LifeCycleState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MountTargetId != nil {
		v := *s.MountTargetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MountTargetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkInterfaceId != nil {
		v := *s.NetworkInterfaceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NetworkInterfaceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OwnerId != nil {
		v := *s.OwnerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OwnerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubnetId != nil {
		v := *s.SubnetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SubnetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A tag is a key-value pair. Allowed characters are letters, white space, and
// numbers that can be represented in UTF-8, and the following characters:+
// - = . _ : /
type Tag struct {
	_ struct{} `type:"structure"`

	// The tag key (String). The key can't start with aws:.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The value of the tag key.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Tag) MarshalFields(e protocol.FieldEncoder) error {
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
