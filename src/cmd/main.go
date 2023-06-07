package main

import (
	"fmt"

	"github.com/google/uuid"
	database "github.com/nathanbahiadev/go-suppliers/src/infraestructure/database/gorm"
	"github.com/nathanbahiadev/go-suppliers/src/infraestructure/web/routes"
)

func main() {
	database.StartDatabase()

	u := uuid.New().String()
	fmt.Print(u)

	routes.StartServer()
}
