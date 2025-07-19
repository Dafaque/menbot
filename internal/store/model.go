package store

const (
	sqliteDbName = "store.db"
)

type Chat struct {
	ID         int64  `json:"id"`
	TgChatID   int64  `json:"tg_chat_id"`
	TgChatName string `json:"tg_chat_name"`
	Authorized bool   `json:"authorized"`
}

type ChatUser struct {
	ID         int64  `json:"id"`
	ChatID     int64  `json:"chat_id"`
	TgChatName string `json:"tg_chat_name"`
	TgChatID   int64  `json:"tg_chat_id"`
	TgUserID   int64  `json:"tg_user_id"`
	TgUserName string `json:"tg_user_name"`
}

type Role struct {
	ID         int64  `json:"id"`
	ChatID     int64  `json:"chat_id"`
	TgChatID   int64  `json:"tg_chat_id"`
	TgChatName string `json:"tg_chat_name"`
	Name       string `json:"name"`
}

type RoleUser struct {
	ID         int64  `json:"id"`
	RoleID     int64  `json:"role_id"`
	RoleName   string `json:"role_name"`
	ChatID     int64  `json:"chat_id"`
	TgUserID   int64  `json:"tg_user_id"`
	TgUserName string `json:"tg_user_name"`
	TgChatName string `json:"tg_chat_name"`
	TgChatID   int64  `json:"tg_chat_id"`
}
