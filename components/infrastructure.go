package components

import (
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// InfrastructureArgs defines the arguments for the Infrastructure component
type InfrastructureArgs struct {
	Environment string
	Project     string
	Droplets    []DropletConfig
}

// DropletConfig defines configuration for individual droplets
type DropletConfig struct {
	Name   string
	Region string
	Size   string
	Image  string
	Tags   []string
}

// Infrastructure represents a complete infrastructure setup
type Infrastructure struct {
	pulumi.ResourceState

	Droplets map[string]*Droplet
}

// NewInfrastructure creates a new Infrastructure component
func NewInfrastructure(ctx *pulumi.Context, name string, args *InfrastructureArgs, opts ...pulumi.ResourceOption) (*Infrastructure, error) {
	component := &Infrastructure{
		Droplets: make(map[string]*Droplet),
	}

	err := ctx.RegisterComponentResource("digitalocean:infrastructure:Infrastructure", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Default environment and project
	environment := args.Environment
	if environment == "" {
		environment = "dev"
	}

	project := args.Project
	if project == "" {
		project = "default"
	}

	// Create droplets
	for i, dropletConfig := range args.Droplets {
		dropletName := dropletConfig.Name
		if dropletName == "" {
			dropletName = fmt.Sprintf("%s-%s-droplet-%d", project, environment, i)
		}

		// Add common tags
		tags := append(dropletConfig.Tags, environment, project, "pulumi-managed")

		componentName := fmt.Sprintf("%s-droplet-%d", name, i)
		
		droplet, err := NewDroplet(ctx, componentName, &DropletArgs{
			Name:   dropletName,
			Region: dropletConfig.Region,
			Size:   dropletConfig.Size,
			Image:  dropletConfig.Image,
			Tags:   tags,
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}

		component.Droplets[dropletName] = droplet
	}

	// Register outputs
	outputs := pulumi.Map{}
	for name, droplet := range component.Droplets {
		outputs[name] = pulumi.Map{
			"name":                droplet.Name,
			"id":                  droplet.ID,
			"ipv4Address":         droplet.IPv4Address,
			"ipv4AddressPrivate":  droplet.IPv4AddressPrivate,
			"status":              droplet.Status,
		}
	}

	if err := ctx.RegisterResourceOutputs(component, outputs); err != nil {
		return nil, err
	}

	return component, nil
}