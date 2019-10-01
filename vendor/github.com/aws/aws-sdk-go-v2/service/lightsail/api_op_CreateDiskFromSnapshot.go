// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDiskFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone where you want to create the disk (e.g., us-east-2a).
	// Choose the same Availability Zone as the Lightsail instance where you want
	// to create the disk.
	//
	// Use the GetRegions operation to list the Availability Zones where Lightsail
	// is currently available.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// The unique Lightsail disk name (e.g., my-disk).
	//
	// DiskName is a required field
	DiskName *string `locationName:"diskName" type:"string" required:"true"`

	// The name of the disk snapshot (e.g., my-snapshot) from which to create the
	// new storage disk.
	//
	// DiskSnapshotName is a required field
	DiskSnapshotName *string `locationName:"diskSnapshotName" type:"string" required:"true"`

	// The size of the disk in GB (e.g., 32).
	//
	// SizeInGb is a required field
	SizeInGb *int64 `locationName:"sizeInGb" type:"integer" required:"true"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateDiskFromSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDiskFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDiskFromSnapshotInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if s.DiskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskName"))
	}

	if s.DiskSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskSnapshotName"))
	}

	if s.SizeInGb == nil {
		invalidParams.Add(aws.NewErrParamRequired("SizeInGb"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDiskFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateDiskFromSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDiskFromSnapshot = "CreateDiskFromSnapshot"

// CreateDiskFromSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates a block storage disk from a disk snapshot that can be attached to
// a Lightsail instance in the same Availability Zone (e.g., us-east-2a). The
// disk is created in the regional endpoint that you send the HTTP request to.
// For more information, see Regions and Availability Zones in Lightsail (https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail).
//
// The create disk from snapshot operation supports tag-based access control
// via request tags and resource tags applied to the resource identified by
// diskSnapshotName. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateDiskFromSnapshotRequest.
//    req := client.CreateDiskFromSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateDiskFromSnapshot
func (c *Client) CreateDiskFromSnapshotRequest(input *CreateDiskFromSnapshotInput) CreateDiskFromSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateDiskFromSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDiskFromSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateDiskFromSnapshotOutput{})
	return CreateDiskFromSnapshotRequest{Request: req, Input: input, Copy: c.CreateDiskFromSnapshotRequest}
}

// CreateDiskFromSnapshotRequest is the request type for the
// CreateDiskFromSnapshot API operation.
type CreateDiskFromSnapshotRequest struct {
	*aws.Request
	Input *CreateDiskFromSnapshotInput
	Copy  func(*CreateDiskFromSnapshotInput) CreateDiskFromSnapshotRequest
}

// Send marshals and sends the CreateDiskFromSnapshot API request.
func (r CreateDiskFromSnapshotRequest) Send(ctx context.Context) (*CreateDiskFromSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDiskFromSnapshotResponse{
		CreateDiskFromSnapshotOutput: r.Request.Data.(*CreateDiskFromSnapshotOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDiskFromSnapshotResponse is the response type for the
// CreateDiskFromSnapshot API operation.
type CreateDiskFromSnapshotResponse struct {
	*CreateDiskFromSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDiskFromSnapshot request.
func (r *CreateDiskFromSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
