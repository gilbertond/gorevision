// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type RemoveLayerVersionPermissionInput struct {
	_ struct{} `type:"structure"`

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// LayerName is a required field
	LayerName *string `location:"uri" locationName:"LayerName" min:"1" type:"string" required:"true"`

	// Only update the policy if the revision ID matches the ID specified. Use this
	// option to avoid modifying a policy that has changed since you last read it.
	RevisionId *string `location:"querystring" locationName:"RevisionId" type:"string"`

	// The identifier that was specified when the statement was added.
	//
	// StatementId is a required field
	StatementId *string `location:"uri" locationName:"StatementId" min:"1" type:"string" required:"true"`

	// The version number.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" type:"long" required:"true"`
}

// String returns the string representation
func (s RemoveLayerVersionPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveLayerVersionPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveLayerVersionPermissionInput"}

	if s.LayerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerName"))
	}
	if s.LayerName != nil && len(*s.LayerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LayerName", 1))
	}

	if s.StatementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatementId"))
	}
	if s.StatementId != nil && len(*s.StatementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatementId", 1))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveLayerVersionPermissionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LayerName != nil {
		v := *s.LayerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LayerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StatementId != nil {
		v := *s.StatementId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "StatementId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionNumber", protocol.Int64Value(v), metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RemoveLayerVersionPermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveLayerVersionPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveLayerVersionPermissionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRemoveLayerVersionPermission = "RemoveLayerVersionPermission"

// RemoveLayerVersionPermissionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Removes a statement from the permissions policy for a version of an AWS Lambda
// layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// For more information, see AddLayerVersionPermission.
//
//    // Example sending a request using RemoveLayerVersionPermissionRequest.
//    req := client.RemoveLayerVersionPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/RemoveLayerVersionPermission
func (c *Client) RemoveLayerVersionPermissionRequest(input *RemoveLayerVersionPermissionInput) RemoveLayerVersionPermissionRequest {
	op := &aws.Operation{
		Name:       opRemoveLayerVersionPermission,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions/{VersionNumber}/policy/{StatementId}",
	}

	if input == nil {
		input = &RemoveLayerVersionPermissionInput{}
	}

	req := c.newRequest(op, input, &RemoveLayerVersionPermissionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveLayerVersionPermissionRequest{Request: req, Input: input, Copy: c.RemoveLayerVersionPermissionRequest}
}

// RemoveLayerVersionPermissionRequest is the request type for the
// RemoveLayerVersionPermission API operation.
type RemoveLayerVersionPermissionRequest struct {
	*aws.Request
	Input *RemoveLayerVersionPermissionInput
	Copy  func(*RemoveLayerVersionPermissionInput) RemoveLayerVersionPermissionRequest
}

// Send marshals and sends the RemoveLayerVersionPermission API request.
func (r RemoveLayerVersionPermissionRequest) Send(ctx context.Context) (*RemoveLayerVersionPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveLayerVersionPermissionResponse{
		RemoveLayerVersionPermissionOutput: r.Request.Data.(*RemoveLayerVersionPermissionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveLayerVersionPermissionResponse is the response type for the
// RemoveLayerVersionPermission API operation.
type RemoveLayerVersionPermissionResponse struct {
	*RemoveLayerVersionPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveLayerVersionPermission request.
func (r *RemoveLayerVersionPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
