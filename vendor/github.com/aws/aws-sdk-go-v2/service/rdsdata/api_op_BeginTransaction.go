// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request parameters represent the input of a request to start a SQL transaction.
type BeginTransactionInput struct {
	_ struct{} `type:"structure"`

	Database *string `locationName:"database" type:"string"`

	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" type:"string" required:"true"`

	Schema *string `locationName:"schema" type:"string"`

	// SecretArn is a required field
	SecretArn *string `locationName:"secretArn" type:"string" required:"true"`
}

// String returns the string representation
func (s BeginTransactionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BeginTransactionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BeginTransactionInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if s.SecretArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BeginTransactionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Database != nil {
		v := *s.Database

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "database", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecretArn != nil {
		v := *s.SecretArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "secretArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The response elements represent the output of a request to start a SQL transaction.
type BeginTransactionOutput struct {
	_ struct{} `type:"structure"`

	TransactionId *string `locationName:"transactionId" type:"string"`
}

// String returns the string representation
func (s BeginTransactionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BeginTransactionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TransactionId != nil {
		v := *s.TransactionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "transactionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opBeginTransaction = "BeginTransaction"

// BeginTransactionRequest returns a request value for making API operation for
// AWS RDS DataService.
//
// Starts a SQL transaction.
//
// A transaction can run for a maximum of 24 hours. A transaction is terminated
// and rolled back automatically after 24 hours.
//
// A transaction times out if no calls use its transaction ID in three minutes.
// If a transaction times out before it's committed, it's rolled back automatically.
//
// DDL statements inside a transaction cause an implicit commit. We recommend
// that you run each DDL statement in a separate ExecuteStatement call with
// continueAfterTimeout enabled.
//
//    // Example sending a request using BeginTransactionRequest.
//    req := client.BeginTransactionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/BeginTransaction
func (c *Client) BeginTransactionRequest(input *BeginTransactionInput) BeginTransactionRequest {
	op := &aws.Operation{
		Name:       opBeginTransaction,
		HTTPMethod: "POST",
		HTTPPath:   "/BeginTransaction",
	}

	if input == nil {
		input = &BeginTransactionInput{}
	}

	req := c.newRequest(op, input, &BeginTransactionOutput{})
	return BeginTransactionRequest{Request: req, Input: input, Copy: c.BeginTransactionRequest}
}

// BeginTransactionRequest is the request type for the
// BeginTransaction API operation.
type BeginTransactionRequest struct {
	*aws.Request
	Input *BeginTransactionInput
	Copy  func(*BeginTransactionInput) BeginTransactionRequest
}

// Send marshals and sends the BeginTransaction API request.
func (r BeginTransactionRequest) Send(ctx context.Context) (*BeginTransactionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BeginTransactionResponse{
		BeginTransactionOutput: r.Request.Data.(*BeginTransactionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BeginTransactionResponse is the response type for the
// BeginTransaction API operation.
type BeginTransactionResponse struct {
	*BeginTransactionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BeginTransaction request.
func (r *BeginTransactionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
