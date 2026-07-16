ALTER TABLE conversation_states RENAME COLUMN state TO current_step;
ALTER TABLE conversation_states RENAME COLUMN data TO payload;
ALTER TABLE conversation_states ADD COLUMN IF NOT EXISTS current_flow VARCHAR(100) NOT NULL DEFAULT '';
