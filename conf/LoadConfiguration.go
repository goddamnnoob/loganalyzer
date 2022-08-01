package conf

import (
	"encoding/json"
	"os"
)

func GetConfiguration(confFilePath string) (*Config, *error) {
	var config Config
	configFile, err := os.Open(confFilePath)
	if err != nil {
		return &config, &err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config, nil
}
