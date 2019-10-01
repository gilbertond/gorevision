// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateJobInput struct {
	_ struct{} `type:"structure"`

	// The ID for the address that you want the Snowball shipped to.
	AddressId *string `min:"40" type:"string"`

	// The ID of a cluster. If you're creating a job for a node in a cluster, you
	// need to provide only this clusterId value. The other job attributes are inherited
	// from the cluster.
	ClusterId *string `min:"39" type:"string"`

	// Defines an optional description of this specific job, for example Important
	// Photos 2016-08-11.
	Description *string `min:"1" type:"string"`

	// The forwarding address ID for a job. This field is not supported in most
	// regions.
	ForwardingAddressId *string `min:"40" type:"string"`

	// Defines the type of job that you're creating.
	JobType JobType `type:"string" enum:"true"`

	// The KmsKeyARN that you want to associate with this job. KmsKeyARNs are created
	// using the CreateKey (http://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html)
	// AWS Key Management Service (KMS) API action.
	KmsKeyARN *string `type:"string"`

	// Defines the Amazon Simple Notification Service (Amazon SNS) notification
	// settings for this job.
	Notification *Notification `type:"structure"`

	// Defines the Amazon S3 buckets associated with this job.
	//
	// With IMPORT jobs, you specify the bucket or buckets that your transferred
	// data will be imported into.
	//
	// With EXPORT jobs, you specify the bucket or buckets that your transferred
	// data will be exported from. Optionally, you can also specify a KeyRange value.
	// If you choose to export a range, you define the length of the range by providing
	// either an inclusive BeginMarker value, an inclusive EndMarker value, or both.
	// Ranges are UTF-8 binary sorted.
	Resources *JobResource `type:"structure"`

	// The RoleARN that you want to associate with this job. RoleArns are created
	// using the CreateRole (http://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
	// AWS Identity and Access Management (IAM) API action.
	RoleARN *string `type:"string"`

	// The shipping speed for this job. This speed doesn't dictate how soon you'll
	// get the Snowball, rather it represents how quickly the Snowball moves to
	// its destination while in transit. Regional shipping speeds are as follows:
	//
	//    * In Australia, you have access to express shipping. Typically, Snowballs
	//    shipped express are delivered in about a day.
	//
	//    * In the European Union (EU), you have access to express shipping. Typically,
	//    Snowballs shipped express are delivered in about a day. In addition, most
	//    countries in the EU have access to standard shipping, which typically
	//    takes less than a week, one way.
	//
	//    * In India, Snowballs are delivered in one to seven days.
	//
	//    * In the US, you have access to one-day shipping and two-day shipping.
	ShippingOption ShippingOption `type:"string" enum:"true"`

	// If your job is being created in one of the US regions, you have the option
	// of specifying what size Snowball you'd like for this job. In all other regions,
	// Snowballs come with 80 TB in storage capacity.
	SnowballCapacityPreference SnowballCapacity `type:"string" enum:"true"`

	// The type of AWS Snowball device to use for this job. The only supported device
	// types for cluster jobs are EDGE, EDGE_C, and EDGE_CG.
	SnowballType SnowballType `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateJobInput"}
	if s.AddressId != nil && len(*s.AddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("AddressId", 40))
	}
	if s.ClusterId != nil && len(*s.ClusterId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("ClusterId", 39))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.ForwardingAddressId != nil && len(*s.ForwardingAddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("ForwardingAddressId", 40))
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

type CreateJobOutput struct {
	_ struct{} `type:"structure"`

	// The automatically generated ID for a job, for example JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string `min:"39" type:"string"`
}

// String returns the string representation
func (s CreateJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateJob = "CreateJob"

// CreateJobRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Creates a job to import or export data between Amazon S3 and your on-premises
// data center. Your AWS account must have the right trust policies and permissions
// in place to create a job for Snowball. If you're creating a job for a node
// in a cluster, you only need to provide the clusterId value; the other job
// attributes are inherited from the cluster.
//
//    // Example sending a request using CreateJobRequest.
//    req := client.CreateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/CreateJob
func (c *Client) CreateJobRequest(input *CreateJobInput) CreateJobRequest {
	op := &aws.Operation{
		Name:       opCreateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateJobInput{}
	}

	req := c.newRequest(op, input, &CreateJobOutput{})
	return CreateJobRequest{Request: req, Input: input, Copy: c.CreateJobRequest}
}

// CreateJobRequest is the request type for the
// CreateJob API operation.
type CreateJobRequest struct {
	*aws.Request
	Input *CreateJobInput
	Copy  func(*CreateJobInput) CreateJobRequest
}

// Send marshals and sends the CreateJob API request.
func (r CreateJobRequest) Send(ctx context.Context) (*CreateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJobResponse{
		CreateJobOutput: r.Request.Data.(*CreateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJobResponse is the response type for the
// CreateJob API operation.
type CreateJobResponse struct {
	*CreateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJob request.
func (r *CreateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
