package component

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Image struct {
	core  common.Core
	image *ebiten.Image
}

func (i *Image) Core() *common.Core {
	return &i.core
}

func (i *Image) Mount() {
	i.Core().OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		if i.image != nil {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
			screen.DrawImage(i.image, op)
		}
	})
}

func (i *Image) Update() {
}

func (i *Image) Destroy() {
}

func (i *Image) SetImage(image *ebiten.Image) {
	if image != nil {
		i.Core().SetSize(common.SizeFromBounds(image.Bounds()))
		i.image = image
		i.Core().ForceFrameRedraw()
	}
}

func NewImage(core common.Core) common.Component {
	return &Image{core: core}
}
