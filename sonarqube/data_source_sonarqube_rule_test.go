package sonarqube

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccSonarqubeRuleDataSourceConfig(rnd string, custom_key string, markdown_description string, name string, template_key string, severity string, status string, type_p string) string {
	return fmt.Sprintf(`
		resource "sonarqube_rule" "%[1]s" {
			custom_key = "%[2]s"
			markdown_description = "%[3]s"
			name = "%[4]s"
			template_key = "%[5]s"
			severity = "%[6]s"
			status = "%[7]s"
			type = "%[8]s"
		}

		data "sonarqube_rule" "%[1]s" {
			key = sonarqube_rule.%[1]s.id
		}`, rnd, custom_key, markdown_description, name, template_key, severity, status, type_p)
}

func TestAccSonarqubeRuleDataSource(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "data.sonarqube_rule." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSonarqubeRuleDataSourceConfig(rnd, "basicRule", "markdown_description", "name", "xml:XPathCheck", "INFO", "READY", "VULNERABILITY"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "markdown_description", "markdown_description"),
					resource.TestCheckResourceAttr(name, "name", "name"),
					resource.TestCheckResourceAttr(name, "template_key", "xml:XPathCheck"),
					resource.TestCheckResourceAttr(name, "severity", "INFO"),
					resource.TestCheckResourceAttr(name, "status", "READY"),
					resource.TestCheckResourceAttr(name, "type", "VULNERABILITY"),
				),
			},
		},
	})
}
