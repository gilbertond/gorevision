// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetDeployment operation.
type GetDeploymentInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of a deployment associated with the IAM user or AWS account.
	//
	// DeploymentId is a required field
	DeploymentId *string `locationName:"deploymentId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentInput"}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetDeployment operation.
type GetDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deployment.
	DeploymentInfo *DeploymentInfo `locationName:"deploymentInfo" type:"structure"`
}

// String returns the string representation
func (s GetDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDeployment = "GetDeployment"

// GetDeploymentRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about a deployment.
//
// The content property of the appSpecContent object in the returned revision
// is always null. Use GetApplicationRevision and the sha256 property of the
// returned appSpecContent object to get the content of the deployment’s AppSpec
// file.
//
//    // Example sending a request using GetDeploymentRequest.
//    req := client.GetDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetDeployment
func (c *Client) GetDeploymentRequest(input *GetDeploymentInput) GetDeploymentRequest {
	op := &aws.Operation{
		Name:       opGetDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeploymentInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentOutput{})
	return GetDeploymentRequest{Request: req, Input: input, Copy: c.GetDeploymentRequest}
}

// GetDeploymentRequest is the request type for the
// GetDeployment API operation.
type GetDeploymentRequest struct {
	*aws.Request
	Input *GetDeploymentInput
	Copy  func(*GetDeploymentInput) GetDeploymentRequest
}

// Send marshals and sends the GetDeployment API request.
func (r GetDeploymentRequest) Send(ctx context.Context) (*GetDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentResponse{
		GetDeploymentOutput: r.Request.Data.(*GetDeploymentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentResponse is the response type for the
// GetDeployment API operation.
type GetDeploymentResponse struct {
	*GetDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeployment request.
func (r *GetDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
