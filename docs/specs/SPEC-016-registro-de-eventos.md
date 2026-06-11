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

- [ ] Entidade Event criada
- [ ] EventRepository criado
- [ ] Casos de uso de registro criados

---

## Persistência

- [ ] Tabela events criada
- [ ] Migration de events criada
- [ ] Persistência implementada

---

## Modelo de Evento

- [ ] Campo pet_id implementado
- [ ] Campo type implementado
- [ ] Campo description implementado
- [ ] Campo source implementado
- [ ] Campo created_at implementado

---

## Tipos de Evento

- [ ] Tipo vomit suportado
- [ ] Tipo itching suportado
- [ ] Tipo panting suportado
- [ ] Tipo medication_given suportado
- [ ] Tipo vet_visit suportado

---

## Associação

- [ ] Associação Pet → Event implementada
- [ ] Validação de pet existente implementada

---

## Registro

- [ ] Registro via parser implementado
- [ ] Registro via IA implementado
- [ ] Registro manual preparado

---

## Histórico

- [ ] Ordenação temporal implementada
- [ ] Consulta de histórico implementada
- [ ] Recuperação de eventos implementada

---

## Imutabilidade

- [ ] Eventos protegidos contra edição
- [ ] Estratégia de correção documentada

---

## Observabilidade

- [ ] Log de criação de evento implementado
- [ ] Log de falha de persistência implementado
- [ ] Log de origem do evento implementado

---

## Tratamento de Erros

- [ ] Evento inválido tratado
- [ ] Falha de persistência tratada
- [ ] Mensagens amigáveis implementadas

---

## Testes

- [ ] Testes de criação implementados
- [ ] Testes de persistência implementados
- [ ] Testes de associação implementados
- [ ] Testes de histórico implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Registro desacoplado da origem do evento