package component

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/event"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/slices"
)

type Button struct {
	core          Core
	background    *Rect
	text          *Text
	counter       int
	clicked       bool
	clickListener func()
}

func (b *Button) Core() *Core {
	return &b.core
}

func (b *Button) Mount() {
	defaultText := "Button"
	b.Core().ApplyStyle("button")
	size := b.Core().GetSize()
	b.background = b.Core().AddChild(NewRect).(*Rect)
	b.background.Core().SetDisplayType(FlexCentered())
	b.text = b.background.Core().AddChild(NewText).(*Text)
	b.background.Core().SetSize(size)
	b.text.SetText(defaultText)
	b.text.Core().SetZ(1)

	b.background.Core().ApplyStyle("button>body")
	b.text.Core().ApplyStyle("button>label")

	b.Core().Events().On(func(e event.Event) {
		switch e := e.(type) {
		case event.MouseClickEvent:
			if slices.Contains(e.Button, ebiten.MouseButtonLeft) {
				b.background.Core().ApplyStyle("button>body:active")
				b.clicked = true
				if b.clickListener != nil {
					b.clickListener()
				}
			}
		}
	})
}

func (b *Button) OnClick(clickFn func()) {
	b.clickListener = clickFn
}

func (b *Button) SetText(text string) {
	b.text.SetText(text)
}

func (b *Button) Text() string {
	return b.text.Text()
}

func (b *Button) Update() {
	if b.clicked {
		if b.counter == 5 {
			b.counter = 0
			b.clicked = false
			b.background.Core().ApplyStyle("button>body")
		}
		b.counter++
	}
}

func (b *Button) Destroy() {
}

func NewButton(core Core) Component {
	core.SetSize(common.Size{
		Width:  120,
		Height: 35,
	})
	return &Button{core: core}
}
