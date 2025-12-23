-- +goose Up
CREATE TABLE friend_requests (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status       TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending','accepted','rejected')),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (sender_id, receiver_id)
);

CREATE INDEX idx_friend_requests_sender_id   ON friend_requests(sender_id);
CREATE INDEX idx_friend_requests_receiver_id ON friend_requests(receiver_id);

CREATE TRIGGER trg_friend_requests_updated
    BEFORE UPDATE ON friend_requests
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS trg_friend_requests_updated ON friend_requests;
DROP TABLE IF EXISTS friend_requests;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
