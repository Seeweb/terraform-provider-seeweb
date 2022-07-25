package seeweb

import (
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebActions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebActionsRead,

		Schema: map[string]*schema.Schema{
			"actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
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
				},
			},
		},
	}
}

func fetchActionList(d *schema.ResourceData, meta interface{}, errCallback func(error, *schema.ResourceData) error) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(15*time.Second, func() *resource.RetryError {
		resp, _, err := client.Action.List()
		if err != nil {
			log.Printf("[INFO] Actions read error. Retrying in %d seconds", retryAfter5Seconds)
			errResp := errCallback(err, d)
			if errResp != nil {
				time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
				return resource.RetryableError(errResp)
			}

			return nil

		}

		if err := flattenActions(d, resp.Actions); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebActionsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb actions")
	err := fetchActionList(d, meta, handleNotFoundError)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func flattenActions(d *schema.ResourceData, actionList []*seeweb.Action) error {
	var actions []map[string]interface{}
	for _, action := range actionList {
		actions = append(actions, map[string]interface{}{
			"id":            action.ID,
			"status":        action.Status,
			"user":          action.User,
			"created_at":    action.CreatedAt.Format(time.RFC3339),
			"started_at":    action.StartedAt.Format(time.RFC3339),
			"completed_at":  action.CompletedAt.Format(time.RFC3339),
			"resource":      action.Resource,
			"resource_type": action.ResourceType,
			"type":          action.Type,
			"progress":      action.Progress,
		})
	}

	d.Set("actions", actions)

	return nil
}
