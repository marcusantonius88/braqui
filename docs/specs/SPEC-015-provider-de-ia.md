# SPEC-015 - Provider de IA

## Objetivo

Definir a estratégia de integração com Inteligência Artificial para o Braqui MVP.

O foco desta spec é:
- interpretar mensagens não reconhecidas pelo parser local;
- reduzir complexidade do sistema;
- manter baixo custo operacional;
- permitir evolução futura.

---

# Contexto

O Braqui utilizará uma abordagem híbrida:

```text
Parser Local
      ↓
Sucesso?
      ↓
SIM → continua
NÃO → IA
```

A IA não será a primeira opção.

Ela atuará como fallback inteligente.

---

# Escopo

## O sistema deve:

- possuir provider de IA;
- receber mensagens para interpretação;
- retornar estrutura padronizada;
- funcionar como fallback do parser local.

---

# Fora do Escopo

Esta spec NÃO contempla:

- memória longa;
- agentes autônomos;
- multi-agent;
- RAG;
- fine-tuning;
- embeddings.

---

# Filosofia Arquitetural

A IA deve ser:

- capability auxiliar;
- desacoplada;
- substituível;
- controlada.

---

# Estratégia Inicial

Provider inicial:

```text
Google Gemini
```

---

# Motivos

- free tier generoso;
- API simples;
- baixo custo;
- boa qualidade.

---

# Estratégia de Provider

O domínio NÃO deve conhecer Gemini.

---

# Interface Esperada

```go
type AIProvider interface {
    Interpret(
        ctx context.Context,
        message string,
    ) (*InterpretationResult, error)
}
```

---

# Implementação Inicial

```text
GeminiProvider
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /infra
        /ai
```

---

# Entrada

Exemplo:

```text
Thor parece mais cansado que o normal
```

---

# Saída

Exemplo:

```json
{
  "type": "fatigue",
  "confidence": "medium"
}
```

---

# Objetivo da Resposta

A IA deve retornar:

- tipo;
- confiança;
- payload opcional.

---

# Prompting

O prompt deve ser:

- determinístico;
- simples;
- focado em classificação.

---

# NÃO fazer

Evitar prompts:

- criativos;
- longos;
- conversacionais.

---

# Exemplo Conceitual

```text
Classifique a mensagem abaixo
e retorne JSON.
```

---

# Estratégia de Custo

A IA deve ser chamada apenas quando:

- parser falhar.

---

# Objetivo

Minimizar:
- custo;
- latência.

---

# Timeout

Toda chamada deve possuir timeout.

---

# Exemplo

```text
5 segundos
```

---

# Fallback

Se a IA falhar:

```text
NOT_INTERPRETED
```

---

# Observabilidade

Registrar:

- chamada para IA;
- sucesso;
- falha;
- tempo de resposta.

---

# NÃO registrar

- prompts completos;
- dados sensíveis.

---

# Tratamento de Erros

## Timeout

Retornar:

```text
NOT_INTERPRETED
```

---

## Erro do Provider

Retornar:

```text
NOT_INTERPRETED
```

---

## Resposta Inválida

Retornar:

```text
NOT_INTERPRETED
```

---

# Critérios de Aceite

## Integração

- provider funciona corretamente.

---

## Estrutura

- resposta padronizada.

---

## Fallback

- falhas tratadas corretamente.

---

## Arquitetura

- domínio desacoplado do Gemini.

---

# Requisitos Técnicos

## Deve existir

- AIProvider;
- GeminiProvider;
- timeout;
- resposta estruturada.

---

# Dependências

Relaciona-se com:
- SPEC-014 - Parser Local de Eventos
- SPEC-013 - Router Conversacional

---

# Considerações Arquiteturais

## Provider Pattern

A IA deve ser facilmente substituível.

---

# Exemplo Futuro

```text
Gemini
Claude
OpenAI
Local Model
```

---

## IA como fallback

A IA nunca deve ser a primeira etapa.

---

## Compatibilidade com Monorepo

O provider deve:
- permanecer isolado em `/apps/api/internal/infra/ai`;
- ser reutilizável por futuros módulos;
- evitar dependência do domínio em SDKs externos.

---

# Objetivo Real do MVP

O foco é:
- interpretar mensagens difíceis;
- reduzir esforço manual;
- complementar parser local.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- múltiplos providers;
- fallback entre providers;
- memória contextual;
- RAG;
- modelos locais;
- classificação avançada.

---

# Implementation Checklist

## Estrutura Base

- [x] Estrutura `/infra/ai` criada
- [x] Interface AIProvider criada
- [x] Contratos de interpretação criados

---

## Gemini

- [x] GeminiProvider criado
- [x] Cliente Gemini configurado
- [x] Integração com API Gemini funcionando

---

## Configuração

- [x] GEMINI_API_KEY configurada
- [x] Integração com Config implementada
- [x] Variáveis de ambiente documentadas

---

## Modelo de Resposta

- [x] InterpretationResult criado
- [x] Campo type implementado
- [x] Campo confidence implementado
- [x] Campo payload implementado

---

## Prompting

- [x] Prompt determinístico criado
- [x] Prompt focado em classificação criado
- [x] Resposta JSON padronizada implementada

---

## Integração com Parser

- [ ] Acionamento após NOT_MATCHED implementado (SPEC-016)
- [ ] Integração com Router preparada (SPEC-016)
- [ ] Integração com Registro de Eventos preparada (SPEC-016)

---

## Timeout

- [x] Timeout configurado
- [x] Cancelamento por contexto implementado
- [x] Tratamento de timeout implementado

---

## Fallback

- [x] Resultado NOT_INTERPRETED implementado
- [x] Tratamento de falha do provider implementado
- [x] Tratamento de resposta inválida implementado

---

## Observabilidade

- [x] Log de chamada para IA implementado
- [x] Log de sucesso implementado
- [x] Log de falha implementado
- [x] Log de latência implementado

---

## Segurança

- [x] API Key protegida
- [x] Prompts sensíveis protegidos
- [x] Dados pessoais minimizados

---

## Testes

- [x] Mock de AIProvider criado
- [x] Testes do GeminiProvider implementados
- [x] Testes de timeout implementados
- [x] Testes de fallback implementados

---

## Qualidade

- [x] Provider desacoplado do domínio
- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Gemini facilmente substituível