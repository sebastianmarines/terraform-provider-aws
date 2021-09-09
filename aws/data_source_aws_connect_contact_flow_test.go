package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAwsConnectContactFlow_basic(t *testing.T) {
	rName := acctest.RandomWithPrefix("resource-test-terraform")
	resourceName := "aws_connect_contact_flow.test"
	datasourceName := "data.aws_connect_contact_flow.test"

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { testAccPreCheck(t) },
		ErrorCheck: testAccErrorCheck(t, connect.EndpointsID),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsConnectContactFlowDataSourceConfig_basic(rName, resourceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "contact_flow_id", resourceName, "contact_flow_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "instance_id", resourceName, "instance_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(datasourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(datasourceName, "content", resourceName, "content"),
					resource.TestCheckResourceAttrPair(datasourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
				),
			},
		},
	})
}

func TestAccDataSourceAwsConnectContactFlow_byname(t *testing.T) {
	rName := acctest.RandomWithPrefix("resource-test-terraform")
	rName2 := acctest.RandomWithPrefix("resource-test-terraform")
	resourceName := "aws_connect_contact_flow.test"
	datasourceName := "data.aws_connect_contact_flow.test"

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { testAccPreCheck(t) },
		ErrorCheck: testAccErrorCheck(t, connect.EndpointsID),
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsConnectContactFlowDataSourceConfig_byname(rName, rName2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "id", resourceName, "id"),
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "contact_flow_id", resourceName, "contact_flow_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "instance_id", resourceName, "instance_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(datasourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(datasourceName, "content", resourceName, "content"),
					resource.TestCheckResourceAttrPair(datasourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(datasourceName, "tags.%", resourceName, "tags.%"),
				),
			},
		},
	})
}

func testAccAwsConnectContactFlowDataSourceBaseConfig(rName, rName2 string) string {
	return fmt.Sprintf(`
resource "aws_connect_instance" "test" {
  identity_management_type = "CONNECT_MANAGED"
  inbound_calls_enabled    = true
  instance_alias           = %[1]q
  outbound_calls_enabled   = true
}

resource "aws_connect_contact_flow" "test" {
  instance_id = aws_connect_instance.test.id
  name        = %[2]q
  description = "Test Contact Flow Description"
  type        = "CONTACT_FLOW"
  content     = <<JSON
	{
		"Version": "2019-10-30",
		"StartAction": "12345678-1234-1234-1234-123456789012",
		"Actions": [ 
			{
				"Identifier": "12345678-1234-1234-1234-123456789012",
				"Type": "MessageParticipant",
				"Transitions": {
					"NextAction": "abcdef-abcd-abcd-abcd-abcdefghijkl",
					"Errors": [],
					"Conditions": []
				},
				"Parameters": {
					"Text": "Thanks for calling the sample flow!"
				}
			},
			{
				"Identifier": "abcdef-abcd-abcd-abcd-abcdefghijkl",
				"Type": "DisconnectParticipant",
				"Transitions": {},
				"Parameters": {}
			}
		]
	}
	JSON
  tags = {
    "Name"        = "Test Contact Flow",
    "Application" = "Terraform",
    "Method"      = "Create"
  }
}
	`, rName, rName2)
}

func testAccAwsConnectContactFlowDataSourceConfig_basic(rName, rName2 string) string {
	return fmt.Sprintf(testAccAwsConnectContactFlowDataSourceBaseConfig(rName, rName2) + `
data "aws_connect_contact_flow" "test" {
  instance_id     = aws_connect_instance.test.id
  contact_flow_id = aws_connect_contact_flow.test.contact_flow_id
}
`)
}

func testAccAwsConnectContactFlowDataSourceConfig_byname(rName, rName2 string) string {
	return fmt.Sprintf(testAccAwsConnectContactFlowDataSourceBaseConfig(rName, rName2) + `
data "aws_connect_contact_flow" "test" {
  instance_id = aws_connect_instance.test.id
  name        = aws_connect_contact_flow.test.name
}
`)
}
