package main

import (
	"account-app/cmd/dao"
	"account-app/cmd/handler"
	"account-app/cmd/service"
	"account-app/internal/initializers"
)

func main() {

	config := initializers.LoadConfig("config.yml")
	//initialize connection to database
	db := initializers.Initialize(config)
	//repositories
	repositories := dao.InitRepositories(db)
	//services
	services := service.Init(repositories)
	//controllers
	handlers := handler.Init(services)
	handlers.RegisterRoutes(config)
}
