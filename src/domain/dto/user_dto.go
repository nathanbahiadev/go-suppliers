package dto

import "github.com/nathanbahiadev/go-suppliers/src/domain/entity"

type UserCreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
}

type UserResponseDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserResponseDtoFromEntity(user *entity.User) UserResponseDto {
	return UserResponseDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email.Value,
	}
}
