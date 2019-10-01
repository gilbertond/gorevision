// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AdminListUserAuthEventsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of authentication events to return.
	MaxResults *int64 `type:"integer"`

	// A pagination token.
	NextToken *string `min:"1" type:"string"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user pool username or an alias.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AdminListUserAuthEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminListUserAuthEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminListUserAuthEventsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AdminListUserAuthEventsOutput struct {
	_ struct{} `type:"structure"`

	// The response object. It includes the EventID, EventType, CreationDate, EventRisk,
	// and EventResponse.
	AuthEvents []AuthEventType `type:"list"`

	// A pagination token.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s AdminListUserAuthEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminListUserAuthEvents = "AdminListUserAuthEvents"

// AdminListUserAuthEventsRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists a history of user activity and any risks detected as part of Amazon
// Cognito advanced security.
//
//    // Example sending a request using AdminListUserAuthEventsRequest.
//    req := client.AdminListUserAuthEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminListUserAuthEvents
func (c *Client) AdminListUserAuthEventsRequest(input *AdminListUserAuthEventsInput) AdminListUserAuthEventsRequest {
	op := &aws.Operation{
		Name:       opAdminListUserAuthEvents,
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
		input = &AdminListUserAuthEventsInput{}
	}

	req := c.newRequest(op, input, &AdminListUserAuthEventsOutput{})
	return AdminListUserAuthEventsRequest{Request: req, Input: input, Copy: c.AdminListUserAuthEventsRequest}
}

// AdminListUserAuthEventsRequest is the request type for the
// AdminListUserAuthEvents API operation.
type AdminListUserAuthEventsRequest struct {
	*aws.Request
	Input *AdminListUserAuthEventsInput
	Copy  func(*AdminListUserAuthEventsInput) AdminListUserAuthEventsRequest
}

// Send marshals and sends the AdminListUserAuthEvents API request.
func (r AdminListUserAuthEventsRequest) Send(ctx context.Context) (*AdminListUserAuthEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminListUserAuthEventsResponse{
		AdminListUserAuthEventsOutput: r.Request.Data.(*AdminListUserAuthEventsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewAdminListUserAuthEventsRequestPaginator returns a paginator for AdminListUserAuthEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.AdminListUserAuthEventsRequest(input)
//   p := cognitoidentityprovider.NewAdminListUserAuthEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewAdminListUserAuthEventsPaginator(req AdminListUserAuthEventsRequest) AdminListUserAuthEventsPaginator {
	return AdminListUserAuthEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *AdminListUserAuthEventsInput
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

// AdminListUserAuthEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type AdminListUserAuthEventsPaginator struct {
	aws.Pager
}

func (p *AdminListUserAuthEventsPaginator) CurrentPage() *AdminListUserAuthEventsOutput {
	return p.Pager.CurrentPage().(*AdminListUserAuthEventsOutput)
}

// AdminListUserAuthEventsResponse is the response type for the
// AdminListUserAuthEvents API operation.
type AdminListUserAuthEventsResponse struct {
	*AdminListUserAuthEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminListUserAuthEvents request.
func (r *AdminListUserAuthEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
