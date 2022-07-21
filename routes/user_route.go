package routes

import (
	"github.com/gin-gonic/gin"
	"user-info-service/controllers"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	userRouter.GET("", controllers.GetAllUser())
	userRouter.GET("/:id", controllers.GetUser())
	userRouter.PUT("/:id", controllers.UpdateUser())
	userRouter.DELETE("/:id", controllers.DeleteUser())
}
