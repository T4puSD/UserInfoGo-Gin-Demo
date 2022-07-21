package authservice

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"user-info-service/config"
	"user-info-service/model"
)

const (
	TOKEN_VALIDITY_TIME         = 1 * time.Hour
	REFRESH_TOKEN_VALIDITY_TIME = 48 * time.Hour
)

type SignedDetails struct {
	Email string
	Name  string
	jwt.StandardClaims
}

func GenerateTokens(user *model.User) (token *string, refreshToken *string) {
	claims := &SignedDetails{
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TOKEN_VALIDITY_TIME).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "rk",
			Subject:   user.Id.Hex(),
		},
	}

	refreshClaims := &SignedDetails{
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_VALIDITY_TIME).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "rk",
			Subject:   user.Id.Hex(),
		},
	}

	return generateToken(claims), generateToken(refreshClaims)
}

func generateToken(claims *SignedDetails) *string {
	token, err := jwt.NewWithClaims(jwt.SigningMethod(jwt.SigningMethodHS256), claims).SignedString([]byte(config.GetEnv().JwtSecretKey))
	if err != nil {
		log.Panicln("Error during creating token: ", err)
	}

	return &token
}

func GetHashedPassword(rawPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

// VerifyPassword return true if rawPassword is correct otherwise return false
func VerifyPassword(rawPassword string, user *model.User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword))
	return err == nil
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}
