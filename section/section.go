package section

import (
	"github.com/AllenDang/giu"
	"github.com/LeRoiDesPoissons/pub-configurator/consumables"
	"github.com/LeRoiDesPoissons/pub-configurator/general"
)

type Section int

const (
	GeneralSection Section = iota
	ConsumablesSection
)

var SelectedSection Section = GeneralSection

func SetSection(s Section) {
	if s != SelectedSection {
		SelectedSection = s
	}
}

func SelectedLayout() giu.Layout {
	var layout giu.Layout

	switch SelectedSection {
	case GeneralSection:
		layout = general.GeneralLayout
	case ConsumablesSection:
		layout = consumables.ConsumablesLayout
	default:
		layout = giu.Layout{
			giu.Label("Create a new or open a config file to start"),
		}
	}

	return layout
}
