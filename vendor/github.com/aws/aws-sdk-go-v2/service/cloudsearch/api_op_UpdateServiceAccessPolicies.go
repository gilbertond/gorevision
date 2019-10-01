// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the UpdateServiceAccessPolicies operation.
// Specifies the name of the domain you want to update and the access rules
// you want to configure.
type UpdateServiceAccessPoliciesInput struct {
	_ struct{} `type:"structure"`

	// The access rules you want to configure. These rules replace any existing
	// rules.
	//
	// AccessPolicies is a required field
	AccessPolicies *string `type:"string" required:"true"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateServiceAccessPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceAccessPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceAccessPoliciesInput"}

	if s.AccessPolicies == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessPolicies"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of an UpdateServiceAccessPolicies request. Contains the new access
// policies.
type UpdateServiceAccessPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The access rules configured for the domain.
	//
	// AccessPolicies is a required field
	AccessPolicies *AccessPoliciesStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateServiceAccessPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateServiceAccessPolicies = "UpdateServiceAccessPolicies"

// UpdateServiceAccessPoliciesRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Configures the access rules that control access to the domain's document
// and search endpoints. For more information, see Configuring Access for an
// Amazon CloudSearch Domain (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html).
//
//    // Example sending a request using UpdateServiceAccessPoliciesRequest.
//    req := client.UpdateServiceAccessPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateServiceAccessPoliciesRequest(input *UpdateServiceAccessPoliciesInput) UpdateServiceAccessPoliciesRequest {
	op := &aws.Operation{
		Name:       opUpdateServiceAccessPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceAccessPoliciesInput{}
	}

	req := c.newRequest(op, input, &UpdateServiceAccessPoliciesOutput{})
	return UpdateServiceAccessPoliciesRequest{Request: req, Input: input, Copy: c.UpdateServiceAccessPoliciesRequest}
}

// UpdateServiceAccessPoliciesRequest is the request type for the
// UpdateServiceAccessPolicies API operation.
type UpdateServiceAccessPoliciesRequest struct {
	*aws.Request
	Input *UpdateServiceAccessPoliciesInput
	Copy  func(*UpdateServiceAccessPoliciesInput) UpdateServiceAccessPoliciesRequest
}

// Send marshals and sends the UpdateServiceAccessPolicies API request.
func (r UpdateServiceAccessPoliciesRequest) Send(ctx context.Context) (*UpdateServiceAccessPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceAccessPoliciesResponse{
		UpdateServiceAccessPoliciesOutput: r.Request.Data.(*UpdateServiceAccessPoliciesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceAccessPoliciesResponse is the response type for the
// UpdateServiceAccessPolicies API operation.
type UpdateServiceAccessPoliciesResponse struct {
	*UpdateServiceAccessPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServiceAccessPolicies request.
func (r *UpdateServiceAccessPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
