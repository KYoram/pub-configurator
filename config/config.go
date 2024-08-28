package config

import (
	"encoding/json"
	"os"

	"github.com/ncruces/zenity"
)

const DefaultConfigName string = "pub.config"

type Config struct {
	Name string
}

func OpenConfig() (Config, error) {
	var message Config

	fileName, err := zenity.SelectFile(
		zenity.Filename(""),
		zenity.FileFilters{
			{
				Name:     "Pub config files (.json)",
				Patterns: []string{"*.json"},
			},
		})

	if err != nil {
		return Config{}, err
	}

	content, err := os.ReadFile(fileName)

	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(content, &message)

	return message, err
}

func GenerateConfig() Config {
	return Config{
		Name: "",
	}
}
