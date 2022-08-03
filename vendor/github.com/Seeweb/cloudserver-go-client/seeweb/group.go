package seeweb

import (
	"fmt"
	"strconv"
)

// GroupService handles the communication with group
// related methods of the Seeweb API.
type GroupService service

type Group struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Notes   string `json:"notes"`
	Enabled bool   `json:"enabled"`
}
type SeewebGroupCreateRequest struct {
	Notes    string `json:"notes"`
	Password string `json:"password"`
}
type SeewebGroupCreateResponse struct {
	Status string `json:"status"`
	Group  *Group `json:"group"`
}

type SeewebGroupListResponse struct {
	Status string   `json:"status"`
	Groups []*Group `json:"groups"`
}

type SeewebGroupDeleteResponse struct {
	Status string `json:"status"`
}

// List lists all existing groups.
func (a *GroupService) List() (*SeewebGroupListResponse, *Response, error) {
	u := "/groups"
	v := new(SeewebGroupListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new group.
func (s *GroupService) Create(createGroupRequest *SeewebGroupCreateRequest) (*SeewebGroupCreateResponse, *Response, error) {
	u := "/groups"
	v := new(SeewebGroupCreateResponse)

	resp, err := s.client.newRequestDo("POST", u, &createGroupRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete removes an existing group.
func (s *GroupService) Delete(id int) (*SeewebGroupDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/groups/%s", strconv.Itoa(id))
	v := new(SeewebGroupDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
