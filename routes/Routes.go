package routes

import (
	"github.com/Vis7044/GinCrud2/controllers"
	"github.com/Vis7044/GinCrud2/middleware"
	"github.com/gin-gonic/gin"
)

func TestRoute(router *gin.Engine, testcontroller *controllers.Testcontroller, authcontroller *controllers.AuthController) {
	test := router.Group("/test") 
	{
		test.GET("/",testcontroller.GetTest)
		test.POST("/",testcontroller.CreateTest)
		test.GET("/:id",testcontroller.GetOne)
		test.PUT("/:id",testcontroller.UpdateTest)
		test.DELETE("/:id",testcontroller.DeleteTest)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register",authcontroller.Register)
		auth.POST("/login",authcontroller.Login)
		auth.GET("/profile", middleware.AuthMiddleware(), authcontroller.Profile)
	}
}