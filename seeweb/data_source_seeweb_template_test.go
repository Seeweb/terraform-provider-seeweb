package seeweb

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceSeewebTemplate_Basic(t *testing.T) {
	templateId := "504"
	templateName := "ei200088"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSeewebTemplateConfig(templateId),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceSeewebTemplateExists("data.seeweb_template.testacc"),
					resource.TestCheckResourceAttr(
						"data.seeweb_template.testacc", "id", templateId),
					resource.TestCheckResourceAttr(
						"data.seeweb_template.testacc", "name", templateName),
				),
			},
		},
	})
}

func TestAccDataSourceSeewebTemplate_NotFound(t *testing.T) {
	templateId := "9999999"
	expectedError := "failed 404 Not Found"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceSeewebTemplateConfig(templateId),
				ExpectError: regexp.MustCompile(expectedError),
			},
		},
	})
}

func testAccDataSourceSeewebTemplateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to get a template data source ID from state")
		}

		return nil
	}
}

func testAccDataSourceSeewebTemplateConfig(id string) string {
	return fmt.Sprintf(`
    data "seeweb_template" "testacc" {
      id = %s
    }
`, id)
}
