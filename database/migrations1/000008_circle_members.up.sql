-- up
CREATE TABLE circle_members (
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    circle_id UUID NOT NULL REFERENCES circles(id) ON DELETE CASCADE,
    user_id   UUID NOT NULL REFERENCES users(id)   ON DELETE CASCADE,
    role      TEXT NOT NULL DEFAULT 'member' CHECK (role IN ('admin','moderator','member')),
    joined_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (circle_id, user_id)
);

CREATE INDEX idx_circle_members_circle_id ON circle_members(circle_id);
CREATE INDEX idx_circle_members_user_id   ON circle_members(user_id);