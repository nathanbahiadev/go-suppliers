package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	ID       string
	Email    Email
	Password Password
	Name     string
}

func (u *User) Validate() error {
	if u.ID == "" {
		return errors.New("user id cannot be blank")
	}

	if err := u.Email.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	if u.Name == "" {
		return errors.New("user name cannot be blank")
	}

	return nil
}

func (u *User) Create() error {
	u.ID = uuid.NewString()
	return u.Validate()
}

type Email struct {
	Value string
}

func (e *Email) Validate() error {
	pattern := `^[a-zA-Z]+[\w\-\.]*@[a-zA-Z]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2})?$`
	re, _ := regexp.Compile(pattern)

	if !re.MatchString(e.Value) {
		return errors.New("email invalid")
	}
	return nil
}

type Password struct {
	Value string
}

func (p *Password) Validate() error {
	return nil
}
