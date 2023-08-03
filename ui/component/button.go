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
	b.core.ApplyStyle("button")
	size := b.core.GetSize()
	b.text = b.core.AddChild(NewText).(*Text)
	b.background = b.core.AddChild(NewRect).(*Rect)
	b.background.Core().SetSize(size)
	b.text.SetText(defaultText)
	b.text.Core().SetZ(1)

	b.background.Core().ApplyStyle("button>body")
	b.text.Core().ApplyStyle("button>label")
}

func (b *Button) SetText(text string) {
	b.text.SetText(text)
	//b.text.Core().CenterIn(b.background.Core())
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
