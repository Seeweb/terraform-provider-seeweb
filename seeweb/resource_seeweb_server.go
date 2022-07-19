package seeweb

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
)

func resourceSeewebServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSeewebServerCreate,
		Read:   resourceSeewebServerRead,
		Update: resourceSeewebServerUpdate,
		Delete: resourceSeewebServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"plan": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
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

func buildServerCreateReq(d *schema.ResourceData) *seeweb.SeewebServerCreateRequest {
	r := seeweb.SeewebServerCreateRequest{
		Plan:     d.Get("plan").(string),
		Location: d.Get("location").(string),
		Image:    d.Get("image").(string),
		Notes:    d.Get("notes").(string),
	}

	if attr, ok := d.GetOk("ssh_key"); ok {
		r.SSHKey = attr.(string)
	}

	return &r
}

func buildServerUpdateReq(d *schema.ResourceData) *seeweb.SeewebServerUpdateRequest {
	r := seeweb.SeewebServerUpdateRequest{
		Note:  d.Get("notes").(string),
		Group: d.Get("group").(string),
	}

	return &r
}

func getServerByName(name string, c *seeweb.Client) (*seeweb.Server, error) {
	resp, _, err := c.Server.List()
	if err != nil {
		return nil, err
	}

	var found *seeweb.Server
	for _, server := range resp.Server {
		if server.Name == name {
			found = server
			break
		}
	}
	return found, nil
}

func fetchServer(d *schema.ResourceData, meta interface{}, errCallback func(error, *schema.ResourceData) error) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(2*time.Minute, func() *resource.RetryError {
		server, err := getServerByName(d.Id(), client)
		if err != nil {
			log.Printf("[WARN] Server read error")
			errResp := errCallback(err, d)
			if errResp != nil {
				time.Sleep(2 * time.Second)
				return resource.RetryableError(errResp)
			}

			return nil

		}

		if server == nil {
			return resource.NonRetryableError(
				fmt.Errorf("Unable to locate any server with the name: %s", d.Id()),
			)
		}

		if err := flattenServer(d, server); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func resourceSeewebServerCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	req := buildServerCreateReq(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating Seeweb server with plan %s, location %s and iamge %s", req.Plan, req.Location, req.Image)

	resp, _, err := client.Server.Create(req)
	if err != nil {
		return err
	}

	d.SetId(resp.Server.Name)

	return fetchServer(d, meta, genError)
}

func resourceSeewebServerRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb server %s", d.Id())
	return fetchServer(d, meta, handleNotFoundError)
}

func resourceSeewebServerUpdate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	req := buildServerUpdateReq(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Updating Seeweb server %s", d.Id())

	_, _, err = client.Server.Update(d.Id(), req)
	if err != nil {
		return err
	}

	return nil
}

func resourceSeewebServerDelete(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(2*time.Minute, func() *resource.RetryError {
		log.Printf("[INFO] Deleting Seeweb server %s", d.Id())
		if _, _, err := client.Server.Delete(d.Id()); err != nil {
			time.Sleep(30 * time.Second)
			return resource.RetryableError(err)
		}

		d.SetId("")
		return nil
	})
}

func flattenServer(d *schema.ResourceData, server *seeweb.Server) error {
	d.Set("name", server.Name)
	d.Set("ipv4", server.Ipv4)
	d.Set("ipv6", server.Ipv6)
	d.Set("so", server.So)
	d.Set("creation_date", server.CreationDate.Format(time.RFC3339))
	d.Set("deletion_date", server.DeletionDate.Format(time.RFC3339))
	d.Set("active_flag", server.ActiveFlag)
	d.Set("status", server.Status)
	d.Set("api_version", server.APIVersion)
	// d.Set("group", server.Group)
	d.Set("user", server.User)

	if server.PlanSize != nil {
		if err := d.Set("plan_size", flattenPlanSize(server.PlanSize)); err != nil {
			return err
		}
	}
	return nil
}

func flattenPlanSize(v *seeweb.PlanSize) interface{} {
	planSize := map[string]interface{}{
		"core": v.Core,
		"ram":  v.RAM,
		"disk": v.Disk,
	}

	return []interface{}{planSize}
}
