package config

import (
	"encoding/json"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

const DefaultConfigName string = "pub.config"

type Config struct {
	Name string
}

func OpenConfig() (Config, error) {
	var message Config
	var err error
	var result string

	result, err = cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open A File",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Pub config file (*config.json)",
				Pattern:     "*.json",
			},
		},
		SelectedFileFilterIndex: 1,
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
