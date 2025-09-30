package main

import (
	"digitalocean-spaces/components"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Example of creating a single droplet with custom configuration
func ExampleSingleDroplet() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a single droplet component
		droplet, err := components.NewDroplet(ctx, "my-custom-droplet", &components.DropletArgs{
			Name:   "custom-app-server",
			Region: "fra1",
			Size:   "s-2vcpu-2gb",
			Image:  "ubuntu-22-04-x64",
			Tags:   []string{"app", "custom", "pulumi"},
		})
		if err != nil {
			return err
		}

		// Export droplet information
		ctx.Export("dropletName", droplet.Name)
		ctx.Export("dropletId", droplet.ID)
		ctx.Export("ipv4Address", droplet.IPv4Address)
		ctx.Export("ipv4AddressPrivate", droplet.IPv4AddressPrivate)
		ctx.Export("status", droplet.Status)

		return nil
	})
}