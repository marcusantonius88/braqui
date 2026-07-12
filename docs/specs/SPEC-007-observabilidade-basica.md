# SPEC-007 - Observabilidade Básica

## Objetivo

Definir a estratégia de observabilidade básica do Braqui MVP.

O foco desta spec é:
- facilitar troubleshooting;
- melhorar entendimento do comportamento da aplicação;
- reduzir tempo de diagnóstico;
- apoiar operação do sistema.

---

# Contexto

O Braqui será executado inicialmente:

- como monólito;
- em uma única instância;
- com baixo volume de usuários.

Mesmo assim será necessário entender:

- falhas;
- comportamento da aplicação;
- integrações;
- execução de jobs;
- fluxo conversacional.

---

# Escopo

## O sistema deve:

- registrar logs estruturados;
- registrar eventos relevantes;
- registrar erros operacionais;
- facilitar troubleshooting.

---

# Fora do Escopo

Esta spec NÃO contempla:

- OpenTelemetry;
- Prometheus;
- Grafana;
- tracing distribuído;
- observabilidade enterprise;
- dashboards avançados.

---

# Filosofia Arquitetural

Observabilidade deve ser:

- simples;
- útil;
- objetiva;
- de baixo custo operacional.

---

# Estratégia Inicial

Inicialmente:

- logs estruturados;
- saída padrão da aplicação;
- integração com logs da plataforma.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /infra
        /logger
```

---

# Logger Centralizado

A aplicação deve possuir:

- logger único;
- formato consistente;
- padronização de mensagens.

---

# Objetivos

Permitir:

- diagnóstico rápido;
- análise de falhas;
- acompanhamento operacional.

---

# Eventos que Devem Gerar Logs

## Startup

Exemplos:

```text
application started
environment loaded
database connected
telegram initialized
```

---

## Shutdown

Exemplos:

```text
application shutting down
```

---

## Integrações

Exemplos:

```text
telegram webhook received
telegram message processed
```

---

## Scheduler

Exemplos:

```text
reminder job executed
summary job executed
```

---

## Banco

Exemplos:

```text
database connection established
database connection failed
```

---

## Erros

Exemplos:

```text
unable to load configuration
unable to connect database
telegram processing failed
```

---

# NÃO fazer

Evitar logs excessivos.

---

# NÃO registrar

- tokens;
- secrets;
- credenciais;
- payloads completos;
- informações sensíveis.

---

# Estrutura de Log

Preferencialmente:

```json
{
  "level": "info",
  "message": "application started"
}
```

---

# Níveis de Log

Inicialmente:

- INFO
- WARN
- ERROR

---

# DEBUG

Inicialmente:
- opcional.

---

# Contextualização

Sempre que possível registrar:

- operação;
- módulo;
- erro;
- contexto mínimo necessário.

---

# Exemplo

```json
{
  "level": "error",
  "module": "telegram",
  "message": "unable to process message"
}
```

---

# Tratamento de Erros

Erros devem:

- ser registrados;
- possuir contexto suficiente;
- facilitar troubleshooting.

---

# Critérios de Aceite

## Logs

- logs estruturados funcionando.

---

## Startup

- eventos de inicialização registrados.

---

## Integrações

- eventos relevantes registrados.

---

## Erros

- falhas relevantes registradas.

---

# Requisitos Técnicos

## Deve existir

- logger centralizado;
- níveis de log;
- logs estruturados;
- logs de erro.

---

# Dependências

Relaciona-se com:
- SPEC-002 - Configuração e Gerenciamento de Ambiente
- SPEC-003 - Deploy e Infraestrutura Inicial
- SPEC-006 - Healthcheck e Diagnóstico Operacional

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- stack ELK;
- tracing distribuído;
- observabilidade avançada.

---

## Diagnóstico rápido

O foco é:
- identificar problemas rapidamente;
- reduzir tempo de investigação.

---

## Compatibilidade com Monorepo

A observabilidade deve:
- permanecer isolada em `/apps/api`;
- ser reutilizável pelos módulos;
- manter padronização em toda a aplicação.

---

# Objetivo Real do MVP

O foco é:
- melhorar suporte;
- melhorar operação;
- facilitar troubleshooting;
- reduzir tempo de resolução de problemas.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- OpenTelemetry;
- Grafana;
- Prometheus;
- métricas customizadas;
- tracing distribuído;
- dashboards operacionais;
- alertas automáticos.

---

# Implementation Checklist

## Estrutura Base

- [x] Estrutura `/infra/logger` criada
- [x] Logger centralizado criado
- [x] Interface de logging definida

---

## Configuração

- [x] Configuração de logger implementada
- [x] Níveis INFO implementados
- [x] Níveis WARN implementados
- [x] Níveis ERROR implementados
- [x] DEBUG preparado para evolução futura

---

## Startup e Shutdown

- [x] Log de startup implementado
- [ ] Log de shutdown implementado
- [x] Log de carregamento de ambiente implementado
- [x] Log de inicialização de dependências implementado

---

## Banco de Dados

- [x] Log de conexão com PostgreSQL implementado
- [x] Log de falha de conexão implementado

---

## Telegram (pendente — depende de SPEC-009)

- [ ] Log de recebimento de webhook implementado
- [ ] Log de processamento de mensagens implementado
- [ ] Log de falhas de processamento implementado

---

## Scheduler (pendente — depende de SPEC-018)

- [ ] Log de execução de jobs implementado
- [ ] Log de falha de jobs implementado

---

## Tratamento de Erros

- [x] Padronização de logs de erro implementada
- [x] Contextualização mínima de erros implementada
- [x] Troubleshooting básico suportado

---

## Segurança

- [x] Secrets removidos dos logs
- [x] Tokens removidos dos logs
- [x] Payloads sensíveis protegidos
- [x] Credenciais protegidas

---

## Estruturação

- [x] Logs estruturados implementados
- [x] Formato JSON implementado
- [x] Contexto de módulo implementado
- [x] Contexto de operação implementado

---

## Testes

- [x] Testes do logger implementados
- [x] Testes de formatação implementados
- [x] Testes de níveis de log implementados

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Logs consistentes em toda a aplicação