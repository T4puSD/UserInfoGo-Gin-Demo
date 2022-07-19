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
}

var env *EnvVariables

func initEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Unable to load dotenv file")
	}

	if env != nil {
		return
	}

	env = &EnvVariables{
		RunMode:             os.Getenv("RUN_MODE"),
		ServerPort:          os.Getenv("SERVER_PORT"),
		MongodbUri:          os.Getenv("MONGODB_URI"),
		MongodbDatabaseName: os.Getenv("MONGODB_DATABASE"),
		TrustedProxies:      strings.Split(os.Getenv("TRUSTED_PROXIES"), ","),
	}
}

func GetEnv() *EnvVariables {
	initEnvVariables()
	if env == nil {
		log.Fatalln("Env variables are not initialized! Please Call The InitEnvVariables Method First")
	}
	return env
}
