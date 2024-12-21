package store

const (
	sqliteDbName = "store.db"
)

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RoleUser struct {
	Id       int    `json:"id"`
	RoleId   int    `json:"role_id"`
	UserTgId string `json:"user_tg_id"`
}
