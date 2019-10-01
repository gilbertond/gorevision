// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a get repository operation.
type GetRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The name of the repository to get information about.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRepositoryInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a get repository operation.
type GetRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the repository.
	RepositoryMetadata *RepositoryMetadata `locationName:"repositoryMetadata" type:"structure"`
}

// String returns the string representation
func (s GetRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRepository = "GetRepository"

// GetRepositoryRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about a repository.
//
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a web page could expose users to potentially malicious
// code. Make sure that you HTML-encode the description field in any application
// that uses this API to display the repository description on a web page.
//
//    // Example sending a request using GetRepositoryRequest.
//    req := client.GetRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetRepository
func (c *Client) GetRepositoryRequest(input *GetRepositoryInput) GetRepositoryRequest {
	op := &aws.Operation{
		Name:       opGetRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRepositoryInput{}
	}

	req := c.newRequest(op, input, &GetRepositoryOutput{})
	return GetRepositoryRequest{Request: req, Input: input, Copy: c.GetRepositoryRequest}
}

// GetRepositoryRequest is the request type for the
// GetRepository API operation.
type GetRepositoryRequest struct {
	*aws.Request
	Input *GetRepositoryInput
	Copy  func(*GetRepositoryInput) GetRepositoryRequest
}

// Send marshals and sends the GetRepository API request.
func (r GetRepositoryRequest) Send(ctx context.Context) (*GetRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRepositoryResponse{
		GetRepositoryOutput: r.Request.Data.(*GetRepositoryOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRepositoryResponse is the response type for the
// GetRepository API operation.
type GetRepositoryResponse struct {
	*GetRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRepository request.
func (r *GetRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
