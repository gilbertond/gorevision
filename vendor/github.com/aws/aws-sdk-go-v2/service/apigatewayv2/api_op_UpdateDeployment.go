// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDeploymentInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// DeploymentId is a required field
	DeploymentId *string `location:"uri" locationName:"deploymentId" type:"string" required:"true"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`
}

// String returns the string representation
func (s UpdateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDeploymentInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"iso8601"`

	// The identifier.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// Represents a deployment status.
	DeploymentStatus DeploymentStatus `locationName:"deploymentStatus" type:"string" enum:"true"`

	DeploymentStatusMessage *string `locationName:"deploymentStatusMessage" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`
}

// String returns the string representation
func (s UpdateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DeploymentStatus) > 0 {
		v := s.DeploymentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DeploymentStatusMessage != nil {
		v := *s.DeploymentStatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentStatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateDeployment = "UpdateDeployment"

// UpdateDeploymentRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates a Deployment.
//
//    // Example sending a request using UpdateDeploymentRequest.
//    req := client.UpdateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateDeployment
func (c *Client) UpdateDeploymentRequest(input *UpdateDeploymentInput) UpdateDeploymentRequest {
	op := &aws.Operation{
		Name:       opUpdateDeployment,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/apis/{apiId}/deployments/{deploymentId}",
	}

	if input == nil {
		input = &UpdateDeploymentInput{}
	}

	req := c.newRequest(op, input, &UpdateDeploymentOutput{})
	return UpdateDeploymentRequest{Request: req, Input: input, Copy: c.UpdateDeploymentRequest}
}

// UpdateDeploymentRequest is the request type for the
// UpdateDeployment API operation.
type UpdateDeploymentRequest struct {
	*aws.Request
	Input *UpdateDeploymentInput
	Copy  func(*UpdateDeploymentInput) UpdateDeploymentRequest
}

// Send marshals and sends the UpdateDeployment API request.
func (r UpdateDeploymentRequest) Send(ctx context.Context) (*UpdateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDeploymentResponse{
		UpdateDeploymentOutput: r.Request.Data.(*UpdateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDeploymentResponse is the response type for the
// UpdateDeployment API operation.
type UpdateDeploymentResponse struct {
	*UpdateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDeployment request.
func (r *UpdateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
