package components

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/event"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/slices"
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
	c.button = c.Core().AddChild(component.NewButton).(*component.Button)
	c.counter = 0
	c.setCurrentCount()

	c.Core().Events().On(func(e event.Event) {
		switch e := e.(type) {
		case event.MouseClickEvent:
			if slices.Contains(e.Button, ebiten.MouseButtonLeft) {
				c.counter++
				c.setCurrentCount()
			}
		}
	})
}

func (c *Counter) setCurrentCount() {
	c.button.SetText(fmt.Sprintf("Clicked %d times", c.counter))
}

func (c *Counter) Update() {
	//Handle update logic target is 60 update calls per second
}

func (c *Counter) Destroy() {
	//Clean up of resources if needed
}

// NewCounter Factory for `AddChild` method `c.Core().AddChild(NewCounter)`
func NewCounter(core component.Core) component.Component {
	core.SetSize(common.Size{
		Width:  120,
		Height: 35,
	})
	return &Counter{core: core}
}
