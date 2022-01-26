package main

import (
	"github.com/hoangggg5/learn-go/config"
	"github.com/hoangggg5/learn-go/models"
	"github.com/hoangggg5/learn-go/routes"
)

func main() {
	config.InitConfig()

	config.Database.AutoMigrate(models.User{})

	routes.InitRouter()
}
