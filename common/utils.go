package common

import (
	"os"
	"log"
	"encoding/json"
)

type configuration struct {
	Email string `json: "email"`
	Password string `json: "password"`
	Oemail string `json: "oemail"`
}

var AppConfig configuration

func InitConfig()  {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
