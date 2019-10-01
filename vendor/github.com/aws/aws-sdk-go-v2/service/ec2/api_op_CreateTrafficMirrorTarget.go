// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTrafficMirrorTargetInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The description of the Traffic Mirror target.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `type:"string"`

	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated
	// with the target.
	NetworkLoadBalancerArn *string `type:"string"`

	// The tags to assign to the Traffic Mirror target.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateTrafficMirrorTargetInput) String() string {
	return awsutil.Prettify(s)
}

type CreateTrafficMirrorTargetOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the Traffic Mirror target.
	TrafficMirrorTarget *TrafficMirrorTarget `locationName:"trafficMirrorTarget" type:"structure"`
}

// String returns the string representation
func (s CreateTrafficMirrorTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTrafficMirrorTarget = "CreateTrafficMirrorTarget"

// CreateTrafficMirrorTargetRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a target for your Traffic Mirror session.
//
// A Traffic Mirror target is the destination for mirrored traffic. The Traffic
// Mirror source and the Traffic Mirror target (monitoring appliances) can be
// in the same VPC, or in different VPCs connected via VPC peering or a transit
// gateway.
//
// A Traffic Mirror target can be a network interface, or a Network Load Balancer.
//
// To use the target in a Traffic Mirror session, use CreateTrafficMirrorSession.
//
//    // Example sending a request using CreateTrafficMirrorTargetRequest.
//    req := client.CreateTrafficMirrorTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTrafficMirrorTarget
func (c *Client) CreateTrafficMirrorTargetRequest(input *CreateTrafficMirrorTargetInput) CreateTrafficMirrorTargetRequest {
	op := &aws.Operation{
		Name:       opCreateTrafficMirrorTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrafficMirrorTargetInput{}
	}

	req := c.newRequest(op, input, &CreateTrafficMirrorTargetOutput{})
	return CreateTrafficMirrorTargetRequest{Request: req, Input: input, Copy: c.CreateTrafficMirrorTargetRequest}
}

// CreateTrafficMirrorTargetRequest is the request type for the
// CreateTrafficMirrorTarget API operation.
type CreateTrafficMirrorTargetRequest struct {
	*aws.Request
	Input *CreateTrafficMirrorTargetInput
	Copy  func(*CreateTrafficMirrorTargetInput) CreateTrafficMirrorTargetRequest
}

// Send marshals and sends the CreateTrafficMirrorTarget API request.
func (r CreateTrafficMirrorTargetRequest) Send(ctx context.Context) (*CreateTrafficMirrorTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrafficMirrorTargetResponse{
		CreateTrafficMirrorTargetOutput: r.Request.Data.(*CreateTrafficMirrorTargetOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrafficMirrorTargetResponse is the response type for the
// CreateTrafficMirrorTarget API operation.
type CreateTrafficMirrorTargetResponse struct {
	*CreateTrafficMirrorTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrafficMirrorTarget request.
func (r *CreateTrafficMirrorTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
