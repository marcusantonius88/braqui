# SPEC-017 - Timeline de Eventos

## Objetivo

Definir a funcionalidade de consulta do histórico de eventos do pet através de uma timeline conversacional.

O foco desta spec é:
- permitir visualização do histórico;
- facilitar acompanhamento do pet;
- dar contexto para o tutor;
- servir de base para insights futuros.

---

# Contexto

Após o registro contínuo de eventos o tutor precisa conseguir visualizar o histórico acumulado.

Exemplos:

```text
Quando foi a última vez que Thor vomitou?
```

```text
Quais foram os últimos eventos registrados?
```

```text
Mostre o histórico do Thor
```

---

# Escopo

## O sistema deve:

- consultar eventos registrados;
- ordenar eventos cronologicamente;
- exibir histórico resumido;
- permitir consultas simples.

---

# Fora do Escopo

Esta spec NÃO contempla:

- dashboards gráficos;
- filtros avançados;
- exportação;
- relatórios complexos;
- analytics avançado.

---

# Filosofia Arquitetural

A timeline deve ser:

- simples;
- legível;
- cronológica;
- conversacional.

---

# Estratégia Inicial

A timeline será apresentada via Telegram.

---

# Fluxo Esperado

```text
Usuário
      ↓
Solicita histórico
      ↓
Braqui consulta eventos
      ↓
Braqui responde timeline
```

---

# Fonte dos Dados

A timeline utilizará:

```text
EventRepository
```

---

# Ordenação

Os eventos devem ser ordenados por:

```text
created_at DESC
```

---

# Exemplo de Resposta

```text
📋 Histórico recente do Thor

• Hoje - Ofegante
• Ontem - Tomou Simparic
• 3 dias atrás - Vômito
• 5 dias atrás - Consulta veterinária
```

---

# Quantidade Inicial

Inicialmente:

```text
últimos 10 eventos
```

---

# Estrutura Esperada

```text
TimelineService
```

---

# Responsabilidade

Converter:

```text
Eventos estruturados
```

em

```text
Resumo amigável
```

---

# Tipos de Eventos Suportados

## Saúde

```text
vomit
itching
panting
fatigue
cough
diarrhea
```

---

## Medicação

```text
medication_given
```

---

## Consulta

```text
vet_visit
```

---

# Formatação

A timeline deve:

- ser curta;
- ser fácil de ler;
- evitar excesso de detalhes.

---

# Exemplo de Evento

Entrada:

```json
{
  "type": "vomit",
  "created_at": "2026-01-01"
}
```

---

# Saída

```text
1 dia atrás - Vômito
```

---

# Consulta por Tipo

Preparação futura:

```text
Mostrar apenas vômitos
```

---

# MVP

Inicialmente:

- histórico geral.

---

# Fluxo Conversacional

Exemplos de gatilho:

```text
histórico
```

```text
timeline
```

```text
últimos eventos
```

---

# Integração com Router

O router deve encaminhar consultas de histórico para a timeline.

---

# Observabilidade

Registrar:

- consulta realizada;
- quantidade de eventos retornados;
- falhas de consulta.

---

# NÃO registrar

- histórico completo nos logs;
- informações sensíveis.

---

# Tratamento de Erros

## Sem eventos

Responder:

```text
Ainda não encontrei eventos registrados para o Thor 🐶
```

---

## Falha de persistência

Responder:

```text
Não consegui consultar o histórico agora.
Tente novamente em alguns minutos.
```

---

# Critérios de Aceite

## Consulta

- histórico recuperado corretamente.

---

## Ordenação

- eventos ordenados corretamente.

---

## Formatação

- timeline amigável ao usuário.

---

## Arquitetura

- timeline desacoplada da persistência.

---

# Requisitos Técnicos

## Deve existir

- TimelineService;
- integração com EventRepository;
- formatação amigável;
- consulta cronológica.

---

# Dependências

Relaciona-se com:
- SPEC-016 - Registro de Eventos
- SPEC-013 - Router Conversacional
- SPEC-009 - Integração com Telegram

---

# Considerações Arquiteturais

## Histórico como produto

O histórico é uma funcionalidade central do Braqui.

Sem histórico:
- não existem insights;
- não existe acompanhamento.

---

## Serviço dedicado

A lógica de montagem da timeline deve ficar isolada.

---

## Compatibilidade com Monorepo

A timeline deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por futuros canais;
- evitar dependência direta do Telegram.

---

# Objetivo Real do MVP

O foco é:
- permitir visualização do histórico;
- aumentar percepção do tutor;
- gerar valor imediato.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- filtros;
- busca por período;
- timeline visual;
- dashboard web;
- exportação PDF;
- histórico avançado.

---

# Implementation Checklist

## Estrutura Base

- [x] TimelineService criado
- [x] Casos de uso de consulta criados
- [x] Integração com EventRepository criada

---

## Consulta

- [x] Consulta dos últimos eventos implementada
- [x] Limite inicial de 10 eventos implementado
- [x] Recuperação por pet implementada

---

## Ordenação

- [x] Ordenação por created_at DESC implementada
- [x] Ordenação validada por testes

---

## Formatação

- [x] Formatação amigável implementada
- [x] Conversão de tipos de eventos implementada
- [x] Conversão de datas relativas implementada

---

## Tipos de Evento

- [x] Exibição de vomit implementada
- [x] Exibição de itching implementada
- [x] Exibição de panting implementada
- [x] Exibição de fatigue implementada
- [x] Exibição de cough implementada
- [x] Exibição de diarrhea implementada
- [x] Exibição de medication_given implementada
- [x] Exibição de vet_visit implementada

---

## Fluxo Conversacional

- [x] Comando "histórico" implementado
- [x] Comando "timeline" implementado
- [x] Comando "últimos eventos" implementado

---

## Integração com Router

- [x] Roteamento para timeline implementado
- [x] Fluxo de consulta integrado

---

## Tratamento de Erros

- [x] Cenário sem eventos implementado
- [x] Falha de consulta tratada
- [x] Mensagens amigáveis implementadas

---

## Observabilidade

- [x] Log de consulta implementado
- [x] Log de quantidade de eventos retornados implementado
- [x] Log de falha de consulta implementado

---

## Testes

- [x] Testes de consulta implementados
- [x] Testes de ordenação implementados
- [x] Testes de formatação implementados
- [x] Testes de cenários sem eventos implementados

---

## Qualidade

- [x] Timeline desacoplada da persistência
- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development