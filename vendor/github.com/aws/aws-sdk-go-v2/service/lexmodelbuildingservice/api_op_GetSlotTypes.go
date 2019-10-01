// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetSlotTypesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of slot types to return in the response. The default is
	// 10.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Substring to match in slot type names. A slot type will be returned if any
	// part of its name matches the substring. For example, "xyz" matches both "xyzabc"
	// and "abcxyz."
	NameContains *string `location:"querystring" locationName:"nameContains" min:"1" type:"string"`

	// A pagination token that fetches the next page of slot types. If the response
	// to this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch next page of slot types, specify the pagination token
	// in the next request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetSlotTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSlotTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSlotTypesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSlotTypesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NameContains != nil {
		v := *s.NameContains

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nameContains", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSlotTypesOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of slot types.
	NextToken *string `locationName:"nextToken" type:"string"`

	// An array of objects, one for each slot type, that provides information such
	// as the name of the slot type, the version, and a description.
	SlotTypes []SlotTypeMetadata `locationName:"slotTypes" type:"list"`
}

// String returns the string representation
func (s GetSlotTypesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSlotTypesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlotTypes != nil {
		v := s.SlotTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "slotTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetSlotTypes = "GetSlotTypes"

// GetSlotTypesRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns slot type information as follows:
//
//    * If you specify the nameContains field, returns the $LATEST version of
//    all slot types that contain the specified string.
//
//    * If you don't specify the nameContains field, returns information about
//    the $LATEST version of all slot types.
//
// The operation requires permission for the lex:GetSlotTypes action.
//
//    // Example sending a request using GetSlotTypesRequest.
//    req := client.GetSlotTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetSlotTypes
func (c *Client) GetSlotTypesRequest(input *GetSlotTypesInput) GetSlotTypesRequest {
	op := &aws.Operation{
		Name:       opGetSlotTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/slottypes/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetSlotTypesInput{}
	}

	req := c.newRequest(op, input, &GetSlotTypesOutput{})
	return GetSlotTypesRequest{Request: req, Input: input, Copy: c.GetSlotTypesRequest}
}

// GetSlotTypesRequest is the request type for the
// GetSlotTypes API operation.
type GetSlotTypesRequest struct {
	*aws.Request
	Input *GetSlotTypesInput
	Copy  func(*GetSlotTypesInput) GetSlotTypesRequest
}

// Send marshals and sends the GetSlotTypes API request.
func (r GetSlotTypesRequest) Send(ctx context.Context) (*GetSlotTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSlotTypesResponse{
		GetSlotTypesOutput: r.Request.Data.(*GetSlotTypesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetSlotTypesRequestPaginator returns a paginator for GetSlotTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetSlotTypesRequest(input)
//   p := lexmodelbuildingservice.NewGetSlotTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetSlotTypesPaginator(req GetSlotTypesRequest) GetSlotTypesPaginator {
	return GetSlotTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetSlotTypesInput
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

// GetSlotTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetSlotTypesPaginator struct {
	aws.Pager
}

func (p *GetSlotTypesPaginator) CurrentPage() *GetSlotTypesOutput {
	return p.Pager.CurrentPage().(*GetSlotTypesOutput)
}

// GetSlotTypesResponse is the response type for the
// GetSlotTypes API operation.
type GetSlotTypesResponse struct {
	*GetSlotTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSlotTypes request.
func (r *GetSlotTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
