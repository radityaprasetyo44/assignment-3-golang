package services

import (
	"assignment3/helpers"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStatus() (filePath string, data gin.H, err error) {
	content, err := ioutil.ReadFile("models/status.json")
	if err != nil {
		return "", nil, errors.New("failed to read file")
	}

	result := map[string]interface{}{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return "", nil, errors.New("failed to unmarshal")
	}

	helpers := helpers.StatusMapper(result["status"].(map[string]interface{}))

	return "status.html", gin.H{
		"water":        result["status"].(map[string]interface{})["water"],
		"wind":         result["status"].(map[string]interface{})["wind"],
		"wind_status":  helpers["wind_status"],
		"water_status": helpers["water_status"],
	}, nil
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
