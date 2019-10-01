// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteBusinessReportScheduleInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the business report schedule.
	//
	// ScheduleArn is a required field
	ScheduleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBusinessReportScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBusinessReportScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBusinessReportScheduleInput"}

	if s.ScheduleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteBusinessReportScheduleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBusinessReportScheduleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteBusinessReportSchedule = "DeleteBusinessReportSchedule"

// DeleteBusinessReportScheduleRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes the recurring report delivery schedule with the specified schedule
// ARN.
//
//    // Example sending a request using DeleteBusinessReportScheduleRequest.
//    req := client.DeleteBusinessReportScheduleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteBusinessReportSchedule
func (c *Client) DeleteBusinessReportScheduleRequest(input *DeleteBusinessReportScheduleInput) DeleteBusinessReportScheduleRequest {
	op := &aws.Operation{
		Name:       opDeleteBusinessReportSchedule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBusinessReportScheduleInput{}
	}

	req := c.newRequest(op, input, &DeleteBusinessReportScheduleOutput{})
	return DeleteBusinessReportScheduleRequest{Request: req, Input: input, Copy: c.DeleteBusinessReportScheduleRequest}
}

// DeleteBusinessReportScheduleRequest is the request type for the
// DeleteBusinessReportSchedule API operation.
type DeleteBusinessReportScheduleRequest struct {
	*aws.Request
	Input *DeleteBusinessReportScheduleInput
	Copy  func(*DeleteBusinessReportScheduleInput) DeleteBusinessReportScheduleRequest
}

// Send marshals and sends the DeleteBusinessReportSchedule API request.
func (r DeleteBusinessReportScheduleRequest) Send(ctx context.Context) (*DeleteBusinessReportScheduleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBusinessReportScheduleResponse{
		DeleteBusinessReportScheduleOutput: r.Request.Data.(*DeleteBusinessReportScheduleOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBusinessReportScheduleResponse is the response type for the
// DeleteBusinessReportSchedule API operation.
type DeleteBusinessReportScheduleResponse struct {
	*DeleteBusinessReportScheduleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBusinessReportSchedule request.
func (r *DeleteBusinessReportScheduleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
