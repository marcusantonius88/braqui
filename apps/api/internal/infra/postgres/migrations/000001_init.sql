CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    telegram_id BIGINT       NOT NULL UNIQUE,
    first_name  VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE pets (
    id        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id   UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name      VARCHAR(255) NOT NULL,
    breed     VARCHAR(255) NOT NULL,
    age       INT          NOT NULL DEFAULT 0,
    weight    NUMERIC(5,2) NOT NULL DEFAULT 0,
    location  VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_pets_user_id ON pets(user_id);

CREATE TABLE events (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pet_id      UUID         NOT NULL REFERENCES pets(id) ON DELETE CASCADE,
    type        VARCHAR(50)  NOT NULL,
    description TEXT         NOT NULL DEFAULT '',
    timestamp   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_events_pet_id ON events(pet_id);
CREATE INDEX idx_events_timestamp ON events(timestamp);

CREATE TABLE reminders (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pet_id          UUID         NOT NULL REFERENCES pets(id) ON DELETE CASCADE,
    type            VARCHAR(50)  NOT NULL,
    description     TEXT         NOT NULL DEFAULT '',
    due_date        TIMESTAMPTZ  NOT NULL,
    repeat_interval VARCHAR(50)  NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reminders_pet_id ON reminders(pet_id);

CREATE TABLE conversation_states (
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id      UUID         NOT NULL REFERENCES users(id) ON DELETE CASCADE UNIQUE,
    current_flow VARCHAR(100) NOT NULL DEFAULT '',
    current_step VARCHAR(100) NOT NULL DEFAULT 'idle',
    payload      JSONB        NOT NULL DEFAULT '{}',
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_conversation_states_user_id ON conversation_states(user_id);
