// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateDelegateToResourceInput struct {
	_ struct{} `type:"structure"`

	// The member (user or group) to associate to the resource.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The organization under which the resource exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The resource for which members (users or groups) are associated.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDelegateToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDelegateToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDelegateToResourceInput"}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateDelegateToResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDelegateToResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateDelegateToResource = "AssociateDelegateToResource"

// AssociateDelegateToResourceRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Adds a member (user or group) to the resource's set of delegates.
//
//    // Example sending a request using AssociateDelegateToResourceRequest.
//    req := client.AssociateDelegateToResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/AssociateDelegateToResource
func (c *Client) AssociateDelegateToResourceRequest(input *AssociateDelegateToResourceInput) AssociateDelegateToResourceRequest {
	op := &aws.Operation{
		Name:       opAssociateDelegateToResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateDelegateToResourceInput{}
	}

	req := c.newRequest(op, input, &AssociateDelegateToResourceOutput{})
	return AssociateDelegateToResourceRequest{Request: req, Input: input, Copy: c.AssociateDelegateToResourceRequest}
}

// AssociateDelegateToResourceRequest is the request type for the
// AssociateDelegateToResource API operation.
type AssociateDelegateToResourceRequest struct {
	*aws.Request
	Input *AssociateDelegateToResourceInput
	Copy  func(*AssociateDelegateToResourceInput) AssociateDelegateToResourceRequest
}

// Send marshals and sends the AssociateDelegateToResource API request.
func (r AssociateDelegateToResourceRequest) Send(ctx context.Context) (*AssociateDelegateToResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDelegateToResourceResponse{
		AssociateDelegateToResourceOutput: r.Request.Data.(*AssociateDelegateToResourceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDelegateToResourceResponse is the response type for the
// AssociateDelegateToResource API operation.
type AssociateDelegateToResourceResponse struct {
	*AssociateDelegateToResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDelegateToResource request.
func (r *AssociateDelegateToResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
