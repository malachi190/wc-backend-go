-- down
DROP TRIGGER IF EXISTS trg_watchlists_updated ON watchlists;
DROP TABLE IF EXISTS watchlists;