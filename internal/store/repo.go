package store

import (
	"database/sql"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/Dafaque/mentbot/db"
	"github.com/Masterminds/squirrel"
	"github.com/mattn/go-sqlite3"
)

type Repository interface {
	ListRoles() ([]Role, error)
	AddRole(roleName string) error
	RemoveRole(roleName string) error
	ListRoleUsers() ([]RoleUser, error)
	AddRoleUser(roleName, userTgId string) error
	RemoveRoleUser(roleName, userTgId string) error
	GetUsersByRole(roleName string) ([]string, error)
	AddSuperuser(userTgId int64) error
	IsSuperuser(userTgId int64) bool
}

const (
	tableRoles      = "roles"
	tableRoleUsers  = "role_users"
	tableSuperusers = "superusers"
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

// MARK: Roles
func (r *repository) ListRoles() ([]Role, error) {
	sb := squirrel.Select("id", "name").From(tableRoles)
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.Id, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *repository) AddRole(roleName string) error {
	sb := squirrel.Insert(tableRoles).Columns("name").Values(roleName)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

var ErrNoAffectedRows = errors.New("no affected rows")

func (r *repository) RemoveRole(roleName string) error {
	sb := squirrel.Delete(tableRoles).Where(squirrel.Eq{"name": roleName})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoAffectedRows
	}
	return nil
}

// MARK: Role users

func (r *repository) ListRoleUsers() ([]RoleUser, error) {
	sb := squirrel.Select("role_id", "user_tg_id").From(tableRoleUsers)
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roleUsers []RoleUser
	for rows.Next() {
		var roleUser RoleUser
		if err := rows.Scan(&roleUser.RoleId, &roleUser.UserTgId); err != nil {
			return nil, err
		}
		roleUsers = append(roleUsers, roleUser)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roleUsers, nil
}

var ErrRoleUserAlreadyExists = errors.New("role user already exists")

func (r *repository) AddRoleUser(roleName, userTgId string) error {
	roleId, err := r.getRoleId(roleName)
	if err != nil {
		return err
	}
	if roleId == -1 {
		return ErrRoleNotFound
	}
	sb := squirrel.Insert(tableRoleUsers).Columns("role_id", "user_tg_id").Values(roleId, userTgId)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		if errors.Is(err, sqlite3.ErrConstraint) {
			return ErrRoleUserAlreadyExists
		}
		return err
	}
	return nil
}

var ErrRoleNotFound = errors.New("role not found")

func (r *repository) RemoveRoleUser(roleName, userTgId string) error {
	roleId, err := r.getRoleId(roleName)
	if err != nil {
		return err
	}
	if roleId == -1 {
		return ErrRoleNotFound
	}
	sb := squirrel.Delete(tableRoleUsers).Where(squirrel.Eq{"role_id": roleId, "user_tg_id": userTgId})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (r *repository) GetUsersByRole(roleName string) ([]string, error) {
	roleId, err := r.getRoleId(roleName)
	if err != nil {
		return nil, err
	}
	sb := squirrel.Select("user_tg_id").From(tableRoleUsers).Where(squirrel.Eq{"role_id": roleId})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) getRoleId(roleName string) (int, error) {
	sb := squirrel.Select("id").From(tableRoles).Where(squirrel.Eq{"name": roleName})
	query, args, err := sb.ToSql()
	if err != nil {
		return -1, err
	}
	var id int
	err = r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *repository) AddSuperuser(userTgId int64) error {
	sb := squirrel.Insert(tableSuperusers).Columns("user_tg_id").Values(userTgId)
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func (r *repository) IsSuperuser(userTgId int64) bool {
	sb := squirrel.Select("user_tg_id").From(tableSuperusers).Where(squirrel.Eq{"user_tg_id": userTgId})
	query, args, err := sb.ToSql()
	if err != nil {
		return false
	}
	var id int
	err = r.db.QueryRow(query, args...).Scan(&id)
	return err == nil //TODO: possible logic error
}
