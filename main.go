package main

import (
	"digitalocean-spaces/components"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Get configuration values
		cfg := config.New(ctx, "")
		
		// Infrastructure configuration
		environment := cfg.Get("environment")
		if environment == "" {
			environment = "dev"
		}
		
		project := cfg.Get("project")
		if project == "" {
			project = "digitalocean-spaces"
		}

		// Droplet configurations
		dropletConfigs := []components.DropletConfig{}

		// Get droplet configuration from config or use defaults
		dropletName := cfg.Get("dropletName")
		if dropletName == "" {
			dropletName = "test-sloth"
		}
		
		region := cfg.Get("region")
		if region == "" {
			region = "nyc3"
		}
		
		size := cfg.Get("size")
		if size == "" {
			size = "s-1vcpu-1gb"
		}
		
		image := cfg.Get("image")
		if image == "" {
			image = "ubuntu-22-04-x64"
		}

		// Add main droplet configuration
		dropletConfigs = append(dropletConfigs, components.DropletConfig{
			Name:   dropletName,
			Region: region,
			Size:   size,
			Image:  image,
			Tags:   []string{"web", "main"},
		})

		// Create infrastructure component
		infrastructure, err := components.NewInfrastructure(ctx, "main-infrastructure", &components.InfrastructureArgs{
			Environment: environment,
			Project:     project,
			Droplets:    dropletConfigs,
		})
		if err != nil {
			return err
		}

		// Export infrastructure information
		for name, droplet := range infrastructure.Droplets {
			ctx.Export(name+"-name", droplet.Name)
			ctx.Export(name+"-id", droplet.ID)
			ctx.Export(name+"-ipv4Address", droplet.IPv4Address)
			ctx.Export(name+"-ipv4AddressPrivate", droplet.IPv4AddressPrivate)
			ctx.Export(name+"-status", droplet.Status)
		}

		return nil
	})
}
