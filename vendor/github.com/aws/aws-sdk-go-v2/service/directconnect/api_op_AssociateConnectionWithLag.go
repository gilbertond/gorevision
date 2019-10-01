// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateConnectionWithLagInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// The ID of the LAG with which to associate the connection.
	//
	// LagId is a required field
	LagId *string `locationName:"lagId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateConnectionWithLagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateConnectionWithLagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateConnectionWithLagInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.LagId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LagId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about an AWS Direct Connect connection.
type AssociateConnectionWithLagOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The bandwidth of the connection.
	Bandwidth *string `locationName:"bandwidth" type:"string"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The name of the connection.
	ConnectionName *string `locationName:"connectionName" type:"string"`

	// The state of the connection. The following are the possible values:
	//
	//    * ordering: The initial state of a hosted connection provisioned on an
	//    interconnect. The connection stays in the ordering state until the owner
	//    of the hosted connection confirms or declines the connection order.
	//
	//    * requested: The initial state of a standard connection. The connection
	//    stays in the requested state until the Letter of Authorization (LOA) is
	//    sent to the customer.
	//
	//    * pending: The connection has been approved and is being initialized.
	//
	//    * available: The network link is up and the connection is ready for use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The connection is being deleted.
	//
	//    * deleted: The connection has been deleted.
	//
	//    * rejected: A hosted connection in the ordering state enters the rejected
	//    state if it is deleted by the customer.
	//
	//    * unknown: The state of the connection is not available.
	ConnectionState ConnectionState `locationName:"connectionState" type:"string" enum:"true"`

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy HasLogicalRedundancy `locationName:"hasLogicalRedundancy" type:"string" enum:"true"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time `locationName:"loaIssueTime" type:"timestamp"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string"`

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The name of the AWS Direct Connect service provider associated with the connection.
	PartnerName *string `locationName:"partnerName" type:"string"`

	// The AWS Region where the connection is located.
	Region *string `locationName:"region" type:"string"`

	// Any tags assigned to the connection.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`

	// The ID of the VLAN.
	Vlan *int64 `locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s AssociateConnectionWithLagOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateConnectionWithLag = "AssociateConnectionWithLag"

// AssociateConnectionWithLagRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Associates an existing connection with a link aggregation group (LAG). The
// connection is interrupted and re-established as a member of the LAG (connectivity
// to AWS is interrupted). The connection must be hosted on the same AWS Direct
// Connect endpoint as the LAG, and its bandwidth must match the bandwidth for
// the LAG. You can re-associate a connection that's currently associated with
// a different LAG; however, if removing the connection would cause the original
// LAG to fall below its setting for minimum number of operational connections,
// the request fails.
//
// Any virtual interfaces that are directly associated with the connection are
// automatically re-associated with the LAG. If the connection was originally
// associated with a different LAG, the virtual interfaces remain associated
// with the original LAG.
//
// For interconnects, any hosted connections are automatically re-associated
// with the LAG. If the interconnect was originally associated with a different
// LAG, the hosted connections remain associated with the original LAG.
//
//    // Example sending a request using AssociateConnectionWithLagRequest.
//    req := client.AssociateConnectionWithLagRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AssociateConnectionWithLag
func (c *Client) AssociateConnectionWithLagRequest(input *AssociateConnectionWithLagInput) AssociateConnectionWithLagRequest {
	op := &aws.Operation{
		Name:       opAssociateConnectionWithLag,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateConnectionWithLagInput{}
	}

	req := c.newRequest(op, input, &AssociateConnectionWithLagOutput{})
	return AssociateConnectionWithLagRequest{Request: req, Input: input, Copy: c.AssociateConnectionWithLagRequest}
}

// AssociateConnectionWithLagRequest is the request type for the
// AssociateConnectionWithLag API operation.
type AssociateConnectionWithLagRequest struct {
	*aws.Request
	Input *AssociateConnectionWithLagInput
	Copy  func(*AssociateConnectionWithLagInput) AssociateConnectionWithLagRequest
}

// Send marshals and sends the AssociateConnectionWithLag API request.
func (r AssociateConnectionWithLagRequest) Send(ctx context.Context) (*AssociateConnectionWithLagResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateConnectionWithLagResponse{
		AssociateConnectionWithLagOutput: r.Request.Data.(*AssociateConnectionWithLagOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateConnectionWithLagResponse is the response type for the
// AssociateConnectionWithLag API operation.
type AssociateConnectionWithLagResponse struct {
	*AssociateConnectionWithLagOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateConnectionWithLag request.
func (r *AssociateConnectionWithLagResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
