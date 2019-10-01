// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	MaxResults *string `location:"querystring" locationName:"maxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentsInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	Items []Deployment `locationName:"items" type:"list"`

	// The next page of elements from this collection. Not valid for the last element
	// of the collection.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "items", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDeployments = "GetDeployments"

// GetDeploymentsRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the Deployments for an API.
//
//    // Example sending a request using GetDeploymentsRequest.
//    req := client.GetDeploymentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDeployments
func (c *Client) GetDeploymentsRequest(input *GetDeploymentsInput) GetDeploymentsRequest {
	op := &aws.Operation{
		Name:       opGetDeployments,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/deployments",
	}

	if input == nil {
		input = &GetDeploymentsInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentsOutput{})
	return GetDeploymentsRequest{Request: req, Input: input, Copy: c.GetDeploymentsRequest}
}

// GetDeploymentsRequest is the request type for the
// GetDeployments API operation.
type GetDeploymentsRequest struct {
	*aws.Request
	Input *GetDeploymentsInput
	Copy  func(*GetDeploymentsInput) GetDeploymentsRequest
}

// Send marshals and sends the GetDeployments API request.
func (r GetDeploymentsRequest) Send(ctx context.Context) (*GetDeploymentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentsResponse{
		GetDeploymentsOutput: r.Request.Data.(*GetDeploymentsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentsResponse is the response type for the
// GetDeployments API operation.
type GetDeploymentsResponse struct {
	*GetDeploymentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeployments request.
func (r *GetDeploymentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
