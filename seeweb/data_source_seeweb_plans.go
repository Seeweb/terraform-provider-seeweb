package seeweb

import (
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebPlans() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebPlansRead,

		Schema: map[string]*schema.Schema{
			"plans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ram": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"disk": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hourly_price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"montly_price": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"windows": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"available": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"available_regions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func fetchPlanListData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(15*time.Second, func() *resource.RetryError {
		resp, _, err := client.Plan.List()
		if err != nil {
			log.Printf("[INFO] Plans read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if err := flattenPlans(d, resp.Plans); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebPlansRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb plans")
	err := fetchPlanListData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func flattenPlans(d *schema.ResourceData, planList []*seeweb.Plan) error {
	var plans []map[string]interface{}
	for _, plan := range planList {
		plans = append(plans, map[string]interface{}{
			"id":                plan.ID,
			"name":              plan.Name,
			"cpu":               plan.CPU,
			"ram":               plan.RAM,
			"disk":              plan.Disk,
			"hourly_price":      plan.HourlyPrice,
			"montly_price":      plan.MontlyPrice,
			"windows":           plan.Windows,
			"available":         plan.Available,
			"available_regions": flattenPlanAvailableRegions(plan.AvailableRegions),
		})
	}

	d.Set("plans", plans)

	return nil
}

func flattenPlanAvailableRegions(v []*seeweb.AvailableRegions) interface{} {
	availableRegions := []map[string]interface{}{}

	for _, ar := range v {
		availableRegions = append(availableRegions, map[string]interface{}{
			"id":          ar.ID,
			"location":    ar.Location,
			"description": ar.Description,
		})
	}
	return availableRegions
}
