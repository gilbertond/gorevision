// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMissionProfileInput struct {
	_ struct{} `type:"structure"`

	ContactPostPassDurationSeconds *int64 `locationName:"contactPostPassDurationSeconds" min:"1" type:"integer"`

	ContactPrePassDurationSeconds *int64 `locationName:"contactPrePassDurationSeconds" min:"1" type:"integer"`

	// DataflowEdges is a required field
	DataflowEdges [][]string `locationName:"dataflowEdges" type:"list" required:"true"`

	// MinimumViableContactDurationSeconds is a required field
	MinimumViableContactDurationSeconds *int64 `locationName:"minimumViableContactDurationSeconds" min:"1" type:"integer" required:"true"`

	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`

	// TrackingConfigArn is a required field
	TrackingConfigArn *string `locationName:"trackingConfigArn" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMissionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMissionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMissionProfileInput"}
	if s.ContactPostPassDurationSeconds != nil && *s.ContactPostPassDurationSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ContactPostPassDurationSeconds", 1))
	}
	if s.ContactPrePassDurationSeconds != nil && *s.ContactPrePassDurationSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ContactPrePassDurationSeconds", 1))
	}

	if s.DataflowEdges == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataflowEdges"))
	}

	if s.MinimumViableContactDurationSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("MinimumViableContactDurationSeconds"))
	}
	if s.MinimumViableContactDurationSeconds != nil && *s.MinimumViableContactDurationSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MinimumViableContactDurationSeconds", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.TrackingConfigArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrackingConfigArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMissionProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ContactPostPassDurationSeconds != nil {
		v := *s.ContactPostPassDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactPostPassDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.ContactPrePassDurationSeconds != nil {
		v := *s.ContactPrePassDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactPrePassDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.DataflowEdges != nil {
		v := s.DataflowEdges

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataflowEdges", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls1 := ls0.List()
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ls0.End()

	}
	if s.MinimumViableContactDurationSeconds != nil {
		v := *s.MinimumViableContactDurationSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "minimumViableContactDurationSeconds", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.TrackingConfigArn != nil {
		v := *s.TrackingConfigArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "trackingConfigArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateMissionProfileOutput struct {
	_ struct{} `type:"structure"`

	MissionProfileId *string `locationName:"missionProfileId" type:"string"`
}

// String returns the string representation
func (s CreateMissionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMissionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MissionProfileId != nil {
		v := *s.MissionProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "missionProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateMissionProfile = "CreateMissionProfile"

// CreateMissionProfileRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Creates a mission profile.
//
// dataflowEdges is a list of lists of strings. Each lower level list of strings
// has two elements: a from ARN and a to ARN.
//
//    // Example sending a request using CreateMissionProfileRequest.
//    req := client.CreateMissionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/CreateMissionProfile
func (c *Client) CreateMissionProfileRequest(input *CreateMissionProfileInput) CreateMissionProfileRequest {
	op := &aws.Operation{
		Name:       opCreateMissionProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/missionprofile",
	}

	if input == nil {
		input = &CreateMissionProfileInput{}
	}

	req := c.newRequest(op, input, &CreateMissionProfileOutput{})
	return CreateMissionProfileRequest{Request: req, Input: input, Copy: c.CreateMissionProfileRequest}
}

// CreateMissionProfileRequest is the request type for the
// CreateMissionProfile API operation.
type CreateMissionProfileRequest struct {
	*aws.Request
	Input *CreateMissionProfileInput
	Copy  func(*CreateMissionProfileInput) CreateMissionProfileRequest
}

// Send marshals and sends the CreateMissionProfile API request.
func (r CreateMissionProfileRequest) Send(ctx context.Context) (*CreateMissionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMissionProfileResponse{
		CreateMissionProfileOutput: r.Request.Data.(*CreateMissionProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMissionProfileResponse is the response type for the
// CreateMissionProfile API operation.
type CreateMissionProfileResponse struct {
	*CreateMissionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMissionProfile request.
func (r *CreateMissionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
