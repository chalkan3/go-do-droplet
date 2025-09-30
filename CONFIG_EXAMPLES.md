# Exemplos de Configuração para Pulumi

## Configuração Básica (Single Droplet)
```bash
# Droplet configuration
pulumi config set dropletName "my-web-server"
pulumi config set region "nyc3"
pulumi config set size "s-1vcpu-1gb"
pulumi config set image "ubuntu-22-04-x64"

# Organization
pulumi config set environment "dev"
pulumi config set project "webapp"
```

## Configuração para Produção
```bash
# Droplet configuration
pulumi config set dropletName "webapp-prod-main"
pulumi config set region "nyc3"
pulumi config set size "s-2vcpu-4gb"
pulumi config set image "ubuntu-22-04-x64"

# Organization
pulumi config set environment "production"
pulumi config set project "webapp"
```

## Configuração para Desenvolvimento
```bash
# Droplet configuration
pulumi config set dropletName "webapp-dev-test"
pulumi config set region "nyc3"
pulumi config set size "s-1vcpu-1gb"
pulumi config set image "ubuntu-22-04-x64"

# Organization
pulumi config set environment "development"
pulumi config set project "webapp"
```

## Configuração Multi-Regional
```bash
# Primary region
pulumi config set dropletName "webapp-east"
pulumi config set region "nyc3"
pulumi config set size "s-2vcpu-2gb"

# Environment
pulumi config set environment "staging"
pulumi config set project "webapp"
```

## Limpar Configurações
```bash
# Remove all configurations
pulumi config rm dropletName
pulumi config rm region
pulumi config rm size
pulumi config rm image
pulumi config rm environment
pulumi config rm project
```