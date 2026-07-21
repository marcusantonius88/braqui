# SPEC-018 - Scheduler e Tarefas Agendadas

## Objetivo

Definir a estratégia de execução de tarefas agendadas do Braqui MVP.

O foco desta spec é:
- envio de lembretes;
- geração de resumos;
- alertas automáticos;
- execução periódica de processos internos.

---

# Contexto

O Braqui não será apenas reativo.

Parte do valor do produto está em comportamentos proativos.

Exemplos:

```text
Lembrar antipulgas
```

```text
Enviar resumo semanal
```

```text
Verificar clima da região
```

Essas ações exigem execução agendada.

---

# Escopo

## O sistema deve:

- executar tarefas periódicas;
- suportar múltiplos jobs;
- permitir expansão futura;
- funcionar dentro da própria aplicação.

---

# Fora do Escopo

Esta spec NÃO contempla:

- filas distribuídas;
- workers dedicados;
- Kafka;
- RabbitMQ;
- processamento distribuído;
- cron externo.

---

# Filosofia Arquitetural

O scheduler deve ser:

- simples;
- previsível;
- desacoplado;
- de baixo custo operacional.

---

# Estratégia Inicial

Inicialmente:

```text
Scheduler interno
```

executando no mesmo processo da API.

---

# Fluxo Esperado

```text
Aplicação inicia
        ↓
Scheduler inicia
        ↓
Jobs são registrados
        ↓
Execução periódica
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /scheduler
```

---

# Responsabilidades

O scheduler deve:

- registrar jobs;
- executar jobs;
- registrar falhas;
- permitir observabilidade.

---

# Jobs Iniciais

## Lembretes

Relacionado a:

```text
SPEC-019
```

---

## Resumo Semanal

Relacionado a:

```text
SPEC-023
```

---

## Alerta Climático

Relacionado a:

```text
SPEC-021
```

---

# Interface Conceitual

```go
type Job interface {
    Name() string
    Execute(ctx context.Context) error
}
```

---

# Registro de Jobs

Exemplo:

```go
scheduler.Register(job)
```

---

# Frequência

Inicialmente:

- diária;
- semanal;
- horária.

---

# Exemplos

## Diário

```text
Verificar lembretes
```

---

## Semanal

```text
Gerar resumo semanal
```

---

## Horário

```text
Verificar clima
```

---

# Startup

O scheduler deve iniciar automaticamente.

---

# Configuração

Deve ser possível desabilitar:

```text
SCHEDULER_ENABLED=false
```

---

# Objetivo

Facilitar:
- desenvolvimento;
- testes;
- troubleshooting.

---

# Tratamento de Falhas

Falha de um job NÃO deve:

- derrubar aplicação;
- interromper outros jobs.

---

# Estratégia

Cada execução deve ser isolada.

---

# Observabilidade

Registrar:

- job iniciado;
- job concluído;
- job falhou;
- duração.

---

# Exemplo

```text
job=weekly_summary status=success
```

---

# NÃO registrar

- dados sensíveis;
- payloads completos.

---

# Critérios de Aceite

## Execução

- jobs executam corretamente.

---

## Isolamento

- falha de um job não afeta outros.

---

## Startup

- scheduler inicia corretamente.

---

## Observabilidade

- execução registrada corretamente.

---

# Requisitos Técnicos

## Deve existir

- scheduler;
- job registry;
- execução periódica;
- observabilidade.

---

# Dependências

Relaciona-se com:
- SPEC-002 - Configuração e Gerenciamento de Ambiente
- SPEC-019 - Lembretes
- SPEC-021 - Alertas Climáticos
- SPEC-023 - Resumo Semanal

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- workers;
- filas;
- infraestrutura distribuída.

---

## Processo único

Scheduler e API coexistem no mesmo runtime.

---

## Compatibilidade com Monorepo

O scheduler deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por futuras funcionalidades;
- não depender diretamente do Telegram.

---

# Objetivo Real do MVP

O foco é:
- permitir comportamento proativo;
- suportar lembretes;
- suportar resumos;
- suportar alertas automáticos.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- filas;
- workers dedicados;
- cron distribuído;
- retry avançado;
- priorização de jobs;
- processamento assíncrono.

---

# Implementation Checklist

## Estrutura Base

- [x] Estrutura `/scheduler` criada
- [x] Scheduler principal criado
- [x] Interface Job criada
- [x] Registry de jobs criado

---

## Inicialização

- [x] Scheduler inicia automaticamente
- [x] Integração com bootstrap implementada
- [x] Shutdown gracioso implementado

---

## Configuração

- [x] Variável `SCHEDULER_ENABLED` implementada
- [x] Scheduler pode ser desabilitado
- [x] Configuração integrada ao Config Loader

---

## Registro de Jobs

- [x] Registro de jobs implementado
- [x] Mecanismo de descoberta/configuração implementado
- [x] Execução periódica configurável

---

## Frequências

- [x] Execução horária suportada
- [x] Execução diária suportada
- [x] Execução semanal suportada

---

## Jobs Iniciais

- [x] Estrutura para ReminderJob criada
- [x] Estrutura para WeeklySummaryJob criada
- [x] Estrutura para ClimateAlertJob criada

---

## Execução

- [x] Execução isolada de jobs implementada
- [x] Execução concorrente avaliada/implementada
- [x] Proteção contra falha global implementada

---

## Tratamento de Falhas

- [x] Falha de job não derruba aplicação
- [x] Falha de job não interrompe outros jobs
- [x] Tratamento de panic implementado

---

## Observabilidade

- [x] Log de job iniciado implementado
- [x] Log de job concluído implementado
- [x] Log de job falhou implementado
- [x] Log de duração implementado

---

## Segurança

- [x] Dados sensíveis removidos dos logs
- [x] Payloads protegidos
- [x] Execuções auditáveis

---

## Testes

- [x] Testes do scheduler implementados
- [x] Testes de registro de jobs implementados
- [x] Testes de execução implementados
- [x] Testes de falha implementados

---

## Qualidade

- [x] Scheduler desacoplado das funcionalidades
- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Base preparada para novos jobs futuros