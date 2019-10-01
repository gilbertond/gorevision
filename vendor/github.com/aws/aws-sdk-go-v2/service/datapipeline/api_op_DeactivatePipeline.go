// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DeactivatePipeline.
type DeactivatePipelineInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether to cancel any running objects. The default is true, which
	// sets the state of any running objects to CANCELED. If this value is false,
	// the pipeline is deactivated after all running objects finish.
	CancelActive *bool `locationName:"cancelActive" type:"boolean"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeactivatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeactivatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeactivatePipelineInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DeactivatePipeline.
type DeactivatePipelineOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeactivatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeactivatePipeline = "DeactivatePipeline"

// DeactivatePipelineRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Deactivates the specified running pipeline. The pipeline is set to the DEACTIVATING
// state until the deactivation process completes.
//
// To resume a deactivated pipeline, use ActivatePipeline. By default, the pipeline
// resumes from the last completed execution. Optionally, you can specify the
// date and time to resume the pipeline.
//
//    // Example sending a request using DeactivatePipelineRequest.
//    req := client.DeactivatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/DeactivatePipeline
func (c *Client) DeactivatePipelineRequest(input *DeactivatePipelineInput) DeactivatePipelineRequest {
	op := &aws.Operation{
		Name:       opDeactivatePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeactivatePipelineInput{}
	}

	req := c.newRequest(op, input, &DeactivatePipelineOutput{})
	return DeactivatePipelineRequest{Request: req, Input: input, Copy: c.DeactivatePipelineRequest}
}

// DeactivatePipelineRequest is the request type for the
// DeactivatePipeline API operation.
type DeactivatePipelineRequest struct {
	*aws.Request
	Input *DeactivatePipelineInput
	Copy  func(*DeactivatePipelineInput) DeactivatePipelineRequest
}

// Send marshals and sends the DeactivatePipeline API request.
func (r DeactivatePipelineRequest) Send(ctx context.Context) (*DeactivatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeactivatePipelineResponse{
		DeactivatePipelineOutput: r.Request.Data.(*DeactivatePipelineOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeactivatePipelineResponse is the response type for the
// DeactivatePipeline API operation.
type DeactivatePipelineResponse struct {
	*DeactivatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeactivatePipeline request.
func (r *DeactivatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
