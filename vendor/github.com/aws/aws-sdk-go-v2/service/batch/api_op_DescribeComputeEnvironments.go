// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeComputeEnvironmentsInput struct {
	_ struct{} `type:"structure"`

	// A list of up to 100 compute environment names or full Amazon Resource Name
	// (ARN) entries.
	ComputeEnvironments []string `locationName:"computeEnvironments" type:"list"`

	// The maximum number of cluster results returned by DescribeComputeEnvironments
	// in paginated output. When this parameter is used, DescribeComputeEnvironments
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeComputeEnvironments request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// DescribeComputeEnvironments returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeComputeEnvironments
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeComputeEnvironmentsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeComputeEnvironmentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ComputeEnvironments != nil {
		v := s.ComputeEnvironments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "computeEnvironments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeComputeEnvironmentsOutput struct {
	_ struct{} `type:"structure"`

	// The list of compute environments.
	ComputeEnvironments []ComputeEnvironmentDetail `locationName:"computeEnvironments" type:"list"`

	// The nextToken value to include in a future DescribeComputeEnvironments request.
	// When the results of a DescribeJobDefinitions request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeComputeEnvironmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeComputeEnvironmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ComputeEnvironments != nil {
		v := s.ComputeEnvironments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "computeEnvironments", metadata)
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

const opDescribeComputeEnvironments = "DescribeComputeEnvironments"

// DescribeComputeEnvironmentsRequest returns a request value for making API operation for
// AWS Batch.
//
// Describes one or more of your compute environments.
//
// If you are using an unmanaged compute environment, you can use the DescribeComputeEnvironment
// operation to determine the ecsClusterArn that you should launch your Amazon
// ECS container instances into.
//
//    // Example sending a request using DescribeComputeEnvironmentsRequest.
//    req := client.DescribeComputeEnvironmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeComputeEnvironments
func (c *Client) DescribeComputeEnvironmentsRequest(input *DescribeComputeEnvironmentsInput) DescribeComputeEnvironmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeComputeEnvironments,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/describecomputeenvironments",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeComputeEnvironmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeComputeEnvironmentsOutput{})
	return DescribeComputeEnvironmentsRequest{Request: req, Input: input, Copy: c.DescribeComputeEnvironmentsRequest}
}

// DescribeComputeEnvironmentsRequest is the request type for the
// DescribeComputeEnvironments API operation.
type DescribeComputeEnvironmentsRequest struct {
	*aws.Request
	Input *DescribeComputeEnvironmentsInput
	Copy  func(*DescribeComputeEnvironmentsInput) DescribeComputeEnvironmentsRequest
}

// Send marshals and sends the DescribeComputeEnvironments API request.
func (r DescribeComputeEnvironmentsRequest) Send(ctx context.Context) (*DescribeComputeEnvironmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeComputeEnvironmentsResponse{
		DescribeComputeEnvironmentsOutput: r.Request.Data.(*DescribeComputeEnvironmentsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeComputeEnvironmentsRequestPaginator returns a paginator for DescribeComputeEnvironments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeComputeEnvironmentsRequest(input)
//   p := batch.NewDescribeComputeEnvironmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeComputeEnvironmentsPaginator(req DescribeComputeEnvironmentsRequest) DescribeComputeEnvironmentsPaginator {
	return DescribeComputeEnvironmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeComputeEnvironmentsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeComputeEnvironmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeComputeEnvironmentsPaginator struct {
	aws.Pager
}

func (p *DescribeComputeEnvironmentsPaginator) CurrentPage() *DescribeComputeEnvironmentsOutput {
	return p.Pager.CurrentPage().(*DescribeComputeEnvironmentsOutput)
}

// DescribeComputeEnvironmentsResponse is the response type for the
// DescribeComputeEnvironments API operation.
type DescribeComputeEnvironmentsResponse struct {
	*DescribeComputeEnvironmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeComputeEnvironments request.
func (r *DescribeComputeEnvironmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
