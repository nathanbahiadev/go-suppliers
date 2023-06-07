package repository

import (
	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/model"
	"gorm.io/gorm"
)

type CompanyRepository struct{}

func (r *CompanyRepository) Create(DB *gorm.DB, c *entity.Company) error {
	return DB.Create(&model.CompanyModel{
		ID:     c.ID,
		Name:   c.Name,
		CNPJ:   c.CNPJ.Value,
		UserID: c.User.ID,
	}).Error
}

func (r *CompanyRepository) Update(c *entity.Company) error {
	return nil
}

func (r *CompanyRepository) Find(DB *gorm.DB, id string) (*entity.Company, error) {
	companyModel := model.CompanyModel{}

	if err := DB.Preload("UserModel").First(&companyModel, "id = ?", id).Error; err != nil {
		return &entity.Company{}, err
	}

	return &entity.Company{
		ID:   companyModel.ID,
		Name: companyModel.Name,
		CNPJ: entity.CNPJ{Value: companyModel.CNPJ},
		User: entity.User{
			ID:    companyModel.UserModel.ID,
			Name:  companyModel.UserModel.Name,
			Email: entity.Email{Value: companyModel.UserModel.Email},
		},
	}, nil
}

func (r *CompanyRepository) FindAll(id string) ([]*entity.Company, error) {
	return nil, nil
}
