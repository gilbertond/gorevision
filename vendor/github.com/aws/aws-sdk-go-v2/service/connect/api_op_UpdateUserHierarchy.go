// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UpdateUserHierarchyInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the hierarchy group to assign to the user.
	HierarchyGroupId *string `type:"string"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The identifier of the user account to assign the hierarchy group to.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"UserId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserHierarchyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserHierarchyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserHierarchyInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserHierarchyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.HierarchyGroupId != nil {
		v := *s.HierarchyGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HierarchyGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateUserHierarchyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateUserHierarchyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserHierarchyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateUserHierarchy = "UpdateUserHierarchy"

// UpdateUserHierarchyRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Assigns the specified hierarchy group to the user.
//
//    // Example sending a request using UpdateUserHierarchyRequest.
//    req := client.UpdateUserHierarchyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/UpdateUserHierarchy
func (c *Client) UpdateUserHierarchyRequest(input *UpdateUserHierarchyInput) UpdateUserHierarchyRequest {
	op := &aws.Operation{
		Name:       opUpdateUserHierarchy,
		HTTPMethod: "POST",
		HTTPPath:   "/users/{InstanceId}/{UserId}/hierarchy",
	}

	if input == nil {
		input = &UpdateUserHierarchyInput{}
	}

	req := c.newRequest(op, input, &UpdateUserHierarchyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateUserHierarchyRequest{Request: req, Input: input, Copy: c.UpdateUserHierarchyRequest}
}

// UpdateUserHierarchyRequest is the request type for the
// UpdateUserHierarchy API operation.
type UpdateUserHierarchyRequest struct {
	*aws.Request
	Input *UpdateUserHierarchyInput
	Copy  func(*UpdateUserHierarchyInput) UpdateUserHierarchyRequest
}

// Send marshals and sends the UpdateUserHierarchy API request.
func (r UpdateUserHierarchyRequest) Send(ctx context.Context) (*UpdateUserHierarchyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserHierarchyResponse{
		UpdateUserHierarchyOutput: r.Request.Data.(*UpdateUserHierarchyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserHierarchyResponse is the response type for the
// UpdateUserHierarchy API operation.
type UpdateUserHierarchyResponse struct {
	*UpdateUserHierarchyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserHierarchy request.
func (r *UpdateUserHierarchyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
