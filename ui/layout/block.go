package layout

import "github.com/constantincuy/go-gui/ui/component"

// BlockLayout This layout does not manipulate any positions or sizes and just forwards
// the next child components for the layout phase
type BlockLayout struct{}

func (l BlockLayout) ProcessLayout(comp component.Component) []*component.Component {
	return comp.Core().Children()
}
