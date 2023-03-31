package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const url = "https://jsonplaceholder.typicode.com/posts"

type PostRequ struct {
	UserID      int    `json:"userId"`
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"statusWater"`
	StatusWind  string `json:"statusWind"`
}

type PostResp struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	for {
		rand.Seed(time.Now().UnixNano())

		userID := rand.Intn(100) + 1
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		statusWater := getStatusWater(water)
		statusWind := getStatusWind(wind)

		req := PostRequ{
			UserID:      userID,
			Water:       water,
			Wind:        wind,
			StatusWater: statusWater,
			StatusWind:  statusWind,
		}
		reqBody, err := json.Marshal(req)
		if err != nil {
			log.Fatalf("Error marshaling request body: %v", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Fatalf("Error performing post request: %v", err)
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		var postResp map[string]interface{}
		if err := json.Unmarshal(respBody, &postResp); err != nil {
			log.Fatalf("Error unmarshaling response body: %v", err)
		}

		fmt.Printf("Request :\n%s\n", reqBody)
		fmt.Printf("Response :\n{\n  \"water\": %d,\n  \"wind\": %d\n}\n", int(postResp["water"].(float64)), int(postResp["wind"].(float64)))
		fmt.Printf("Status water : %s\n", statusWater)
		fmt.Printf("Status wind : %s\n", statusWind)

		time.Sleep(15 * time.Second)
	}
}

func getStatusWater(water int) string {
	if water < 5 {
		return "aman"
	} else if water >= 6 && water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getStatusWind(wind int) string {
	if wind < 6 {
		return "aman"
	} else if wind >= 7 && wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
