package conf

import (
	"encoding/json"
	"os"
)

func GetConfiguration() (*Config, *error) {
	var config Config
	configFile, err := os.Open("../config.json")
	if err != nil {
		return &config, &err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config, nil
}
