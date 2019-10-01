// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBlockInput struct {
	_ struct{} `type:"structure"`

	// The location of the block that you want to request. An address is an Amazon
	// Ion structure that has two fields: strandId and sequenceNo.
	//
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:14}
	//
	// BlockAddress is a required field
	BlockAddress *ValueHolder `type:"structure" required:"true"`

	// The latest block location covered by the digest for which to request a proof.
	// An address is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	//
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:49}
	DigestTipAddress *ValueHolder `type:"structure"`

	// The name of the ledger.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBlockInput"}

	if s.BlockAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("BlockAddress"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.BlockAddress != nil {
		if err := s.BlockAddress.Validate(); err != nil {
			invalidParams.AddNested("BlockAddress", err.(aws.ErrInvalidParams))
		}
	}
	if s.DigestTipAddress != nil {
		if err := s.DigestTipAddress.Validate(); err != nil {
			invalidParams.AddNested("DigestTipAddress", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBlockInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BlockAddress != nil {
		v := s.BlockAddress

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BlockAddress", v, metadata)
	}
	if s.DigestTipAddress != nil {
		v := s.DigestTipAddress

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DigestTipAddress", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetBlockOutput struct {
	_ struct{} `type:"structure"`

	// The block data object in Amazon Ion format.
	//
	// Block is a required field
	Block *ValueHolder `type:"structure" required:"true"`

	// The proof object in Amazon Ion format returned by a GetBlock request. A proof
	// contains the list of hash values required to recalculate the specified digest
	// using a Merkle tree, starting with the specified block.
	Proof *ValueHolder `type:"structure"`
}

// String returns the string representation
func (s GetBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBlockOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Block != nil {
		v := s.Block

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Block", v, metadata)
	}
	if s.Proof != nil {
		v := s.Proof

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Proof", v, metadata)
	}
	return nil
}

const opGetBlock = "GetBlock"

// GetBlockRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Returns a journal block object at a specified address in a ledger. Also returns
// a proof of the specified block for verification if DigestTipAddress is provided.
//
// If the specified ledger doesn't exist or is in DELETING status, then throws
// ResourceNotFoundException.
//
// If the specified ledger is in CREATING status, then throws ResourcePreconditionNotMetException.
//
// If no block exists with the specified address, then throws InvalidParameterException.
//
//    // Example sending a request using GetBlockRequest.
//    req := client.GetBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/GetBlock
func (c *Client) GetBlockRequest(input *GetBlockInput) GetBlockRequest {
	op := &aws.Operation{
		Name:       opGetBlock,
		HTTPMethod: "POST",
		HTTPPath:   "/ledgers/{name}/block",
	}

	if input == nil {
		input = &GetBlockInput{}
	}

	req := c.newRequest(op, input, &GetBlockOutput{})
	return GetBlockRequest{Request: req, Input: input, Copy: c.GetBlockRequest}
}

// GetBlockRequest is the request type for the
// GetBlock API operation.
type GetBlockRequest struct {
	*aws.Request
	Input *GetBlockInput
	Copy  func(*GetBlockInput) GetBlockRequest
}

// Send marshals and sends the GetBlock API request.
func (r GetBlockRequest) Send(ctx context.Context) (*GetBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBlockResponse{
		GetBlockOutput: r.Request.Data.(*GetBlockOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBlockResponse is the response type for the
// GetBlock API operation.
type GetBlockResponse struct {
	*GetBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBlock request.
func (r *GetBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
