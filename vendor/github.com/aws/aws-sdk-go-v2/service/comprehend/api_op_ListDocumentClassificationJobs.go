// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDocumentClassificationJobsInput struct {
	_ struct{} `type:"structure"`

	// Filters the jobs that are returned. You can filter jobs on their names, status,
	// or the date and time that they were submitted. You can only set one filter
	// at a time.
	Filter *DocumentClassificationJobFilter `type:"structure"`

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDocumentClassificationJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDocumentClassificationJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDocumentClassificationJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDocumentClassificationJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list containing the properties of each job returned.
	DocumentClassificationJobPropertiesList []DocumentClassificationJobProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDocumentClassificationJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDocumentClassificationJobs = "ListDocumentClassificationJobs"

// ListDocumentClassificationJobsRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets a list of the documentation classification jobs that you have submitted.
//
//    // Example sending a request using ListDocumentClassificationJobsRequest.
//    req := client.ListDocumentClassificationJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListDocumentClassificationJobs
func (c *Client) ListDocumentClassificationJobsRequest(input *ListDocumentClassificationJobsInput) ListDocumentClassificationJobsRequest {
	op := &aws.Operation{
		Name:       opListDocumentClassificationJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDocumentClassificationJobsInput{}
	}

	req := c.newRequest(op, input, &ListDocumentClassificationJobsOutput{})
	return ListDocumentClassificationJobsRequest{Request: req, Input: input, Copy: c.ListDocumentClassificationJobsRequest}
}

// ListDocumentClassificationJobsRequest is the request type for the
// ListDocumentClassificationJobs API operation.
type ListDocumentClassificationJobsRequest struct {
	*aws.Request
	Input *ListDocumentClassificationJobsInput
	Copy  func(*ListDocumentClassificationJobsInput) ListDocumentClassificationJobsRequest
}

// Send marshals and sends the ListDocumentClassificationJobs API request.
func (r ListDocumentClassificationJobsRequest) Send(ctx context.Context) (*ListDocumentClassificationJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDocumentClassificationJobsResponse{
		ListDocumentClassificationJobsOutput: r.Request.Data.(*ListDocumentClassificationJobsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDocumentClassificationJobsRequestPaginator returns a paginator for ListDocumentClassificationJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDocumentClassificationJobsRequest(input)
//   p := comprehend.NewListDocumentClassificationJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDocumentClassificationJobsPaginator(req ListDocumentClassificationJobsRequest) ListDocumentClassificationJobsPaginator {
	return ListDocumentClassificationJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDocumentClassificationJobsInput
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

// ListDocumentClassificationJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDocumentClassificationJobsPaginator struct {
	aws.Pager
}

func (p *ListDocumentClassificationJobsPaginator) CurrentPage() *ListDocumentClassificationJobsOutput {
	return p.Pager.CurrentPage().(*ListDocumentClassificationJobsOutput)
}

// ListDocumentClassificationJobsResponse is the response type for the
// ListDocumentClassificationJobs API operation.
type ListDocumentClassificationJobsResponse struct {
	*ListDocumentClassificationJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDocumentClassificationJobs request.
func (r *ListDocumentClassificationJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
