package sonarqube

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("sonarqube_qualityprofile_activate_rule", &resource.Sweeper{
		Name: "sonarqube_qualityprofile_activate_rule",
		F:    testSweeepSonarqibeQualityprofileActivateRuleSweeper,
	})
}

func testSweeepSonarqibeQualityprofileActivateRuleSweeper(r string) error {
	return nil
}

func testAccSonarqubeQualityprofileActivateRuleBasicConfig(rnd string, key string, rule string, severity string) string {
	return fmt.Sprintf(`
		resource "sonarqube_qualityprofile_activate_rule" "%[1]s" {
			key = "%[2]s"
			rule = "%[3]s"
			severity = "%[4]s"
		}`, rnd, key, rule, severity)
}

func TestAccSonarqubeQualityprofileActivateRuleBasic(t *testing.T) {
	rnd := generateRandomResourceName()
	name := "sonarqube_quality_profile_activate_rule" + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSonarqubeQualityprofileActivateRuleBasicConfig(rnd, "key", "rule", "severity"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "key", "key"),
					resource.TestCheckResourceAttr(name, "rule", "rule"),
					resource.TestCheckResourceAttr(name, "severity", "severity"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "key", "key"),
					resource.TestCheckResourceAttr(name, "rule", "rule"),
					resource.TestCheckResourceAttr(name, "params", "key1=v1"),
					resource.TestCheckResourceAttr(name, "reset", "false"),
					resource.TestCheckResourceAttr(name, "severity", "BLOCKER"),
				),
			},
		},
	})
}