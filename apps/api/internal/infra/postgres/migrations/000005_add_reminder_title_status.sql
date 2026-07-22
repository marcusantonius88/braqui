ALTER TABLE reminders
    ADD COLUMN title  VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN status VARCHAR(20)  NOT NULL DEFAULT 'pending';

CREATE INDEX idx_reminders_status ON reminders(status);
