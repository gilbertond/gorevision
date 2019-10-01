// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RegenerateSecurityTokenInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The bot ID.
	//
	// BotId is a required field
	BotId *string `location:"uri" locationName:"botId" type:"string" required:"true"`
}

// String returns the string representation
func (s RegenerateSecurityTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegenerateSecurityTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegenerateSecurityTokenInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.BotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegenerateSecurityTokenInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotId != nil {
		v := *s.BotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RegenerateSecurityTokenOutput struct {
	_ struct{} `type:"structure"`

	// A resource that allows Enterprise account administrators to configure an
	// interface to receive events from Amazon Chime.
	Bot *Bot `type:"structure"`
}

// String returns the string representation
func (s RegenerateSecurityTokenOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegenerateSecurityTokenOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bot != nil {
		v := s.Bot

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Bot", v, metadata)
	}
	return nil
}

const opRegenerateSecurityToken = "RegenerateSecurityToken"

// RegenerateSecurityTokenRequest returns a request value for making API operation for
// Amazon Chime.
//
// Regenerates the security token for a bot.
//
//    // Example sending a request using RegenerateSecurityTokenRequest.
//    req := client.RegenerateSecurityTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/RegenerateSecurityToken
func (c *Client) RegenerateSecurityTokenRequest(input *RegenerateSecurityTokenInput) RegenerateSecurityTokenRequest {
	op := &aws.Operation{
		Name:       opRegenerateSecurityToken,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/bots/{botId}?operation=regenerate-security-token",
	}

	if input == nil {
		input = &RegenerateSecurityTokenInput{}
	}

	req := c.newRequest(op, input, &RegenerateSecurityTokenOutput{})
	return RegenerateSecurityTokenRequest{Request: req, Input: input, Copy: c.RegenerateSecurityTokenRequest}
}

// RegenerateSecurityTokenRequest is the request type for the
// RegenerateSecurityToken API operation.
type RegenerateSecurityTokenRequest struct {
	*aws.Request
	Input *RegenerateSecurityTokenInput
	Copy  func(*RegenerateSecurityTokenInput) RegenerateSecurityTokenRequest
}

// Send marshals and sends the RegenerateSecurityToken API request.
func (r RegenerateSecurityTokenRequest) Send(ctx context.Context) (*RegenerateSecurityTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegenerateSecurityTokenResponse{
		RegenerateSecurityTokenOutput: r.Request.Data.(*RegenerateSecurityTokenOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegenerateSecurityTokenResponse is the response type for the
// RegenerateSecurityToken API operation.
type RegenerateSecurityTokenResponse struct {
	*RegenerateSecurityTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegenerateSecurityToken request.
func (r *RegenerateSecurityTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
