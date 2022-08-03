package seeweb

import (
	"fmt"
	"time"
)

// ActionService handles the communication with action
// related methods of the Seeweb API.
type ActionService service

type Action struct {
	ID           float64   `json:"id"`
	Status       string    `json:"status"`
	User         string    `json:"user"`
	CreatedAt    time.Time `json:"created_at"`
	StartedAt    time.Time `json:"started_at"`
	CompletedAt  time.Time `json:"completed_at"`
	Resource     string    `json:"resource"`
	ResourceType string    `json:"resource_type"`
	Type         string    `json:"type"`
	Progress     int       `json:"progress"`
}

type SeewebActionListResponse struct {
	Status  string    `json:"status"`
	Actions []*Action `json:"actions"`
}

type SeewebActionGetResponse struct {
	Status string  `json:"status"`
	Action *Action `json:"action"`
}

// Get retrieves information about an action.
func (a *ActionService) Get(id string) (*SeewebActionGetResponse, *Response, error) {
	u := fmt.Sprintf("/actions/%s", id)
	v := new(SeewebActionGetResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// List lists all existing servers.
func (a *ActionService) List() (*SeewebActionListResponse, *Response, error) {
	u := "/actions"
	v := new(SeewebActionListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
