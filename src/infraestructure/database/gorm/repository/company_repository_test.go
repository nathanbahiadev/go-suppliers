package repository_test

import (
	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/model"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/repository"
	"github.com/stretchr/testify/assert"
)

func (suite *DatabaseSuiteTest) TestCreateCompany() {
	companyRepository := repository.CompanyRepository{DB: suite.DB}
	userRepository := repository.UserRepository{DB: suite.DB}

	var err error

	user := entity.User{
		Email: entity.Email{Value: "user1@email.com"},
		Name:  "User",
	}
	user.Create()

	err = userRepository.Create(&user)
	assert.Nil(suite.T(), err)

	companyCreated := entity.Company{
		Name: "Company",
		CNPJ: entity.CNPJ{
			Value: "00.000.000/0001-00",
		},
		User: user,
	}
	companyCreated.Create()

	err = companyRepository.Create(&companyCreated)
	assert.Nil(suite.T(), err)

	companyModel := &model.CompanyModel{}
	suite.DB.First(companyModel, "id = ?", companyCreated.ID)

	assert.Equal(suite.T(), companyModel.ID, companyCreated.ID)
	assert.Equal(suite.T(), companyModel.Name, companyCreated.Name)
	assert.Equal(suite.T(), companyModel.CNPJ, companyCreated.CNPJ.Value)
	assert.Equal(suite.T(), companyModel.UserID, companyCreated.User.ID)
}

func (suite *DatabaseSuiteTest) TestFindCompany() {
	companyRepository := repository.CompanyRepository{DB: suite.DB}
	userRepository := repository.UserRepository{DB: suite.DB}

	var err error

	user := entity.User{
		Email: entity.Email{Value: "user1@email.com"},
		Name:  "User",
	}
	user.Create()

	err = userRepository.Create(&user)
	assert.Nil(suite.T(), err)

	companyCreated := entity.Company{
		Name: "Company",
		CNPJ: entity.CNPJ{
			Value: "00.000.000/0001-00",
		},
		User: user,
	}
	companyCreated.Create()

	err = companyRepository.Create(&companyCreated)
	assert.Nil(suite.T(), err)

	companyModel := &model.CompanyModel{}
	suite.DB.First(companyModel, "id = ?", companyCreated.ID)

	companyFounded, err := companyRepository.Find(companyCreated.ID)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), companyModel.ID, companyFounded.ID, "company ids aren't equal")
	assert.Equal(suite.T(), companyModel.Name, companyFounded.Name, "company names aren't equal")
	assert.Equal(suite.T(), companyModel.CNPJ, companyFounded.CNPJ.Value, "company cnpjs aren't equal")
	assert.Equal(suite.T(), companyModel.UserID, companyFounded.User.ID, "company user ids aren't equal")
}

func (suite *DatabaseSuiteTest) TestFindCompanyShouldRaiseAnErrorWhenCompanyDoesNotExists() {
	companyRepository := repository.CompanyRepository{DB: suite.DB}
	company, err := companyRepository.Find("1")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err.Error(), "record not found")
	assert.Equal(suite.T(), company.ID, "")
}
