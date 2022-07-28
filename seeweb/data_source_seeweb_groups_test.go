package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebGroups_Basic(t *testing.T) {
	notes := fmt.Sprintf("%s::group created during acceptance tests", TEST_RESOURCE_PREFIX)
	password := "secret"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebGroupsConfig(notes, password),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebGroupsExists("data.seeweb_groups.testacc"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_groups.testacc", "groups.#"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"data.seeweb_groups.testacc",
						"groups.*",
						map[string]string{
							"notes": notes,
						}),
				),
			},
		},
	})
}

func testAccDataSourceSeewebGroupsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a groups data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebGroupsConfig(notes, password string) string {
	return fmt.Sprintf(`
resource "seeweb_group" "testacc" {
  notes        = "%s"
  password       = "%s"
}

data "seeweb_groups" "testacc" {
  depends_on = [seeweb_group.testacc]
}
`, notes, password)
}
