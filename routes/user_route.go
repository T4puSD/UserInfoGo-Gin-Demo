package routes

import (
	"github.com/gin-gonic/gin"
	"user-info-service/controllers"
	"user-info-service/middlewares/security"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	userRouter.Use(security.ValidateAuthentication())
	userRouter.GET("", controllers.GetAllUser())
	userRouter.GET("/:id", controllers.GetUser())
	userRouter.PUT("/:id", controllers.UpdateUser())
	userRouter.DELETE("/:id", controllers.DeleteUser())
}
