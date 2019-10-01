// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeExecutionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the execution to describe.
	//
	// ExecutionArn is a required field
	ExecutionArn *string `locationName:"executionArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeExecutionInput"}

	if s.ExecutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExecutionArn"))
	}
	if s.ExecutionArn != nil && len(*s.ExecutionArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExecutionArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// ExecutionArn is a required field
	ExecutionArn *string `locationName:"executionArn" min:"1" type:"string" required:"true"`

	// The string that contains the JSON input data of the execution.
	//
	// Input is a required field
	Input *string `locationName:"input" type:"string" required:"true"`

	// The name of the execution.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	Name *string `locationName:"name" min:"1" type:"string"`

	// The JSON output data of the execution.
	//
	// This field is set only if the execution succeeds. If the execution fails,
	// this field is null.
	Output *string `locationName:"output" type:"string"`

	// The date the execution is started.
	//
	// StartDate is a required field
	StartDate *time.Time `locationName:"startDate" type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) of the executed stated machine.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`

	// The current status of the execution.
	//
	// Status is a required field
	Status ExecutionStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// If the execution has already ended, the date the execution stopped.
	StopDate *time.Time `locationName:"stopDate" type:"timestamp"`
}

// String returns the string representation
func (s DescribeExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeExecution = "DescribeExecution"

// DescribeExecutionRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Describes an execution.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
//    // Example sending a request using DescribeExecutionRequest.
//    req := client.DescribeExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DescribeExecution
func (c *Client) DescribeExecutionRequest(input *DescribeExecutionInput) DescribeExecutionRequest {
	op := &aws.Operation{
		Name:       opDescribeExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeExecutionInput{}
	}

	req := c.newRequest(op, input, &DescribeExecutionOutput{})
	return DescribeExecutionRequest{Request: req, Input: input, Copy: c.DescribeExecutionRequest}
}

// DescribeExecutionRequest is the request type for the
// DescribeExecution API operation.
type DescribeExecutionRequest struct {
	*aws.Request
	Input *DescribeExecutionInput
	Copy  func(*DescribeExecutionInput) DescribeExecutionRequest
}

// Send marshals and sends the DescribeExecution API request.
func (r DescribeExecutionRequest) Send(ctx context.Context) (*DescribeExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeExecutionResponse{
		DescribeExecutionOutput: r.Request.Data.(*DescribeExecutionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeExecutionResponse is the response type for the
// DescribeExecution API operation.
type DescribeExecutionResponse struct {
	*DescribeExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeExecution request.
func (r *DescribeExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
