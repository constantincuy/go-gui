package window

import (
	"gode/ui/common"
	"gode/ui/component"
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
}

func NewDefaultWindow(c component.Component) Window {
	c.Mount()
	return &DefaultWindow{view: c, color: color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}}
}
