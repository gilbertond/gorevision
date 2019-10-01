// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeMaintenanceWindowsForTargetInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The type of resource you want to retrieve information about. For example,
	// "INSTANCE".
	//
	// ResourceType is a required field
	ResourceType MaintenanceWindowResourceType `type:"string" required:"true" enum:"true"`

	// The instance ID or key/value pair to retrieve information about.
	//
	// Targets is a required field
	Targets []Target `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowsForTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMaintenanceWindowsForTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMaintenanceWindowsForTargetInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeMaintenanceWindowsForTargetOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. (You use this token in the
	// next call.)
	NextToken *string `type:"string"`

	// Information about the maintenance window targets and tasks an instance is
	// associated with.
	WindowIdentities []MaintenanceWindowIdentityForTarget `type:"list"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowsForTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMaintenanceWindowsForTarget = "DescribeMaintenanceWindowsForTarget"

// DescribeMaintenanceWindowsForTargetRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves information about the maintenance window targets or tasks that
// an instance is associated with.
//
//    // Example sending a request using DescribeMaintenanceWindowsForTargetRequest.
//    req := client.DescribeMaintenanceWindowsForTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeMaintenanceWindowsForTarget
func (c *Client) DescribeMaintenanceWindowsForTargetRequest(input *DescribeMaintenanceWindowsForTargetInput) DescribeMaintenanceWindowsForTargetRequest {
	op := &aws.Operation{
		Name:       opDescribeMaintenanceWindowsForTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMaintenanceWindowsForTargetInput{}
	}

	req := c.newRequest(op, input, &DescribeMaintenanceWindowsForTargetOutput{})
	return DescribeMaintenanceWindowsForTargetRequest{Request: req, Input: input, Copy: c.DescribeMaintenanceWindowsForTargetRequest}
}

// DescribeMaintenanceWindowsForTargetRequest is the request type for the
// DescribeMaintenanceWindowsForTarget API operation.
type DescribeMaintenanceWindowsForTargetRequest struct {
	*aws.Request
	Input *DescribeMaintenanceWindowsForTargetInput
	Copy  func(*DescribeMaintenanceWindowsForTargetInput) DescribeMaintenanceWindowsForTargetRequest
}

// Send marshals and sends the DescribeMaintenanceWindowsForTarget API request.
func (r DescribeMaintenanceWindowsForTargetRequest) Send(ctx context.Context) (*DescribeMaintenanceWindowsForTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMaintenanceWindowsForTargetResponse{
		DescribeMaintenanceWindowsForTargetOutput: r.Request.Data.(*DescribeMaintenanceWindowsForTargetOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMaintenanceWindowsForTargetResponse is the response type for the
// DescribeMaintenanceWindowsForTarget API operation.
type DescribeMaintenanceWindowsForTargetResponse struct {
	*DescribeMaintenanceWindowsForTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMaintenanceWindowsForTarget request.
func (r *DescribeMaintenanceWindowsForTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
