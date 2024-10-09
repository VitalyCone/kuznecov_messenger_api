package dtos

import "github.com/VitalyCone/kuznecov_messenger_api/internal/app/model"

type CreateUserDto struct {
	Username string `json:"username"`
}

func (u *CreateUserDto) CreateUserDtoToModel() *model.User {
	return &model.User{
		Username: u.Username,
	}
}

type ModifyUserDto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (u *ModifyUserDto) ModifyUserDtoToModel() *model.User {
	return &model.User{
		ID:       u.ID,
		Username: u.Username,
	}
}
