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

- [ ] Estrutura `/scheduler` criada
- [ ] Scheduler principal criado
- [ ] Interface Job criada
- [ ] Registry de jobs criado

---

## Inicialização

- [ ] Scheduler inicia automaticamente
- [ ] Integração com bootstrap implementada
- [ ] Shutdown gracioso implementado

---

## Configuração

- [ ] Variável `SCHEDULER_ENABLED` implementada
- [ ] Scheduler pode ser desabilitado
- [ ] Configuração integrada ao Config Loader

---

## Registro de Jobs

- [ ] Registro de jobs implementado
- [ ] Mecanismo de descoberta/configuração implementado
- [ ] Execução periódica configurável

---

## Frequências

- [ ] Execução horária suportada
- [ ] Execução diária suportada
- [ ] Execução semanal suportada

---

## Jobs Iniciais

- [ ] Estrutura para ReminderJob criada
- [ ] Estrutura para WeeklySummaryJob criada
- [ ] Estrutura para ClimateAlertJob criada

---

## Execução

- [ ] Execução isolada de jobs implementada
- [ ] Execução concorrente avaliada/implementada
- [ ] Proteção contra falha global implementada

---

## Tratamento de Falhas

- [ ] Falha de job não derruba aplicação
- [ ] Falha de job não interrompe outros jobs
- [ ] Tratamento de panic implementado

---

## Observabilidade

- [ ] Log de job iniciado implementado
- [ ] Log de job concluído implementado
- [ ] Log de job falhou implementado
- [ ] Log de duração implementado

---

## Segurança

- [ ] Dados sensíveis removidos dos logs
- [ ] Payloads protegidos
- [ ] Execuções auditáveis

---

## Testes

- [ ] Testes do scheduler implementados
- [ ] Testes de registro de jobs implementados
- [ ] Testes de execução implementados
- [ ] Testes de falha implementados

---

## Qualidade

- [ ] Scheduler desacoplado das funcionalidades
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Base preparada para novos jobs futuros