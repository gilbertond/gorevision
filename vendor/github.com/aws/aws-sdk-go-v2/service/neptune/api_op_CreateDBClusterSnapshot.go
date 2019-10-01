// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDBClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DB cluster to create a snapshot for. This parameter
	// is not case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster.
	//
	// Example: my-cluster1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The identifier of the DB cluster snapshot. This parameter is stored as a
	// lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1-snapshot1
	//
	// DBClusterSnapshotIdentifier is a required field
	DBClusterSnapshotIdentifier *string `type:"string" required:"true"`

	// The tags to be assigned to the DB cluster snapshot.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateDBClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBClusterSnapshotInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.DBClusterSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details for an Amazon Neptune DB cluster snapshot
	//
	// This data type is used as a response element in the DescribeDBClusterSnapshots
	// action.
	DBClusterSnapshot *DBClusterSnapshot `type:"structure"`
}

// String returns the string representation
func (s CreateDBClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBClusterSnapshot = "CreateDBClusterSnapshot"

// CreateDBClusterSnapshotRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Creates a snapshot of a DB cluster.
//
//    // Example sending a request using CreateDBClusterSnapshotRequest.
//    req := client.CreateDBClusterSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBClusterSnapshot
func (c *Client) CreateDBClusterSnapshotRequest(input *CreateDBClusterSnapshotInput) CreateDBClusterSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateDBClusterSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBClusterSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateDBClusterSnapshotOutput{})
	return CreateDBClusterSnapshotRequest{Request: req, Input: input, Copy: c.CreateDBClusterSnapshotRequest}
}

// CreateDBClusterSnapshotRequest is the request type for the
// CreateDBClusterSnapshot API operation.
type CreateDBClusterSnapshotRequest struct {
	*aws.Request
	Input *CreateDBClusterSnapshotInput
	Copy  func(*CreateDBClusterSnapshotInput) CreateDBClusterSnapshotRequest
}

// Send marshals and sends the CreateDBClusterSnapshot API request.
func (r CreateDBClusterSnapshotRequest) Send(ctx context.Context) (*CreateDBClusterSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBClusterSnapshotResponse{
		CreateDBClusterSnapshotOutput: r.Request.Data.(*CreateDBClusterSnapshotOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBClusterSnapshotResponse is the response type for the
// CreateDBClusterSnapshot API operation.
type CreateDBClusterSnapshotResponse struct {
	*CreateDBClusterSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBClusterSnapshot request.
func (r *CreateDBClusterSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
