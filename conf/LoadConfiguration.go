package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfiguration() (*Config, *error) {
	confFilePath := "config.json"
	var config Config
	configFile, err := os.Open(confFilePath)
	if err != nil {
		fmt.Println(err)
		return &config, &err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return &config, nil
}
