// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Creates an AWS Managed Microsoft AD directory.
type CreateMicrosoftADInput struct {
	_ struct{} `type:"structure"`

	// A textual description for the directory. This label will appear on the AWS
	// console Directory Details page after the directory is created.
	Description *string `type:"string"`

	// AWS Managed Microsoft AD is available in two editions: Standard and Enterprise.
	// Enterprise is the default.
	Edition DirectoryEdition `type:"string" enum:"true"`

	// The fully qualified domain name for the directory, such as corp.example.com.
	// This name will resolve inside your VPC only. It does not need to be publicly
	// resolvable.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The password for the default administrative user named Admin.
	//
	// Password is a required field
	Password *string `type:"string" required:"true"`

	// The NetBIOS name for your domain. A short identifier for your domain, such
	// as CORP. If you don't specify a NetBIOS name, it will default to the first
	// part of your directory DNS. For example, CORP for the directory DNS corp.example.com.
	ShortName *string `type:"string"`

	// The tags to be assigned to the AWS Managed Microsoft AD directory.
	Tags []Tag `type:"list"`

	// Contains VPC information for the CreateDirectory or CreateMicrosoftAD operation.
	//
	// VpcSettings is a required field
	VpcSettings *DirectoryVpcSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateMicrosoftADInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMicrosoftADInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMicrosoftADInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if s.VpcSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcSettings"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcSettings != nil {
		if err := s.VpcSettings.Validate(); err != nil {
			invalidParams.AddNested("VpcSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result of a CreateMicrosoftAD request.
type CreateMicrosoftADOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory that was created.
	DirectoryId *string `type:"string"`
}

// String returns the string representation
func (s CreateMicrosoftADOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateMicrosoftAD = "CreateMicrosoftAD"

// CreateMicrosoftADRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Creates an AWS Managed Microsoft AD directory.
//
// Before you call CreateMicrosoftAD, ensure that all of the required permissions
// have been explicitly granted through a policy. For details about what permissions
// are required to run the CreateMicrosoftAD operation, see AWS Directory Service
// API Permissions: Actions, Resources, and Conditions Reference (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
//
//    // Example sending a request using CreateMicrosoftADRequest.
//    req := client.CreateMicrosoftADRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateMicrosoftAD
func (c *Client) CreateMicrosoftADRequest(input *CreateMicrosoftADInput) CreateMicrosoftADRequest {
	op := &aws.Operation{
		Name:       opCreateMicrosoftAD,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMicrosoftADInput{}
	}

	req := c.newRequest(op, input, &CreateMicrosoftADOutput{})
	return CreateMicrosoftADRequest{Request: req, Input: input, Copy: c.CreateMicrosoftADRequest}
}

// CreateMicrosoftADRequest is the request type for the
// CreateMicrosoftAD API operation.
type CreateMicrosoftADRequest struct {
	*aws.Request
	Input *CreateMicrosoftADInput
	Copy  func(*CreateMicrosoftADInput) CreateMicrosoftADRequest
}

// Send marshals and sends the CreateMicrosoftAD API request.
func (r CreateMicrosoftADRequest) Send(ctx context.Context) (*CreateMicrosoftADResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMicrosoftADResponse{
		CreateMicrosoftADOutput: r.Request.Data.(*CreateMicrosoftADOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMicrosoftADResponse is the response type for the
// CreateMicrosoftAD API operation.
type CreateMicrosoftADResponse struct {
	*CreateMicrosoftADOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMicrosoftAD request.
func (r *CreateMicrosoftADResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
