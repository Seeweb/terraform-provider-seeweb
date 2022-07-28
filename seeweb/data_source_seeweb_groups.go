package seeweb

import (
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebGroupsRead,

		Schema: map[string]*schema.Schema{
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
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
				},
			},
		},
	}
}

func fetchGroupListData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(15*time.Second, func() *resource.RetryError {
		resp, _, err := client.Group.List()
		if err != nil {
			log.Printf("[INFO] Groups read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if err := flattenGroups(d, resp.Groups); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebGroupsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb groups")
	err := fetchGroupListData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func flattenGroups(d *schema.ResourceData, groupList []*seeweb.Group) error {
	var groups []map[string]interface{}
	for _, group := range groupList {
		groups = append(groups, map[string]interface{}{
			"name":    group.Name,
			"id":      group.ID,
			"notes":   group.Notes,
			"enabled": group.Enabled,
		})
	}

	d.Set("groups", groups)

	return nil
}
