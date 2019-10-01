// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteFlowInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFlowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFlowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFlowInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFlowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful DeleteFlow request.
type DeleteFlowOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that was deleted.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The status of the flow when the DeleteFlow process begins.
	Status Status `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteFlowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFlowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDeleteFlow = "DeleteFlow"

// DeleteFlowRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Deletes a flow. Before you can delete a flow, you must stop the flow.
//
//    // Example sending a request using DeleteFlowRequest.
//    req := client.DeleteFlowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/DeleteFlow
func (c *Client) DeleteFlowRequest(input *DeleteFlowInput) DeleteFlowRequest {
	op := &aws.Operation{
		Name:       opDeleteFlow,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/flows/{flowArn}",
	}

	if input == nil {
		input = &DeleteFlowInput{}
	}

	req := c.newRequest(op, input, &DeleteFlowOutput{})
	return DeleteFlowRequest{Request: req, Input: input, Copy: c.DeleteFlowRequest}
}

// DeleteFlowRequest is the request type for the
// DeleteFlow API operation.
type DeleteFlowRequest struct {
	*aws.Request
	Input *DeleteFlowInput
	Copy  func(*DeleteFlowInput) DeleteFlowRequest
}

// Send marshals and sends the DeleteFlow API request.
func (r DeleteFlowRequest) Send(ctx context.Context) (*DeleteFlowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFlowResponse{
		DeleteFlowOutput: r.Request.Data.(*DeleteFlowOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFlowResponse is the response type for the
// DeleteFlow API operation.
type DeleteFlowResponse struct {
	*DeleteFlowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFlow request.
func (r *DeleteFlowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
