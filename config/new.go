package config

import (
	"encoding/json"
	"log"
	"os"
)

func New() *Config {

	file, err := os.ReadFile("./js-strings.json")
	if err != nil {
		log.Fatal("can't find \"js-strings.json\" file: ", err)
	}

	c := new(Config)

	err = json.Unmarshal(file, c)
	if err != nil {
		log.Fatal("can't decode json file: ", err)
	}

	return c
}
