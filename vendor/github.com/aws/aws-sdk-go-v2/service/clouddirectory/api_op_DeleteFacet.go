// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteFacetInput struct {
	_ struct{} `type:"structure"`

	// The name of the facet to delete.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Facet. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFacetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFacetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFacetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFacetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteFacetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFacetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteFacetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteFacet = "DeleteFacet"

// DeleteFacetRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Deletes a given Facet. All attributes and Rules that are associated with
// the facet will be deleted. Only development schema facets are allowed deletion.
//
//    // Example sending a request using DeleteFacetRequest.
//    req := client.DeleteFacetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteFacet
func (c *Client) DeleteFacetRequest(input *DeleteFacetInput) DeleteFacetRequest {
	op := &aws.Operation{
		Name:       opDeleteFacet,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/facet/delete",
	}

	if input == nil {
		input = &DeleteFacetInput{}
	}

	req := c.newRequest(op, input, &DeleteFacetOutput{})
	return DeleteFacetRequest{Request: req, Input: input, Copy: c.DeleteFacetRequest}
}

// DeleteFacetRequest is the request type for the
// DeleteFacet API operation.
type DeleteFacetRequest struct {
	*aws.Request
	Input *DeleteFacetInput
	Copy  func(*DeleteFacetInput) DeleteFacetRequest
}

// Send marshals and sends the DeleteFacet API request.
func (r DeleteFacetRequest) Send(ctx context.Context) (*DeleteFacetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFacetResponse{
		DeleteFacetOutput: r.Request.Data.(*DeleteFacetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFacetResponse is the response type for the
// DeleteFacet API operation.
type DeleteFacetResponse struct {
	*DeleteFacetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFacet request.
func (r *DeleteFacetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
