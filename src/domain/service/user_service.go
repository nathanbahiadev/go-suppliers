package service

import (
	"github.com/nathanbahiadev/go-suppliers/src/domain/dto"
	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/nathanbahiadev/go-suppliers/src/domain/repository"
)

func CreateUserService(
	userRepository repository.UserRepositoryInterface,
	userCreateDto dto.UserCreateDto,
) (*entity.User, error) {
	user := &entity.User{
		Name:     userCreateDto.Name,
		Email:    entity.Email{Value: userCreateDto.Email},
		Password: entity.Password{Value: userCreateDto.Password},
	}

	if err := user.Create(); err != nil {
		return nil, err
	}

	if err := userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
