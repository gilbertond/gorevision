// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdatePullRequestStatusInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the pull request. To get this ID, use ListPullRequests.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The status of the pull request. The only valid operations are to update the
	// status from OPEN to OPEN, OPEN to CLOSED or from from CLOSED to CLOSED.
	//
	// PullRequestStatus is a required field
	PullRequestStatus PullRequestStatusEnum `locationName:"pullRequestStatus" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdatePullRequestStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePullRequestStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePullRequestStatusInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}
	if len(s.PullRequestStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdatePullRequestStatusOutput struct {
	_ struct{} `type:"structure"`

	// Information about the pull request.
	//
	// PullRequest is a required field
	PullRequest *PullRequest `locationName:"pullRequest" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePullRequestStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdatePullRequestStatus = "UpdatePullRequestStatus"

// UpdatePullRequestStatusRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Updates the status of a pull request.
//
//    // Example sending a request using UpdatePullRequestStatusRequest.
//    req := client.UpdatePullRequestStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdatePullRequestStatus
func (c *Client) UpdatePullRequestStatusRequest(input *UpdatePullRequestStatusInput) UpdatePullRequestStatusRequest {
	op := &aws.Operation{
		Name:       opUpdatePullRequestStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePullRequestStatusInput{}
	}

	req := c.newRequest(op, input, &UpdatePullRequestStatusOutput{})
	return UpdatePullRequestStatusRequest{Request: req, Input: input, Copy: c.UpdatePullRequestStatusRequest}
}

// UpdatePullRequestStatusRequest is the request type for the
// UpdatePullRequestStatus API operation.
type UpdatePullRequestStatusRequest struct {
	*aws.Request
	Input *UpdatePullRequestStatusInput
	Copy  func(*UpdatePullRequestStatusInput) UpdatePullRequestStatusRequest
}

// Send marshals and sends the UpdatePullRequestStatus API request.
func (r UpdatePullRequestStatusRequest) Send(ctx context.Context) (*UpdatePullRequestStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePullRequestStatusResponse{
		UpdatePullRequestStatusOutput: r.Request.Data.(*UpdatePullRequestStatusOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePullRequestStatusResponse is the response type for the
// UpdatePullRequestStatus API operation.
type UpdatePullRequestStatusResponse struct {
	*UpdatePullRequestStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePullRequestStatus request.
func (r *UpdatePullRequestStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
