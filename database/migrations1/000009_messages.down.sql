-- down
DROP TRIGGER IF EXISTS trg_messages_updated ON messages;
DROP TABLE IF EXISTS messages;