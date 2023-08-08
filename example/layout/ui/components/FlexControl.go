package components

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
)

type LayoutListener func(justify, align component.FlexPosition)

type FlexControl struct {
	core     component.Core
	listener func(justify, align component.FlexPosition)
}

func (f *FlexControl) Core() *component.Core {
	return &f.core
}

func (f *FlexControl) Mount() {
	f.Core().SetDisplayType(component.FlexCentered().UseGap(15).UseDirection(component.FlexColumn))
	headline := f.Core().AddChild(component.NewText).(*component.Text)
	headline.SetFontSize(30)
	headline.SetText("Flex Layout")

	grid := f.Core().AddChild(component.NewGridContainer(3, 5)).(*component.Container)
	size := grid.Core().GetSize()
	size.Width = 375
	grid.Core().SetSize(size)
	options := []component.FlexPosition{component.FlexStart, component.FlexCenter, component.FlexEnd}
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

func NewFlexControl(core component.Core) component.Component {
	core.SetSize(common.Size{
		Width:  375,
		Height: 200,
	})
	return &FlexControl{core: core}
}
