-- +goose Up
CREATE TABLE circles (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name         TEXT NOT NULL,
    description  TEXT,
    avatar       TEXT,
    is_private   BOOLEAN NOT NULL DEFAULT true,
    created_by_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_circles_created_by_id ON circles(created_by_id);

CREATE TRIGGER trg_circles_updated
    BEFORE UPDATE ON circles
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS trg_circles_updated ON circles;
DROP TABLE IF EXISTS circles;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
