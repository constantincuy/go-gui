package component

import (
	"github.com/constantincuy/go-gui/ui/common"
)

type Button struct {
	core       Core
	background *Rect
	text       *Text
	counter    int
}

func (b *Button) Core() *Core {
	return &b.core
}

func (b *Button) Mount() {
	defaultText := "Button"
	b.Core().ApplyStyle("button")
	size := b.Core().GetSize()
	b.background = b.Core().AddChild(NewRect).(*Rect)
	b.background.Core().SetDisplayType(FlexLayoutCentered())
	b.text = b.background.Core().AddChild(NewText).(*Text)
	b.background.Core().SetSize(size)
	b.text.SetText(defaultText)
	b.text.Core().SetZ(1)

	b.background.Core().ApplyStyle("button>body")
	b.text.Core().ApplyStyle("button>label")
}

func (b *Button) SetText(text string) {
	b.text.SetText(text)
}

func (b *Button) Text() string {
	return b.text.Text()
}

func (b *Button) Update() {
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
