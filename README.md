# DigitalOcean Droplet com Pulumi - Componentizado

Este projeto demonstra como criar droplets DigitalOcean usando Pulumi e Go com componentes reutilizáveis.

## 🧩 Arquitetura Componentizada

O projeto foi refatorado para usar componentes Pulumi, tornando-o mais modular e reutilizável:

### Componentes Disponíveis

1. **Droplet Component** (`components/droplet.go`)
   - Encapsula a criação de um droplet DigitalOcean
   - Valores padrão inteligentes
   - Tags automáticas
   - Configuração flexível

2. **Infrastructure Component** (`components/infrastructure.go`)
   - Gerencia múltiplos droplets
   - Organização por ambiente (dev, prod, staging)
   - Tags padronizadas por projeto
   - Configuração de alta nivel

## 📁 Estrutura do Projeto

```
digitalocean-spaces/
├── main.go                          # Aplicação principal
├── components/
│   ├── droplet.go                   # Componente de droplet
│   └── infrastructure.go            # Componente de infraestrutura
├── examples/
│   ├── single_droplet.go           # Exemplo de droplet único
│   └── multiple_droplets.go        # Exemplo de múltiplos droplets
├── Pulumi.yaml                     # Configuração do projeto
├── Pulumi.dev.yaml                 # Configurações do stack
└── README.md                       # Esta documentação
```

## 🚀 Uso Básico

### 1. Droplet Único

```go
droplet, err := components.NewDroplet(ctx, "my-droplet", &components.DropletArgs{
    Name:   "web-server",
    Region: "nyc3",
    Size:   "s-1vcpu-1gb",
    Image:  "ubuntu-22-04-x64",
    Tags:   []string{"web", "production"},
})
```

### 2. Infraestrutura Completa

```go
infrastructure, err := components.NewInfrastructure(ctx, "main-infra", &components.InfrastructureArgs{
    Environment: "production",
    Project:     "webapp",
    Droplets: []components.DropletConfig{
        {
            Name:   "web-server-1",
            Region: "nyc3",
            Size:   "s-2vcpu-4gb",
            Tags:   []string{"web", "frontend"},
        },
        {
            Name:   "api-server",
            Region: "nyc3", 
            Size:   "s-4vcpu-8gb",
            Tags:   []string{"api", "backend"},
        },
    },
})
```

## ⚙️ Configuração

### Variáveis de Ambiente
```bash
export DIGITALOCEAN_TOKEN=your_token_here
```

### Configuração Pulumi
```bash
# Configurações básicas
pulumi config set dropletName my-server
pulumi config set region nyc3
pulumi config set size s-1vcpu-1gb
pulumi config set image ubuntu-22-04-x64

# Configurações de organização
pulumi config set environment dev
pulumi config set project my-project
```

## 🎯 Exemplos de Uso

### Deploy Simples
```bash
# Deploy com configurações padrão
pulumi up
```

### Deploy Multi-Ambiente
Veja `examples/multiple_droplets.go` para um exemplo completo de como configurar diferentes ambientes (dev, staging, prod) com diferentes configurações.

## 🔧 Personalização

### Tamanhos de Droplet Disponíveis
- `s-1vcpu-1gb` (1 vCPU, 1GB RAM) - $6/mês
- `s-1vcpu-2gb` (1 vCPU, 2GB RAM) - $12/mês  
- `s-2vcpu-2gb` (2 vCPUs, 2GB RAM) - $18/mês
- `s-2vcpu-4gb` (2 vCPUs, 4GB RAM) - $24/mês
- `s-4vcpu-8gb` (4 vCPUs, 8GB RAM) - $48/mês

### Regiões Disponíveis
- **Estados Unidos**: `nyc1`, `nyc2`, `nyc3`, `sfo1`, `sfo2`, `sfo3`
- **Europa**: `ams2`, `ams3`, `fra1`, `lon1`
- **Ásia**: `sgp1`, `blr1`
- **Canadá**: `tor1`

### Imagens Suportadas
- `ubuntu-22-04-x64` (Ubuntu 22.04 LTS)
- `ubuntu-20-04-x64` (Ubuntu 20.04 LTS)
- `debian-11-x64` (Debian 11)
- `centos-stream-9-x64` (CentOS Stream 9)
- `fedora-39-x64` (Fedora 39)

## 🧪 Testando

```bash
# Compilar o projeto
go build

# Validar configuração
pulumi preview

# Deploy
pulumi up

# Limpeza
pulumi destroy
```

## 🔄 Vantagens da Componentização

1. **Reutilização**: Componentes podem ser usados em múltiplos projetos
2. **Encapsulamento**: Lógica complexa escondida atrás de interfaces simples
3. **Padronização**: Tags e configurações consistentes
4. **Manutenibilidade**: Mudanças centralizadas nos componentes
5. **Testabilidade**: Componentes podem ser testados independentemente
6. **Composição**: Combine componentes para criar infraestruturas complexas

## 📖 Próximos Passos

- Adicionar componentes para Load Balancers
- Implementar componentes de rede (VPC, Firewall)
- Criar templates para diferentes tipos de aplicação
- Adicionar testes automatizados para os componentes