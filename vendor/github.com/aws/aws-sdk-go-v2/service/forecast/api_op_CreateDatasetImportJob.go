// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDatasetImportJobInput struct {
	_ struct{} `type:"structure"`

	// The location of the training data to import and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data.
	//
	// DataSource is a required field
	DataSource *DataSource `type:"structure" required:"true"`

	// The Amazon Resource Name (ARN) of the Amazon Forecast dataset that you want
	// to import data to.
	//
	// DatasetArn is a required field
	DatasetArn *string `type:"string" required:"true"`

	// The name for the dataset import job. It is recommended to include the current
	// timestamp in the name to guard against getting a ResourceAlreadyExistsException
	// exception, for example, 20190721DatasetImport.
	//
	// DatasetImportJobName is a required field
	DatasetImportJobName *string `min:"1" type:"string" required:"true"`

	// The format of timestamps in the dataset. Two formats are supported, dependent
	// on the DataFrequency specified when the dataset was created.
	//
	//    * "yyyy-MM-dd" For data frequencies: Y, M, W, and D
	//
	//    * "yyyy-MM-dd HH:mm:ss" For data frequencies: H, 30min, 15min, and 1min;
	//    and optionally, for: Y, M, W, and D
	TimestampFormat *string `type:"string"`
}

// String returns the string representation
func (s CreateDatasetImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatasetImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatasetImportJobInput"}

	if s.DataSource == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSource"))
	}

	if s.DatasetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetArn"))
	}

	if s.DatasetImportJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetImportJobName"))
	}
	if s.DatasetImportJobName != nil && len(*s.DatasetImportJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetImportJobName", 1))
	}
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			invalidParams.AddNested("DataSource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDatasetImportJobOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset import job.
	DatasetImportJobArn *string `type:"string"`
}

// String returns the string representation
func (s CreateDatasetImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDatasetImportJob = "CreateDatasetImportJob"

// CreateDatasetImportJobRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Imports your training data to an Amazon Forecast dataset. You provide the
// location of your training data in an Amazon Simple Storage Service (Amazon
// S3) bucket and the Amazon Resource Name (ARN) of the dataset that you want
// to import the data to.
//
// You must specify a DataSource object that includes an AWS Identity and Access
// Management (IAM) role that Amazon Forecast can assume to access the data.
// For more information, see aws-forecast-iam-roles.
//
// Two properties of the training data are optionally specified:
//
//    * The delimiter that separates the data fields. The default delimiter
//    is a comma (,), which is the only supported delimiter in this release.
//
//    * The format of timestamps. If the format is not specified, Amazon Forecast
//    expects the format to be "yyyy-MM-dd HH:mm:ss".
//
// When Amazon Forecast uploads your training data, it verifies that the data
// was collected at the DataFrequency specified when the target dataset was
// created. For more information, see CreateDataset and howitworks-datasets-groups.
// Amazon Forecast also verifies the delimiter and timestamp format.
//
// You can use the ListDatasetImportJobs operation to get a list of all your
// dataset import jobs, filtered by specified criteria.
//
// To get a list of all your dataset import jobs, filtered by the specified
// criteria, use the ListDatasetGroups operation.
//
//    // Example sending a request using CreateDatasetImportJobRequest.
//    req := client.CreateDatasetImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/CreateDatasetImportJob
func (c *Client) CreateDatasetImportJobRequest(input *CreateDatasetImportJobInput) CreateDatasetImportJobRequest {
	op := &aws.Operation{
		Name:       opCreateDatasetImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDatasetImportJobInput{}
	}

	req := c.newRequest(op, input, &CreateDatasetImportJobOutput{})
	return CreateDatasetImportJobRequest{Request: req, Input: input, Copy: c.CreateDatasetImportJobRequest}
}

// CreateDatasetImportJobRequest is the request type for the
// CreateDatasetImportJob API operation.
type CreateDatasetImportJobRequest struct {
	*aws.Request
	Input *CreateDatasetImportJobInput
	Copy  func(*CreateDatasetImportJobInput) CreateDatasetImportJobRequest
}

// Send marshals and sends the CreateDatasetImportJob API request.
func (r CreateDatasetImportJobRequest) Send(ctx context.Context) (*CreateDatasetImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetImportJobResponse{
		CreateDatasetImportJobOutput: r.Request.Data.(*CreateDatasetImportJobOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetImportJobResponse is the response type for the
// CreateDatasetImportJob API operation.
type CreateDatasetImportJobResponse struct {
	*CreateDatasetImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatasetImportJob request.
func (r *CreateDatasetImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
