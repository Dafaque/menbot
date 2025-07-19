CREATE TABLE IF NOT EXISTS chats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tg_chat_id BIGINT NOT NULL,
    tg_chat_name VARCHAR(255) NOT NULL,
    authorized BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_chat_tg_id_idx ON chats (tg_chat_id);
