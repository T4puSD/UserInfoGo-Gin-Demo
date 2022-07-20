package authservice

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
	"user-info-service/config"
	"user-info-service/model"
)

type SignedDetails struct {
	Email string
	Name  string
	jwt.StandardClaims
}

func GenerateTokens(user *model.User) (token string, refreshToken string) {
	claims := SignedDetails{
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "rk",
			Subject:   user.Id.String(),
		},
	}

	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "rk",
			Subject:   user.Id.String(),
		},
	}

	return generateToken(claims), generateToken(refreshClaims)
}

func generateToken(claims SignedDetails) string {
	token, err := jwt.NewWithClaims(jwt.SigningMethod(jwt.SigningMethodEdDSA), claims).SignedString(config.GetEnv().JwtSecretKey)
	if err != nil {
		log.Panicln("Error during creating token", err)
	}

	return token
}
