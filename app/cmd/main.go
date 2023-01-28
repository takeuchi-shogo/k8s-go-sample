package main

import (
	"github.com/takeuchi-shogo/k8s-go-sample/config"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/database"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/routes"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/server"
)

func main() {

	// k8s.LeadConfigMap()
	config := config.NewConfig(".")
	db := database.NewDB(config)
	routes := routes.NewRoutes(db)
	server := server.NewServer(routes)

	server.Run(":" + config.ServerPort)
}
