package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-info-service/model"
	"user-info-service/responses"
	"user-info-service/services/userService"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := userService.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, getErrorUserResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": user},
		})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		// validation of request json body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
			return
		}

		err := userService.Save(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
		}
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
			return
		}

		if err := userService.Update(id, &user); err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "updated"},
		})
	}
}

func getErrorUserResponse(statusCode int, err error) *responses.UserResponse {
	return &responses.UserResponse{
		Status:  statusCode,
		Message: "error",
		Data:    map[string]interface{}{"data": err.Error()},
	}
}
