// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBackupVaultNotificationsInput struct {
	_ struct{} `type:"structure"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `location:"uri" locationName:"backupVaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBackupVaultNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackupVaultNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBackupVaultNotificationsInput"}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupVaultNotificationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetBackupVaultNotificationsOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string `type:"string"`

	// An array of events that indicate the status of jobs to back up resources
	// to the backup vault.
	BackupVaultEvents []BackupVaultEvent `type:"list"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string `type:"string"`

	// An ARN that uniquely identifies an Amazon Simple Notification Service (Amazon
	// SNS) topic; for example, arn:aws:sns:us-west-2:111122223333:MyTopic.
	SNSTopicArn *string `type:"string"`
}

// String returns the string representation
func (s GetBackupVaultNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupVaultNotificationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupVaultArn != nil {
		v := *s.BackupVaultArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupVaultEvents != nil {
		v := s.BackupVaultEvents

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BackupVaultEvents", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SNSTopicArn != nil {
		v := *s.SNSTopicArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SNSTopicArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBackupVaultNotifications = "GetBackupVaultNotifications"

// GetBackupVaultNotificationsRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns event notifications for the specified backup vault.
//
//    // Example sending a request using GetBackupVaultNotificationsRequest.
//    req := client.GetBackupVaultNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupVaultNotifications
func (c *Client) GetBackupVaultNotificationsRequest(input *GetBackupVaultNotificationsInput) GetBackupVaultNotificationsRequest {
	op := &aws.Operation{
		Name:       opGetBackupVaultNotifications,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-vaults/{backupVaultName}/notification-configuration",
	}

	if input == nil {
		input = &GetBackupVaultNotificationsInput{}
	}

	req := c.newRequest(op, input, &GetBackupVaultNotificationsOutput{})
	return GetBackupVaultNotificationsRequest{Request: req, Input: input, Copy: c.GetBackupVaultNotificationsRequest}
}

// GetBackupVaultNotificationsRequest is the request type for the
// GetBackupVaultNotifications API operation.
type GetBackupVaultNotificationsRequest struct {
	*aws.Request
	Input *GetBackupVaultNotificationsInput
	Copy  func(*GetBackupVaultNotificationsInput) GetBackupVaultNotificationsRequest
}

// Send marshals and sends the GetBackupVaultNotifications API request.
func (r GetBackupVaultNotificationsRequest) Send(ctx context.Context) (*GetBackupVaultNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupVaultNotificationsResponse{
		GetBackupVaultNotificationsOutput: r.Request.Data.(*GetBackupVaultNotificationsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupVaultNotificationsResponse is the response type for the
// GetBackupVaultNotifications API operation.
type GetBackupVaultNotificationsResponse struct {
	*GetBackupVaultNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupVaultNotifications request.
func (r *GetBackupVaultNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
