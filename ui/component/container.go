package component

import "github.com/constantincuy/go-gui/ui/common"

type Container struct {
	core    common.Core
	layout  common.State[common.LayoutOptions]
	initial common.LayoutOptions
}

func (c *Container) Core() *common.Core {
	return &c.core
}

func (c *Container) Mount() {
	c.layout = common.NewState(c.initial)

	c.layout.OnChange(func(new common.LayoutOptions) {
		c.Core().SetDisplayType(new)
	})
}

func (c *Container) Add(factory common.Factory) common.Component {
	return c.Core().AddChild(factory)
}

func (c *Container) SetLayout(options common.LayoutOptions) {
	c.layout.SetState(options)
}

func (c *Container) Update() {}

func (c *Container) Destroy() {}

func NewFlexContainer(initial common.FlexLayout) common.Factory {
	return func(core common.Core) common.Component {
		return &Container{initial: initial}
	}
}

func NewGridContainer(cols int, gap int) common.Factory {
	return func(core common.Core) common.Component {
		gridOp := common.NewGridLayout().(common.GridLayout)
		return &Container{initial: gridOp.UseColumns(cols).UseGap(gap)}
	}
}
