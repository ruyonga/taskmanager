package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server, mongoDBHost, DBUser, DBPwd, Database string
}

var AppConfig configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	defer file.Close()

	decorder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decorder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
