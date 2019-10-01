// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DescribeCertificate operation.
type DescribeCertificateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the certificate. (The last part of the certificate ARN contains
	// the certificate ID.)
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"certificateId" min:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCertificateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output of the DescribeCertificate operation.
type DescribeCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The description of the certificate.
	CertificateDescription *CertificateDescription `locationName:"certificateDescription" type:"structure"`
}

// String returns the string representation
func (s DescribeCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateDescription != nil {
		v := s.CertificateDescription

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "certificateDescription", v, metadata)
	}
	return nil
}

const opDescribeCertificate = "DescribeCertificate"

// DescribeCertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the specified certificate.
//
//    // Example sending a request using DescribeCertificateRequest.
//    req := client.DescribeCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeCertificateRequest(input *DescribeCertificateInput) DescribeCertificateRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/certificates/{certificateId}",
	}

	if input == nil {
		input = &DescribeCertificateInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificateOutput{})
	return DescribeCertificateRequest{Request: req, Input: input, Copy: c.DescribeCertificateRequest}
}

// DescribeCertificateRequest is the request type for the
// DescribeCertificate API operation.
type DescribeCertificateRequest struct {
	*aws.Request
	Input *DescribeCertificateInput
	Copy  func(*DescribeCertificateInput) DescribeCertificateRequest
}

// Send marshals and sends the DescribeCertificate API request.
func (r DescribeCertificateRequest) Send(ctx context.Context) (*DescribeCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificateResponse{
		DescribeCertificateOutput: r.Request.Data.(*DescribeCertificateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificateResponse is the response type for the
// DescribeCertificate API operation.
type DescribeCertificateResponse struct {
	*DescribeCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificate request.
func (r *DescribeCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
