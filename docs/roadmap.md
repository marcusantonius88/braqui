# Roadmap - Braqui

## Objetivo

Este documento descreve a evolução planejada do Braqui.

O roadmap deve servir como direcionamento estratégico e NÃO como compromisso rígido de entrega.

As prioridades poderão mudar conforme:
- feedback dos usuários;
- retenção;
- comportamento de uso;
- validação do MVP;
- limitações técnicas;
- aprendizados do domínio.

---

# Filosofia do Roadmap

O Braqui seguirá evolução incremental baseada em:

- simplicidade;
- validação contínua;
- baixo custo operacional;
- feedback real de usuários;
- desenvolvimento orientado a domínio;
- entregas pequenas e iterativas.

O foco inicial NÃO é escala.
O foco inicial é:
- retenção;
- utilidade;
- frequência de uso;
- valor percebido.

---

# Fase 0 — Fundação

Objetivo:
Criar a base técnica e arquitetural do projeto.

## Entregas

- estrutura inicial do projeto;
- documentação inicial;
- arquitetura base;
- setup do Telegram Bot;
- setup PostgreSQL;
- setup deploy;
- pipeline básica;
- estrutura modular;
- providers iniciais.

## Resultado esperado

Base pronta para desenvolvimento incremental via specs.

---

# Fase 1 — MVP Conversacional

Objetivo:
Validar uso contínuo através de interação conversacional simples.

## Entregas

### Cadastro de pet

- nome;
- raça;
- idade;
- peso.

---

### Registro de eventos

Permitir registrar:
- vômito;
- coceira;
- ofegância;
- medicação;
- passeio;
- alimentação.

---

### Timeline de eventos

- histórico simples;
- ordenação cronológica;
- persistência básica.

---

### Alertas climáticos

- temperatura atual;
- risco climático;
- alertas simples para calor excessivo.

---

### Lembretes

- medicação;
- antipulgas;
- vacinas;
- banho.

---

### Parsing híbrido

- regras simples;
- IA apenas para mensagens ambíguas.

---

### Respostas contextuais

- mensagens curtas;
- tom acolhedor;
- confirmações rápidas.

---

## Resultado esperado

Validar:
- retenção;
- frequência de uso;
- comportamento dos usuários;
- valor percebido.

---

# Fase 2 — Insights Inteligentes

Objetivo:
Transformar eventos acumulados em informações úteis.

## Entregas

### Insights automáticos

Exemplos:
- aumento de episódios respiratórios;
- frequência de coceira;
- comportamento após troca de ração.

---

### Resumo semanal

Exemplo:
- passeios;
- episódios registrados;
- lembretes pendentes;
- tendências básicas.

---

### Score diário simples

Baseado em:
- clima;
- atividade;
- eventos recentes.

---

### Contextualização de riscos

Exemplos:
- calor excessivo;
- esforço físico;
- padrão respiratório.

---

## Resultado esperado

Aumentar:
- retenção;
- percepção de inteligência;
- sensação de acompanhamento contínuo.

---

# Fase 3 — Inteligência Contextual

Objetivo:
Melhorar interpretação e personalização.

## Possíveis entregas

### Melhorias de IA

- interpretação contextual;
- correlação de eventos;
- respostas mais naturais;
- análise longitudinal.

---

### Perfil comportamental do pet

- padrões individuais;
- comportamento normal;
- sensibilidade climática.

---

### Insights personalizados

Exemplos:
- alterações de comportamento;
- aumento de risco;
- padrões recorrentes.

---

## Resultado esperado

Transformar o Braqui em um copiloto mais contextual e inteligente.

---

# Fase 4 — Expansão de Plataforma

Objetivo:
Expandir experiência além do Telegram.

## Possíveis entregas

### Aplicativo mobile

- timeline visual;
- gráficos;
- dashboards simples.

---

### Multi canal

- WhatsApp;
- app mobile;
- web dashboard.

---

### Upload de imagens

Exemplos:
- pele;
- ouvido;
- exames;
- receitas;
- medicamentos.

---

### Analytics

- tendências;
- saúde longitudinal;
- evolução histórica.

---

# Fase 5 — Ecossistema

Objetivo:
Expandir funcionalidades e monetização.

## Possíveis entregas

### Comunidade

- grupos;
- compartilhamento de experiências;
- recomendações.

---

### Marketplace

- produtos;
- rações;
- parceiros;
- planos pet.

---

### Veterinários parceiros

- indicações;
- consultas;
- conteúdos especializados.

---

### Assinatura premium

Possíveis recursos:
- insights avançados;
- histórico completo;
- IA avançada;
- relatórios;
- múltiplos pets.

---

# Roadmap Técnico

## Curto prazo

- monólito modular;
- PostgreSQL;
- Telegram;
- Gemini;
- deploy simples.

---

## Médio prazo

- melhorias de observabilidade;
- filas;
- cache;
- otimizações de IA;
- analytics.

---

## Longo prazo

- multi tenancy;
- escalabilidade horizontal;
- event driven;
- múltiplos canais;
- infraestrutura avançada.

---

# Não Objetivos Atuais

O Braqui NÃO pretende inicialmente:

- substituir veterinários;
- fornecer diagnóstico médico;
- ser rede social pet;
- competir com apps genéricos de pet;
- focar em escala antes da validação.

---

# Métricas de Validação

O MVP será considerado promissor caso consiga demonstrar:

- usuários recorrentes;
- registros frequentes;
- baixa fricção;
- retenção consistente;
- engajamento com alertas;
- percepção clara de valor.

---

# Filosofia de Evolução

O Braqui deve evoluir baseado em:

- comportamento real dos usuários;
- observação contínua;
- aprendizado incremental;
- simplicidade;
- foco em problemas reais.

O produto deve crescer apenas quando houver evidência clara de necessidade.