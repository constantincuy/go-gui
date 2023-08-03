package components

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/event"
)

type Counter struct {
	core    component.Core
	counter int
	button  *component.Button
}

func (c *Counter) Core() *component.Core {
	return &c.core
}

func (c *Counter) Mount() {
	button := component.NewButton()
	c.Core().AddChild(&button)
	c.button = button.(*component.Button)
	c.counter = 0
	c.setCurrentCount()

	c.Core().Events().On(func(e event.Event) {
		switch e.(type) {
		case event.MouseClickEvent:
			c.counter++
			c.setCurrentCount()
		}
	})
}

func (c *Counter) setCurrentCount() {
	c.button.SetText(fmt.Sprintln("Clicked ", c.counter))
}

func (c *Counter) Update() {}

func (c *Counter) Destroy() {}

func NewCounter() component.Component {
	return &Counter{core: component.NewCore(common.Size{Width: 120, Height: 35})}
}
