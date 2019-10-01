// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDeleteImportDataInput struct {
	_ struct{} `type:"structure"`

	// The IDs for the import tasks that you want to delete.
	//
	// ImportTaskIds is a required field
	ImportTaskIds []string `locationName:"importTaskIds" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDeleteImportDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeleteImportDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeleteImportDataInput"}

	if s.ImportTaskIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImportTaskIds"))
	}
	if s.ImportTaskIds != nil && len(s.ImportTaskIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImportTaskIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeleteImportDataOutput struct {
	_ struct{} `type:"structure"`

	// Error messages returned for each import task that you deleted as a response
	// for this command.
	Errors []BatchDeleteImportDataError `locationName:"errors" type:"list"`
}

// String returns the string representation
func (s BatchDeleteImportDataOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeleteImportData = "BatchDeleteImportData"

// BatchDeleteImportDataRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Deletes one or more import tasks, each identified by their import ID. Each
// import task has a number of records that can identify servers or applications.
//
// AWS Application Discovery Service has built-in matching logic that will identify
// when discovered servers match existing entries that you've previously discovered,
// the information for the already-existing discovered server is updated. When
// you delete an import task that contains records that were used to match,
// the information in those matched records that comes from the deleted records
// will also be deleted.
//
//    // Example sending a request using BatchDeleteImportDataRequest.
//    req := client.BatchDeleteImportDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/BatchDeleteImportData
func (c *Client) BatchDeleteImportDataRequest(input *BatchDeleteImportDataInput) BatchDeleteImportDataRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteImportData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteImportDataInput{}
	}

	req := c.newRequest(op, input, &BatchDeleteImportDataOutput{})
	return BatchDeleteImportDataRequest{Request: req, Input: input, Copy: c.BatchDeleteImportDataRequest}
}

// BatchDeleteImportDataRequest is the request type for the
// BatchDeleteImportData API operation.
type BatchDeleteImportDataRequest struct {
	*aws.Request
	Input *BatchDeleteImportDataInput
	Copy  func(*BatchDeleteImportDataInput) BatchDeleteImportDataRequest
}

// Send marshals and sends the BatchDeleteImportData API request.
func (r BatchDeleteImportDataRequest) Send(ctx context.Context) (*BatchDeleteImportDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteImportDataResponse{
		BatchDeleteImportDataOutput: r.Request.Data.(*BatchDeleteImportDataOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteImportDataResponse is the response type for the
// BatchDeleteImportData API operation.
type BatchDeleteImportDataResponse struct {
	*BatchDeleteImportDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteImportData request.
func (r *BatchDeleteImportDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
