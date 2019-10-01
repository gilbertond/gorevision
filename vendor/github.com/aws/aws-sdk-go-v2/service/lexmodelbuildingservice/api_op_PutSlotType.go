// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutSlotTypeInput struct {
	_ struct{} `type:"structure"`

	// Identifies a specific revision of the $LATEST version.
	//
	// When you create a new slot type, leave the checksum field blank. If you specify
	// a checksum you get a BadRequestException exception.
	//
	// When you want to update a slot type, set the checksum field to the checksum
	// of the most recent revision of the $LATEST version. If you don't specify
	// the checksum field, or if the checksum does not match the $LATEST version,
	// you get a PreconditionFailedException exception.
	Checksum *string `locationName:"checksum" type:"string"`

	CreateVersion *bool `locationName:"createVersion" type:"boolean"`

	// A description of the slot type.
	Description *string `locationName:"description" type:"string"`

	// A list of EnumerationValue objects that defines the values that the slot
	// type can take. Each value can have a list of synonyms, which are additional
	// values that help train the machine learning model about the values that it
	// resolves for a slot.
	//
	// When Amazon Lex resolves a slot value, it generates a resolution list that
	// contains up to five possible values for the slot. If you are using a Lambda
	// function, this resolution list is passed to the function. If you are not
	// using a Lambda function you can choose to return the value that the user
	// entered or the first value in the resolution list as the slot value. The
	// valueSelectionStrategy field indicates the option to use.
	EnumerationValues []EnumerationValue `locationName:"enumerationValues" min:"1" type:"list"`

	// The name of the slot type. The name is not case sensitive.
	//
	// The name can't match a built-in slot type name, or a built-in slot type name
	// with "AMAZON." removed. For example, because there is a built-in slot type
	// called AMAZON.DATE, you can't create a custom slot type called DATE.
	//
	// For a list of built-in slot types, see Slot Type Reference (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference)
	// in the Alexa Skills Kit.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// Determines the slot resolution strategy that Amazon Lex uses to return slot
	// type values. The field can be set to one of the following values:
	//
	//    * ORIGINAL_VALUE - Returns the value entered by the user, if the user
	//    value is similar to the slot value.
	//
	//    * TOP_RESOLUTION - If there is a resolution list for the slot, return
	//    the first value in the resolution list as the slot type value. If there
	//    is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is ORIGINAL_VALUE.
	ValueSelectionStrategy SlotValueSelectionStrategy `locationName:"valueSelectionStrategy" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutSlotTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutSlotTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutSlotTypeInput"}
	if s.EnumerationValues != nil && len(s.EnumerationValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnumerationValues", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.EnumerationValues != nil {
		for i, v := range s.EnumerationValues {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EnumerationValues", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutSlotTypeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "checksum", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreateVersion != nil {
		v := *s.CreateVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createVersion", protocol.BoolValue(v), metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnumerationValues != nil {
		v := s.EnumerationValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "enumerationValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ValueSelectionStrategy) > 0 {
		v := s.ValueSelectionStrategy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "valueSelectionStrategy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutSlotTypeOutput struct {
	_ struct{} `type:"structure"`

	// Checksum of the $LATEST version of the slot type.
	Checksum *string `locationName:"checksum" type:"string"`

	CreateVersion *bool `locationName:"createVersion" type:"boolean"`

	// The date that the slot type was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// A description of the slot type.
	Description *string `locationName:"description" type:"string"`

	// A list of EnumerationValue objects that defines the values that the slot
	// type can take.
	EnumerationValues []EnumerationValue `locationName:"enumerationValues" min:"1" type:"list"`

	// The date that the slot type was updated. When you create a slot type, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp"`

	// The name of the slot type.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The slot resolution strategy that Amazon Lex uses to determine the value
	// of the slot. For more information, see PutSlotType.
	ValueSelectionStrategy SlotValueSelectionStrategy `locationName:"valueSelectionStrategy" type:"string" enum:"true"`

	// The version of the slot type. For a new slot type, the version is always
	// $LATEST.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s PutSlotTypeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutSlotTypeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "checksum", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreateVersion != nil {
		v := *s.CreateVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createVersion", protocol.BoolValue(v), metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnumerationValues != nil {
		v := s.EnumerationValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "enumerationValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ValueSelectionStrategy) > 0 {
		v := s.ValueSelectionStrategy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "valueSelectionStrategy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opPutSlotType = "PutSlotType"

// PutSlotTypeRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Creates a custom slot type or replaces an existing custom slot type.
//
// To create a custom slot type, specify a name for the slot type and a set
// of enumeration values, which are the values that a slot of this type can
// assume. For more information, see how-it-works.
//
// If you specify the name of an existing slot type, the fields in the request
// replace the existing values in the $LATEST version of the slot type. Amazon
// Lex removes the fields that you don't provide in the request. If you don't
// specify required fields, Amazon Lex throws an exception. When you update
// the $LATEST version of a slot type, if a bot uses the $LATEST version of
// an intent that contains the slot type, the bot's status field is set to NOT_BUILT.
//
// This operation requires permissions for the lex:PutSlotType action.
//
//    // Example sending a request using PutSlotTypeRequest.
//    req := client.PutSlotTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/PutSlotType
func (c *Client) PutSlotTypeRequest(input *PutSlotTypeInput) PutSlotTypeRequest {
	op := &aws.Operation{
		Name:       opPutSlotType,
		HTTPMethod: "PUT",
		HTTPPath:   "/slottypes/{name}/versions/$LATEST",
	}

	if input == nil {
		input = &PutSlotTypeInput{}
	}

	req := c.newRequest(op, input, &PutSlotTypeOutput{})
	return PutSlotTypeRequest{Request: req, Input: input, Copy: c.PutSlotTypeRequest}
}

// PutSlotTypeRequest is the request type for the
// PutSlotType API operation.
type PutSlotTypeRequest struct {
	*aws.Request
	Input *PutSlotTypeInput
	Copy  func(*PutSlotTypeInput) PutSlotTypeRequest
}

// Send marshals and sends the PutSlotType API request.
func (r PutSlotTypeRequest) Send(ctx context.Context) (*PutSlotTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutSlotTypeResponse{
		PutSlotTypeOutput: r.Request.Data.(*PutSlotTypeOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutSlotTypeResponse is the response type for the
// PutSlotType API operation.
type PutSlotTypeResponse struct {
	*PutSlotTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutSlotType request.
func (r *PutSlotTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
