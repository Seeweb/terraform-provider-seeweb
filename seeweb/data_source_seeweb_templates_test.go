package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebTemplates_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebTemplatesConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebTemplatesExists("data.seeweb_templates.testacc"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_templates.testacc", "templates.#"),
				),
			},
		},
	})
}

func testAccDataSourceSeewebTemplatesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a templates data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebTemplatesConfig() string {
	return fmt.Sprintf(`
    data "seeweb_templates" "testacc" {}
`)
}
