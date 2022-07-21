package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"user-info-service/model"
	"user-info-service/requests"
	"user-info-service/responses"
	"user-info-service/services/authservice"
	"user-info-service/services/userService"
)

var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest requests.LoginRequest

		// binding json to struct
		if err := c.BindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
			return
		}

		// validation of the request
		if err := validate.Struct(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
			return
		}

		user, err := userService.GetByEmail(&loginRequest.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
			return
		}
		if verified := authservice.VerifyPassword(loginRequest.Password, user); !verified {
			c.JSON(http.StatusUnauthorized, getErrorUserResponse(http.StatusUnauthorized, errors.New("Username or password doesn't match")))
			return
		}

		token, refreshToken := authservice.GenerateTokens(user)
		c.JSON(http.StatusOK, responses.TokenResponse{
			Token:        *token,
			RefreshToken: *refreshToken,
		})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		// binding of request json body to user model
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
			return
		}

		insertResult, err := userService.RegisterNewUser(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, getErrorUserResponse(http.StatusBadRequest, err))
		}

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "Registration Successful!",
			Data:    map[string]interface{}{"data": insertResult},
		})
	}
}
