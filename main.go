package main

import (
	"github.com/AllenDang/giu"
	"github.com/LeRoiDesPoissons/pub-configurator/section"
)

var window = giu.SingleWindowWithMenuBar()

func loop() {
	window.Layout(
		menubar,
		giu.SplitLayout(
			giu.DirectionVertical,
			&pos,
			giu.Layout{
				giu.Label("Sections"),
				giu.Column(
					giu.Button("General").Disabled(true).OnClick(func() {
						section.SetSection(section.GeneralSection)
					}),
					giu.Button("Consumables").Disabled(true).OnClick(func() {
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
