package routes

import (
	"github.com/gin-gonic/gin"
	"user-info-service/controllers"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/login", controllers.Login())
	router.POST("/register", controllers.Register())
}
