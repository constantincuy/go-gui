package component

import (
	"gode/ui/common"
	"image"
)

type Button struct {
	core       Core
	background *Box
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
	box := NewBox(image.Point{
		X: 0,
		Y: 0,
	}, size)
	text := NewText(defaultText)
	b.core.AddChild(&text)
	b.core.AddChild(&box)
	b.background = box.(*Box)
	text.Core().SetZ(1)
	//text.Core().CenterIn(b.background.Core())
	b.text = text.(*Text)

	box.Core().ApplyStyle("button>body")
	text.Core().ApplyStyle("button>label")
}

func (b *Button) SetText(text string) {
	b.text.SetText(text)
	//b.text.Core().CenterIn(b.background.Core())
}

func (b *Button) GetText() string {
	return b.text.GetText()
}

func (b *Button) Update() {
}

func (b *Button) Destroy() {
}

func NewButton() Component {
	return &Button{core: NewCore(common.Size{
		Width:  120,
		Height: 35,
	})}
}
