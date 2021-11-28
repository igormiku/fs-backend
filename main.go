package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"toggle_backend/api"
	"toggle_backend/config"
	"toggle_backend/services"
	"toggle_backend/storage"
)

var cfg config.Config

func readConfig(cfg *config.Config) {
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}

func main() {

	readConfig(&cfg)

	db := storage.InitializeMockDB()
	//init FileDB or MockDB
	if cfg.DBType == "FileDB" {
		db = storage.InitializeBuntDB(cfg.DBFile)
	}

	defer db.Close()

	//create featureToggle service
	featureToggleService := services.NewFeatureToggleService(db)

	//TODO: interface API too, or add more services as parameters going forward
	api := api.NewAPI(featureToggleService)

	if cfg.APIBase != "" {
		api.SetupAPI(cfg.APIBase)
	}

	log.Fatal(http.ListenAndServe(cfg.Listen, api.Router))

	//TODO make separate Server package to create and run listener

}
