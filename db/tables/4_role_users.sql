CREATE TABLE IF NOT EXISTS role_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_id INT NOT NULL,
    chat_id INT NOT NULL,
    chat_user_id INT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE,
    FOREIGN KEY (chat_user_id) REFERENCES chat_users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_role_users_role_id ON role_users (role_id);
CREATE UNIQUE INDEX IF NOT EXISTS unique_user_role_per_chat_idx ON role_users (role_id, chat_id, chat_user_id);
