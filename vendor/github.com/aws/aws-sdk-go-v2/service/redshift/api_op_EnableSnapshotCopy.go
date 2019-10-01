// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableSnapshotCopyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the source cluster to copy snapshots from.
	//
	// Constraints: Must be the valid name of an existing cluster that does not
	// already have cross-region snapshot copy enabled.
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`

	// The destination AWS Region that you want to copy snapshots to.
	//
	// Constraints: Must be the name of a valid AWS Region. For more information,
	// see Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region)
	// in the Amazon Web Services General Reference.
	//
	// DestinationRegion is a required field
	DestinationRegion *string `type:"string" required:"true"`

	// The number of days to retain newly copied snapshots in the destination AWS
	// Region after they are copied from the source AWS Region. If the value is
	// -1, the manual snapshot is retained indefinitely.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int64 `type:"integer"`

	// The number of days to retain automated snapshots in the destination region
	// after they are copied from the source region.
	//
	// Default: 7.
	//
	// Constraints: Must be at least 1 and no more than 35.
	RetentionPeriod *int64 `type:"integer"`

	// The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted
	// cluster are copied to the destination region.
	SnapshotCopyGrantName *string `type:"string"`
}

// String returns the string representation
func (s EnableSnapshotCopyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableSnapshotCopyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableSnapshotCopyInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if s.DestinationRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationRegion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableSnapshotCopyOutput struct {
	_ struct{} `type:"structure"`

	// Describes a cluster.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s EnableSnapshotCopyOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableSnapshotCopy = "EnableSnapshotCopy"

// EnableSnapshotCopyRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Enables the automatic copy of snapshots from one region to another region
// for a specified cluster.
//
//    // Example sending a request using EnableSnapshotCopyRequest.
//    req := client.EnableSnapshotCopyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/EnableSnapshotCopy
func (c *Client) EnableSnapshotCopyRequest(input *EnableSnapshotCopyInput) EnableSnapshotCopyRequest {
	op := &aws.Operation{
		Name:       opEnableSnapshotCopy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableSnapshotCopyInput{}
	}

	req := c.newRequest(op, input, &EnableSnapshotCopyOutput{})
	return EnableSnapshotCopyRequest{Request: req, Input: input, Copy: c.EnableSnapshotCopyRequest}
}

// EnableSnapshotCopyRequest is the request type for the
// EnableSnapshotCopy API operation.
type EnableSnapshotCopyRequest struct {
	*aws.Request
	Input *EnableSnapshotCopyInput
	Copy  func(*EnableSnapshotCopyInput) EnableSnapshotCopyRequest
}

// Send marshals and sends the EnableSnapshotCopy API request.
func (r EnableSnapshotCopyRequest) Send(ctx context.Context) (*EnableSnapshotCopyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableSnapshotCopyResponse{
		EnableSnapshotCopyOutput: r.Request.Data.(*EnableSnapshotCopyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableSnapshotCopyResponse is the response type for the
// EnableSnapshotCopy API operation.
type EnableSnapshotCopyResponse struct {
	*EnableSnapshotCopyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableSnapshotCopy request.
func (r *EnableSnapshotCopyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
