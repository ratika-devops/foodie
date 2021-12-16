package cache

import (
	"encoding/json"
	"foodies/gosrc/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var store struct {
	data []model.FoodTruck
	mu   sync.RWMutex
}

func InitCache() {
	err := loadCacheFromFile()
	if err != nil {
		err := LoadCacheFromURL()
		if err != nil {
			log.Fatal("Failed to initalize cache. Exiting")
		}
	}
}

func loadCacheFromFile() (err error) {
	file := "../data/foodtrrucks.json"
	log.Println("Starting to load cache from file:  ", file)
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Println("Failed to open cache file: "+file+" Error: ", err)
		return
	}
	log.Println("Successfully opened cache file: "+file+" Error: ", err)
	defer jsonFile.Close()

	fileInBytes, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Println("Failed to read cache file: "+file+" Error: ", err)
		return
	}

	store.mu.Lock()
	defer store.mu.Unlock()
	err = json.Unmarshal(fileInBytes, &store.data)
	if err != nil {
		log.Println("Failed to parse data from file: "+file+" Error: ", err)
		return
	}

	log.Println("Successfully initialized cache from file:  ", file)
	return
}

func LoadCacheFromURL() (err error) {
	URL := "https://data.sfgov.org/api/id/rqzj-sfat.json"

	log.Println("Trying cache URL: " + URL)
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("Failed to access URL: "+URL+" Error: ", err)
		return
	}

	defer resp.Body.Close()

	log.Println("Trying to read response from URL: " + URL)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response from URL: "+URL+" Error: ", err)
		return
	}

	log.Println("Trying to parse response from URL: " + URL)
	store.mu.Lock()
	defer store.mu.Unlock()
	err = json.Unmarshal(body, &store.data)
	if err != nil {
		log.Println("Failed to parse response from URL: "+URL+" Error: ", err)
		return
	}

	log.Println("Successfully initialized cache from URL: ", URL)
	return
}
