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
	view.Core().SetDisplayType(component.FlexLayout{
		JustifyContent: component.FlexCenter,
		AlignItems:     component.FlexCenter,
		Gap:            5,
	})
	for i := 0; i < 5; i++ {
		ex := view.Core().AddChild(components.NewCounter)
		ex.Core().SetPositionXY(6+((i%10)*126), 6+((i/10)*41))
	}
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {}

func NewMainView(core component.Core) component.Component {
	return &MainView{core: core}
}
