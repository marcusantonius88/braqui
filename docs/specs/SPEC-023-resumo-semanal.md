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

- [ ] Estrutura `/summary` criada
- [ ] SummaryService criado
- [ ] WeeklySummaryJob criado

---

## Modelo

- [ ] Estrutura WeeklySummary criada
- [ ] Campo TotalEvents implementado
- [ ] Campo StartDate implementado
- [ ] Campo EndDate implementado

---

## Integração com Histórico

- [ ] Integração com EventRepository implementada
- [ ] Consulta dos últimos 7 dias implementada
- [ ] Recuperação de eventos por pet implementada

---

## Agregação

- [ ] Contagem total de eventos implementada
- [ ] Contagem de vômitos implementada
- [ ] Contagem de coceiras implementada
- [ ] Contagem de ofegância implementada
- [ ] Contagem de medicações implementada
- [ ] Contagem de consultas implementada

---

## Formatação

- [ ] Template de resumo criado
- [ ] Mensagem amigável implementada
- [ ] Formatação textual implementada

---

## Scheduler

- [ ] Integração com WeeklySummaryJob implementada
- [ ] Execução semanal implementada
- [ ] Registro do job implementado

---

## Telegram

- [ ] Envio do resumo implementado
- [ ] Integração com TelegramGateway implementada
- [ ] Tratamento de falha de envio implementado

---

## Regras de Negócio

- [ ] Cenário sem eventos implementado
- [ ] Cenário com eventos implementado
- [ ] Usuário sem pet tratado

---

## Integração Futura

- [ ] Estrutura preparada para Insights
- [ ] Estrutura preparada para IA futura

---

## Observabilidade

- [ ] Log de resumo gerado implementado
- [ ] Log de quantidade de eventos implementado
- [ ] Log de envio implementado
- [ ] Log de falhas implementado

---

## Tratamento de Erros

- [ ] Falha de consulta tratada
- [ ] Falha de envio tratada
- [ ] Recuperação operacional implementada

---

## Testes

- [ ] Testes de agregação implementados
- [ ] Testes de formatação implementados
- [ ] Testes de scheduler implementados
- [ ] Testes de envio implementados
- [ ] Testes de cenários sem eventos implementados

---

## Qualidade

- [ ] SummaryService desacoplado da persistência
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Base preparada para evolução futura