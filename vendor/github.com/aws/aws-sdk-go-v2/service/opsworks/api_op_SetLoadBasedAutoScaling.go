// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type SetLoadBasedAutoScalingInput struct {
	_ struct{} `type:"structure"`

	// An AutoScalingThresholds object with the downscaling threshold configuration.
	// If the load falls below these thresholds for a specified amount of time,
	// AWS OpsWorks Stacks stops a specified number of instances.
	DownScaling *AutoScalingThresholds `type:"structure"`

	// Enables load-based auto scaling for the layer.
	Enable *bool `type:"boolean"`

	// The layer ID.
	//
	// LayerId is a required field
	LayerId *string `type:"string" required:"true"`

	// An AutoScalingThresholds object with the upscaling threshold configuration.
	// If the load exceeds these thresholds for a specified amount of time, AWS
	// OpsWorks Stacks starts a specified number of instances.
	UpScaling *AutoScalingThresholds `type:"structure"`
}

// String returns the string representation
func (s SetLoadBasedAutoScalingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetLoadBasedAutoScalingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetLoadBasedAutoScalingInput"}

	if s.LayerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerId"))
	}
	if s.DownScaling != nil {
		if err := s.DownScaling.Validate(); err != nil {
			invalidParams.AddNested("DownScaling", err.(aws.ErrInvalidParams))
		}
	}
	if s.UpScaling != nil {
		if err := s.UpScaling.Validate(); err != nil {
			invalidParams.AddNested("UpScaling", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetLoadBasedAutoScalingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetLoadBasedAutoScalingOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetLoadBasedAutoScaling = "SetLoadBasedAutoScaling"

// SetLoadBasedAutoScalingRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Specify the load-based auto scaling configuration for a specified layer.
// For more information, see Managing Load with Time-based and Load-based Instances
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-autoscaling.html).
//
// To use load-based auto scaling, you must create a set of load-based auto
// scaling instances. Load-based auto scaling operates only on the instances
// from that set, so you must ensure that you have created enough instances
// to handle the maximum anticipated load.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using SetLoadBasedAutoScalingRequest.
//    req := client.SetLoadBasedAutoScalingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/SetLoadBasedAutoScaling
func (c *Client) SetLoadBasedAutoScalingRequest(input *SetLoadBasedAutoScalingInput) SetLoadBasedAutoScalingRequest {
	op := &aws.Operation{
		Name:       opSetLoadBasedAutoScaling,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetLoadBasedAutoScalingInput{}
	}

	req := c.newRequest(op, input, &SetLoadBasedAutoScalingOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetLoadBasedAutoScalingRequest{Request: req, Input: input, Copy: c.SetLoadBasedAutoScalingRequest}
}

// SetLoadBasedAutoScalingRequest is the request type for the
// SetLoadBasedAutoScaling API operation.
type SetLoadBasedAutoScalingRequest struct {
	*aws.Request
	Input *SetLoadBasedAutoScalingInput
	Copy  func(*SetLoadBasedAutoScalingInput) SetLoadBasedAutoScalingRequest
}

// Send marshals and sends the SetLoadBasedAutoScaling API request.
func (r SetLoadBasedAutoScalingRequest) Send(ctx context.Context) (*SetLoadBasedAutoScalingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetLoadBasedAutoScalingResponse{
		SetLoadBasedAutoScalingOutput: r.Request.Data.(*SetLoadBasedAutoScalingOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetLoadBasedAutoScalingResponse is the response type for the
// SetLoadBasedAutoScaling API operation.
type SetLoadBasedAutoScalingResponse struct {
	*SetLoadBasedAutoScalingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetLoadBasedAutoScaling request.
func (r *SetLoadBasedAutoScalingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
