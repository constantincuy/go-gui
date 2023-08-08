package components

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
)

type LayoutListener func(justify, align common.FlexPosition)

type FlexControl struct {
	core     common.Core
	listener func(justify, align common.FlexPosition)
}

func (f *FlexControl) Core() *common.Core {
	return &f.core
}

func (f *FlexControl) Mount() {
	f.Core().SetDisplayType(common.FlexCentered().UseGap(15).UseDirection(common.FlexColumn))
	headline := f.Core().AddChild(component.NewText).(*component.Text)
	headline.SetFontSize(30)
	headline.SetText("Flex Layout")

	grid := f.Core().AddChild(component.NewGridContainer(3, 5)).(*component.Container)
	size := grid.Core().GetSize()
	size.Width = 375
	grid.Core().SetSize(size)
	options := []common.FlexPosition{common.FlexStart, common.FlexCenter, common.FlexEnd}
	labels := []string{"Top Left", "Top Center", "Top Right", "Center Left", "Center Center", "Center Right", "Bottom Left", "Bottom Center", "Bottom Right"}

	for i, label := range labels {
		col := i % 3
		row := i / 3
		btn := grid.Add(component.NewButton).(*component.Button)
		btn.SetText(label)
		btn.OnClick(func() {
			justify := options[col]
			align := options[row]
			f.listener(justify, align)
		})
	}
}

func (f *FlexControl) OnLayoutChange(l LayoutListener) {
	f.listener = l
}

func (f *FlexControl) Update() {}

func (f *FlexControl) Destroy() {}

func NewFlexControl(core common.Core) common.Component {
	core.SetSize(common.Size{
		Width:  375,
		Height: 200,
	})
	return &FlexControl{core: core}
}
