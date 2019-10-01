// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RunTaskInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster on which
	// to run your task. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the specified task to place on your cluster.
	// You can specify up to 10 tasks per call.
	Count *int64 `locationName:"count" type:"integer"`

	// Specifies whether to enable Amazon ECS managed tags for the task. For more
	// information, see Tagging Your Amazon ECS Resources (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// in the Amazon Elastic Container Service Developer Guide.
	EnableECSManagedTags *bool `locationName:"enableECSManagedTags" type:"boolean"`

	// The name of the task group to associate with the task. The default value
	// is the family name of the task definition (for example, family:my-family-name).
	Group *string `locationName:"group" type:"string"`

	// The launch type on which to run your task. For more information, see Amazon
	// ECS Launch Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LaunchType LaunchType `locationName:"launchType" type:"string" enum:"true"`

	// The network configuration for the task. This parameter is required for task
	// definitions that use the awsvpc network mode to receive their own elastic
	// network interface, and it is not supported for other network modes. For more
	// information, see Task Networking (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
	// in the Amazon Elastic Container Service Developer Guide.
	NetworkConfiguration *NetworkConfiguration `locationName:"networkConfiguration" type:"structure"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	//
	// A total of 8192 characters are allowed for overrides. This limit includes
	// the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// An array of placement constraint objects to use for the task. You can specify
	// up to 10 constraints per task (including constraints in the task definition
	// and those specified at runtime).
	PlacementConstraints []PlacementConstraint `locationName:"placementConstraints" type:"list"`

	// The placement strategy objects to use for the task. You can specify a maximum
	// of five strategy rules per task.
	PlacementStrategy []PlacementStrategy `locationName:"placementStrategy" type:"list"`

	// The platform version the task should run. A platform version is only specified
	// for tasks using the Fargate launch type. If one is not specified, the LATEST
	// platform version is used by default. For more information, see AWS Fargate
	// Platform Versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string `locationName:"platformVersion" type:"string"`

	// Specifies whether to propagate the tags from the task definition to the task.
	// If no value is specified, the tags are not propagated. Tags can only be propagated
	// to the task during task creation. To add tags to a task after task creation,
	// use the TagResource API action.
	//
	// An error will be received if you specify the SERVICE option when running
	// a task.
	PropagateTags PropagateTags `locationName:"propagateTags" type:"string" enum:"true"`

	// An optional tag specified when a task is started. For example, if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value. Up to 36 letters (uppercase and lowercase), numbers,
	// hyphens, and underscores are allowed.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The metadata that you apply to the task to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per resource - 50
	//
	//    * For each resource, each tag key must be unique, and each tag key can
	//    have only one value.
	//
	//    * Maximum key length - 128 Unicode characters in UTF-8
	//
	//    * Maximum value length - 256 Unicode characters in UTF-8
	//
	//    * If your tagging schema is used across multiple services and resources,
	//    remember that other services may have restrictions on allowed characters.
	//    Generally allowed characters are: letters, numbers, and spaces representable
	//    in UTF-8, and the following characters: + - = . _ : / @.
	//
	//    * Tag keys and values are case-sensitive.
	//
	//    * Do not use aws:, AWS:, or any upper or lowercase combination of such
	//    as a prefix for either keys or values as it is reserved for AWS use. You
	//    cannot edit or delete tag keys or values with this prefix. Tags with this
	//    prefix do not count against your tags per resource limit.
	Tags []Tag `locationName:"tags" type:"list"`

	// The family and revision (family:revision) or full ARN of the task definition
	// to run. If a revision is not specified, the latest ACTIVE revision is used.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s RunTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunTaskInput"}

	if s.TaskDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskDefinition"))
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Overrides != nil {
		if err := s.Overrides.Validate(); err != nil {
			invalidParams.AddNested("Overrides", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RunTaskOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were run. The tasks that were successfully
	// placed on your cluster are described here.
	Tasks []Task `locationName:"tasks" type:"list"`
}

// String returns the string representation
func (s RunTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opRunTask = "RunTask"

// RunTaskRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Starts a new task using the specified task definition.
//
// You can allow Amazon ECS to place tasks for you, or you can customize how
// Amazon ECS places tasks using placement constraints and placement strategies.
// For more information, see Scheduling Tasks (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html)
// in the Amazon Elastic Container Service Developer Guide.
//
// Alternatively, you can use StartTask to use your own scheduler or place tasks
// manually on specific container instances.
//
// The Amazon ECS API follows an eventual consistency model, due to the distributed
// nature of the system supporting the API. This means that the result of an
// API command you run that affects your Amazon ECS resources might not be immediately
// visible to all subsequent commands you run. Keep this in mind when you carry
// out an API command that immediately follows a previous API command.
//
// To manage eventual consistency, you can do the following:
//
//    * Confirm the state of the resource before you run a command to modify
//    it. Run the DescribeTasks command using an exponential backoff algorithm
//    to ensure that you allow enough time for the previous command to propagate
//    through the system. To do this, run the DescribeTasks command repeatedly,
//    starting with a couple of seconds of wait time and increasing gradually
//    up to five minutes of wait time.
//
//    * Add wait time between subsequent commands, even if the DescribeTasks
//    command returns an accurate response. Apply an exponential backoff algorithm
//    starting with a couple of seconds of wait time, and increase gradually
//    up to about five minutes of wait time.
//
//    // Example sending a request using RunTaskRequest.
//    req := client.RunTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/RunTask
func (c *Client) RunTaskRequest(input *RunTaskInput) RunTaskRequest {
	op := &aws.Operation{
		Name:       opRunTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RunTaskInput{}
	}

	req := c.newRequest(op, input, &RunTaskOutput{})
	return RunTaskRequest{Request: req, Input: input, Copy: c.RunTaskRequest}
}

// RunTaskRequest is the request type for the
// RunTask API operation.
type RunTaskRequest struct {
	*aws.Request
	Input *RunTaskInput
	Copy  func(*RunTaskInput) RunTaskRequest
}

// Send marshals and sends the RunTask API request.
func (r RunTaskRequest) Send(ctx context.Context) (*RunTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RunTaskResponse{
		RunTaskOutput: r.Request.Data.(*RunTaskOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RunTaskResponse is the response type for the
// RunTask API operation.
type RunTaskResponse struct {
	*RunTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RunTask request.
func (r *RunTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
