package component

import (
	"github.com/constantincuy/go-gui/ui/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type Rect struct {
	core          Core
	updateCounter int
	border        int
	borderColor   color.RGBA
	background    color.RGBA
}

func (box *Rect) Core() *Core {
	return &box.core
}

func (box *Rect) Mount() {
	box.Core().ApplyStyle("rect")
	box.Core().OnStyleChange(func(property theme.Property) {
		switch property.Name {
		case "border":
			px, _ := property.AsPX()
			box.border = px
		case "border-color":
			bc, _ := property.AsColor()
			box.borderColor = bc
		case "background":
			bg, _ := property.AsColor()
			box.background = bg
		}
	})
	box.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		if box.border > 0 {
			vector.DrawFilledRect(screen, float32(bounds.Min.X), float32(bounds.Min.Y), float32(box.core.GetSize().Width), float32(box.core.GetSize().Height), box.borderColor, false)
		}
		vector.DrawFilledRect(screen, float32(bounds.Min.X+box.border), float32(bounds.Min.Y+box.border), float32(box.core.GetSize().Width-(box.border*2)), float32(box.core.GetSize().Height-(box.border*2)), box.background, false)
	})
}

func (box *Rect) Destroy() {}

func (box *Rect) SetColor(c color.RGBA) {
	box.Core().ApplyColorProperty("background", c)
}

func (box *Rect) GetColor() color.RGBA {
	return box.background
}

func (box *Rect) SetBorder(b int) {
	box.Core().ApplyPixelProperty("border", b)
}

func (box *Rect) GetBorder() int {
	return box.border
}

func (box *Rect) SetBorderColor(c color.RGBA) {
	box.Core().ApplyColorProperty("border-color", c)
}

func (box *Rect) GetBorderColor() color.RGBA {
	return box.borderColor
}

func (box *Rect) Update() {}

func NewRect(core Core) Component {
	return &Rect{core: core}
}
