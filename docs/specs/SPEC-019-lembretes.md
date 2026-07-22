# SPEC-019 - Lembretes

## Objetivo

Definir a funcionalidade de lembretes do Braqui MVP.

O foco desta spec é:
- ajudar o tutor a lembrar compromissos importantes;
- reduzir esquecimentos;
- aumentar valor recorrente do produto;
- estimular engajamento contínuo.

---

# Contexto

Muitos cuidados com cães braquicefálicos dependem de recorrência.

Exemplos:

```text
Antipulgas
```

```text
Medicamentos
```

```text
Consulta veterinária
```

```text
Vacinação
```

O Braqui deve ajudar o tutor a lembrar dessas atividades.

---

# Escopo

## O sistema deve:

- criar lembretes;
- armazenar lembretes;
- executar lembretes agendados;
- enviar notificações via Telegram.

---

# Fora do Escopo

Esta spec NÃO contempla:

- integração com calendário;
- sincronização externa;
- recorrência avançada;
- notificações push mobile;
- e-mail.

---

# Filosofia Arquitetural

Lembretes devem ser:

- simples;
- confiáveis;
- fáceis de criar;
- fáceis de entender.

---

# Estrutura Esperada

```text
Reminder
```

---

# Campos Iniciais

```text
id
pet_id
title
description
due_at
status
created_at
updated_at
```

---

# Significado dos Campos

## title

Exemplo:

```text
Dar Simparic
```

---

## due_at

Data prevista.

---

## status

Valores:

```text
pending
completed
cancelled
```

---

# Fluxo Esperado

```text
Usuário cria lembrete
        ↓
Persistência
        ↓
Scheduler verifica
        ↓
Telegram envia mensagem
```

---

# Exemplos de Criação

```text
Me lembre de dar Simparic daqui a 30 dias
```

---

```text
Lembrar consulta veterinária dia 15
```

---

# MVP

Inicialmente:

- criação simples;
- data única.

---

# Recorrência

Fora do MVP.

---

# Persistência

Os lembretes devem ser armazenados.

---

# Estrutura Esperada

```text
ReminderRepository
```

---

# Integração com Scheduler

Relacionamento:

```text
Scheduler
      ↓
ReminderRepository
      ↓
Telegram
```

---

# Mensagem de Lembrete

Exemplo:

```text
🐶 Lembrete do Thor

Hoje é dia de dar Simparic.
```

---

# Conclusão

Após envio:

```text
completed
```

ou

```text
pending
```

dependendo da estratégia escolhida.

---

# Estratégia Inicial

Inicialmente:

```text
pending → enviado
```

---

# Criação Conversacional

O usuário poderá criar lembretes através do chat.

---

# Exemplo

```text
Lembre consulta veterinária para amanhã
```

---

# Integração com Router

Mensagens relacionadas a lembretes devem ser encaminhadas para o fluxo adequado.

---

# Observabilidade

Registrar:

- lembrete criado;
- lembrete enviado;
- falha de envio.

---

# NÃO registrar

- dados sensíveis;
- payloads completos.

---

# Tratamento de Erros

## Data inválida

Responder:

```text
Não consegui entender a data informada.
```

---

## Falha de envio

Registrar erro e permitir nova tentativa futura.

---

# Critérios de Aceite

## Criação

- lembrete criado corretamente.

---

## Persistência

- lembrete armazenado corretamente.

---

## Scheduler

- lembrete executado corretamente.

---

## Telegram

- mensagem enviada corretamente.

---

# Requisitos Técnicos

## Deve existir

- entidade Reminder;
- ReminderRepository;
- integração com Scheduler;
- integração com Telegram.

---

# Dependências

Relaciona-se com:
- SPEC-018 - Scheduler e Tarefas Agendadas
- SPEC-009 - Integração com Telegram
- SPEC-005 - Persistência e Repositories

---

# Considerações Arquiteturais

## Valor recorrente

Lembretes aumentam:
- retenção;
- recorrência;
- utilidade percebida.

---

## Simplicidade primeiro

O MVP NÃO precisa:
- recorrência complexa;
- calendário;
- múltiplos canais.

---

## Compatibilidade com Monorepo

O módulo de lembretes deve:
- permanecer em `/apps/api`;
- ser reutilizável por futuros canais;
- não depender diretamente do Telegram.

---

# Objetivo Real do MVP

O foco é:
- ajudar o tutor;
- gerar valor contínuo;
- aumentar engajamento.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- recorrência;
- calendário;
- múltiplos canais;
- lembretes inteligentes;
- lembretes baseados em eventos.

---

# Implementation Checklist

## Estrutura Base

- [x] Entidade Reminder criada
- [x] ReminderRepository criado
- [x] Casos de uso de lembretes criados

---

## Persistência

- [x] Tabela reminders criada
- [x] Migration de reminders criada
- [x] Persistência implementada

---

## Modelo de Dados

- [x] Campo pet_id implementado
- [x] Campo title implementado
- [x] Campo description implementado
- [x] Campo due_at implementado
- [x] Campo status implementado

---

## Status

- [x] Status pending implementado
- [x] Status completed implementado
- [x] Status cancelled implementado

---

## Criação de Lembretes

- [x] Fluxo conversacional implementado
- [x] Criação manual implementada
- [x] Validação de dados implementada

---

## Interpretação de Datas

- [x] Suporte a "amanhã"
- [x] Suporte a datas explícitas (dia X, daqui a X dias, em X dias)
- [x] Validação de datas implementada

---

## Integração com Scheduler

- [x] Consulta de lembretes pendentes implementada
- [x] Execução automática implementada
- [x] Integração com jobs implementada

---

## Integração com Telegram

- [x] Envio de lembrete implementado
- [x] Template de mensagem implementado
- [x] Tratamento de falhas implementado

---

## Atualização de Status

- [x] Atualização após envio implementada
- [x] Estratégia de reprocessamento definida (mantém completed)

---

## Integração Conversacional

- [x] Integração com Router implementada
- [x] Gatilhos de lembrete implementados (/remind)

---

## Observabilidade

- [x] Log de criação implementado
- [x] Log de envio implementado
- [x] Log de falha implementado

---

## Tratamento de Erros

- [x] Data inválida tratada
- [x] Falha de persistência tratada
- [x] Falha de envio tratada

---

## Testes

- [x] Testes de criação implementados
- [x] Testes de persistência implementados
- [x] Testes de scheduler implementados
- [x] Testes de envio implementados

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Lembretes desacoplados do Telegram