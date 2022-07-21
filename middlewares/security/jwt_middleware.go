package security

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-info-service/responses"
	"user-info-service/services/authservice"
	"user-info-service/utils"
)

const (
	BEARER_KEYWORD           = "Bearer"
	AUTHORIZATION_HEADER_KEY = "Authorization"
)

func ValidateAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromRequest(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.GetErrorUserResponse(http.StatusUnauthorized, err))
			c.Abort()
			return
		}

		_, err = authservice.VerifyToken(*token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GetErrorUserResponse(http.StatusUnauthorized, err))
			c.Abort()
			return
		}
		c.Next()
	}
}

func extractTokenFromRequest(c *gin.Context) (*string, error) {
	noAuthorizationHeaderError := errors.New("No Authorization Header Provided!")

	bearer := c.GetHeader(AUTHORIZATION_HEADER_KEY)
	if utils.IsBlank(bearer) {
		return nil, noAuthorizationHeaderError
	}

	token := bearer[len(BEARER_KEYWORD)+1:]
	if utils.IsBlank(token) {
		return nil, noAuthorizationHeaderError
	}

	return &token, nil
}
