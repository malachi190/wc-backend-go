-- +goose Up
CREATE TABLE recommendations (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    to_user_id   UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tmdb_id      INTEGER NOT NULL,
    media_type   TEXT NOT NULL CHECK (media_type IN ('movie','tv')),
    title        TEXT NOT NULL,
    poster_path  TEXT,
    message      TEXT,
    status       TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending','accepted','rejected')),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_recommendations_from_user_id ON recommendations(from_user_id);
CREATE INDEX idx_recommendations_to_user_id   ON recommendations(to_user_id);

CREATE TRIGGER trg_recommendations_updated
    BEFORE UPDATE ON recommendations
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS trg_recommendations_updated ON recommendations;
DROP TABLE IF EXISTS recommendations;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
