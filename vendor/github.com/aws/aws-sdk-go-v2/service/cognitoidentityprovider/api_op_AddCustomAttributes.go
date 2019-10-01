// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to add custom attributes.
type AddCustomAttributesInput struct {
	_ struct{} `type:"structure"`

	// An array of custom attributes, such as Mutable and Name.
	//
	// CustomAttributes is a required field
	CustomAttributes []SchemaAttributeType `min:"1" type:"list" required:"true"`

	// The user pool ID for the user pool where you want to add custom attributes.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AddCustomAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddCustomAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddCustomAttributesInput"}

	if s.CustomAttributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomAttributes"))
	}
	if s.CustomAttributes != nil && len(s.CustomAttributes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomAttributes", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.CustomAttributes != nil {
		for i, v := range s.CustomAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CustomAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server for the request to add custom attributes.
type AddCustomAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddCustomAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddCustomAttributes = "AddCustomAttributes"

// AddCustomAttributesRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Adds additional user attributes to the user pool schema.
//
//    // Example sending a request using AddCustomAttributesRequest.
//    req := client.AddCustomAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AddCustomAttributes
func (c *Client) AddCustomAttributesRequest(input *AddCustomAttributesInput) AddCustomAttributesRequest {
	op := &aws.Operation{
		Name:       opAddCustomAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddCustomAttributesInput{}
	}

	req := c.newRequest(op, input, &AddCustomAttributesOutput{})
	return AddCustomAttributesRequest{Request: req, Input: input, Copy: c.AddCustomAttributesRequest}
}

// AddCustomAttributesRequest is the request type for the
// AddCustomAttributes API operation.
type AddCustomAttributesRequest struct {
	*aws.Request
	Input *AddCustomAttributesInput
	Copy  func(*AddCustomAttributesInput) AddCustomAttributesRequest
}

// Send marshals and sends the AddCustomAttributes API request.
func (r AddCustomAttributesRequest) Send(ctx context.Context) (*AddCustomAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddCustomAttributesResponse{
		AddCustomAttributesOutput: r.Request.Data.(*AddCustomAttributesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddCustomAttributesResponse is the response type for the
// AddCustomAttributes API operation.
type AddCustomAttributesResponse struct {
	*AddCustomAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddCustomAttributes request.
func (r *AddCustomAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
