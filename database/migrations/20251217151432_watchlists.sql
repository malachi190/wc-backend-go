-- +goose Up
CREATE TABLE watchlists (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tmdb_id      INTEGER NOT NULL,
    media_type   TEXT NOT NULL CHECK (media_type IN ('movie','tv')),
    title        TEXT NOT NULL,
    poster_path  TEXT,
    status       TEXT NOT NULL DEFAULT 'plan_to_watch',
    rating       REAL CHECK (rating BETWEEN 0 AND 10),
    notes        TEXT,
    added_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS trg_watchlists_updated ON watchlists;
DROP TABLE IF EXISTS watchlists;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
