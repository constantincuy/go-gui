package window

import (
	"gode/ui/common"
	"gode/ui/component"
	"image/color"
)

type Window interface {
	GetTitle() string
	GetSize() common.Size
	SetBackground(c color.RGBA)
	GetBackground() color.RGBA
	GetView() *component.Component
	Layout(outsideWidth, outsideHeight int)
}
