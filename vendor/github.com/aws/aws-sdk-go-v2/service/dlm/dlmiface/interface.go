// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dlmiface provides an interface to enable mocking the Amazon Data Lifecycle Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dlmiface

import (
	"github.com/aws/aws-sdk-go-v2/service/dlm"
)

// ClientAPI provides an interface to enable mocking the
// dlm.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon DLM.
//    func myFunc(svc dlmiface.ClientAPI) bool {
//        // Make svc.CreateLifecyclePolicy request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := dlm.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        dlmiface.ClientPI
//    }
//    func (m *mockClientClient) CreateLifecyclePolicy(input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
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
	CreateLifecyclePolicyRequest(*dlm.CreateLifecyclePolicyInput) dlm.CreateLifecyclePolicyRequest

	DeleteLifecyclePolicyRequest(*dlm.DeleteLifecyclePolicyInput) dlm.DeleteLifecyclePolicyRequest

	GetLifecyclePoliciesRequest(*dlm.GetLifecyclePoliciesInput) dlm.GetLifecyclePoliciesRequest

	GetLifecyclePolicyRequest(*dlm.GetLifecyclePolicyInput) dlm.GetLifecyclePolicyRequest

	UpdateLifecyclePolicyRequest(*dlm.UpdateLifecyclePolicyInput) dlm.UpdateLifecyclePolicyRequest
}

var _ ClientAPI = (*dlm.Client)(nil)
