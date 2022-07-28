package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebGroup_Basic(t *testing.T) {
	notes := fmt.Sprintf("%s::group created during acceptance tests", TEST_RESOURCE_PREFIX)
	password := "secret"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebGroupConfig(notes, password),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebGroupExists("data.seeweb_group.testacc"),
					resource.TestCheckResourceAttr(
						"data.seeweb_group.testacc", "notes", notes),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_group.testacc", "id"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_group.testacc", "name"),
					resource.TestCheckResourceAttr(
						"data.seeweb_group.testacc", "enabled", "true"),
				),
			},
		},
	})
}

func testAccDataSourceSeewebGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a group data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebGroupConfig(notes, password string) string {
	return fmt.Sprintf(`
resource "seeweb_group" "testacc" {
  notes        = "%s"
  password       = "%s"
}

data "seeweb_group" "testacc" {
  depends_on = [seeweb_group.testacc]
  id = seeweb_group.testacc.id
}
`, notes, password)
}
