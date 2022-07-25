package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebActions_Basic(t *testing.T) {
	actionId := "10535"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebActionsConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebActionsExists("data.seeweb_actions.testacc"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_actions.testacc", "actions.#"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"data.seeweb_actions.testacc",
						"actions.*",
						map[string]string{
							"id": actionId,
						}),
				),
			},
		},
	})
}

func testAccDataSourceSeewebActionsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a actions data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebActionsConfig() string {
	return fmt.Sprintf(`
    data "seeweb_actions" "testacc" {}
`)
}
