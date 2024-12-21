package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Dafaque/tgment/internal/config"
	"github.com/Dafaque/tgment/internal/store"
)

type Handler interface {
	ListRoles() string
	AddRole(reqUserTgId int64, roleName string) string
	RemoveRole(reqUserTgId int64, roleName string) string
	ListRoleUsers(reqUserTgId int64) string
	AddRoleUser(reqUserTgId int64, roleName, userTgId string) string
	RemoveRoleUser(reqUserTgId int64, roleName, userTgId string) string
	AddSuperuser(token string, userTgId int64) string
	GetUsersByRole(roleName string) string
	IsUper(userTgId int64) bool
}

type handler struct {
	repo    store.Repository
	suToken string
}

func New(repo store.Repository, cfg *config.Config) Handler {
	return &handler{repo: repo, suToken: cfg.SuToken}
}

func (h *handler) ListRoles() string {
	roles, err := h.repo.ListRoles()
	if err != nil {
		return err.Error()
	}
	if len(roles) == 0 {
		return "No roles found"
	}
	buf := &strings.Builder{}
	buf.WriteString("Roles:\n")
	for _, role := range roles {
		fmt.Fprintf(buf, "	\\- %s\n", role.Name)
	}
	return buf.String()
}

func (h *handler) AddRole(reqUserTgId int64, roleName string) string {
	if !h.IsUper(reqUserTgId) {
		return "You are not allowed to add roles"
	}
	err := h.repo.AddRole(roleName)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "UNIQUE constraint failed") {
			return fmt.Sprintf("Role *%s* already exists", roleName)
		}
		return msg
	}
	return fmt.Sprintf("Role %s added", roleName)
}

func (h *handler) RemoveRole(reqUserTgId int64, roleName string) string {
	if !h.IsUper(reqUserTgId) {
		return "You are not allowed to remove roles"
	}
	err := h.repo.RemoveRole(roleName)
	if err != nil {
		if errors.Is(err, store.ErrNoAffectedRows) {
			return fmt.Sprintf("Role *%s* not found", roleName)
		}
		return err.Error()
	}
	return fmt.Sprintf("Role %s removed", roleName)
}

var msgUndefinedRole = "Undefined role %d"

func (h *handler) ListRoleUsers(reqUserTgId int64) string {
	if !h.IsUper(reqUserTgId) {
		return "You are not allowed to list role users"
	}
	roles, err := h.repo.ListRoles()
	if err != nil {

		return err.Error()
	}

	users, err := h.repo.ListRoleUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return "No users assigned to any role"
		}
		return err.Error()
	}

	roleIdMap := make(map[int]string)
	for _, role := range roles {
		roleIdMap[role.Id] = role.Name
	}
	usersByRoles := make(map[string][]string)

	for _, user := range users {
		roleName, found := roleIdMap[user.RoleId]
		if !found {
			undefinedRoleKey := fmt.Sprintf(msgUndefinedRole, user.RoleId)
			if _, ok := usersByRoles[undefinedRoleKey]; !ok {
				usersByRoles[undefinedRoleKey] = []string{}
			}
			usersByRoles[undefinedRoleKey] = append(usersByRoles[undefinedRoleKey], user.UserTgId)
		}
		usersByRoles[roleName] = append(usersByRoles[roleName], user.UserTgId)
	}

	if len(usersByRoles) == 0 {
		return "No users assigned to any role"
	}

	buf := &strings.Builder{}
	for role, users := range usersByRoles {
		fmt.Fprintf(buf, "%s:\n", role)
		for _, user := range users {
			fmt.Fprintf(buf, "	\\- %s\n", user)
		}
	}
	return buf.String()
}

func (h *handler) AddRoleUser(reqUserTgId int64, roleName, userTgId string) string {
	if !h.IsUper(reqUserTgId) {
		return "You are not allowed to add users to roles"
	}
	err := h.repo.AddRoleUser(roleName, userTgId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Sprintf("Role *%s* not found", roleName)
		}
		msg := err.Error()
		if strings.Contains(msg, "UNIQUE constraint failed") {
			return fmt.Sprintf("User *%s* already assigned to role *%s*", userTgId, roleName)
		}
		return msg
	}
	return fmt.Sprintf("User %s added to role", userTgId)
}

func (h *handler) RemoveRoleUser(reqUserTgId int64, roleName, userTgId string) string {
	if !h.IsUper(reqUserTgId) {
		return "You are not allowed to remove users from roles"
	}
	err := h.repo.RemoveRoleUser(roleName, userTgId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Sprintf("Role *%s* or user *%s* not found", roleName, userTgId)
		}
		return err.Error()
	}
	return "User removed from role"
}

func (h *handler) AddSuperuser(token string, userTgId int64) string {
	if token != h.suToken {
		return "Invalid superuser token"
	}
	err := h.repo.AddSuperuser(userTgId)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return "You are already a superuser"
		}
		return err.Error()
	}
	return "You are now a superuser\\! Use help command to see available commands"
}

func (h *handler) GetUsersByRole(roleName string) string {
	users, err := h.repo.GetUsersByRole(roleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "No users assigned to role"
		}
		return err.Error()
	}
	return strings.Join(users, ", ")
}

func (h *handler) IsUper(userTgId int64) bool {
	return h.repo.IsSuperuser(userTgId)
}
