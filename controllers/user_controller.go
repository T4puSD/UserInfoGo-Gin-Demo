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
			c.JSON(http.StatusInternalServerError, responses.GetErrorUserResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": user},
		})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.GetErrorUserResponse(http.StatusBadRequest, err))
			return
		}

		if err := userService.Update(id, &user); err != nil {
			c.JSON(http.StatusBadRequest, responses.GetErrorUserResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "updated"},
		})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := userService.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GetErrorUserResponse(http.StatusInternalServerError, err))
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "deleted",
			Data:    map[string]interface{}{"data": "deleted"},
		})
	}
}

func GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GetErrorUserResponse(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": users},
		})
	}
}
