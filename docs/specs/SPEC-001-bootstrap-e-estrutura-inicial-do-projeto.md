# SPEC-001 - Bootstrap e Estrutura Inicial do Projeto

## Objetivo

Definir a estrutura inicial do projeto Braqui e estabelecer os padrões arquiteturais fundamentais do código.

Esta spec será responsável por:
- bootstrap inicial;
- organização de diretórios;
- modularização;
- boundaries arquiteturais;
- setup inicial Go;
- fundação do desenvolvimento incremental;
- estratégia de monorepo.

---

# Contexto

O Braqui será desenvolvido utilizando:
- Spec Driven Development (SDD);
- IA-assisted development;
- OpenCode;
- arquitetura modular;
- Clean Architecture leve;
- abordagem hexagonal leve.

Sem uma estrutura inicial clara:
- o código tende a degradar rapidamente;
- a IA tende a criar inconsistências;
- boundaries arquiteturais ficam difusos;
- manutenção se torna difícil.

---

# Escopo

## O sistema deve possuir:

- estrutura base do projeto;
- organização modular;
- setup inicial Go;
- diretórios padronizados;
- boundaries arquiteturais explícitos;
- estrutura monorepo preparada para evolução futura.

---

# Fora do Escopo

Esta spec NÃO contempla:

- regras de negócio;
- integrações externas;
- persistência;
- Telegram;
- IA;
- scheduler;
- implementação de features.

---

# Filosofia Arquitetural

A estrutura do projeto deve priorizar:
- simplicidade;
- clareza;
- baixo acoplamento;
- alta coesão;
- modularidade explícita;
- evolução incremental.

---

# Estratégia Arquitetural

O Braqui seguirá:
- monorepo;
- monólito modular inicialmente;
- arquitetura em camadas leves;
- boundaries explícitos;
- domínio desacoplado.

---

# Estratégia de Repositório

O Braqui utilizará:
- monorepo;
- organização orientada a aplicações;
- separação explícita entre apps e documentação.

---

# Objetivos da Estratégia

A estrutura deve:
- facilitar IA-assisted development;
- facilitar evolução incremental;
- facilitar compartilhamento de contexto;
- suportar futuras aplicações;
- manter simplicidade operacional.

---

# Estrutura Inicial Esperada

```text
/apps
/docs
/docker
/scripts
/tests
```

---

# Estrutura Apps

Responsável pelas aplicações do monorepo.

---

# Estrutura Inicial Esperada

```text
/apps
  /api
```

---

# Evoluções Futuras Possíveis

Fora do MVP:

```text
/apps
  /api
  /dashboard
  /admin
```

---

# Estrutura da API

Responsável pela aplicação backend principal.

---

# Estrutura Esperada

```text
/apps
  /api
    /cmd
    /internal
```

---

# Estrutura do CMD

Responsável pelo entrypoint da aplicação.

---

# Exemplo

```text
/apps
  /api
    /cmd
      /braqui
        main.go
```

---

# Estrutura Internal

Responsável pelo código interno da aplicação.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /application
      /domain
      /interfaces
      /infra
```

---

# Domain

Responsável por:
- entidades;
- regras de domínio;
- contratos;
- comportamento puro.

---

# Domain NÃO deve conhecer

- PostgreSQL;
- Telegram;
- HTTP;
- IA;
- frameworks;
- Docker.

---

# Application

Responsável por:
- casos de uso;
- orquestração;
- coordenação do domínio.

---

# Infra

Responsável por:
- banco;
- providers;
- integrações externas;
- Telegram;
- IA;
- clima;
- scheduler.

---

# Interfaces

Responsável por:
- HTTP;
- webhooks;
- adapters;
- handlers;
- entrada e saída externa.

---

# Organização Modular

O sistema deve ser dividido por módulos.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /pet
      /event
      /reminder
      /timeline
      /conversation
      /router
      /climate
      /insight
      /summary
```

---

# Filosofia Modular

Cada módulo deve possuir:
- alta coesão;
- baixo acoplamento;
- responsabilidade clara.

---

# Estrutura Docs

Responsável por:
- visão;
- arquitetura;
- roadmap;
- specs;
- tasks;
- playbook.

---

# Estrutura Esperada

```text
/docs
  vision.md
  architecture.md
  roadmap.md
  playbook.md

  /specs
  /tasks
```

---

# Estrutura Tests

Responsável por:
- testes de integração;
- testes E2E futuros.

---

# Estrutura Esperada

```text
/tests
  /integration
  /e2e
```

---

# Estrutura Docker

Responsável por:
- arquivos Docker;
- configurações locais;
- infraestrutura de desenvolvimento.

---

# Estrutura Esperada

```text
/docker
  /postgres
```

---

# Setup Inicial Go

O projeto deve possuir:

```text
/apps/api/go.mod
/apps/api/go.sum
```

---

# Main.go

Inicialmente:
- bootstrap simples;
- carregamento de config;
- inicialização da aplicação.

---

# Estratégia Inicial

Inicialmente:
- monólito único;
- um único processo;
- sem microsserviços.

---

# Convenções de Naming

## Use Cases

```text
CreatePet
RegisterEvent
SendReminder
```

---

## Repositories

```text
PetRepository
EventRepository
```

---

## Providers

```text
AIProvider
ClimateProvider
TelegramGateway
```

---

# Filosofia de Código

Preferir:
- código explícito;
- simplicidade;
- clareza;
- composição.

Evitar:
- abstrações prematuras;
- genericismo excessivo;
- magia.

---

# NÃO fazer

O bootstrap inicial NÃO deve:
- implementar regras de negócio;
- criar arquitetura complexa;
- criar microsserviços;
- adicionar dependências desnecessárias.

---

# Critérios de Aceite

## Estrutura Base

- estrutura de diretórios criada corretamente.

---

## Setup Go

- projeto Go inicializado corretamente.

---

## Modularização

- módulos iniciais organizados corretamente.

---

## Arquitetura

- boundaries arquiteturais respeitados.

---

## Monorepo

- estrutura preparada para múltiplas aplicações futuras.

---

## Documentação

- diretório docs organizado corretamente.

---

# Tratamento de Erros

## Estrutura inconsistente

O projeto deve:
- manter padronização;
- evitar diretórios duplicados;
- evitar código fora das boundaries definidas.

---

# Requisitos Técnicos

## Deve existir

- go.mod;
- main.go;
- estrutura modular;
- diretórios padronizados;
- organização arquitetural inicial;
- estrutura monorepo.

---

# Dependências

Esta é a spec fundacional do projeto.

Todas as outras specs dependem implicitamente dela.

---

# Considerações Arquiteturais

## Estrutura primeiro

A fundação arquitetural deve existir antes:
- da persistência;
- do Telegram;
- da IA;
- das regras de negócio.

---

## IA-assisted development

A estrutura foi desenhada para funcionar bem com:
- OpenCode;
- Cursor;
- agentes de IA;
- desenvolvimento incremental.

---

## Monorepo pragmático

O Braqui utilizará:
- monorepo simples;
- sem complexidade prematura;
- sem ferramentas enterprise inicialmente.

---

# Objetivo Real do MVP

O foco é:
- criar fundação sólida;
- evitar degradação arquitetural;
- melhorar consistência;
- facilitar evolução incremental.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- dashboard web;
- admin panel;
- mobile app;
- packages compartilhados;
- múltiplas aplicações;
- workers separados;
- arquitetura distribuída.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/apps` criada
- [ ] Estrutura `/docs` criada
- [ ] Estrutura `/docker` criada
- [ ] Estrutura `/scripts` criada
- [ ] Estrutura `/tests` criada

---

## Estrutura Monorepo

- [ ] Estrutura `/apps/api` criada
- [ ] Estrutura preparada para múltiplas aplicações
- [ ] Estratégia monorepo documentada

---

## Setup Go

- [ ] `go.mod` criado
- [ ] `go.sum` criado
- [ ] `main.go` criado
- [ ] Projeto Go inicializado corretamente

---

## Organização Arquitetural

- [ ] Estrutura `/application` criada
- [ ] Estrutura `/domain` criada
- [ ] Estrutura `/infra` criada
- [ ] Estrutura `/interfaces` criada

---

## Organização Modular

- [ ] Módulo `pet` criado
- [ ] Módulo `event` criado
- [ ] Módulo `reminder` criado
- [ ] Módulo `timeline` criado
- [ ] Módulo `conversation` criado
- [ ] Módulo `router` criado
- [ ] Módulo `climate` criado
- [ ] Módulo `insight` criado
- [ ] Módulo `summary` criado

---

## Documentação

- [ ] `vision.md` criado
- [ ] `architecture.md` criado
- [ ] `roadmap.md` criado
- [ ] `playbook.md` criado
- [ ] Diretório `/specs` criado
- [ ] Diretório `/tasks` criado

---

## Convenções

- [ ] Convenções arquiteturais documentadas
- [ ] Boundaries definidos corretamente
- [ ] Naming conventions padronizadas

---

## Qualidade

- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com OpenCode
- [ ] Estrutura compatível com IA-assisted development