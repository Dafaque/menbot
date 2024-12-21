CREATE TABLE IF NOT EXISTS superusers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_tg_id TEXT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_superusers_user_tg_id ON superusers (user_tg_id);