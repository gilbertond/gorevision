// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudtrailiface provides an interface to enable mocking the AWS CloudTrail service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudtrailiface

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

// ClientAPI provides an interface to enable mocking the
// cloudtrail.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // CloudTrail.
//    func myFunc(svc cloudtrailiface.ClientAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudtrail.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cloudtrailiface.ClientPI
//    }
//    func (m *mockClientClient) AddTags(input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AddTagsRequest(*cloudtrail.AddTagsInput) cloudtrail.AddTagsRequest

	CreateTrailRequest(*cloudtrail.CreateTrailInput) cloudtrail.CreateTrailRequest

	DeleteTrailRequest(*cloudtrail.DeleteTrailInput) cloudtrail.DeleteTrailRequest

	DescribeTrailsRequest(*cloudtrail.DescribeTrailsInput) cloudtrail.DescribeTrailsRequest

	GetEventSelectorsRequest(*cloudtrail.GetEventSelectorsInput) cloudtrail.GetEventSelectorsRequest

	GetTrailStatusRequest(*cloudtrail.GetTrailStatusInput) cloudtrail.GetTrailStatusRequest

	ListPublicKeysRequest(*cloudtrail.ListPublicKeysInput) cloudtrail.ListPublicKeysRequest

	ListTagsRequest(*cloudtrail.ListTagsInput) cloudtrail.ListTagsRequest

	LookupEventsRequest(*cloudtrail.LookupEventsInput) cloudtrail.LookupEventsRequest

	PutEventSelectorsRequest(*cloudtrail.PutEventSelectorsInput) cloudtrail.PutEventSelectorsRequest

	RemoveTagsRequest(*cloudtrail.RemoveTagsInput) cloudtrail.RemoveTagsRequest

	StartLoggingRequest(*cloudtrail.StartLoggingInput) cloudtrail.StartLoggingRequest

	StopLoggingRequest(*cloudtrail.StopLoggingInput) cloudtrail.StopLoggingRequest

	UpdateTrailRequest(*cloudtrail.UpdateTrailInput) cloudtrail.UpdateTrailRequest
}

var _ ClientAPI = (*cloudtrail.Client)(nil)
