-- +goose Up
CREATE TABLE messages (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    circle_id    UUID NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    sender_id    UUID NOT NULL REFERENCES users(id)   ON DELETE CASCADE,
    content      TEXT NOT NULL,
    tmdb_id      INTEGER,
    media_type   TEXT CHECK (media_type IN ('movie','tv')),
    attachments  JSONB,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_messages_circle_id ON messages(circle_id);
CREATE INDEX idx_messages_sender_id ON messages(sender_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);

CREATE TRIGGER trg_messages_updated
    BEFORE UPDATE ON messages
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS trg_messages_updated ON messages;
DROP TABLE IF EXISTS messages;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
