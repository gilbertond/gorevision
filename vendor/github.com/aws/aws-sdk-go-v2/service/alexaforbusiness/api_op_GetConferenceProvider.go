// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetConferenceProviderInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created conference provider.
	//
	// ConferenceProviderArn is a required field
	ConferenceProviderArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetConferenceProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConferenceProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConferenceProviderInput"}

	if s.ConferenceProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetConferenceProviderOutput struct {
	_ struct{} `type:"structure"`

	// The conference provider.
	ConferenceProvider *ConferenceProvider `type:"structure"`
}

// String returns the string representation
func (s GetConferenceProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConferenceProvider = "GetConferenceProvider"

// GetConferenceProviderRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets details about a specific conference provider.
//
//    // Example sending a request using GetConferenceProviderRequest.
//    req := client.GetConferenceProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetConferenceProvider
func (c *Client) GetConferenceProviderRequest(input *GetConferenceProviderInput) GetConferenceProviderRequest {
	op := &aws.Operation{
		Name:       opGetConferenceProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConferenceProviderInput{}
	}

	req := c.newRequest(op, input, &GetConferenceProviderOutput{})
	return GetConferenceProviderRequest{Request: req, Input: input, Copy: c.GetConferenceProviderRequest}
}

// GetConferenceProviderRequest is the request type for the
// GetConferenceProvider API operation.
type GetConferenceProviderRequest struct {
	*aws.Request
	Input *GetConferenceProviderInput
	Copy  func(*GetConferenceProviderInput) GetConferenceProviderRequest
}

// Send marshals and sends the GetConferenceProvider API request.
func (r GetConferenceProviderRequest) Send(ctx context.Context) (*GetConferenceProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConferenceProviderResponse{
		GetConferenceProviderOutput: r.Request.Data.(*GetConferenceProviderOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConferenceProviderResponse is the response type for the
// GetConferenceProvider API operation.
type GetConferenceProviderResponse struct {
	*GetConferenceProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConferenceProvider request.
func (r *GetConferenceProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
