package seeweb

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebTemplateRead,

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
	}
}

func fetchTemplateData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(30*time.Second, func() *resource.RetryError {
		id := d.Get("id").(int)
		resp, _, err := client.Template.Get(strconv.Itoa(id))
		if err != nil {
			log.Printf("[INFO] Template read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if resp.Template == nil {
			return resource.NonRetryableError(
				fmt.Errorf("Unable to locate any template with the id: %d", id),
			)
		}

		if err := flattenTemplate(d, resp.Template); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebTemplateRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb template %d", d.Get("id").(int))
	err := fetchTemplateData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(d.Get("id").(int)))
	return nil
}

func flattenTemplate(d *schema.ResourceData, template *seeweb.Template) error {
	if _, ok := d.GetOk("id"); !ok {
		d.Set("id", template.ID)
	}
	d.Set("name", template.Name)
	d.Set("creation_date", template.CreationDate.Format(time.RFC3339))
	d.Set("active_flag", template.ActiveFlag)
	d.Set("status", template.Status)
	d.Set("uuid", template.UUID)
	d.Set("notes", template.Notes)

	return nil
}
