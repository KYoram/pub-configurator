package main

import (
	"fmt"

	"github.com/AllenDang/giu"
	"github.com/LeRoiDesPoissons/pub-configurator/config"
	"github.com/LeRoiDesPoissons/pub-configurator/section"
)

var (
	pos           float32       = 200
	Config        config.Config = config.Config{}
	ReflectConfig config.Config = Config
	Filthy        bool          = false
)

func loop() {
	giu.SingleWindow().Layout(
		giu.SplitLayout(
			giu.DirectionVertical,
			&pos,
			giu.Layout{
				giu.Label("File"),
				giu.PopupModal("select").Layout(
					giu.Label("Select file"),
					giu.Row(
						giu.Button("Open").OnClick(func() { giu.CloseCurrentPopup() }),
						giu.Button("Cancel").OnClick(func() { giu.CloseCurrentPopup() }),
					),
				),
				giu.Column(
					giu.Row(
						giu.Button("New").OnClick(func() {
							if Filthy {
								// handle filty
							} else {
								Config = config.GenerateConfig()
								ReflectConfig = Config
							}
						}),
						giu.Button("Open").OnClick(func() {
							if Filthy {
								// handle filty
							}

							giu.OpenPopup("select")

							fmt.Println("ja")

							// config, err := config.OpenConfig()

							// if err != nil {
							// 	// handle error
							// } else {
							// 	Config = config
							// 	ReflectConfig = Config
							// }
						}),
					),
					giu.Row(
						giu.Button("Save"),
						giu.Button("Save as..."),
					),
				),
				giu.Separator(),
				giu.Label("Sections"),
				giu.Column(
					giu.Button("General").OnClick(func() {
						section.SetSection(section.GeneralSection)
					}),
					giu.Button("Consumables").OnClick(func() {
						section.SetSection(section.ConsumablesSection)
					}),
				),
			},
			section.SelectedLayout(),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Pub Configurator", 800, 600, 0)
	wnd.Run(loop)
}
