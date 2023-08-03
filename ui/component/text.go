package component

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
)

type Text struct {
	core          Core
	updateCounter int
	color         color.Color
	text          string
}

func (text *Text) Mount() {
	text.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		ebitenutil.DebugPrintAt(screen, text.text, bounds.Min.X, bounds.Min.Y)
	})
}

func (text *Text) Destroy() {}

func (text *Text) Core() *Core {
	return &text.core
}

func (text *Text) SetColor(c color.Color) {
	text.core.SetDirty(true)
	text.color = c
}

func (text *Text) GetColor() color.Color {
	return text.color
}

func (text *Text) SetText(t string) {
	text.core.SetSize(common.Size{
		Width:  len(t) * 6,
		Height: 16,
	})
	text.text = t
	text.core.ForceFrameRedraw()
}

func (text *Text) GetText() string {
	return text.text
}

func (text *Text) Update() {}

func NewText(core Core) Component {
	core.SetSize(common.Size{
		Width:  6,
		Height: 16,
	})
	col := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	return &Text{core: core, color: col}
}
