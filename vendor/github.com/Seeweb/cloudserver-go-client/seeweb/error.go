package seeweb

import (
	"errors"
	"fmt"
)

var (
	// ErrNoToken is returned by NewClient if a user
	// passed an empty/missing token.
	ErrNoToken = errors.New("an empty token was provided")

	// ErrAuthFailure is returned by NewClient if a user
	// passed an invalid token and failed validation against the Seeweb API.
	ErrAuthFailure = errors.New("failed to authenticate using the provided Seeweb token")
)

// Error represents an error response from the Seeweb API.
type Error struct {
	ErrorResponse *Response
	ErrorCode     int    `json:"error_code,omitempty"`
	Status        string `json:"status,omitempty"`
	Message       string `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s API call to %s failed %v. Code: %d, Status: %s, Message: %s", e.ErrorResponse.Response.Request.Method, e.ErrorResponse.Response.Request.URL.String(), e.ErrorResponse.Response.Status, e.ErrorCode, e.Status, e.Message)
}
