package views

import (
	"github.com/constantincuy/go-gui/example/counter/ui/components"
	"github.com/constantincuy/go-gui/ui/component"
)

type MainView struct {
	core component.Core
}

func (view *MainView) Core() *component.Core {
	return &view.core
}

func (view *MainView) Mount() {
	view.Core().SetDisplayType(component.FlexCentered().UseGap(5).UseDirection(component.FlexColumn))
	for i := 0; i < 5; i++ {
		view.Core().AddChild(components.NewCounter)
	}
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {}

func NewMainView(core component.Core) component.Component {
	return &MainView{core: core}
}
