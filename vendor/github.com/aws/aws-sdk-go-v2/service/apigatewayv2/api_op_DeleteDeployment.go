// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteDeploymentInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// DeploymentId is a required field
	DeploymentId *string `location:"uri" locationName:"deploymentId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDeploymentInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDeploymentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDeployment = "DeleteDeployment"

// DeleteDeploymentRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes a Deployment.
//
//    // Example sending a request using DeleteDeploymentRequest.
//    req := client.DeleteDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteDeployment
func (c *Client) DeleteDeploymentRequest(input *DeleteDeploymentInput) DeleteDeploymentRequest {
	op := &aws.Operation{
		Name:       opDeleteDeployment,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/deployments/{deploymentId}",
	}

	if input == nil {
		input = &DeleteDeploymentInput{}
	}

	req := c.newRequest(op, input, &DeleteDeploymentOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDeploymentRequest{Request: req, Input: input, Copy: c.DeleteDeploymentRequest}
}

// DeleteDeploymentRequest is the request type for the
// DeleteDeployment API operation.
type DeleteDeploymentRequest struct {
	*aws.Request
	Input *DeleteDeploymentInput
	Copy  func(*DeleteDeploymentInput) DeleteDeploymentRequest
}

// Send marshals and sends the DeleteDeployment API request.
func (r DeleteDeploymentRequest) Send(ctx context.Context) (*DeleteDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDeploymentResponse{
		DeleteDeploymentOutput: r.Request.Data.(*DeleteDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDeploymentResponse is the response type for the
// DeleteDeployment API operation.
type DeleteDeploymentResponse struct {
	*DeleteDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDeployment request.
func (r *DeleteDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
