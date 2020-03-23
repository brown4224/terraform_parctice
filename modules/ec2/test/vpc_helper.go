
package test

import (
	"testing"
	"log"
	// "fmt"
	"github.com/stretchr/testify/require"
	terratest "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// From https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/ec2/ec2_create_security_group.go

func createPublicRoute(t *testing.T, vpcId string, routeTableId string, region string) {
	ec2Client := terratest.NewEc2Client(t, region)

	createIGWOut, igerr := ec2Client.CreateInternetGateway(&ec2.CreateInternetGatewayInput{})
	require.NoError(t, igerr)

	_, aigerr := ec2Client.AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: createIGWOut.InternetGateway.InternetGatewayId,
		VpcId:             aws.String(vpcId),
	})
	require.NoError(t, aigerr)

	_, err := ec2Client.CreateRoute(&ec2.CreateRouteInput{
		RouteTableId:         aws.String(routeTableId),
		DestinationCidrBlock: aws.String("0.0.0.0/0"),
		GatewayId:            createIGWOut.InternetGateway.InternetGatewayId,
	})

	require.NoError(t, err)
}

func createRouteTable(t *testing.T, vpcId string, region string) ec2.RouteTable {
	ec2Client := terratest.NewEc2Client(t, region)

	createRouteTableOutput, err := ec2Client.CreateRouteTable(&ec2.CreateRouteTableInput{
		VpcId: aws.String(vpcId),
	})

	require.NoError(t, err)
	return *createRouteTableOutput.RouteTable
}

func createSubnet(t *testing.T, vpcId string, routeTableId string, region string) ec2.Subnet {
	ec2Client := terratest.NewEc2Client(t, region)

	createSubnetOutput, err := ec2Client.CreateSubnet(&ec2.CreateSubnetInput{
		CidrBlock: aws.String("10.10.1.0/24"),
		VpcId:     aws.String(vpcId),
	})
	require.NoError(t, err)

	_, err = ec2Client.AssociateRouteTable(&ec2.AssociateRouteTableInput{
		RouteTableId: aws.String(routeTableId),
		SubnetId:     aws.String(*createSubnetOutput.Subnet.SubnetId),
	})
	require.NoError(t, err)

	return *createSubnetOutput.Subnet
}

func createSecurityGroup(t *testing.T, vpcID string, region string) ec2.CreateSecurityGroupOutput{
	name := "Go Security Group"
	desc := "Go Group Description"
	
	ec2Client := terratest.NewEc2Client(t, region)
	sg , err := ec2Client.CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
        GroupName:   &name,
        Description: &desc,
		VpcId:       &vpcID,	
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case "InvalidVpcID.NotFound":
                log.Println("Unable to find VPC with ID.", vpcID)
            case "Invali*dGroup.Duplicate":
                log.Println("Security group already exists.", name)
            }
        }
        log.Println("Unable to create security group", err)
	}
	return *sg

}
func configureSecurityGroup(t *testing.T, sgid string, region string, ingressIP string) {
	ec2Client := terratest.NewEc2Client(t, region)
	_, err := ec2Client.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: &sgid,
		// GroupName: &name,
        IpPermissions: []*ec2.IpPermission{
            // Can use setters to simplify seting multiple values without the
            // needing to use aws.String or associated helper utilities.
            (&ec2.IpPermission{}).
                SetIpProtocol("tcp").
                SetFromPort(22).
                SetToPort(22).
                SetIpRanges([]*ec2.IpRange{
                    (&ec2.IpRange{}).
                        SetCidrIp(ingressIP),
                }),
        },
	})
	require.NoError(t, err)

}
func destroySecurityGroup(t *testing.T, sgId string, region string) {
	ec2Client := terratest.NewEc2Client(t, region)
	input := &ec2.DeleteSecurityGroupInput{
		GroupId: &sgId,
	}

	_, err := ec2Client.DeleteSecurityGroup(input)
	require.NoError(t, err)


}

func createVpc(t *testing.T, region string) ec2.Vpc {
	ec2Client := terratest.NewEc2Client(t, region)

	createVpcOutput, err := ec2Client.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: aws.String("10.10.0.0/16"),
	})

	require.NoError(t, err)
	return *createVpcOutput.Vpc
}

func deleteRouteTables(t *testing.T, vpcId string, region string) {
	ec2Client := terratest.NewEc2Client(t, region)

	vpcIDFilterName := "vpc-id"
	vpcIDFilter := ec2.Filter{Name: &vpcIDFilterName, Values: []*string{&vpcId}}

	// "You can't delete the main route table."
	mainRTFilterName := "association.main"
	mainRTFilterValue := "false"
	notMainRTFilter := ec2.Filter{Name: &mainRTFilterName, Values: []*string{&mainRTFilterValue}}

	filters := []*ec2.Filter{&vpcIDFilter, &notMainRTFilter}

	rtOutput, err := ec2Client.DescribeRouteTables(&ec2.DescribeRouteTablesInput{Filters: filters})
	require.NoError(t, err)

	for _, rt := range rtOutput.RouteTables {

		// "You must disassociate the route table from any subnets before you can delete it."
		for _, assoc := range rt.Associations {
			_, disassocErr := ec2Client.DisassociateRouteTable(&ec2.DisassociateRouteTableInput{
				AssociationId: assoc.RouteTableAssociationId,
			})
			require.NoError(t, disassocErr)
		}

		_, err := ec2Client.DeleteRouteTable(&ec2.DeleteRouteTableInput{
			RouteTableId: rt.RouteTableId,
		})
		require.NoError(t, err)
	}
}

func deleteSubnets(t *testing.T, vpcId string, region string) {
	ec2Client := terratest.NewEc2Client(t, region)
	vpcIDFilterName := "vpc-id"
	vpcIDFilter := ec2.Filter{Name: &vpcIDFilterName, Values: []*string{&vpcId}}

	subnetsOutput, err := ec2Client.DescribeSubnets(&ec2.DescribeSubnetsInput{Filters: []*ec2.Filter{&vpcIDFilter}})
	require.NoError(t, err)

	for _, subnet := range subnetsOutput.Subnets {
		_, err := ec2Client.DeleteSubnet(&ec2.DeleteSubnetInput{
			SubnetId: subnet.SubnetId,
		})
		require.NoError(t, err)
	}
}

func deleteInternetGateways(t *testing.T, vpcId string, region string) {
	ec2Client := terratest.NewEc2Client(t, region)
	vpcIDFilterName := "attachment.vpc-id"
	vpcIDFilter := ec2.Filter{Name: &vpcIDFilterName, Values: []*string{&vpcId}}

	igwOutput, err := ec2Client.DescribeInternetGateways(&ec2.DescribeInternetGatewaysInput{Filters: []*ec2.Filter{&vpcIDFilter}})
	require.NoError(t, err)

	for _, igw := range igwOutput.InternetGateways {

		_, detachErr := ec2Client.DetachInternetGateway(&ec2.DetachInternetGatewayInput{
			InternetGatewayId: igw.InternetGatewayId,
			VpcId:             aws.String(vpcId),
		})
		require.NoError(t, detachErr)

		_, err := ec2Client.DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
			InternetGatewayId: igw.InternetGatewayId,
		})
		require.NoError(t, err)
	}
}

func deleteVpc(t *testing.T, vpcId string, region string, sgId string) {
	ec2Client := terratest.NewEc2Client(t, region)
	destroySecurityGroup(t, sgId, region)
	deleteRouteTables(t, vpcId, region)
	deleteSubnets(t, vpcId, region)
	deleteInternetGateways(t, vpcId, region)

	_, err := ec2Client.DeleteVpc(&ec2.DeleteVpcInput{
		VpcId: aws.String(vpcId),
	})
	if err != nil {
		log.Println("Error deleting VPC: ", err)
	}
	require.NoError(t, err)
}
