package pipeline

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gode/ui/window"
)

type Pipeline interface {
	Render(screen *ebiten.Image, window window.Window)
}
