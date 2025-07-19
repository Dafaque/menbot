package api

import (
	"context"
	"log"

	api "github.com/Dafaque/mentbot/internal/api/gen"
)

// CreateRole implements api.StrictServerInterface.
func (h *handler) CreateRole(ctx context.Context, request api.CreateRoleRequestObject) (api.CreateRoleResponseObject, error) {
	err := h.store.AddRole(ctx, request.ChatId, request.Body.Name)
	if err != nil {
		return nil, err
	}
	return api.CreateRole200Response{}, nil
}

// GetChatRoles implements api.StrictServerInterface.
func (h *handler) GetChatRoles(ctx context.Context, request api.GetChatRolesRequestObject) (api.GetChatRolesResponseObject, error) {
	roles, err := h.store.ListRoles(ctx, request.ChatId)
	if err != nil {
		return nil, err
	}
	response := api.GetChatRoles200JSONResponse{RoleListResponseJSONResponse: []api.Role{}}
	for _, role := range roles {
		response.RoleListResponseJSONResponse = append(response.RoleListResponseJSONResponse, api.Role{
			Id:         int(role.ID),
			ChatId:     int(role.ChatID),
			TgChatId:   int(role.TgChatID),
			TgChatName: role.TgChatName,
			Name:       role.Name,
		})
	}
	return response, nil
}

// GetRoleUsers implements api.StrictServerInterface.
func (h *handler) GetRoleUsers(ctx context.Context, request api.GetRoleUsersRequestObject) (api.GetRoleUsersResponseObject, error) {
	users, err := h.store.ListRoleUsers(ctx, request.ChatId, request.RoleId)
	if err != nil {
		return nil, err
	}
	response := api.GetRoleUsers200JSONResponse{UserListResponseJSONResponse: []api.User{}}
	for _, user := range users {
		response.UserListResponseJSONResponse = append(response.UserListResponseJSONResponse, api.User{
			Id:         int(user.ID),
			ChatId:     int(user.ChatID),
			TgChatId:   int(user.TgChatID),
			TgChatName: user.TgChatName,
			TgUserId:   int(user.TgUserID),
			TgUserName: user.TgUserName,
		})
	}
	return response, nil
}

// AssignRole implements api.StrictServerInterface.
func (h *handler) AssignRole(ctx context.Context, request api.AssignRoleRequestObject) (api.AssignRoleResponseObject, error) {
	for _, user := range *request.Body {
		var err error
		if user.Assign {
			err = h.store.AssignRole(ctx, request.ChatId, request.RoleId, int64(user.UserId))
		} else {
			err = h.store.UnassignRole(ctx, request.ChatId, request.RoleId, int64(user.UserId))
		}
		if err != nil {
			log.Println(err)
		}
	}
	return api.AssignRole200Response{}, nil
}
