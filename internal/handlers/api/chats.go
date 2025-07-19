package api

import (
	"context"

	api "github.com/Dafaque/mentbot/internal/api/gen"
)

// AuthorizeChat implements api.StrictServerInterface.
func (h *handler) AuthorizeChat(ctx context.Context, request api.AuthorizeChatRequestObject) (api.AuthorizeChatResponseObject, error) {
	err := h.store.SetChatAuthorized(ctx, request.ChatId, request.Params.Authorized)
	if err != nil {
		return nil, err
	}
	return api.AuthorizeChat200Response{}, nil
}

// GetChats implements api.StrictServerInterface.
func (h *handler) GetChats(ctx context.Context, request api.GetChatsRequestObject) (api.GetChatsResponseObject, error) {
	chats, err := h.store.ListChats(ctx)
	if err != nil {
		return nil, err
	}
	response := api.GetChats200JSONResponse{ChatListResponseJSONResponse: []api.Chat{}}
	for _, chat := range chats {
		response.ChatListResponseJSONResponse = append(response.ChatListResponseJSONResponse, api.Chat{
			Id:         int(chat.ID),
			TgChatId:   int(chat.TgChatID),
			TgChatName: chat.TgChatName,
			Authorized: chat.Authorized,
		})
	}
	return response, nil
}

// GetChatUsers implements api.StrictServerInterface.
func (h *handler) GetChatUsers(ctx context.Context, request api.GetChatUsersRequestObject) (api.GetChatUsersResponseObject, error) {
	users, err := h.store.ListChatUsers(ctx, request.ChatId)
	if err != nil {
		return nil, err
	}
	response := api.GetChatUsers200JSONResponse{UserListResponseJSONResponse: []api.User{}}
	for _, user := range users {
		response.UserListResponseJSONResponse = append(response.UserListResponseJSONResponse, api.User{
			Id:         int(user.ID),
			ChatId:     int(user.ChatID),
			TgChatName: user.TgChatName,
			TgChatId:   int(user.TgChatID),
			TgUserId:   int(user.TgUserID),
			TgUserName: user.TgUserName,
		})
	}
	return response, nil
}

// GetChat implements api.StrictServerInterface.
func (h *handler) GetChat(ctx context.Context, request api.GetChatRequestObject) (api.GetChatResponseObject, error) {
	chat, err := h.store.GetChatByID(ctx, request.ChatId)
	if err != nil {
		return nil, err
	}

	response := api.Chat{
		Id:         int(chat.ID),
		TgChatId:   int(chat.TgChatID),
		TgChatName: chat.TgChatName,
		Authorized: chat.Authorized,
	}
	return api.GetChat200JSONResponse{ChatResponseJSONResponse: api.ChatResponseJSONResponse(response)}, nil
}
