package tg

import (
	"context"

	"github.com/Dafaque/mentbot/internal/config"
	"github.com/Dafaque/mentbot/internal/store"
)

type Handler interface {
	NewChat(ctx context.Context, tgChatID, tgUserID int64, tgChatName, tgUserName string) error
	Register(ctx context.Context, tgChatID, tgUserID int64, tgUserName string) error
	RoleUsers(ctx context.Context, tgChatID int64, roleName string) ([]string, error)
	Subscribe(ctx context.Context, tgChatID, tgUserID int64, roleName string) error
	RolesForBotCommands(ctx context.Context) (map[int64][]string, error)
}

type handler struct {
	repo store.Repository
}

func New(repo store.Repository, cfg *config.Config) Handler {
	return &handler{repo: repo}
}

func (h *handler) NewChat(
	ctx context.Context,
	tgChatID, tgUserID int64,
	tgChatName, tgUserName string,
) error {
	err := h.repo.AddChat(ctx, tgChatID, tgChatName)
	if err != nil {
		return err
	}
	chat, err := h.repo.FindChat(ctx, tgChatID)
	if err != nil {
		return err
	}
	return h.repo.AddChatUser(ctx, chat.ID, tgUserID, tgUserName)
}

func (h *handler) Register(ctx context.Context, tgChatID int64, tgUserID int64, tgUserName string) error {
	chat, err := h.repo.FindChat(ctx, tgChatID)
	if err != nil {
		return err
	}
	return h.repo.AddChatUser(ctx, chat.ID, tgUserID, tgUserName)
}

func (h *handler) RoleUsers(ctx context.Context, tgChatID int64, roleName string) ([]string, error) {
	chat, err := h.repo.FindChat(ctx, tgChatID)
	if err != nil {
		return nil, err
	}
	role, err := h.repo.FindRole(ctx, chat.ID, roleName)
	if err != nil {
		return nil, err
	}
	roleUsers, err := h.repo.ListRoleUsers(ctx, chat.ID, role.ID)
	if err != nil {
		return nil, err
	}
	users := make([]string, len(roleUsers))
	for i, roleUser := range roleUsers {
		users[i] = roleUser.TgUserName
	}
	return users, nil
}

func (h *handler) Subscribe(ctx context.Context, tgChatID, tgUserID int64, roleName string) error {
	chat, err := h.repo.FindChat(ctx, tgChatID)
	if err != nil {
		return err
	}
	role, err := h.repo.FindRole(ctx, chat.ID, roleName)
	if err != nil {
		return err
	}
	chatUser, err := h.repo.FindChatUser(ctx, chat.ID, tgUserID)
	if err != nil {
		return err
	}
	return h.repo.AssignRole(ctx, chat.ID, role.ID, chatUser.ID)

}

func (h *handler) RolesForBotCommands(ctx context.Context) (map[int64][]string, error) {
	roles, err := h.repo.ListAllRolesForBotCommands(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
