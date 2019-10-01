// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListResourcesForTagOptionInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The resource type.
	//
	//    * Portfolio
	//
	//    * Product
	ResourceType *string `type:"string"`

	// The TagOption identifier.
	//
	// TagOptionId is a required field
	TagOptionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResourcesForTagOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourcesForTagOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourcesForTagOptionInput"}

	if s.TagOptionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagOptionId"))
	}
	if s.TagOptionId != nil && len(*s.TagOptionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagOptionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResourcesForTagOptionOutput struct {
	_ struct{} `type:"structure"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// Information about the resources.
	ResourceDetails []ResourceDetail `type:"list"`
}

// String returns the string representation
func (s ListResourcesForTagOptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResourcesForTagOption = "ListResourcesForTagOption"

// ListResourcesForTagOptionRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the resources associated with the specified TagOption.
//
//    // Example sending a request using ListResourcesForTagOptionRequest.
//    req := client.ListResourcesForTagOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListResourcesForTagOption
func (c *Client) ListResourcesForTagOptionRequest(input *ListResourcesForTagOptionInput) ListResourcesForTagOptionRequest {
	op := &aws.Operation{
		Name:       opListResourcesForTagOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"PageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListResourcesForTagOptionInput{}
	}

	req := c.newRequest(op, input, &ListResourcesForTagOptionOutput{})
	return ListResourcesForTagOptionRequest{Request: req, Input: input, Copy: c.ListResourcesForTagOptionRequest}
}

// ListResourcesForTagOptionRequest is the request type for the
// ListResourcesForTagOption API operation.
type ListResourcesForTagOptionRequest struct {
	*aws.Request
	Input *ListResourcesForTagOptionInput
	Copy  func(*ListResourcesForTagOptionInput) ListResourcesForTagOptionRequest
}

// Send marshals and sends the ListResourcesForTagOption API request.
func (r ListResourcesForTagOptionRequest) Send(ctx context.Context) (*ListResourcesForTagOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourcesForTagOptionResponse{
		ListResourcesForTagOptionOutput: r.Request.Data.(*ListResourcesForTagOptionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResourcesForTagOptionRequestPaginator returns a paginator for ListResourcesForTagOption.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResourcesForTagOptionRequest(input)
//   p := servicecatalog.NewListResourcesForTagOptionRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResourcesForTagOptionPaginator(req ListResourcesForTagOptionRequest) ListResourcesForTagOptionPaginator {
	return ListResourcesForTagOptionPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListResourcesForTagOptionInput
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

// ListResourcesForTagOptionPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResourcesForTagOptionPaginator struct {
	aws.Pager
}

func (p *ListResourcesForTagOptionPaginator) CurrentPage() *ListResourcesForTagOptionOutput {
	return p.Pager.CurrentPage().(*ListResourcesForTagOptionOutput)
}

// ListResourcesForTagOptionResponse is the response type for the
// ListResourcesForTagOption API operation.
type ListResourcesForTagOptionResponse struct {
	*ListResourcesForTagOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourcesForTagOption request.
func (r *ListResourcesForTagOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
