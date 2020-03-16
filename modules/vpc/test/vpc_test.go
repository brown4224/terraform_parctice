package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerragruntVPC(t *testing.T) {
	t.Parallel()

	awsRegion := "us-east-1"
	name       := "terragrunt_test"
	vpc_cidr := "10.1.0.0/16"
	subnet_cidr := "10.1.0.0/24"
	ingress_cidr := []string{"136.55.41.30/32"}
	ports := make(map[string]string)
	ports["ssh"] = "22"
	ports["rdp"] = "3389"


	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"aws_region" : awsRegion,
			"name" : name,
			"vpc_cidr" : vpc_cidr,
			"subnet_cidr" : subnet_cidr,
			"ingress_cidr" : ingress_cidr,
			"ports" : ports,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	validate(t, terraformOptions, awsRegion)
}

func validate(t *testing.T, options *terraform.Options, awsRegion string){
	publicSubnetId := terraform.Output(t, options, "subnet_id")
	assert.True(t, aws.IsPublicSubnet(t, publicSubnetId, awsRegion))

}