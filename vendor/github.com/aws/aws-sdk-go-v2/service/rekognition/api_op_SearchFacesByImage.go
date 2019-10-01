// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchFacesByImageInput struct {
	_ struct{} `type:"structure"`

	// ID of the collection to search.
	//
	// CollectionId is a required field
	CollectionId *string `min:"1" type:"string" required:"true"`

	// (Optional) Specifies the minimum confidence in the face match to return.
	// For example, don't return any matches where confidence in matches is less
	// than 70%. The default value is 80%.
	FaceMatchThreshold *float64 `type:"float"`

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS
	// CLI to call Amazon Rekognition operations, passing base64-encoded image bytes
	// is not supported.
	//
	// If you are using an AWS SDK to call Amazon Rekognition, you might not need
	// to base64-encode image bytes passed using the Bytes field. For more information,
	// see Images in the Amazon Rekognition developer guide.
	//
	// Image is a required field
	Image *Image `type:"structure" required:"true"`

	// Maximum number of faces to return. The operation returns the maximum number
	// of faces with the highest confidence in the match.
	MaxFaces *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s SearchFacesByImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchFacesByImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchFacesByImageInput"}

	if s.CollectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CollectionId"))
	}
	if s.CollectionId != nil && len(*s.CollectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CollectionId", 1))
	}

	if s.Image == nil {
		invalidParams.Add(aws.NewErrParamRequired("Image"))
	}
	if s.MaxFaces != nil && *s.MaxFaces < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxFaces", 1))
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			invalidParams.AddNested("Image", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchFacesByImageOutput struct {
	_ struct{} `type:"structure"`

	// An array of faces that match the input face, along with the confidence in
	// the match.
	FaceMatches []FaceMatch `type:"list"`

	// Version number of the face detection model associated with the input collection
	// (CollectionId).
	FaceModelVersion *string `type:"string"`

	// The bounding box around the face in the input image that Amazon Rekognition
	// used for the search.
	SearchedFaceBoundingBox *BoundingBox `type:"structure"`

	// The level of confidence that the searchedFaceBoundingBox, contains a face.
	SearchedFaceConfidence *float64 `type:"float"`
}

// String returns the string representation
func (s SearchFacesByImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchFacesByImage = "SearchFacesByImage"

// SearchFacesByImageRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// For a given input image, first detects the largest face in the image, and
// then searches the specified collection for matching faces. The operation
// compares the features of the input face with faces in the specified collection.
//
// To search for all faces in an input image, you might first call the IndexFaces
// operation, and then use the face IDs returned in subsequent calls to the
// SearchFaces operation.
//
// You can also call the DetectFaces operation and use the bounding boxes in
// the response to make face crops, which then you can pass in to the SearchFacesByImage
// operation.
//
// You pass the input image either as base64-encoded image bytes or as a reference
// to an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must
// be either a PNG or JPEG formatted file.
//
// The response returns an array of faces that match, ordered by similarity
// score with the highest similarity first. More specifically, it is an array
// of metadata for each face match found. Along with the metadata, the response
// also includes a similarity indicating how similar the face is to the input
// face. In the response, the operation also returns the bounding box (and a
// confidence level that the bounding box contains a face) of the face that
// Amazon Rekognition used for the input image.
//
// For an example, Searching for a Face Using an Image in the Amazon Rekognition
// Developer Guide.
//
// This operation requires permissions to perform the rekognition:SearchFacesByImage
// action.
//
//    // Example sending a request using SearchFacesByImageRequest.
//    req := client.SearchFacesByImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SearchFacesByImageRequest(input *SearchFacesByImageInput) SearchFacesByImageRequest {
	op := &aws.Operation{
		Name:       opSearchFacesByImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SearchFacesByImageInput{}
	}

	req := c.newRequest(op, input, &SearchFacesByImageOutput{})
	return SearchFacesByImageRequest{Request: req, Input: input, Copy: c.SearchFacesByImageRequest}
}

// SearchFacesByImageRequest is the request type for the
// SearchFacesByImage API operation.
type SearchFacesByImageRequest struct {
	*aws.Request
	Input *SearchFacesByImageInput
	Copy  func(*SearchFacesByImageInput) SearchFacesByImageRequest
}

// Send marshals and sends the SearchFacesByImage API request.
func (r SearchFacesByImageRequest) Send(ctx context.Context) (*SearchFacesByImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchFacesByImageResponse{
		SearchFacesByImageOutput: r.Request.Data.(*SearchFacesByImageOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SearchFacesByImageResponse is the response type for the
// SearchFacesByImage API operation.
type SearchFacesByImageResponse struct {
	*SearchFacesByImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchFacesByImage request.
func (r *SearchFacesByImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
