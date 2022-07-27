package seeweb

import (
	"log"
	"strconv"
	"time"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSeewebServers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSeewebServersRead,

		Schema: map[string]*schema.Schema{
			"servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func fetchServerListData(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	return resource.Retry(15*time.Second, func() *resource.RetryError {
		resp, _, err := client.Server.List()
		if err != nil {
			log.Printf("[INFO] Servers read error. Retrying in %d seconds", retryAfter5Seconds)
			time.Sleep(time.Duration(retryAfter5Seconds) * time.Second)
			return resource.RetryableError(err)
		}

		if err := flattenServers(d, resp.Server); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
}

func dataSourceSeewebServersRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Seeweb servers")
	err := fetchServerListData(d, meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

func flattenServers(d *schema.ResourceData, serverList []*seeweb.Server) error {
	var servers []map[string]interface{}
	for _, server := range serverList {
		servers = append(servers, map[string]interface{}{
			"name":          server.Name,
			"plan":          server.Plan,
			"location":      server.Location,
			"notes":         server.Notes,
			"ssh_key":       server.SSHKey,
			"ipv4":          server.Ipv4,
			"ipv6":          server.Ipv6,
			"so":            server.So,
			"creation_date": server.CreationDate.Format(time.RFC3339),
			"deletion_date": server.DeletionDate.Format(time.RFC3339),
			"active_flag":   server.ActiveFlag,
			"status":        server.Status,
			"api_version":   server.APIVersion,
			"group":         server.Group,
			"user":          server.User,
			"plan_size":     flattenPlanSize(server.PlanSize),
		})
	}

	d.Set("servers", servers)

	return nil
}
