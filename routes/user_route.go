package routes

import (
	"github.com/gin-gonic/gin"
	"user-info-service/controllers"
)

func UserRoutes(route *gin.Engine) {
	route.GET("/users/:id", controllers.GetUser())
	route.POST("/users", controllers.CreateUser())
	route.PUT("/users/:id", controllers.UpdateUser())
}
