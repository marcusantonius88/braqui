# SPEC-004 - Containerização e Ambiente Local

## Objetivo

Definir a estratégia de containerização e ambiente local do Braqui MVP utilizando Docker.

O foco desta spec é:
- padronizar ambiente de desenvolvimento;
- simplificar setup;
- facilitar deploy;
- melhorar experiência de desenvolvimento;
- garantir reproducibilidade do ambiente.

---

# Contexto

O Braqui utilizará:
- Go;
- PostgreSQL;
- Telegram webhook;
- scheduler;
- integrações externas;
- IA;
- variáveis de ambiente.

Sem containerização:
- onboarding fica inconsistente;
- setup local se torna frágil;
- diferenças de ambiente aumentam;
- debugging operacional fica mais difícil.

---

# Escopo

## O sistema deve:

- utilizar Docker;
- utilizar Docker Compose;
- suportar ambiente local completo;
- possuir containers desacoplados;
- permitir execução com único comando.

---

# Fora do Escopo

Esta spec NÃO contempla:

- Kubernetes;
- Docker Swarm;
- infraestrutura distribuída;
- autoscaling;
- service mesh;
- múltiplos ambientes complexos.

---

# Filosofia Arquitetural

A containerização deve:
- ser simples;
- previsível;
- rápida;
- amigável para desenvolvimento.

O objetivo NÃO é criar infraestrutura enterprise.

---

# Estratégia Monorepo

O ambiente local deverá iniciar inicialmente:

```text
/apps/api
```

A estrutura deve permitir evolução futura para:

```text
/apps
  /api
  /dashboard
  /admin
```

sem necessidade de reestruturação significativa.

---

# Containers Iniciais

Inicialmente:

- api
- postgres

---

# Estrutura Esperada

```text
/docker
  /postgres

/apps
  /api
    Dockerfile

docker-compose.yml
```

---

# Dockerfile

O projeto deve possuir:
- Dockerfile principal;
- build reproduzível;
- runtime padronizado.

---

# Estratégia Inicial

Inicialmente:
- multi-stage build simples.

---

# Exemplo Conceitual

```dockerfile
FROM golang:1.xx AS builder
```

---

# Docker Compose

O projeto deve possuir:

```text
docker-compose.yml
```

---

# Objetivo Principal

Permitir:

```bash
docker compose up
```

e ter:
- aplicação;
- PostgreSQL;
- ambiente funcional.

---

# Serviços Iniciais

## API

Responsável por:
- webhook Telegram;
- scheduler;
- APIs;
- domínio;
- integração IA.

---

## PostgreSQL

Responsável por:
- persistência local;
- desenvolvimento;
- testes locais.

---

# Rede

Containers devem:
- comunicar-se via rede interna do Docker.

---

# Volumes

O PostgreSQL deve utilizar volume persistente.

---

# Exemplo Conceitual

```yaml
volumes:
  postgres_data:
```

---

# Variáveis de Ambiente

O Compose deve suportar:

```text
.env
```

---

# Exemplos

```text
DATABASE_URL=
TELEGRAM_BOT_TOKEN=
GEMINI_API_KEY=
OPENWEATHER_API_KEY=
```

---

# Hot Reload

Inicialmente:
- opcional.

---

# Possíveis Ferramentas

Exemplos:
- air
- reflex

---

# Objetivo do Hot Reload

Melhorar:
- produtividade;
- desenvolvimento iterativo;
- integração com OpenCode.

---

# Healthcheck

Containers devem possuir:
- healthcheck simples.

---

# Exemplo

```text
GET /health
```

---

# Estratégia de Banco

Inicialmente:
- PostgreSQL único local.

---

# Migrações

O ambiente deve permitir:
- execução de migrations.

---

# Estratégia Inicial

Pode ser:
- manual;
OU
- container auxiliar futuramente.

---

# Logs

Containers devem:
- expor logs facilmente;
- facilitar debugging.

---

# Comandos Esperados

## Subir ambiente

```bash
docker compose up
```

---

## Derrubar ambiente

```bash
docker compose down
```

---

## Rebuild

```bash
docker compose up --build
```

---

# NÃO fazer

O MVP NÃO deve:
- criar múltiplos containers desnecessários;
- separar scheduler;
- criar infraestrutura distribuída;
- adicionar complexidade prematura.

---

# Estratégia de Desenvolvimento

O ambiente local deve:
- ser reproduzível;
- funcionar em diferentes máquinas;
- reduzir "works on my machine".

---

# Observabilidade

Inicialmente:
- logs simples via Docker.

---

# Segurança

## NÃO fazer

O projeto NÃO deve:
- commitar secrets;
- commitar `.env` real;
- expor credenciais em imagens.

---

# Critérios de Aceite

## Build

- aplicação builda corretamente via Docker.

---

## Ambiente local

- ambiente sobe corretamente com Docker Compose.

---

## Banco

- PostgreSQL funciona corretamente.

---

## Comunicação

- API consegue acessar PostgreSQL.

---

## Persistência

- volume persiste dados corretamente.

---

## Developer Experience

- setup local funciona com poucos comandos.

---

# Tratamento de Erros

## Falha de build

Logs devem:
- ser claros;
- facilitar troubleshooting.

---

## Banco indisponível

Aplicação deve:
- falhar claramente;
- registrar logs amigáveis.

---

# Requisitos Técnicos

## Deve existir

- Dockerfile;
- docker-compose.yml;
- volume PostgreSQL;
- variáveis de ambiente;
- rede Docker;
- documentação de setup.

---

# Dependências

Relaciona-se com:
- SPEC-001 - Bootstrap e Estrutura Inicial do Projeto
- SPEC-002 - Configuração e Gerenciamento de Ambiente
- SPEC-003 - Deploy e Infraestrutura Inicial
- SPEC-006 - Healthcheck e Diagnóstico Operacional

---

# Considerações Arquiteturais

## Docker como capability operacional

A containerização faz parte da:
- experiência de desenvolvimento;
- padronização operacional;
- evolução do projeto.

---

## Simplicidade primeiro

O MVP deve priorizar:
- poucos containers;
- setup rápido;
- baixa complexidade operacional.

---

## Compatibilidade com Monorepo

A estrutura deve:
- funcionar naturalmente dentro do monorepo;
- suportar novas aplicações no futuro;
- evitar acoplamento excessivo ao backend atual.

---

# Objetivo Real do MVP

O foco é:
- acelerar desenvolvimento;
- simplificar onboarding;
- facilitar deploy;
- melhorar experiência com IA-assisted development.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- containers separados;
- workers;
- Redis;
- filas;
- observabilidade avançada;
- ambientes multi-stage;
- Kubernetes;
- CI/CD avançado.

---

# Implementation Checklist

## Estrutura Docker

- [x] Dockerfile criado em `/apps/api` (SPEC-003)
- [x] Estrutura `/docker` criada
- [x] Estrutura `/docker/postgres` criada (com init.sql — extensão uuid-ossp)
- [x] docker-compose.yml criado

---

## Build

- [x] Multi-stage build implementado (SPEC-003)
- [x] Build local funcionando (SPEC-003)
- [x] Runtime da aplicação configurado (SPEC-003)

---

## Ambiente Local

- [x] Ambiente completo sobe com `docker compose up`
- [x] API inicia corretamente
- [x] PostgreSQL inicia corretamente

---

## Banco de Dados

- [x] Container PostgreSQL configurado
- [x] Volume persistente configurado
- [ ] Persistência validada (requer SPEC-005 — operações de banco)

---

## Rede

- [x] Rede interna Docker configurada (bridge `braqui`)
- [x] Comunicação API ↔ PostgreSQL funcionando (ambos na mesma rede)

---

## Variáveis de Ambiente

- [x] Integração com `.env` implementada (Docker Compose carrega root `.env`)
- [x] Variáveis propagadas para containers
- [x] `.env.example` documentado (SPEC-002)

---

## Healthcheck

- [x] Healthcheck da API configurado (SPEC-003 — Dockerfile HEALTHCHECK)
- [x] Healthcheck do PostgreSQL configurado (pg_isready)

---

## Logs

- [x] Logs acessíveis via Docker
- [x] Logs de startup implementados
- [x] Logs de erro implementados

---

## Migrações

- [ ] Estratégia de migrations definida (pendente — manual por enquanto)
- [ ] Execução local de migrations funcionando (pendente — manual)

---

## Segurança

- [x] `.env` ignorado pelo Git
- [x] Secrets não expostos em imagens
- [x] Secrets não expostos nos logs

---

## Developer Experience

- [ ] Setup documentado (documentação de setup ainda não escrita)
- [x] Ambiente reproduzível validado
- [x] Workflow compatível com OpenCode

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com evolução futura
- [x] Estrutura compatível com IA-assisted development