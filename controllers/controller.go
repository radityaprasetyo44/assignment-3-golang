package controllers

import (
	"assignment3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	filePath, data, err := services.GetStatus()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "failed to get status",
		})
		return
	}

	c.HTML(http.StatusOK, filePath, data)
}