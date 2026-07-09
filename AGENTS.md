# AGENTS.md — Braqui

## Project status

Pre-code. The repository contains only documentation (vision, architecture, roadmap, playbook, 23 specs). No `go.mod`, no `main.go`, no Dockerfile, no `docker-compose.yml`, no CI/CD, no tests — zero lines of application code. The `apps/api/` directory tree does not exist yet.

## Architecture

- Monorepo, modular monolith (Go), Clean Architecture / hexagonal-arch light.
- Only `apps/api` will be implemented for MVP. `apps/dashboard` and `apps/admin` are future.
- Layer dependency: `domain` → `application` → `interfaces` ← `infra`. Domain must import zero external packages (no PostgreSQL, Telegram, HTTP, AI, web frameworks).
- Modules (each under `apps/api/internal/{module}`): `pet`, `event`, `reminder`, `timeline`, `conversation`, `router`, `climate`, `insight`, `summary`.
- Entrypoint: `apps/api/cmd/braqui/main.go`.
- Intent: local parser first, AI (Gemini) fallback, friendly fallback if both fail.

## Naming conventions

- Use cases: `CreatePet`, `RegisterEvent`, `GenerateInsights`, `SendReminder` (VerbNoun).
- Repositories: `PetRepository`, `EventRepository`, `ReminderRepository`.
- Providers/gateways: `AIProvider`, `ClimateProvider`, `TelegramGateway`.
- Handlers: `TelegramWebhookHandler`, `ReminderHandler`, `TimelineHandler`.

## Development process

- Spec-Driven Development (SDD). Before writing code, read `docs/vision.md`, `docs/architecture.md`, `docs/playbook.md`, and the relevant spec.
- Each spec ends with an `Implementation Checklist` (`- [ ]` items). Implement only pending items from the relevant spec.
- Project documentation is in Brazilian Portuguese. Specs are numbered `SPEC-NNN-{slug}.md`.
- All docs are the single source of truth until code exists.

## Testing

- Prefer unit tests. Mock external dependencies (`PetRepository`, `TelegramGateway`, `AIProvider`, `ClimateProvider`).
- Integration tests use PostgreSQL via Docker.
- No test framework chosen yet.

## Docker

- Expected containers (not yet created): `api` (Go multi-stage build) and `postgres`.
- Single command: `docker compose up`.
- `.env` file for secrets (`DATABASE_URL`, `TELEGRAM_BOT_TOKEN`, `GEMINI_API_KEY`, `OPENWEATHER_API_KEY`).
- `docker/` dir for auxiliary Docker assets.

## Deploy targets (planned, not configured)

- Render / Railway / Fly.io with managed PostgreSQL.

## Constraints

- No Kubernetes, no Kafka, no microsservices, no CQRS, no event sourcing, no premature enterprise architecture.
- No AI controlling system flow — AI is optional, decoupled, and replaceable.
- Handlers must not contain business logic. Use cases orchestrate domain.
- Repositories must not contain business logic. SQL stays in infra.
- Never commit secrets or real `.env` files.
- Avoid clever code, premature abstractions, excessive generics, and magic.
