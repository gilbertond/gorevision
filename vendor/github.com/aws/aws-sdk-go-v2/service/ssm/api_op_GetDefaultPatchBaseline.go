// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDefaultPatchBaselineInput struct {
	_ struct{} `type:"structure"`

	// Returns the default patch baseline for the specified operating system.
	OperatingSystem OperatingSystem `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetDefaultPatchBaselineInput) String() string {
	return awsutil.Prettify(s)
}

type GetDefaultPatchBaselineOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the default patch baseline.
	BaselineId *string `min:"20" type:"string"`

	// The operating system for the returned patch baseline.
	OperatingSystem OperatingSystem `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetDefaultPatchBaselineOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDefaultPatchBaseline = "GetDefaultPatchBaseline"

// GetDefaultPatchBaselineRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the default patch baseline. Note that Systems Manager supports
// creating multiple default patch baselines. For example, you can create a
// default patch baseline for each operating system.
//
// If you do not specify an operating system value, the default patch baseline
// for Windows is returned.
//
//    // Example sending a request using GetDefaultPatchBaselineRequest.
//    req := client.GetDefaultPatchBaselineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetDefaultPatchBaseline
func (c *Client) GetDefaultPatchBaselineRequest(input *GetDefaultPatchBaselineInput) GetDefaultPatchBaselineRequest {
	op := &aws.Operation{
		Name:       opGetDefaultPatchBaseline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDefaultPatchBaselineInput{}
	}

	req := c.newRequest(op, input, &GetDefaultPatchBaselineOutput{})
	return GetDefaultPatchBaselineRequest{Request: req, Input: input, Copy: c.GetDefaultPatchBaselineRequest}
}

// GetDefaultPatchBaselineRequest is the request type for the
// GetDefaultPatchBaseline API operation.
type GetDefaultPatchBaselineRequest struct {
	*aws.Request
	Input *GetDefaultPatchBaselineInput
	Copy  func(*GetDefaultPatchBaselineInput) GetDefaultPatchBaselineRequest
}

// Send marshals and sends the GetDefaultPatchBaseline API request.
func (r GetDefaultPatchBaselineRequest) Send(ctx context.Context) (*GetDefaultPatchBaselineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDefaultPatchBaselineResponse{
		GetDefaultPatchBaselineOutput: r.Request.Data.(*GetDefaultPatchBaselineOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDefaultPatchBaselineResponse is the response type for the
// GetDefaultPatchBaseline API operation.
type GetDefaultPatchBaselineResponse struct {
	*GetDefaultPatchBaselineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDefaultPatchBaseline request.
func (r *GetDefaultPatchBaselineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
