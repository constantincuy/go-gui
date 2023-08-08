package component

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type Rect struct {
	core          common.Core
	updateCounter int
	border        int
	borderColor   color.RGBA
	background    color.RGBA
}

func (rect *Rect) Core() *common.Core {
	return &rect.core
}

func (rect *Rect) Mount() {
	rect.Core().ApplyStyle("rect")
	rect.Core().OnStyleChange(func(property theme.Property) {
		switch property.Name {
		case "border":
			px, _ := property.AsPX()
			rect.border = px
		case "border-color":
			bc, _ := property.AsColor()
			rect.borderColor = bc
		case "background":
			bg, _ := property.AsColor()
			rect.background = bg
		}
	})
	rect.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		if rect.border > 0 {
			vector.DrawFilledRect(screen, float32(bounds.Min.X), float32(bounds.Min.Y), float32(rect.core.GetSize().Width), float32(rect.core.GetSize().Height), rect.borderColor, false)
		}
		vector.DrawFilledRect(screen, float32(bounds.Min.X+rect.border), float32(bounds.Min.Y+rect.border), float32(rect.core.GetSize().Width-(rect.border*2)), float32(rect.core.GetSize().Height-(rect.border*2)), rect.background, false)
	})
}

func (rect *Rect) Destroy() {}

func (rect *Rect) SetColor(c color.RGBA) {
	rect.Core().ApplyColorProperty("background", c)
}

func (rect *Rect) GetColor() color.RGBA {
	return rect.background
}

func (rect *Rect) SetBorder(b int) {
	rect.Core().ApplyPixelProperty("border", b)
}

func (rect *Rect) GetBorder() int {
	return rect.border
}

func (rect *Rect) SetBorderColor(c color.RGBA) {
	rect.Core().ApplyColorProperty("border-color", c)
}

func (rect *Rect) GetBorderColor() color.RGBA {
	return rect.borderColor
}

func (rect *Rect) Update() {}

func NewRect(core common.Core) common.Component {
	return &Rect{core: core}
}
