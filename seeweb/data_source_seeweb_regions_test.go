package seeweb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebRegions_Basic(t *testing.T) {
	expectedLocation := "it-fr2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebRegionsConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebRegionsExists("data.seeweb_regions.testacc"),
					resource.TestCheckResourceAttrSet(
						"data.seeweb_regions.testacc", "regions.#"),
					resource.TestCheckTypeSetElemNestedAttrs(
						"data.seeweb_regions.testacc",
						"regions.*",
						map[string]string{
							"location": expectedLocation,
						}),
				),
			},
		},
	})
}

func testAccDataSourceSeewebRegionsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a regions data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebRegionsConfig() string {
	return fmt.Sprintf(`
    data "seeweb_regions" "testacc" {}
`)
}
