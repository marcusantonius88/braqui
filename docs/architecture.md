# Architecture - Braqui

## Visão Arquitetural

O Braqui será construído como:
- monorepo;
- monólito modular inicialmente;
- orientado a domínio;
- preparado para evolução incremental;
- otimizado para desenvolvimento assistido por IA.

A arquitetura prioriza:
- simplicidade;
- clareza;
- baixo acoplamento;
- alta coesão;
- facilidade de manutenção;
- baixo custo operacional.

---

# Filosofia Arquitetural

O objetivo NÃO é:
- criar arquitetura enterprise prematuramente;
- introduzir complexidade desnecessária;
- otimizar escala inexistente.

O objetivo É:
- construir uma base sólida;
- evoluir incrementalmente;
- permitir adaptação rápida;
- manter o código compreensível.

---

# Estratégia de Repositório

O Braqui utilizará:
- monorepo;
- organização orientada a aplicações;
- compartilhamento centralizado de documentação e specs.

---

# Objetivos do Monorepo

A estratégia de monorepo existe para:
- facilitar IA-assisted development;
- centralizar contexto;
- simplificar evolução;
- reduzir fragmentação;
- manter consistência arquitetural.

---

# Estrutura Inicial do Monorepo

```text
/apps
/docs
/docker
/scripts
/tests
```

---

# Estrutura Apps

Responsável pelas aplicações do projeto.

---

# Estrutura Inicial Esperada

```text
/apps
  /api
```

---

# Evoluções Futuras Possíveis

```text
/apps
  /api
  /dashboard
  /admin
```

---

# Estratégia Inicial

Inicialmente:
- apenas `/apps/api` existirá;
- um único backend;
- um único processo;
- um único deploy.

---

# Estrutura da API

```text
/apps
  /api
    /cmd
    /internal
```

---

# CMD

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

# Internal

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

# Estratégia Arquitetural

O Braqui utilizará:
- Clean Architecture leve;
- arquitetura hexagonal leve;
- modularização explícita;
- boundaries claros.

---

# Domain Layer

Responsável por:
- entidades;
- regras de negócio;
- contratos;
- comportamento puro.

---

# Domain NÃO conhece

- PostgreSQL;
- Telegram;
- HTTP;
- Docker;
- IA;
- frameworks;
- detalhes externos.

---

# Application Layer

Responsável por:
- casos de uso;
- orquestração;
- coordenação entre módulos.

---

# Infra Layer

Responsável por:
- banco de dados;
- Telegram;
- IA;
- clima;
- providers externos;
- scheduler;
- Docker integration.

---

# Interfaces Layer

Responsável por:
- HTTP;
- webhook Telegram;
- handlers;
- adapters externos;
- entrada e saída do sistema.

---

# Organização Modular

O sistema será organizado por módulos de domínio.

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

Cada módulo deve:
- possuir responsabilidade clara;
- ser altamente coeso;
- possuir baixo acoplamento;
- evitar dependências cíclicas.

---

# Estratégia de IA

O Braqui utilizará abordagem híbrida:
- parser local primeiro;
- IA apenas quando necessário.

---

# Fluxo Esperado

1. parser local tenta interpretar
2. caso falhe:
   - IA é acionada
3. caso IA falhe:
   - fallback amigável

---

# Filosofia de IA

A IA:
- é capability auxiliar;
- NÃO controla o sistema;
- NÃO substitui regras explícitas;
- deve ser desacoplada.

---

# Estratégia Conversacional

O Braqui será:
- conversacional;
- contextual;
- proativo;
- orientado a acompanhamento contínuo.

---

# Canal Inicial

Inicialmente:
- Telegram Bot API.

---

# Estratégia de Persistência

Inicialmente:
- PostgreSQL;
- repositories explícitos;
- SQL simples;
- sem ORM complexo.

---

# Estratégia Operacional

Inicialmente:
- monólito único;
- um único deploy;
- um único processo;
- scheduler interno.

---

# Containerização

O projeto utilizará:
- Docker;
- Docker Compose;
- ambiente local reproduzível.

---

# Objetivo da Containerização

- facilitar onboarding;
- facilitar OpenCode;
- reduzir inconsistências;
- simplificar deploy.

---

# Estrutura Docker

```text
/docker
```

---

# Estratégia de Desenvolvimento

O Braqui seguirá:
- Spec Driven Development (SDD);
- implementação incremental;
- specs pequenas;
- tasks pequenas;
- contexto reduzido para IA.

---

# Fluxo de Desenvolvimento

1. vision.md
2. architecture.md
3. specs
4. tasks
5. implementação incremental

---

# Estratégia de Testes

Priorizar:
- testes unitários;
- parser;
- casos de uso;
- fluxo conversacional;
- repositories.

---

# Estratégia de Observabilidade

Inicialmente:
- logs estruturados;
- healthcheck;
- troubleshooting simples.

---

# Estratégia de Deploy

Inicialmente:
- deploy simples;
- Render/Railway/Fly.io;
- PostgreSQL gerenciado.

---

# Filosofia de Escalabilidade

Escalabilidade será tratada:
- apenas após validação real;
- baseada em uso;
- baseada em retenção.

---

# NÃO fazer inicialmente

O MVP NÃO precisa:
- microsserviços;
- Kubernetes;
- Kafka;
- CQRS complexo;
- event sourcing;
- infraestrutura distribuída.

---

# Boundaries Arquiteturais

## Domain

NÃO depende de:
- infra;
- interfaces;
- providers.

---

## Application

Pode depender de:
- domain;
- contratos/interfaces.

---

## Infra

Pode implementar:
- repositories;
- providers;
- gateways.

---

## Interfaces

Pode depender de:
- application;
- contratos.

---

# Filosofia de Código

Preferir:
- código explícito;
- simplicidade;
- clareza;
- composição;
- pragmatismo.

Evitar:
- abstrações prematuras;
- genericismo excessivo;
- clever code;
- magia.

---

# Objetivo Real da Arquitetura

A arquitetura existe para:
- acelerar evolução;
- reduzir degradação;
- facilitar IA-assisted development;
- manter consistência;
- facilitar manutenção.

---

# Evoluções Futuras Possíveis

Fora do MVP:
- dashboard web;
- admin panel;
- mobile app;
- múltiplos workers;
- Redis;
- filas;
- observabilidade avançada;
- arquitetura distribuída;
- packages compartilhados.