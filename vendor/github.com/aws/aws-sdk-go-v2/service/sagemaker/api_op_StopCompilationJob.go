// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopCompilationJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the model compilation job to stop.
	//
	// CompilationJobName is a required field
	CompilationJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopCompilationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopCompilationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopCompilationJobInput"}

	if s.CompilationJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CompilationJobName"))
	}
	if s.CompilationJobName != nil && len(*s.CompilationJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CompilationJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopCompilationJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopCompilationJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopCompilationJob = "StopCompilationJob"

// StopCompilationJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Stops a model compilation job.
//
// To stop a job, Amazon SageMaker sends the algorithm the SIGTERM signal. This
// gracefully shuts the job down. If the job hasn't stopped, it sends the SIGKILL
// signal.
//
// When it receives a StopCompilationJob request, Amazon SageMaker changes the
// CompilationJobSummary$CompilationJobStatus of the job to Stopping. After
// Amazon SageMaker stops the job, it sets the CompilationJobSummary$CompilationJobStatus
// to Stopped.
//
//    // Example sending a request using StopCompilationJobRequest.
//    req := client.StopCompilationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopCompilationJob
func (c *Client) StopCompilationJobRequest(input *StopCompilationJobInput) StopCompilationJobRequest {
	op := &aws.Operation{
		Name:       opStopCompilationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopCompilationJobInput{}
	}

	req := c.newRequest(op, input, &StopCompilationJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StopCompilationJobRequest{Request: req, Input: input, Copy: c.StopCompilationJobRequest}
}

// StopCompilationJobRequest is the request type for the
// StopCompilationJob API operation.
type StopCompilationJobRequest struct {
	*aws.Request
	Input *StopCompilationJobInput
	Copy  func(*StopCompilationJobInput) StopCompilationJobRequest
}

// Send marshals and sends the StopCompilationJob API request.
func (r StopCompilationJobRequest) Send(ctx context.Context) (*StopCompilationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopCompilationJobResponse{
		StopCompilationJobOutput: r.Request.Data.(*StopCompilationJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopCompilationJobResponse is the response type for the
// StopCompilationJob API operation.
type StopCompilationJobResponse struct {
	*StopCompilationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopCompilationJob request.
func (r *StopCompilationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
