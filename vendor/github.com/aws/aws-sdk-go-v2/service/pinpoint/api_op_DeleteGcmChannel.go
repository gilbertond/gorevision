// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteGcmChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGcmChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGcmChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGcmChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGcmChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGcmChannelOutput struct {
	_ struct{} `type:"structure" payload:"GCMChannelResponse"`

	// Provides information about the status and settings of the GCM channel for
	// an application. The GCM channel enables Amazon Pinpoint to send push notifications
	// through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging
	// (GCM), service.
	//
	// GCMChannelResponse is a required field
	GCMChannelResponse *GCMChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteGcmChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGcmChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GCMChannelResponse != nil {
		v := s.GCMChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "GCMChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteGcmChannel = "DeleteGcmChannel"

// DeleteGcmChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Disables the GCM channel for an application and deletes any existing settings
// for the channel.
//
//    // Example sending a request using DeleteGcmChannelRequest.
//    req := client.DeleteGcmChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteGcmChannel
func (c *Client) DeleteGcmChannelRequest(input *DeleteGcmChannelInput) DeleteGcmChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteGcmChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/gcm",
	}

	if input == nil {
		input = &DeleteGcmChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteGcmChannelOutput{})
	return DeleteGcmChannelRequest{Request: req, Input: input, Copy: c.DeleteGcmChannelRequest}
}

// DeleteGcmChannelRequest is the request type for the
// DeleteGcmChannel API operation.
type DeleteGcmChannelRequest struct {
	*aws.Request
	Input *DeleteGcmChannelInput
	Copy  func(*DeleteGcmChannelInput) DeleteGcmChannelRequest
}

// Send marshals and sends the DeleteGcmChannel API request.
func (r DeleteGcmChannelRequest) Send(ctx context.Context) (*DeleteGcmChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGcmChannelResponse{
		DeleteGcmChannelOutput: r.Request.Data.(*DeleteGcmChannelOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGcmChannelResponse is the response type for the
// DeleteGcmChannel API operation.
type DeleteGcmChannelResponse struct {
	*DeleteGcmChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGcmChannel request.
func (r *DeleteGcmChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
