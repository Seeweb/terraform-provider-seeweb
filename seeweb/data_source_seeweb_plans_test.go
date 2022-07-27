package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebPlans_Basic(t *testing.T) {
	expectedPlanName := "eCS1"
	expectedLocation := "it-fr2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebPlansConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebPlansExists("data.seeweb_plans.testacc"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_plans.testacc", "plans.#"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"data.seeweb_plans.testacc",
						"plans.*",
						map[string]string{
							"name": expectedPlanName,
						}),
					resource.TestCheckTypeSetElemNestedAttrs(
						"data.seeweb_plans.testacc",
						"plans.*.available_regions.*",
						map[string]string{
							"location": expectedLocation,
						}),
				),
			},
		},
	})
}

func testAccDataSourceSeewebPlansExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a plans data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebPlansConfig() string {
	return fmt.Sprintf(`
    data "seeweb_plans" "testacc" {}
`)
}
