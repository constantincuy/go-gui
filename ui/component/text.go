package component

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/font"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"image/color"
)

type Text struct {
	core          Core
	updateCounter int
	color         color.Color
	text          string
	font          string
	size          float64
	lineHeight    float64
}

func (t *Text) Mount() {
	t.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		ff, _ := font.Manager.GetFontFace(t.font, t.size, t.lineHeight)
		text.Draw(screen, t.text, ff, bounds.Min.X-int(t.size), bounds.Min.Y+3, t.color)
	})
}

func (t *Text) Destroy() {}

func (t *Text) Core() *Core {
	return &t.core
}

func (t *Text) SetColor(c color.Color) {
	t.Core().ForceFrameRedraw()
	t.color = c
}

func (t *Text) Color() color.Color {
	return t.color
}

func (t *Text) SetText(text string) {
	t.core.SetSize(common.Size{
		Width:  len(text) * 6,
		Height: int(t.lineHeight),
	})
	t.text = text
	t.core.ForceFrameRedraw()
}

func (t *Text) Text() string {
	return t.text
}

func (t *Text) SetFont(font string) {
	t.font = font
	t.recalculateSize()
	t.core.ForceFrameRedraw()
}

func (t *Text) Font() string {
	return t.font
}

func (t *Text) SetFontSize(size float64) {
	t.size = size
	t.recalculateSize()
	t.core.ForceFrameRedraw()
}

func (t *Text) FontSize() float64 {
	return t.size
}

func (t *Text) SetLineHeight(size float64) {
	t.lineHeight = size
	t.recalculateSize()
	t.core.ForceFrameRedraw()
}

func (t *Text) LineHeight() float64 {
	return t.lineHeight
}

func (t *Text) Update() {}

func (t *Text) recalculateSize() {
	t.core.SetSize(common.Size{
		Width:  len(t.text) * 6, // TODO: Calculate width based on glyphs
		Height: int(t.lineHeight),
	})
}

func NewText(core Core) Component {
	core.SetSize(common.Size{
		Width:  6,
		Height: 16,
	})
	col := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	return &Text{core: core, color: col, size: 14, font: "Segoe-UI", lineHeight: 14}
}
