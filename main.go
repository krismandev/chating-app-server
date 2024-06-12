package main

import (
	"flag"
	"os"
	"server-chat/app"
	"server-chat/config"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func main() {

	var logFile *os.File
	LoadConfig("config/config.yml", config.Config.Log.FileNamePrefix, config.Config.Log.Level, logFile)

	db := config.NewDatabase()
	validator := validator.New()

	gin := gin.Default()

	app.InitApp(&app.Application{
		DB:            db,
		Validate:      validator,
		Gin:           gin,
		WsConnections: make([]*config.WebSocketConnection, 0),
	})

	gin.Run(":" + config.Config.Port)
	logrus.Infof("App Running... Url: %v Port: %v", config.Config.AppUrl, config.Config.Port)

}

// LoadConfig to load config from config.yml file
func LoadConfig(configPath, logFileName, logLevel string, logFile *os.File) {
	configFile := flag.String("config", configPath, "main configuration file")
	config.InitConfig(configFile)
	flag.Parse()
	logrus.Infof("Reads configuration from %s", *configFile)
	config.InitLog(config.Config.Log.FileNamePrefix, logLevel, logFile)
}
