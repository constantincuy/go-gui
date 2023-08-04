package component

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type Rect struct {
	core          Core
	updateCounter int
}

func (box *Rect) Core() *Core {
	return &box.core
}

func (box *Rect) Mount() {
	box.Core().ApplyStyle("rect")
	box.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		vector.DrawFilledRect(screen, float32(bounds.Min.X), float32(bounds.Min.Y), float32(box.core.GetSize().Width), float32(box.core.GetSize().Height), box.getBackgroundOrDefault(), false)
	})
}

func (box *Rect) Destroy() {}

func (box *Rect) SetColor(c color.Color) {
	box.core.ForceFrameRedraw()
	r, g, b, _ := c.RGBA()
	style := box.Core().Style()
	(*style)["background"] = theme.Property{
		Name:  "background",
		Value: fmt.Sprintf("#%02x%02x%02x", r, g, b),
	}
}

func (box *Rect) GetColor() color.Color {
	return box.getBackgroundOrDefault()
}

func (box *Rect) getBackgroundOrDefault() color.RGBA {
	prop, exists := box.Core().style["background"]
	if exists {
		c, e := prop.AsColor()
		if e == nil {
			return c
		}
	}

	return color.RGBA{
		R: 0xff,
		G: 0xff,
		B: 0xff,
		A: 0xff,
	}
}

func (box *Rect) Update() {}

func NewRect(core Core) Component {
	return &Rect{core: core}
}
