// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetModelTemplateInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// ModelId is a required field
	ModelId *string `location:"uri" locationName:"modelId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetModelTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetModelTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetModelTemplateInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.ModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetModelTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ModelId != nil {
		v := *s.ModelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "modelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetModelTemplateOutput struct {
	_ struct{} `type:"structure"`

	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s GetModelTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetModelTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetModelTemplate = "GetModelTemplate"

// GetModelTemplateRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a model template.
//
//    // Example sending a request using GetModelTemplateRequest.
//    req := client.GetModelTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetModelTemplate
func (c *Client) GetModelTemplateRequest(input *GetModelTemplateInput) GetModelTemplateRequest {
	op := &aws.Operation{
		Name:       opGetModelTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/models/{modelId}/template",
	}

	if input == nil {
		input = &GetModelTemplateInput{}
	}

	req := c.newRequest(op, input, &GetModelTemplateOutput{})
	return GetModelTemplateRequest{Request: req, Input: input, Copy: c.GetModelTemplateRequest}
}

// GetModelTemplateRequest is the request type for the
// GetModelTemplate API operation.
type GetModelTemplateRequest struct {
	*aws.Request
	Input *GetModelTemplateInput
	Copy  func(*GetModelTemplateInput) GetModelTemplateRequest
}

// Send marshals and sends the GetModelTemplate API request.
func (r GetModelTemplateRequest) Send(ctx context.Context) (*GetModelTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetModelTemplateResponse{
		GetModelTemplateOutput: r.Request.Data.(*GetModelTemplateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetModelTemplateResponse is the response type for the
// GetModelTemplate API operation.
type GetModelTemplateResponse struct {
	*GetModelTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetModelTemplate request.
func (r *GetModelTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
