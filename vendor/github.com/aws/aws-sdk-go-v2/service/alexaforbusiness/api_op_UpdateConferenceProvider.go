// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateConferenceProviderInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the conference provider.
	//
	// ConferenceProviderArn is a required field
	ConferenceProviderArn *string `type:"string" required:"true"`

	// The type of the conference provider.
	//
	// ConferenceProviderType is a required field
	ConferenceProviderType ConferenceProviderType `type:"string" required:"true" enum:"true"`

	// The IP endpoint and protocol for calling.
	IPDialIn *IPDialIn `type:"structure"`

	// The meeting settings for the conference provider.
	//
	// MeetingSetting is a required field
	MeetingSetting *MeetingSetting `type:"structure" required:"true"`

	// The information for PSTN conferencing.
	PSTNDialIn *PSTNDialIn `type:"structure"`
}

// String returns the string representation
func (s UpdateConferenceProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConferenceProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConferenceProviderInput"}

	if s.ConferenceProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderArn"))
	}
	if len(s.ConferenceProviderType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderType"))
	}

	if s.MeetingSetting == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingSetting"))
	}
	if s.IPDialIn != nil {
		if err := s.IPDialIn.Validate(); err != nil {
			invalidParams.AddNested("IPDialIn", err.(aws.ErrInvalidParams))
		}
	}
	if s.MeetingSetting != nil {
		if err := s.MeetingSetting.Validate(); err != nil {
			invalidParams.AddNested("MeetingSetting", err.(aws.ErrInvalidParams))
		}
	}
	if s.PSTNDialIn != nil {
		if err := s.PSTNDialIn.Validate(); err != nil {
			invalidParams.AddNested("PSTNDialIn", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateConferenceProviderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConferenceProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateConferenceProvider = "UpdateConferenceProvider"

// UpdateConferenceProviderRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates an existing conference provider's settings.
//
//    // Example sending a request using UpdateConferenceProviderRequest.
//    req := client.UpdateConferenceProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateConferenceProvider
func (c *Client) UpdateConferenceProviderRequest(input *UpdateConferenceProviderInput) UpdateConferenceProviderRequest {
	op := &aws.Operation{
		Name:       opUpdateConferenceProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateConferenceProviderInput{}
	}

	req := c.newRequest(op, input, &UpdateConferenceProviderOutput{})
	return UpdateConferenceProviderRequest{Request: req, Input: input, Copy: c.UpdateConferenceProviderRequest}
}

// UpdateConferenceProviderRequest is the request type for the
// UpdateConferenceProvider API operation.
type UpdateConferenceProviderRequest struct {
	*aws.Request
	Input *UpdateConferenceProviderInput
	Copy  func(*UpdateConferenceProviderInput) UpdateConferenceProviderRequest
}

// Send marshals and sends the UpdateConferenceProvider API request.
func (r UpdateConferenceProviderRequest) Send(ctx context.Context) (*UpdateConferenceProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConferenceProviderResponse{
		UpdateConferenceProviderOutput: r.Request.Data.(*UpdateConferenceProviderOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConferenceProviderResponse is the response type for the
// UpdateConferenceProvider API operation.
type UpdateConferenceProviderResponse struct {
	*UpdateConferenceProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConferenceProvider request.
func (r *UpdateConferenceProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
