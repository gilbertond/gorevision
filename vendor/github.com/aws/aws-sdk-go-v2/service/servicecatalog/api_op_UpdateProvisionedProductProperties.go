// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateProvisionedProductPropertiesInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The idempotency token that uniquely identifies the provisioning product update
	// request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The identifier of the provisioned product.
	//
	// ProvisionedProductId is a required field
	ProvisionedProductId *string `min:"1" type:"string" required:"true"`

	// A map that contains the provisioned product properties to be updated.
	//
	// The OWNER key only accepts user ARNs. The owner is the user that is allowed
	// to see, update, terminate, and execute service actions in the provisioned
	// product.
	//
	// The administrator can change the owner of a provisioned product to another
	// IAM user within the same account. Both end user owners and administrators
	// can see ownership history of the provisioned product using the ListRecordHistory
	// API. The new owner can describe all past records for the provisioned product
	// using the DescribeRecord API. The previous owner can no longer use DescribeRecord,
	// but can still see the product's history from when he was an owner using ListRecordHistory.
	//
	// If a provisioned product ownership is assigned to an end user, they can see
	// and perform any action through the API or Service Catalog console such as
	// update, terminate, and execute service actions. If an end user provisions
	// a product and the owner is updated to someone else, they will no longer be
	// able to see or perform any actions through API or the Service Catalog console
	// on that provisioned product.
	//
	// ProvisionedProductProperties is a required field
	ProvisionedProductProperties map[string]string `min:"1" type:"map" required:"true"`
}

// String returns the string representation
func (s UpdateProvisionedProductPropertiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProvisionedProductPropertiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProvisionedProductPropertiesInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.ProvisionedProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisionedProductId"))
	}
	if s.ProvisionedProductId != nil && len(*s.ProvisionedProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionedProductId", 1))
	}

	if s.ProvisionedProductProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisionedProductProperties"))
	}
	if s.ProvisionedProductProperties != nil && len(s.ProvisionedProductProperties) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionedProductProperties", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateProvisionedProductPropertiesOutput struct {
	_ struct{} `type:"structure"`

	// The provisioned product identifier.
	ProvisionedProductId *string `min:"1" type:"string"`

	// A map that contains the properties updated.
	ProvisionedProductProperties map[string]string `min:"1" type:"map"`

	// The identifier of the record.
	RecordId *string `min:"1" type:"string"`

	// The status of the request.
	Status RecordStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateProvisionedProductPropertiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateProvisionedProductProperties = "UpdateProvisionedProductProperties"

// UpdateProvisionedProductPropertiesRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Requests updates to the properties of the specified provisioned product.
//
//    // Example sending a request using UpdateProvisionedProductPropertiesRequest.
//    req := client.UpdateProvisionedProductPropertiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateProvisionedProductProperties
func (c *Client) UpdateProvisionedProductPropertiesRequest(input *UpdateProvisionedProductPropertiesInput) UpdateProvisionedProductPropertiesRequest {
	op := &aws.Operation{
		Name:       opUpdateProvisionedProductProperties,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateProvisionedProductPropertiesInput{}
	}

	req := c.newRequest(op, input, &UpdateProvisionedProductPropertiesOutput{})
	return UpdateProvisionedProductPropertiesRequest{Request: req, Input: input, Copy: c.UpdateProvisionedProductPropertiesRequest}
}

// UpdateProvisionedProductPropertiesRequest is the request type for the
// UpdateProvisionedProductProperties API operation.
type UpdateProvisionedProductPropertiesRequest struct {
	*aws.Request
	Input *UpdateProvisionedProductPropertiesInput
	Copy  func(*UpdateProvisionedProductPropertiesInput) UpdateProvisionedProductPropertiesRequest
}

// Send marshals and sends the UpdateProvisionedProductProperties API request.
func (r UpdateProvisionedProductPropertiesRequest) Send(ctx context.Context) (*UpdateProvisionedProductPropertiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProvisionedProductPropertiesResponse{
		UpdateProvisionedProductPropertiesOutput: r.Request.Data.(*UpdateProvisionedProductPropertiesOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProvisionedProductPropertiesResponse is the response type for the
// UpdateProvisionedProductProperties API operation.
type UpdateProvisionedProductPropertiesResponse struct {
	*UpdateProvisionedProductPropertiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProvisionedProductProperties request.
func (r *UpdateProvisionedProductPropertiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
