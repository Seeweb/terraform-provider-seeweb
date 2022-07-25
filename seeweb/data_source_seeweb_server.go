package seeweb

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebServerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"plan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_size": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"core": {
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
					},
				},
			},
			"so": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deletion_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active_flag": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSeewebServerRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb server %s", d.Get("name").(string))
	err := fetchServer(d, meta, handleNotFoundError)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))
	return nil
}
