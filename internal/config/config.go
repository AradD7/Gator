package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl 				string `json:"db_url"`
	CurrentUserName 	string `json:"current_user_name"`
}



func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}


func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}


func write(cfg  Config) error {
	configFullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	return nil
}


func getConfigFilePath() (string, error) {
	configFilePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return configFilePath + "/" + configFileName, nil
}

