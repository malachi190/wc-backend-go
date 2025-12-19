-- down
DROP TRIGGER IF EXISTS trg_friend_requests_updated ON friend_requests;
DROP TABLE IF EXISTS friend_requests;