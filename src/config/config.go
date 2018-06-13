package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"io/ioutil"
	"log"
)

// Config is variable for config
var (
	Config   Configuration
	FilePath = "config.json"
)

// Configuration is a singleton object for application config
type Configuration struct {
	ListenURL   string        `json:"ListenURL"`
	LogFilePath string        `json:"LogFilePath"`
	Database    database.Info `json:"Database"`
}

// Load loads config once
func Load() error {
	err := readFromJSON(FilePath)
	if err != nil {
		return errors.New("configuration not found. Please specify configuration")
	}

	return nil
}

// readFromJSON reads config data from JSON-file
func readFromJSON(configFilePath string) error {
	log.Printf("Looking for JSON config file (%s)", configFilePath)

	contents, err := ioutil.ReadFile(configFilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
	} else {
		log.Printf("Configuration has been read from JSON (%s) successfully\n", configFilePath)
	}

	return err
}
