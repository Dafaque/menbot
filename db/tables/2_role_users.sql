CREATE TABLE IF NOT EXISTS role_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_id INT NOT NULL,
    user_tg_id TEXT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_role_users_role_id ON role_users (role_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_role_users_role_id_user_tg_id ON role_users (role_id, user_tg_id);
