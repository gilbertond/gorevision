// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteNamespaceInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNamespaceInput) String() string {
	return awsutil.Prettify(s)
}

type DeleteNamespaceOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the namespace to be deleted.
	NamespaceArn *string `locationName:"namespaceArn" type:"string"`

	// The name of the namespace to be deleted.
	NamespaceName *string `locationName:"namespaceName" type:"string"`
}

// String returns the string representation
func (s DeleteNamespaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNamespace = "DeleteNamespace"

// DeleteNamespaceRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Deletes the specified namespace. This action deletes all of the entities
// in the namespace. Delete the systems and flows that use entities in the namespace
// before performing this action.
//
//    // Example sending a request using DeleteNamespaceRequest.
//    req := client.DeleteNamespaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeleteNamespace
func (c *Client) DeleteNamespaceRequest(input *DeleteNamespaceInput) DeleteNamespaceRequest {
	op := &aws.Operation{
		Name:       opDeleteNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNamespaceInput{}
	}

	req := c.newRequest(op, input, &DeleteNamespaceOutput{})
	return DeleteNamespaceRequest{Request: req, Input: input, Copy: c.DeleteNamespaceRequest}
}

// DeleteNamespaceRequest is the request type for the
// DeleteNamespace API operation.
type DeleteNamespaceRequest struct {
	*aws.Request
	Input *DeleteNamespaceInput
	Copy  func(*DeleteNamespaceInput) DeleteNamespaceRequest
}

// Send marshals and sends the DeleteNamespace API request.
func (r DeleteNamespaceRequest) Send(ctx context.Context) (*DeleteNamespaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNamespaceResponse{
		DeleteNamespaceOutput: r.Request.Data.(*DeleteNamespaceOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNamespaceResponse is the response type for the
// DeleteNamespace API operation.
type DeleteNamespaceResponse struct {
	*DeleteNamespaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNamespace request.
func (r *DeleteNamespaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
