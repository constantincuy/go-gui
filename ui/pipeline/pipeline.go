package pipeline

import (
	"github.com/constantincuy/go-gui/ui/window"
	"github.com/hajimehoshi/ebiten/v2"
)

type Pipeline interface {
	Render(screen *ebiten.Image, window window.Window)
}
