# SPEC-010 - Autenticação e Identificação de Usuário

## Objetivo

Definir a estratégia de identificação e autenticação dos usuários do Braqui MVP.

O foco desta spec é:
- identificar tutores;
- associar pets aos tutores;
- manter contexto conversacional;
- simplificar onboarding;
- evitar complexidade desnecessária.

---

# Contexto

O Braqui será utilizado inicialmente através do Telegram.

Diferente de aplicações tradicionais:

- não haverá login;
- não haverá senha;
- não haverá cadastro manual.

A identidade do usuário será derivada do próprio Telegram.

---

# Escopo

## O sistema deve:

- identificar usuários;
- criar cadastro automaticamente;
- associar pets ao tutor;
- manter contexto persistente.

---

# Fora do Escopo

Esta spec NÃO contempla:

- login e senha;
- OAuth;
- autenticação social;
- SSO;
- MFA;
- recuperação de senha.

---

# Filosofia Arquitetural

A autenticação deve ser:

- invisível;
- simples;
- automática;
- transparente para o usuário.

---

# Estratégia Inicial

A identificação será baseada em:

```text
telegram_user_id
```

---

# Fluxo Inicial

```text
Mensagem recebida
        ↓
telegram_user_id
        ↓
Usuário existe?
        ↓
SIM → continua
NÃO → cria usuário
```

---

# Objetivo

Eliminar fricção no onboarding.

---

# Cadastro Automático

Quando um usuário conversar pela primeira vez:

- usuário será criado automaticamente;
- onboarding será iniciado;
- contexto será persistido.

---

# Estrutura Esperada

```text
User
```

---

# Campos Iniciais

```text
id
telegram_user_id
first_name
username
created_at
updated_at
```

---

# Regras de Negócio

## Um usuário

Pode possuir:

- múltiplos pets futuramente.

---

## MVP

Inicialmente:

- um único pet por usuário.

---

# Identificação

Toda mensagem recebida deve:

- localizar usuário;
OU
- criar usuário automaticamente.

---

# Persistência

Usuários devem ser armazenados no PostgreSQL.

---

# Estrutura Esperada

```text
UserRepository
```

---

# Fluxo de Onboarding

Primeira mensagem:

```text
Olá
```

---

# Sistema

```text
Olá 👋

Sou o Braqui.

Qual o nome do seu cão?
```

---

# Associação de Pet

Após onboarding:

```text
Usuário
      ↓
Pet
```

---

# Contexto Conversacional

A identificação deve permitir:

- recuperar histórico;
- recuperar pet associado;
- recuperar lembretes;
- recuperar eventos.

---

# Segurança

O sistema NÃO deve:

- armazenar senhas;
- solicitar credenciais;
- expor identificadores internos.

---

# Privacidade

Armazenar apenas:

- informações necessárias;
- contexto mínimo.

---

# Observabilidade

Registrar:

- criação de usuário;
- onboarding iniciado;
- onboarding concluído.

---

# NÃO registrar

- informações sensíveis;
- conteúdo desnecessário.

---

# Tratamento de Erros

## Usuário não encontrado

Sistema deve:

- criar automaticamente.

---

## Falha de persistência

Sistema deve:

- registrar erro;
- informar falha amigável.

---

# Critérios de Aceite

## Identificação

- usuário identificado corretamente.

---

## Cadastro

- usuário criado automaticamente.

---

## Persistência

- usuário armazenado corretamente.

---

## Contexto

- contexto recuperado corretamente.

---

# Requisitos Técnicos

## Deve existir

- entidade User;
- UserRepository;
- identificação automática;
- onboarding inicial.

---

# Dependências

Relaciona-se com:
- SPEC-005 - Persistência e Repositories
- SPEC-009 - Integração com Telegram

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- autenticação tradicional;
- login;
- senha.

---

## Experiência sem atrito

O usuário deve:
- começar a usar imediatamente;
- sem configuração adicional.

---

## Compatibilidade com Monorepo

A identificação deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por futuros canais;
- evitar dependência direta do Telegram no domínio.

---

# Objetivo Real do MVP

O foco é:
- identificar usuários;
- iniciar onboarding automaticamente;
- manter histórico consistente;
- reduzir fricção.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- múltiplos pets;
- múltiplos canais;
- contas compartilhadas;
- dashboard web;
- autenticação complementar;
- gerenciamento de perfil.

---

# Implementation Checklist

## Estrutura Base

- [x] Entidade User criada
- [x] UserRepository criado
- [x] Casos de uso de identificação criados

---

## Persistência

- [x] Tabela users criada
- [x] Migration de users criada
- [x] Persistência de usuários implementada

---

## Identificação

- [x] Extração de telegram_user_id implementada
- [x] Busca de usuário implementada
- [x] Criação automática de usuário implementada

---

## Cadastro Automático

- [x] Fluxo de primeiro acesso implementado
- [x] Criação automática de usuário validada
- [x] Usuário persistido corretamente

---

## Onboarding

- [x] Início do onboarding implementado
- [x] Primeira mensagem de boas-vindas implementada
- [ ] Fluxo inicial de cadastro concluído (pendente — depende de SPEC-011 cadastro de pet)

---

## Associação de Contexto

- [x] Recuperação de usuário implementada
- [ ] Recuperação de pet associado implementada (pendente — depende de SPEC-011)
- [ ] Recuperação de histórico implementada (pendente — depende de specs de funcionalidade)

---

## Regras de Negócio

- [ ] Regra de um pet por usuário implementada (pendente — depende de SPEC-011)
- [x] Preparação para múltiplos pets futura documentada

---

## Observabilidade

- [x] Log de criação de usuário implementado
- [x] Log de onboarding iniciado implementado
- [ ] Log de onboarding concluído implementado (pendente — depende de SPEC-011)

---

## Segurança

- [x] Nenhuma senha armazenada
- [x] Nenhuma credencial solicitada
- [x] Identificadores internos protegidos

---

## Tratamento de Erros

- [x] Falhas de persistência tratadas
- [x] Mensagens amigáveis implementadas
- [x] Logs de erro implementados

---

## Testes

- [x] Testes de identificação implementados
- [x] Testes de criação automática implementados
- [x] Testes de onboarding implementados
- [x] Testes de persistência implementados

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Domínio desacoplado do Telegram