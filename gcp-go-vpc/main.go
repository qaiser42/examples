package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create an empty VPC Network
		network, err := compute.NewNetwork(ctx, "asd", &compute.NetworkArgs{
			// AutoCreateSubnetworks = false makes it a "Custom Mode" VPC (empty)
			// This is recommended for production to control IP ranges.
			AutoCreateSubnetworks: pulumi.Bool(false),
			Description:           pulumi.String("An empty custom VPC created via Pulumi"),
			// RoutingMode can be REGIONAL or GLOBAL
			RoutingMode: pulumi.String("REGIONAL"),
		})
		if err != nil {
			return err
		}

		// Export the Network ID and Name
		ctx.Export("networkId", network.ID())

		return nil
	})
}
