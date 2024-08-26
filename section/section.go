package section

import (
	"github.com/AllenDang/giu"
	"github.com/KYoram/pub-configurator/consumables"
	"github.com/KYoram/pub-configurator/general"
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
	layout := giu.Layout{
		giu.Label("Unknown section"),
	}

	switch SelectedSection {
	case GeneralSection:
		layout = general.GeneralLayout
	case ConsumablesSection:
		layout = consumables.ConsumablesLayout
	}

	return layout
}
