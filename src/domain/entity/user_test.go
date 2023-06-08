package entity_test

import (
	"fmt"
	"testing"

	"github.com/nathanbahiadev/go-suppliers/src/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidEmails(t *testing.T) {
	emails := []entity.Email{
		{Value: "user@email.com"},
		{Value: "user.lastname@email.com"},
		{Value: "user-lastname@email.com.br"},
		{Value: "user_lastname@email.com.br"},
		{Value: "user_1@email.com"},
		{Value: "u1-lastname@email.com"},
		{Value: "u.user.lastname@email.com.br"},
	}

	for _, email := range emails {
		err := email.Validate()
		assert.Nil(t, err, fmt.Sprintf("%s should be valid", email.Value))
	}
}

func TestInvalidEmails(t *testing.T) {
	emails := []entity.Email{
		{Value: "useremail.com"},
		{Value: ".lastname@email.com"},
		{Value: "-lastname@email.com"},
		{Value: "_lastname@email.com"},
		{Value: "1@email.com"},
		{Value: "user@email"},
		{Value: "user@email.com.com"},
	}

	for _, email := range emails {
		err := email.Validate()
		assert.NotNil(t, err, fmt.Sprintf("%s should be invalid", email.Value))
	}
}

func TestValidPasswords(t *testing.T) {
	passwords := []entity.Password{
		{Value: "12345678Na!"},
		{Value: "abcdefgH!!2"},
		{Value: "1234efgh#Ça"},
		{Value: "a@a#a$2%Ava"},
		{Value: "@@@@@@@@1Na"},
	}

	for _, password := range passwords {
		err := password.Validate()
		assert.Nil(t, err, fmt.Sprintf("%s should be valid", password.Value))
	}
}

func TestInvalidPasswords(t *testing.T) {
	passwords := []entity.Password{
		{Value: "testtest"},
		{Value: "        "},
		{Value: "12345678"},
		{Value: "Naaaaaaa"},
		{Value: "aaaaaaa1"},
		{Value: "@@@@@@@@"},
		{Value: "N@@@@@@@"},
	}

	for _, password := range passwords {
		err := password.Validate()
		assert.NotNil(t, err, fmt.Sprintf("%s should be invalid", password.Value))
	}
}

func TestValidateUserWhenIDIsEmptyShouldRaiseAnError(t *testing.T) {
	user := entity.User{}

	err := user.Validate()
	assert.Equal(t, err.Error(), "user id cannot be blank")
}

func TestValidateUserWhenEmailIsInvalidShouldRaiseAnError(t *testing.T) {
	user := entity.User{
		ID: "1",
	}

	err := user.Validate()
	assert.Equal(t, err.Error(), "email invalid")
}

func TestValidateUserWhenNameIsEmptyShouldRaiseAnError(t *testing.T) {
	user := entity.User{
		ID:       "1",
		Email:    entity.Email{Value: "user@email.com"},
		Password: entity.Password{Value: "12345678Na!"},
	}

	err := user.Validate()
	assert.Equal(t, err.Error(), "user name cannot be blank")
}

func TestCreateUser(t *testing.T) {
	user := entity.User{
		Email:    entity.Email{Value: "user@email.com"},
		Name:     "User",
		Password: entity.Password{Value: "12345678Na!"},
	}
	err := user.Create()

	assert.Nil(t, err)
	assert.NotEqual(t, user.ID, "")
}
