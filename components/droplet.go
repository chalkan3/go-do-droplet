package components

import (
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DropletArgs defines the arguments for the Droplet component
type DropletArgs struct {
	Name   string
	Region string
	Size   string
	Image  string
	Tags   []string
}

// Droplet represents a reusable DigitalOcean droplet component
type Droplet struct {
	pulumi.ResourceState

	Droplet              *digitalocean.Droplet
	Name                 pulumi.StringOutput `pulumi:"name"`
	ID                   pulumi.IDOutput     `pulumi:"id"`
	IPv4Address          pulumi.StringOutput `pulumi:"ipv4Address"`
	IPv4AddressPrivate   pulumi.StringOutput `pulumi:"ipv4AddressPrivate"`
	Status               pulumi.StringOutput `pulumi:"status"`
}

// NewDroplet creates a new Droplet component
func NewDroplet(ctx *pulumi.Context, name string, args *DropletArgs, opts ...pulumi.ResourceOption) (*Droplet, error) {
	component := &Droplet{}
	
	err := ctx.RegisterComponentResource("digitalocean:droplet:Droplet", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Set default values
	dropletName := args.Name
	if dropletName == "" {
		dropletName = "default-droplet"
	}

	region := args.Region
	if region == "" {
		region = "nyc3"
	}

	size := args.Size
	if size == "" {
		size = "s-1vcpu-1gb"
	}

	image := args.Image
	if image == "" {
		image = "ubuntu-22-04-x64"
	}

	// Create the droplet
	droplet, err := digitalocean.NewDroplet(ctx, name+"-droplet", &digitalocean.DropletArgs{
		Name:   pulumi.String(dropletName),
		Region: pulumi.String(region),
		Size:   pulumi.String(size),
		Image:  pulumi.String(image),
		Tags:   pulumi.ToStringArray(args.Tags),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Set component properties
	component.Droplet = droplet
	component.Name = droplet.Name
	component.ID = droplet.ID()
	component.IPv4Address = droplet.Ipv4Address
	component.IPv4AddressPrivate = droplet.Ipv4AddressPrivate
	component.Status = droplet.Status

	// Register outputs
	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"name":                droplet.Name,
		"id":                  droplet.ID(),
		"ipv4Address":         droplet.Ipv4Address,
		"ipv4AddressPrivate":  droplet.Ipv4AddressPrivate,
		"status":              droplet.Status,
	}); err != nil {
		return nil, err
	}

	return component, nil
}