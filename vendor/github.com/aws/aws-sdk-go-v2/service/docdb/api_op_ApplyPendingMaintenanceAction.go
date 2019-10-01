// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to ApplyPendingMaintenanceAction.
type ApplyPendingMaintenanceActionInput struct {
	_ struct{} `type:"structure"`

	// The pending maintenance action to apply to this resource.
	//
	// Valid values: system-update, db-upgrade
	//
	// ApplyAction is a required field
	ApplyAction *string `type:"string" required:"true"`

	// A value that specifies the type of opt-in request or undoes an opt-in request.
	// An opt-in request of type immediate can't be undone.
	//
	// Valid values:
	//
	//    * immediate - Apply the maintenance action immediately.
	//
	//    * next-maintenance - Apply the maintenance action during the next maintenance
	//    window for the resource.
	//
	//    * undo-opt-in - Cancel any existing next-maintenance opt-in requests.
	//
	// OptInType is a required field
	OptInType *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the resource that the pending maintenance
	// action applies to.
	//
	// ResourceIdentifier is a required field
	ResourceIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplyPendingMaintenanceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplyPendingMaintenanceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplyPendingMaintenanceActionInput"}

	if s.ApplyAction == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyAction"))
	}

	if s.OptInType == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptInType"))
	}

	if s.ResourceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ApplyPendingMaintenanceActionOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of ApplyPendingMaintenanceAction.
	ResourcePendingMaintenanceActions *ResourcePendingMaintenanceActions `type:"structure"`
}

// String returns the string representation
func (s ApplyPendingMaintenanceActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opApplyPendingMaintenanceAction = "ApplyPendingMaintenanceAction"

// ApplyPendingMaintenanceActionRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Applies a pending maintenance action to a resource (for example, to a DB
// instance).
//
//    // Example sending a request using ApplyPendingMaintenanceActionRequest.
//    req := client.ApplyPendingMaintenanceActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/ApplyPendingMaintenanceAction
func (c *Client) ApplyPendingMaintenanceActionRequest(input *ApplyPendingMaintenanceActionInput) ApplyPendingMaintenanceActionRequest {
	op := &aws.Operation{
		Name:       opApplyPendingMaintenanceAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ApplyPendingMaintenanceActionInput{}
	}

	req := c.newRequest(op, input, &ApplyPendingMaintenanceActionOutput{})
	return ApplyPendingMaintenanceActionRequest{Request: req, Input: input, Copy: c.ApplyPendingMaintenanceActionRequest}
}

// ApplyPendingMaintenanceActionRequest is the request type for the
// ApplyPendingMaintenanceAction API operation.
type ApplyPendingMaintenanceActionRequest struct {
	*aws.Request
	Input *ApplyPendingMaintenanceActionInput
	Copy  func(*ApplyPendingMaintenanceActionInput) ApplyPendingMaintenanceActionRequest
}

// Send marshals and sends the ApplyPendingMaintenanceAction API request.
func (r ApplyPendingMaintenanceActionRequest) Send(ctx context.Context) (*ApplyPendingMaintenanceActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ApplyPendingMaintenanceActionResponse{
		ApplyPendingMaintenanceActionOutput: r.Request.Data.(*ApplyPendingMaintenanceActionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ApplyPendingMaintenanceActionResponse is the response type for the
// ApplyPendingMaintenanceAction API operation.
type ApplyPendingMaintenanceActionResponse struct {
	*ApplyPendingMaintenanceActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ApplyPendingMaintenanceAction request.
func (r *ApplyPendingMaintenanceActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
