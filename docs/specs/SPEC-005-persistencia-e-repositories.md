# SPEC-005 - Persistência e Repositories

## Objetivo

Definir a estratégia de persistência de dados do Braqui MVP e padronizar o acesso ao banco através de repositories.

O foco desta spec é:
- isolamento da persistência;
- baixo acoplamento;
- simplicidade;
- testabilidade;
- aderência à arquitetura definida.

---

# Contexto

O Braqui armazenará informações relacionadas a:

- usuários;
- pets;
- eventos;
- lembretes;
- estado conversacional;
- resumos;
- insights.

O sistema precisa manter:
- histórico persistente;
- consultas simples;
- evolução incremental do modelo de dados.

---

# Escopo

## O sistema deve:

- utilizar PostgreSQL;
- utilizar repositories explícitos;
- isolar acesso ao banco;
- suportar migrations;
- permitir testes desacoplados.

---

# Fora do Escopo

Esta spec NÃO contempla:

- event sourcing;
- CQRS;
- banco distribuído;
- cache distribuído;
- replicação;
- sharding;
- ORM complexo.

---

# Filosofia Arquitetural

Persistência deve:
- ser simples;
- previsível;
- explícita;
- desacoplada do domínio.

---

# Estratégia de Banco

Inicialmente:

- PostgreSQL único;
- instância única;
- sem replicação.

---

# Estratégia de Acesso

Todo acesso ao banco deve ocorrer através de repositories.

---

# NÃO fazer

Proibido:

- SQL em handlers;
- SQL em use cases;
- SQL em domínio;
- acesso direto ao banco fora da camada infra.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /infra
        /postgres
          /repositories
          /migrations
```

---

# Organização dos Repositories

Exemplo:

```text
PetRepository
UserRepository
EventRepository
ReminderRepository
ConversationStateRepository
```

---

# Interface dos Repositories

As interfaces devem ficar próximas ao domínio ou aplicação.

---

# Exemplo Conceitual

```go
type PetRepository interface {
    Create(ctx context.Context, pet Pet) error
    FindByID(ctx context.Context, id string) (*Pet, error)
}
```

---

# Implementações

Devem ficar em:

```text
/apps/api/internal/infra/postgres/repositories
```

---

# Estratégia de Migrations

O banco deve ser versionado.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /infra
        /postgres
          /migrations
```

---

# Objetivos das Migrations

Permitir:
- criação do banco;
- evolução incremental;
- ambiente reproduzível.

---

# Entidades Iniciais

## User

Representa o tutor.

---

## Pet

Representa o cão.

---

## Event

Representa registros de saúde.

---

## Reminder

Representa lembretes agendados.

---

## ConversationState

Representa estado da conversa.

---

# Estratégia de Modelagem

Inicialmente:
- UUID como identificador;
- timestamps de criação;
- timestamps de atualização.

---

# Exemplo Conceitual

```go
type BaseEntity struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

---

# Transações

Inicialmente:
- apenas quando necessário.

Evitar:
- abstrações complexas;
- unit of work sofisticado.

---

# Tratamento de Erros

Erros de persistência devem:
- ser propagados;
- ser tratados na camada superior.

---

# Exemplo

```text
pet not found
```

---

# Observabilidade

Operações críticas devem:
- registrar erros;
- facilitar troubleshooting.

---

# NÃO fazer

Não registrar:
- dados sensíveis;
- tokens;
- segredos.

---

# Critérios de Aceite

## Banco

- PostgreSQL funcionando corretamente.

---

## Repositories

- acesso ao banco ocorre exclusivamente via repositories.

---

## Migrations

- banco pode ser recriado a partir das migrations.

---

## Arquitetura

- domínio permanece desacoplado.

---

## Testabilidade

- repositories podem ser mockados.

---

# Requisitos Técnicos

## Deve existir

- conexão PostgreSQL;
- repositories;
- migrations;
- interfaces;
- testes de persistência.

---

# Dependências

Relaciona-se com:
- SPEC-001 - Bootstrap e Estrutura Inicial do Projeto
- SPEC-002 - Configuração e Gerenciamento de Ambiente
- SPEC-004 - Containerização e Ambiente Local

---

# Considerações Arquiteturais

## Repository Pattern

O padrão repository existe para:
- desacoplar domínio;
- facilitar testes;
- permitir evolução futura.

---

## Simplicidade primeiro

O MVP NÃO precisa:
- ORM complexo;
- abstrações sofisticadas;
- frameworks pesados de persistência.

---

## Compatibilidade com Monorepo

A persistência deve:
- permanecer isolada em `/apps/api`;
- evitar dependências externas desnecessárias;
- suportar futuras aplicações do monorepo.

---

# Objetivo Real do MVP

O foco é:
- armazenar dados de forma confiável;
- manter arquitetura limpa;
- facilitar evolução incremental;
- melhorar testabilidade.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- Redis;
- cache;
- read replicas;
- observabilidade avançada;
- múltiplos bancos;
- event sourcing;
- CQRS.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/infra/postgres` criada
- [ ] Estrutura `/repositories` criada
- [ ] Estrutura `/migrations` criada

---

## Banco de Dados

- [ ] Conexão PostgreSQL implementada
- [ ] Pool de conexões configurado
- [ ] Healthcheck de banco implementado

---

## Migrations

- [ ] Ferramenta de migrations definida
- [ ] Migration inicial criada
- [ ] Execução de migrations funcionando
- [ ] Recriação completa do banco validada

---

## Entidades

- [ ] Entidade User criada
- [ ] Entidade Pet criada
- [ ] Entidade Event criada
- [ ] Entidade Reminder criada
- [ ] Entidade ConversationState criada

---

## Repositories

- [ ] Interface UserRepository criada
- [ ] Interface PetRepository criada
- [ ] Interface EventRepository criada
- [ ] Interface ReminderRepository criada
- [ ] Interface ConversationStateRepository criada

---

## Implementações PostgreSQL

- [ ] PostgreSQLUserRepository implementado
- [ ] PostgreSQLPetRepository implementado
- [ ] PostgreSQLEventRepository implementado
- [ ] PostgreSQLReminderRepository implementado
- [ ] PostgreSQLConversationStateRepository implementado

---

## Arquitetura

- [ ] Domínio desacoplado da persistência
- [ ] Use cases utilizando interfaces
- [ ] Infra implementando repositories
- [ ] SQL isolado na camada infra

---

## Tratamento de Erros

- [ ] Erros de persistência padronizados
- [ ] Erro "not found" implementado
- [ ] Logs de erro implementados

---

## Observabilidade

- [ ] Logs de conexão implementados
- [ ] Logs de erro implementados
- [ ] Troubleshooting básico implementado

---

## Testes

- [ ] Mocks dos repositories criados
- [ ] Testes unitários dos repositories implementados
- [ ] Testes de integração PostgreSQL implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development