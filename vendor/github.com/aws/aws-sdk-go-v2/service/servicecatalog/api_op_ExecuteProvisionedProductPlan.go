// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ExecuteProvisionedProductPlanInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The plan identifier.
	//
	// PlanId is a required field
	PlanId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ExecuteProvisionedProductPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteProvisionedProductPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExecuteProvisionedProductPlanInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.PlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlanId"))
	}
	if s.PlanId != nil && len(*s.PlanId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlanId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ExecuteProvisionedProductPlanOutput struct {
	_ struct{} `type:"structure"`

	// Information about the result of provisioning the product.
	RecordDetail *RecordDetail `type:"structure"`
}

// String returns the string representation
func (s ExecuteProvisionedProductPlanOutput) String() string {
	return awsutil.Prettify(s)
}

const opExecuteProvisionedProductPlan = "ExecuteProvisionedProductPlan"

// ExecuteProvisionedProductPlanRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Provisions or modifies a product based on the resource changes for the specified
// plan.
//
//    // Example sending a request using ExecuteProvisionedProductPlanRequest.
//    req := client.ExecuteProvisionedProductPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ExecuteProvisionedProductPlan
func (c *Client) ExecuteProvisionedProductPlanRequest(input *ExecuteProvisionedProductPlanInput) ExecuteProvisionedProductPlanRequest {
	op := &aws.Operation{
		Name:       opExecuteProvisionedProductPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExecuteProvisionedProductPlanInput{}
	}

	req := c.newRequest(op, input, &ExecuteProvisionedProductPlanOutput{})
	return ExecuteProvisionedProductPlanRequest{Request: req, Input: input, Copy: c.ExecuteProvisionedProductPlanRequest}
}

// ExecuteProvisionedProductPlanRequest is the request type for the
// ExecuteProvisionedProductPlan API operation.
type ExecuteProvisionedProductPlanRequest struct {
	*aws.Request
	Input *ExecuteProvisionedProductPlanInput
	Copy  func(*ExecuteProvisionedProductPlanInput) ExecuteProvisionedProductPlanRequest
}

// Send marshals and sends the ExecuteProvisionedProductPlan API request.
func (r ExecuteProvisionedProductPlanRequest) Send(ctx context.Context) (*ExecuteProvisionedProductPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExecuteProvisionedProductPlanResponse{
		ExecuteProvisionedProductPlanOutput: r.Request.Data.(*ExecuteProvisionedProductPlanOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExecuteProvisionedProductPlanResponse is the response type for the
// ExecuteProvisionedProductPlan API operation.
type ExecuteProvisionedProductPlanResponse struct {
	*ExecuteProvisionedProductPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExecuteProvisionedProductPlan request.
func (r *ExecuteProvisionedProductPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
