package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanbahiadev/go-suppliers/src/domain/dto"
	"github.com/nathanbahiadev/go-suppliers/src/domain/service"
	database "github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm/repository"
)

func CreateUserController(c *gin.Context) {
	userCreateDto := &dto.UserCreateDto{}
	userRepository := &repository.UserRepository{DB: database.DB}

	if err := c.ShouldBindJSON(userCreateDto); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid data",
			"error":   err.Error(),
		})
		return
	}

	if userCreated, err := service.CreateUserService(userRepository, *userCreateDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on create user",
			"error":   err.Error(),
		})
		return
	} else {
		userResponseDto := dto.UserResponseDtoFromEntity(userCreated)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created",
			"detail":  userResponseDto,
		})
		return
	}
}
