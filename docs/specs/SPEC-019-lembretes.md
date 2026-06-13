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

- [ ] Entidade Reminder criada
- [ ] ReminderRepository criado
- [ ] Casos de uso de lembretes criados

---

## Persistência

- [ ] Tabela reminders criada
- [ ] Migration de reminders criada
- [ ] Persistência implementada

---

## Modelo de Dados

- [ ] Campo pet_id implementado
- [ ] Campo title implementado
- [ ] Campo description implementado
- [ ] Campo due_at implementado
- [ ] Campo status implementado

---

## Status

- [ ] Status pending implementado
- [ ] Status completed implementado
- [ ] Status cancelled implementado

---

## Criação de Lembretes

- [ ] Fluxo conversacional implementado
- [ ] Criação manual implementada
- [ ] Validação de dados implementada

---

## Interpretação de Datas

- [ ] Suporte a "amanhã"
- [ ] Suporte a datas explícitas
- [ ] Validação de datas implementada

---

## Integração com Scheduler

- [ ] Consulta de lembretes pendentes implementada
- [ ] Execução automática implementada
- [ ] Integração com jobs implementada

---

## Integração com Telegram

- [ ] Envio de lembrete implementado
- [ ] Template de mensagem implementado
- [ ] Tratamento de falhas implementado

---

## Atualização de Status

- [ ] Atualização após envio implementada
- [ ] Estratégia de reprocessamento definida

---

## Integração Conversacional

- [ ] Integração com Router implementada
- [ ] Gatilhos de lembrete implementados

---

## Observabilidade

- [ ] Log de criação implementado
- [ ] Log de envio implementado
- [ ] Log de falha implementado

---

## Tratamento de Erros

- [ ] Data inválida tratada
- [ ] Falha de persistência tratada
- [ ] Falha de envio tratada

---

## Testes

- [ ] Testes de criação implementados
- [ ] Testes de persistência implementados
- [ ] Testes de scheduler implementados
- [ ] Testes de envio implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Lembretes desacoplados do Telegram