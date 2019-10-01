// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteCampaignInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// CampaignId is a required field
	CampaignId *string `location:"uri" locationName:"campaign-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCampaignInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.CampaignId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCampaignInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CampaignId != nil {
		v := *s.CampaignId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "campaign-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteCampaignOutput struct {
	_ struct{} `type:"structure" payload:"CampaignResponse"`

	// Provides information about the status, configuration, and other settings
	// for a campaign.
	//
	// CampaignResponse is a required field
	CampaignResponse *CampaignResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteCampaignOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCampaignOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CampaignResponse != nil {
		v := s.CampaignResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CampaignResponse", v, metadata)
	}
	return nil
}

const opDeleteCampaign = "DeleteCampaign"

// DeleteCampaignRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes a campaign from an application.
//
//    // Example sending a request using DeleteCampaignRequest.
//    req := client.DeleteCampaignRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteCampaign
func (c *Client) DeleteCampaignRequest(input *DeleteCampaignInput) DeleteCampaignRequest {
	op := &aws.Operation{
		Name:       opDeleteCampaign,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/campaigns/{campaign-id}",
	}

	if input == nil {
		input = &DeleteCampaignInput{}
	}

	req := c.newRequest(op, input, &DeleteCampaignOutput{})
	return DeleteCampaignRequest{Request: req, Input: input, Copy: c.DeleteCampaignRequest}
}

// DeleteCampaignRequest is the request type for the
// DeleteCampaign API operation.
type DeleteCampaignRequest struct {
	*aws.Request
	Input *DeleteCampaignInput
	Copy  func(*DeleteCampaignInput) DeleteCampaignRequest
}

// Send marshals and sends the DeleteCampaign API request.
func (r DeleteCampaignRequest) Send(ctx context.Context) (*DeleteCampaignResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCampaignResponse{
		DeleteCampaignOutput: r.Request.Data.(*DeleteCampaignOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCampaignResponse is the response type for the
// DeleteCampaign API operation.
type DeleteCampaignResponse struct {
	*DeleteCampaignOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCampaign request.
func (r *DeleteCampaignResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
