package main

import (
	"digitalocean-spaces/components"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Example of using components to create multiple droplets for different environments
func ExampleMultipleDroplets() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create development infrastructure
		devInfra, err := components.NewInfrastructure(ctx, "dev-infrastructure", &components.InfrastructureArgs{
			Environment: "dev",
			Project:     "webapp",
			Droplets: []components.DropletConfig{
				{
					Name:   "webapp-dev-web",
					Region: "nyc3",
					Size:   "s-1vcpu-1gb",
					Image:  "ubuntu-22-04-x64",
					Tags:   []string{"web", "frontend"},
				},
				{
					Name:   "webapp-dev-api",
					Region: "nyc3",
					Size:   "s-1vcpu-2gb",
					Image:  "ubuntu-22-04-x64",
					Tags:   []string{"api", "backend"},
				},
			},
		})
		if err != nil {
			return err
		}

		// Create production infrastructure
		prodInfra, err := components.NewInfrastructure(ctx, "prod-infrastructure", &components.InfrastructureArgs{
			Environment: "prod",
			Project:     "webapp",
			Droplets: []components.DropletConfig{
				{
					Name:   "webapp-prod-web-1",
					Region: "nyc3",
					Size:   "s-2vcpu-4gb",
					Image:  "ubuntu-22-04-x64",
					Tags:   []string{"web", "frontend", "loadbalanced"},
				},
				{
					Name:   "webapp-prod-web-2",
					Region: "sfo3",
					Size:   "s-2vcpu-4gb",
					Image:  "ubuntu-22-04-x64",
					Tags:   []string{"web", "frontend", "loadbalanced"},
				},
				{
					Name:   "webapp-prod-api",
					Region: "nyc3",
					Size:   "s-4vcpu-8gb",
					Image:  "ubuntu-22-04-x64",
					Tags:   []string{"api", "backend"},
				},
			},
		})
		if err != nil {
			return err
		}

		// Export development infrastructure
		for name, droplet := range devInfra.Droplets {
			ctx.Export("dev-"+name+"-ip", droplet.IPv4Address)
		}

		// Export production infrastructure
		for name, droplet := range prodInfra.Droplets {
			ctx.Export("prod-"+name+"-ip", droplet.IPv4Address)
		}

		return nil
	})
}