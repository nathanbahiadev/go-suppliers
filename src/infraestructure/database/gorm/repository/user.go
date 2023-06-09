package repository

import (
	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(u *entity.User) error {
	return r.DB.Create(&model.UserModel{
		ID:       u.ID,
		Email:    u.Email.Value,
		Password: u.Password.Value,
		Name:     u.Name,
	}).Error
}

func (r *UserRepository) Update(u *entity.User) error {
	return nil
}

func (r *UserRepository) Find(id string) (*entity.User, error) {
	userModel := model.UserModel{}

	if err := r.DB.First(&userModel, "id = ?", id).Error; err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: entity.Email{Value: userModel.Email},
	}, nil
}
