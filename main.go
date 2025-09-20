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
	// Database Initialization
	config.ConnectDatabase()
	// test Controller 
	test_repo := repository.NewTestRepository(config.DB)
	test_service := services.NewTestService(test_repo)
	test_controller := controllers.Init(test_service)
	// auth Controllter
	auth_repo := repository.NewAuthRepository(config.DB)
	auth_service := services.NewAuthService(auth_repo)
	auth_controller := controllers.NewAuthService(auth_service)

	// routers
	router := gin.Default()
	routes.TestRoute(router, test_controller,auth_controller)
	// server
	router.Run(":8080")
}