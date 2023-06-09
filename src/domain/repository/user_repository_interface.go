package repository

import "github.com/nathanbahiadev/go-suppliers/src/domain/entity"

type UserRepositoryInterface interface {
	Create(u *entity.User) error
	Update(u *entity.User) error
	Find(id string) (*entity.User, error)
}
