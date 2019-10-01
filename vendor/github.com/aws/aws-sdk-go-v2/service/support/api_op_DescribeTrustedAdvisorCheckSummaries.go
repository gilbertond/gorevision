// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTrustedAdvisorCheckSummariesInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the Trusted Advisor checks.
	//
	// CheckIds is a required field
	CheckIds []string `locationName:"checkIds" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorCheckSummariesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrustedAdvisorCheckSummariesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrustedAdvisorCheckSummariesInput"}

	if s.CheckIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("CheckIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The summaries of the Trusted Advisor checks returned by the DescribeTrustedAdvisorCheckSummaries
// operation.
type DescribeTrustedAdvisorCheckSummariesOutput struct {
	_ struct{} `type:"structure"`

	// The summary information for the requested Trusted Advisor checks.
	//
	// Summaries is a required field
	Summaries []TrustedAdvisorCheckSummary `locationName:"summaries" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorCheckSummariesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrustedAdvisorCheckSummaries = "DescribeTrustedAdvisorCheckSummaries"

// DescribeTrustedAdvisorCheckSummariesRequest returns a request value for making API operation for
// AWS Support.
//
// Returns the summaries of the results of the Trusted Advisor checks that have
// the specified check IDs. Check IDs can be obtained by calling DescribeTrustedAdvisorChecks.
//
// The response contains an array of TrustedAdvisorCheckSummary objects.
//
//    // Example sending a request using DescribeTrustedAdvisorCheckSummariesRequest.
//    req := client.DescribeTrustedAdvisorCheckSummariesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorCheckSummaries
func (c *Client) DescribeTrustedAdvisorCheckSummariesRequest(input *DescribeTrustedAdvisorCheckSummariesInput) DescribeTrustedAdvisorCheckSummariesRequest {
	op := &aws.Operation{
		Name:       opDescribeTrustedAdvisorCheckSummaries,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTrustedAdvisorCheckSummariesInput{}
	}

	req := c.newRequest(op, input, &DescribeTrustedAdvisorCheckSummariesOutput{})
	return DescribeTrustedAdvisorCheckSummariesRequest{Request: req, Input: input, Copy: c.DescribeTrustedAdvisorCheckSummariesRequest}
}

// DescribeTrustedAdvisorCheckSummariesRequest is the request type for the
// DescribeTrustedAdvisorCheckSummaries API operation.
type DescribeTrustedAdvisorCheckSummariesRequest struct {
	*aws.Request
	Input *DescribeTrustedAdvisorCheckSummariesInput
	Copy  func(*DescribeTrustedAdvisorCheckSummariesInput) DescribeTrustedAdvisorCheckSummariesRequest
}

// Send marshals and sends the DescribeTrustedAdvisorCheckSummaries API request.
func (r DescribeTrustedAdvisorCheckSummariesRequest) Send(ctx context.Context) (*DescribeTrustedAdvisorCheckSummariesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrustedAdvisorCheckSummariesResponse{
		DescribeTrustedAdvisorCheckSummariesOutput: r.Request.Data.(*DescribeTrustedAdvisorCheckSummariesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTrustedAdvisorCheckSummariesResponse is the response type for the
// DescribeTrustedAdvisorCheckSummaries API operation.
type DescribeTrustedAdvisorCheckSummariesResponse struct {
	*DescribeTrustedAdvisorCheckSummariesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrustedAdvisorCheckSummaries request.
func (r *DescribeTrustedAdvisorCheckSummariesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
