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

- [ ] Endpoint `GET /health` criado
- [ ] Handler de healthcheck implementado
- [ ] Rota registrada

---

## Resposta

- [ ] Resposta JSON padronizada implementada
- [ ] Status HTTP correto implementado
- [ ] Estrutura de resposta documentada

---

## Inicialização

- [ ] Log de startup implementado
- [ ] Log de ambiente implementado
- [ ] Log de inicialização de dependências implementado

---

## Diagnóstico

- [ ] Mensagens de erro amigáveis implementadas
- [ ] Troubleshooting básico implementado
- [ ] Logs operacionais implementados

---

## Banco de Dados

- [ ] Verificação básica de conexão implementada
- [ ] Healthcheck de PostgreSQL preparado para evolução futura

---

## Segurança

- [ ] Secrets removidos dos logs
- [ ] Credenciais protegidas
- [ ] Payloads sensíveis não registrados

---

## Docker e Deploy

- [ ] Endpoint compatível com Docker Healthcheck
- [ ] Endpoint compatível com provedores cloud
- [ ] Endpoint validado em ambiente local

---

## Testes

- [ ] Testes do endpoint implementados
- [ ] Testes de resposta implementados
- [ ] Testes de falha implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development