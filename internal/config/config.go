package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// ////////////////////////////////////////////////////

func Read() (Config, error) {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonBody, err := os.Open(jsonPath)
	if err != nil {
		return Config{}, err
	}
	defer jsonBody.Close()

	decoder := json.NewDecoder(jsonBody)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// ////////////////////////////////////////////////////

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

// ////////////////////////////////////////////////////

func write(cfg Config) error {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonBody, err := os.Create(jsonPath)
	if err != nil {
		return err
	}
	defer jsonBody.Close()

	encoder := json.NewEncoder(jsonBody)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

// ////////////////////////////////////////////////////

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	pathFile := fmt.Sprintf("%s/%s", homeDir, configFileName)
	return pathFile, nil
}
