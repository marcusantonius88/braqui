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

Após o registro contínuo de eventos, o tutor precisa conseguir visualizar o histórico acumulado.

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

- [ ] TimelineService criado
- [ ] Casos de uso de consulta criados
- [ ] Integração com EventRepository criada

---

## Consulta

- [ ] Consulta dos últimos eventos implementada
- [ ] Limite inicial de 10 eventos implementado
- [ ] Recuperação por pet implementada

---

## Ordenação

- [ ] Ordenação por created_at DESC implementada
- [ ] Ordenação validada por testes

---

## Formatação

- [ ] Formatação amigável implementada
- [ ] Conversão de tipos de eventos implementada
- [ ] Conversão de datas relativas implementada

---

## Tipos de Evento

- [ ] Exibição de vomit implementada
- [ ] Exibição de itching implementada
- [ ] Exibição de panting implementada
- [ ] Exibição de fatigue implementada
- [ ] Exibição de cough implementada
- [ ] Exibição de diarrhea implementada
- [ ] Exibição de medication_given implementada
- [ ] Exibição de vet_visit implementada

---

## Fluxo Conversacional

- [ ] Comando "histórico" implementado
- [ ] Comando "timeline" implementado
- [ ] Comando "últimos eventos" implementado

---

## Integração com Router

- [ ] Roteamento para timeline implementado
- [ ] Fluxo de consulta integrado

---

## Tratamento de Erros

- [ ] Cenário sem eventos implementado
- [ ] Falha de consulta tratada
- [ ] Mensagens amigáveis implementadas

---

## Observabilidade

- [ ] Log de consulta implementado
- [ ] Log de quantidade de eventos retornados implementado
- [ ] Log de falha de consulta implementado

---

## Testes

- [ ] Testes de consulta implementados
- [ ] Testes de ordenação implementados
- [ ] Testes de formatação implementados
- [ ] Testes de cenários sem eventos implementados

---

## Qualidade

- [ ] Timeline desacoplada da persistência
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development