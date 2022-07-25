package seeweb

import (
	"fmt"
	"log"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebAction() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebActionRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"started_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"completed_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func fetchAction(d *schema.ResourceData, meta interface{}, errCallback func(error, *schema.ResourceData) error) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(30*time.Second, func() *resource.RetryError {
		id := d.Get("id").(string)
		resp, _, err := client.Action.Get(id)
		if err != nil {
			log.Printf("[INFO] Action read error. Retrying in %d seconds", retryAfter5Seconds)
			errResp := errCallback(err, d)
			if errResp != nil {
				time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
				return resource.RetryableError(errResp)
			}

			return nil

		}

		if resp.Action == nil {
			return resource.NonRetryableError(
				fmt.Errorf("Unable to locate any action with the id: %s", id),
			)
		}

		if err := flattenAction(d, resp.Action); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebActionRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb action %s", d.Get("id").(string))
	err := fetchAction(d, meta, handleNotFoundError)
	if err != nil {
		return err
	}

	d.SetId(d.Get("id").(string))
	return nil
}

func flattenAction(d *schema.ResourceData, action *seeweb.Action) error {
	if _, ok := d.GetOk("id"); !ok {
		d.Set("id", action.ID)
	}
	d.Set("status", action.Status)
	d.Set("user", action.User)
	d.Set("created_at", action.CreatedAt.Format(time.RFC3339))
	d.Set("started_at", action.StartedAt.Format(time.RFC3339))
	d.Set("completed_at", action.CompletedAt.Format(time.RFC3339))
	d.Set("resource", action.Resource)
	d.Set("resource_type", action.ResourceType)
	d.Set("type", action.Type)
	d.Set("progress", action.Progress)

	return nil
}
