// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBucketLoggingInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketLoggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketLoggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLoggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetBucketLoggingOutput struct {
	_ struct{} `type:"structure"`

	// Describes where logs are stored and the prefix that Amazon S3 assigns to
	// all log object keys for a bucket. For more information, see PUT Bucket logging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html)
	// in the Amazon Simple Storage Service API Reference.
	LoggingEnabled *LoggingEnabled `type:"structure"`
}

// String returns the string representation
func (s GetBucketLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLoggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LoggingEnabled != nil {
		v := s.LoggingEnabled

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LoggingEnabled", v, metadata)
	}
	return nil
}

const opGetBucketLogging = "GetBucketLogging"

// GetBucketLoggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the logging status of a bucket and the permissions users have to
// view and modify that status. To use GET, you must be the bucket owner.
//
//    // Example sending a request using GetBucketLoggingRequest.
//    req := client.GetBucketLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLogging
func (c *Client) GetBucketLoggingRequest(input *GetBucketLoggingInput) GetBucketLoggingRequest {
	op := &aws.Operation{
		Name:       opGetBucketLogging,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?logging",
	}

	if input == nil {
		input = &GetBucketLoggingInput{}
	}

	req := c.newRequest(op, input, &GetBucketLoggingOutput{})
	return GetBucketLoggingRequest{Request: req, Input: input, Copy: c.GetBucketLoggingRequest}
}

// GetBucketLoggingRequest is the request type for the
// GetBucketLogging API operation.
type GetBucketLoggingRequest struct {
	*aws.Request
	Input *GetBucketLoggingInput
	Copy  func(*GetBucketLoggingInput) GetBucketLoggingRequest
}

// Send marshals and sends the GetBucketLogging API request.
func (r GetBucketLoggingRequest) Send(ctx context.Context) (*GetBucketLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketLoggingResponse{
		GetBucketLoggingOutput: r.Request.Data.(*GetBucketLoggingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketLoggingResponse is the response type for the
// GetBucketLogging API operation.
type GetBucketLoggingResponse struct {
	*GetBucketLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketLogging request.
func (r *GetBucketLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
