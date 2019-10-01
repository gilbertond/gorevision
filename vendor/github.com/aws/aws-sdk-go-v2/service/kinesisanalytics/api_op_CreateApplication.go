// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// TBD
type CreateApplicationInput struct {
	_ struct{} `type:"structure"`

	// One or more SQL statements that read input data, transform it, and generate
	// output. For example, you can write a SQL statement that reads data from one
	// in-application stream, generates a running average of the number of advertisement
	// clicks by vendor, and insert resulting rows in another in-application stream
	// using pumps. For more information about the typical pattern, see Application
	// Code (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html).
	//
	// You can provide such series of SQL statements, where output of one statement
	// can be used as the input for the next statement. You store intermediate results
	// by creating in-application streams and pumps.
	//
	// Note that the application code must create the streams with names specified
	// in the Outputs. For example, if your Outputs defines output streams named
	// ExampleOutputStream1 and ExampleOutputStream2, then your application code
	// must create these streams.
	ApplicationCode *string `type:"string"`

	// Summary description of the application.
	ApplicationDescription *string `type:"string"`

	// Name of your Amazon Kinesis Analytics application (for example, sample-app).
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Use this parameter to configure a CloudWatch log stream to monitor application
	// configuration errors. For more information, see Working with Amazon CloudWatch
	// Logs (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html).
	CloudWatchLoggingOptions []CloudWatchLoggingOption `type:"list"`

	// Use this parameter to configure the application input.
	//
	// You can configure your application to receive input from a single streaming
	// source. In this configuration, you map this streaming source to an in-application
	// stream that is created. Your application code can then query the in-application
	// stream like a table (you can think of it as a constantly updating table).
	//
	// For the streaming source, you provide its Amazon Resource Name (ARN) and
	// format of data on the stream (for example, JSON, CSV, etc.). You also must
	// provide an IAM role that Amazon Kinesis Analytics can assume to read this
	// stream on your behalf.
	//
	// To create the in-application stream, you need to specify a schema to transform
	// your data into a schematized version used in SQL. In the schema, you provide
	// the necessary mapping of the data elements in the streaming source to record
	// columns in the in-app stream.
	Inputs []Input `type:"list"`

	// You can configure application output to write data from any of the in-application
	// streams to up to three destinations.
	//
	// These destinations can be Amazon Kinesis streams, Amazon Kinesis Firehose
	// delivery streams, AWS Lambda destinations, or any combination of the three.
	//
	// In the configuration, you specify the in-application stream name, the destination
	// stream or Lambda function Amazon Resource Name (ARN), and the format to use
	// when writing data. You must also provide an IAM role that Amazon Kinesis
	// Analytics can assume to write to the destination stream or Lambda function
	// on your behalf.
	//
	// In the output configuration, you also provide the output stream or Lambda
	// function ARN. For stream destinations, you provide the format of data in
	// the stream (for example, JSON, CSV). You also must provide an IAM role that
	// Amazon Kinesis Analytics can assume to write to the stream or Lambda function
	// on your behalf.
	Outputs []Output `type:"list"`

	// A list of one or more tags to assign to the application. A tag is a key-value
	// pair that identifies an application. Note that the maximum number of application
	// tags includes system tags. The maximum number of user-defined application
	// tags is 50. For more information, see Using Tagging (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html).
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.CloudWatchLoggingOptions != nil {
		for i, v := range s.CloudWatchLoggingOptions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CloudWatchLoggingOptions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Inputs != nil {
		for i, v := range s.Inputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Inputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Outputs != nil {
		for i, v := range s.Outputs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Outputs", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// TBD
type CreateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// In response to your CreateApplication request, Amazon Kinesis Analytics returns
	// a response with a summary of the application it created, including the application
	// Amazon Resource Name (ARN), name, and status.
	//
	// ApplicationSummary is a required field
	ApplicationSummary *ApplicationSummary `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApplication = "CreateApplication"

// CreateApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Creates an Amazon Kinesis Analytics application. You can configure each application
// with one streaming source as input, application code to process the input,
// and up to three destinations where you want Amazon Kinesis Analytics to write
// the output data from your application. For an overview, see How it Works
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works.html).
//
// In the input configuration, you map the streaming source to an in-application
// stream, which you can think of as a constantly updating table. In the mapping,
// you must provide a schema for the in-application stream and map each data
// column in the in-application stream to a data element in the streaming source.
//
// Your application code is one or more SQL statements that read input data,
// transform it, and generate output. Your application code can create one or
// more SQL artifacts like SQL streams or pumps.
//
// In the output configuration, you can configure the application to write data
// from in-application streams created in your applications to up to three destinations.
//
// To read data from your source stream or write data to destination streams,
// Amazon Kinesis Analytics needs your permissions. You grant these permissions
// by creating IAM roles. This operation requires permissions to perform the
// kinesisanalytics:CreateApplication action.
//
// For introductory exercises to create an Amazon Kinesis Analytics application,
// see Getting Started (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/getting-started.html).
//
//    // Example sending a request using CreateApplicationRequest.
//    req := client.CreateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/CreateApplication
func (c *Client) CreateApplicationRequest(input *CreateApplicationInput) CreateApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationOutput{})
	return CreateApplicationRequest{Request: req, Input: input, Copy: c.CreateApplicationRequest}
}

// CreateApplicationRequest is the request type for the
// CreateApplication API operation.
type CreateApplicationRequest struct {
	*aws.Request
	Input *CreateApplicationInput
	Copy  func(*CreateApplicationInput) CreateApplicationRequest
}

// Send marshals and sends the CreateApplication API request.
func (r CreateApplicationRequest) Send(ctx context.Context) (*CreateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationResponse{
		CreateApplicationOutput: r.Request.Data.(*CreateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationResponse is the response type for the
// CreateApplication API operation.
type CreateApplicationResponse struct {
	*CreateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplication request.
func (r *CreateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
