package seeweb

import (
	"fmt"
	"time"
)

// TemplateService handles the communication with template
// related methods of the Seeweb API.
type TemplateService service

type Template struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
	ActiveFlag   bool      `json:"active_flag"`
	Status       string    `json:"status"`
	UUID         string    `json:"uuid"`
	Notes        string    `json:"notes"`
}

type SeewebTemplateListResponse struct {
	Status    string      `json:"status"`
	Templates []*Template `json:"templates"`
}

type SeewebTemplateGetResponse struct {
	Status   string    `json:"status"`
	Template *Template `json:"template"`
}

// Get retrieves information about an template.
func (a *TemplateService) Get(id string) (*SeewebTemplateGetResponse, *Response, error) {
	u := fmt.Sprintf("/templates/%s", id)
	v := new(SeewebTemplateGetResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// List lists all existing plans.
func (a *TemplateService) List() (*SeewebTemplateListResponse, *Response, error) {
	u := "/templates"
	v := new(SeewebTemplateListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
