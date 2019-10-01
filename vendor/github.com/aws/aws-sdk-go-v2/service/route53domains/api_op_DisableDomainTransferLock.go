// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The DisableDomainTransferLock request includes the following element.
type DisableDomainTransferLockInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to remove the transfer lock for.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableDomainTransferLockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableDomainTransferLockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableDomainTransferLockInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The DisableDomainTransferLock response includes the following element.
type DisableDomainTransferLockOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for tracking the progress of the request. To use this ID to query
	// the operation status, use GetOperationDetail.
	//
	// OperationId is a required field
	OperationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableDomainTransferLockOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableDomainTransferLock = "DisableDomainTransferLock"

// DisableDomainTransferLockRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation removes the transfer lock on the domain (specifically the
// clientTransferProhibited status) to allow domain transfers. We recommend
// you refrain from performing this action unless you intend to transfer the
// domain to a different registrar. Successful submission returns an operation
// ID that you can use to track the progress and completion of the action. If
// the request is not completed successfully, the domain registrant will be
// notified by email.
//
//    // Example sending a request using DisableDomainTransferLockRequest.
//    req := client.DisableDomainTransferLockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/DisableDomainTransferLock
func (c *Client) DisableDomainTransferLockRequest(input *DisableDomainTransferLockInput) DisableDomainTransferLockRequest {
	op := &aws.Operation{
		Name:       opDisableDomainTransferLock,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableDomainTransferLockInput{}
	}

	req := c.newRequest(op, input, &DisableDomainTransferLockOutput{})
	return DisableDomainTransferLockRequest{Request: req, Input: input, Copy: c.DisableDomainTransferLockRequest}
}

// DisableDomainTransferLockRequest is the request type for the
// DisableDomainTransferLock API operation.
type DisableDomainTransferLockRequest struct {
	*aws.Request
	Input *DisableDomainTransferLockInput
	Copy  func(*DisableDomainTransferLockInput) DisableDomainTransferLockRequest
}

// Send marshals and sends the DisableDomainTransferLock API request.
func (r DisableDomainTransferLockRequest) Send(ctx context.Context) (*DisableDomainTransferLockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableDomainTransferLockResponse{
		DisableDomainTransferLockOutput: r.Request.Data.(*DisableDomainTransferLockOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableDomainTransferLockResponse is the response type for the
// DisableDomainTransferLock API operation.
type DisableDomainTransferLockResponse struct {
	*DisableDomainTransferLockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableDomainTransferLock request.
func (r *DisableDomainTransferLockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
