package config

import (
	"encoding/json"
	"log"
	"os"
)

func Get() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	configuration := Config{}
	if err := json.NewDecoder(file).Decode(&configuration); err != nil {
		log.Fatal(err)
	}
	return configuration
}
