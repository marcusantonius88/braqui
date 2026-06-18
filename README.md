![Braqui Banner](docs/assets/banner-braqui.png)

![Go](https://img.shields.io/badge/Go-Latest-blue) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791) ![Docker](https://img.shields.io/badge/Docker-Container-2496ED) ![Telegram](https://img.shields.io/badge/Telegram-Bot%20API-0088cc) ![Gemini](https://img.shields.io/badge/Gemini-AI-white) ![OpenWeather](https://img.shields.io/badge/OpenWeather-API-orange)

# 🐶 Braqui

**Saúde inteligente para cães braquicefálicos.**

Braqui é uma plataforma focada no acompanhamento da saúde e rotina de cães braquicefálicos, como Buldogue Francês, Pug, Shih Tzu, Boston Terrier, Bulldog Inglês e raças similares.

O objetivo é ajudar tutores a registrarem eventos importantes da rotina do pet, receberem lembretes, alertas climáticos e acompanharem a evolução da saúde do animal através de uma interface conversacional simples via Telegram.

---

# 🎯 Problema

Cães braquicefálicos possuem necessidades específicas relacionadas principalmente a:

* respiração;
* calor excessivo;
* ofegância;
* obesidade;
* alergias;
* problemas dermatológicos;
* rotina medicamentosa.

Grande parte dos tutores não possui um histórico organizado dos acontecimentos relacionados ao pet.

Informações importantes acabam se perdendo:

* episódios de vômito;
* crises de coceira;
* consultas veterinárias;
* medicamentos administrados;
* mudanças de comportamento.

O Braqui nasce para resolver esse problema.

---

# 🚀 MVP

O MVP possui foco em:

* cadastro do tutor;
* cadastro do pet;
* registro de eventos;
* histórico de eventos;
* insights básicos;
* lembretes;
* alertas climáticos;
* resumo semanal.

Tudo através de uma experiência conversacional no Telegram.

---

# 🏗️ Arquitetura

O projeto utiliza arquitetura modular organizada em monorepo.

```text
/apps
  /api
  /dashboard
  /admin

/docs
  /specs

docker-compose.yml
README.md
```

Inicialmente apenas a aplicação `api` será implementada.

As aplicações `dashboard` e `admin` permanecem previstas para futuras evoluções.

---

# ⚙️ Stack Tecnológica

## Backend

* Go
* PostgreSQL
* Docker
* Docker Compose

## Integrações

* Telegram Bot API
* Google Gemini
* OpenWeather

## Infraestrutura

* Docker
* GitHub
* Deploy Cloud (Render / Railway / Fly.io)

---

# 🐳 Ambiente Local

Subir ambiente:

```bash
docker compose up
```

Parar ambiente:

```bash
docker compose down
```

Rebuild:

```bash
docker compose up --build
```

---

# 📚 Documentação

Toda a documentação do projeto encontra-se na pasta:

```text
/docs
```

Arquivos principais:

```text
docs/
├── vision.md
├── architecture.md
├── roadmap.md
├── playbook.md
└── specs/
```

---

# 🧠 Spec Driven Development (SDD)

O Braqui está sendo desenvolvido utilizando a abordagem **Spec Driven Development (SDD)**.

Ao invés de implementar funcionalidades diretamente a partir de ideias ou prompts extensos, todo o projeto foi decomposto em especificações independentes e rastreáveis.

Cada funcionalidade possui:

* objetivo;
* contexto;
* escopo;
* critérios de aceite;
* checklist de implementação.

Exemplo:

```text
SPEC-001-bootstrap-e-estrutura-inicial-do-projeto.md
SPEC-002-configuracao-e-gerenciamento-de-ambiente.md
SPEC-003-deploy-e-infraestrutura-inicial.md
...
SPEC-023-resumo-semanal.md
```

---

# 🤖 IA-Assisted Development

O desenvolvimento do Braqui utiliza Inteligência Artificial como ferramenta de engenharia de software.

O fluxo adotado é:

```text
Vision
    ↓
Architecture
    ↓
Roadmap
    ↓
Specs
    ↓
Implementação
    ↓
Checklist da Spec
```

Cada implementação deve:

1. Ler a documentação.
2. Implementar apenas uma Spec por vez.
3. Atualizar o checklist da Spec.
4. Executar testes.
5. Registrar pendências.

---

# 💻 OpenCode

O projeto foi estruturado especificamente para funcionar bem com ferramentas de IA como:

* OpenCode
* Cursor
* Windsurf
* Kiro
* GitHub Copilot

A estratégia adotada evita prompts gigantes e prioriza:

* contexto persistente;
* especificações pequenas;
* implementação incremental;
* validação contínua.

---

# 📋 Roadmap Inicial

## Fundação

* Bootstrap
* Configuração
* Infraestrutura
* Docker
* Persistência

## Plataforma

* Telegram
* Onboarding
* Cadastro de Pet
* Estado Conversacional
* Router

## Inteligência

* Parser Local
* IA
* Registro de Eventos
* Timeline

## Automação

* Scheduler
* Lembretes
* Alertas Climáticos

## Valor ao Usuário

* Insights
* Resumo Semanal

---

# 🔒 Status

Projeto em desenvolvimento.

Atualmente em fase de implementação do MVP.

---

# 📄 Licença

MIT
