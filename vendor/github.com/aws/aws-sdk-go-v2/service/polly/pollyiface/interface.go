// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pollyiface provides an interface to enable mocking the Amazon Polly service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pollyiface

import (
	"github.com/aws/aws-sdk-go-v2/service/polly"
)

// ClientAPI provides an interface to enable mocking the
// polly.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Polly.
//    func myFunc(svc pollyiface.ClientAPI) bool {
//        // Make svc.DeleteLexicon request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := polly.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        pollyiface.ClientPI
//    }
//    func (m *mockClientClient) DeleteLexicon(input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error) {
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
	DeleteLexiconRequest(*polly.DeleteLexiconInput) polly.DeleteLexiconRequest

	DescribeVoicesRequest(*polly.DescribeVoicesInput) polly.DescribeVoicesRequest

	GetLexiconRequest(*polly.GetLexiconInput) polly.GetLexiconRequest

	GetSpeechSynthesisTaskRequest(*polly.GetSpeechSynthesisTaskInput) polly.GetSpeechSynthesisTaskRequest

	ListLexiconsRequest(*polly.ListLexiconsInput) polly.ListLexiconsRequest

	ListSpeechSynthesisTasksRequest(*polly.ListSpeechSynthesisTasksInput) polly.ListSpeechSynthesisTasksRequest

	PutLexiconRequest(*polly.PutLexiconInput) polly.PutLexiconRequest

	StartSpeechSynthesisTaskRequest(*polly.StartSpeechSynthesisTaskInput) polly.StartSpeechSynthesisTaskRequest

	SynthesizeSpeechRequest(*polly.SynthesizeSpeechInput) polly.SynthesizeSpeechRequest
}

var _ ClientAPI = (*polly.Client)(nil)
