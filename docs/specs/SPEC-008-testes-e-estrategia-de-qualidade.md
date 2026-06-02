# SPEC-008 - Testes e Estratégia de Qualidade

## Objetivo

Definir a estratégia de testes e qualidade do Braqui MVP.

O foco desta spec é:
- reduzir regressões;
- aumentar confiança nas mudanças;
- facilitar refatorações;
- garantir estabilidade das funcionalidades críticas.

---

# Contexto

O Braqui será desenvolvido com:
- Spec Driven Development (SDD);
- IA-assisted development;
- implementação incremental.

Nesse contexto, testes são fundamentais para:
- validar comportamento;
- reduzir erros introduzidos por IA;
- garantir evolução segura.

---

# Escopo

## O sistema deve possuir:

- testes unitários;
- testes de integração;
- estratégia de mocks;
- cobertura dos fluxos críticos.

---

# Fora do Escopo

Esta spec NÃO contempla:

- testes de carga;
- testes de performance;
- testes de segurança avançados;
- testes de caos;
- cobertura extrema.

---

# Filosofia Arquitetural

Os testes devem ser:

- simples;
- rápidos;
- confiáveis;
- fáceis de manter.

---

# Estratégia Inicial

Priorizar:

1. testes unitários
2. testes de integração
3. testes E2E futuramente

---

# Pirâmide de Testes

Prioridade:

```text
Testes Unitários
↑↑↑↑↑

Testes Integração
↑↑

Testes E2E
↑
```

---

# Estrutura Esperada

```text
/tests
  /integration
  /e2e
```

---

# Testes Unitários

Devem validar:

- entidades;
- casos de uso;
- parsers;
- regras de negócio;
- serviços internos.

---

# Exemplo

```text
CreatePet
RegisterEvent
GenerateInsights
```

---

# O que NÃO testar

Evitar:

- banco real;
- Telegram real;
- APIs externas reais.

---

# Mocks

Dependências externas devem ser mockadas.

---

# Exemplos

```text
PetRepository
TelegramGateway
AIProvider
ClimateProvider
```

---

# Objetivo dos Mocks

Permitir:

- velocidade;
- isolamento;
- previsibilidade.

---

# Testes de Integração

Devem validar:

- PostgreSQL;
- repositories;
- migrations;
- persistência.

---

# Ambiente

Preferencialmente:

- PostgreSQL em Docker.

---

# Testes E2E

Inicialmente:
- opcionais.

---

# Futuramente

Poderão validar:

- fluxo Telegram;
- onboarding;
- lembretes;
- registro de eventos.

---

# Cobertura Prioritária

## Alta Prioridade

- onboarding;
- parser de eventos;
- lembretes;
- estado conversacional;
- insights.

---

## Média Prioridade

- integração Telegram;
- scheduler;
- clima.

---

## Baixa Prioridade

- logs;
- observabilidade.

---

# Estratégia para IA-Assisted Development

Toda funcionalidade implementada por IA deve possuir:

- validação;
- testes relevantes;
- cobertura mínima adequada.

---

# Critérios de Aceite

## Unitários

- regras críticas cobertas.

---

## Integração

- persistência validada.

---

## Mocks

- dependências externas isoladas.

---

## Qualidade

- regressões reduzidas.

---

# Requisitos Técnicos

## Deve existir

- estrutura de testes;
- mocks;
- testes unitários;
- testes de integração.

---

# Dependências

Relaciona-se com:
- SPEC-001 - Bootstrap e Estrutura Inicial do Projeto
- SPEC-005 - Persistência e Repositories
- SPEC-007 - Observabilidade Básica

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- cobertura de 100%;
- frameworks complexos;
- pipelines sofisticados.

---

## Velocidade

Os testes devem:
- executar rapidamente;
- facilitar feedback rápido.

---

## Compatibilidade com Monorepo

A estratégia de testes deve:
- funcionar dentro de `/apps/api`;
- suportar futuras aplicações;
- permitir expansão gradual.

---

# Objetivo Real do MVP

O foco é:
- aumentar confiança;
- reduzir regressões;
- melhorar qualidade;
- facilitar evolução incremental.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- cobertura avançada;
- testes E2E completos;
- testes de carga;
- testes de performance;
- quality gates;
- mutation testing.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/tests` criada
- [ ] Estrutura `/tests/integration` criada
- [ ] Estrutura `/tests/e2e` criada

---

## Testes Unitários

- [ ] Estratégia de testes unitários definida
- [ ] Testes de entidades implementados
- [ ] Testes de casos de uso implementados
- [ ] Testes de parsers implementados
- [ ] Testes de regras de negócio implementados

---

## Mocks

- [ ] Mock de PetRepository criado
- [ ] Mock de EventRepository criado
- [ ] Mock de ReminderRepository criado
- [ ] Mock de TelegramGateway criado
- [ ] Mock de AIProvider criado
- [ ] Mock de ClimateProvider criado

---

## Testes de Integração

- [ ] Estratégia de integração definida
- [ ] Testes de PostgreSQL implementados
- [ ] Testes de repositories implementados
- [ ] Testes de migrations implementados

---

## Fluxos Críticos

- [ ] Testes de onboarding implementados
- [ ] Testes de parser de eventos implementados
- [ ] Testes de lembretes implementados
- [ ] Testes de estado conversacional implementados
- [ ] Testes de insights implementados

---

## Dependências Externas

- [ ] Telegram isolado em testes
- [ ] IA isolada em testes
- [ ] Clima isolado em testes
- [ ] Banco real evitado em testes unitários

---

## Qualidade

- [ ] Estratégia de regressão definida
- [ ] Cobertura mínima dos fluxos críticos atingida
- [ ] Execução rápida dos testes validada

---

## Integração com Desenvolvimento

- [ ] Estratégia compatível com SDD
- [ ] Estratégia compatível com OpenCode
- [ ] Estratégia compatível com IA-assisted development

---

## Futuro

- [ ] Base preparada para testes E2E
- [ ] Base preparada para quality gates futuros

---

## Documentação

- [ ] Estratégia de testes documentada
- [ ] Convenções de testes documentadas
- [ ] Boas práticas de testes documentadas