package main

import (
	alb "github.com/jfr992/pulumi-packages/alb"
	asg "github.com/jfr992/pulumi-packages/asg"
	network "github.com/jfr992/pulumi-packages/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		configFile := "infra.yaml"

		userdata := "userdata.sh"

		vpcID, privateSubnetIds, publicSubnetIds, err := network.CreateNetwork(ctx, configFile)
		if err != nil {
			return err
		}

		targetGroupArn, securityGroupID, err := alb.CreateALB(ctx, configFile, vpcID, privateSubnetIds)
		if err != nil {
			return err
		}

		err = asg.CreateASG(ctx, configFile, userdata, vpcID, privateSubnetIds, targetGroupArn, securityGroupID)
		if err != nil {
			return err
		}

		ctx.Export("vpcID", vpcID)
		ctx.Export("privateSubnetIds", privateSubnetIds)
		ctx.Export("publicSubnetIds", publicSubnetIds)
		ctx.Export("targetGroupArn", targetGroupArn)
		ctx.Export("securityGroupID", securityGroupID)

		return nil
	})
}
