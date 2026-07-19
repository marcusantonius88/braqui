# SPEC-016 - Registro de Eventos

## Objetivo

Definir a estratégia de registro de eventos relacionados à saúde, comportamento e rotina do pet.

O foco desta spec é:
- armazenar histórico;
- criar base para insights;
- criar base para lembretes;
- permitir acompanhamento contínuo.

---

# Contexto

O principal valor do Braqui está na construção de histórico ao longo do tempo.

Exemplos de eventos:

```text
Thor vomitou
```

```text
Thor tomou simparic
```

```text
Thor está muito ofegante
```

```text
Thor foi ao veterinário
```

Esses eventos devem ser registrados de forma estruturada.

---

# Escopo

## O sistema deve:

- registrar eventos;
- associar eventos ao pet;
- armazenar histórico;
- permitir consultas futuras.

---

# Fora do Escopo

Esta spec NÃO contempla:

- diagnósticos veterinários;
- recomendações médicas;
- análise avançada;
- machine learning.

---

# Filosofia Arquitetural

Eventos devem ser:

- simples;
- estruturados;
- imutáveis;
- auditáveis.

---

# Estrutura Esperada

```text
Event
```

---

# Campos Iniciais

```text
id
pet_id
type
description
source
created_at
```

---

# Significado dos Campos

## pet_id

Pet associado.

---

## type

Tipo do evento.

Exemplos:

```text
vomit
itching
panting
medication_given
vet_visit
```

---

## description

Mensagem original.

Exemplo:

```text
Thor vomitou hoje pela manhã
```

---

## source

Origem do registro.

Exemplo:

```text
parser
ai
manual
```

---

# Estratégia de Registro

Fluxo:

```text
Mensagem
      ↓
Parser
      ↓
IA (opcional)
      ↓
Evento
      ↓
Persistência
```

---

# Associação

Relacionamento:

```text
User
      ↓
Pet
      ↓
Event
```

---

# Persistência

Eventos devem ser armazenados permanentemente.

---

# Estrutura Esperada

```text
EventRepository
```

---

# Exemplos

## Evento de Saúde

```json
{
  "type": "vomit"
}
```

---

## Evento de Medicação

```json
{
  "type": "medication_given"
}
```

---

## Evento de Consulta

```json
{
  "type": "vet_visit"
}
```

---

# Imutabilidade

Eventos não devem ser alterados após criação.

Correções futuras devem gerar novos eventos.

---

# Estratégia de Histórico

Todos os eventos devem possuir:

```text
created_at
```

para permitir ordenação temporal.

---

# Observabilidade

Registrar:

- evento criado;
- falha de persistência;
- origem do evento.

---

# NÃO registrar

- informações sensíveis;
- conteúdo excessivo.

---

# Tratamento de Erros

## Evento inválido

Sistema deve:

- rejeitar criação;
- registrar erro.

---

## Falha de persistência

Sistema deve:

- registrar erro;
- informar usuário.

---

# Critérios de Aceite

## Registro

- eventos registrados corretamente.

---

## Associação

- evento associado ao pet correto.

---

## Persistência

- histórico armazenado corretamente.

---

## Arquitetura

- registro desacoplado da origem.

---

# Requisitos Técnicos

## Deve existir

- entidade Event;
- EventRepository;
- persistência;
- histórico.

---

# Dependências

Relaciona-se com:
- SPEC-014 - Parser Local de Eventos
- SPEC-015 - Provider de IA
- SPEC-005 - Persistência e Repositories

---

# Considerações Arquiteturais

## Histórico como ativo principal

O histórico é a base para:

- insights;
- lembretes;
- análises futuras.

---

## Imutabilidade

Eventos representam fatos históricos.

---

## Compatibilidade com Monorepo

O registro deve:
- permanecer isolado em `/apps/api`;
- ser reutilizável por futuros canais;
- não depender diretamente do Telegram.

---

# Objetivo Real do MVP

O foco é:
- construir histórico;
- criar memória operacional do pet;
- preparar funcionalidades futuras.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- anexos;
- fotos;
- localização de eventos;
- severidade;
- tags;
- categorização avançada.

---

# Implementation Checklist

## Estrutura Base

- [x] Entidade Event criada
- [x] EventRepository criado
- [x] Casos de uso de registro criados

---

## Persistência

- [x] Tabela events criada
- [x] Migration de events criada
- [x] Persistência implementada

---

## Modelo de Evento

- [x] Campo pet_id implementado
- [x] Campo type implementado
- [x] Campo description implementado
- [x] Campo source implementado
- [x] Campo created_at implementado

---

## Tipos de Evento

- [x] Tipo vomit suportado
- [x] Tipo itching suportado
- [x] Tipo panting suportado
- [x] Tipo medication_given suportado
- [x] Tipo vet_visit suportado

---

## Associação

- [x] Associação Pet → Event implementada
- [x] Validação de pet existente implementada

---

## Registro

- [x] Registro via parser implementado
- [x] Registro via IA implementado
- [x] Registro manual preparado

---

## Histórico

- [x] Ordenação temporal implementada
- [x] Consulta de histórico implementada
- [x] Recuperação de eventos implementada

---

## Imutabilidade

- [x] Eventos protegidos contra edição
- [x] Estratégia de correção documentada

---

## Observabilidade

- [x] Log de criação de evento implementado
- [x] Log de falha de persistência implementado
- [x] Log de origem do evento implementado

---

## Tratamento de Erros

- [x] Evento inválido tratado
- [x] Falha de persistência tratada
- [x] Mensagens amigáveis implementadas

---

## Testes

- [x] Testes de criação implementados
- [x] Testes de persistência implementados
- [x] Testes de associação implementados
- [x] Testes de histórico implementados

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Registro desacoplado da origem do evento