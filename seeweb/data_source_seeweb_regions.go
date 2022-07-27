package seeweb

import (
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebRegionsRead,

		Schema: map[string]*schema.Schema{
			"regions": {
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
	}
}

func fetchRegionListData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(15*time.Second, func() *resource.RetryError {
		resp, _, err := client.Region.List()
		if err != nil {
			log.Printf("[INFO] Regions read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if err := flattenRegions(d, resp.Regions); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebRegionsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb regions")
	err := fetchRegionListData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func flattenRegions(d *schema.ResourceData, regionList []*seeweb.Region) error {
	var regions []map[string]interface{}
	for _, region := range regionList {
		regions = append(regions, map[string]interface{}{
			"id":          region.ID,
			"location":    region.Location,
			"description": region.Description,
		})
	}

	d.Set("regions", regions)

	return nil
}
