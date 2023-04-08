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

func posting() {
	var water, wind int
	var statusWater, statusWind string

	data := map[string]interface{}{
		"water": rand.Intn(100),
		"wind":  rand.Intn(100),
	}

	client := &http.Client{}
	requestJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	water = data["water"].(int)
	wind = data["wind"].(int)

	if water < 5 {
		statusWater = "aman"
	} else if water >= 6 && water < 9 {
		statusWater = "siaga"
	} else {
		statusWater = "bahaya"
	}

	if wind < 6 {
		statusWind = "aman"
	} else if wind >= 7 && wind < 16 {
		statusWind = "siaga"
	} else {
		statusWind = "bahaya"
	}

	log.Println(string(body))
	fmt.Println("status water : ", statusWater)
	fmt.Println("status wind : ", statusWind)
}

func main() {
	limit := time.Now().Add(time.Minute * 1)
	for i := 0; i < 100; i++ {
		if time.Now().After(limit) {
			break
		}
		posting()
		time.Sleep(time.Second * 15)

	}
}
