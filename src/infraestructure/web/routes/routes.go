package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/web/controllers"
)

func StartServer() {
	r := gin.Default()
	r.GET("/", controllers.CheckStatus)
	r.Run()
}
