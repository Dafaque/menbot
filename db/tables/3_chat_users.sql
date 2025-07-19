CREATE TABLE IF NOT EXISTS chat_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INT NOT NULL,
    tg_user_id BIGINT NOT NULL,
    tg_user_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_chat_user_idx ON chat_users (chat_id, tg_user_id);