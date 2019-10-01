// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateExportTaskInput struct {
	_ struct{} `type:"structure"`

	// The name of S3 bucket for the exported log data. The bucket must be in the
	// same AWS region.
	//
	// Destination is a required field
	Destination *string `locationName:"destination" min:"1" type:"string" required:"true"`

	// The prefix used as the start of the key for every object exported. If you
	// don't specify a value, the default is exportedlogs.
	DestinationPrefix *string `locationName:"destinationPrefix" type:"string"`

	// The start time of the range for the request, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC. Events with a timestamp earlier than this
	// time are not exported.
	//
	// From is a required field
	From *int64 `locationName:"from" type:"long" required:"true"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// Export only log streams that match the provided prefix. If you don't specify
	// a value, no prefix filter is applied.
	LogStreamNamePrefix *string `locationName:"logStreamNamePrefix" min:"1" type:"string"`

	// The name of the export task.
	TaskName *string `locationName:"taskName" min:"1" type:"string"`

	// The end time of the range for the request, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC. Events with a timestamp later than this time
	// are not exported.
	//
	// To is a required field
	To *int64 `locationName:"to" type:"long" required:"true"`
}

// String returns the string representation
func (s CreateExportTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateExportTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateExportTaskInput"}

	if s.Destination == nil {
		invalidParams.Add(aws.NewErrParamRequired("Destination"))
	}
	if s.Destination != nil && len(*s.Destination) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Destination", 1))
	}

	if s.From == nil {
		invalidParams.Add(aws.NewErrParamRequired("From"))
	}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}
	if s.LogStreamNamePrefix != nil && len(*s.LogStreamNamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamNamePrefix", 1))
	}
	if s.TaskName != nil && len(*s.TaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskName", 1))
	}

	if s.To == nil {
		invalidParams.Add(aws.NewErrParamRequired("To"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateExportTaskOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the export task.
	TaskId *string `locationName:"taskId" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateExportTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateExportTask = "CreateExportTask"

// CreateExportTaskRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Creates an export task, which allows you to efficiently export data from
// a log group to an Amazon S3 bucket.
//
// This is an asynchronous call. If all the required information is provided,
// this operation initiates an export task and responds with the ID of the task.
// After the task has started, you can use DescribeExportTasks to get the status
// of the export task. Each account can only have one active (RUNNING or PENDING)
// export task at a time. To cancel an export task, use CancelExportTask.
//
// You can export logs from multiple log groups or multiple time ranges to the
// same S3 bucket. To separate out log data for each export task, you can specify
// a prefix to be used as the Amazon S3 key prefix for all exported objects.
//
// Exporting to S3 buckets that are encrypted with AES-256 is supported. Exporting
// to S3 buckets encrypted with SSE-KMS is not supported.
//
//    // Example sending a request using CreateExportTaskRequest.
//    req := client.CreateExportTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/CreateExportTask
func (c *Client) CreateExportTaskRequest(input *CreateExportTaskInput) CreateExportTaskRequest {
	op := &aws.Operation{
		Name:       opCreateExportTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateExportTaskInput{}
	}

	req := c.newRequest(op, input, &CreateExportTaskOutput{})
	return CreateExportTaskRequest{Request: req, Input: input, Copy: c.CreateExportTaskRequest}
}

// CreateExportTaskRequest is the request type for the
// CreateExportTask API operation.
type CreateExportTaskRequest struct {
	*aws.Request
	Input *CreateExportTaskInput
	Copy  func(*CreateExportTaskInput) CreateExportTaskRequest
}

// Send marshals and sends the CreateExportTask API request.
func (r CreateExportTaskRequest) Send(ctx context.Context) (*CreateExportTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateExportTaskResponse{
		CreateExportTaskOutput: r.Request.Data.(*CreateExportTaskOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateExportTaskResponse is the response type for the
// CreateExportTask API operation.
type CreateExportTaskResponse struct {
	*CreateExportTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateExportTask request.
func (r *CreateExportTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
