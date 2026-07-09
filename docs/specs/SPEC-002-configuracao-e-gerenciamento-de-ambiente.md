# SPEC-002 - Configuração e Gerenciamento de Ambiente

## Objetivo

Definir a estratégia de configuração do Braqui MVP para garantir:
- simplicidade operacional;
- segurança básica;
- facilidade de deploy;
- padronização de ambientes;
- compatibilidade com desenvolvimento local e cloud.

---

# Contexto

O Braqui utilizará:
- Telegram API;
- Gemini API;
- OpenWeather API;
- PostgreSQL;
- scheduler;
- webhooks.

O sistema precisará gerenciar:
- secrets;
- URLs;
- tokens;
- parâmetros de execução;
- configurações por ambiente.

---

# Escopo

## O sistema deve:

- carregar variáveis de ambiente;
- validar configurações obrigatórias;
- suportar múltiplos ambientes;
- evitar configurações hardcoded.

---

# Fora do Escopo

Esta spec NÃO contempla:

- secret manager enterprise;
- vault;
- rotação automática de segredo;
- feature flags avançadas;
- configuração distribuída;
- service discovery.

---

# Filosofia Arquitetural

Configuração deve ser:
- explícita;
- previsível;
- simples;
- centralizada.

O sistema NÃO deve:
- depender de arquivos mágicos;
- espalhar configuração pelo código;
- possuir valores sensíveis hardcoded.

---

# Ambientes Iniciais

Inicialmente:

```text
- local
- development
- production
```

---

# Estratégia Inicial

Inicialmente:
- variáveis de ambiente.

---

# Exemplo

```text
TELEGRAM_BOT_TOKEN=
DATABASE_URL=
GEMINI_API_KEY=
OPENWEATHER_API_KEY=
APP_ENV=
PORT=
```

---

# Configurações Obrigatórias

## Telegram

```text
TELEGRAM_BOT_TOKEN
```

---

## Banco

```text
DATABASE_URL
```

---

## Ambiente

```text
APP_ENV
```

---

# Configurações Opcionais

## IA

```text
GEMINI_API_KEY
```

---

## Clima

```text
OPENWEATHER_API_KEY
```

---

## Scheduler

```text
SCHEDULER_ENABLED
```

---

# Estratégia de Inicialização

Na inicialização:
- validar configs obrigatórias;
- falhar rapidamente em caso de ausência.

---

# Exemplo de Erro

```text
missing required environment variable: TELEGRAM_BOT_TOKEN
```

---

# Estrutura Esperada

## Package Config

Responsável por:
- carregar variáveis;
- validar configuração;
- expor objeto central de config.

---

# Estrutura Esperada no Monorepo

```text
/apps
  /api
    /internal
      /infra
        /config
```

---

# Exemplo Conceitual

```go
type Config struct {
    AppEnv string

    TelegramBotToken string

    DatabaseURL string

    GeminiAPIKey string

    OpenWeatherAPIKey string
}
```

---

# Estratégia Local

Inicialmente:
- `.env`

---

# Exemplo

```text
.env
.env.example
```

---

# Segurança

## NÃO fazer

O projeto NÃO deve:
- commitar secrets;
- commitar tokens reais;
- expor credenciais em logs.

---

# .env.example

O repositório deve possuir:
- exemplo de configuração;
- sem valores reais.

---

# Configuração por Ambiente

O sistema deve permitir:
- valores diferentes por ambiente;
- deploy simples em cloud providers.

---

# Observabilidade

Na inicialização:
- logar ambiente atual;
- NÃO logar segredos.

---

# Exemplo permitido

```text
starting braqui in production mode
```

---

# Exemplo proibido

```text
telegram token: 123abc...
```

---

# Critérios de Aceite

## Carregamento

- sistema carrega configs corretamente.

---

## Validação

- sistema falha caso configs obrigatórias estejam ausentes.

---

## Segurança

- secrets não aparecem em logs.

---

## Ambientes

- sistema suporta múltiplos ambientes.

---

# Tratamento de Erros

## Variável ausente

Sistema deve:
- falhar rapidamente;
- exibir mensagem clara.

---

## Configuração inválida

Sistema deve:
- impedir inicialização;
- registrar erro amigável.

---

# Requisitos Técnicos

## Deve existir

- package config;
- loader de env;
- validação de configuração;
- `.env.example`;
- centralização de configs.

---

# Dependências

Relaciona-se com:
- SPEC-001 - Bootstrap e Estrutura Inicial do Projeto

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- Consul;
- Vault;
- Kubernetes ConfigMap;
- configuração distribuída.

---

## Configuração centralizada

Toda configuração deve passar:
- por um único ponto;
- por validação explícita.

---

## Compatibilidade com Monorepo

A estratégia de configuração deve:
- funcionar bem dentro da estrutura monorepo;
- manter baixo acoplamento;
- facilitar futuras aplicações dentro de `/apps`.

---

# Objetivo Real do MVP

O foco é:
- facilitar setup;
- reduzir erros operacionais;
- simplificar deploy;
- manter segurança básica.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- secret manager;
- feature flags;
- config remota;
- multi environment avançado;
- rotação automática de segredos;
- configuração dinâmica.

---

# Implementation Checklist

## Estrutura Base

- [x] Package config criado
- [x] Estrutura centralizada de configuração criada
- [x] Config struct implementada
- [x] Loader de variáveis implementado

---

## Variáveis de Ambiente

- [x] Suporte a `.env` implementado (parser próprio, sem dependências)
- [x] Suporte a `.env.example` implementado
- [x] Variáveis obrigatórias definidas (TELEGRAM_BOT_TOKEN, DATABASE_URL, APP_ENV)
- [x] Variáveis opcionais definidas (GEMINI_API_KEY, OPENWEATHER_API_KEY, SCHEDULER_ENABLED, PORT)

---

## Validação

- [x] Validação de variáveis obrigatórias implementada
- [x] Fail fast na inicialização implementado
- [x] Mensagens de erro amigáveis implementadas

---

## Segurança

- [x] Secrets removidos dos logs (apenas `starting braqui in X mode` é logado)
- [x] Tokens protegidos contra exposição
- [x] `.env` adicionado ao `.gitignore`

---

## Ambientes

- [x] Suporte a ambiente local implementado
- [x] Suporte a development implementado
- [x] Suporte a production implementado

---

## Observabilidade

- [x] Log de inicialização implementado
- [x] Log de ambiente atual implementado

---

## Testes

- [x] Testes unitários do loader implementados
- [x] Testes de validação implementados