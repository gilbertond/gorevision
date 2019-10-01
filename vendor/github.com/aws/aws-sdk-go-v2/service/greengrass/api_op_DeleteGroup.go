// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteGroupInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGroupInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteGroup = "DeleteGroup"

// DeleteGroupRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Deletes a group.
//
//    // Example sending a request using DeleteGroupRequest.
//    req := client.DeleteGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteGroup
func (c *Client) DeleteGroupRequest(input *DeleteGroupInput) DeleteGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/groups/{GroupId}",
	}

	if input == nil {
		input = &DeleteGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteGroupOutput{})
	return DeleteGroupRequest{Request: req, Input: input, Copy: c.DeleteGroupRequest}
}

// DeleteGroupRequest is the request type for the
// DeleteGroup API operation.
type DeleteGroupRequest struct {
	*aws.Request
	Input *DeleteGroupInput
	Copy  func(*DeleteGroupInput) DeleteGroupRequest
}

// Send marshals and sends the DeleteGroup API request.
func (r DeleteGroupRequest) Send(ctx context.Context) (*DeleteGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupResponse{
		DeleteGroupOutput: r.Request.Data.(*DeleteGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupResponse is the response type for the
// DeleteGroup API operation.
type DeleteGroupResponse struct {
	*DeleteGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroup request.
func (r *DeleteGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
