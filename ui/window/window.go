package window

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
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
