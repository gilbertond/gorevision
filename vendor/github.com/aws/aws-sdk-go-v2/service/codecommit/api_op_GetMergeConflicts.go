// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMergeConflictsInput struct {
	_ struct{} `type:"structure"`

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL
	// is used, which will return a not mergeable result if the same file has differences
	// in both branches. If LINE_LEVEL is specified, a conflict will be considered
	// not mergeable if the same file in both branches has differences on the same
	// line.
	ConflictDetailLevel ConflictDetailLevelTypeEnum `locationName:"conflictDetailLevel" type:"string" enum:"true"`

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation
	// will be successful.
	ConflictResolutionStrategy ConflictResolutionStrategyTypeEnum `locationName:"conflictResolutionStrategy" type:"string" enum:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit. For example, a branch name or a full commit ID.
	//
	// DestinationCommitSpecifier is a required field
	DestinationCommitSpecifier *string `locationName:"destinationCommitSpecifier" type:"string" required:"true"`

	// The maximum number of files to include in the output.
	MaxConflictFiles *int64 `locationName:"maxConflictFiles" type:"integer"`

	// The merge option or strategy you want to use to merge the code.
	//
	// MergeOption is a required field
	MergeOption MergeOptionTypeEnum `locationName:"mergeOption" type:"string" required:"true" enum:"true"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository where the pull request was created.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit. For example, a branch name or a full commit ID.
	//
	// SourceCommitSpecifier is a required field
	SourceCommitSpecifier *string `locationName:"sourceCommitSpecifier" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMergeConflictsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMergeConflictsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMergeConflictsInput"}

	if s.DestinationCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCommitSpecifier"))
	}
	if len(s.MergeOption) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MergeOption"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if s.SourceCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceCommitSpecifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMergeConflictsOutput struct {
	_ struct{} `type:"structure"`

	// The commit ID of the merge base.
	BaseCommitId *string `locationName:"baseCommitId" type:"string"`

	// A list of metadata for any conflicting files. If the specified merge strategy
	// is FAST_FORWARD_MERGE, this list will always be empty.
	//
	// ConflictMetadataList is a required field
	ConflictMetadataList []ConflictMetadata `locationName:"conflictMetadataList" type:"list" required:"true"`

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// DestinationCommitId is a required field
	DestinationCommitId *string `locationName:"destinationCommitId" type:"string" required:"true"`

	// A Boolean value that indicates whether the code is mergeable by the specified
	// merge option.
	//
	// Mergeable is a required field
	Mergeable *bool `locationName:"mergeable" type:"boolean" required:"true"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The commit ID of the source commit specifier that was used in the merge evaluation.
	//
	// SourceCommitId is a required field
	SourceCommitId *string `locationName:"sourceCommitId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMergeConflictsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMergeConflicts = "GetMergeConflicts"

// GetMergeConflictsRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about merge conflicts between the before and after commit
// IDs for a pull request in a repository.
//
//    // Example sending a request using GetMergeConflictsRequest.
//    req := client.GetMergeConflictsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetMergeConflicts
func (c *Client) GetMergeConflictsRequest(input *GetMergeConflictsInput) GetMergeConflictsRequest {
	op := &aws.Operation{
		Name:       opGetMergeConflicts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxConflictFiles",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetMergeConflictsInput{}
	}

	req := c.newRequest(op, input, &GetMergeConflictsOutput{})
	return GetMergeConflictsRequest{Request: req, Input: input, Copy: c.GetMergeConflictsRequest}
}

// GetMergeConflictsRequest is the request type for the
// GetMergeConflicts API operation.
type GetMergeConflictsRequest struct {
	*aws.Request
	Input *GetMergeConflictsInput
	Copy  func(*GetMergeConflictsInput) GetMergeConflictsRequest
}

// Send marshals and sends the GetMergeConflicts API request.
func (r GetMergeConflictsRequest) Send(ctx context.Context) (*GetMergeConflictsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMergeConflictsResponse{
		GetMergeConflictsOutput: r.Request.Data.(*GetMergeConflictsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetMergeConflictsRequestPaginator returns a paginator for GetMergeConflicts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetMergeConflictsRequest(input)
//   p := codecommit.NewGetMergeConflictsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetMergeConflictsPaginator(req GetMergeConflictsRequest) GetMergeConflictsPaginator {
	return GetMergeConflictsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetMergeConflictsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetMergeConflictsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetMergeConflictsPaginator struct {
	aws.Pager
}

func (p *GetMergeConflictsPaginator) CurrentPage() *GetMergeConflictsOutput {
	return p.Pager.CurrentPage().(*GetMergeConflictsOutput)
}

// GetMergeConflictsResponse is the response type for the
// GetMergeConflicts API operation.
type GetMergeConflictsResponse struct {
	*GetMergeConflictsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMergeConflicts request.
func (r *GetMergeConflictsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
