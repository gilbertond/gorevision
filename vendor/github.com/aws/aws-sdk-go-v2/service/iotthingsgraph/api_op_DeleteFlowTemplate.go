// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFlowTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the workflow to be deleted.
	//
	// The ID should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:workflow:WORKFLOWNAME
	//
	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFlowTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFlowTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFlowTemplateInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFlowTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFlowTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFlowTemplate = "DeleteFlowTemplate"

// DeleteFlowTemplateRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Deletes a workflow. Any new system or deployment that contains this workflow
// will fail to update or deploy. Existing deployments that contain the workflow
// will continue to run (since they use a snapshot of the workflow taken at
// the time of deployment).
//
//    // Example sending a request using DeleteFlowTemplateRequest.
//    req := client.DeleteFlowTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeleteFlowTemplate
func (c *Client) DeleteFlowTemplateRequest(input *DeleteFlowTemplateInput) DeleteFlowTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteFlowTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFlowTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteFlowTemplateOutput{})
	return DeleteFlowTemplateRequest{Request: req, Input: input, Copy: c.DeleteFlowTemplateRequest}
}

// DeleteFlowTemplateRequest is the request type for the
// DeleteFlowTemplate API operation.
type DeleteFlowTemplateRequest struct {
	*aws.Request
	Input *DeleteFlowTemplateInput
	Copy  func(*DeleteFlowTemplateInput) DeleteFlowTemplateRequest
}

// Send marshals and sends the DeleteFlowTemplate API request.
func (r DeleteFlowTemplateRequest) Send(ctx context.Context) (*DeleteFlowTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFlowTemplateResponse{
		DeleteFlowTemplateOutput: r.Request.Data.(*DeleteFlowTemplateOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFlowTemplateResponse is the response type for the
// DeleteFlowTemplate API operation.
type DeleteFlowTemplateResponse struct {
	*DeleteFlowTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFlowTemplate request.
func (r *DeleteFlowTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
