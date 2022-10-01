package controllers

import (
	"assignment3/services"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

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

func Scheduler() {
	for {
		log.Println("Scheduler: ----- Change Status -----")

		newStatus := map[string]interface{}{
			"status": map[string]interface{}{
				"water": rand.Intn(100-1) + 1,
				"wind":  rand.Intn(100-1) + 1,
			},
		}
		content, err := json.Marshal(newStatus)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("models/status.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(15 * time.Second)
	}
}
