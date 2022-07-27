package seeweb

import (
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebTemplatesRead,

		Schema: map[string]*schema.Schema{
			"templates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_date": {
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
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"notes": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func fetchTemplateListData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(15*time.Second, func() *resource.RetryError {
		resp, _, err := client.Template.List()
		if err != nil {
			log.Printf("[INFO] Templates read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if err := flattenTemplates(d, resp.Templates); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb templates")
	err := fetchTemplateListData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func flattenTemplates(d *schema.ResourceData, templateList []*seeweb.Template) error {
	var templates []map[string]interface{}
	for _, template := range templateList {
		templates = append(templates, map[string]interface{}{
			"id":            template.ID,
			"name":          template.Name,
			"creation_date": template.CreationDate.Format(time.RFC3339),
			"active_flag":   template.ActiveFlag,
			"status":        template.Status,
			"uuid":          template.UUID,
			"notes":         template.Notes,
		})
	}

	d.Set("templates", templates)

	return nil
}
