package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebServers_Basic(t *testing.T) {
	plan := "ECS1"
	location := "it-fr2"
	image := "centos-7"
	notes := fmt.Sprintf("%s::server created during acceptance tests", TEST_RESOURCE_PREFIX)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebServersConfig(plan, location, image, notes),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebServersExists("data.seeweb_servers.testacc"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_servers.testacc", "servers.#"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"data.seeweb_servers.testacc",
						"servers.*",
						map[string]string{
							"notes": notes,
						}),
				),
			},
		},
	})
}

func testAccDataSourceSeewebServersExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a servers data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebServersConfig(plan, location, image, notes string) string {
	return fmt.Sprintf(`
    resource "seeweb_server" "testacc" {
      plan        = "%s"
      location       = "%s"
      image       = "%s"
      notes       = "%s"
    }

    data "seeweb_servers" "testacc" {
      depends_on = [seeweb_server.testacc]
    }
`, plan, location, image, notes)
}
