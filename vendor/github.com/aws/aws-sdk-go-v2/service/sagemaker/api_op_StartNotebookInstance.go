// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StartNotebookInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the notebook instance to start.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartNotebookInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartNotebookInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartNotebookInstanceInput"}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartNotebookInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartNotebookInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartNotebookInstance = "StartNotebookInstance"

// StartNotebookInstanceRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Launches an ML compute instance with the latest version of the libraries
// and attaches your ML storage volume. After configuring the notebook instance,
// Amazon SageMaker sets the notebook instance status to InService. A notebook
// instance's status must be InService before you can connect to your Jupyter
// notebook.
//
//    // Example sending a request using StartNotebookInstanceRequest.
//    req := client.StartNotebookInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StartNotebookInstance
func (c *Client) StartNotebookInstanceRequest(input *StartNotebookInstanceInput) StartNotebookInstanceRequest {
	op := &aws.Operation{
		Name:       opStartNotebookInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartNotebookInstanceInput{}
	}

	req := c.newRequest(op, input, &StartNotebookInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartNotebookInstanceRequest{Request: req, Input: input, Copy: c.StartNotebookInstanceRequest}
}

// StartNotebookInstanceRequest is the request type for the
// StartNotebookInstance API operation.
type StartNotebookInstanceRequest struct {
	*aws.Request
	Input *StartNotebookInstanceInput
	Copy  func(*StartNotebookInstanceInput) StartNotebookInstanceRequest
}

// Send marshals and sends the StartNotebookInstance API request.
func (r StartNotebookInstanceRequest) Send(ctx context.Context) (*StartNotebookInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartNotebookInstanceResponse{
		StartNotebookInstanceOutput: r.Request.Data.(*StartNotebookInstanceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartNotebookInstanceResponse is the response type for the
// StartNotebookInstance API operation.
type StartNotebookInstanceResponse struct {
	*StartNotebookInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartNotebookInstance request.
func (r *StartNotebookInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
