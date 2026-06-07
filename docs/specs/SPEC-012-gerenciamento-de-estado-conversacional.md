# SPEC-012 - Gerenciamento de Estado Conversacional

## Objetivo

Definir a estratégia de gerenciamento de estado conversacional do Braqui MVP.

O foco desta spec é:
- manter contexto entre mensagens;
- permitir fluxos guiados;
- suportar onboarding;
- suportar interações multi-etapas;
- evitar perda de contexto.

---

# Contexto

O Braqui é uma aplicação conversacional.

Diferente de aplicações tradicionais:
- não existe formulário completo;
- não existe tela de cadastro;
- não existe wizard visual.

O estado da conversa deve ser persistido para permitir continuidade.

---

# Escopo

## O sistema deve:

- armazenar estado conversacional;
- recuperar contexto;
- continuar fluxos interrompidos;
- permitir conversas multi-etapas.

---

# Fora do Escopo

Esta spec NÃO contempla:

- memória de longo prazo baseada em IA;
- histórico completo de chat;
- contexto semântico avançado;
- LLM memory.

---

# Filosofia Arquitetural

O estado conversacional deve ser:

- explícito;
- persistente;
- previsível;
- simples.

---

# Estratégia Inicial

Cada usuário possuirá:

```text
ConversationState
```

---

# Objetivo

Permitir:

- onboarding;
- cadastro de pet;
- registro guiado;
- futuras interações complexas.

---

# Estrutura Esperada

```text
ConversationState
```

---

# Campos Iniciais

```text
id
user_id
current_flow
current_step
payload
created_at
updated_at
```

---

# Significado dos Campos

## current_flow

Representa o fluxo atual.

Exemplos:

```text
onboarding
register_pet
register_event
```

---

## current_step

Representa etapa atual.

Exemplos:

```text
ask_name
ask_breed
ask_age
```

---

## payload

Armazena respostas intermediárias.

Exemplo:

```json
{
  "name": "Thor"
}
```

---

# Persistência

O estado deve ser persistido.

---

# Estrutura Esperada

```text
ConversationStateRepository
```

---

# Fluxo Esperado

```text
Pergunta
      ↓
Resposta
      ↓
Atualiza estado
      ↓
Próxima pergunta
```

---

# Exemplo

Sistema:

```text
Qual o nome do seu cão?
```

---

# Estado

```json
{
  "flow": "register_pet",
  "step": "ask_name"
}
```

---

# Usuário

```text
Thor
```

---

# Novo Estado

```json
{
  "flow": "register_pet",
  "step": "ask_breed",
  "payload": {
    "name": "Thor"
  }
}
```

---

# Recuperação

Caso o usuário envie nova mensagem:

- sistema recupera estado;
- continua fluxo corretamente.

---

# Interrupções

O sistema deve suportar:

- reinício da aplicação;
- mensagens atrasadas;
- retomada de fluxo.

---

# Finalização

Ao concluir fluxo:

- estado deve ser removido;
OU
- marcado como concluído.

---

# Estratégia Inicial

Preferencialmente:

```text
estado ativo único por usuário
```

---

# Regras de Negócio

## Usuário sem estado

Fluxo normal.

---

## Usuário com estado

Continuar fluxo existente.

---

## Fluxo concluído

Encerrar estado.

---

# Observabilidade

Registrar:

- fluxo iniciado;
- etapa alterada;
- fluxo concluído.

---

# NÃO registrar

- payloads sensíveis;
- conteúdo excessivo.

---

# Tratamento de Erros

## Estado inexistente

Sistema deve:

- iniciar fluxo adequado.

---

## Estado inválido

Sistema deve:

- registrar erro;
- reiniciar fluxo quando necessário.

---

# Critérios de Aceite

## Persistência

- estado salvo corretamente.

---

## Recuperação

- estado recuperado corretamente.

---

## Continuidade

- fluxo continua corretamente.

---

## Finalização

- estado encerrado corretamente.

---

# Requisitos Técnicos

## Deve existir

- entidade ConversationState;
- repository;
- persistência;
- recuperação de contexto.

---

# Dependências

Relaciona-se com:
- SPEC-005 - Persistência e Repositories
- SPEC-010 - Autenticação e Identificação de Usuário
- SPEC-011 - Cadastro de Pet

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- memória baseada em IA;
- contexto infinito;
- histórico completo.

---

## Estado explícito

Toda continuidade deve depender:
- de estado persistido;
- de regras claras.

---

## Compatibilidade com Monorepo

O gerenciamento de estado deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por futuros fluxos;
- evitar dependência direta do Telegram.

---

# Objetivo Real do MVP

O foco é:
- permitir onboarding;
- permitir fluxos guiados;
- manter contexto consistente;
- melhorar experiência conversacional.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- múltiplos fluxos simultâneos;
- contexto avançado;
- memória semântica;
- IA contextual;
- histórico conversacional completo.

---

# Implementation Checklist

## Estrutura Base

- [ ] Entidade ConversationState criada
- [ ] ConversationStateRepository criado
- [ ] Casos de uso de gerenciamento de estado criados

---

## Persistência

- [ ] Tabela conversation_states criada
- [ ] Migration criada
- [ ] Persistência implementada

---

## Estado Conversacional

- [ ] Campo current_flow implementado
- [ ] Campo current_step implementado
- [ ] Campo payload implementado

---

## Recuperação de Estado

- [ ] Recuperação por usuário implementada
- [ ] Continuidade de fluxo implementada
- [ ] Retomada após reinício implementada

---

## Fluxos

- [ ] Fluxo onboarding suportado
- [ ] Fluxo register_pet suportado
- [ ] Fluxo register_event preparado

---

## Atualização de Estado

- [ ] Atualização de etapas implementada
- [ ] Atualização de payload implementada
- [ ] Transição entre etapas implementada

---

## Finalização

- [ ] Encerramento de fluxo implementado
- [ ] Limpeza de estado implementada
- [ ] Marcação de conclusão implementada

---

## Regras de Negócio

- [ ] Usuário sem estado tratado
- [ ] Usuário com estado tratado
- [ ] Estado inválido tratado

---

## Observabilidade

- [ ] Log de fluxo iniciado implementado
- [ ] Log de mudança de etapa implementado
- [ ] Log de fluxo concluído implementado

---

## Tratamento de Erros

- [ ] Estado inexistente tratado
- [ ] Estado inválido tratado
- [ ] Recuperação de fluxo implementada

---

## Testes

- [ ] Testes de persistência implementados
- [ ] Testes de recuperação implementados
- [ ] Testes de continuidade implementados
- [ ] Testes de finalização implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Estado desacoplado do Telegram