package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-info-service/config"
	"user-info-service/db"
	"user-info-service/routes"
)

func main() {
	gin.SetMode(config.GetEnv().RunMode)

	// initiating the router
	router := gin.Default()

	if err := router.SetTrustedProxies(config.GetEnv().TrustedProxies); err != nil {
		log.Fatalln("Unable to set Trusted Proxies", err)
	}

	// initiate database connection
	db.ConnectDB()

	// plugging different routes
	routes.UserRoutes(router)

	err := router.Run(":" + config.GetEnv().ServerPort)
	if err != nil {
		log.Fatal("Unable to start the Gin Server on Port. Err: ", err.Error())
	}
}
