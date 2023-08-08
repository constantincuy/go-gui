package window

import (
	"github.com/constantincuy/go-gui/ui/common"
	"image/color"
)

type Window interface {
	GetTitle() string
	GetSize() common.Size
	SetBackground(c color.RGBA)
	GetBackground() color.RGBA
	GetView() *common.Component
	Layout(outsideWidth, outsideHeight int)
}
