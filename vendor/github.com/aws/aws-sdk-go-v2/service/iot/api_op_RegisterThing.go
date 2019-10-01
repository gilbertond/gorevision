// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RegisterThingInput struct {
	_ struct{} `type:"structure"`

	// The parameters for provisioning a thing. See Programmatic Provisioning (https://docs.aws.amazon.com/iot/latest/developerguide/programmatic-provisioning.html)
	// for more information.
	Parameters map[string]string `locationName:"parameters" type:"map"`

	// The provisioning template. See Programmatic Provisioning (https://docs.aws.amazon.com/iot/latest/developerguide/programmatic-provisioning.html)
	// for more information.
	//
	// TemplateBody is a required field
	TemplateBody *string `locationName:"templateBody" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterThingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterThingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterThingInput"}

	if s.TemplateBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateBody"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterThingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Parameters != nil {
		v := s.Parameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "parameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TemplateBody != nil {
		v := *s.TemplateBody

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateBody", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RegisterThingOutput struct {
	_ struct{} `type:"structure"`

	// .
	CertificatePem *string `locationName:"certificatePem" min:"1" type:"string"`

	// ARNs for the generated resources.
	ResourceArns map[string]string `locationName:"resourceArns" type:"map"`
}

// String returns the string representation
func (s RegisterThingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterThingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificatePem != nil {
		v := *s.CertificatePem

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificatePem", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArns != nil {
		v := s.ResourceArns

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "resourceArns", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opRegisterThing = "RegisterThing"

// RegisterThingRequest returns a request value for making API operation for
// AWS IoT.
//
// Provisions a thing.
//
//    // Example sending a request using RegisterThingRequest.
//    req := client.RegisterThingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RegisterThingRequest(input *RegisterThingInput) RegisterThingRequest {
	op := &aws.Operation{
		Name:       opRegisterThing,
		HTTPMethod: "POST",
		HTTPPath:   "/things",
	}

	if input == nil {
		input = &RegisterThingInput{}
	}

	req := c.newRequest(op, input, &RegisterThingOutput{})
	return RegisterThingRequest{Request: req, Input: input, Copy: c.RegisterThingRequest}
}

// RegisterThingRequest is the request type for the
// RegisterThing API operation.
type RegisterThingRequest struct {
	*aws.Request
	Input *RegisterThingInput
	Copy  func(*RegisterThingInput) RegisterThingRequest
}

// Send marshals and sends the RegisterThing API request.
func (r RegisterThingRequest) Send(ctx context.Context) (*RegisterThingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterThingResponse{
		RegisterThingOutput: r.Request.Data.(*RegisterThingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterThingResponse is the response type for the
// RegisterThing API operation.
type RegisterThingResponse struct {
	*RegisterThingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterThing request.
func (r *RegisterThingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
