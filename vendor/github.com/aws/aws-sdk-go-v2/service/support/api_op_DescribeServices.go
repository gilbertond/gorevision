// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeServicesInput struct {
	_ struct{} `type:"structure"`

	// The ISO 639-1 code for the language in which AWS provides support. AWS Support
	// currently supports English ("en") and Japanese ("ja"). Language parameters
	// must be passed explicitly for operations that take them.
	Language *string `locationName:"language" type:"string"`

	// A JSON-formatted list of service codes available for AWS services.
	ServiceCodeList []string `locationName:"serviceCodeList" type:"list"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// The list of AWS services returned by the DescribeServices operation.
type DescribeServicesOutput struct {
	_ struct{} `type:"structure"`

	// A JSON-formatted list of AWS services.
	Services []Service `locationName:"services" type:"list"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeServices = "DescribeServices"

// DescribeServicesRequest returns a request value for making API operation for
// AWS Support.
//
// Returns the current list of AWS services and a list of service categories
// that applies to each one. You then use service names and categories in your
// CreateCase requests. Each AWS service has its own set of categories.
//
// The service codes and category codes correspond to the values that are displayed
// in the Service and Category drop-down lists on the AWS Support Center Create
// Case (https://console.aws.amazon.com/support/home#/case/create) page. The
// values in those fields, however, do not necessarily match the service codes
// and categories returned by the DescribeServices request. Always use the service
// codes and categories obtained programmatically. This practice ensures that
// you always have the most recent set of service and category codes.
//
//    // Example sending a request using DescribeServicesRequest.
//    req := client.DescribeServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeServices
func (c *Client) DescribeServicesRequest(input *DescribeServicesInput) DescribeServicesRequest {
	op := &aws.Operation{
		Name:       opDescribeServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServicesInput{}
	}

	req := c.newRequest(op, input, &DescribeServicesOutput{})
	return DescribeServicesRequest{Request: req, Input: input, Copy: c.DescribeServicesRequest}
}

// DescribeServicesRequest is the request type for the
// DescribeServices API operation.
type DescribeServicesRequest struct {
	*aws.Request
	Input *DescribeServicesInput
	Copy  func(*DescribeServicesInput) DescribeServicesRequest
}

// Send marshals and sends the DescribeServices API request.
func (r DescribeServicesRequest) Send(ctx context.Context) (*DescribeServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeServicesResponse{
		DescribeServicesOutput: r.Request.Data.(*DescribeServicesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeServicesResponse is the response type for the
// DescribeServices API operation.
type DescribeServicesResponse struct {
	*DescribeServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeServices request.
func (r *DescribeServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
