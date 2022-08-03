package seeweb

// RegionService handles the communication with region
// related methods of the Seeweb API.
type RegionService service

type Region struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	Description string `json:"description"`
}
type SeewebRegionListResponse struct {
	Status  string    `json:"status"`
	Regions []*Region `json:"regions"`
}

// List lists all existing regions.
func (a *RegionService) List() (*SeewebRegionListResponse, *Response, error) {
	u := "/regions"
	v := new(SeewebRegionListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
