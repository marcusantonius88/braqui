# SPEC-009 - Integração com Telegram

## Objetivo

Definir a integração do Braqui com o Telegram, que será o canal principal de interação do MVP.

O foco desta spec é:
- comunicação com usuários;
- recebimento de mensagens;
- envio de respostas;
- integração conversacional;
- identificação do tutor.

---

# Contexto

O Telegram será a interface principal do Braqui.

O usuário irá:
- cadastrar seu pet;
- registrar eventos;
- receber lembretes;
- receber alertas;
- consultar informações;

através de conversas com o bot.

---

# Escopo

## O sistema deve:

- receber mensagens;
- enviar mensagens;
- identificar usuários;
- processar comandos;
- suportar webhook.

---

# Fora do Escopo

Esta spec NÃO contempla:

- WhatsApp;
- Instagram;
- Facebook Messenger;
- Discord;
- múltiplos canais simultaneamente.

---

# Filosofia Arquitetural

Telegram deve ser:
- apenas um canal;
- desacoplado do domínio;
- facilmente substituível.

---

# Estratégia Inicial

Canal único:

```text
Telegram Bot API
```

---

# Fluxo Básico

```text
Usuário
    ↓
Telegram
    ↓
Webhook
    ↓
Braqui
    ↓
Resposta
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /interfaces
        /telegram

      /infra
        /telegram
```

---

# Responsabilidades

## Interfaces

Responsável por:
- receber webhook;
- validar payload;
- encaminhar mensagens.

---

## Infra

Responsável por:
- integração Telegram;
- envio de mensagens;
- chamadas HTTP.

---

# Webhook

O sistema deve receber:

```http
POST /telegram/webhook
```

---

# Processamento Inicial

Inicialmente:

- mensagens de texto;
- comandos básicos.

---

# Comandos Iniciais

```text
/start
/help
```

---

# Identificação do Usuário

O Telegram ID será utilizado como identificador inicial.

---

# Exemplo

```text
telegram_user_id
```

---

# Mensagens

Inicialmente:
- texto simples.

---

# Fora do MVP

- imagens;
- vídeos;
- áudio;
- documentos.

---

# Gateway Telegram

Deve existir um contrato explícito.

---

# Exemplo Conceitual

```go
type TelegramGateway interface {
    SendMessage(
        ctx context.Context,
        chatID string,
        text string,
    ) error
}
```

---

# Estratégia de Resposta

As respostas devem ser:

- curtas;
- claras;
- amigáveis;
- contextuais.

---

# Exemplo

```text
Thor foi cadastrado com sucesso 🐶
```

---

# Tratamento de Erros

Falhas do Telegram devem:

- registrar logs;
- permitir reprocessamento futuro;
- não derrubar a aplicação.

---

# Observabilidade

Registrar:

- webhook recebido;
- mensagem processada;
- erro de integração.

---

# Segurança

Validar:

- origem do webhook;
- payload recebido.

---

# NÃO fazer

Não registrar:

- informações sensíveis;
- payload completo do usuário.

---

# Critérios de Aceite

## Recebimento

- mensagens chegam corretamente.

---

## Envio

- respostas são enviadas corretamente.

---

## Identificação

- usuário identificado corretamente.

---

## Arquitetura

- Telegram desacoplado do domínio.

---

# Requisitos Técnicos

## Deve existir

- webhook;
- gateway Telegram;
- handler Telegram;
- integração de envio.

---

# Dependências

Relaciona-se com:
- SPEC-002 - Configuração e Gerenciamento de Ambiente
- SPEC-003 - Deploy e Infraestrutura Inicial
- SPEC-007 - Observabilidade Básica

---

# Considerações Arquiteturais

## Canal desacoplado

O domínio NÃO deve conhecer Telegram.

---

## Evolução futura

No futuro outros canais poderão existir.

---

## Compatibilidade com Monorepo

A integração deve:
- permanecer isolada em `/apps/api`;
- evitar dependências de outros apps;
- ser facilmente reutilizável.

---

# Objetivo Real do MVP

O foco é:
- permitir interação conversacional;
- reduzir atrito do usuário;
- acelerar validação do produto.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- WhatsApp;
- Discord;
- múltiplos canais;
- mídia;
- comandos avançados;
- menus interativos.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/interfaces/telegram` criada
- [ ] Estrutura `/infra/telegram` criada
- [ ] Organização arquitetural definida

---

## Webhook

- [ ] Endpoint `POST /telegram/webhook` criado
- [ ] Handler Telegram implementado
- [ ] Recebimento de mensagens funcionando
- [ ] Validação básica do payload implementada

---

## Gateway

- [ ] Interface TelegramGateway criada
- [ ] Implementação TelegramGateway criada
- [ ] Cliente HTTP Telegram configurado

---

## Envio de Mensagens

- [ ] Envio de mensagens implementado
- [ ] Tratamento de erros de envio implementado
- [ ] Timeout configurado

---

## Recebimento de Mensagens

- [ ] Processamento de mensagens de texto implementado
- [ ] Extração de chat_id implementada
- [ ] Extração de telegram_user_id implementada

---

## Comandos

- [ ] Comando `/start` implementado
- [ ] Comando `/help` implementado

---

## Identificação

- [ ] Telegram ID utilizado como identificador
- [ ] Associação de usuário implementada
- [ ] Fluxo de identificação validado

---

## Observabilidade

- [ ] Log de webhook recebido implementado
- [ ] Log de mensagem processada implementado
- [ ] Log de falhas implementado

---

## Segurança

- [ ] Validação de origem implementada
- [ ] Payload sanitizado
- [ ] Dados sensíveis protegidos

---

## Tratamento de Erros

- [ ] Falhas de integração tratadas
- [ ] Erros registrados corretamente
- [ ] Aplicação resiliente a falhas do Telegram

---

## Testes

- [ ] Mock de TelegramGateway criado
- [ ] Testes do handler implementados
- [ ] Testes do gateway implementados
- [ ] Testes de recebimento implementados
- [ ] Testes de envio implementados

---

## Qualidade

- [ ] Telegram desacoplado do domínio
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development