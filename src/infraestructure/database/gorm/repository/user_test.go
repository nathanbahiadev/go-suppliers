package repository_test

import (
	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/repository"
	"github.com/stretchr/testify/assert"
)

func (suite *DatabaseSuiteTest) TestCreateUser() {
	userRepository := repository.UserRepository{}

	user := entity.User{
		Email: entity.Email{Value: "user1@email.com"},
		Name:  "User",
	}
	user.Create()

	err := userRepository.Create(suite.DB, &user)
	assert.Nil(suite.T(), err)
}

func (suite *DatabaseSuiteTest) TestFindUser() {
	userRepository := repository.UserRepository{}

	userCreated := entity.User{
		Email: entity.Email{Value: "user@email.com"},
		Name:  "User",
	}
	userCreated.Create()

	err := userRepository.Create(suite.DB, &userCreated)
	assert.Nil(suite.T(), err)

	userFounded, err := userRepository.Find(suite.DB, userCreated.ID)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), userFounded.ID, userCreated.ID)
	assert.Equal(suite.T(), userFounded.Name, userCreated.Name)
	assert.Equal(suite.T(), userFounded.Email.Value, userCreated.Email.Value)
}

func (suite *DatabaseSuiteTest) TestFindUserShouldRaiseAnErrorWhenUserDoesNotExists() {
	userRepository := repository.UserRepository{}
	user, err := userRepository.Find(suite.DB, "1")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err.Error(), "record not found")
	assert.Equal(suite.T(), user.ID, "")
}
