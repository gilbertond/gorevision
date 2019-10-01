// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UntagInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource from which to remove tags.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"Arn" min:"12" type:"string" required:"true"`

	// The keys of the tags to be removed.
	//
	// Keys is a required field
	Keys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s UntagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 12))
	}

	if s.Keys == nil {
		invalidParams.Add(aws.NewErrParamRequired("Keys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Keys != nil {
		v := s.Keys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Keys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UntagOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource from which tags have been removed.
	Arn *string `min:"12" type:"string"`

	// The keys of tags that have been removed.
	Keys []string `type:"list"`
}

// String returns the string representation
func (s UntagOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Keys != nil {
		v := s.Keys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Keys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opUntag = "Untag"

// UntagRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Deletes specified tags from a specified resource.
//
//    // Example sending a request using UntagRequest.
//    req := client.UntagRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/Untag
func (c *Client) UntagRequest(input *UntagInput) UntagRequest {
	op := &aws.Operation{
		Name:       opUntag,
		HTTPMethod: "PATCH",
		HTTPPath:   "/resources/{Arn}/tags",
	}

	if input == nil {
		input = &UntagInput{}
	}

	req := c.newRequest(op, input, &UntagOutput{})
	return UntagRequest{Request: req, Input: input, Copy: c.UntagRequest}
}

// UntagRequest is the request type for the
// Untag API operation.
type UntagRequest struct {
	*aws.Request
	Input *UntagInput
	Copy  func(*UntagInput) UntagRequest
}

// Send marshals and sends the Untag API request.
func (r UntagRequest) Send(ctx context.Context) (*UntagResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResponse{
		UntagOutput: r.Request.Data.(*UntagOutput),
		response:    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResponse is the response type for the
// Untag API operation.
type UntagResponse struct {
	*UntagOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Untag request.
func (r *UntagResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
