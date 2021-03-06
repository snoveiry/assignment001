package main

import (
	"os"

	"github.com/snoveiry/assignment001"
	"github.com/snoveiry/assignment001/docs"
	"github.com/snoveiry/assignment001/logger"
)

// @title Assignment001 API
// @version 1.0.0
// @description This is the documentation for the Assignment001 API.

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {

	as := assignment001.New()

	as.Config.Port = os.Getenv("PORT")
	logger.Info(nil, as.Config.Port)

	as.Config.BaseURL = os.Getenv("BASE_URL")
	logger.Info(nil, as.Config.BaseURL)

	docs.SwaggerInfo.Version = "1.0.0"
	
	as.Run()
}
