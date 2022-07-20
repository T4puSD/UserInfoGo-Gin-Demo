package routes

import (
	"github.com/gin-gonic/gin"
	"user-info-service/controllers"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/users", controllers.CreateUser())
	route.GET("/users/:id", controllers.GetUser())
	route.PUT("/users/:id", controllers.UpdateUser())
	route.DELETE("/users/:id", controllers.DeleteUser())
}
