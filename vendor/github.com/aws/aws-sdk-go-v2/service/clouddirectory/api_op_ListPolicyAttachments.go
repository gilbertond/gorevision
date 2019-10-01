// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListPolicyAttachmentsInput struct {
	_ struct{} `type:"structure"`

	// Represents the manner and timing in which the successful write or update
	// of an object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// objects reside. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The reference that identifies the policy object.
	//
	// PolicyReference is a required field
	PolicyReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListPolicyAttachmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPolicyAttachmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPolicyAttachmentsInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.PolicyReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPolicyAttachmentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyReference != nil {
		v := s.PolicyReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PolicyReference", v, metadata)
	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListPolicyAttachmentsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// A list of ObjectIdentifiers to which the policy is attached.
	ObjectIdentifiers []string `type:"list"`
}

// String returns the string representation
func (s ListPolicyAttachmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPolicyAttachmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectIdentifiers != nil {
		v := s.ObjectIdentifiers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ObjectIdentifiers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListPolicyAttachments = "ListPolicyAttachments"

// ListPolicyAttachmentsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns all of the ObjectIdentifiers to which a given policy is attached.
//
//    // Example sending a request using ListPolicyAttachmentsRequest.
//    req := client.ListPolicyAttachmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListPolicyAttachments
func (c *Client) ListPolicyAttachmentsRequest(input *ListPolicyAttachmentsInput) ListPolicyAttachmentsRequest {
	op := &aws.Operation{
		Name:       opListPolicyAttachments,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/policy/attachment",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPolicyAttachmentsInput{}
	}

	req := c.newRequest(op, input, &ListPolicyAttachmentsOutput{})
	return ListPolicyAttachmentsRequest{Request: req, Input: input, Copy: c.ListPolicyAttachmentsRequest}
}

// ListPolicyAttachmentsRequest is the request type for the
// ListPolicyAttachments API operation.
type ListPolicyAttachmentsRequest struct {
	*aws.Request
	Input *ListPolicyAttachmentsInput
	Copy  func(*ListPolicyAttachmentsInput) ListPolicyAttachmentsRequest
}

// Send marshals and sends the ListPolicyAttachments API request.
func (r ListPolicyAttachmentsRequest) Send(ctx context.Context) (*ListPolicyAttachmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPolicyAttachmentsResponse{
		ListPolicyAttachmentsOutput: r.Request.Data.(*ListPolicyAttachmentsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPolicyAttachmentsRequestPaginator returns a paginator for ListPolicyAttachments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPolicyAttachmentsRequest(input)
//   p := clouddirectory.NewListPolicyAttachmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPolicyAttachmentsPaginator(req ListPolicyAttachmentsRequest) ListPolicyAttachmentsPaginator {
	return ListPolicyAttachmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPolicyAttachmentsInput
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

// ListPolicyAttachmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPolicyAttachmentsPaginator struct {
	aws.Pager
}

func (p *ListPolicyAttachmentsPaginator) CurrentPage() *ListPolicyAttachmentsOutput {
	return p.Pager.CurrentPage().(*ListPolicyAttachmentsOutput)
}

// ListPolicyAttachmentsResponse is the response type for the
// ListPolicyAttachments API operation.
type ListPolicyAttachmentsResponse struct {
	*ListPolicyAttachmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicyAttachments request.
func (r *ListPolicyAttachmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
