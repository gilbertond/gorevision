// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisableDomainAutoRenewInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to disable automatic renewal for.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableDomainAutoRenewInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableDomainAutoRenewInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableDomainAutoRenewInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisableDomainAutoRenewOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableDomainAutoRenewOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableDomainAutoRenew = "DisableDomainAutoRenew"

// DisableDomainAutoRenewRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation disables automatic renewal of domain registration for the
// specified domain.
//
//    // Example sending a request using DisableDomainAutoRenewRequest.
//    req := client.DisableDomainAutoRenewRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DisableDomainAutoRenew
func (c *Client) DisableDomainAutoRenewRequest(input *DisableDomainAutoRenewInput) DisableDomainAutoRenewRequest {
	op := &aws.Operation{
		Name:       opDisableDomainAutoRenew,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableDomainAutoRenewInput{}
	}

	req := c.newRequest(op, input, &DisableDomainAutoRenewOutput{})
	return DisableDomainAutoRenewRequest{Request: req, Input: input, Copy: c.DisableDomainAutoRenewRequest}
}

// DisableDomainAutoRenewRequest is the request type for the
// DisableDomainAutoRenew API operation.
type DisableDomainAutoRenewRequest struct {
	*aws.Request
	Input *DisableDomainAutoRenewInput
	Copy  func(*DisableDomainAutoRenewInput) DisableDomainAutoRenewRequest
}

// Send marshals and sends the DisableDomainAutoRenew API request.
func (r DisableDomainAutoRenewRequest) Send(ctx context.Context) (*DisableDomainAutoRenewResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableDomainAutoRenewResponse{
		DisableDomainAutoRenewOutput: r.Request.Data.(*DisableDomainAutoRenewOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableDomainAutoRenewResponse is the response type for the
// DisableDomainAutoRenew API operation.
type DisableDomainAutoRenewResponse struct {
	*DisableDomainAutoRenewOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableDomainAutoRenew request.
func (r *DisableDomainAutoRenewResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
