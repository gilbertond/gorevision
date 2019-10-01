// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAccountLimitsInput struct {
	_ struct{} `type:"structure"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `type:"string"`

	// The maximum number of results to return with this call.
	PageSize *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DescribeAccountLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAccountLimitsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAccountLimitsInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAccountLimitsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the limits.
	Limits []Limit `type:"list"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s DescribeAccountLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAccountLimits = "DescribeAccountLimits"

// DescribeAccountLimitsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the current Elastic Load Balancing resource limits for your AWS
// account.
//
// For more information, see Limits for Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using DescribeAccountLimitsRequest.
//    req := client.DescribeAccountLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DescribeAccountLimits
func (c *Client) DescribeAccountLimitsRequest(input *DescribeAccountLimitsInput) DescribeAccountLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeAccountLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAccountLimitsInput{}
	}

	req := c.newRequest(op, input, &DescribeAccountLimitsOutput{})
	return DescribeAccountLimitsRequest{Request: req, Input: input, Copy: c.DescribeAccountLimitsRequest}
}

// DescribeAccountLimitsRequest is the request type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsRequest struct {
	*aws.Request
	Input *DescribeAccountLimitsInput
	Copy  func(*DescribeAccountLimitsInput) DescribeAccountLimitsRequest
}

// Send marshals and sends the DescribeAccountLimits API request.
func (r DescribeAccountLimitsRequest) Send(ctx context.Context) (*DescribeAccountLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAccountLimitsResponse{
		DescribeAccountLimitsOutput: r.Request.Data.(*DescribeAccountLimitsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAccountLimitsResponse is the response type for the
// DescribeAccountLimits API operation.
type DescribeAccountLimitsResponse struct {
	*DescribeAccountLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAccountLimits request.
func (r *DescribeAccountLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
