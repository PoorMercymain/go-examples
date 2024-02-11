BEGIN;
CREATE INDEX IF NOT EXISTS users_idx ON users USING BTREE (username, user_description);
COMMIT;