package main

import (
	"github.com/AllenDang/giu"
	"github.com/LeRoiDesPoissons/pub-configurator/config"
)

var menubar = giu.MenuBar().Layout(
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

		result, err := config.OpenConfig()

		if err != nil {
			// handle err
		}

		Config = result
		ReflectConfig = Config
	}),
	giu.MenuItem("Save").Shortcut("Ctrl+S"),
	giu.MenuItem("Save as..."),
)
