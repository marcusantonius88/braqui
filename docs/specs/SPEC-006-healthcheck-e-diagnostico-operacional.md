# SPEC-006 - Healthcheck e Diagnóstico Operacional

## Objetivo

Definir mecanismos básicos de healthcheck e diagnóstico operacional para o Braqui MVP.

O foco desta spec é:
- facilitar troubleshooting;
- simplificar monitoramento;
- melhorar observabilidade operacional;
- acelerar identificação de falhas.

---

# Contexto

O Braqui dependerá de:

- PostgreSQL;
- Telegram;
- scheduler;
- APIs externas;
- infraestrutura cloud.

Mesmo sendo um MVP, será necessário identificar rapidamente:

- falhas de inicialização;
- falhas de conexão;
- falhas de dependências;
- indisponibilidades.

---

# Escopo

## O sistema deve:

- expor endpoint de healthcheck;
- validar dependências críticas;
- fornecer informações básicas de diagnóstico;
- facilitar troubleshooting.

---

# Fora do Escopo

Esta spec NÃO contempla:

- observabilidade enterprise;
- OpenTelemetry;
- Prometheus;
- Grafana;
- distributed tracing;
- APM avançado.

---

# Filosofia Arquitetural

O healthcheck deve ser:
- simples;
- rápido;
- previsível;
- fácil de consumir.

---

# Endpoint Principal

O sistema deve expor:

```text
GET /health
```

---

# Resposta Esperada

```json
{
  "status": "ok"
}
```

---

# Objetivo

Permitir que:
- Docker;
- plataformas cloud;
- desenvolvedores;

validem rapidamente se a aplicação está viva.

---

# Healthcheck de Aplicação

Deve validar:

- aplicação iniciada;
- runtime funcional.

---

# Healthcheck de Banco

Opcional inicialmente.

Quando implementado deve validar:

- conexão PostgreSQL;
- capacidade de comunicação.

---

# Exemplo Futuro

```json
{
  "status": "ok",
  "database": "ok"
}
```

---

# Readiness

Inicialmente:
- não necessário.

---

# Liveness

Inicialmente:
- endpoint único.

---

# Diagnóstico Operacional

O sistema deve fornecer:

- logs de startup;
- logs de erro;
- logs de dependências críticas.

---

# Informações de Inicialização

Exemplo:

```text
starting braqui
environment: production
database connected
telegram initialized
```

---

# NÃO fazer

Não registrar:

- tokens;
- secrets;
- credenciais;
- payloads sensíveis.

---

# Tratamento de Falhas

## Banco indisponível

A aplicação deve:
- registrar erro;
- falhar rapidamente quando necessário.

---

## Configuração inválida

A aplicação deve:
- impedir startup;
- registrar erro amigável.

---

## Dependência externa indisponível

A aplicação deve:
- registrar erro;
- permitir troubleshooting.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /interfaces
        /http
```

---

# Handler de Healthcheck

Responsável por:

- responder healthcheck;
- expor status da aplicação.

---

# Exemplo Conceitual

```go
type HealthHandler struct{}
```

---

# Critérios de Aceite

## Endpoint

- endpoint `/health` disponível.

---

## Resposta

- endpoint retorna sucesso corretamente.

---

## Logs

- startup gera logs úteis.

---

## Troubleshooting

- falhas são facilmente identificáveis.

---

# Requisitos Técnicos

## Deve existir

- endpoint healthcheck;
- logs básicos;
- mensagens claras de erro.

---

# Dependências

Relaciona-se com:
- SPEC-002 - Configuração e Gerenciamento de Ambiente
- SPEC-003 - Deploy e Infraestrutura Inicial
- SPEC-004 - Containerização e Ambiente Local

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- métricas avançadas;
- tracing;
- dashboards.

---

## Diagnóstico rápido

O foco é:
- descobrir problemas rapidamente;
- reduzir tempo de troubleshooting.

---

## Compatibilidade com Monorepo

O healthcheck deve:
- estar isolado em `/apps/api`;
- não depender de futuras aplicações;
- ser facilmente reutilizável.

---

# Objetivo Real do MVP

O foco é:
- facilitar operação;
- facilitar deploy;
- melhorar suporte;
- melhorar experiência de desenvolvimento.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- readiness endpoint;
- métricas;
- Prometheus;
- Grafana;
- OpenTelemetry;
- tracing distribuído;
- dashboards operacionais.

---

# Implementation Checklist

## Endpoint

- [x] Endpoint `GET /health` criado
- [x] Handler de healthcheck implementado
- [x] Rota registrada

---

## Resposta

- [x] Resposta JSON padronizada implementada
- [x] Status HTTP correto implementado
- [x] Estrutura de resposta documentada

---

## Inicialização

- [x] Log de startup implementado
- [x] Log de ambiente implementado
- [x] Log de inicialização de dependências implementado

---

## Diagnóstico

- [x] Mensagens de erro amigáveis implementadas
- [x] Troubleshooting básico implementado
- [x] Logs operacionais implementados

---

## Banco de Dados

- [x] Verificação básica de conexão implementada
- [x] Healthcheck de PostgreSQL preparado para evolução futura

---

## Segurança

- [x] Secrets removidos dos logs
- [x] Credenciais protegidas
- [x] Payloads sensíveis não registrados

---

## Docker e Deploy

- [x] Endpoint compatível com Docker Healthcheck
- [x] Endpoint compatível com provedores cloud
- [x] Endpoint validado em ambiente local

---

## Testes

- [x] Testes do endpoint implementados
- [x] Testes de resposta implementados
- [x] Testes de falha implementados

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development