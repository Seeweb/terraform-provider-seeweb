package seeweb

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebGroupRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func fetchGroupData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(30*time.Second, func() *resource.RetryError {
		id := d.Get("id").(int)
		group, err := getGroupByID(id, client)
		if err != nil {
			log.Printf("[INFO] Group read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if group == nil {
			return resource.NonRetryableError(
				fmt.Errorf("Unable to locate any group with the id: %d", id),
			)
		}

		if err := flattenGroup(d, group); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb group %d", d.Get("id").(int))
	err := fetchGroupData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(d.Get("id").(int)))
	return nil
}
