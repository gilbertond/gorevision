// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSkillsInput struct {
	_ struct{} `type:"structure"`

	// Whether the skill is enabled under the user's account, or if it requires
	// linking to be used.
	EnablementType EnablementTypeFilter `type:"string" enum:"true"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved. Required.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string `min:"1" type:"string"`

	// The ARN of the skill group for which to list enabled skills. Required.
	SkillGroupArn *string `type:"string"`

	// Whether the skill is publicly available or is a private skill.
	SkillType SkillTypeFilter `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListSkillsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSkillsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSkillsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSkillsOutput struct {
	_ struct{} `type:"structure"`

	// The token returned to indicate that there is more data available.
	NextToken *string `min:"1" type:"string"`

	// The list of enabled skills requested. Required.
	SkillSummaries []SkillSummary `type:"list"`
}

// String returns the string representation
func (s ListSkillsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSkills = "ListSkills"

// ListSkillsRequest returns a request value for making API operation for
// Alexa For Business.
//
// Lists all enabled skills in a specific skill group.
//
//    // Example sending a request using ListSkillsRequest.
//    req := client.ListSkillsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ListSkills
func (c *Client) ListSkillsRequest(input *ListSkillsInput) ListSkillsRequest {
	op := &aws.Operation{
		Name:       opListSkills,
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
		input = &ListSkillsInput{}
	}

	req := c.newRequest(op, input, &ListSkillsOutput{})
	return ListSkillsRequest{Request: req, Input: input, Copy: c.ListSkillsRequest}
}

// ListSkillsRequest is the request type for the
// ListSkills API operation.
type ListSkillsRequest struct {
	*aws.Request
	Input *ListSkillsInput
	Copy  func(*ListSkillsInput) ListSkillsRequest
}

// Send marshals and sends the ListSkills API request.
func (r ListSkillsRequest) Send(ctx context.Context) (*ListSkillsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSkillsResponse{
		ListSkillsOutput: r.Request.Data.(*ListSkillsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSkillsRequestPaginator returns a paginator for ListSkills.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSkillsRequest(input)
//   p := alexaforbusiness.NewListSkillsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSkillsPaginator(req ListSkillsRequest) ListSkillsPaginator {
	return ListSkillsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSkillsInput
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

// ListSkillsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSkillsPaginator struct {
	aws.Pager
}

func (p *ListSkillsPaginator) CurrentPage() *ListSkillsOutput {
	return p.Pager.CurrentPage().(*ListSkillsOutput)
}

// ListSkillsResponse is the response type for the
// ListSkills API operation.
type ListSkillsResponse struct {
	*ListSkillsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSkills request.
func (r *ListSkillsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
