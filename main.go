package main

import (
	"github.com/Vis7044/GinCrud2/config"
	"github.com/Vis7044/GinCrud2/controllers"
	"github.com/Vis7044/GinCrud2/repository"
	"github.com/Vis7044/GinCrud2/routes"
	"github.com/Vis7044/GinCrud2/services"
	"github.com/gin-gonic/gin"
)

func main() {
	
	config.ConnectDatabase()
	
	repo := repository.NewTestRepository(config.DB)
	service := services.NewTestService(repo)
	controller := controllers.Init(service)

	router := gin.Default()
	routes.TestRoute(router, controller)
	router.Run(":8080")
}