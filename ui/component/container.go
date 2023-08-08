package component

type Container struct {
	core    Core
	layout  State[LayoutOptions]
	initial LayoutOptions
}

func (c *Container) Core() *Core {
	return &c.core
}

func (c *Container) Mount() {
	c.layout = NewState(c.initial)

	c.layout.OnChange(func(new LayoutOptions) {
		c.Core().SetDisplayType(new)
	})
}

func (c *Container) Add(factory Factory) Component {
	return c.Core().AddChild(factory)
}

func (c *Container) SetLayout(options LayoutOptions) {
	c.layout.SetState(options)
}

func (c *Container) Update() {}

func (c *Container) Destroy() {}

func NewFlexContainer(initial FlexLayout) Factory {
	return func(core Core) Component {
		return &Container{initial: initial}
	}
}

func NewGridContainer(cols int, gap int) Factory {
	return func(core Core) Component {
		gridOp := NewGridLayout().(GridLayout)
		return &Container{initial: gridOp.UseColumns(cols).UseGap(gap)}
	}
}
