package views

import (
	_ "embed"
	"github.com/constantincuy/go-gui/example/layout/ui/components"
	"github.com/constantincuy/go-gui/ui/common"
)

type MainView struct {
	core   common.Core
	layout common.State[common.FlexLayout]
}

func (view *MainView) Core() *common.Core {
	return &view.core
}

func (view *MainView) Mount() {
	view.layout = common.NewState(common.FlexCentered())
	view.layout.OnChange(func(newLayout common.FlexLayout) { view.Core().SetDisplayType(newLayout) })

	fc := view.Core().AddChild(components.NewFlexControl).(*components.FlexControl)

	fc.OnLayoutChange(func(justify, align common.FlexPosition) {
		cur := view.layout.Get()
		view.layout.SetState(cur.Justify(justify).Align(align))
	})
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {}

func NewMainView(core common.Core) common.Component {
	return &MainView{core: core}
}
