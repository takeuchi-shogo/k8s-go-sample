package main

import (
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/config"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/database"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/routes"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/server"
)

func main() {

	config := config.NewConfig(".")
	fmt.Println(config)
	db := database.NewDB()
	routes := routes.NewRoutes(db)
	server := server.NewServer(routes)

	server.Run(":" + config.ServerPort)
}
