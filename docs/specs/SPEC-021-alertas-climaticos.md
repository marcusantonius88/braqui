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

- [x] Estrutura `/climate` criada (`internal/climate/`)
- [x] Interface ClimateProvider criada (`Provider`)
- [x] Casos de uso climáticos criados (`Service.CheckAndAlert`)

---

## OpenWeather

- [x] Cliente OpenWeather criado (`OpenWeatherProvider`)
- [x] Integração com API implementada (REST via net/http, units=metric)
- [x] OPENWEATHER_API_KEY configurada (config + .env.example + docker-compose)

---

## Dados Climáticos

- [x] Consulta de temperatura implementada
- [x] Consulta de umidade implementada
- [x] Modelo WeatherData criado (Temperature, Humidity, City)

---

## Regras de Risco

- [x] Regra de calor elevado (>= 30°C) implementada (RiskHigh)
- [x] Regra de calor crítico (>= 35°C) implementada (RiskCritical)
- [x] Estrutura para novas regras preparada (RiskLevel enum + EvaluateRisk)

---

## Localização

- [x] Integração com Pet.city implementada (via PetRepository.FindAllWithLocation)
- [x] Tratamento de ausência de cidade implementado (pula pets sem city)
- [x] Tratamento de cidade inválida implementado (erro OpenWeather → log)

---

## Alertas

- [x] Template de alerta criado (FormatAlert)
- [x] Geração de alerta implementada
- [x] Envio via Telegram implementado

---

## Scheduler

- [x] Integração com ClimateAlertJob implementada
- [x] Execução diária implementada (frequência Daily)

---

## Anti-Spam

- [x] Controle de alerta diário implementado (alertedToday map[string]string)
- [x] Evitar alertas duplicados implementado (mesmo pet ignorado se já alertado hoje)

---

## Observabilidade

- [x] Log de consulta climática implementado (erro/sucesso)
- [x] Log de alerta enviado implementado
- [x] Log de falha implementado

---

## Tratamento de Erros

- [x] Falha do provider tratada (log + continue)
- [x] Falha de envio tratada (log + continue)
- [x] Recuperação operacional implementada (panic recovery do scheduler)

---

## Testes

- [x] Mock de ClimateProvider criado (mockWeatherProvider)
- [x] Testes de regras climáticas implementados (RiskLevel thresholds)
- [x] Testes de geração de alertas implementados (FormatAssert)
- [x] Testes de integração implementados (end-to-end com Docker: job roda, logs corretos, app não quebra)

---

## Qualidade

- [x] ClimateProvider desacoplado do domínio (interface local em climate/)
- [x] Estrutura compatível com monorepo
- [x] Estrutura compatível com SDD
- [x] Estrutura compatível com IA-assisted development
- [x] Base preparada para novas regras climáticas (RiskLevel enum + EvaluateRisk switch)