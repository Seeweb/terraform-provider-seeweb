package seeweb

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("seeweb_group", &resource.Sweeper{
		Name: "seeweb_group",
		F:    testSweepGroup,
	})
}

func testSweepGroup(region string) error {
	if os.Getenv("SEEWEB_TOKEN") == "" {
		return fmt.Errorf("$SEEWEB_TOKEN must be set")
	}

	config := &Config{
		Token: os.Getenv("SEEWEB_TOKEN"),
	}

	client, err := config.Client()
	if err != nil {
		return err
	}

	resp, _, err := client.Group.List()
	if err != nil {
		return err
	}

	for _, group := range resp.Groups {
		if strings.HasPrefix(group.Notes, TEST_RESOURCE_PREFIX) {
			log.Printf("[INFO] Destroying group %d", group.ID)
			if _, _, err := client.Group.Delete(group.ID); err != nil {
				return err
			}
		}
	}

	return nil
}

func TestAccSeewebGroup_Basic(t *testing.T) {
	notes := fmt.Sprintf("%s::group created during acceptance tests", TEST_RESOURCE_PREFIX)
	password := "secret"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSeewebGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSeewebGroupConfig(notes, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSeewebGroupExists("seeweb_group.testacc"),
					resource.TestCheckResourceAttr(
						"seeweb_group.testacc", "notes", notes),
					resource.TestCheckResourceAttrSet(
						"seeweb_group.testacc", "id"),
					resource.TestCheckResourceAttrSet(
						"seeweb_group.testacc", "name"),
					resource.TestCheckResourceAttr(
						"seeweb_group.testacc", "enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckSeewebGroupDestroy(s *terraform.State) error {
	client, _ := testAccProvider.Meta().(*Config).Client()
	for _, r := range s.RootModule().Resources {
		if r.Type != "seeweb_group" {
			continue
		}

		id, err := strconv.Atoi(r.Primary.ID)
		if err != nil {
			return err
		}
		group, err := getGroupByID(id, client)
		if err != nil {
			return err
		}
		if group != nil {
			return fmt.Errorf("Group still %q exists", r.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSeewebGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Group ID is set")
		}

		client, _ := testAccProvider.Meta().(*Config).Client()

		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		group, err := getGroupByID(id, client)
		if err != nil {
			return err
		}

		if strconv.Itoa(group.ID) != rs.Primary.ID {
			return fmt.Errorf("Group not found: %v", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckSeewebGroupConfig(notes, password string) string {
	return fmt.Sprintf(`
resource "seeweb_group" "testacc" {
  notes        = "%s"
  password       = "%s"
}

`, notes, password)
}
