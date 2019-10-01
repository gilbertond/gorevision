// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteResolverInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The resolver field name.
	//
	// FieldName is a required field
	FieldName *string `location:"uri" locationName:"fieldName" type:"string" required:"true"`

	// The name of the resolver type.
	//
	// TypeName is a required field
	TypeName *string `location:"uri" locationName:"typeName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResolverInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResolverInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResolverInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.FieldName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldName"))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResolverInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FieldName != nil {
		v := *s.FieldName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "fieldName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TypeName != nil {
		v := *s.TypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "typeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteResolverOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResolverOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResolverOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteResolver = "DeleteResolver"

// DeleteResolverRequest returns a request value for making API operation for
// AWS AppSync.
//
// Deletes a Resolver object.
//
//    // Example sending a request using DeleteResolverRequest.
//    req := client.DeleteResolverRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/DeleteResolver
func (c *Client) DeleteResolverRequest(input *DeleteResolverInput) DeleteResolverRequest {
	op := &aws.Operation{
		Name:       opDeleteResolver,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apis/{apiId}/types/{typeName}/resolvers/{fieldName}",
	}

	if input == nil {
		input = &DeleteResolverInput{}
	}

	req := c.newRequest(op, input, &DeleteResolverOutput{})
	return DeleteResolverRequest{Request: req, Input: input, Copy: c.DeleteResolverRequest}
}

// DeleteResolverRequest is the request type for the
// DeleteResolver API operation.
type DeleteResolverRequest struct {
	*aws.Request
	Input *DeleteResolverInput
	Copy  func(*DeleteResolverInput) DeleteResolverRequest
}

// Send marshals and sends the DeleteResolver API request.
func (r DeleteResolverRequest) Send(ctx context.Context) (*DeleteResolverResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResolverResponse{
		DeleteResolverOutput: r.Request.Data.(*DeleteResolverOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResolverResponse is the response type for the
// DeleteResolver API operation.
type DeleteResolverResponse struct {
	*DeleteResolverOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResolver request.
func (r *DeleteResolverResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
