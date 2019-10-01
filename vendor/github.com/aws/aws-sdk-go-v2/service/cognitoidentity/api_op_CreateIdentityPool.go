// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the CreateIdentityPool action.
type CreateIdentityPoolInput struct {
	_ struct{} `type:"structure"`

	// TRUE if the identity pool supports unauthenticated logins.
	//
	// AllowUnauthenticatedIdentities is a required field
	AllowUnauthenticatedIdentities *bool `type:"boolean" required:"true"`

	// An array of Amazon Cognito user pools and their client IDs.
	CognitoIdentityProviders []CognitoIdentityProvider `type:"list"`

	// The "domain" by which Cognito will refer to your users. This name acts as
	// a placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use
	// letters as well as period (.), underscore (_), and dash (-).
	//
	// Once you have set a developer provider name, you cannot change it. Please
	// take care in setting this parameter.
	DeveloperProviderName *string `min:"1" type:"string"`

	// A string that you provide.
	//
	// IdentityPoolName is a required field
	IdentityPoolName *string `min:"1" type:"string" required:"true"`

	// Tags to assign to the identity pool. A tag is a label that you can apply
	// to identity pools to categorize and manage them in different ways, such as
	// by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]string `type:"map"`

	// A list of OpendID Connect provider ARNs.
	OpenIdConnectProviderARNs []string `type:"list"`

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []string `type:"list"`

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateIdentityPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIdentityPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIdentityPoolInput"}

	if s.AllowUnauthenticatedIdentities == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllowUnauthenticatedIdentities"))
	}
	if s.DeveloperProviderName != nil && len(*s.DeveloperProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeveloperProviderName", 1))
	}

	if s.IdentityPoolName == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolName"))
	}
	if s.IdentityPoolName != nil && len(*s.IdentityPoolName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolName", 1))
	}
	if s.CognitoIdentityProviders != nil {
		for i, v := range s.CognitoIdentityProviders {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CognitoIdentityProviders", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object representing an Amazon Cognito identity pool.
type CreateIdentityPoolOutput struct {
	_ struct{} `type:"structure"`

	// TRUE if the identity pool supports unauthenticated logins.
	//
	// AllowUnauthenticatedIdentities is a required field
	AllowUnauthenticatedIdentities *bool `type:"boolean" required:"true"`

	// A list representing an Amazon Cognito user pool and its client ID.
	CognitoIdentityProviders []CognitoIdentityProvider `type:"list"`

	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string `min:"1" type:"string"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`

	// A string that you provide.
	//
	// IdentityPoolName is a required field
	IdentityPoolName *string `min:"1" type:"string" required:"true"`

	// The tags that are assigned to the identity pool. A tag is a label that you
	// can apply to identity pools to categorize and manage them in different ways,
	// such as by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]string `type:"map"`

	// A list of OpendID Connect provider ARNs.
	OpenIdConnectProviderARNs []string `type:"list"`

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []string `type:"list"`

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateIdentityPoolOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateIdentityPool = "CreateIdentityPool"

// CreateIdentityPoolRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your AWS account. The limit on identity pools
// is 60 per account. The keys for SupportedLoginProviders are as follows:
//
//    * Facebook: graph.facebook.com
//
//    * Google: accounts.google.com
//
//    * Amazon: www.amazon.com
//
//    * Twitter: api.twitter.com
//
//    * Digits: www.digits.com
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using CreateIdentityPoolRequest.
//    req := client.CreateIdentityPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/CreateIdentityPool
func (c *Client) CreateIdentityPoolRequest(input *CreateIdentityPoolInput) CreateIdentityPoolRequest {
	op := &aws.Operation{
		Name:       opCreateIdentityPool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateIdentityPoolInput{}
	}

	req := c.newRequest(op, input, &CreateIdentityPoolOutput{})
	return CreateIdentityPoolRequest{Request: req, Input: input, Copy: c.CreateIdentityPoolRequest}
}

// CreateIdentityPoolRequest is the request type for the
// CreateIdentityPool API operation.
type CreateIdentityPoolRequest struct {
	*aws.Request
	Input *CreateIdentityPoolInput
	Copy  func(*CreateIdentityPoolInput) CreateIdentityPoolRequest
}

// Send marshals and sends the CreateIdentityPool API request.
func (r CreateIdentityPoolRequest) Send(ctx context.Context) (*CreateIdentityPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIdentityPoolResponse{
		CreateIdentityPoolOutput: r.Request.Data.(*CreateIdentityPoolOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIdentityPoolResponse is the response type for the
// CreateIdentityPool API operation.
type CreateIdentityPoolResponse struct {
	*CreateIdentityPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIdentityPool request.
func (r *CreateIdentityPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
