package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/web/controllers"
)

func StartServer() {
	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	apiv1.POST("/users/", controllers.CreateUserController)

	r.Run()
}
