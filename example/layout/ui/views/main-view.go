package views

import (
	_ "embed"
	"github.com/constantincuy/go-gui/example/layout/ui/components"
	"github.com/constantincuy/go-gui/ui/component"
)

type MainView struct {
	core   component.Core
	layout component.State[component.FlexLayout]
}

func (view *MainView) Core() *component.Core {
	return &view.core
}

func (view *MainView) Mount() {
	view.layout = component.NewState(component.FlexCentered())
	view.layout.OnChange(func(newLayout component.FlexLayout) { view.Core().SetDisplayType(newLayout) })

	fc := view.Core().AddChild(components.NewFlexControl).(*components.FlexControl)

	fc.OnLayoutChange(func(justify, align component.FlexPosition) {
		cur := view.layout.Get()
		view.layout.SetState(cur.Justify(justify).Align(align))
	})
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {}

func NewMainView(core component.Core) component.Component {
	return &MainView{core: core}
}
