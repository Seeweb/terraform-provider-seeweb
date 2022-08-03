package seeweb

// PlanService handles the communication with plan
// related methods of the Seeweb API.
type PlanService service

type AvailableRegions struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	Description string `json:"description"`
}
type Plan struct {
	ID               int                 `json:"id"`
	Name             string              `json:"name"`
	CPU              string              `json:"cpu"`
	RAM              string              `json:"ram"`
	Disk             string              `json:"disk"`
	HourlyPrice      float64             `json:"hourly_price"`
	MontlyPrice      float64             `json:"montly_price"`
	Windows          bool                `json:"windows"`
	Available        bool                `json:"available"`
	AvailableRegions []*AvailableRegions `json:"available_regions"`
}
type SeewebPlanListResponse struct {
	Status string  `json:"status"`
	Plans  []*Plan `json:"plans"`
}

// List lists all existing plans.
func (a *PlanService) List() (*SeewebPlanListResponse, *Response, error) {
	u := "/plans"
	v := new(SeewebPlanListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
