package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	fpath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error reading home dir: %v", err)
	}

	data, err := os.ReadFile(fpath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, fmt.Errorf("unmarshal error: %v", err)
	}

	return conf, nil
}

func (cfg *Config) SetUser(name string) (Config, error) {

	if name == "" {
		return Config{}, errors.New("empty names are not allowed")
	}

	fpath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(fpath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, fmt.Errorf("unmarshal error: %v", err)
	}

	conf.CurrentUserName = name

	if err := write(conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error reading home dir: %v", err)
	}

	fpath := homeDir + "/" + configFileName

	return fpath, nil
}

func write(cfg Config) error {
	newConf, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	fpath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(fpath, newConf, os.ModeAppend); err != nil {
		return err
	}

	return nil
}
