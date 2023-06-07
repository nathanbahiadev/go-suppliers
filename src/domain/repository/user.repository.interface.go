package repository

import "github.com/nathanbahiadev/go-suppliers/src/domain/entity"

type UserRepositoryInterface interface {
	Create(DB interface{}, u *entity.User) error
	Update(DB interface{}, u *entity.User) error
	Find(DB interface{}, id string) (*entity.User, error)
	FindAll(DB interface{}, id string) ([]*entity.User, error)
}
