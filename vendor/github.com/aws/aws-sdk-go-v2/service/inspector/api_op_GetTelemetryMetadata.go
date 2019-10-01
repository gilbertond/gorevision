// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTelemetryMetadataInput struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the assessment run that has the telemetry data that
	// you want to obtain.
	//
	// AssessmentRunArn is a required field
	AssessmentRunArn *string `locationName:"assessmentRunArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTelemetryMetadataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTelemetryMetadataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTelemetryMetadataInput"}

	if s.AssessmentRunArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssessmentRunArn"))
	}
	if s.AssessmentRunArn != nil && len(*s.AssessmentRunArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssessmentRunArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTelemetryMetadataOutput struct {
	_ struct{} `type:"structure"`

	// Telemetry details.
	//
	// TelemetryMetadata is a required field
	TelemetryMetadata []TelemetryMetadata `locationName:"telemetryMetadata" type:"list" required:"true"`
}

// String returns the string representation
func (s GetTelemetryMetadataOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTelemetryMetadata = "GetTelemetryMetadata"

// GetTelemetryMetadataRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Information about the data that is collected for the specified assessment
// run.
//
//    // Example sending a request using GetTelemetryMetadataRequest.
//    req := client.GetTelemetryMetadataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/GetTelemetryMetadata
func (c *Client) GetTelemetryMetadataRequest(input *GetTelemetryMetadataInput) GetTelemetryMetadataRequest {
	op := &aws.Operation{
		Name:       opGetTelemetryMetadata,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTelemetryMetadataInput{}
	}

	req := c.newRequest(op, input, &GetTelemetryMetadataOutput{})
	return GetTelemetryMetadataRequest{Request: req, Input: input, Copy: c.GetTelemetryMetadataRequest}
}

// GetTelemetryMetadataRequest is the request type for the
// GetTelemetryMetadata API operation.
type GetTelemetryMetadataRequest struct {
	*aws.Request
	Input *GetTelemetryMetadataInput
	Copy  func(*GetTelemetryMetadataInput) GetTelemetryMetadataRequest
}

// Send marshals and sends the GetTelemetryMetadata API request.
func (r GetTelemetryMetadataRequest) Send(ctx context.Context) (*GetTelemetryMetadataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTelemetryMetadataResponse{
		GetTelemetryMetadataOutput: r.Request.Data.(*GetTelemetryMetadataOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTelemetryMetadataResponse is the response type for the
// GetTelemetryMetadata API operation.
type GetTelemetryMetadataResponse struct {
	*GetTelemetryMetadataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTelemetryMetadata request.
func (r *GetTelemetryMetadataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
