-- up
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

CREATE INDEX idx_watchlists_user_id  ON watchlists(user_id);
CREATE INDEX idx_watchlists_tmdb_id  ON watchlists(tmdb_id);

CREATE TRIGGER trg_watchlists_updated
    BEFORE UPDATE ON watchlists
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();