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

- [ ] Estrutura do módulo `event` criada
- [ ] Interface de parser criada
- [ ] Implementação inicial do parser criada

---

## Modelo de Resultado

- [ ] Estrutura ParseResult criada
- [ ] Campo type implementado
- [ ] Campo confidence implementado
- [ ] Campo payload implementado

---

## Tipos de Eventos

- [ ] Evento vomit implementado
- [ ] Evento diarrhea implementado
- [ ] Evento itching implementado
- [ ] Evento cough implementado
- [ ] Evento fatigue implementado
- [ ] Evento panting implementado
- [ ] Evento medication_given implementado
- [ ] Evento weight_update implementado
- [ ] Evento vet_visit implementado

---

## Regras de Matching

- [ ] Matching por palavras-chave implementado
- [ ] Matching por expressões implementado
- [ ] Regras explícitas documentadas

---

## Vômito

- [ ] Reconhecimento de "vomitou"
- [ ] Reconhecimento de "vomitando"
- [ ] Reconhecimento de "teve vômito"

---

## Coceira

- [ ] Reconhecimento de "coçando"
- [ ] Reconhecimento de "muita coceira"
- [ ] Reconhecimento de "coceira"

---

## Ofegância

- [ ] Reconhecimento de "ofegante"
- [ ] Reconhecimento de "respiração acelerada"

---

## Confiança

- [ ] Nível HIGH implementado
- [ ] Nível MEDIUM implementado
- [ ] Nível LOW implementado

---

## Fallback

- [ ] Resultado NOT_MATCHED implementado
- [ ] Encaminhamento para IA preparado
- [ ] Integração com router preparada

---

## Observabilidade

- [ ] Log de evento identificado implementado
- [ ] Log de confiança implementado
- [ ] Log de falha de interpretação implementado

---

## Tratamento de Erros

- [ ] Mensagem vazia tratada
- [ ] Evento desconhecido tratado
- [ ] Falhas internas tratadas

---

## Testes

- [ ] Testes de matching implementados
- [ ] Testes de confiança implementados
- [ ] Testes de fallback implementados
- [ ] Testes de eventos suportados implementados

---

## Qualidade

- [ ] Parser desacoplado da IA
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development