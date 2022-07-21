package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebServer_Basic(t *testing.T) {
	plan := "ECS1"
	location := "it-fr2"
	image := "centos-7"
	notes := fmt.Sprintf("%s::server created during acceptance tests", TEST_RESOURCE_PREFIX)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebServerConfig(plan, location, image, notes),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebServerExists("data.seeweb_server.testacc"),
					resource.TestCheckResourceAttr(
						"data.seeweb_server.testacc", "plan", "eCS1"),
					resource.TestCheckResourceAttr(
						"data.seeweb_server.testacc", "location", location),
					resource.TestCheckResourceAttr(
						"data.seeweb_server.testacc", "notes", notes),
				),
			},
		},
	})
}

func testAccDataSourceSeewebServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a server data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebServerConfig(plan, location, image, notes string) string {
	return fmt.Sprintf(`
    resource "seeweb_server" "testacc" {
      plan        = "%s"
      location       = "%s"
      image       = "%s"
      notes       = "%s"
    }

    data "seeweb_server" "testacc" {
      depends_on = [seeweb_server.testacc]
      name = seeweb_server.testacc.name
    }
`, plan, location, image, notes)
}
