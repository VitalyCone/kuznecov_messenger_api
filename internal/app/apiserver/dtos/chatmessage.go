package dtos

import (
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/store"
)

type CreateChatMessageDto struct {
	ChatId int    `json:"chat"`
	UserId int    `json:"user"`
	Text   string `json:"text"`
}

func (c *CreateChatMessageDto) CreateChatMessageDtoToModel(store *store.Store) (*model.ChatMessage, error) {

	chat, err := store.Chat().GetById(c.ChatId)
	if err != nil {
		return nil, err
	}

	user, err := store.User().GetUserByID(c.UserId)
	if err != nil {
		return nil, err
	}

	return &model.ChatMessage{
		Chat: *chat,
		User: *user,
		Text: c.Text,
	}, nil
}
