// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type CreatePlayerSessionsInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the game session to add players to.
	//
	// GameSessionId is a required field
	GameSessionId *string `min:"1" type:"string" required:"true"`

	// Map of string pairs, each specifying a player ID and a set of developer-defined
	// information related to the player. Amazon GameLift does not use this data,
	// so it can be formatted as needed for use in the game. Player data strings
	// for player IDs not included in the PlayerIds parameter are ignored.
	PlayerDataMap map[string]string `type:"map"`

	// List of unique identifiers for the players to be added.
	//
	// PlayerIds is a required field
	PlayerIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreatePlayerSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePlayerSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePlayerSessionsInput"}

	if s.GameSessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionId"))
	}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}

	if s.PlayerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlayerIds"))
	}
	if s.PlayerIds != nil && len(s.PlayerIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlayerIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type CreatePlayerSessionsOutput struct {
	_ struct{} `type:"structure"`

	// Collection of player session objects created for the added players.
	PlayerSessions []PlayerSession `type:"list"`
}

// String returns the string representation
func (s CreatePlayerSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePlayerSessions = "CreatePlayerSessions"

// CreatePlayerSessionsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Reserves open slots in a game session for a group of players. Before players
// can be added, a game session must have an ACTIVE status, have a creation
// policy of ALLOW_ALL, and have an open player slot. To add a single player
// to a game session, use CreatePlayerSession. When a player connects to the
// game server and references a player session ID, the game server contacts
// the Amazon GameLift service to validate the player reservation and accept
// the player.
//
// To create player sessions, specify a game session ID, a list of player IDs,
// and optionally a set of player data strings. If successful, a slot is reserved
// in the game session for each player and a set of new PlayerSession objects
// is returned. Player sessions cannot be updated.
//
// Available in Amazon GameLift Local.
//
//    * CreatePlayerSession
//
//    * CreatePlayerSessions
//
//    * DescribePlayerSessions
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using CreatePlayerSessionsRequest.
//    req := client.CreatePlayerSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreatePlayerSessions
func (c *Client) CreatePlayerSessionsRequest(input *CreatePlayerSessionsInput) CreatePlayerSessionsRequest {
	op := &aws.Operation{
		Name:       opCreatePlayerSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePlayerSessionsInput{}
	}

	req := c.newRequest(op, input, &CreatePlayerSessionsOutput{})
	return CreatePlayerSessionsRequest{Request: req, Input: input, Copy: c.CreatePlayerSessionsRequest}
}

// CreatePlayerSessionsRequest is the request type for the
// CreatePlayerSessions API operation.
type CreatePlayerSessionsRequest struct {
	*aws.Request
	Input *CreatePlayerSessionsInput
	Copy  func(*CreatePlayerSessionsInput) CreatePlayerSessionsRequest
}

// Send marshals and sends the CreatePlayerSessions API request.
func (r CreatePlayerSessionsRequest) Send(ctx context.Context) (*CreatePlayerSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlayerSessionsResponse{
		CreatePlayerSessionsOutput: r.Request.Data.(*CreatePlayerSessionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlayerSessionsResponse is the response type for the
// CreatePlayerSessions API operation.
type CreatePlayerSessionsResponse struct {
	*CreatePlayerSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlayerSessions request.
func (r *CreatePlayerSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
