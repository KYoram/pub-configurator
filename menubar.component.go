package main

import (
	"github.com/AllenDang/giu"
	"github.com/LeRoiDesPoissons/pub-configurator/config"
)

func generateMenuBar() *giu.MenuBarWidget {
	return giu.MenuBar().Layout(
		giu.MenuItem("New").Shortcut("Ctrl+N").OnClick(func() {
			if Filthy {
				// handle filty
			}

			Config = config.GenerateConfig()
			ReflectConfig = Config
		}),
		giu.MenuItem("Open").Shortcut("Ctrl+O").OnClick(func() {
			if Filthy {
				// handle filthy
			}

			loadedConfig, err := config.OpenConfig()

			if err != nil {
				// handle err
			}

			Config = loadedConfig
			ReflectConfig = Config
		}),
		giu.MenuItem("Save").Shortcut("Ctrl+S"),
		giu.MenuItem("Save as..."),
	)

}
