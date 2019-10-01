// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The portfolio identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePortfolioInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePortfolioOutput struct {
	_ struct{} `type:"structure"`

	// Information about the associated budgets.
	Budgets []BudgetDetail `type:"list"`

	// Information about the portfolio.
	PortfolioDetail *PortfolioDetail `type:"structure"`

	// Information about the TagOptions associated with the portfolio.
	TagOptions []TagOptionDetail `type:"list"`

	// Information about the tags associated with the portfolio.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s DescribePortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePortfolio = "DescribePortfolio"

// DescribePortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified portfolio.
//
//    // Example sending a request using DescribePortfolioRequest.
//    req := client.DescribePortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribePortfolio
func (c *Client) DescribePortfolioRequest(input *DescribePortfolioInput) DescribePortfolioRequest {
	op := &aws.Operation{
		Name:       opDescribePortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePortfolioInput{}
	}

	req := c.newRequest(op, input, &DescribePortfolioOutput{})
	return DescribePortfolioRequest{Request: req, Input: input, Copy: c.DescribePortfolioRequest}
}

// DescribePortfolioRequest is the request type for the
// DescribePortfolio API operation.
type DescribePortfolioRequest struct {
	*aws.Request
	Input *DescribePortfolioInput
	Copy  func(*DescribePortfolioInput) DescribePortfolioRequest
}

// Send marshals and sends the DescribePortfolio API request.
func (r DescribePortfolioRequest) Send(ctx context.Context) (*DescribePortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePortfolioResponse{
		DescribePortfolioOutput: r.Request.Data.(*DescribePortfolioOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePortfolioResponse is the response type for the
// DescribePortfolio API operation.
type DescribePortfolioResponse struct {
	*DescribePortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePortfolio request.
func (r *DescribePortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
