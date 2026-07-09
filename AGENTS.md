# AGENTS.md

# Braqui — Guia para Agentes de IA

Este documento define como agentes de IA (OpenCode, Cursor, Windsurf, GitHub Copilot, ChatGPT e similares) devem trabalhar neste repositório.

---

# Objetivo

O Braqui é uma plataforma voltada ao acompanhamento da saúde de cães braquicefálicos através de uma experiência conversacional.

O foco do MVP é permitir que tutores registrem eventos importantes da rotina do pet e recebam informações úteis através do Telegram.

O agente NÃO deve assumir funcionalidades que não estejam descritas na documentação oficial.

---

# Fonte de Verdade

Em caso de conflito entre documentos, a seguinte ordem de prioridade deve ser respeitada:

1. SPEC correspondente
2. docs/playbook.md
3. docs/architecture.md
4. docs/vision.md
5. docs/roadmap.md
6. README.md
7. AGENTS.md

Este documento apenas complementa os demais.

---

# Processo de Desenvolvimento

Este projeto utiliza **Spec Driven Development (SDD)**.

Todo desenvolvimento deve seguir obrigatoriamente o fluxo:

```text
README
    ↓
Vision
    ↓
Architecture
    ↓
Roadmap
    ↓
Playbook
    ↓
SPEC
    ↓
Implementação
    ↓
Testes
    ↓
Atualização do Checklist
```

Nunca implemente código antes de compreender completamente a SPEC correspondente.

---

# Regras Gerais

Sempre:

- implementar apenas UMA SPEC por vez;
- respeitar rigorosamente o escopo da SPEC;
- manter o projeto simples;
- seguir a arquitetura definida;
- escrever código legível;
- criar testes quando previsto pela SPEC;
- atualizar o "Implementation Checklist" antes de finalizar.

Nunca:

- antecipar funcionalidades de outras SPECs;
- modificar SPECs sem solicitação explícita;
- remover itens do checklist;
- criar funcionalidades "porque serão úteis depois";
- criar abstrações desnecessárias;
- adicionar dependências sem necessidade.

---

# Status Atual do Projeto

Neste momento o projeto encontra-se em fase de bootstrap.

Atualmente existem apenas documentos.

Ainda NÃO existem:

- código Go;
- go.mod;
- main.go;
- Dockerfile;
- docker-compose.yml;
- migrations;
- testes;
- CI/CD.

Toda implementação deve partir dessa premissa.

---

# Arquitetura

O projeto utiliza:

- Monorepo
- Modular Monolith
- Clean Architecture (Hexagonal Light)

Estrutura principal:

```text
/apps
    /api
    /dashboard
    /admin
```

No MVP apenas:

```text
/apps/api
```

será implementado.

Dashboard e Admin pertencem ao roadmap futuro.

---

# Camadas

A direção das dependências deve permanecer:

```text
Domain
    ↓
Application
    ↓
Interfaces
    ↓
Infrastructure
```

Nunca inverter essa direção.

O domínio NÃO pode conhecer:

- PostgreSQL
- Telegram
- HTTP
- Gemini
- OpenWeather
- Frameworks

---

# Organização dos Módulos

Cada módulo deverá permanecer isolado.

Exemplo:

```text
pet
event
conversation
router
timeline
reminder
climate
insight
summary
```

Evite dependências cruzadas.

---

# Convenções de Nomenclatura

Use Cases

```text
CreatePet
RegisterEvent
GenerateInsights
SendReminder
```

Repositories

```text
PetRepository
EventRepository
ReminderRepository
```

Providers

```text
AIProvider
ClimateProvider
TelegramGateway
```

Handlers

```text
TelegramWebhookHandler
TimelineHandler
ReminderHandler
```

Utilize nomes explícitos.

---

# Filosofia

Sempre priorizar:

- simplicidade;
- clareza;
- baixo acoplamento;
- alta coesão;
- código fácil de manter.

Evitar:

- overengineering;
- design patterns desnecessários;
- abstrações prematuras;
- código "inteligente demais".

---

# Inteligência Artificial

A IA é uma funcionalidade opcional.

Ela NÃO controla o fluxo do sistema.

Fluxo esperado:

```text
Parser Local
        ↓
Sucesso
        ↓
Fim

Falha
        ↓
Gemini
        ↓
Fim
```

Nunca utilizar IA como única estratégia de processamento.

Toda IA deve ser desacoplada.

---

# Docker

Todo ambiente deve funcionar através de:

```bash
docker compose up
```

Containers esperados:

- api
- postgres

Não criar dependências que exijam instalação manual.

---

# Banco de Dados

Banco inicial:

```text
PostgreSQL
```

Toda persistência deve ocorrer através de repositories.

Nenhuma regra de negócio deve existir dentro das consultas SQL.

---

# Testes

Priorizar:

- testes unitários.

Mockar:

- repositories;
- providers;
- gateways.

Testes de integração devem utilizar PostgreSQL via Docker.

---

# Observabilidade

Sempre utilizar logging simples.

Não registrar:

- tokens;
- secrets;
- payloads completos;
- informações sensíveis.

---

# Segurança

Nunca:

- commitar arquivos .env;
- commitar credenciais;
- inserir tokens diretamente no código.

Toda configuração deve ser externa.

---

# Checklist

Toda SPEC possui uma seção:

```text
Implementation Checklist
```

Ao concluir uma implementação:

- marcar os itens concluídos;
- deixar pendentes os itens não implementados;
- justificar qualquer pendência.

Nunca remover itens do checklist.

---

# Fluxo Esperado

Para cada SPEC:

```text
Ler documentação
        ↓
Planejar implementação
        ↓
Aguardar aprovação (quando solicitado)
        ↓
Implementar
        ↓
Executar testes
        ↓
Atualizar checklist
        ↓
Finalizar
```

---

# Commits

Cada commit deve representar apenas uma SPEC.

Exemplo:

```text
feat(spec-001): bootstrap e estrutura inicial do projeto
```

Evitar commits contendo múltiplas SPECs.

---

# Restrições Arquiteturais

Não implementar:

- Microservices
- Kafka
- RabbitMQ
- CQRS
- Event Sourcing
- Kubernetes
- DDD complexo
- Arquiteturas distribuídas

O MVP deve permanecer simples.

---

# Objetivo do Agente

O agente deve atuar como um engenheiro de software experiente.

Priorizar:

- previsibilidade;
- legibilidade;
- simplicidade;
- rastreabilidade;
- facilidade de manutenção.

Quando existir dúvida entre duas soluções, escolher sempre a mais simples.

O sucesso deste projeto não é medido pela quantidade de código produzido, mas pela aderência às SPECs e pela qualidade da implementação.