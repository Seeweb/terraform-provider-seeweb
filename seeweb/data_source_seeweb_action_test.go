package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebAction_Basic(t *testing.T) {
	actionId := "10535"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebActionConfig(actionId),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebActionExists("data.seeweb_action.testacc"),
					resource.TestCheckResourceAttr(
						"data.seeweb_action.testacc", "id", actionId),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "status"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "user"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "created_at"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "started_at"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "completed_at"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "resource"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "resource_type"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "type"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_action.testacc", "progress"),
				),
			},
		},
	})
}

func testAccDataSourceSeewebActionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a action data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebActionConfig(id string) string {
	return fmt.Sprintf(`
    data "seeweb_action" "testacc" {
      id = %s
    }
`, id)
}
