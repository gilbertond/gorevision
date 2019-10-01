// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDatastoresInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in this request.
	//
	// The default value is 100.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatastoresInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatastoresInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatastoresInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDatastoresInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDatastoresOutput struct {
	_ struct{} `type:"structure"`

	// A list of "DatastoreSummary" objects.
	DatastoreSummaries []DatastoreSummary `locationName:"datastoreSummaries" type:"list"`

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatastoresOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDatastoresOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DatastoreSummaries != nil {
		v := s.DatastoreSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "datastoreSummaries", metadata)
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

const opListDatastores = "ListDatastores"

// ListDatastoresRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves a list of data stores.
//
//    // Example sending a request using ListDatastoresRequest.
//    req := client.ListDatastoresRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/ListDatastores
func (c *Client) ListDatastoresRequest(input *ListDatastoresInput) ListDatastoresRequest {
	op := &aws.Operation{
		Name:       opListDatastores,
		HTTPMethod: "GET",
		HTTPPath:   "/datastores",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDatastoresInput{}
	}

	req := c.newRequest(op, input, &ListDatastoresOutput{})
	return ListDatastoresRequest{Request: req, Input: input, Copy: c.ListDatastoresRequest}
}

// ListDatastoresRequest is the request type for the
// ListDatastores API operation.
type ListDatastoresRequest struct {
	*aws.Request
	Input *ListDatastoresInput
	Copy  func(*ListDatastoresInput) ListDatastoresRequest
}

// Send marshals and sends the ListDatastores API request.
func (r ListDatastoresRequest) Send(ctx context.Context) (*ListDatastoresResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatastoresResponse{
		ListDatastoresOutput: r.Request.Data.(*ListDatastoresOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDatastoresRequestPaginator returns a paginator for ListDatastores.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDatastoresRequest(input)
//   p := iotanalytics.NewListDatastoresRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDatastoresPaginator(req ListDatastoresRequest) ListDatastoresPaginator {
	return ListDatastoresPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDatastoresInput
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

// ListDatastoresPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDatastoresPaginator struct {
	aws.Pager
}

func (p *ListDatastoresPaginator) CurrentPage() *ListDatastoresOutput {
	return p.Pager.CurrentPage().(*ListDatastoresOutput)
}

// ListDatastoresResponse is the response type for the
// ListDatastores API operation.
type ListDatastoresResponse struct {
	*ListDatastoresOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatastores request.
func (r *ListDatastoresResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
