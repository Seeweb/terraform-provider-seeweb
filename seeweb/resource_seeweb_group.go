package seeweb

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
)

func resourceSeewebGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSeewebGroupCreate,
		Read:   resourceSeewebGroupRead,
		Delete: resourceSeewebGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"notes": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func buildGroupCreateReq(d *schema.ResourceData) *seeweb.SeewebGroupCreateRequest {
	r := seeweb.SeewebGroupCreateRequest{
		Notes:    d.Get("notes").(string),
		Password: d.Get("password").(string),
	}

	return &r
}

func getGroupByID(id int, c *seeweb.Client) (*seeweb.Group, error) {
	resp, _, err := c.Group.List()
	if err != nil {
		return nil, err
	}

	var found *seeweb.Group
	for _, group := range resp.Groups {
		if group.ID == id {
			found = group
			break
		}
	}
	return found, nil
}

func fetchGroup(d *schema.ResourceData, meta interface{}, errCallback func(error, *schema.ResourceData) error) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(30*time.Second, func() *resource.RetryError {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}
		group, err := getGroupByID(id, client)
		if err != nil {
			log.Printf("[INFO] Group read error. Retrying in %d seconds", retryAfter5Seconds)
			errResp := errCallback(err, d)
			if errResp != nil {
				time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
				return resource.RetryableError(errResp)
			}

			return nil
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

func resourceSeewebGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	req := buildGroupCreateReq(d)
	if err != nil {
		return err
	}

	log.Println("[INFO] Creating Seeweb group")

	resp, _, err := client.Group.Create(req)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.Group.ID))

	return fetchGroup(d, meta, genError)
}

func resourceSeewebGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb group %s", d.Id())
	return fetchGroup(d, meta, handleNotFoundError)
}

func resourceSeewebGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(1*time.Minute, func() *resource.RetryError {
		log.Printf("[INFO] Deleting Seeweb group %s", d.Id())
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if _, _, err := client.Group.Delete(id); err != nil {
			log.Printf("[INFO] Group deletion error. Retrying in %d seconds", retryAfter30Seconds)
			time.Sleep(time.Duration(retryAfter30Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		d.SetId("")
		return nil
	})
}

func flattenGroup(d *schema.ResourceData, group *seeweb.Group) error {
	if _, ok := d.GetOk("notes"); !ok {
		d.Set("notes", group.Notes)
	}
	d.Set("name", group.Name)
	d.Set("enabled", group.Enabled)
	return nil
}
