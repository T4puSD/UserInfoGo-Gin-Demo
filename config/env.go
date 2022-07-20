package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type EnvVariables struct {
	RunMode             string
	ServerPort          string
	MongodbUri          string
	MongodbDatabaseName string
	TrustedProxies      []string
	JwtSecretKey        string
}

var env *EnvVariables

func initEnvVariables() {
	if env != nil {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalln("Unable to load dotenv file")
	}

	env = &EnvVariables{
		RunMode:             os.Getenv("RUN_MODE"),
		ServerPort:          os.Getenv("SERVER_PORT"),
		MongodbUri:          os.Getenv("MONGODB_URI"),
		MongodbDatabaseName: os.Getenv("MONGODB_DATABASE"),
		TrustedProxies:      strings.Split(os.Getenv("TRUSTED_PROXIES"), ","),
		JwtSecretKey:        os.Getenv("JWT_SECRET_KEY"),
	}
}

func GetEnv() *EnvVariables {
	initEnvVariables()
	if env == nil {
		log.Fatalln("Env variables are not initialized! Please Call The InitEnvVariables Method First")
	}
	return env
}
