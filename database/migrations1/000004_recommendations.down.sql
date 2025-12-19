-- down
DROP TRIGGER IF EXISTS trg_recommendations_updated ON recommendations;
DROP TABLE IF EXISTS recommendations;