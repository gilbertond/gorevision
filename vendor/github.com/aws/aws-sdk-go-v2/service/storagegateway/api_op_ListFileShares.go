// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// ListFileShareInput
type ListFileSharesInput struct {
	_ struct{} `type:"structure"`

	// The Amazon resource Name (ARN) of the gateway whose file shares you want
	// to list. If this field is not present, all file shares under your account
	// are listed.
	GatewayARN *string `min:"50" type:"string"`

	// The maximum number of file shares to return in the response. The value must
	// be an integer with a value greater than zero. Optional.
	Limit *int64 `min:"1" type:"integer"`

	// Opaque pagination token returned from a previous ListFileShares operation.
	// If present, Marker specifies where to continue the list from after a previous
	// call to ListFileShares. Optional.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFileSharesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFileSharesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFileSharesInput"}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// ListFileShareOutput
type ListFileSharesOutput struct {
	_ struct{} `type:"structure"`

	// An array of information about the file gateway's file shares.
	FileShareInfoList []FileShareInfo `type:"list"`

	// If the request includes Marker, the response returns that value in this field.
	Marker *string `min:"1" type:"string"`

	// If a value is present, there are more file shares to return. In a subsequent
	// request, use NextMarker as the value for Marker to retrieve the next set
	// of file shares.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFileSharesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListFileShares = "ListFileShares"

// ListFileSharesRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Gets a list of the file shares for a specific file gateway, or the list of
// file shares that belong to the calling user account. This operation is only
// supported for file gateways.
//
//    // Example sending a request using ListFileSharesRequest.
//    req := client.ListFileSharesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListFileShares
func (c *Client) ListFileSharesRequest(input *ListFileSharesInput) ListFileSharesRequest {
	op := &aws.Operation{
		Name:       opListFileShares,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFileSharesInput{}
	}

	req := c.newRequest(op, input, &ListFileSharesOutput{})
	return ListFileSharesRequest{Request: req, Input: input, Copy: c.ListFileSharesRequest}
}

// ListFileSharesRequest is the request type for the
// ListFileShares API operation.
type ListFileSharesRequest struct {
	*aws.Request
	Input *ListFileSharesInput
	Copy  func(*ListFileSharesInput) ListFileSharesRequest
}

// Send marshals and sends the ListFileShares API request.
func (r ListFileSharesRequest) Send(ctx context.Context) (*ListFileSharesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFileSharesResponse{
		ListFileSharesOutput: r.Request.Data.(*ListFileSharesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFileSharesRequestPaginator returns a paginator for ListFileShares.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFileSharesRequest(input)
//   p := storagegateway.NewListFileSharesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFileSharesPaginator(req ListFileSharesRequest) ListFileSharesPaginator {
	return ListFileSharesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFileSharesInput
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

// ListFileSharesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFileSharesPaginator struct {
	aws.Pager
}

func (p *ListFileSharesPaginator) CurrentPage() *ListFileSharesOutput {
	return p.Pager.CurrentPage().(*ListFileSharesOutput)
}

// ListFileSharesResponse is the response type for the
// ListFileShares API operation.
type ListFileSharesResponse struct {
	*ListFileSharesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFileShares request.
func (r *ListFileSharesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
