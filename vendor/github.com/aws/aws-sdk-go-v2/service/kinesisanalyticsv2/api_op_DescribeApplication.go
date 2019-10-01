// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Displays verbose information about a Kinesis Data Analytics application,
	// including the application's job plan.
	IncludeAdditionalDetails *bool `type:"boolean"`
}

// String returns the string representation
func (s DescribeApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeApplicationOutput struct {
	_ struct{} `type:"structure"`

	// Provides a description of the application, such as the application's Amazon
	// Resource Name (ARN), status, and latest version.
	//
	// ApplicationDetail is a required field
	ApplicationDetail *ApplicationDetail `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeApplication = "DescribeApplication"

// DescribeApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Returns information about a specific Amazon Kinesis Data Analytics application.
//
// If you want to retrieve a list of all applications in your account, use the
// ListApplications operation.
//
//    // Example sending a request using DescribeApplicationRequest.
//    req := client.DescribeApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DescribeApplication
func (c *Client) DescribeApplicationRequest(input *DescribeApplicationInput) DescribeApplicationRequest {
	op := &aws.Operation{
		Name:       opDescribeApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeApplicationInput{}
	}

	req := c.newRequest(op, input, &DescribeApplicationOutput{})
	return DescribeApplicationRequest{Request: req, Input: input, Copy: c.DescribeApplicationRequest}
}

// DescribeApplicationRequest is the request type for the
// DescribeApplication API operation.
type DescribeApplicationRequest struct {
	*aws.Request
	Input *DescribeApplicationInput
	Copy  func(*DescribeApplicationInput) DescribeApplicationRequest
}

// Send marshals and sends the DescribeApplication API request.
func (r DescribeApplicationRequest) Send(ctx context.Context) (*DescribeApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeApplicationResponse{
		DescribeApplicationOutput: r.Request.Data.(*DescribeApplicationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeApplicationResponse is the response type for the
// DescribeApplication API operation.
type DescribeApplicationResponse struct {
	*DescribeApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApplication request.
func (r *DescribeApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
