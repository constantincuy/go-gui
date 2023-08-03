package window

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
	"image/color"
)

type DefaultWindow struct {
	view  component.Component
	size  common.Size
	color color.RGBA
}

func (win *DefaultWindow) SetBackground(c color.RGBA) {
	win.color = c
}
func (win *DefaultWindow) GetBackground() color.RGBA {
	return win.color
}

func (win *DefaultWindow) GetTitle() string {
	return "Default"
}

func (win *DefaultWindow) GetSize() common.Size {
	return win.size
}

func (win *DefaultWindow) GetView() *component.Component {
	return &win.view
}

func (win *DefaultWindow) Layout(outsideWidth, outsideHeight int) {
	win.size = common.Size{
		Width:  outsideWidth,
		Height: outsideHeight,
	}
	win.view.Core().SetSize(win.size)
}

func NewDefaultWindow(viewFactory func(core component.Core) component.Component) Window {
	view := viewFactory(component.NewCore())
	view.Mount()
	return &DefaultWindow{view: view, color: color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}}
}
