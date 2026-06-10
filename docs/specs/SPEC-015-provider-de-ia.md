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

- [ ] Estrutura `/infra/ai` criada
- [ ] Interface AIProvider criada
- [ ] Contratos de interpretação criados

---

## Gemini

- [ ] GeminiProvider criado
- [ ] Cliente Gemini configurado
- [ ] Integração com API Gemini funcionando

---

## Configuração

- [ ] GEMINI_API_KEY configurada
- [ ] Integração com Config implementada
- [ ] Variáveis de ambiente documentadas

---

## Modelo de Resposta

- [ ] InterpretationResult criado
- [ ] Campo type implementado
- [ ] Campo confidence implementado
- [ ] Campo payload implementado

---

## Prompting

- [ ] Prompt determinístico criado
- [ ] Prompt focado em classificação criado
- [ ] Resposta JSON padronizada implementada

---

## Integração com Parser

- [ ] Acionamento após NOT_MATCHED implementado
- [ ] Integração com Router preparada
- [ ] Integração com Registro de Eventos preparada

---

## Timeout

- [ ] Timeout configurado
- [ ] Cancelamento por contexto implementado
- [ ] Tratamento de timeout implementado

---

## Fallback

- [ ] Resultado NOT_INTERPRETED implementado
- [ ] Tratamento de falha do provider implementado
- [ ] Tratamento de resposta inválida implementado

---

## Observabilidade

- [ ] Log de chamada para IA implementado
- [ ] Log de sucesso implementado
- [ ] Log de falha implementado
- [ ] Log de latência implementado

---

## Segurança

- [ ] API Key protegida
- [ ] Prompts sensíveis protegidos
- [ ] Dados pessoais minimizados

---

## Testes

- [ ] Mock de AIProvider criado
- [ ] Testes do GeminiProvider implementados
- [ ] Testes de timeout implementados
- [ ] Testes de fallback implementados

---

## Qualidade

- [ ] Provider desacoplado do domínio
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Gemini facilmente substituível