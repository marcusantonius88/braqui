# SPEC-011 - Cadastro de Pet

## Objetivo

Definir o fluxo de cadastro do pet no Braqui MVP.

O foco desta spec é:
- coletar informações essenciais do cão;
- concluir o onboarding;
- criar contexto inicial para o Braqui;
- permitir personalização dos insights e alertas.

---

# Contexto

O Braqui depende de informações básicas do pet para:

- personalizar respostas;
- interpretar eventos;
- gerar lembretes;
- emitir alertas climáticos;
- construir histórico.

Sem cadastro do pet, o sistema perde grande parte do seu valor.

---

# Escopo

## O sistema deve:

- cadastrar um pet;
- associar pet ao tutor;
- validar informações mínimas;
- concluir onboarding inicial.

---

# Fora do Escopo

Esta spec NÃO contempla:

- múltiplos pets;
- upload de fotos;
- pedigree;
- histórico veterinário completo;
- documentos;
- vacinação detalhada.

---

# Filosofia Arquitetural

O cadastro deve ser:

- simples;
- rápido;
- conversacional;
- sem formulários complexos.

---

# Estratégia Inicial

O cadastro ocorrerá diretamente na conversa.

---

# Fluxo Esperado

```text
Qual o nome do seu cão?
        ↓
Qual a raça?
        ↓
Qual a idade?
        ↓
Qual o peso?
        ↓
Em qual cidade você mora?
        ↓
Cadastro concluído
```

---

# Informações Obrigatórias

## Nome

Exemplo:

```text
Thor
```

---

## Raça

Exemplo:

```text
Buldogue Francês
```

---

## Idade

Exemplo:

```text
3 anos
```

---

## Peso

Exemplo:

```text
12 kg
```

---

## Cidade

Exemplo:

```text
João Pessoa
```

---

# Objetivo da Cidade

Permitir:

- alertas climáticos;
- contexto ambiental;
- personalização futura.

---

# Estrutura Esperada

```text
Pet
```

---

# Campos Iniciais

```text
id
user_id
name
breed
age
weight
city
created_at
updated_at
```

---

# Associação

Relacionamento:

```text
User
   ↓
 Pet
```

---

# MVP

Inicialmente:

- um pet por usuário.

---

# Regras de Negócio

## Usuário sem pet

Deve iniciar onboarding.

---

## Usuário com pet

Não deve repetir onboarding.

---

## Cadastro concluído

Permite utilização completa do sistema.

---

# Validações

## Nome

Obrigatório.

---

## Raça

Obrigatória.

---

## Idade

Obrigatória.

---

## Peso

Obrigatório.

---

## Cidade

Obrigatória.

---

# Tratamento Conversacional

O sistema deve:

- fazer uma pergunta por vez;
- manter contexto;
- recuperar estado da conversa.

---

# Exemplo

Usuário:

```text
Thor
```

Sistema:

```text
Qual a raça do Thor?
```

---

# Confirmação Final

Após cadastro:

```text
Perfeito 🐶

Cadastro do Thor concluído com sucesso.
```

---

# Persistência

Os dados devem ser armazenados no PostgreSQL.

---

# Estrutura Esperada

```text
PetRepository
```

---

# Observabilidade

Registrar:

- início do cadastro;
- conclusão do cadastro;
- falhas de persistência.

---

# NÃO registrar

- informações desnecessárias;
- conteúdo excessivo da conversa.

---

# Critérios de Aceite

## Cadastro

- pet criado corretamente.

---

## Associação

- pet associado ao usuário.

---

## Persistência

- dados armazenados corretamente.

---

## Conversação

- onboarding ocorre corretamente.

---

# Requisitos Técnicos

## Deve existir

- entidade Pet;
- PetRepository;
- fluxo conversacional;
- persistência.

---

# Dependências

Relaciona-se com:
- SPEC-010 - Autenticação e Identificação de Usuário
- SPEC-012 - Gerenciamento de Estado Conversacional
- SPEC-005 - Persistência e Repositories

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- múltiplos pets;
- anexos;
- cadastro avançado.

---

## Conversação guiada

O onboarding deve:
- reduzir atrito;
- evitar perguntas excessivas;
- ser natural.

---

## Compatibilidade com Monorepo

O cadastro deve:
- permanecer em `/apps/api`;
- utilizar os módulos existentes;
- evitar dependência direta do Telegram no domínio.

---

# Objetivo Real do MVP

O foco é:
- obter contexto do pet;
- concluir onboarding;
- habilitar funcionalidades futuras.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- múltiplos pets;
- fotos;
- peso histórico;
- dados veterinários;
- documentos;
- vacinação detalhada.

---

# Implementation Checklist

## Estrutura Base

- [x] Entidade Pet criada
- [x] PetRepository criado
- [x] Casos de uso de cadastro criados

---

## Persistência

- [x] Tabela pets criada
- [x] Migration de pets criada
- [x] Persistência de pets implementada

---

## Associação

- [x] Associação User → Pet implementada
- [x] Relacionamento persistido corretamente

---

## Fluxo Conversacional

- [x] Pergunta de nome implementada
- [x] Pergunta de raça implementada
- [x] Pergunta de idade implementada
- [x] Pergunta de peso implementada
- [x] Pergunta de cidade implementada

---

## Estado Conversacional

- [x] Avanço entre etapas implementado
- [x] Recuperação de estado implementada
- [x] Continuidade de onboarding implementada

---

## Validações

- [x] Validação de nome implementada
- [x] Validação de raça implementada
- [x] Validação de idade implementada
- [x] Validação de peso implementada
- [x] Validação de cidade implementada

---

## Conclusão de Cadastro

- [x] Confirmação final implementada
- [x] Finalização de onboarding implementada
- [x] Liberação das funcionalidades implementada

---

## Observabilidade

- [x] Log de início de cadastro implementado
- [x] Log de conclusão de cadastro implementado
- [x] Log de falhas implementado

---

## Tratamento de Erros

- [x] Falhas de persistência tratadas
- [x] Mensagens amigáveis implementadas
- [x] Recuperação de fluxo implementada

---

## Testes

- [x] Testes de cadastro implementados
- [x] Testes de validação implementados
- [x] Testes de persistência implementados
- [x] Testes de onboarding implementados

---

## Qualidade

- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Domínio desacoplado do Telegram