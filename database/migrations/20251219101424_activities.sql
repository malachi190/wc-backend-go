-- +goose Up
CREATE TABLE activities (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type        TEXT NOT NULL CHECK (type IN ('added_to_watchlist','completed_movie','rated_movie','joined_circle','left_circle','created_circle','deleted_circle','sent_friend_request','accepted_friend_request','rejected_friend_request','sent_recommendation','accepted_recommendation','rejected_recommendation','updated_profile','reached_milestone','watched_episode','reviewed_movie','reviewed_show','followed_user','unfollowed_user','circle_invite_accepted','circle_invite_rejected','new_message','media_comment','media_like','media_unlike','list_comment','list_like','list_unlike','circle_post','circle_comment','circle_post_like','circle_post_unlike','watchlist_comment','watchlist_like','watchlist_unlike','custom_activity')),
    tmdb_id     INTEGER,
    media_type  TEXT CHECK (media_type IN ('movie','tv')),
    title       TEXT,
    description TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_activities_user_id    ON activities(user_id);
CREATE INDEX idx_activities_created_at ON activities(created_at);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS activities;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
