// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector that specifies the GuardDuty service whose findings
	// you want to list.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// Represents the criteria used for querying findings.
	FindingCriteria *FindingCriteria `locationName:"findingCriteria" type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls
	// to the action fill nextToken in the request with the value of NextToken from
	// the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Represents the criteria used for sorting findings.
	SortCriteria *SortCriteria `locationName:"sortCriteria" type:"structure"`
}

// String returns the string representation
func (s ListFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFindingsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFindingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FindingCriteria != nil {
		v := s.FindingCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "findingCriteria", v, metadata)
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
	if s.SortCriteria != nil {
		v := s.SortCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sortCriteria", v, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListFindingsOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the findings you are listing.
	//
	// FindingIds is a required field
	FindingIds []string `locationName:"findingIds" type:"list" required:"true"`

	// Pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListFindingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFindingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FindingIds != nil {
		v := s.FindingIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "findingIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
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

const opListFindings = "ListFindings"

// ListFindingsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Lists Amazon GuardDuty findings for the specified detector ID.
//
//    // Example sending a request using ListFindingsRequest.
//    req := client.ListFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListFindings
func (c *Client) ListFindingsRequest(input *ListFindingsInput) ListFindingsRequest {
	op := &aws.Operation{
		Name:       opListFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/findings",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFindingsInput{}
	}

	req := c.newRequest(op, input, &ListFindingsOutput{})
	return ListFindingsRequest{Request: req, Input: input, Copy: c.ListFindingsRequest}
}

// ListFindingsRequest is the request type for the
// ListFindings API operation.
type ListFindingsRequest struct {
	*aws.Request
	Input *ListFindingsInput
	Copy  func(*ListFindingsInput) ListFindingsRequest
}

// Send marshals and sends the ListFindings API request.
func (r ListFindingsRequest) Send(ctx context.Context) (*ListFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFindingsResponse{
		ListFindingsOutput: r.Request.Data.(*ListFindingsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFindingsRequestPaginator returns a paginator for ListFindings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFindingsRequest(input)
//   p := guardduty.NewListFindingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFindingsPaginator(req ListFindingsRequest) ListFindingsPaginator {
	return ListFindingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFindingsInput
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

// ListFindingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFindingsPaginator struct {
	aws.Pager
}

func (p *ListFindingsPaginator) CurrentPage() *ListFindingsOutput {
	return p.Pager.CurrentPage().(*ListFindingsOutput)
}

// ListFindingsResponse is the response type for the
// ListFindings API operation.
type ListFindingsResponse struct {
	*ListFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFindings request.
func (r *ListFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
