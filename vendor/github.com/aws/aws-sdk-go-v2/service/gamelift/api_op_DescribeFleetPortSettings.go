// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeFleetPortSettingsInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet to retrieve port settings for.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFleetPortSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetPortSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetPortSettingsInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeFleetPortSettingsOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains port settings for the requested fleet ID.
	InboundPermissions []IpPermission `type:"list"`
}

// String returns the string representation
func (s DescribeFleetPortSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFleetPortSettings = "DescribeFleetPortSettings"

// DescribeFleetPortSettingsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the inbound connection permissions for a fleet. Connection permissions
// include a range of IP addresses and port settings that incoming traffic can
// use to access server processes in the fleet. To get a fleet's inbound connection
// permissions, specify a fleet ID. If successful, a collection of IpPermission
// objects is returned for the requested fleet ID. If the requested fleet has
// been deleted, the result set is empty.
//
// Learn more
//
//  Working with Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html).
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * Describe fleets: DescribeFleetAttributes DescribeFleetCapacity DescribeFleetPortSettings
//    DescribeFleetUtilization DescribeRuntimeConfiguration DescribeEC2InstanceLimits
//    DescribeFleetEvents
//
//    * Update fleets: UpdateFleetAttributes UpdateFleetCapacity UpdateFleetPortSettings
//    UpdateRuntimeConfiguration
//
//    * Manage fleet actions: StartFleetActions StopFleetActions
//
//    // Example sending a request using DescribeFleetPortSettingsRequest.
//    req := client.DescribeFleetPortSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeFleetPortSettings
func (c *Client) DescribeFleetPortSettingsRequest(input *DescribeFleetPortSettingsInput) DescribeFleetPortSettingsRequest {
	op := &aws.Operation{
		Name:       opDescribeFleetPortSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeFleetPortSettingsInput{}
	}

	req := c.newRequest(op, input, &DescribeFleetPortSettingsOutput{})
	return DescribeFleetPortSettingsRequest{Request: req, Input: input, Copy: c.DescribeFleetPortSettingsRequest}
}

// DescribeFleetPortSettingsRequest is the request type for the
// DescribeFleetPortSettings API operation.
type DescribeFleetPortSettingsRequest struct {
	*aws.Request
	Input *DescribeFleetPortSettingsInput
	Copy  func(*DescribeFleetPortSettingsInput) DescribeFleetPortSettingsRequest
}

// Send marshals and sends the DescribeFleetPortSettings API request.
func (r DescribeFleetPortSettingsRequest) Send(ctx context.Context) (*DescribeFleetPortSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFleetPortSettingsResponse{
		DescribeFleetPortSettingsOutput: r.Request.Data.(*DescribeFleetPortSettingsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFleetPortSettingsResponse is the response type for the
// DescribeFleetPortSettings API operation.
type DescribeFleetPortSettingsResponse struct {
	*DescribeFleetPortSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFleetPortSettings request.
func (r *DescribeFleetPortSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
