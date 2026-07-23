# SPEC-023 - Resumo Semanal

## Objetivo

Definir a funcionalidade de geração automática de resumos semanais do histórico do pet.

O foco desta spec é:
- aumentar percepção de valor do Braqui;
- manter o tutor engajado;
- destacar eventos importantes;
- transformar histórico em acompanhamento contínuo.

---

# Contexto

Ao longo da semana, diversos eventos podem ser registrados:

```text
Vômitos
```

```text
Coceiras
```

```text
Ofegância
```

```text
Medicações
```

```text
Consultas
```

O tutor dificilmente acompanhará tudo manualmente.

O Braqui deve consolidar essas informações.

---

# Escopo

## O sistema deve:

- consultar eventos da semana;
- gerar resumo amigável;
- enviar automaticamente ao tutor;
- utilizar scheduler.

---

# Fora do Escopo

Esta spec NÃO contempla:

- relatórios PDF;
- dashboards;
- gráficos;
- IA generativa avançada;
- análise preditiva.

---

# Filosofia Arquitetural

O resumo semanal deve ser:

- simples;
- objetivo;
- útil;
- fácil de consumir.

---

# Estratégia Inicial

Fluxo:

```text
Scheduler
      ↓
Consulta eventos
      ↓
Gera resumo
      ↓
Telegram
```

---

# Período

Inicialmente:

```text
últimos 7 dias
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /summary
```

---

# Fonte dos Dados

Utilizar:

```text
EventRepository
```

---

# Informações do Resumo

## Quantidade de Eventos

Exemplo:

```text
12 eventos registrados
```

---

## Vômitos

Exemplo:

```text
2 episódios de vômito
```

---

## Coceiras

Exemplo:

```text
4 episódios de coceira
```

---

## Ofegância

Exemplo:

```text
3 episódios de ofegância
```

---

## Medicações

Exemplo:

```text
1 medicação registrada
```

---

## Consultas

Exemplo:

```text
1 consulta veterinária registrada
```

---

# Estrutura Conceitual

```go
type WeeklySummary struct {
    TotalEvents int
    StartDate   time.Time
    EndDate     time.Time
}
```

---

# Exemplo de Resposta

```text
📊 Resumo semanal do Thor

• 12 eventos registrados

• 2 episódios de vômito
• 4 episódios de coceira
• 3 episódios de ofegância

• 1 medicação registrada

Continue registrando eventos para que eu possa acompanhar melhor a saúde do Thor 🐶
```

---

# Frequência

Inicialmente:

```text
1 vez por semana
```

---

# Scheduler

O resumo será executado através de:

```text
WeeklySummaryJob
```

---

# Regras de Negócio

## Sem eventos

Enviar:

```text
Nenhum evento foi registrado esta semana.
```

---

## Com eventos

Enviar resumo consolidado.

---

# Integração com Insights

Preparação futura:

```text
Resumo
      ↓
Insights
```

---

# MVP

Inicialmente:

- apenas agregação simples.

---

# Observabilidade

Registrar:

- resumo gerado;
- quantidade de eventos;
- envio realizado;
- falhas.

---

# NÃO registrar

- histórico completo;
- dados sensíveis.

---

# Tratamento de Erros

## Falha de consulta

Registrar erro.

---

## Falha de envio

Registrar erro.

---

## Usuário sem pet

Ignorar processamento.

---

# Critérios de Aceite

## Consulta

- eventos recuperados corretamente.

---

## Agregação

- totais calculados corretamente.

---

## Telegram

- resumo enviado corretamente.

---

## Arquitetura

- resumo desacoplado da persistência.

---

# Requisitos Técnicos

## Deve existir

- SummaryService;
- WeeklySummaryJob;
- agregação semanal;
- integração Telegram.

---

# Dependências

Relaciona-se com:
- SPEC-016 - Registro de Eventos
- SPEC-018 - Scheduler e Tarefas Agendadas
- SPEC-009 - Integração com Telegram
- SPEC-022 - Insights Básicos

---

# Considerações Arquiteturais

## Valor percebido

O resumo semanal é uma funcionalidade de retenção.

Ele lembra o tutor que:
- existe histórico;
- existe acompanhamento;
- existe inteligência no sistema.

---

## Simplicidade primeiro

O MVP NÃO precisa:
- IA;
- linguagem natural avançada;
- análises complexas.

---

## Compatibilidade com Monorepo

O módulo de resumo deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por futuros canais;
- evitar dependência direta do Telegram.

---

# Objetivo Real do MVP

O foco é:
- aumentar engajamento;
- reforçar valor do produto;
- consolidar histórico semanal.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- IA para geração textual;
- gráficos;
- PDF;
- insights avançados;
- comparações históricas;
- score semanal de saúde.

---

# Implementation Checklist

## Estrutura Base

- [x] Estrutura `/summary` criada (`internal/summary/`)
- [x] SummaryService criado (`Service.GenerateAndSend`)
- [x] WeeklySummaryJob criado (real, com `NewWeeklySummaryJob(svc, log)`)

---

## Modelo

- [x] Estrutura WeeklySummary criada (modelo implícito em `formatSummary`)
- [x] Campo TotalEvents implementado (contagem agregada)
- [x] Campo StartDate implementado (7 dias atrás via `time.Now().AddDate(0,0,-7)`)
- [x] Campo EndDate implementado (`time.Now()`)

---

## Integração com Histórico

- [x] Integração com EventRepository implementada
- [x] Consulta dos últimos 7 dias implementada (filtro in-memory via `Timestamp`)
- [x] Recuperação de eventos por pet implementada

---

## Agregação

- [x] Contagem total de eventos implementada
- [x] Contagem de vômitos implementada
- [x] Contagem de coceiras implementada
- [x] Contagem de ofegância implementada
- [x] Contagem de medicações implementada
- [x] Contagem de consultas implementada (vet_visit)

---

## Formatação

- [x] Template de resumo criado (`📊 Resumo semanal do {pet}\n\n• {total}\n\n• {cat1}\n• {cat2}\n\n{footer}`)
- [x] Mensagem amigável implementada (singular/plural corretos, labels em PT-BR)
- [x] Formatação textual implementada

---

## Scheduler

- [x] Integração com WeeklySummaryJob implementada
- [x] Execução semanal implementada (frequência `scheduler.Weekly`)
- [x] Registro do job implementado

---

## Telegram

- [x] Envio do resumo implementado
- [x] Integração com TelegramGateway implementada
- [x] Tratamento de falha de envio implementado (log + continue)

---

## Regras de Negócio

- [x] Cenário sem eventos implementado ("Nenhum evento foi registrado esta semana para o {pet}.")
- [x] Cenário com eventos implementado (resumo consolidado)
- [x] Usuário sem pet tratado (FindAll retorna vazio → 0 mensagens)

---

## Integração Futura

- [x] Estrutura preparada para Insights (SummaryService via interface SummaryService no scheduler)
- [x] Estrutura preparada para IA futura (serviço desacoplado via interface)

---

## Observabilidade

- [x] Log de resumo gerado implementado
- [x] Log de quantidade de eventos implementado
- [x] Log de envio implementado
- [x] Log de falhas implementado

---

## Tratamento de Erros

- [x] Falha de consulta tratada (log + continue)
- [x] Falha de envio tratada (log + continue)
- [x] Recuperação operacional implementada (panic recovery do scheduler)

---

## Testes

- [x] Testes de agregação implementados (contagem por tipo)
- [x] Testes de formatação implementados (singular/plural, estrutura do template)
- [x] Testes de scheduler implementados (job executa sem erro)
- [x] Testes de envio implementados (mensagens enviadas corretamente)
- [x] Testes de cenários sem eventos implementados

---

## Qualidade

- [x] SummaryService desacoplado da persistência (interfaces locais)
- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Base preparada para evolução futura