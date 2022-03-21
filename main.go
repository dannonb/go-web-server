package main

import (
	"github.com/dannonb/web-server/config"
	"github.com/dannonb/web-server/routes"
	"github.com/gin-gonic/gin"

)

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Hello World",
		})
	})
	
	// connect database
	config.ConnectDB()

	// routes
	routes.UserRoute(router)

	// start server
	router.Run()

	
}