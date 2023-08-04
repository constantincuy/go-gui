package views

import (
	"bytes"
	_ "embed"
	"github.com/constantincuy/go-gui/example/counter/ui/components"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var (
	//go:embed gopher.png
	GopherPNG []byte
)

type MainView struct {
	core   component.Core
	gopher *ebiten.Image
}

func (view *MainView) Core() *component.Core {
	return &view.core
}

func (view *MainView) Mount() {
	decoded, _, err := image.Decode(bytes.NewReader(GopherPNG))
	if err != nil {
		log.Fatal(err)
	}
	view.gopher = ebiten.NewImageFromImage(decoded)

	view.Core().SetDisplayType(component.FlexCentered().UseGap(5).UseDirection(component.FlexRow))
	img := view.Core().AddChild(component.NewImage).(*component.Image)
	img.SetImage(view.gopher)
	view.Core().AddChild(components.NewCounter)
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {
	//Dispose the image when the component is destroyed to prevent memory leaks
	view.gopher.Dispose()
}

func NewMainView(core component.Core) component.Component {
	return &MainView{core: core}
}
