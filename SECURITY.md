# Security Guidelines

## 🔐 Protegendo Credenciais

### ⚠️ NUNCA faça commit de:
- Tokens de API do DigitalOcean
- Chaves SSH privadas
- Senhas ou segredos
- Arquivos `.env` com credenciais

### ✅ Como configurar credenciais de forma segura:

#### 1. Variáveis de Ambiente
```bash
# Adicione ao seu ~/.bashrc ou ~/.zshrc
export DIGITALOCEAN_TOKEN="your_token_here"

# Ou use um arquivo .env (não commitado)
echo "DIGITALOCEAN_TOKEN=your_token_here" > .env
source .env
```

#### 2. Pulumi Config com Secrets
```bash
# Para dados sensíveis, use --secret
pulumi config set digitalocean:token --secret

# Para dados não sensíveis
pulumi config set region nyc3
pulumi config set environment dev
```

#### 3. CI/CD Secrets
- **GitHub Actions**: Use GitHub Secrets
- **GitLab CI**: Use GitLab Variables
- **CircleCI**: Use CircleCI Environment Variables

### 🛡️ Arquivo de Exemplo (.env.example)
```bash
# DigitalOcean API Token (obtenha em: https://cloud.digitalocean.com/account/api/tokens)
DIGITALOCEAN_TOKEN=your_token_here

# Opcional: Configurações específicas do Spaces (se usar Spaces)
# SPACES_ACCESS_KEY=your_spaces_access_key
# SPACES_SECRET_KEY=your_spaces_secret_key
```

### 📋 Checklist de Segurança
- [ ] Token não está commitado no código
- [ ] .env está no .gitignore
- [ ] Credenciais são passadas via variáveis de ambiente
- [ ] Secrets do Pulumi estão marcados como --secret
- [ ] CI/CD usa secrets seguros

### 🔍 Como verificar se não há vazamentos
```bash
# Procurar por tokens no repositório
git log --all -p | grep -i "dop_v1\|token\|secret"

# Verificar arquivos atuais
grep -r "dop_v1\|secret\|token" . --exclude-dir=.git
```