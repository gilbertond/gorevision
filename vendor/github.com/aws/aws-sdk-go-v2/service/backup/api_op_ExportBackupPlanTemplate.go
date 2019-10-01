// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ExportBackupPlanTemplateInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`
}

// String returns the string representation
func (s ExportBackupPlanTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportBackupPlanTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportBackupPlanTemplateInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExportBackupPlanTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ExportBackupPlanTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The body of a backup plan template in JSON format.
	//
	// This is a signed JSON document that cannot be modified before being passed
	// to GetBackupPlanFromJSON.
	BackupPlanTemplateJson *string `type:"string"`
}

// String returns the string representation
func (s ExportBackupPlanTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExportBackupPlanTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlanTemplateJson != nil {
		v := *s.BackupPlanTemplateJson

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanTemplateJson", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opExportBackupPlanTemplate = "ExportBackupPlanTemplate"

// ExportBackupPlanTemplateRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns the backup plan that is specified by the plan ID as a backup template.
//
//    // Example sending a request using ExportBackupPlanTemplateRequest.
//    req := client.ExportBackupPlanTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ExportBackupPlanTemplate
func (c *Client) ExportBackupPlanTemplateRequest(input *ExportBackupPlanTemplateInput) ExportBackupPlanTemplateRequest {
	op := &aws.Operation{
		Name:       opExportBackupPlanTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/toTemplate/",
	}

	if input == nil {
		input = &ExportBackupPlanTemplateInput{}
	}

	req := c.newRequest(op, input, &ExportBackupPlanTemplateOutput{})
	return ExportBackupPlanTemplateRequest{Request: req, Input: input, Copy: c.ExportBackupPlanTemplateRequest}
}

// ExportBackupPlanTemplateRequest is the request type for the
// ExportBackupPlanTemplate API operation.
type ExportBackupPlanTemplateRequest struct {
	*aws.Request
	Input *ExportBackupPlanTemplateInput
	Copy  func(*ExportBackupPlanTemplateInput) ExportBackupPlanTemplateRequest
}

// Send marshals and sends the ExportBackupPlanTemplate API request.
func (r ExportBackupPlanTemplateRequest) Send(ctx context.Context) (*ExportBackupPlanTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportBackupPlanTemplateResponse{
		ExportBackupPlanTemplateOutput: r.Request.Data.(*ExportBackupPlanTemplateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportBackupPlanTemplateResponse is the response type for the
// ExportBackupPlanTemplate API operation.
type ExportBackupPlanTemplateResponse struct {
	*ExportBackupPlanTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportBackupPlanTemplate request.
func (r *ExportBackupPlanTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
