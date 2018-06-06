package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	FilePath = "../config.json"

	err := Load()
	if err != nil {
		t.Errorf("Reading configuration failed: %s\n", err)
	}
}

func TestLoadConfigWithWrongPath(t *testing.T) {
	FilePath = ""

	err := Load()
	if err == nil {
		t.Errorf("Error while reading config from wrong path: %s\n", err)
	}
}

func TestReadConfigFromJSON(t *testing.T) {
	configFilePath := "../config.json"

	err := readFromJSON(configFilePath)
	if err != nil {
		t.Errorf("Reading configuration failed: %s\n", err)
	}
}

func TestReadConfigFromJSONWithWrongPath(t *testing.T) {
	configFilePath := ""

	err := readFromJSON(configFilePath)
	if err == nil {
		t.Errorf("Error while reading config from wrong path: %s\n", err)
	}
}
