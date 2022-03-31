package main

import (
	"fmt"
	"github.com/Toflex/Wallet-API/app"
	"github.com/Toflex/Wallet-API/configs"
	"github.com/Toflex/Wallet-API/databases"
	"github.com/Toflex/Wallet-API/docs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)


}


func main() {
	APP:=&app.App{}
	APP.Config = configs.Configuration()
	APP.Router = gin.Default()

	docs.SwaggerInfo.Title="Wallet API"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Description="Wallet management API"
	docs.SwaggerInfo.Version="version 1"

	database:=&databases.Database{}
	database.Config = APP.Config
	database.InitSQLDB()
	database.InitRedisDB()

	APP.DB = database.DB
	APP.RDB = database.RDB

	// Set up the routes
	APP.SetupRoutes()

	log.Info(fmt.Sprintf("Server started %s:%s", APP.Config.ServerHost, APP.Config.ServerPort))
	APP.Router.Run(fmt.Sprintf("%s:%s", APP.Config.ServerHost, APP.Config.ServerPort))
}

