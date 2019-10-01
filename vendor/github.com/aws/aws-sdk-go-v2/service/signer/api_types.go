// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Points to an S3Destination object that contains information about your S3
// bucket.
type Destination struct {
	_ struct{} `type:"structure"`

	// The S3Destination object.
	S3 *S3Destination `locationName:"s3" type:"structure"`
}

// String returns the string representation
func (s Destination) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Destination) MarshalFields(e protocol.FieldEncoder) error {
	if s.S3 != nil {
		v := s.S3

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "s3", v, metadata)
	}
	return nil
}

// The encryption algorithm options that are available to an AWS Signer job.
type EncryptionAlgorithmOptions struct {
	_ struct{} `type:"structure"`

	// The set of accepted encryption algorithms that are allowed in an AWS Signer
	// job.
	//
	// AllowedValues is a required field
	AllowedValues []EncryptionAlgorithm `locationName:"allowedValues" type:"list" required:"true"`

	// The default encryption algorithm that is used by an AWS Signer job.
	//
	// DefaultValue is a required field
	DefaultValue EncryptionAlgorithm `locationName:"defaultValue" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s EncryptionAlgorithmOptions) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EncryptionAlgorithmOptions) MarshalFields(e protocol.FieldEncoder) error {
	if s.AllowedValues != nil {
		v := s.AllowedValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "allowedValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.DefaultValue) > 0 {
		v := s.DefaultValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultValue", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The hash algorithms that are available to an AWS Signer job.
type HashAlgorithmOptions struct {
	_ struct{} `type:"structure"`

	// The set of accepted hash algorithms allowed in an AWS Signer job.
	//
	// AllowedValues is a required field
	AllowedValues []HashAlgorithm `locationName:"allowedValues" type:"list" required:"true"`

	// The default hash algorithm that is used in an AWS Signer job.
	//
	// DefaultValue is a required field
	DefaultValue HashAlgorithm `locationName:"defaultValue" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s HashAlgorithmOptions) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HashAlgorithmOptions) MarshalFields(e protocol.FieldEncoder) error {
	if s.AllowedValues != nil {
		v := s.AllowedValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "allowedValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.DefaultValue) > 0 {
		v := s.DefaultValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultValue", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The name and prefix of the S3 bucket where AWS Signer saves your signed objects.
type S3Destination struct {
	_ struct{} `type:"structure"`

	// Name of the S3 bucket.
	BucketName *string `locationName:"bucketName" type:"string"`

	// An Amazon S3 prefix that you can use to limit responses to those that begin
	// with the specified prefix.
	Prefix *string `locationName:"prefix" type:"string"`
}

// String returns the string representation
func (s S3Destination) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s S3Destination) MarshalFields(e protocol.FieldEncoder) error {
	if s.BucketName != nil {
		v := *s.BucketName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "bucketName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "prefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The S3 bucket name and key where AWS Signer saved your signed code image.
type S3SignedObject struct {
	_ struct{} `type:"structure"`

	// Name of the S3 bucket.
	BucketName *string `locationName:"bucketName" type:"string"`

	// Key name that uniquely identifies a signed code image in your bucket.
	Key *string `locationName:"key" type:"string"`
}

// String returns the string representation
func (s S3SignedObject) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s S3SignedObject) MarshalFields(e protocol.FieldEncoder) error {
	if s.BucketName != nil {
		v := *s.BucketName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "bucketName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the S3 bucket where you saved your unsigned code.
type S3Source struct {
	_ struct{} `type:"structure"`

	// Name of the S3 bucket.
	//
	// BucketName is a required field
	BucketName *string `locationName:"bucketName" type:"string" required:"true"`

	// Key name of the bucket object that contains your unsigned code.
	//
	// Key is a required field
	Key *string `locationName:"key" type:"string" required:"true"`

	// Version of your source image in your version enabled S3 bucket.
	//
	// Version is a required field
	Version *string `locationName:"version" type:"string" required:"true"`
}

// String returns the string representation
func (s S3Source) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *S3Source) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "S3Source"}

	if s.BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketName"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s S3Source) MarshalFields(e protocol.FieldEncoder) error {
	if s.BucketName != nil {
		v := *s.BucketName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "bucketName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Points to an S3SignedObject object that contains information about your signed
// code image.
type SignedObject struct {
	_ struct{} `type:"structure"`

	// The S3SignedObject.
	S3 *S3SignedObject `locationName:"s3" type:"structure"`
}

// String returns the string representation
func (s SignedObject) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SignedObject) MarshalFields(e protocol.FieldEncoder) error {
	if s.S3 != nil {
		v := s.S3

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "s3", v, metadata)
	}
	return nil
}

// The configuration of an AWS Signer operation.
type SigningConfiguration struct {
	_ struct{} `type:"structure"`

	// The encryption algorithm options that are available for an AWS Signer job.
	//
	// EncryptionAlgorithmOptions is a required field
	EncryptionAlgorithmOptions *EncryptionAlgorithmOptions `locationName:"encryptionAlgorithmOptions" type:"structure" required:"true"`

	// The hash algorithm options that are available for an AWS Signer job.
	//
	// HashAlgorithmOptions is a required field
	HashAlgorithmOptions *HashAlgorithmOptions `locationName:"hashAlgorithmOptions" type:"structure" required:"true"`
}

// String returns the string representation
func (s SigningConfiguration) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningConfiguration) MarshalFields(e protocol.FieldEncoder) error {
	if s.EncryptionAlgorithmOptions != nil {
		v := s.EncryptionAlgorithmOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryptionAlgorithmOptions", v, metadata)
	}
	if s.HashAlgorithmOptions != nil {
		v := s.HashAlgorithmOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hashAlgorithmOptions", v, metadata)
	}
	return nil
}

// A signing configuration that overrides the default encryption or hash algorithm
// of a signing job.
type SigningConfigurationOverrides struct {
	_ struct{} `type:"structure"`

	// A specified override of the default encryption algorithm that is used in
	// an AWS Signer job.
	EncryptionAlgorithm EncryptionAlgorithm `locationName:"encryptionAlgorithm" type:"string" enum:"true"`

	// A specified override of the default hash algorithm that is used in an AWS
	// Signer job.
	HashAlgorithm HashAlgorithm `locationName:"hashAlgorithm" type:"string" enum:"true"`
}

// String returns the string representation
func (s SigningConfigurationOverrides) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningConfigurationOverrides) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.EncryptionAlgorithm) > 0 {
		v := s.EncryptionAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "encryptionAlgorithm", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.HashAlgorithm) > 0 {
		v := s.HashAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "hashAlgorithm", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The image format of an AWS Signer platform or profile.
type SigningImageFormat struct {
	_ struct{} `type:"structure"`

	// The default format of an AWS Signer signing image.
	//
	// DefaultFormat is a required field
	DefaultFormat ImageFormat `locationName:"defaultFormat" type:"string" required:"true" enum:"true"`

	// The supported formats of an AWS Signer signing image.
	//
	// SupportedFormats is a required field
	SupportedFormats []ImageFormat `locationName:"supportedFormats" type:"list" required:"true"`
}

// String returns the string representation
func (s SigningImageFormat) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningImageFormat) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DefaultFormat) > 0 {
		v := s.DefaultFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SupportedFormats != nil {
		v := s.SupportedFormats

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "supportedFormats", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Contains information about a signing job.
type SigningJob struct {
	_ struct{} `type:"structure"`

	// The date and time that the signing job was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The ID of the signing job.
	JobId *string `locationName:"jobId" type:"string"`

	// A SignedObject structure that contains information about a signing job's
	// signed code image.
	SignedObject *SignedObject `locationName:"signedObject" type:"structure"`

	// A SigningMaterial object that contains the Amazon Resource Name (ARN) of
	// the certificate used for the signing job.
	SigningMaterial *SigningMaterial `locationName:"signingMaterial" type:"structure"`

	// A Source that contains information about a signing job's code image source.
	Source *Source `locationName:"source" type:"structure"`

	// The status of the signing job.
	Status SigningStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s SigningJob) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningJob) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SignedObject != nil {
		v := s.SignedObject

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signedObject", v, metadata)
	}
	if s.SigningMaterial != nil {
		v := s.SigningMaterial

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingMaterial", v, metadata)
	}
	if s.Source != nil {
		v := s.Source

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "source", v, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// The ACM certificate that is used to sign your code.
type SigningMaterial struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the certificates that is used to sign your
	// code.
	//
	// CertificateArn is a required field
	CertificateArn *string `locationName:"certificateArn" type:"string" required:"true"`
}

// String returns the string representation
func (s SigningMaterial) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SigningMaterial) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SigningMaterial"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningMaterial) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateArn != nil {
		v := *s.CertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains information about the signing configurations and parameters that
// is used to perform an AWS Signer job.
type SigningPlatform struct {
	_ struct{} `type:"structure"`

	// The category of an AWS Signer platform.
	Category Category `locationName:"category" type:"string" enum:"true"`

	// The display name of an AWS Signer platform.
	DisplayName *string `locationName:"displayName" type:"string"`

	// The maximum size (in MB) of code that can be signed by a AWS Signer platform.
	MaxSizeInMB *int64 `locationName:"maxSizeInMB" type:"integer"`

	// Any partner entities linked to an AWS Signer platform.
	Partner *string `locationName:"partner" type:"string"`

	// The ID of an AWS Signer platform.
	PlatformId *string `locationName:"platformId" type:"string"`

	// The configuration of an AWS Signer platform. This includes the designated
	// hash algorithm and encryption algorithm of a signing platform.
	SigningConfiguration *SigningConfiguration `locationName:"signingConfiguration" type:"structure"`

	// The signing image format that is used by an AWS Signer platform.
	SigningImageFormat *SigningImageFormat `locationName:"signingImageFormat" type:"structure"`

	// The types of targets that can be signed by an AWS Signer platform.
	Target *string `locationName:"target" type:"string"`
}

// String returns the string representation
func (s SigningPlatform) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningPlatform) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Category) > 0 {
		v := s.Category

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "category", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DisplayName != nil {
		v := *s.DisplayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "displayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxSizeInMB != nil {
		v := *s.MaxSizeInMB

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxSizeInMB", protocol.Int64Value(v), metadata)
	}
	if s.Partner != nil {
		v := *s.Partner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "partner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlatformId != nil {
		v := *s.PlatformId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platformId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SigningConfiguration != nil {
		v := s.SigningConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingConfiguration", v, metadata)
	}
	if s.SigningImageFormat != nil {
		v := s.SigningImageFormat

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingImageFormat", v, metadata)
	}
	if s.Target != nil {
		v := *s.Target

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "target", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Any overrides that are applied to the signing configuration of an AWS Signer
// platform.
type SigningPlatformOverrides struct {
	_ struct{} `type:"structure"`

	// A signing configuration that overrides the default encryption or hash algorithm
	// of a signing job.
	SigningConfiguration *SigningConfigurationOverrides `locationName:"signingConfiguration" type:"structure"`
}

// String returns the string representation
func (s SigningPlatformOverrides) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningPlatformOverrides) MarshalFields(e protocol.FieldEncoder) error {
	if s.SigningConfiguration != nil {
		v := s.SigningConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingConfiguration", v, metadata)
	}
	return nil
}

// Contains information about the ACM certificates and AWS Signer configuration
// parameters that can be used by a given AWS Signer user.
type SigningProfile struct {
	_ struct{} `type:"structure"`

	// The ID of a platform that is available for use by a signing profile.
	PlatformId *string `locationName:"platformId" type:"string"`

	// The name of the AWS Signer profile.
	ProfileName *string `locationName:"profileName" min:"2" type:"string"`

	// The ACM certificate that is available for use by a signing profile.
	SigningMaterial *SigningMaterial `locationName:"signingMaterial" type:"structure"`

	// The parameters that are available for use by an AWS Signer user.
	SigningParameters map[string]string `locationName:"signingParameters" type:"map"`

	// The status of an AWS Signer profile.
	Status SigningProfileStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s SigningProfile) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SigningProfile) MarshalFields(e protocol.FieldEncoder) error {
	if s.PlatformId != nil {
		v := *s.PlatformId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "platformId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfileName != nil {
		v := *s.ProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "profileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SigningMaterial != nil {
		v := s.SigningMaterial

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "signingMaterial", v, metadata)
	}
	if s.SigningParameters != nil {
		v := s.SigningParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "signingParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// An S3Source object that contains information about the S3 bucket where you
// saved your unsigned code.
type Source struct {
	_ struct{} `type:"structure"`

	// The S3Source object.
	S3 *S3Source `locationName:"s3" type:"structure"`
}

// String returns the string representation
func (s Source) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Source) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Source"}
	if s.S3 != nil {
		if err := s.S3.Validate(); err != nil {
			invalidParams.AddNested("S3", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Source) MarshalFields(e protocol.FieldEncoder) error {
	if s.S3 != nil {
		v := s.S3

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "s3", v, metadata)
	}
	return nil
}
