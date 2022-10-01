package main

import (
	"assignment3/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	go func() {
		controllers.Scheduler()
	}()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/status", controllers.GetStatus)

	router.Run(fmt.Sprintf("localhost:%v", 9000))
}
