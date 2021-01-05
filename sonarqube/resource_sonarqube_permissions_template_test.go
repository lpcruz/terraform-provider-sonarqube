package sonarqube

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("sonarqube_permission_template", &resource.Sweeper{
		Name: "sonarqube_permission_template",
		F:    testSweepPermissionTemplateSweeper,
	})
}

// TODO: implement sweeper to clean up permission_template: https://www.terraform.io/docs/extend/testing/acceptance-tests/sweepers.html
func testSweepPermissionTemplateSweeper(r string) error {
	return nil
}

func testAccSonarqubePermissionTemplateBasicConfig(rnd string, name string, description string, projectKeyPattern string) string {
	return fmt.Sprintf(`
		resource "sonarqube_permission_template" "%[1]s" {
		  name                = "%[2]s"
		  description         = "%[3]s"
		  project_key_pattern = "%[4]s"
		}`, rnd, name, description, projectKeyPattern)
}

func TestAccSonarqubePermissionTemplateBasic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "sonarqube_permission_template." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSonarqubePermissionTemplateBasicConfig(rnd, "testAccSonarqubePermissionTemplate", "These are internal projects", "internal.*"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "testAccSonarqubePermissionTemplate"),
					resource.TestCheckResourceAttr(name, "description", "These are internal projects"),
					resource.TestCheckResourceAttr(name, "project_key_pattern", "internal.*"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", "testAccSonarqubePermissionTemplate"),
					resource.TestCheckResourceAttr(name, "description", "These are internal projects"),
					resource.TestCheckResourceAttr(name, "project_key_pattern", "internal.*"),
				),
			},
		},
	})
}
