package test

import (
	"testing"
	// "fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEc2(t *testing.T) {
	t.Parallel()

	//  Vars
	awsRegion := "us-east-1"
	name       := "terragrunt_test"
	ami := "ami-0b940a1059f928462"
    instance := "t2.micro"
	keypair := "dev-ansable-key-01"

	//  Example that uses the default VPC
	// vpc := aws.GetDefaultVpc(t, awsRegion)
	// default_subnets, err := aws.GetSubnetsForVpcE(t, vpc.Id, awsRegion)
	// if err != nil {
	// 	fmt.Errorf("There was an error getting default subnets: %v", err)
	//   }
	// subnetId := default_subnets[0].Id
	// ispub, err := aws.IsPublicSubnetE(t, subnetId, awsRegion)
	// if err != nil {
	// 	fmt.Errorf("There with the public subnet: %v", err)
	//   }
	// if ispub == false {
	// 	fmt.Errorf("The subnet is private: %v", err)
	//   }


	// Example Create VPC from scratch
	ingress := "136.55.41.30/32"
	vpc := createVpc(t, awsRegion)
	routeTable := createRouteTable(t, *vpc.VpcId, awsRegion)
	subnet := createSubnet(t, *vpc.VpcId, *routeTable.RouteTableId, awsRegion)
	assert.False(t, aws.IsPublicSubnet(t, *subnet.SubnetId, awsRegion))
	sg := createSecurityGroup(t, *vpc.VpcId, awsRegion)
	configureSecurityGroup(t, *sg.GroupId, awsRegion, ingress)

	// Set Networking dependencies
	subnetId := *subnet.SubnetId
	security_group_ids := []string{*sg.GroupId}

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"aws_region" : awsRegion,
			"name" : name,
			"ami" : ami,
			"instance" : instance,
			"keypair" : keypair,
			"subnet_id" : subnetId,
			"security_group_ids" :security_group_ids,
		},
	}
	defer func() {
		terraform.Destroy(t, terraformOptions)
		deleteVpc(t, *vpc.VpcId, awsRegion, *sg.GroupId)
	}()

	terraform.InitAndApply(t, terraformOptions)
	// validate(t, terraformOptions, awsRegion)
}

func validate(t *testing.T, options *terraform.Options, awsRegion string){
	// publicSubnetId := terraform.Output(t, options, "subnet_id")
	// assert.True(t, aws.IsPublicSubnet(t, publicSubnetId, awsRegion))
}