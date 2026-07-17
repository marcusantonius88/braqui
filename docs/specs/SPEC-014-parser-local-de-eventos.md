# SPEC-014 - Parser Local de Eventos

## Objetivo

Definir a estratégia de interpretação local de mensagens relacionadas à saúde e rotina do pet.

O foco desta spec é:
- reduzir dependência de IA;
- reduzir custos operacionais;
- aumentar previsibilidade;
- identificar eventos comuns de forma rápida.

---

# Contexto

Grande parte das mensagens dos usuários seguirá padrões simples.

Exemplos:

```text
Thor vomitou
```

```text
Thor tomou simparic
```

```text
Thor está coçando muito
```

```text
Thor teve diarreia
```

Não faz sentido utilizar IA para todos os casos.

---

# Escopo

## O sistema deve:

- identificar eventos conhecidos;
- extrair informações básicas;
- classificar eventos;
- produzir resultado estruturado.

---

# Fora do Escopo

Esta spec NÃO contempla:

- IA;
- interpretação semântica avançada;
- memória contextual;
- análise clínica;
- inferências complexas.

---

# Filosofia Arquitetural

Parser local deve ser:

- rápido;
- barato;
- previsível;
- determinístico.

---

# Estratégia Inicial

Sempre tentar:

```text
Parser Local
      ↓
Sucesso?
      ↓
SIM → registra evento
NÃO → IA
```

---

# Benefícios

- menor custo;
- menor latência;
- maior previsibilidade.

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /event
```

---

# Entrada

Exemplo:

```text
Thor vomitou hoje
```

---

# Saída

Exemplo:

```json
{
  "type": "vomit",
  "confidence": "high"
}
```

---

# Tipos de Eventos Iniciais

## Saúde

```text
vomit
diarrhea
itching
cough
fatigue
panting
```

---

## Medicação

```text
medication_given
```

---

## Peso

```text
weight_update
```

---

## Consulta

```text
vet_visit
```

---

# Estratégia de Matching

Inicialmente:

- palavras-chave;
- expressões conhecidas;
- regras explícitas.

---

# Exemplos

## Vômito

```text
vomitou
vomitando
teve vômito
```

---

## Coceira

```text
coçando
muita coceira
coceira
```

---

## Ofegância

```text
ofegante
respiração acelerada
```

---

# Resultado da Interpretação

Deve retornar:

```text
tipo
confiança
payload
```

---

# Exemplo

```json
{
  "type": "itching",
  "confidence": "high",
  "payload": {}
}
```

---

# Confiança

Valores:

```text
high
medium
low
```

---

# Fallback

Caso não encontre correspondência:

```text
NOT_MATCHED
```

---

# Fluxo

```text
Mensagem
      ↓
Parser
      ↓
Evento identificado?
      ↓
SIM → registra
NÃO → IA
```

---

# Integração com Router

O router poderá utilizar:

```text
ParseResult
```

para decidir próximos passos.

---

# Observabilidade

Registrar:

- evento identificado;
- confiança;
- falhas de interpretação.

---

# NÃO registrar

- conteúdo completo da conversa;
- dados sensíveis.

---

# Tratamento de Erros

## Mensagem vazia

Retornar:

```text
NOT_MATCHED
```

---

## Evento desconhecido

Retornar:

```text
NOT_MATCHED
```

---

# Critérios de Aceite

## Interpretação

- eventos comuns identificados corretamente.

---

## Estrutura

- resultado estruturado retornado.

---

## Fallback

- mensagens desconhecidas encaminhadas corretamente.

---

## Arquitetura

- parser desacoplado da IA.

---

# Requisitos Técnicos

## Deve existir

- parser;
- regras explícitas;
- tipos de eventos;
- resultado estruturado.

---

# Dependências

Relaciona-se com:
- SPEC-013 - Router Conversacional
- SPEC-016 - Registro de Eventos
- SPEC-015 - Provider de IA

---

# Considerações Arquiteturais

## Local First

O parser deve sempre ser tentado antes da IA.

---

## IA como fallback

IA só deve ser utilizada:
- quando necessário;
- quando o parser falhar.

---

## Compatibilidade com Monorepo

O parser deve:
- permanecer isolado em `/apps/api`;
- ser reutilizável por futuros canais;
- não depender diretamente do Telegram.

---

# Objetivo Real do MVP

O foco é:
- reduzir custo;
- reduzir latência;
- aumentar previsibilidade;
- registrar eventos rapidamente.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- NLP local;
- classificação avançada;
- extração de entidades;
- modelos locais;
- regras especializadas por raça.

---

# Implementation Checklist

## Estrutura Base

- [x] Estrutura do módulo `parser` criada
- [x] Interface de parser criada (função pública `Parse`)
- [x] Implementação inicial do parser criada

---

## Modelo de Resultado

- [x] Estrutura ParseResult criada
- [x] Campo type implementado
- [x] Campo confidence implementado
- [x] Campo payload implementado

---

## Tipos de Eventos

- [x] Evento vomit implementado
- [x] Evento diarrhea implementado
- [x] Evento itching implementado
- [x] Evento cough implementado
- [x] Evento fatigue implementado
- [x] Evento panting implementado
- [x] Evento medication_given implementado
- [x] Evento weight_update implementado
- [x] Evento vet_visit implementado

---

## Regras de Matching

- [x] Matching por palavras-chave implementado
- [x] Matching por expressões implementado
- [x] Regras explícitas documentadas (em código)

---

## Vômito

- [x] Reconhecimento de "vomitou"
- [x] Reconhecimento de "vomitando"
- [x] Reconhecimento de "teve vômito"

---

## Coceira

- [x] Reconhecimento de "coçando"
- [x] Reconhecimento de "muita coceira"
- [x] Reconhecimento de "coceira"

---

## Ofegância

- [x] Reconhecimento de "ofegante"
- [x] Reconhecimento de "respiração acelerada"

---

## Confiança

- [x] Nível HIGH implementado
- [x] Nível MEDIUM implementado
- [x] Nível LOW implementado

---

## Fallback

- [x] Resultado NOT_MATCHED implementado
- [x] Encaminhamento para IA preparado (estrutura pronta para SPEC-015)
- [x] Integração com router preparada (parser pode ser chamado pelo router)

---

## Observabilidade

- [x] Log de evento identificado implementado
- [x] Log de confiança implementado
- [x] Log de falha de interpretação implementado

---

## Tratamento de Erros

- [x] Mensagem vazia tratada
- [x] Evento desconhecido tratado
- [x] Falhas internas tratadas

---

## Testes

- [x] Testes de matching implementados
- [x] Testes de confiança implementados
- [x] Testes de fallback implementados
- [x] Testes de eventos suportados implementados

---

## Qualidade

- [x] Parser desacoplado da IA
- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development