package store

import (
	"context"
	"database/sql"
	"errors"
	_ "errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/Dafaque/mentbot/db"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	AddChat(ctx context.Context, tgChatID int64, tgChatName string) error
	RemoveChat(ctx context.Context, chatID int64) error
	ListChats(ctx context.Context) ([]Chat, error)
	FindChat(ctx context.Context, tgChatID int64) (Chat, error)
	SetChatAuthorized(ctx context.Context, chatID int64, authorized bool) error
	GetChatByID(ctx context.Context, chatID int64) (Chat, error)

	AddChatUser(ctx context.Context, chatID, tgUserID int64, tgUserName string) error
	RemoveChatUser(ctx context.Context, recordID int64) error
	ListChatUsers(ctx context.Context, chatID int64) ([]ChatUser, error)
	FindChatUser(ctx context.Context, chatID, tgUserID int64) (ChatUser, error)

	AddRole(ctx context.Context, chatID int64, roleName string) error
	RemoveRole(ctx context.Context, recordID int64) error
	ListRoles(ctx context.Context, chatID int64) ([]Role, error)
	FindRole(ctx context.Context, chatID int64, roleName string) (Role, error)
	ListAllRolesForBotCommands(ctx context.Context) (map[int64][]string, error)

	AssignRole(ctx context.Context, chatID, roleID, chatUserID int64) error
	UnassignRole(ctx context.Context, chatID, roleID, chatUserID int64) error
	ListRoleUsers(ctx context.Context, chatID, roleID int64) ([]RoleUser, error)

	Done() error
}

const (
	tableRoles     = "roles"
	tableRoleUsers = "role_users"
	tableChats     = "chats"
	tableChatUsers = "chat_users"
)

type repository struct {
	db *sql.DB
}

func New(appPath string) (Repository, error) {
	if _, err := os.Stat(appPath); os.IsNotExist(err) {
		err = os.MkdirAll(appPath, 0755)
		if err != nil {
			return nil, err
		}
	}
	d, err := sql.Open("sqlite3", filepath.Join(appPath, sqliteDbName))
	if err != nil {
		return nil, err
	}
	err = d.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Creating tables...")
	err = fs.WalkDir(db.Tables, "tables", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		content, err := db.Tables.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = d.Exec(string(content))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	log.Println("Tables created")
	return &repository{db: d}, nil
}

// MARK: - Chats
func (r *repository) AddChat(ctx context.Context, tgChatID int64, tgChatName string) error {
	sb := squirrel.Insert(tableChats).Columns("tg_chat_id", "tg_chat_name").Values(tgChatID, tgChatName)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) RemoveChat(ctx context.Context, recordID int64) error {
	sb := squirrel.Delete(tableChats).Where(squirrel.Eq{"id": recordID})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) ListChats(ctx context.Context) ([]Chat, error) {
	sb := squirrel.Select(
		"id",
		"tg_chat_id",
		"tg_chat_name",
		"authorized",
	).
		From(tableChats)
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	chats := []Chat{}
	for rows.Next() {
		var chat Chat
		err = rows.Scan(&chat.ID, &chat.TgChatID, &chat.TgChatName, &chat.Authorized)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func (r *repository) FindChat(ctx context.Context, tgChatID int64) (Chat, error) {
	sb := squirrel.Select(
		"id",
		"tg_chat_id",
		"tg_chat_name",
	).From(tableChats).
		Where(squirrel.Eq{"tg_chat_id": tgChatID}).
		Limit(1)
	query, args, err := sb.ToSql()
	if err != nil {
		return Chat{}, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return Chat{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return Chat{}, errors.New("chat not found")
	}
	var chat Chat
	err = rows.Scan(&chat.ID, &chat.TgChatID, &chat.TgChatName)
	if err != nil {
		return Chat{}, err
	}
	return chat, nil
}

func (r *repository) SetChatAuthorized(ctx context.Context, chatID int64, authorized bool) error {
	sb := squirrel.Update(tableChats).Set("authorized", authorized).Where(squirrel.Eq{"id": chatID})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) GetChatByID(ctx context.Context, chatID int64) (Chat, error) {
	sb := squirrel.Select("id", "tg_chat_id", "tg_chat_name", "authorized").From(tableChats).Where(squirrel.Eq{"id": chatID})
	query, args, err := sb.ToSql()
	if err != nil {
		return Chat{}, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return Chat{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return Chat{}, errors.New("chat not found")
	}
	var chat Chat
	err = rows.Scan(&chat.ID, &chat.TgChatID, &chat.TgChatName, &chat.Authorized)
	if err != nil {
		return Chat{}, err
	}
	return chat, nil
}

// MARK: - Chat Users
func (r *repository) AddChatUser(ctx context.Context, chatID int64, tgUserID int64, tgUserName string) error {
	sb := squirrel.Insert(tableChatUsers).
		Columns("chat_id", "tg_user_id", "tg_user_name").
		Values(chatID, tgUserID, tgUserName)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) RemoveChatUser(ctx context.Context, recordID int64) error {
	sb := squirrel.Delete(tableChatUsers).Where(squirrel.Eq{"id": recordID})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) ListChatUsers(ctx context.Context, chatID int64) ([]ChatUser, error) {
	sb := squirrel.Select(
		"cu.id",
		"cu.chat_id",
		"cu.tg_user_id",
		"cu.tg_user_name",
		"c.tg_chat_id",
		"c.tg_chat_name",
	).From(tableChatUsers + " AS cu").
		Join(tableChats + " AS c ON cu.chat_id = c.id").
		Where(squirrel.Eq{"chat_id": chatID})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	chatUsers := []ChatUser{}
	for rows.Next() {
		var chatUser ChatUser
		err = rows.Scan(
			&chatUser.ID,
			&chatUser.ChatID,
			&chatUser.TgUserID,
			&chatUser.TgUserName,
			&chatUser.TgChatID,
			&chatUser.TgChatName,
		)
		if err != nil {
			return nil, err
		}
		chatUsers = append(chatUsers, chatUser)
	}
	return chatUsers, nil
}

func (r *repository) FindChatUser(ctx context.Context, chatID, tgUserID int64) (ChatUser, error) {
	sb := squirrel.Select(
		"cu.id",
		"cu.chat_id",
		"cu.tg_user_id",
		"cu.tg_user_name",
		"c.tg_chat_id",
		"c.tg_chat_name",
	).From(tableChatUsers + " AS cu").
		Join(tableChats + " AS c ON cu.chat_id = c.id").
		Where(squirrel.Eq{"chat_id": chatID, "tg_user_id": tgUserID}).
		Limit(1)
	query, args, err := sb.ToSql()
	if err != nil {
		return ChatUser{}, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return ChatUser{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return ChatUser{}, errors.New("chat user not found")
	}
	var chatUser ChatUser
	err = rows.Scan(
		&chatUser.ID,
		&chatUser.ChatID,
		&chatUser.TgUserID,
		&chatUser.TgUserName,
		&chatUser.TgChatID,
		&chatUser.TgChatName,
	)
	if err != nil {
		return ChatUser{}, err
	}
	return chatUser, nil
}

// MARK: - Roles
func (r *repository) AddRole(ctx context.Context, chatID int64, roleName string) error {
	sb := squirrel.Insert(tableRoles).Columns("chat_id", "name").Values(chatID, roleName)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) RemoveRole(ctx context.Context, recordID int64) error {
	sb := squirrel.Delete(tableRoles).Where(squirrel.Eq{"id": recordID})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) ListRoles(ctx context.Context, chatID int64) ([]Role, error) {
	sb := squirrel.Select(
		"r.id",
		"r.chat_id",
		"r.name",
		"c.tg_chat_id",
		"c.tg_chat_name",
	).From(tableRoles + " AS r").
		Join(tableChats + " AS c ON r.chat_id = c.id").
		Where(squirrel.Eq{"chat_id": chatID})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles := []Role{}
	for rows.Next() {
		var role Role
		err = rows.Scan(
			&role.ID,
			&role.ChatID,
			&role.Name,
			&role.TgChatID,
			&role.TgChatName,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *repository) FindRole(ctx context.Context, chatID int64, roleName string) (Role, error) {
	sb := squirrel.Select(
		"r.id",
		"r.chat_id",
		"r.name",
		"c.tg_chat_id",
		"c.tg_chat_name",
	).From(tableRoles + " AS r").
		Join(tableChats + " AS c ON r.chat_id = c.id").
		Where(squirrel.Eq{"r.chat_id": chatID, "r.name": roleName}).
		Limit(1)
	query, args, err := sb.ToSql()
	if err != nil {
		return Role{}, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return Role{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return Role{}, errors.New("role not found")
	}
	var role Role
	err = rows.Scan(
		&role.ID,
		&role.ChatID,
		&role.Name,
		&role.TgChatID,
		&role.TgChatName,
	)
	if err != nil {
		return Role{}, err
	}
	return role, nil
}

func (r *repository) ListAllRolesForBotCommands(ctx context.Context) (map[int64][]string, error) {
	sb := squirrel.
		Select(
			"r.name",
			"c.tg_chat_id",
		).From(tableRoles + " AS r").
		Join(tableChats + " AS c ON r.chat_id = c.id")
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles := map[int64][]string{}
	for rows.Next() {
		var roleName string
		var tgChatID int64
		err = rows.Scan(&roleName, &tgChatID)
		if err != nil {
			return nil, err
		}
		if _, ok := roles[tgChatID]; !ok {
			roles[tgChatID] = []string{}
		}
		roles[tgChatID] = append(roles[tgChatID], roleName)
	}
	return roles, nil
}

// MARK: - Role Users
func (r *repository) AssignRole(ctx context.Context, chatID, roleID, chatUserID int64) error {
	sb := squirrel.Insert(tableRoleUsers).Columns("chat_id", "role_id", "chat_user_id").Values(chatID, roleID, chatUserID)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) ListRoleUsers(ctx context.Context, chatID, roleID int64) ([]RoleUser, error) {
	sb := squirrel.Select(
		"ru.id",
		"ru.chat_id",
		"ru.role_id",
		"cu.tg_user_id",
		"cu.tg_user_name",
		"c.tg_chat_id",
		"c.tg_chat_name",
	).From(tableRoleUsers + " AS ru").
		Join(tableChatUsers + " AS cu ON ru.chat_user_id = cu.id").
		Join(tableChats + " AS c ON ru.chat_id = c.id").
		Where(squirrel.Eq{"ru.chat_id": chatID, "ru.role_id": roleID})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roleUsers := []RoleUser{}
	for rows.Next() {
		var roleUser RoleUser
		err = rows.Scan(
			&roleUser.ID,
			&roleUser.ChatID,
			&roleUser.RoleID,
			&roleUser.TgUserID,
			&roleUser.TgUserName,
			&roleUser.TgChatID,
			&roleUser.TgChatName,
		)
		if err != nil {
			return nil, err
		}
		roleUsers = append(roleUsers, roleUser)
	}
	return roleUsers, nil
}

func (r *repository) UnassignRole(ctx context.Context, chatID, roleID, chatUserID int64) error {
	sb := squirrel.Delete(tableRoleUsers).Where(squirrel.Eq{"chat_id": chatID, "role_id": roleID, "chat_user_id": chatUserID})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *repository) Done() error {
	return r.db.Close()
}
