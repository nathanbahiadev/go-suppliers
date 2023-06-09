package entity

import (
	"errors"
	"regexp"
	"unicode"

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
	var hasUpper, hasLower, hasNumber, hasPunc bool

	if len(p.Value) < 8 {
		return errors.New("password length must be greater than or equal to 8")
	}

	for _, char := range p.Value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasPunc = true
		}
	}

	if hasUpper && hasLower && hasNumber && hasPunc {
		return nil
	}

	return errors.New("password must contain at least 1 uppercase letter, 1 lowercase letter, a symbol or punctuation and 1 number and its length must be greater than or equal to 8")
}
