// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteChannelInput struct {
	_ struct{} `type:"structure"`

	// ChannelId is a required field
	ChannelId *string `location:"uri" locationName:"channelId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteChannelInput"}

	if s.ChannelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteChannelOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	// A standard channel has two encoding pipelines and a single pipeline channel
	// only has one.
	ChannelClass ChannelClass `locationName:"channelClass" type:"string" enum:"true"`

	Destinations []OutputDestination `locationName:"destinations" type:"list"`

	EgressEndpoints []ChannelEgressEndpoint `locationName:"egressEndpoints" type:"list"`

	// Encoder Settings
	EncoderSettings *EncoderSettings `locationName:"encoderSettings" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	InputAttachments []InputAttachment `locationName:"inputAttachments" type:"list"`

	InputSpecification *InputSpecification `locationName:"inputSpecification" type:"structure"`

	// The log level the user wants for their channel.
	LogLevel LogLevel `locationName:"logLevel" type:"string" enum:"true"`

	Name *string `locationName:"name" type:"string"`

	PipelineDetails []PipelineDetail `locationName:"pipelineDetails" type:"list"`

	PipelinesRunningCount *int64 `locationName:"pipelinesRunningCount" type:"integer"`

	RoleArn *string `locationName:"roleArn" type:"string"`

	State ChannelState `locationName:"state" type:"string" enum:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s DeleteChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ChannelClass) > 0 {
		v := s.ChannelClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelClass", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Destinations != nil {
		v := s.Destinations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "destinations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.EgressEndpoints != nil {
		v := s.EgressEndpoints

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "egressEndpoints", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.EncoderSettings != nil {
		v := s.EncoderSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encoderSettings", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InputAttachments != nil {
		v := s.InputAttachments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "inputAttachments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.InputSpecification != nil {
		v := s.InputSpecification

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "inputSpecification", v, metadata)
	}
	if len(s.LogLevel) > 0 {
		v := s.LogLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "logLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PipelineDetails != nil {
		v := s.PipelineDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "pipelineDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.PipelinesRunningCount != nil {
		v := *s.PipelinesRunningCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "pipelinesRunningCount", protocol.Int64Value(v), metadata)
	}
	if s.RoleArn != nil {
		v := *s.RoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "roleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "state", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opDeleteChannel = "DeleteChannel"

// DeleteChannelRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Starts deletion of channel. The associated outputs are also deleted.
//
//    // Example sending a request using DeleteChannelRequest.
//    req := client.DeleteChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DeleteChannel
func (c *Client) DeleteChannelRequest(input *DeleteChannelInput) DeleteChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/prod/channels/{channelId}",
	}

	if input == nil {
		input = &DeleteChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteChannelOutput{})
	return DeleteChannelRequest{Request: req, Input: input, Copy: c.DeleteChannelRequest}
}

// DeleteChannelRequest is the request type for the
// DeleteChannel API operation.
type DeleteChannelRequest struct {
	*aws.Request
	Input *DeleteChannelInput
	Copy  func(*DeleteChannelInput) DeleteChannelRequest
}

// Send marshals and sends the DeleteChannel API request.
func (r DeleteChannelRequest) Send(ctx context.Context) (*DeleteChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteChannelResponse{
		DeleteChannelOutput: r.Request.Data.(*DeleteChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteChannelResponse is the response type for the
// DeleteChannel API operation.
type DeleteChannelResponse struct {
	*DeleteChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteChannel request.
func (r *DeleteChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
