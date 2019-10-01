// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the list projects operation.
type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	// Optional. If no Amazon Resource Name (ARN) is specified, then AWS Device
	// Farm returns a list of all projects for the AWS account. You can also specify
	// a project ARN.
	Arn *string `locationName:"arn" min:"32" type:"string"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProjectsInput"}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list projects request.
type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned, which can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the projects.
	Projects []Project `locationName:"projects" type:"list"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about projects.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListProjects
func (c *Client) ListProjectsRequest(input *ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProjectsInput{}
	}

	req := c.newRequest(op, input, &ListProjectsOutput{})
	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *ListProjectsInput
	Copy  func(*ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProjectsRequestPaginator returns a paginator for ListProjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProjectsRequest(input)
//   p := devicefarm.NewListProjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProjectsPaginator(req ListProjectsRequest) ListProjectsPaginator {
	return ListProjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProjectsInput
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

// ListProjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProjectsPaginator struct {
	aws.Pager
}

func (p *ListProjectsPaginator) CurrentPage() *ListProjectsOutput {
	return p.Pager.CurrentPage().(*ListProjectsOutput)
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
