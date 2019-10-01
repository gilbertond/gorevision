// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetServiceQuotaInput struct {
	_ struct{} `type:"structure"`

	// Identifies the service quota you want to select.
	//
	// QuotaCode is a required field
	QuotaCode *string `min:"1" type:"string" required:"true"`

	// Specifies the service that you want to use.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetServiceQuotaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServiceQuotaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServiceQuotaInput"}

	if s.QuotaCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("QuotaCode"))
	}
	if s.QuotaCode != nil && len(*s.QuotaCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QuotaCode", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetServiceQuotaOutput struct {
	_ struct{} `type:"structure"`

	// Returns the ServiceQuota object which contains all values for a quota.
	Quota *ServiceQuota `type:"structure"`
}

// String returns the string representation
func (s GetServiceQuotaOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServiceQuota = "GetServiceQuota"

// GetServiceQuotaRequest returns a request value for making API operation for
// Service Quotas.
//
// Returns the details for the specified service quota. This operation provides
// a different Value than the GetAWSDefaultServiceQuota operation. This operation
// returns the applied value for each quota. GetAWSDefaultServiceQuota returns
// the default AWS value for each quota.
//
//    // Example sending a request using GetServiceQuotaRequest.
//    req := client.GetServiceQuotaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/GetServiceQuota
func (c *Client) GetServiceQuotaRequest(input *GetServiceQuotaInput) GetServiceQuotaRequest {
	op := &aws.Operation{
		Name:       opGetServiceQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServiceQuotaInput{}
	}

	req := c.newRequest(op, input, &GetServiceQuotaOutput{})
	return GetServiceQuotaRequest{Request: req, Input: input, Copy: c.GetServiceQuotaRequest}
}

// GetServiceQuotaRequest is the request type for the
// GetServiceQuota API operation.
type GetServiceQuotaRequest struct {
	*aws.Request
	Input *GetServiceQuotaInput
	Copy  func(*GetServiceQuotaInput) GetServiceQuotaRequest
}

// Send marshals and sends the GetServiceQuota API request.
func (r GetServiceQuotaRequest) Send(ctx context.Context) (*GetServiceQuotaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceQuotaResponse{
		GetServiceQuotaOutput: r.Request.Data.(*GetServiceQuotaOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceQuotaResponse is the response type for the
// GetServiceQuota API operation.
type GetServiceQuotaResponse struct {
	*GetServiceQuotaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServiceQuota request.
func (r *GetServiceQuotaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
