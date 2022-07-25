package seeweb

import (
	"fmt"
	"log"
	"runtime"

	"github.com/Seeweb/cloudserver-go-client/seeweb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider represents a resource provider in Terraform
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEEWEB_TOKEN", nil),
			},

			"api_url_override": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"seeweb_server":  dataSourceSeewebServer(),
			"seeweb_action":  dataSourceSeewebAction(),
			"seeweb_actions": dataSourceSeewebActions(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"seeweb_server": resourceSeewebServer(),
		},
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return p
}

func isErrCode(err error, code int) bool {
	if e, ok := err.(*seeweb.Error); ok && e.ErrorResponse.Response.StatusCode == code {
		return true
	}

	return false
}

func genError(err error, d *schema.ResourceData) error {
	return fmt.Errorf("Error reading: %s: %s", d.Id(), err)
}

func handleNotFoundError(err error, d *schema.ResourceData) error {
	if isErrCode(err, 404) {
		log.Printf("[WARN] Removing %s because it's gone", d.Id())
		d.SetId("")
		return nil
	}
	return genError(err, d)
}

func providerConfigure(data *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		ApiUrl:         "https://api.seeweb.it/ecs/v2",
		Token:          data.Get("token").(string),
		UserAgent:      fmt.Sprintf("(%s %s) Terraform/%s", runtime.GOOS, runtime.GOARCH, terraformVersion),
		ApiUrlOverride: data.Get("api_url_override").(string),
	}

	log.Println("[INFO] Initializing Seeweb client")
	return &config, nil
}
