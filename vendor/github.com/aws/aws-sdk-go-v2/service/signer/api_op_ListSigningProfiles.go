// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListSigningProfilesInput struct {
	_ struct{} `type:"structure"`

	// Designates whether to include profiles with the status of CANCELED.
	IncludeCanceled *bool `location:"querystring" locationName:"includeCanceled" type:"boolean"`

	// The maximum number of profiles to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Value for specifying the next set of paginated results to return. After you
	// receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListSigningProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSigningProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSigningProfilesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSigningProfilesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IncludeCanceled != nil {
		v := *s.IncludeCanceled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeCanceled", protocol.BoolValue(v), metadata)
	}
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

type ListSigningProfilesOutput struct {
	_ struct{} `type:"structure"`

	// Value for specifying the next set of paginated results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of profiles that are available in the AWS account. This includes profiles
	// with the status of CANCELED if the includeCanceled parameter is set to true.
	Profiles []SigningProfile `locationName:"profiles" type:"list"`
}

// String returns the string representation
func (s ListSigningProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSigningProfilesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Profiles != nil {
		v := s.Profiles

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "profiles", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListSigningProfiles = "ListSigningProfiles"

// ListSigningProfilesRequest returns a request value for making API operation for
// AWS Signer.
//
// Lists all available signing profiles in your AWS account. Returns only profiles
// with an ACTIVE status unless the includeCanceled request field is set to
// true. If additional jobs remain to be listed, AWS Signer returns a nextToken
// value. Use this value in subsequent calls to ListSigningJobs to fetch the
// remaining values. You can continue calling ListSigningJobs with your maxResults
// parameter and with new values that AWS Signer returns in the nextToken parameter
// until all of your signing jobs have been returned.
//
//    // Example sending a request using ListSigningProfilesRequest.
//    req := client.ListSigningProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/ListSigningProfiles
func (c *Client) ListSigningProfilesRequest(input *ListSigningProfilesInput) ListSigningProfilesRequest {
	op := &aws.Operation{
		Name:       opListSigningProfiles,
		HTTPMethod: "GET",
		HTTPPath:   "/signing-profiles",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSigningProfilesInput{}
	}

	req := c.newRequest(op, input, &ListSigningProfilesOutput{})
	return ListSigningProfilesRequest{Request: req, Input: input, Copy: c.ListSigningProfilesRequest}
}

// ListSigningProfilesRequest is the request type for the
// ListSigningProfiles API operation.
type ListSigningProfilesRequest struct {
	*aws.Request
	Input *ListSigningProfilesInput
	Copy  func(*ListSigningProfilesInput) ListSigningProfilesRequest
}

// Send marshals and sends the ListSigningProfiles API request.
func (r ListSigningProfilesRequest) Send(ctx context.Context) (*ListSigningProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSigningProfilesResponse{
		ListSigningProfilesOutput: r.Request.Data.(*ListSigningProfilesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSigningProfilesRequestPaginator returns a paginator for ListSigningProfiles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSigningProfilesRequest(input)
//   p := signer.NewListSigningProfilesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSigningProfilesPaginator(req ListSigningProfilesRequest) ListSigningProfilesPaginator {
	return ListSigningProfilesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSigningProfilesInput
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

// ListSigningProfilesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSigningProfilesPaginator struct {
	aws.Pager
}

func (p *ListSigningProfilesPaginator) CurrentPage() *ListSigningProfilesOutput {
	return p.Pager.CurrentPage().(*ListSigningProfilesOutput)
}

// ListSigningProfilesResponse is the response type for the
// ListSigningProfiles API operation.
type ListSigningProfilesResponse struct {
	*ListSigningProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSigningProfiles request.
func (r *ListSigningProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
