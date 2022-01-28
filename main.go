package main

import (
	"github.com/hoanggggg5/learn-go/config"
	"github.com/hoanggggg5/learn-go/models"
	route "github.com/hoanggggg5/learn-go/routes"
)

func main() {
	config.InitConfig()

	config.Database.AutoMigrate(models.User{})

	route.InitRouter()
}
