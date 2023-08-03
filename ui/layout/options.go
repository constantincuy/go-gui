package layout

import "github.com/constantincuy/go-gui/ui/component"

type Options interface {
	ProcessLayout(comp *component.Component) []*component.Component
}
