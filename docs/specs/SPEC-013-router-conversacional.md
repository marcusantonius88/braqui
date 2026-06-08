# SPEC-013 - Router Conversacional

## Objetivo

Definir a estratégia de roteamento das mensagens recebidas pelo Braqui.

O foco desta spec é:
- direcionar mensagens para o fluxo correto;
- integrar estado conversacional;
- desacoplar regras de roteamento;
- preparar a evolução futura da aplicação.

---

# Contexto

Após receber uma mensagem do Telegram, o Braqui precisa decidir:

- o usuário está em onboarding?
- o usuário está cadastrando um pet?
- o usuário está registrando um evento?
- o usuário está solicitando ajuda?
- a mensagem deve ser enviada para IA?
- a mensagem deve ser interpretada localmente?

Essa decisão será responsabilidade do Router Conversacional.

---

# Escopo

## O sistema deve:

- receber mensagens;
- identificar contexto atual;
- direcionar para fluxo correto;
- permitir evolução incremental.

---

# Fora do Escopo

Esta spec NÃO contempla:

- interpretação semântica avançada;
- processamento por IA;
- regras específicas de negócio;
- parser de eventos.

---

# Filosofia Arquitetural

O router deve ser:

- simples;
- explícito;
- previsível;
- desacoplado.

---

# Responsabilidade Principal

Receber:

```text
Mensagem
```

e decidir:

```text
Quem deve processar?
```

---

# Fluxo Esperado

```text
Mensagem
      ↓
Router
      ↓
Fluxo adequado
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /router
```

---

# Estratégia Inicial

A decisão deve considerar:

1. estado conversacional
2. comandos
3. fluxo padrão

---

# Ordem de Prioridade

## 1. Estado Conversacional

Se existir:

```text
ConversationState
```

o fluxo ativo possui prioridade.

---

# Exemplo

```text
Pergunta:
Qual a raça do seu cão?

Usuário:
Buldogue Francês
```

A resposta deve continuar onboarding.

---

## 2. Comandos

Exemplos:

```text
/start
/help
```

---

## 3. Fluxo Padrão

Mensagens sem contexto específico.

---

# Interface Conceitual

```go
type ConversationRouter interface {
    Route(
        ctx context.Context,
        message Message,
    ) error
}
```

---

# Fluxos Iniciais

## Onboarding

```text
register_pet
```

---

## Comandos

```text
/start
/help
```

---

## Registro de Eventos

Preparação para:

```text
Thor vomitou
```

---

# Estratégia de Crescimento

Novos fluxos devem ser adicionados sem alterar excessivamente o router.

---

# Exemplo Futuro

```text
register_event
reminder
summary
```

---

# Resultado do Router

O router deve encaminhar para:

- handler;
- fluxo;
- caso de uso apropriado.

---

# NÃO fazer

O router NÃO deve:

- executar regras de negócio;
- acessar banco diretamente;
- chamar APIs externas diretamente.

---

# Observabilidade

Registrar:

- rota selecionada;
- fluxo ativo;
- erros de roteamento.

---

# Exemplo

```text
route=register_pet
```

---

# Tratamento de Erros

## Fluxo desconhecido

Sistema deve:

- registrar erro;
- responder amigavelmente.

---

## Estado inconsistente

Sistema deve:

- registrar erro;
- reiniciar fluxo quando necessário.

---

# Critérios de Aceite

## Roteamento

- mensagens encaminhadas corretamente.

---

## Estado

- estado conversacional respeitado.

---

## Comandos

- comandos roteados corretamente.

---

## Arquitetura

- router desacoplado das regras de negócio.

---

# Requisitos Técnicos

## Deve existir

- router;
- contratos;
- roteamento por estado;
- roteamento por comandos.

---

# Dependências

Relaciona-se com:
- SPEC-009 - Integração com Telegram
- SPEC-012 - Gerenciamento de Estado Conversacional

---

# Considerações Arquiteturais

## Router é orquestrador

Sua responsabilidade é:
- decidir;
- encaminhar.

Não executar regras.

---

## Crescimento sustentável

Novos fluxos devem ser adicionados:
- com baixo acoplamento;
- sem reescrever o router.

---

## Compatibilidade com Monorepo

O router deve:
- permanecer em `/apps/api`;
- servir como ponto central de orquestração;
- ser independente do canal Telegram.

---

# Objetivo Real do MVP

O foco é:
- organizar fluxo conversacional;
- evitar lógica espalhada;
- preparar crescimento futuro.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- roteamento por intenção;
- roteamento híbrido com IA;
- múltiplos canais;
- fluxo contextual avançado;
- engine de conversação.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/router` criada
- [ ] Interface ConversationRouter criada
- [ ] Implementação inicial do router criada

---

## Fluxo Principal

- [ ] Recebimento de mensagens integrado
- [ ] Encaminhamento para router implementado
- [ ] Pipeline de roteamento criado

---

## Estado Conversacional

- [ ] Integração com ConversationState implementada
- [ ] Prioridade para fluxo ativo implementada
- [ ] Continuidade de onboarding implementada

---

## Comandos

- [ ] Roteamento de `/start` implementado
- [ ] Roteamento de `/help` implementado
- [ ] Estrutura para novos comandos criada

---

## Fluxo Padrão

- [ ] Tratamento de mensagens sem contexto implementado
- [ ] Fallback padrão implementado

---

## Registro de Eventos

- [ ] Estrutura para register_event preparada
- [ ] Integração futura documentada

---

## Tratamento de Erros

- [ ] Fluxo desconhecido tratado
- [ ] Estado inconsistente tratado
- [ ] Mensagens amigáveis implementadas

---

## Observabilidade

- [ ] Log de rota selecionada implementado
- [ ] Log de fluxo ativo implementado
- [ ] Log de erro de roteamento implementado

---

## Arquitetura

- [ ] Router desacoplado do domínio
- [ ] Router desacoplado da persistência
- [ ] Router desacoplado de APIs externas

---

## Extensibilidade

- [ ] Estrutura preparada para novos fluxos
- [ ] Estrutura preparada para múltiplos canais
- [ ] Estrutura preparada para IA futura

---

## Testes

- [ ] Testes de roteamento implementados
- [ ] Testes de prioridade de estado implementados
- [ ] Testes de comandos implementados
- [ ] Testes de fallback implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Crescimento incremental validado