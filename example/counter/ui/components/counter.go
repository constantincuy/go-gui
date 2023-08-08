package components

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/component"
)

type Counter struct {
	core         common.Core
	counterState common.State[int]
	button       *component.Button
}

func (c *Counter) Core() *common.Core {
	return &c.core
}

func (c *Counter) Mount() {
	c.button = c.Core().AddChild(component.NewButton).(*component.Button)
	c.counterState = common.NewState(0)
	c.counterState.OnChange(c.setCurrentCount)

	c.button.OnClick(func() {
		c.counterState.SetState(c.counterState.Get() + 1)
	})
}

func (c *Counter) setCurrentCount(count int) {
	c.button.SetText(fmt.Sprintf("Clicked %d times", count))
}

func (c *Counter) Update() {
	//Handle update logic target is 60 update calls per second
}

func (c *Counter) Destroy() {
	//Clean up of resources if needed
}

// NewCounter Factory for `AddChild` method `c.Core().AddChild(NewCounter)`
func NewCounter(core common.Core) common.Component {
	core.SetSize(common.Size{
		Width:  120,
		Height: 35,
	})
	return &Counter{core: core}
}
