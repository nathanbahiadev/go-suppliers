package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type Company struct {
	ID   string
	User User
	CNPJ CNPJ
	Name string
}

func (c *Company) Validate() error {
	if c.ID == "" {
		return errors.New("company id cannot be blank")
	}

	if c.User.ID == "" {
		return errors.New("company user id cannot be blank")
	}

	if err := c.CNPJ.Validate(); err != nil {
		return err
	}

	if c.Name == "" {
		return errors.New("company name cannot be blank")
	}

	return nil
}

func (c *Company) Create() error {
	c.ID = uuid.NewString()
	return c.Validate()
}

type CNPJ struct {
	Value string
}

func (c *CNPJ) Validate() error {
	pattern := `^[\d]{2}.[\d]{3}.[\d]{3}\/[\d]{4}-[\d]{2}$`
	re, _ := regexp.Compile(pattern)

	if !re.MatchString(c.Value) {
		return errors.New("cnpj is invalid")
	}

	return nil
}
