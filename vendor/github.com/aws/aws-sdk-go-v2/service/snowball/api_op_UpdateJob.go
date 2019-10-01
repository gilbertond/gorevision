// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateJobInput struct {
	_ struct{} `type:"structure"`

	// The ID of the updated Address object.
	AddressId *string `min:"40" type:"string"`

	// The updated description of this job's JobMetadata object.
	Description *string `min:"1" type:"string"`

	// The updated ID for the forwarding address for a job. This field is not supported
	// in most regions.
	ForwardingAddressId *string `min:"40" type:"string"`

	// The job ID of the job that you want to update, for example JID123e4567-e89b-12d3-a456-426655440000.
	//
	// JobId is a required field
	JobId *string `min:"39" type:"string" required:"true"`

	// The new or updated Notification object.
	Notification *Notification `type:"structure"`

	// The updated JobResource object, or the updated JobResource object.
	Resources *JobResource `type:"structure"`

	// The new role Amazon Resource Name (ARN) that you want to associate with this
	// job. To create a role ARN, use the CreateRole (http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)AWS
	// Identity and Access Management (IAM) API action.
	RoleARN *string `type:"string"`

	// The updated shipping option value of this job's ShippingDetails object.
	ShippingOption ShippingOption `type:"string" enum:"true"`

	// The updated SnowballCapacityPreference of this job's JobMetadata object.
	// The 50 TB Snowballs are only available in the US regions.
	SnowballCapacityPreference SnowballCapacity `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobInput"}
	if s.AddressId != nil && len(*s.AddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("AddressId", 40))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.ForwardingAddressId != nil && len(*s.ForwardingAddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("ForwardingAddressId", 40))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 39))
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			invalidParams.AddNested("Resources", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateJob = "UpdateJob"

// UpdateJobRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// While a job's JobState value is New, you can update some of the information
// associated with a job. Once the job changes to a different job state, usually
// within 60 minutes of the job being created, this action is no longer available.
//
//    // Example sending a request using UpdateJobRequest.
//    req := client.UpdateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/UpdateJob
func (c *Client) UpdateJobRequest(input *UpdateJobInput) UpdateJobRequest {
	op := &aws.Operation{
		Name:       opUpdateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateJobInput{}
	}

	req := c.newRequest(op, input, &UpdateJobOutput{})
	return UpdateJobRequest{Request: req, Input: input, Copy: c.UpdateJobRequest}
}

// UpdateJobRequest is the request type for the
// UpdateJob API operation.
type UpdateJobRequest struct {
	*aws.Request
	Input *UpdateJobInput
	Copy  func(*UpdateJobInput) UpdateJobRequest
}

// Send marshals and sends the UpdateJob API request.
func (r UpdateJobRequest) Send(ctx context.Context) (*UpdateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJobResponse{
		UpdateJobOutput: r.Request.Data.(*UpdateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJobResponse is the response type for the
// UpdateJob API operation.
type UpdateJobResponse struct {
	*UpdateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJob request.
func (r *UpdateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
