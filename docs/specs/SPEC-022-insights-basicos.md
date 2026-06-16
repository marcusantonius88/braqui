# SPEC-022 - Insights Básicos

## Objetivo

Definir a geração de insights simples baseados no histórico de eventos do pet.

O foco desta spec é:
- gerar valor a partir dos dados registrados;
- identificar padrões simples;
- aumentar engajamento;
- transformar eventos em informação útil.

---

# Contexto

Registrar eventos é importante.

Mas o verdadeiro valor surge quando o Braqui consegue ajudar o tutor a perceber padrões.

Exemplos:

```text
Thor apresentou três episódios de vômito este mês.
```

```text
Thor está mais ofegante do que o normal.
```

```text
Não há registros de medicação nos últimos 60 dias.
```

---

# Escopo

## O sistema deve:

- analisar eventos históricos;
- gerar insights simples;
- apresentar mensagens amigáveis;
- utilizar regras explícitas.

---

# Fora do Escopo

Esta spec NÃO contempla:

- machine learning;
- diagnóstico veterinário;
- recomendações médicas;
- IA preditiva;
- análise estatística avançada.

---

# Filosofia Arquitetural

Os insights devem ser:

- simples;
- explicáveis;
- auditáveis;
- previsíveis.

---

# Estratégia Inicial

Fluxo:

```text
Histórico
      ↓
Regras
      ↓
Insights
      ↓
Usuário
```

---

# Objetivo

Transformar:

```text
Dados
```

em

```text
Informação útil
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /insights
```

---

# Fonte dos Dados

Utilizar:

```text
EventRepository
```

---

# Tipos de Insights Iniciais

## Frequência de Vômito

Exemplo:

```text
Thor apresentou 3 episódios de vômito nos últimos 30 dias.
```

---

## Frequência de Coceira

Exemplo:

```text
Thor apresentou episódios frequentes de coceira recentemente.
```

---

## Ofegância Frequente

Exemplo:

```text
Thor apresentou muitos registros de ofegância nos últimos dias.
```

---

## Ausência de Medicação

Exemplo:

```text
Não encontrei registros recentes de antipulgas.
```

---

# Regras Iniciais

## Vômito

```text
>= 3 eventos em 30 dias
```

---

## Coceira

```text
>= 3 eventos em 30 dias
```

---

## Ofegância

```text
>= 5 eventos em 15 dias
```

---

## Medicação

```text
0 eventos em 60 dias
```

---

# Estrutura Conceitual

```go
type Insight struct {
    Type        string
    Message     string
    GeneratedAt time.Time
}
```

---

# Estratégia de Execução

Inicialmente:

- sob demanda.

---

# Exemplo

```text
Quais insights você encontrou?
```

---

# Evolução Futura

Possibilidade de geração automática.

---

# Integração com Timeline

Insights devem utilizar os eventos já registrados.

---

# Fluxo Esperado

```text
Usuário solicita insights
      ↓
Consulta histórico
      ↓
Executa regras
      ↓
Retorna insights
```

---

# Resposta ao Usuário

Exemplo:

```text
📊 Insights do Thor

• Foram registrados 3 episódios de vômito nos últimos 30 dias.

• Foram registrados 5 episódios de ofegância nos últimos 15 dias.
```

---

# Observabilidade

Registrar:

- geração de insights;
- quantidade de insights encontrados;
- falhas de processamento.

---

# NÃO registrar

- histórico completo;
- dados sensíveis.

---

# Tratamento de Erros

## Sem eventos

Responder:

```text
Ainda não há dados suficientes para gerar insights.
```

---

## Nenhum insight encontrado

Responder:

```text
Não encontrei padrões relevantes no momento.
```

---

# Critérios de Aceite

## Análise

- histórico analisado corretamente.

---

## Regras

- insights gerados corretamente.

---

## Resposta

- insights apresentados de forma amigável.

---

## Arquitetura

- regras desacopladas da persistência.

---

# Requisitos Técnicos

## Deve existir

- InsightService;
- regras de análise;
- geração de insights;
- integração com histórico.

---

# Dependências

Relaciona-se com:
- SPEC-016 - Registro de Eventos
- SPEC-017 - Timeline de Eventos
- SPEC-014 - Parser Local de Eventos

---

# Considerações Arquiteturais

## Regras explícitas

Todo insight deve ser explicável.

---

## Sem diagnósticos

O sistema NÃO deve:
- diagnosticar;
- prescrever;
- substituir veterinários.

---

## Compatibilidade com Monorepo

O módulo de insights deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por futuros canais;
- evitar dependência direta do Telegram.

---

# Objetivo Real do MVP

O foco é:
- gerar valor com os dados;
- aumentar retenção;
- demonstrar inteligência do produto.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- IA para insights;
- padrões avançados;
- correlações climáticas;
- análise comportamental;
- score de saúde.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/insights` criada
- [ ] InsightService criado
- [ ] Casos de uso de insights criados

---

## Modelo

- [ ] Estrutura Insight criada
- [ ] Campo Type implementado
- [ ] Campo Message implementado
- [ ] Campo GeneratedAt implementado

---

## Integração com Histórico

- [ ] Integração com EventRepository implementada
- [ ] Consulta de eventos implementada
- [ ] Filtragem por período implementada

---

## Insight de Vômito

- [ ] Regra de >= 3 eventos em 30 dias implementada
- [ ] Mensagem amigável implementada

---

## Insight de Coceira

- [ ] Regra de >= 3 eventos em 30 dias implementada
- [ ] Mensagem amigável implementada

---

## Insight de Ofegância

- [ ] Regra de >= 5 eventos em 15 dias implementada
- [ ] Mensagem amigável implementada

---

## Insight de Medicação

- [ ] Regra de ausência de medicação em 60 dias implementada
- [ ] Mensagem amigável implementada

---

## Geração

- [ ] Geração sob demanda implementada
- [ ] Agregação de múltiplos insights implementada
- [ ] Ordenação de insights implementada

---

## Fluxo Conversacional

- [ ] Comando de insights implementado
- [ ] Integração com Router implementada

---

## Respostas

- [ ] Template de resposta criado
- [ ] Formatação amigável implementada

---

## Tratamento de Erros

- [ ] Cenário sem eventos implementado
- [ ] Cenário sem insights implementado
- [ ] Falhas de processamento tratadas

---

## Observabilidade

- [ ] Log de geração implementado
- [ ] Log de quantidade de insights implementado
- [ ] Log de falhas implementado

---

## Testes

- [ ] Testes de regras implementados
- [ ] Testes de geração implementados
- [ ] Testes de agregação implementados
- [ ] Testes de cenários vazios implementados

---

## Qualidade

- [ ] Regras desacopladas da persistência
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Insights explicáveis e auditáveis