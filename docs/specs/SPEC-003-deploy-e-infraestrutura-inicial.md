# SPEC-003 - Deploy e Infraestrutura Inicial

## Objetivo

Definir a estratégia inicial de deploy e infraestrutura do Braqui MVP.

O foco desta spec é:
- simplicidade operacional;
- baixo custo;
- facilidade de manutenção;
- velocidade de entrega;
- deploy contínuo simples.

---

# Contexto

O Braqui MVP será:
- pequeno;
- monolítico inicialmente;
- orientado a validação rápida;
- com baixa carga inicial.

O sistema NÃO precisa inicialmente de:
- infraestrutura distribuída;
- Kubernetes;
- alta escalabilidade;
- múltiplos serviços.

---

# Escopo

## O sistema deve:

- possuir deploy simples;
- suportar webhook do Telegram;
- possuir banco PostgreSQL;
- executar scheduler;
- suportar variáveis de ambiente;
- permitir evolução incremental.

---

# Fora do Escopo

Esta spec NÃO contempla:

- Kubernetes;
- autoscaling avançado;
- multi-region;
- service mesh;
- infraestrutura distribuída;
- alta disponibilidade enterprise;
- multi cloud.

---

# Filosofia Arquitetural

Infraestrutura deve:
- ser simples;
- barata;
- previsível;
- fácil de operar.

O objetivo NÃO é otimizar escala prematuramente.

---

# Estratégia Inicial

Inicialmente:
- uma única aplicação backend.

A aplicação será responsável por:
- Telegram webhook;
- scheduler;
- domínio;
- integrações;
- persistência.

---

# Estratégia Monorepo

O deploy inicialmente utilizará apenas:

```text
/apps/api
```

A arquitetura deve permitir evolução futura para:
- dashboard;
- admin panel;
- aplicações adicionais.

---

# Hospedagem Inicial

Possíveis opções:

- Render
- Railway
- Fly.io

---

# Critérios de Escolha

Priorizar:
- simplicidade;
- free tier;
- deploy rápido;
- suporte a variáveis de ambiente;
- PostgreSQL gerenciado.

---

# Banco de Dados

Inicialmente:
- PostgreSQL gerenciado.

---

# Deploy

Inicialmente:
- deploy automático via GitHub.

---

# Fluxo Esperado

1. push no GitHub
2. pipeline executa
3. build da aplicação
4. deploy automático

---

# Estratégia de Build

Inicialmente:
- Docker.

---

# Estrutura Esperada

```text
/apps/api
  Dockerfile
```

---

# Webhook Telegram

A infraestrutura deve suportar:
- HTTPS público;
- endpoint acessível externamente.

---

# Scheduler

O scheduler será executado:
- dentro da própria aplicação inicialmente.

---

# Estratégia Inicial de Runtime

Single process:
- API;
- scheduler;
- webhook.

---

# NÃO fazer

O MVP NÃO deve:
- separar workers;
- criar múltiplos serviços;
- introduzir filas;
- usar Kubernetes;
- usar Terraform inicialmente.

---

# Observabilidade

Inicialmente:
- logs da plataforma;
- logs estruturados da aplicação.

---

# Segurança

## Variáveis sensíveis

Devem permanecer:
- em environment variables;
- fora do código.

---

# HTTPS

Obrigatório para:
- Telegram webhook.

---

# Healthcheck

A aplicação deve expor endpoint simples.

---

# Exemplo

```text
GET /health
```

---

# Resposta esperada

```json
{
  "status": "ok"
}
```

---

# Critérios de Aceite

## Deploy

- aplicação sobe corretamente.

---

## Webhook

- Telegram consegue acessar webhook.

---

## Banco

- aplicação conecta no PostgreSQL.

---

## Scheduler

- jobs executam corretamente.

---

## Healthcheck

- endpoint responde corretamente.

---

# Tratamento de Erros

## Falha de conexão banco

Sistema deve:
- falhar rapidamente;
- registrar logs claros.

---

## Variáveis ausentes

Sistema deve:
- impedir inicialização.

---

# Requisitos Técnicos

## Deve existir

- Dockerfile;
- configuração de ambiente;
- healthcheck;
- deploy configurável;
- conexão PostgreSQL;
- webhook público HTTPS.

---

# Dependências

Relaciona-se com:
- SPEC-001 - Bootstrap e Estrutura Inicial do Projeto
- SPEC-002 - Configuração e Gerenciamento de Ambiente

---

# Considerações Arquiteturais

## Monólito primeiro

O MVP deve priorizar:
- simplicidade;
- velocidade;
- facilidade operacional.

---

## Escala depois

Escalabilidade só deve ser considerada:
- após validação real;
- após crescimento consistente.

---

## Compatibilidade com Monorepo

A infraestrutura deve:
- funcionar bem dentro da estratégia monorepo;
- suportar futuras aplicações;
- manter simplicidade operacional.

---

# Objetivo Real do MVP

O foco é:
- colocar o Braqui online rapidamente;
- validar produto;
- reduzir complexidade operacional;
- permitir evolução incremental.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- workers separados;
- filas;
- autoscaling;
- Kubernetes;
- observabilidade avançada;
- multi-region;
- CDN;
- edge workloads;
- infraestrutura distribuída.

---

# Implementation Checklist

## Estrutura Base

- [x] Estrutura de deploy inicial criada
- [x] Estrutura compatível com monorepo definida
- [x] Estrutura `/apps/api` preparada para deploy

---

## Build

- [x] Dockerfile criado (multi-stage: golang:1.23-alpine → alpine:3.20)
- [x] Build da aplicação funcionando
- [x] Runtime da aplicação configurado

---

## Deploy

- [ ] Deploy automático via GitHub configurado (requer configurar provider: Render/Railway/Fly.io)
- [x] Pipeline inicial criado (.github/workflows/ci.yml — build + test + vet)
- [ ] Estratégia de deploy documentada (pendente — definir provider primeiro)

---

## Infraestrutura

- [ ] PostgreSQL configurado (SPEC-005)
- [x] Variáveis de ambiente configuradas (SPEC-002)
- [ ] HTTPS configurado (requer deploy real)
- [ ] Endpoint público configurado (requer deploy real)

---

## Telegram

- [ ] Webhook Telegram configurado (SPEC-009)
- [ ] Endpoint webhook exposto publicamente (SPEC-009)

---

## Scheduler

- [ ] Scheduler integrado ao runtime principal (SPEC-018)
- [ ] Execução de jobs funcionando (SPEC-018)

---

## Healthcheck

- [x] Endpoint `/health` criado (retorna `{"status":"ok"}`)
- [x] Healthcheck retornando status corretamente

---

## Observabilidade

- [x] Logs básicos configurados
- [ ] Logs estruturados configurados (SPEC-007)
- [x] Logs de inicialização implementados (SPEC-002)

---

## Segurança

- [x] Secrets protegidos
- [x] Variáveis sensíveis removidas do código
- [x] Configuração segura de ambiente implementada

---

## Qualidade

- [x] Deploy reproduzível funcionando (Docker build verificado)
- [x] Estrutura compatível com evolução incremental
- [x] Infraestrutura compatível com IA-assisted development