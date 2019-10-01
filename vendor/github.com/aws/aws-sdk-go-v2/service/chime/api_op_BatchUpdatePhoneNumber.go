// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type BatchUpdatePhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The request containing the phone number IDs and product types to update.
	//
	// UpdatePhoneNumberRequestItems is a required field
	UpdatePhoneNumberRequestItems []UpdatePhoneNumberRequestItem `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchUpdatePhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchUpdatePhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchUpdatePhoneNumberInput"}

	if s.UpdatePhoneNumberRequestItems == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdatePhoneNumberRequestItems"))
	}
	if s.UpdatePhoneNumberRequestItems != nil {
		for i, v := range s.UpdatePhoneNumberRequestItems {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UpdatePhoneNumberRequestItems", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdatePhoneNumberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.UpdatePhoneNumberRequestItems != nil {
		v := s.UpdatePhoneNumberRequestItems

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UpdatePhoneNumberRequestItems", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type BatchUpdatePhoneNumberOutput struct {
	_ struct{} `type:"structure"`

	// If the action fails for one or more of the phone numbers in the request,
	// a list of the phone numbers is returned, along with error codes and error
	// messages.
	PhoneNumberErrors []PhoneNumberError `type:"list"`
}

// String returns the string representation
func (s BatchUpdatePhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdatePhoneNumberOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PhoneNumberErrors != nil {
		v := s.PhoneNumberErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "PhoneNumberErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchUpdatePhoneNumber = "BatchUpdatePhoneNumber"

// BatchUpdatePhoneNumberRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates phone number product types. Choose from Amazon Chime Business Calling
// and Amazon Chime Voice Connector product types. For toll-free numbers, you
// can use only the Amazon Chime Voice Connector product type.
//
//    // Example sending a request using BatchUpdatePhoneNumberRequest.
//    req := client.BatchUpdatePhoneNumberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchUpdatePhoneNumber
func (c *Client) BatchUpdatePhoneNumberRequest(input *BatchUpdatePhoneNumberInput) BatchUpdatePhoneNumberRequest {
	op := &aws.Operation{
		Name:       opBatchUpdatePhoneNumber,
		HTTPMethod: "POST",
		HTTPPath:   "/phone-numbers?operation=batch-update",
	}

	if input == nil {
		input = &BatchUpdatePhoneNumberInput{}
	}

	req := c.newRequest(op, input, &BatchUpdatePhoneNumberOutput{})
	return BatchUpdatePhoneNumberRequest{Request: req, Input: input, Copy: c.BatchUpdatePhoneNumberRequest}
}

// BatchUpdatePhoneNumberRequest is the request type for the
// BatchUpdatePhoneNumber API operation.
type BatchUpdatePhoneNumberRequest struct {
	*aws.Request
	Input *BatchUpdatePhoneNumberInput
	Copy  func(*BatchUpdatePhoneNumberInput) BatchUpdatePhoneNumberRequest
}

// Send marshals and sends the BatchUpdatePhoneNumber API request.
func (r BatchUpdatePhoneNumberRequest) Send(ctx context.Context) (*BatchUpdatePhoneNumberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchUpdatePhoneNumberResponse{
		BatchUpdatePhoneNumberOutput: r.Request.Data.(*BatchUpdatePhoneNumberOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchUpdatePhoneNumberResponse is the response type for the
// BatchUpdatePhoneNumber API operation.
type BatchUpdatePhoneNumberResponse struct {
	*BatchUpdatePhoneNumberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchUpdatePhoneNumber request.
func (r *BatchUpdatePhoneNumberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
