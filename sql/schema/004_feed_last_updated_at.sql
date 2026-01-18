-- +goose up
ALTER TABLE feeds ADD COLUMN IF NOT EXISTS
last_fetched_at TIMESTAMP;