package component

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/font"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/tinne26/etxt/efixed"
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
	init          bool
}

func (t *Text) Mount() {
	t.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		textRenderer := t.prepareRenderer()
		textRenderer.SetTarget(screen)
		textRenderer.Draw(t.text, bounds.Min.X, bounds.Min.Y)
	})
}

func (t *Text) prepareRenderer() *etxt.Renderer {
	textRenderer := font.Manager.TextRenderer(t.font)
	if textRenderer != nil {
		textRenderer.SetSizePx(int(t.size))
		textRenderer.SetLineHeight(t.lineHeight)
	}

	return textRenderer
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
	t.text = text
	t.recalculateSize()
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

func (t *Text) Update() {
	if !t.init {
		textRenderer := font.Manager.TextRenderer(t.font)
		if textRenderer != nil {
			t.recalculateSize()
			t.init = true
		}
	}
}

func (t *Text) recalculateSize() {
	r := t.prepareRenderer()
	if r != nil {
		bounds := r.SelectionRect(t.text)
		t.core.SetSize(common.Size{
			Width:  int(efixed.ToFloat64(bounds.Width)),
			Height: int(efixed.ToFloat64(bounds.Height)),
		})
	}
}

func NewText(core Core) Component {
	col := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	text := &Text{core: core, color: col, size: 12, font: "Segoe UI", lineHeight: 12}
	return text
}
