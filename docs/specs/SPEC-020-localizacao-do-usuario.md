# SPEC-020 - Localização do Usuário

## Objetivo

Definir a estratégia de captura e utilização da localização do tutor para personalização de funcionalidades do Braqui.

O foco desta spec é:
- permitir alertas climáticos;
- contextualizar informações ambientais;
- melhorar precisão dos insights;
- suportar futuras funcionalidades geográficas.

---

# Contexto

Cães braquicefálicos possuem forte sensibilidade às condições climáticas.

Exemplos:

- calor excessivo;
- baixa qualidade do ar;
- alta umidade;
- sensação térmica elevada.

Para que o Braqui gere alertas úteis, ele precisa conhecer a localização do tutor.

---

# Escopo

## O sistema deve:

- armazenar localização do tutor;
- associar localização ao pet;
- permitir consulta de clima;
- suportar personalização futura.

---

# Fora do Escopo

Esta spec NÃO contempla:

- GPS em tempo real;
- rastreamento contínuo;
- geofencing;
- mapas;
- compartilhamento de localização em tempo real.

---

# Filosofia Arquitetural

A localização deve ser:

- simples;
- opcionalmente refinável;
- suficiente para alertas climáticos.

---

# Estratégia Inicial

A localização será capturada durante o onboarding.

---

# Exemplo

```text
Em qual cidade você mora?
```

---

# MVP

Inicialmente:

```text
cidade
```

será suficiente.

---

# Exemplo

```text
João Pessoa
```

---

# Objetivo

Permitir integração com provedores de clima.

---

# Estrutura Inicial

A localização ficará armazenada no cadastro do pet.

---

# Campos Iniciais

```text
city
```

---

# Evolução Futura

Possibilidade de adicionar:

```text
state
country
latitude
longitude
```

---

# Fluxo Esperado

```text
Onboarding
      ↓
Cidade informada
      ↓
Persistência
      ↓
Consulta climática futura
```

---

# Persistência

A cidade deve ser armazenada junto ao pet.

---

# Estrutura Relacionada

```text
Pet
```

---

# Atualização

O tutor poderá alterar a localização futuramente.

---

# Exemplo

```text
Mude minha cidade para Recife
```

---

# MVP

Inicialmente:

- atualização manual simples.

---

# Integração com Clima

A localização será utilizada por:

```text
Climate Provider
```

---

# Exemplo

```text
João Pessoa
      ↓
OpenWeather
      ↓
Dados climáticos
```

---

# Validações

## Cidade

Obrigatória durante onboarding.

---

# Tratamento Inicial

Aceitar texto livre.

---

# Normalização

Inicialmente:

- opcional.

---

# Observabilidade

Registrar:

- localização cadastrada;
- localização atualizada;
- falhas de consulta.

---

# NÃO registrar

- localização precisa;
- dados excessivos.

---

# Privacidade

Armazenar apenas:

- cidade.

---

# Objetivo

Minimizar coleta de dados.

---

# Tratamento de Erros

## Cidade inválida

Responder:

```text
Não consegui localizar essa cidade.
```

---

## Falha de consulta climática

Registrar erro e continuar funcionamento.

---

# Critérios de Aceite

## Cadastro

- cidade armazenada corretamente.

---

## Consulta

- localização recuperada corretamente.

---

## Integração

- localização disponível para módulos climáticos.

---

## Arquitetura

- localização desacoplada do provider climático.

---

# Requisitos Técnicos

## Deve existir

- armazenamento de cidade;
- recuperação da cidade;
- atualização da cidade;
- integração futura com clima.

---

# Dependências

Relaciona-se com:
- SPEC-011 - Cadastro de Pet
- SPEC-021 - Alertas Climáticos

---

# Considerações Arquiteturais

## Menor dado possível

O MVP deve armazenar apenas o necessário.

---

## Privacidade

Evitar:
- coordenadas;
- rastreamento;
- monitoramento contínuo.

---

## Compatibilidade com Monorepo

A funcionalidade deve:
- permanecer dentro de `/apps/api`;
- ser reutilizável por módulos climáticos;
- evitar acoplamento ao provider externo.

---

# Objetivo Real do MVP

O foco é:
- habilitar alertas climáticos;
- personalizar experiência;
- minimizar coleta de dados.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- latitude;
- longitude;
- múltiplas localizações;
- GPS;
- qualidade do ar;
- geofencing.

---

# Implementation Checklist

## Estrutura Base

- [ ] Campo city definido no domínio
- [ ] Estrutura de localização criada
- [ ] Casos de uso de localização criados

---

## Persistência

- [ ] Campo city adicionado à entidade Pet
- [ ] Migration atualizada
- [ ] Persistência implementada

---

## Onboarding

- [ ] Pergunta de cidade implementada
- [ ] Captura da cidade implementada
- [ ] Integração com cadastro do pet implementada

---

## Recuperação

- [ ] Recuperação da cidade implementada
- [ ] Disponibilização para outros módulos implementada

---

## Atualização

- [ ] Atualização manual da cidade implementada
- [ ] Fluxo conversacional de alteração implementado

---

## Validação

- [ ] Cidade obrigatória no onboarding
- [ ] Validação básica implementada
- [ ] Tratamento de cidade inválida implementado

---

## Integração Climática

- [ ] Estrutura preparada para Climate Provider
- [ ] Cidade disponível para consultas climáticas
- [ ] Integração futura documentada

---

## Privacidade

- [ ] Apenas cidade armazenada
- [ ] Nenhuma coordenada armazenada
- [ ] Nenhum rastreamento implementado

---

## Observabilidade

- [ ] Log de cadastro de localização implementado
- [ ] Log de atualização implementado
- [ ] Log de falhas implementado

---

## Tratamento de Erros

- [ ] Cidade inválida tratada
- [ ] Falha de atualização tratada
- [ ] Falha de consulta tratada

---

## Testes

- [ ] Testes de cadastro implementados
- [ ] Testes de recuperação implementados
- [ ] Testes de atualização implementados
- [ ] Testes de validação implementados

---

## Qualidade

- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Localização desacoplada do provider climático