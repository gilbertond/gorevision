// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotthingsgraphiface provides an interface to enable mocking the AWS IoT Things Graph service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotthingsgraphiface

import (
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
)

// ClientAPI provides an interface to enable mocking the
// iotthingsgraph.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Things Graph.
//    func myFunc(svc iotthingsgraphiface.ClientAPI) bool {
//        // Make svc.AssociateEntityToThing request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := iotthingsgraph.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        iotthingsgraphiface.ClientPI
//    }
//    func (m *mockClientClient) AssociateEntityToThing(input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
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
	AssociateEntityToThingRequest(*iotthingsgraph.AssociateEntityToThingInput) iotthingsgraph.AssociateEntityToThingRequest

	CreateFlowTemplateRequest(*iotthingsgraph.CreateFlowTemplateInput) iotthingsgraph.CreateFlowTemplateRequest

	CreateSystemInstanceRequest(*iotthingsgraph.CreateSystemInstanceInput) iotthingsgraph.CreateSystemInstanceRequest

	CreateSystemTemplateRequest(*iotthingsgraph.CreateSystemTemplateInput) iotthingsgraph.CreateSystemTemplateRequest

	DeleteFlowTemplateRequest(*iotthingsgraph.DeleteFlowTemplateInput) iotthingsgraph.DeleteFlowTemplateRequest

	DeleteNamespaceRequest(*iotthingsgraph.DeleteNamespaceInput) iotthingsgraph.DeleteNamespaceRequest

	DeleteSystemInstanceRequest(*iotthingsgraph.DeleteSystemInstanceInput) iotthingsgraph.DeleteSystemInstanceRequest

	DeleteSystemTemplateRequest(*iotthingsgraph.DeleteSystemTemplateInput) iotthingsgraph.DeleteSystemTemplateRequest

	DeploySystemInstanceRequest(*iotthingsgraph.DeploySystemInstanceInput) iotthingsgraph.DeploySystemInstanceRequest

	DeprecateFlowTemplateRequest(*iotthingsgraph.DeprecateFlowTemplateInput) iotthingsgraph.DeprecateFlowTemplateRequest

	DeprecateSystemTemplateRequest(*iotthingsgraph.DeprecateSystemTemplateInput) iotthingsgraph.DeprecateSystemTemplateRequest

	DescribeNamespaceRequest(*iotthingsgraph.DescribeNamespaceInput) iotthingsgraph.DescribeNamespaceRequest

	DissociateEntityFromThingRequest(*iotthingsgraph.DissociateEntityFromThingInput) iotthingsgraph.DissociateEntityFromThingRequest

	GetEntitiesRequest(*iotthingsgraph.GetEntitiesInput) iotthingsgraph.GetEntitiesRequest

	GetFlowTemplateRequest(*iotthingsgraph.GetFlowTemplateInput) iotthingsgraph.GetFlowTemplateRequest

	GetFlowTemplateRevisionsRequest(*iotthingsgraph.GetFlowTemplateRevisionsInput) iotthingsgraph.GetFlowTemplateRevisionsRequest

	GetNamespaceDeletionStatusRequest(*iotthingsgraph.GetNamespaceDeletionStatusInput) iotthingsgraph.GetNamespaceDeletionStatusRequest

	GetSystemInstanceRequest(*iotthingsgraph.GetSystemInstanceInput) iotthingsgraph.GetSystemInstanceRequest

	GetSystemTemplateRequest(*iotthingsgraph.GetSystemTemplateInput) iotthingsgraph.GetSystemTemplateRequest

	GetSystemTemplateRevisionsRequest(*iotthingsgraph.GetSystemTemplateRevisionsInput) iotthingsgraph.GetSystemTemplateRevisionsRequest

	GetUploadStatusRequest(*iotthingsgraph.GetUploadStatusInput) iotthingsgraph.GetUploadStatusRequest

	ListFlowExecutionMessagesRequest(*iotthingsgraph.ListFlowExecutionMessagesInput) iotthingsgraph.ListFlowExecutionMessagesRequest

	ListTagsForResourceRequest(*iotthingsgraph.ListTagsForResourceInput) iotthingsgraph.ListTagsForResourceRequest

	SearchEntitiesRequest(*iotthingsgraph.SearchEntitiesInput) iotthingsgraph.SearchEntitiesRequest

	SearchFlowExecutionsRequest(*iotthingsgraph.SearchFlowExecutionsInput) iotthingsgraph.SearchFlowExecutionsRequest

	SearchFlowTemplatesRequest(*iotthingsgraph.SearchFlowTemplatesInput) iotthingsgraph.SearchFlowTemplatesRequest

	SearchSystemInstancesRequest(*iotthingsgraph.SearchSystemInstancesInput) iotthingsgraph.SearchSystemInstancesRequest

	SearchSystemTemplatesRequest(*iotthingsgraph.SearchSystemTemplatesInput) iotthingsgraph.SearchSystemTemplatesRequest

	SearchThingsRequest(*iotthingsgraph.SearchThingsInput) iotthingsgraph.SearchThingsRequest

	TagResourceRequest(*iotthingsgraph.TagResourceInput) iotthingsgraph.TagResourceRequest

	UndeploySystemInstanceRequest(*iotthingsgraph.UndeploySystemInstanceInput) iotthingsgraph.UndeploySystemInstanceRequest

	UntagResourceRequest(*iotthingsgraph.UntagResourceInput) iotthingsgraph.UntagResourceRequest

	UpdateFlowTemplateRequest(*iotthingsgraph.UpdateFlowTemplateInput) iotthingsgraph.UpdateFlowTemplateRequest

	UpdateSystemTemplateRequest(*iotthingsgraph.UpdateSystemTemplateInput) iotthingsgraph.UpdateSystemTemplateRequest

	UploadEntityDefinitionsRequest(*iotthingsgraph.UploadEntityDefinitionsInput) iotthingsgraph.UploadEntityDefinitionsRequest
}

var _ ClientAPI = (*iotthingsgraph.Client)(nil)
