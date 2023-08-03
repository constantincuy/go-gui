package component

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gode/ui/common"
	"gode/ui/theme"
	"image"
	"image/color"
)

type Box struct {
	core          Core
	updateCounter int
}

func (box *Box) Core() *Core {
	return &box.core
}

func (box *Box) Mount() {
	box.Core().ApplyStyle("box")
	box.core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		vector.DrawFilledRect(screen, float32(bounds.Min.X), float32(bounds.Min.Y), float32(box.core.GetSize().Width), float32(box.core.GetSize().Height), box.getBackgroundOrDefault(), false)
	})
}

func (box *Box) Destroy() {}

func (box *Box) SetColor(c color.Color) {
	box.core.SetDirty(true)
	r, g, b, _ := c.RGBA()
	style := box.Core().Style()
	(*style)["background"] = theme.Property{
		Name:  "background",
		Value: fmt.Sprintf("#%02x%02x%02x", r, g, b),
	}
}

func (box *Box) GetColor() color.Color {
	return box.getBackgroundOrDefault()
}

func (box *Box) getBackgroundOrDefault() color.RGBA {
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

func (box *Box) Update() {}

func NewBox(position image.Point, size common.Size) Component {
	core := NewCore(size)
	core.SetPosition(position)
	return &Box{core: core}
}
