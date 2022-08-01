package seeweb

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	TEST_RESOURCE_PREFIX = "DESTROY-AFTER-TESTACC"
)

func init() {
	resource.AddTestSweepers("seeweb_server", &resource.Sweeper{
		Name: "seeweb_server",
		F:    testSweepServer,
	})
}

func testSweepServer(region string) error {
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

	resp, _, err := client.Server.List()
	if err != nil {
		return err
	}

	for _, server := range resp.Server {
		if strings.HasPrefix(server.Notes, TEST_RESOURCE_PREFIX) {
			log.Printf("[INFO] Destroying server %s", server.Name)
			if _, _, err := client.Server.Delete(server.Name); err != nil {
				return err
			}
		}
	}

	return nil
}

func TestAccSeewebServer_Basic(t *testing.T) {
	plan := "ECS1"
	location := "it-fr2"
	image := "centos-7"
	notes := fmt.Sprintf("%s::server created during acceptance tests", TEST_RESOURCE_PREFIX)
	notesUpdated := fmt.Sprintf("%s::server updated during acceptance tests", TEST_RESOURCE_PREFIX)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSeewebServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSeewebServerConfig(plan, location, image, notes),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSeewebServerExists("seeweb_server.testacc"),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "plan", plan),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "location", location),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "image", image),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "notes", notes),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "active_flag", "false"),
				),
			},
			{
				Config: testAccCheckSeewebServerConfig(plan, location, image, notesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSeewebServerExists("seeweb_server.testacc"),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "plan", plan),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "location", location),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "image", image),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "notes", notesUpdated),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "active_flag", "false"),
				),
			},
		},
	})
}

func TestAccSeewebServer_WithGroup(t *testing.T) {
	plan := "ECS1"
	location := "it-fr2"
	image := "centos-7"
	serverNotes := fmt.Sprintf("%s::server created during acceptance tests", TEST_RESOURCE_PREFIX)
	groupNotes := fmt.Sprintf("%s::group created during acceptance tests", TEST_RESOURCE_PREFIX)
	groupPass := "secret"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSeewebServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSeewebServerWithGroupConfig(groupNotes, groupPass, plan, location, image, serverNotes),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSeewebServerExists("seeweb_server.testacc"),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "plan", plan),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "location", location),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "image", image),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "notes", serverNotes),
					resource.TestCheckResourceAttrSet(
						"seeweb_server.testacc", "group"),
				),
			},
			{
				Config: testAccCheckSeewebServerWithGroupUpdatedConfig(groupNotes, groupPass, plan, location, image, serverNotes),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSeewebServerExists("seeweb_server.testacc"),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "plan", plan),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "location", location),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "image", image),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "notes", serverNotes),
					resource.TestCheckResourceAttr(
						"seeweb_server.testacc", "group", "nogroup"),
				),
			},
		},
	})
}

func testAccCheckSeewebServerDestroy(s *terraform.State) error {
	client, _ := testAccProvider.Meta().(*Config).Client()
	for _, r := range s.RootModule().Resources {
		if r.Type != "seeweb_server" {
			continue
		}

		server, err := getServerByName(r.Primary.ID, client)
		if err != nil {
			return err
		}
		if server != nil {
			return fmt.Errorf("Server still %q exists", r.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSeewebServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Server ID is set")
		}

		client, _ := testAccProvider.Meta().(*Config).Client()

		server, err := getServerByName(rs.Primary.ID, client)
		if err != nil {
			return err
		}

		if server.Name != rs.Primary.ID {
			return fmt.Errorf("Server not found: %v", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckSeewebServerConfig(plan, location, image, notes string) string {
	return fmt.Sprintf(`
    resource "seeweb_server" "testacc" {
      plan        = "%s"
      location       = "%s"
      image       = "%s"
      notes       = "%s"
    }

`, plan, location, image, notes)
}

func testAccCheckSeewebServerWithGroupConfig(groupNotes, groupPass, plan, location, image, serverNotes string) string {
	return fmt.Sprintf(`
    resource "seeweb_group" "testacc" {
      notes        = "%s"
      password       = "%s"
    }

    resource "seeweb_server" "testacc" {
      depends_on = [seeweb_group.testacc]

      plan        = "%s"
      location       = "%s"
      image       = "%s"
      notes       = "%s"
      group = seeweb_group.testacc.name
    }

`, groupNotes, groupPass, plan, location, image, serverNotes)
}

func testAccCheckSeewebServerWithGroupUpdatedConfig(groupNotes, groupPass, plan, location, image, serverNotes string) string {
	return fmt.Sprintf(`
    resource "seeweb_group" "testacc" {
      notes        = "%s"
      password       = "%s"
    }

    resource "seeweb_server" "testacc" {
      depends_on = [seeweb_group.testacc]

      plan        = "%s"
      location       = "%s"
      image       = "%s"
      notes       = "%s"
      group = "nogroup"
    }

`, groupNotes, groupPass, plan, location, image, serverNotes)
}
