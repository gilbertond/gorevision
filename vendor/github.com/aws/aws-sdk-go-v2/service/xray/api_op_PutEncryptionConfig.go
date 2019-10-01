// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutEncryptionConfigInput struct {
	_ struct{} `type:"structure"`

	// An AWS KMS customer master key (CMK) in one of the following formats:
	//
	//    * Alias - The name of the key. For example, alias/MyKey.
	//
	//    * Key ID - The KMS key ID of the key. For example, ae4aa6d49-a4d8-9df9-a475-4ff6d7898456.
	//
	//    * ARN - The full Amazon Resource Name of the key ID or alias. For example,
	//    arn:aws:kms:us-east-2:123456789012:key/ae4aa6d49-a4d8-9df9-a475-4ff6d7898456.
	//    Use this format to specify a key in a different account.
	//
	// Omit this key if you set Type to NONE.
	KeyId *string `min:"1" type:"string"`

	// The type of encryption. Set to KMS to use your own key for encryption. Set
	// to NONE for default encryption.
	//
	// Type is a required field
	Type EncryptionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s PutEncryptionConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEncryptionConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEncryptionConfigInput"}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEncryptionConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.KeyId != nil {
		v := *s.KeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type PutEncryptionConfigOutput struct {
	_ struct{} `type:"structure"`

	// The new encryption configuration.
	EncryptionConfig *EncryptionConfig `type:"structure"`
}

// String returns the string representation
func (s PutEncryptionConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEncryptionConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EncryptionConfig != nil {
		v := s.EncryptionConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EncryptionConfig", v, metadata)
	}
	return nil
}

const opPutEncryptionConfig = "PutEncryptionConfig"

// PutEncryptionConfigRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Updates the encryption configuration for X-Ray data.
//
//    // Example sending a request using PutEncryptionConfigRequest.
//    req := client.PutEncryptionConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/PutEncryptionConfig
func (c *Client) PutEncryptionConfigRequest(input *PutEncryptionConfigInput) PutEncryptionConfigRequest {
	op := &aws.Operation{
		Name:       opPutEncryptionConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/PutEncryptionConfig",
	}

	if input == nil {
		input = &PutEncryptionConfigInput{}
	}

	req := c.newRequest(op, input, &PutEncryptionConfigOutput{})
	return PutEncryptionConfigRequest{Request: req, Input: input, Copy: c.PutEncryptionConfigRequest}
}

// PutEncryptionConfigRequest is the request type for the
// PutEncryptionConfig API operation.
type PutEncryptionConfigRequest struct {
	*aws.Request
	Input *PutEncryptionConfigInput
	Copy  func(*PutEncryptionConfigInput) PutEncryptionConfigRequest
}

// Send marshals and sends the PutEncryptionConfig API request.
func (r PutEncryptionConfigRequest) Send(ctx context.Context) (*PutEncryptionConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEncryptionConfigResponse{
		PutEncryptionConfigOutput: r.Request.Data.(*PutEncryptionConfigOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEncryptionConfigResponse is the response type for the
// PutEncryptionConfig API operation.
type PutEncryptionConfigResponse struct {
	*PutEncryptionConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEncryptionConfig request.
func (r *PutEncryptionConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
