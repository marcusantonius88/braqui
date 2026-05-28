# PLAYBOOK - Braqui

## Objetivo

Este documento define:
- padrões;
- convenções;
- princípios;
- regras arquiteturais;
- diretrizes de desenvolvimento;

do projeto Braqui.

O objetivo é garantir:
- consistência;
- previsibilidade;
- simplicidade;
- facilidade de evolução;
- compatibilidade com desenvolvimento assistido por IA.

---

# Filosofia do Projeto

O Braqui deve priorizar:

- simplicidade;
- clareza;
- modularidade;
- baixo acoplamento;
- domínio explícito;
- evolução incremental;
- baixo custo operacional;
- experiência conversacional fluida.

---

# Filosofia de Produto

O Braqui NÃO é:
- chatbot genérico;
- assistente veterinário;
- IA livre conversacional.

O Braqui É:
- um copiloto de saúde contínua;
- contextual;
- proativo;
- especializado em cães braquicefálicos.

---

# Filosofia de IA

A IA deve ser:
- opcional;
- desacoplada;
- substituível;
- utilizada apenas quando necessário.

A IA NÃO deve:
- controlar o fluxo do sistema;
- substituir regras básicas;
- virar dependência central.

---

# Estratégia de Desenvolvimento

O projeto segue:
> Spec Driven Development (SDD)

Fluxo esperado:

1. Vision
2. Architecture
3. Specs
4. Tasks
5. Implementação

---

# Organização do Projeto

Estrutura inicial esperada:

```text
/internal
  /domain
  /application
  /infra
  /interfaces

/docs
  vision.md
  architecture.md
  roadmap.md

  /specs
```

---

# Organização Modular

O sistema deve ser organizado em módulos coesos.

Exemplo:

```text
/internal
  /pet
  /event
  /reminder
  /climate
  /insight
  /ai
```

---

# Regras Arquiteturais

## Domain

A camada de domínio:
- NÃO conhece infraestrutura;
- NÃO conhece banco;
- NÃO conhece Telegram;
- NÃO conhece IA;
- NÃO conhece frameworks.

---

# Domain NÃO deve importar

- postgres
- telegram
- http
- openai/gemini sdk
- ORM
- framework web

---

# Application

A camada de aplicação:
- orquestra casos de uso;
- coordena domínio;
- utiliza interfaces;
- NÃO contém detalhes de infraestrutura.

---

# Infra

A camada infra:
- implementa providers;
- integra APIs externas;
- implementa repositories;
- integra banco;
- integra Telegram.

---

# Interfaces

Responsável por:
- HTTP;
- webhook Telegram;
- scheduler triggers;
- adapters externos.

---

# Regra de Ouro

## Handlers NÃO possuem regra de negócio

Handlers devem apenas:
- validar entrada;
- delegar execução;
- retornar resposta.

---

# Regra de Ouro

## Repositories NÃO possuem regra de negócio

Repositories devem apenas:
- persistir;
- consultar;
- recuperar dados.

---

# Regra de Ouro

## Use Cases concentram comportamento da aplicação

Toda orquestração deve ocorrer:
- na camada application;
- em casos de uso explícitos.

---

# Filosofia de Código

Preferir:
- código explícito;
- simplicidade;
- clareza;
- composição;
- pequenas abstrações.

Evitar:
- abstração prematura;
- genericismo excessivo;
- clever code;
- magia.

---

# Convenções de Naming

## Use Cases

Formato:

```text
CreatePet
RegisterEvent
GenerateInsights
SendReminder
```

---

# Repositories

Formato:

```text
PetRepository
EventRepository
ReminderRepository
```

---

# Providers

Formato:

```text
AIProvider
ClimateProvider
TelegramGateway
```

---

# Handlers

Formato:

```text
TelegramWebhookHandler
ReminderHandler
TimelineHandler
```

---

# Estratégia de IA

Fluxo esperado:

1. parser local primeiro
2. IA apenas se necessário
3. fallback amigável

---

# NÃO fazer

A IA NÃO deve:
- responder diretamente usuários livremente;
- gerar textos enormes;
- inventar diagnósticos;
- controlar fluxo principal.

---

# Estratégia Conversacional

As mensagens do Braqui devem ser:
- curtas;
- acolhedoras;
- simples;
- contextuais;
- não alarmistas.

---

# NÃO fazer

Evitar:
- textões;
- linguagem médica;
- tom robótico;
- excesso de explicação.

---

# Exemplo Correto

```text
Thor apresentou mais episódios de cansaço esta semana 🐶
```

---

# Exemplo Incorreto

```text
Com base nas informações fornecidas...
```

---

# Persistência

Acesso ao banco deve ocorrer:
- exclusivamente via repositories.

---

# NÃO fazer

Proibido:
- SQL em handlers;
- SQL em domínio;
- acesso direto ao banco fora infra.

---

# Estratégia de Testes

Priorizar:
- testes unitários;
- parsing;
- casos de uso;
- fluxos críticos.

---

# NÃO fazer

Evitar:
- dependência de APIs reais;
- dependência de IA real;
- flaky tests.

---

# Estratégia de Observabilidade

Logs devem ser:
- estruturados;
- claros;
- contextuais.

---

# NÃO logar

- tokens;
- secrets;
- dados sensíveis;
- payloads excessivos.

---

# Estratégia de Docker

O ambiente deve funcionar com:

```bash
docker compose up
```

---

# Filosofia Operacional

O MVP deve priorizar:
- simplicidade;
- poucos serviços;
- baixo custo operacional;
- deploy rápido.

---

# NÃO fazer

O MVP NÃO precisa inicialmente:
- Kubernetes;
- microsserviços;
- Kafka;
- filas distribuídas;
- event sourcing;
- CQRS complexo.

---

# Filosofia de Evolução

O Braqui deve evoluir:
- incrementalmente;
- baseado em feedback real;
- baseado em retenção;
- baseado em comportamento dos usuários.

---

# Regra de Ouro

## NÃO implementar complexidade antes da necessidade real.

---

# Desenvolvimento Assistido por IA

O projeto foi estruturado para funcionar bem com:
- OpenCode;
- Cursor;
- agentes de IA;
- desenvolvimento orientado a specs.

---

# Diretrizes para IA

Sempre:
- seguir specs;
- respeitar arquitetura;
- evitar abstrações desnecessárias;
- manter baixo acoplamento;
- manter simplicidade.

---

# Checklist de Implementação das Specs

Toda spec deve possuir ao final uma seção:

```md
# Implementation Checklist
```

---

# Objetivo da Checklist

A checklist deve:
- refletir o estado real da implementação;
- ser objetiva;
- possuir granularidade pequena;
- ser facilmente interpretável por humanos e IA;
- funcionar como rastreamento incremental da spec.

---

# Formato Obrigatório

## Implementado

```md
- [x] Estrutura de pastas criada
```

---

## Pendente

```md
- [ ] Testes unitários implementados
```

---

# Regras da Checklist

A checklist:
- deve ser atualizada após implementação;
- deve refletir o estado real do código;
- deve evitar itens genéricos demais;
- deve possuir itens pequenos e verificáveis.

---

# Objetivo da Checklist no Workflow com IA

A checklist existe para:
- reduzir retrabalho;
- melhorar continuidade de contexto;
- facilitar retomada de desenvolvimento;
- permitir implementação incremental;
- facilitar revisão de progresso;
- melhorar interação com OpenCode.

---

# Exemplo de Uso com IA

Exemplo de prompt:

```text
Leia a SPEC-001 e implemente apenas os itens pendentes do Implementation Checklist.
```

---

# Antes de implementar qualquer task

Sempre verificar:
- vision.md
- architecture.md
- spec correspondente
- regras deste playbook

---

# Objetivo Final

O Braqui deve se manter:
- simples;
- elegante;
- modular;
- evolutivo;
- fácil de entender;
- fácil de manter.

A arquitetura deve servir ao produto.
Nunca o contrário.