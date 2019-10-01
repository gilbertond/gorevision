// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetPatchBaselineInput struct {
	_ struct{} `type:"structure"`

	// The ID of the patch baseline to retrieve.
	//
	// BaselineId is a required field
	BaselineId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPatchBaselineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPatchBaselineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPatchBaselineInput"}

	if s.BaselineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BaselineId"))
	}
	if s.BaselineId != nil && len(*s.BaselineId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("BaselineId", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetPatchBaselineOutput struct {
	_ struct{} `type:"structure"`

	// A set of rules used to include patches in the baseline.
	ApprovalRules *PatchRuleGroup `type:"structure"`

	// A list of explicitly approved patches for the baseline.
	ApprovedPatches []string `type:"list"`

	// Returns the specified compliance severity level for approved patches in the
	// patch baseline.
	ApprovedPatchesComplianceLevel PatchComplianceLevel `type:"string" enum:"true"`

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the instances. The default value is 'false'. Applies
	// to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `type:"boolean"`

	// The ID of the retrieved patch baseline.
	BaselineId *string `min:"20" type:"string"`

	// The date the patch baseline was created.
	CreatedDate *time.Time `type:"timestamp"`

	// A description of the patch baseline.
	Description *string `min:"1" type:"string"`

	// A set of global filters used to exclude patches from the baseline.
	GlobalFilters *PatchFilterGroup `type:"structure"`

	// The date the patch baseline was last modified.
	ModifiedDate *time.Time `type:"timestamp"`

	// The name of the patch baseline.
	Name *string `min:"3" type:"string"`

	// Returns the operating system specified for the patch baseline.
	OperatingSystem OperatingSystem `type:"string" enum:"true"`

	// Patch groups included in the patch baseline.
	PatchGroups []string `type:"list"`

	// A list of explicitly rejected patches for the baseline.
	RejectedPatches []string `type:"list"`

	// The action specified to take on patches included in the RejectedPatches list.
	// A patch can be allowed only if it is a dependency of another package, or
	// blocked entirely along with packages that include it as a dependency.
	RejectedPatchesAction PatchAction `type:"string" enum:"true"`

	// Information about the patches to use to update the instances, including target
	// operating systems and source repositories. Applies to Linux instances only.
	Sources []PatchSource `type:"list"`
}

// String returns the string representation
func (s GetPatchBaselineOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPatchBaseline = "GetPatchBaseline"

// GetPatchBaselineRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves information about a patch baseline.
//
//    // Example sending a request using GetPatchBaselineRequest.
//    req := client.GetPatchBaselineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetPatchBaseline
func (c *Client) GetPatchBaselineRequest(input *GetPatchBaselineInput) GetPatchBaselineRequest {
	op := &aws.Operation{
		Name:       opGetPatchBaseline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPatchBaselineInput{}
	}

	req := c.newRequest(op, input, &GetPatchBaselineOutput{})
	return GetPatchBaselineRequest{Request: req, Input: input, Copy: c.GetPatchBaselineRequest}
}

// GetPatchBaselineRequest is the request type for the
// GetPatchBaseline API operation.
type GetPatchBaselineRequest struct {
	*aws.Request
	Input *GetPatchBaselineInput
	Copy  func(*GetPatchBaselineInput) GetPatchBaselineRequest
}

// Send marshals and sends the GetPatchBaseline API request.
func (r GetPatchBaselineRequest) Send(ctx context.Context) (*GetPatchBaselineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPatchBaselineResponse{
		GetPatchBaselineOutput: r.Request.Data.(*GetPatchBaselineOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPatchBaselineResponse is the response type for the
// GetPatchBaseline API operation.
type GetPatchBaselineResponse struct {
	*GetPatchBaselineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPatchBaseline request.
func (r *GetPatchBaselineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
