package views

import (
	"bytes"
	_ "embed"
	"github.com/constantincuy/go-gui/example/counter/ui/components"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var (
	//go:embed go-logo.png
	GoLogoPNG []byte
)

type MainView struct {
	core common.Core
	logo *ebiten.Image
}

func (view *MainView) Core() *common.Core {
	return &view.core
}

func (view *MainView) Mount() {
	decoded, _, err := image.Decode(bytes.NewReader(GoLogoPNG))
	if err != nil {
		log.Fatal(err)
	}
	view.logo = ebiten.NewImageFromImage(decoded)

	view.Core().SetDisplayType(common.FlexCentered().UseGap(60).UseDirection(common.FlexColumn))
	headline := view.Core().AddChild(component.NewText).(*component.Text)
	headline.SetFontSize(30)
	headline.SetLineHeight(30)
	headline.SetText("Welcome to Go-Gui")
	img := view.Core().AddChild(component.NewImage).(*component.Image)
	img.SetImage(view.logo)
	view.Core().AddChild(components.NewCounter)
}

func (view *MainView) Update() {}

func (view *MainView) Destroy() {}

func NewMainView(core common.Core) common.Component {
	return &MainView{core: core}
}
