package api

import (
	"context"

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
