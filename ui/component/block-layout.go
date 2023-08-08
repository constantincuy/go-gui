package component

// BlockLayout This layout does not manipulate any positions or sizes and just forwards
// the next child components for the layout phase
type BlockLayout struct{}

func (l BlockLayout) ProcessLayout(comp Component) []*Component {
	return comp.Core().Children()
}

func NewBlockLayout() LayoutOptions {
	return BlockLayout{}
}
