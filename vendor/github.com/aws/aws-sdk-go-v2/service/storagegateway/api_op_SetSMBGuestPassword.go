// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// SetSMBGuestPasswordInput
type SetSMBGuestPasswordInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file gateway the SMB file share is
	// associated with.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The password that you want to set for your SMB Server.
	//
	// Password is a required field
	Password *string `min:"6" type:"string" required:"true"`
}

// String returns the string representation
func (s SetSMBGuestPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetSMBGuestPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetSMBGuestPasswordInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetSMBGuestPasswordOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s SetSMBGuestPasswordOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetSMBGuestPassword = "SetSMBGuestPassword"

// SetSMBGuestPasswordRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Sets the password for the guest user smbguest. The smbguest user is the user
// when the authentication method for the file share is set to GuestAccess.
//
//    // Example sending a request using SetSMBGuestPasswordRequest.
//    req := client.SetSMBGuestPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/SetSMBGuestPassword
func (c *Client) SetSMBGuestPasswordRequest(input *SetSMBGuestPasswordInput) SetSMBGuestPasswordRequest {
	op := &aws.Operation{
		Name:       opSetSMBGuestPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetSMBGuestPasswordInput{}
	}

	req := c.newRequest(op, input, &SetSMBGuestPasswordOutput{})
	return SetSMBGuestPasswordRequest{Request: req, Input: input, Copy: c.SetSMBGuestPasswordRequest}
}

// SetSMBGuestPasswordRequest is the request type for the
// SetSMBGuestPassword API operation.
type SetSMBGuestPasswordRequest struct {
	*aws.Request
	Input *SetSMBGuestPasswordInput
	Copy  func(*SetSMBGuestPasswordInput) SetSMBGuestPasswordRequest
}

// Send marshals and sends the SetSMBGuestPassword API request.
func (r SetSMBGuestPasswordRequest) Send(ctx context.Context) (*SetSMBGuestPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetSMBGuestPasswordResponse{
		SetSMBGuestPasswordOutput: r.Request.Data.(*SetSMBGuestPasswordOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetSMBGuestPasswordResponse is the response type for the
// SetSMBGuestPassword API operation.
type SetSMBGuestPasswordResponse struct {
	*SetSMBGuestPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetSMBGuestPassword request.
func (r *SetSMBGuestPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
