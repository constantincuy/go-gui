package views

import (
	_ "embed"
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
)

type MainView struct {
	core common.Core
}

func (view *MainView) Core() *common.Core {
	return &view.core
}

func (view *MainView) Mount() {
	view.Core().SetDisplayType(common.FlexCentered().UseGap(20).UseDirection(common.FlexColumn))
	headline := view.Core().AddChild(component.NewText).(*component.Text)
	headline.SetFontSize(30)
	headline.SetLineHeight(30)
	headline.SetText("Dynamic buttons")
	addBtn := view.Core().AddChild(component.NewButton).(*component.Button)
	addBtn.SetText("Add New Button")
	addBtn.OnClick(func() {
		view.createRemoveButton(len(view.Core().Children()) - 2)
	})
	for i := 0; i < 5; i++ {
		view.createRemoveButton(i)
	}
}

func (view *MainView) createRemoveButton(i int) {
	comp := view.Core().AddChild(component.NewButton)
	btn := comp.(*component.Button)
	btn.SetText(fmt.Sprintf("Remove Me %d", i+1))
	btn.OnClick(func() {
		view.Core().RemoveChild(&comp)
	})
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {}

func NewMainView(core common.Core) common.Component {
	return &MainView{core: core}
}
