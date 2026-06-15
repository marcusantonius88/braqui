# SPEC-021 - Alertas Climáticos

## Objetivo

Definir a funcionalidade de monitoramento climático e geração de alertas para cães braquicefálicos.

O foco desta spec é:
- proteger o pet em condições climáticas adversas;
- gerar alertas proativos;
- aumentar o valor percebido do Braqui;
- utilizar o contexto geográfico do tutor.

---

# Contexto

Cães braquicefálicos possuem maior sensibilidade a:

- calor;
- umidade;
- sensação térmica elevada;
- baixa ventilação.

Essas condições podem aumentar o risco de:

- ofegância excessiva;
- exaustão térmica;
- desconforto respiratório;
- agravamento de sintomas.

---

# Escopo

## O sistema deve:

- consultar condições climáticas;
- identificar situações de risco;
- enviar alertas ao tutor;
- funcionar automaticamente.

---

# Fora do Escopo

Esta spec NÃO contempla:

- previsões avançadas;
- radar meteorológico;
- qualidade do ar;
- alertas por GPS;
- recomendações veterinárias.

---

# Filosofia Arquitetural

Os alertas devem ser:

- simples;
- úteis;
- não alarmistas;
- contextuais.

---

# Estratégia Inicial

Fluxo:

```text
Cidade
      ↓
Provider Climático
      ↓
Regras de Risco
      ↓
Alerta
      ↓
Telegram
```

---

# Fonte da Localização

Utilizar:

```text
Pet.city
```

---

# Provider Climático

Inicialmente:

```text
OpenWeather
```

---

# Estrutura Esperada

```text
/apps
  /api
    /internal
      /climate
```

---

# Informações Climáticas

Inicialmente:

- temperatura;
- umidade.

---

# Regras Iniciais

## Calor Elevado

Temperatura:

```text
>= 30°C
```

---

## Calor Crítico

Temperatura:

```text
>= 35°C
```

---

## Umidade Elevada

Preparação futura.

---

# Exemplo de Alerta

```text
🐶 Atenção

Hoje está muito quente em João Pessoa.

Cães braquicefálicos podem sofrer mais nesses dias.
```

---

# Frequência

Inicialmente:

```text
1 vez por dia
```

---

# Scheduler

A verificação deve ser executada pelo scheduler.

---

# Integração

Relacionamento:

```text
Scheduler
      ↓
Climate Provider
      ↓
Telegram
```

---

# Estratégia Anti-Spam

Evitar:

- múltiplos alertas iguais;
- repetição excessiva.

---

# MVP

Inicialmente:

```text
1 alerta climático por dia
```

---

# Regras de Negócio

## Sem cidade cadastrada

Não enviar alerta.

---

## Clima normal

Não enviar alerta.

---

## Clima de risco

Enviar alerta.

---

# Estrutura Conceitual

```go
type ClimateProvider interface {
    GetCurrentWeather(
        city string,
    )
}
```

---

# Observabilidade

Registrar:

- consulta climática;
- alerta enviado;
- falha de integração.

---

# NÃO registrar

- dados excessivos;
- payloads completos.

---

# Tratamento de Erros

## Cidade inválida

Registrar erro.

---

## Provider indisponível

Registrar erro e tentar novamente futuramente.

---

## Falha de envio

Registrar erro.

---

# Critérios de Aceite

## Consulta

- clima consultado corretamente.

---

## Regras

- situações de risco identificadas.

---

## Telegram

- alertas enviados corretamente.

---

## Arquitetura

- provider desacoplado do domínio.

---

# Requisitos Técnicos

## Deve existir

- ClimateProvider;
- integração OpenWeather;
- regras climáticas;
- integração com scheduler.

---

# Dependências

Relaciona-se com:
- SPEC-020 - Localização do Usuário
- SPEC-018 - Scheduler e Tarefas Agendadas
- SPEC-009 - Integração com Telegram

---

# Considerações Arquiteturais

## Simplicidade primeiro

O MVP NÃO precisa:
- previsão de vários dias;
- machine learning;
- análise avançada.

---

## Valor imediato

O alerta climático é uma das funcionalidades mais alinhadas ao nicho braquicefálico.

---

## Compatibilidade com Monorepo

O módulo climático deve:
- permanecer em `/apps/api`;
- ser reutilizável por insights futuros;
- evitar dependência direta do Telegram.

---

# Objetivo Real do MVP

O foco é:
- proteger o pet;
- gerar valor diário;
- diferenciar o Braqui.

---

# Possíveis Evoluções Futuras

Fora do MVP:
- qualidade do ar;
- previsão climática;
- índice de calor;
- umidade avançada;
- recomendações contextualizadas.

---

# Implementation Checklist

## Estrutura Base

- [ ] Estrutura `/climate` criada
- [ ] Interface ClimateProvider criada
- [ ] Casos de uso climáticos criados

---

## OpenWeather

- [ ] Cliente OpenWeather criado
- [ ] Integração com API implementada
- [ ] OPENWEATHER_API_KEY configurada

---

## Dados Climáticos

- [ ] Consulta de temperatura implementada
- [ ] Consulta de umidade implementada
- [ ] Modelo WeatherData criado

---

## Regras de Risco

- [ ] Regra de calor elevado (>= 30°C) implementada
- [ ] Regra de calor crítico (>= 35°C) implementada
- [ ] Estrutura para novas regras preparada

---

## Localização

- [ ] Integração com Pet.city implementada
- [ ] Tratamento de ausência de cidade implementado
- [ ] Tratamento de cidade inválida implementado

---

## Alertas

- [ ] Template de alerta criado
- [ ] Geração de alerta implementada
- [ ] Envio via Telegram implementado

---

## Scheduler

- [ ] Integração com ClimateAlertJob implementada
- [ ] Execução diária implementada

---

## Anti-Spam

- [ ] Controle de alerta diário implementado
- [ ] Evitar alertas duplicados implementado

---

## Observabilidade

- [ ] Log de consulta climática implementado
- [ ] Log de alerta enviado implementado
- [ ] Log de falha implementado

---

## Tratamento de Erros

- [ ] Falha do provider tratada
- [ ] Falha de envio tratada
- [ ] Recuperação operacional implementada

---

## Testes

- [ ] Mock de ClimateProvider criado
- [ ] Testes de regras climáticas implementados
- [ ] Testes de geração de alertas implementados
- [ ] Testes de integração implementados

---

## Qualidade

- [ ] ClimateProvider desacoplado do domínio
- [ ] Estrutura compatível com monorepo
- [ ] Estrutura compatível com SDD
- [ ] Estrutura compatível com IA-assisted development
- [ ] Base preparada para novas regras climáticas