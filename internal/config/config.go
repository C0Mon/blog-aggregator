package config

import (
	"encoding/json"
	"os"
)

const configFileName = "aggrigatorconfig.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	configPath := getConfigPath()
	data, err := os.ReadFile(configPath)
	check(err)
	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	check(err)
	return cfg
}

func (cfg *Config) SetUser(newUser string) {
	cfg.CurrentUserName = newUser
	data, err := json.Marshal(cfg)
	check(err)

	configPath := getConfigPath()
	os.WriteFile(configPath, data, os.FileMode(0777))
}

func getConfigPath() string {
	homePath, err := os.UserHomeDir()
	check(err)
	return homePath + "/" + configFileName
}
