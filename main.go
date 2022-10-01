package main

import (
	"assignment3/controllers"
	"assignment3/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	go func() {
		services.Scheduler()
	}()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/status", controllers.GetStatus)

	router.Run(fmt.Sprintf("localhost:%v", 9000))
}
