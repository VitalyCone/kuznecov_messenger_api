package dtos

import (
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/store"
)

type CreateChatDto struct {
	User1Id int `json:"user1_id"`
	User2Id int `json:"user2_id"`
}

func (c *CreateChatDto) CreateChatDtoToModel(s *store.Store) (*model.Chat, error) {
	user := s.User()
	u1, err := user.GetUserByID(c.User1Id)
	if err != nil {
		return nil, err
	}

	u2, err := user.GetUserByID(c.User2Id)
	if err != nil {
		return nil, err
	}

	return &model.Chat{
		User1: *u1,
		User2: *u2,
	}, nil
}
