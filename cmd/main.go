package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kirtikamal/sms-verify/api"
)

func main() {
	router := gin.Default()

	//initialize the config file
	app := api.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8000")
}
