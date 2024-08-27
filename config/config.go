package config

import (
	"encoding/json"

	"github.com/ncruces/zenity"
)

const DefaultConfigName string = "pub.config"

type Config struct {
	Name string
}

func OpenConfig() (Config, error) {
	var message Config

	result, err := zenity.SelectFile(
		zenity.Filename(""),
		zenity.FileFilters{
			{
				Name:     "Pub config files (.json)",
				Patterns: []string{"*.json"},
				CaseFold: false,
			},
		})

	if err == nil {
		err = json.Unmarshal([]byte(result), &message)
	}

	return message, err
}

func GenerateConfig() Config {
	return Config{
		Name: "",
	}
}
