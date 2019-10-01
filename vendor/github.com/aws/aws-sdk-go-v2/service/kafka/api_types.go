// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Includes all client authentication information.
type Authentication struct {
	_ struct{} `type:"structure"`

	// Details for ClientAuthentication using TLS.
	Tls *Tls `locationName:"tls" type:"structure"`
}

// String returns the string representation
func (s Authentication) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Authentication) MarshalFields(e protocol.FieldEncoder) error {
	if s.Tls != nil {
		v := s.Tls

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "tls", v, metadata)
	}
	return nil
}

// Specifies the EBS volume upgrade information. The broker identifier must
// be set to the keyword ALL. This means the changes apply to all the brokers
// in the cluster.
type BrokerEBSVolumeInfo struct {
	_ struct{} `type:"structure"`

	// The ID of the broker to update.
	//
	// KafkaBrokerNodeId is a required field
	KafkaBrokerNodeId *string `locationName:"kafkaBrokerNodeId" type:"string" required:"true"`

	// Size of the EBS volume to update.
	//
	// VolumeSizeGB is a required field
	VolumeSizeGB *int64 `locationName:"volumeSizeGB" type:"integer" required:"true"`
}

// String returns the string representation
func (s BrokerEBSVolumeInfo) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BrokerEBSVolumeInfo) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BrokerEBSVolumeInfo"}

	if s.KafkaBrokerNodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KafkaBrokerNodeId"))
	}

	if s.VolumeSizeGB == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeSizeGB"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerEBSVolumeInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.KafkaBrokerNodeId != nil {
		v := *s.KafkaBrokerNodeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "kafkaBrokerNodeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VolumeSizeGB != nil {
		v := *s.VolumeSizeGB

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "volumeSizeGB", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Describes the setup to be used for Kafka broker nodes in the cluster.
type BrokerNodeGroupInfo struct {
	_ struct{} `type:"structure"`

	// The distribution of broker nodes across Availability Zones.
	BrokerAZDistribution BrokerAZDistribution `locationName:"brokerAZDistribution" type:"string" enum:"true"`

	// The list of subnets to connect to in the client virtual private cloud (VPC).
	// AWS creates elastic network interfaces inside these subnets. Client applications
	// use elastic network interfaces to produce and consume data. Client subnets
	// can't be in Availability Zone us-east-1e.
	//
	// ClientSubnets is a required field
	ClientSubnets []string `locationName:"clientSubnets" type:"list" required:"true"`

	// The type of Amazon EC2 instances to use for Kafka brokers. The following
	// instance types are allowed: kafka.m5.large, kafka.m5.xlarge, kafka.m5.2xlarge,kafka.m5.4xlarge,
	// kafka.m5.12xlarge, and kafka.m5.24xlarge.
	//
	// InstanceType is a required field
	InstanceType *string `locationName:"instanceType" min:"5" type:"string" required:"true"`

	// The AWS security groups to associate with the elastic network interfaces
	// in order to specify who can connect to and communicate with the Amazon MSK
	// cluster. If you don't specify a security group, Amazon MSK uses the default
	// security group associated with the VPC.
	SecurityGroups []string `locationName:"securityGroups" type:"list"`

	// Contains information about storage volumes attached to MSK broker nodes.
	StorageInfo *StorageInfo `locationName:"storageInfo" type:"structure"`
}

// String returns the string representation
func (s BrokerNodeGroupInfo) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BrokerNodeGroupInfo) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BrokerNodeGroupInfo"}

	if s.ClientSubnets == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientSubnets"))
	}

	if s.InstanceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceType"))
	}
	if s.InstanceType != nil && len(*s.InstanceType) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceType", 5))
	}
	if s.StorageInfo != nil {
		if err := s.StorageInfo.Validate(); err != nil {
			invalidParams.AddNested("StorageInfo", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerNodeGroupInfo) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.BrokerAZDistribution) > 0 {
		v := s.BrokerAZDistribution

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerAZDistribution", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ClientSubnets != nil {
		v := s.ClientSubnets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "clientSubnets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.InstanceType != nil {
		v := *s.InstanceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "instanceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityGroups != nil {
		v := s.SecurityGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "securityGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.StorageInfo != nil {
		v := s.StorageInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "storageInfo", v, metadata)
	}
	return nil
}

// BrokerNodeInfo
type BrokerNodeInfo struct {
	_ struct{} `type:"structure"`

	// The attached elastic network interface of the broker.
	AttachedENIId *string `locationName:"attachedENIId" type:"string"`

	// The ID of the broker.
	BrokerId *float64 `locationName:"brokerId" type:"double"`

	// The client subnet to which this broker node belongs.
	ClientSubnet *string `locationName:"clientSubnet" type:"string"`

	// The virtual private cloud (VPC) of the client.
	ClientVpcIpAddress *string `locationName:"clientVpcIpAddress" type:"string"`

	// Information about the version of software currently deployed on the Kafka
	// brokers in the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo `locationName:"currentBrokerSoftwareInfo" type:"structure"`

	// Endpoints for accessing the broker.
	Endpoints []string `locationName:"endpoints" type:"list"`
}

// String returns the string representation
func (s BrokerNodeInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerNodeInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttachedENIId != nil {
		v := *s.AttachedENIId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "attachedENIId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerId", protocol.Float64Value(v), metadata)
	}
	if s.ClientSubnet != nil {
		v := *s.ClientSubnet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientSubnet", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientVpcIpAddress != nil {
		v := *s.ClientVpcIpAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientVpcIpAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CurrentBrokerSoftwareInfo != nil {
		v := s.CurrentBrokerSoftwareInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "currentBrokerSoftwareInfo", v, metadata)
	}
	if s.Endpoints != nil {
		v := s.Endpoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "endpoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Information about the current software installed on the cluster.
type BrokerSoftwareInfo struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the configuration used for the cluster.
	// This field isn't visible in this preview release.
	ConfigurationArn *string `locationName:"configurationArn" type:"string"`

	// The revision of the configuration to use. This field isn't visible in this
	// preview release.
	ConfigurationRevision *int64 `locationName:"configurationRevision" type:"long"`

	// The version of Apache Kafka.
	KafkaVersion *string `locationName:"kafkaVersion" type:"string"`
}

// String returns the string representation
func (s BrokerSoftwareInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BrokerSoftwareInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigurationArn != nil {
		v := *s.ConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationRevision != nil {
		v := *s.ConfigurationRevision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configurationRevision", protocol.Int64Value(v), metadata)
	}
	if s.KafkaVersion != nil {
		v := *s.KafkaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "kafkaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Returns information about a cluster.
type ClusterInfo struct {
	_ struct{} `type:"structure"`

	// Arn of active cluster operation.
	ActiveOperationArn *string `locationName:"activeOperationArn" type:"string"`

	// Information about the broker nodes.
	BrokerNodeGroupInfo *BrokerNodeGroupInfo `locationName:"brokerNodeGroupInfo" type:"structure"`

	// Includes all client authentication information.
	ClientAuthentication *Authentication `locationName:"clientAuthentication" type:"structure"`

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The name of the cluster.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The time when the cluster was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601"`

	// Information about the version of software currently deployed on the Kafka
	// brokers in the cluster.
	CurrentBrokerSoftwareInfo *BrokerSoftwareInfo `locationName:"currentBrokerSoftwareInfo" type:"structure"`

	// The current version of the MSK cluster.
	CurrentVersion *string `locationName:"currentVersion" type:"string"`

	// Includes all encryption-related information.
	EncryptionInfo *EncryptionInfo `locationName:"encryptionInfo" type:"structure"`

	// Specifies which metrics are gathered for the MSK cluster. This property has
	// three possible values: DEFAULT, PER_BROKER, and PER_TOPIC_PER_BROKER. For
	// a list of the metrics associated with each of these three levels of monitoring,
	// see Monitoring (https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html).
	EnhancedMonitoring EnhancedMonitoring `locationName:"enhancedMonitoring" type:"string" enum:"true"`

	// The number of broker nodes in the cluster.
	NumberOfBrokerNodes *int64 `locationName:"numberOfBrokerNodes" type:"integer"`

	// The state of the cluster. The possible states are CREATING, ACTIVE, and FAILED.
	State ClusterState `locationName:"state" type:"string" enum:"true"`

	// Tags attached to the cluster.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The connection string to use to connect to the Apache ZooKeeper cluster.
	ZookeeperConnectString *string `locationName:"zookeeperConnectString" type:"string"`
}

// String returns the string representation
func (s ClusterInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ClusterInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActiveOperationArn != nil {
		v := *s.ActiveOperationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "activeOperationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerNodeGroupInfo != nil {
		v := s.BrokerNodeGroupInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "brokerNodeGroupInfo", v, metadata)
	}
	if s.ClientAuthentication != nil {
		v := s.ClientAuthentication

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "clientAuthentication", v, metadata)
	}
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.CurrentBrokerSoftwareInfo != nil {
		v := s.CurrentBrokerSoftwareInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "currentBrokerSoftwareInfo", v, metadata)
	}
	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EncryptionInfo != nil {
		v := s.EncryptionInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryptionInfo", v, metadata)
	}
	if len(s.EnhancedMonitoring) > 0 {
		v := s.EnhancedMonitoring

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enhancedMonitoring", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.NumberOfBrokerNodes != nil {
		v := *s.NumberOfBrokerNodes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "numberOfBrokerNodes", protocol.Int64Value(v), metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ZookeeperConnectString != nil {
		v := *s.ZookeeperConnectString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "zookeeperConnectString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Returns information about a cluster operation.
type ClusterOperationInfo struct {
	_ struct{} `type:"structure"`

	// The ID of the API request that triggered this operation.
	ClientRequestId *string `locationName:"clientRequestId" type:"string"`

	// ARN of the cluster.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The time at which operation was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601"`

	// The time at which the operation finished.
	EndTime *time.Time `locationName:"endTime" type:"timestamp" timestampFormat:"iso8601"`

	// Describes the error if the operation fails.
	ErrorInfo *ErrorInfo `locationName:"errorInfo" type:"structure"`

	// ARN of the cluster operation.
	OperationArn *string `locationName:"operationArn" type:"string"`

	// State of the cluster operation.
	OperationState *string `locationName:"operationState" type:"string"`

	// Type of the cluster operation.
	OperationType *string `locationName:"operationType" type:"string"`

	// Information about cluster attributes before a cluster is updated.
	SourceClusterInfo *MutableClusterInfo `locationName:"sourceClusterInfo" type:"structure"`

	// Information about cluster attributes after a cluster is updated.
	TargetClusterInfo *MutableClusterInfo `locationName:"targetClusterInfo" type:"structure"`
}

// String returns the string representation
func (s ClusterOperationInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ClusterOperationInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientRequestId != nil {
		v := *s.ClientRequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientRequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.ErrorInfo != nil {
		v := s.ErrorInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "errorInfo", v, metadata)
	}
	if s.OperationArn != nil {
		v := *s.OperationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "operationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OperationState != nil {
		v := *s.OperationState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "operationState", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OperationType != nil {
		v := *s.OperationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "operationType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceClusterInfo != nil {
		v := s.SourceClusterInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sourceClusterInfo", v, metadata)
	}
	if s.TargetClusterInfo != nil {
		v := s.TargetClusterInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "targetClusterInfo", v, metadata)
	}
	return nil
}

// Represents an MSK Configuration.
type Configuration struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the configuration.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" type:"string" required:"true"`

	// CreationTime is a required field
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The description of the configuration.
	//
	// Description is a required field
	Description *string `locationName:"description" type:"string" required:"true"`

	// An array of the versions of Apache Kafka with which you can use this MSK
	// configuration. You can use this configuration for an MSK cluster only if
	// the Apache Kafka version specified for the cluster appears in this array.
	//
	// KafkaVersions is a required field
	KafkaVersions []string `locationName:"kafkaVersions" type:"list" required:"true"`

	// Latest revision of the configuration.
	//
	// LatestRevision is a required field
	LatestRevision *ConfigurationRevision `locationName:"latestRevision" type:"structure" required:"true"`

	// The name of the configuration.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s Configuration) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Configuration) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KafkaVersions != nil {
		v := s.KafkaVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "kafkaVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.LatestRevision != nil {
		v := s.LatestRevision

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "latestRevision", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Specifies the configuration to use for the brokers.
type ConfigurationInfo struct {
	_ struct{} `type:"structure"`

	// ARN of the configuration to use.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" type:"string" required:"true"`

	// The revision of the configuration to use.
	//
	// Revision is a required field
	Revision *int64 `locationName:"revision" type:"long" required:"true"`
}

// String returns the string representation
func (s ConfigurationInfo) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfigurationInfo) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfigurationInfo"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if s.Revision == nil {
		invalidParams.Add(aws.NewErrParamRequired("Revision"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigurationInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Describes a configuration revision.
type ConfigurationRevision struct {
	_ struct{} `type:"structure"`

	// The time when the configuration revision was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The description of the configuration revision.
	Description *string `locationName:"description" type:"string"`

	// The revision number.
	//
	// Revision is a required field
	Revision *int64 `locationName:"revision" type:"long" required:"true"`
}

// String returns the string representation
func (s ConfigurationRevision) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigurationRevision) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Revision != nil {
		v := *s.Revision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revision", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Contains information about the EBS storage volumes attached to Kafka broker
// nodes.
type EBSStorageInfo struct {
	_ struct{} `type:"structure"`

	// The size in GiB of the EBS volume for the data drive on each broker node.
	VolumeSize *int64 `locationName:"volumeSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s EBSStorageInfo) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EBSStorageInfo) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EBSStorageInfo"}
	if s.VolumeSize != nil && *s.VolumeSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("VolumeSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EBSStorageInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.VolumeSize != nil {
		v := *s.VolumeSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "volumeSize", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The data-volume encryption details.
type EncryptionAtRest struct {
	_ struct{} `type:"structure"`

	// The ARN of the AWS KMS key for encrypting data at rest. If you don't specify
	// a KMS key, MSK creates one for you and uses it.
	//
	// DataVolumeKMSKeyId is a required field
	DataVolumeKMSKeyId *string `locationName:"dataVolumeKMSKeyId" type:"string" required:"true"`
}

// String returns the string representation
func (s EncryptionAtRest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EncryptionAtRest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EncryptionAtRest"}

	if s.DataVolumeKMSKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataVolumeKMSKeyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EncryptionAtRest) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataVolumeKMSKeyId != nil {
		v := *s.DataVolumeKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dataVolumeKMSKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The settings for encrypting data in transit.
type EncryptionInTransit struct {
	_ struct{} `type:"structure"`

	// Indicates the encryption setting for data in transit between clients and
	// brokers. The following are the possible values.
	//
	// TLS means that client-broker communication is enabled with TLS only.
	//
	// TLS_PLAINTEXT means that client-broker communication is enabled for both
	// TLS-encrypted, as well as plaintext data.
	//
	// PLAINTEXT means that client-broker communication is enabled in plaintext
	// only.
	//
	// The default value is TLS_PLAINTEXT.
	ClientBroker Broker `locationName:"clientBroker" type:"string" enum:"true"`

	// When set to true, it indicates that data communication among the broker nodes
	// of the cluster is encrypted. When set to false, the communication happens
	// in plaintext.
	//
	// The default value is true.
	InCluster *bool `locationName:"inCluster" type:"boolean"`
}

// String returns the string representation
func (s EncryptionInTransit) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EncryptionInTransit) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ClientBroker) > 0 {
		v := s.ClientBroker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientBroker", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.InCluster != nil {
		v := *s.InCluster

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "inCluster", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Includes encryption-related information, such as the AWS KMS key used for
// encrypting data at rest and whether you want MSK to encrypt your data in
// transit.
type EncryptionInfo struct {
	_ struct{} `type:"structure"`

	// The data-volume encryption details.
	EncryptionAtRest *EncryptionAtRest `locationName:"encryptionAtRest" type:"structure"`

	// The details for encryption in transit.
	EncryptionInTransit *EncryptionInTransit `locationName:"encryptionInTransit" type:"structure"`
}

// String returns the string representation
func (s EncryptionInfo) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EncryptionInfo) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EncryptionInfo"}
	if s.EncryptionAtRest != nil {
		if err := s.EncryptionAtRest.Validate(); err != nil {
			invalidParams.AddNested("EncryptionAtRest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EncryptionInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.EncryptionAtRest != nil {
		v := s.EncryptionAtRest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryptionAtRest", v, metadata)
	}
	if s.EncryptionInTransit != nil {
		v := s.EncryptionInTransit

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryptionInTransit", v, metadata)
	}
	return nil
}

// Returns information about an error state of the cluster.
type ErrorInfo struct {
	_ struct{} `type:"structure"`

	// A number describing the error programmatically.
	ErrorCode *string `locationName:"errorCode" type:"string"`

	// An optional field to provide more details about the error.
	ErrorString *string `locationName:"errorString" type:"string"`
}

// String returns the string representation
func (s ErrorInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ErrorInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.ErrorCode != nil {
		v := *s.ErrorCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "errorCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ErrorString != nil {
		v := *s.ErrorString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "errorString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about cluster attributes that can be updated via update APIs.
type MutableClusterInfo struct {
	_ struct{} `type:"structure"`

	// Specifies the size of the EBS volume and the ID of the associated broker.
	BrokerEBSVolumeInfo []BrokerEBSVolumeInfo `locationName:"brokerEBSVolumeInfo" type:"list"`

	// Information about the changes in the configuration of the brokers.
	ConfigurationInfo *ConfigurationInfo `locationName:"configurationInfo" type:"structure"`

	// The number of broker nodes in the cluster.
	NumberOfBrokerNodes *int64 `locationName:"numberOfBrokerNodes" type:"integer"`
}

// String returns the string representation
func (s MutableClusterInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s MutableClusterInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.BrokerEBSVolumeInfo != nil {
		v := s.BrokerEBSVolumeInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "brokerEBSVolumeInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ConfigurationInfo != nil {
		v := s.ConfigurationInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configurationInfo", v, metadata)
	}
	if s.NumberOfBrokerNodes != nil {
		v := *s.NumberOfBrokerNodes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "numberOfBrokerNodes", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The node information object.
type NodeInfo struct {
	_ struct{} `type:"structure"`

	// The start time.
	AddedToClusterTime *string `locationName:"addedToClusterTime" type:"string"`

	// The broker node info.
	BrokerNodeInfo *BrokerNodeInfo `locationName:"brokerNodeInfo" type:"structure"`

	// The instance type.
	InstanceType *string `locationName:"instanceType" type:"string"`

	// The Amazon Resource Name (ARN) of the node.
	NodeARN *string `locationName:"nodeARN" type:"string"`

	// The node type.
	NodeType NodeType `locationName:"nodeType" type:"string" enum:"true"`

	// The ZookeeperNodeInfo.
	ZookeeperNodeInfo *ZookeeperNodeInfo `locationName:"zookeeperNodeInfo" type:"structure"`
}

// String returns the string representation
func (s NodeInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s NodeInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.AddedToClusterTime != nil {
		v := *s.AddedToClusterTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "addedToClusterTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerNodeInfo != nil {
		v := s.BrokerNodeInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "brokerNodeInfo", v, metadata)
	}
	if s.InstanceType != nil {
		v := *s.InstanceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "instanceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodeARN != nil {
		v := *s.NodeARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nodeARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.NodeType) > 0 {
		v := s.NodeType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nodeType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ZookeeperNodeInfo != nil {
		v := s.ZookeeperNodeInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "zookeeperNodeInfo", v, metadata)
	}
	return nil
}

// Contains information about storage volumes attached to MSK broker nodes.
type StorageInfo struct {
	_ struct{} `type:"structure"`

	// EBS volume information.
	EbsStorageInfo *EBSStorageInfo `locationName:"ebsStorageInfo" type:"structure"`
}

// String returns the string representation
func (s StorageInfo) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StorageInfo) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StorageInfo"}
	if s.EbsStorageInfo != nil {
		if err := s.EbsStorageInfo.Validate(); err != nil {
			invalidParams.AddNested("EbsStorageInfo", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StorageInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.EbsStorageInfo != nil {
		v := s.EbsStorageInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ebsStorageInfo", v, metadata)
	}
	return nil
}

// Details for client authentication using TLS.
type Tls struct {
	_ struct{} `type:"structure"`

	// List of ACM Certificate Authority ARNs.
	CertificateAuthorityArnList []string `locationName:"certificateAuthorityArnList" type:"list"`
}

// String returns the string representation
func (s Tls) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Tls) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateAuthorityArnList != nil {
		v := s.CertificateAuthorityArnList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "certificateAuthorityArnList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Zookeeper node information.
type ZookeeperNodeInfo struct {
	_ struct{} `type:"structure"`

	// The attached elastic network interface of the broker.
	AttachedENIId *string `locationName:"attachedENIId" type:"string"`

	// The virtual private cloud (VPC) IP address of the client.
	ClientVpcIpAddress *string `locationName:"clientVpcIpAddress" type:"string"`

	// Endpoints for accessing the ZooKeeper.
	Endpoints []string `locationName:"endpoints" type:"list"`

	// The role-specific ID for Zookeeper.
	ZookeeperId *float64 `locationName:"zookeeperId" type:"double"`

	// The version of Zookeeper.
	ZookeeperVersion *string `locationName:"zookeeperVersion" type:"string"`
}

// String returns the string representation
func (s ZookeeperNodeInfo) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ZookeeperNodeInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttachedENIId != nil {
		v := *s.AttachedENIId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "attachedENIId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientVpcIpAddress != nil {
		v := *s.ClientVpcIpAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientVpcIpAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Endpoints != nil {
		v := s.Endpoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "endpoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ZookeeperId != nil {
		v := *s.ZookeeperId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "zookeeperId", protocol.Float64Value(v), metadata)
	}
	if s.ZookeeperVersion != nil {
		v := *s.ZookeeperVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "zookeeperVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
